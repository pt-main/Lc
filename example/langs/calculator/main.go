package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/parsing/stringParsing/parser3"
	"github.com/pt-main/lc/tooling/astools"
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
		for i := 1; i < len(children); i += 2 {
			if i+1 >= len(children) {
				break
			}
			opNode := &children[i]
			termNode := &children[i+1]
			termVal, err := evalExpr(termNode)
			if err != nil {
				return 0, err
			}
			switch opNode.Switch {
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
		for i := 1; i < len(children); i += 2 {
			if i+1 >= len(children) {
				break
			}
			opNode := &children[i]
			factorNode := &children[i+1]
			factorVal, err := evalExpr(factorNode)
			if err != nil {
				return 0, err
			}
			switch opNode.Switch {
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
			numStr := child.Raw
			return strconv.ParseFloat(numStr, 64)
		}
		if child.Switch == "LPAREN" && len(children) >= 3 {
			return evalExpr(&children[1])
		}
		return 0, errors.New("unknown factor")
	default:
		return 0, errors.New("unknown node type: " + node.Switch)
	}
}

func main() {
	fmt.Println("Lc version -", lc.Version)

	if len(os.Args) < 2 {
		fmt.Println("Usage: calc <expression>")
		fmt.Println("Example: calc '(2 ** 3) + 4'")
		os.Exit(1)
	}
	expr := os.Args[1]
	lexer := createLexer()
	grammar := createGrammar()
	parser := parser3.NewParser(lexer, grammar, "expr", []string{"WHITESPACE"})
	parsed, err := parser.Parse(expr)
	if err != nil {
		fmt.Println("Parse error:\n", parser3.FormatErrorPretty(err))
		os.Exit(1)
	}
	if len(parsed) == 0 {
		fmt.Println("No nodes parsed")
		os.Exit(1)
	}
	result, err2 := evalExpr(&parsed[0])
	if err2 != nil {
		fmt.Println("Eval error:", err)
		os.Exit(1)
	}
	fmt.Printf("Result: %v\n", result)
}

/*
macbook@MacBook-Pro lc % go run ./example/calculator '(2+3)**4'
Lc version - 1.5.1
Result: 625
macbook@MacBook-Pro lc % go run ./example/calculator '2*3+4'
Lc version - 1.5.1
Result: 10
macbook@MacBook-Pro lc % go run ./example/calculator '2-3*4'
Lc version - 1.5.1
Result: -10
macbook@MacBook-Pro lc % go run ./example/calculator '(2-3*4)+(5*2-1)*2'
Lc version - 1.5.1
Result: 8
*/
