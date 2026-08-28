package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc"
	enginepkg "github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/parsing/stringParsing/parser3"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/tooling/astools"
)

func main() {
	lexerRules := []stringParsing.LexerRule{
		{Type: "NUMBER", Pattern: regexp2.MustCompile(`\d+`, 0)},
		{Type: "IDENT", Pattern: regexp2.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`, 0)},
		{Type: "ASSIGN", Pattern: regexp2.MustCompile(`:=`, 0)},
		{Type: "PLUS", Pattern: regexp2.MustCompile(`\+`, 0)},
		{Type: "MINUS", Pattern: regexp2.MustCompile(`-`, 0)},
		{Type: "MUL", Pattern: regexp2.MustCompile(`\*`, 0)},
		{Type: "DIV", Pattern: regexp2.MustCompile(`/`, 0)},
		{Type: "LPAREN", Pattern: regexp2.MustCompile(`\(`, 0)},
		{Type: "RPAREN", Pattern: regexp2.MustCompile(`\)`, 0)},
		{Type: "WHITESPACE", Pattern: regexp2.MustCompile(`\s+`, 0)},
	}
	lexer := stringParsing.NewLexer(lexerRules, &stringParsing.LexerConfig{
		UseBracketBalance: false,
	})

	grammar := parser3.Grammar{
		"program": {
			Name: "program",
			Expr: parser3.SequenceExpr{
				Exprs: []parser3.Expr{
					parser3.RepeatExpr{Expr: parser3.NamedExpr{RuleName: "assign"}, Min: 0},
					parser3.NamedExpr{RuleName: "expr"}}}},
		"assign": {
			Name: "assign",
			Expr: parser3.NodeExpr{
				NodeType: "assign",
				Expr: parser3.SequenceExpr{
					Exprs: []parser3.Expr{
						parser3.TokenExpr{TokenType: "IDENT"},
						parser3.TokenExpr{TokenType: "ASSIGN"},
						parser3.NamedExpr{RuleName: "expr"}}}}},
		"expr": {
			Name: "expr",
			Expr: parser3.NodeExpr{
				NodeType: "expr",
				Expr: parser3.SequenceExpr{
					Exprs: []parser3.Expr{
						parser3.NamedExpr{RuleName: "term"},
						parser3.RepeatExpr{
							Expr: parser3.SequenceExpr{
								Exprs: []parser3.Expr{
									parser3.ChoiceExpr{
										Alternatives: []parser3.Expr{
											parser3.TokenExpr{TokenType: "PLUS"}, parser3.TokenExpr{TokenType: "MINUS"},
										},
									}, parser3.NamedExpr{RuleName: "term"},
								},
							}, Min: 0}}}}},
		"term": {
			Name: "term",
			Expr: parser3.NodeExpr{
				NodeType: "term",
				Expr: parser3.SequenceExpr{
					Exprs: []parser3.Expr{
						parser3.NamedExpr{RuleName: "factor"},
						parser3.RepeatExpr{
							Expr: parser3.SequenceExpr{
								Exprs: []parser3.Expr{
									parser3.ChoiceExpr{
										Alternatives: []parser3.Expr{
											parser3.TokenExpr{TokenType: "MUL"},
											parser3.TokenExpr{TokenType: "DIV"},
										},
									}, parser3.NamedExpr{RuleName: "factor"},
								},
							}, Min: 0}}}}},
		"factor": {
			Name: "factor",
			Expr: parser3.NodeExpr{
				NodeType: "factor",
				Expr: parser3.ChoiceExpr{
					Alternatives: []parser3.Expr{
						parser3.SequenceExpr{
							Exprs: []parser3.Expr{
								parser3.TokenExpr{TokenType: "MINUS"}, parser3.NamedExpr{RuleName: "factor"},
							}},
						parser3.TokenExpr{TokenType: "NUMBER"},
						parser3.TokenExpr{TokenType: "IDENT"},
						parser3.SequenceExpr{
							Exprs: []parser3.Expr{
								parser3.TokenExpr{TokenType: "LPAREN"},
								parser3.NamedExpr{RuleName: "expr"},
								parser3.TokenExpr{TokenType: "RPAREN"},
							}}}}}}}

	rawParser := parser3.NewParser(lexer, grammar, "program", []string{"WHITESPACE"})
	adapter := &parser3.Adapter{Parser: rawParser}

	engine, err := lc.NewEngineBuilder(public.StringEngineType, public.StringResType).
		WithPipeline([]string{"main"}).
		WithStringParser(adapter).
		WithDefaultEvents(true).
		WithContext(context.Background()).
		Build()
	if err != nil {
		panic(err)
	}

	err = engine.NewCommandString("assign", func(se enginepkg.StringEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
		children := astools.GetChildren(node)
		if len(children) < 3 {
			return core.Err("ASSIGN", "invalid assign node")
		}
		varName := children[0].Raw
		exprNode := &children[2]
		val, err := evalExpr(exprNode, se.GetUep().Scope)
		if err != nil {
			return core.Wrap("ASSIGN", err, "eval error")
		}
		se.GetUep().Scope[varName] = val
		return nil
	}, "assign variable")

	if err != nil {
		panic(err)
	}

	err = engine.NewCommandString("expr", func(se enginepkg.StringEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
		val, err := evalExpr(node, se.GetUep().Scope)
		if err != nil {
			return core.Wrap("EXPR", err, "eval error")
		}

		return se.GetUep().Generator.AddString(fmt.Sprintf("%d", val), "main")
	}, "final expression")

	if err != nil {
		panic(err)
	}

	if len(os.Args) != 2 {
		fmt.Println("Can't calculate: expr is not found")
		return
	}
	input := os.Args[1]

	err = engine.ProcessString(input)
	if err != nil {
		panic(err)
	}

	uep, _ := engine.GetUEP()
	out, err := core.GetStringRes(uep.Generator, "\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

func evalExpr(node *stringParsing.ParsedNode, scope core.ScopeType) (int, error) {
	switch node.Switch {
	case "expr":
		children := astools.GetChildren(node)
		if len(children) == 0 {
			return 0, fmt.Errorf("empty expr")
		}
		left, err := evalExpr(&children[0], scope)
		if err != nil {
			return 0, err
		}

		for i := 1; i < len(children); i += 2 {
			op := children[i].Switch
			right, err := evalExpr(&children[i+1], scope)
			if err != nil {
				return 0, err
			}
			switch op {
			case "PLUS":
				left += right
			case "MINUS":
				left -= right
			default:
				return 0, fmt.Errorf("unknown op: %s", op)
			}
		}
		return left, nil

	case "term":
		children := astools.GetChildren(node)
		if len(children) == 0 {
			return 0, fmt.Errorf("empty term")
		}
		left, err := evalExpr(&children[0], scope)
		if err != nil {
			return 0, err
		}
		for i := 1; i < len(children); i += 2 {
			op := children[i].Switch
			right, err := evalExpr(&children[i+1], scope)
			if err != nil {
				return 0, err
			}
			switch op {
			case "MUL":
				left *= right
			case "DIV":
				if right == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				left /= right
			default:
				return 0, fmt.Errorf("unknown op: %s", op)
			}
		}
		return left, nil

	case "factor":
		children := astools.GetChildren(node)
		if len(children) == 0 {
			return 0, fmt.Errorf("empty factor")
		}
		first := &children[0]
		if first.Switch == "MINUS" {
			if len(children) < 2 {
				return 0, fmt.Errorf("malformed unary minus")
			}
			val, err := evalExpr(&children[1], scope)
			if err != nil {
				return 0, err
			}
			return -val, nil
		}
		switch first.Switch {
		case "NUMBER":
			return strconv.Atoi(first.Raw)
		case "IDENT":
			val, ok := scope[first.Raw]
			if !ok {
				return 0, fmt.Errorf("undefined variable: %s", first.Raw)
			}
			iVal, ok := val.(int)
			if !ok {
				return 0, fmt.Errorf("variable %s is not int", first.Raw)
			}
			return iVal, nil
		case "LPAREN":
			if len(children) < 3 {
				return 0, fmt.Errorf("malformed parentheses")
			}
			return evalExpr(&children[1], scope)
		default:
			return 0, fmt.Errorf("unexpected factor: %s", first.Switch)
		}
	default:
		return 0, fmt.Errorf("unexpected node type: %s", node.Switch)
	}
}
