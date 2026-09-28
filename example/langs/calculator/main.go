package main

import (
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc/v2"
	enginepkg "github.com/pt-main/lc/v2/engine"
	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/parsing/stringParsing"
	"github.com/pt-main/lc/v2/parsing/stringParsing/parser3"
	"github.com/pt-main/lc/v2/public"
	"github.com/pt-main/lc/v2/tooling/astools"
)

func createLexer() *stringParsing.Lexer {
	rules := []stringParsing.LexerRule{
		{Type: "NUMBER", Pattern: regexp2.MustCompile(`\d+(\.\d+)?`, 0)},
		{Type: "PLUS", Pattern: regexp2.MustCompile(`\+`, 0)},
		{Type: "MINUS", Pattern: regexp2.MustCompile(`-`, 0)},
		{Type: "POW", Pattern: regexp2.MustCompile(`\*\*`, 0)},
		{Type: "MUL", Pattern: regexp2.MustCompile(`\*`, 0)},
		{Type: "DIV", Pattern: regexp2.MustCompile(`/`, 0)},
		{Type: "LPAREN", Pattern: regexp2.MustCompile(`\(`, 0)},
		{Type: "RPAREN", Pattern: regexp2.MustCompile(`\)`, 0)},
		{Type: "WHITESPACE", Pattern: regexp2.MustCompile(`\s+`, 0)},
	}
	config := &stringParsing.LexerConfig{
		UseBracketBalance: false,
	}
	return stringParsing.NewLexer(rules, config)
}

func createGrammar() parser3.Grammar {
	return parser3.Grammar{
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
											parser3.TokenExpr{TokenType: "PLUS"},
											parser3.TokenExpr{TokenType: "MINUS"},
										},
									},
									parser3.NamedExpr{RuleName: "term"},
								},
							},
							Min: 0,
						},
					},
				},
			},
		},
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
											parser3.TokenExpr{TokenType: "POW"},
										},
									},
									parser3.NamedExpr{RuleName: "factor"},
								},
							},
							Min: 0,
						},
					},
				},
			},
		},
		"factor": {
			Name: "factor",
			Expr: parser3.NodeExpr{
				NodeType: "factor",
				Expr: parser3.ChoiceExpr{
					Alternatives: []parser3.Expr{
						parser3.TokenExpr{TokenType: "NUMBER"},
						parser3.SequenceExpr{
							Exprs: []parser3.Expr{
								parser3.TokenExpr{TokenType: "LPAREN"},
								parser3.NamedExpr{RuleName: "expr"},
								parser3.TokenExpr{TokenType: "RPAREN"},
							},
						},
					},
				},
			},
		},
	}
}

func evalExpr(node *stringParsing.ParsedNode) (float64, error) {
	switch node.Switch {
	case "expr":
		children := astools.GetChildren(node)
		if len(children) == 0 {
			return 0, errors.New("empty expr")
		}
		val, err := evalExpr(&children[0])
		if err != nil {
			return 0, err
		}
		for i := 1; i+1 < len(children); i += 2 {
			termVal, err := evalExpr(&children[i+1])
			if err != nil {
				return 0, err
			}
			switch children[i].Switch {
			case "PLUS":
				val += termVal
			case "MINUS":
				val -= termVal
			}
		}
		return val, nil
	case "term":
		children := astools.GetChildren(node)
		if len(children) == 0 {
			return 0, errors.New("empty term")
		}
		val, err := evalExpr(&children[0])
		if err != nil {
			return 0, err
		}
		for i := 1; i+1 < len(children); i += 2 {
			factorVal, err := evalExpr(&children[i+1])
			if err != nil {
				return 0, err
			}
			switch children[i].Switch {
			case "MUL":
				val *= factorVal
			case "DIV":
				if factorVal == 0 {
					return 0, errors.New("division by zero")
				}
				val /= factorVal
			case "POW":
				val = math.Pow(val, factorVal)
			}
		}
		return val, nil
	case "factor":
		children := astools.GetChildren(node)
		if len(children) == 0 {
			return 0, errors.New("empty factor")
		}
		child := &children[0]
		if child.Switch == "NUMBER" {
			return strconv.ParseFloat(child.Raw, 64)
		}
		if child.Switch == "LPAREN" && len(children) >= 3 {
			return evalExpr(&children[1])
		}
		return 0, errors.New("unknown factor: " + child.Switch)
	default:
		return 0, errors.New("unknown node type: " + node.Switch)
	}
}

func buildEngine() *lc.EngineUniversal {
	adapter := &parser3.Adapter{
		Parser: parser3.NewParser(createLexer(), createGrammar(), "expr", []string{"WHITESPACE"}),
	}

	engine, err := lc.NewEngineBuilder(public.StringEngineType, public.StringResType).
		WithPipeline([]string{"main"}).
		WithStringParser(adapter).
		WithDefaultEvents(true).
		WithContext(context.Background()).
		Build()
	if err != nil {
		panic(err)
	}

	err = engine.NewCommandString("expr", func(se enginepkg.StringEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
		val, err := evalExpr(node)
		if err != nil {
			return core.Wrap("EXPR", err, "eval error")
		}
		return se.GetUep().Generator.AddString(strconv.FormatFloat(val, 'f', -1, 64), "main")
	}, "evaluate expression")
	if err != nil {
		panic(err)
	}

	return engine
}

func main() {
	fmt.Println("Lc version -", lc.Version)

	if len(os.Args) < 2 {
		fmt.Println("Usage: calc <expression>")
		fmt.Println("Example: calc '(2 ** 3) + 4'")
		os.Exit(1)
	}

	engine := buildEngine()
	if err := engine.ProcessString(os.Args[1]); err != nil {
		if perr, ok := parser3.AsParseError(err); ok {
			fmt.Println("Parse error:\n", parser3.FormatErrorPretty(perr))
		} else {
			fmt.Println("Eval error:\n", err.Format())
		}
		os.Exit(1)
	}

	uep, err := engine.GetUEP()
	if err != nil {
		panic(err)
	}
	out, err := core.GetStringRes(uep.Generator, "\n")
	if err != nil {
		panic(err)
	}
	fmt.Println("Result:", out)
}
