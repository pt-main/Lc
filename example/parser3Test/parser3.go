package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/parsing/stringParsing/parser3"
)

func main() {

	rules := []stringParsing.LexerRule{
		{Type: "NUMBER", Pattern: regexp2.MustCompile(`\d+`, 0)},
		{Type: "PLUS", Pattern: regexp2.MustCompile(`\+`, 0)},
		{Type: "MINUS", Pattern: regexp2.MustCompile(`-`, 0)},
		{Type: "STAR", Pattern: regexp2.MustCompile(`\*`, 0)},
		{Type: "SLASH", Pattern: regexp2.MustCompile(`/`, 0)},
		{Type: "LPAREN", Pattern: regexp2.MustCompile(`\(`, 0)},
		{Type: "RPAREN", Pattern: regexp2.MustCompile(`\)`, 0)},
		{Type: "WHITESPACE", Pattern: regexp2.MustCompile(`\s+`, 0)},
	}
	lexer := stringParsing.NewLexer(rules, &stringParsing.LexerConfig{
		UseBracketBalance: false,
	})

	grammar := parser3.Grammar{
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
											parser3.TokenExpr{TokenType: "STAR"},
											parser3.TokenExpr{TokenType: "SLASH"},
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

	parser := parser3.NewParser(lexer, grammar, "expr", []string{"WHITESPACE"})

	expression := "3 + 5 * (2 - 1)"
	nodes, err := parser.Parse(expression)
	if err != nil {
		log.Fatalf("ParsingError: %v", err)
	}
	if len(nodes) == 0 {
		fmt.Println("Has no nodes")
		return
	}
	root := nodes[0]

	fmt.Printf("Raw expr: %q\n", expression)
	fmt.Println("\nAST (as text):")
	printAST(root, 0)

	cleanRoot := cleanNode(root)
	jsonData, err := json.MarshalIndent(cleanRoot, "", "  ")
	if err != nil {
		log.Fatalf("Json error: %v", err)
	}
	fmt.Println("\nAST (JSON):")
	fmt.Println(string(jsonData))
}

func printAST(node stringParsing.ParsedNode, indent int) {
	prefix := strings.Repeat("  ", indent)
	fmt.Printf("%s[%s] Raw: %q\n", prefix, node.Switch, node.Raw)
	if children, ok := node.Metadata["children"].([]stringParsing.ParsedNode); ok {
		for _, child := range children {
			printAST(child, indent+1)
		}
	}
}

func cleanNode(node stringParsing.ParsedNode) stringParsing.ParsedNode {
	newMeta := make(map[string]interface{})
	for k, v := range node.Metadata {
		if k == "__prev" || k == "__next" {
			continue
		}
		if k == "children" {
			if children, ok := v.([]stringParsing.ParsedNode); ok {
				newChildren := make([]stringParsing.ParsedNode, len(children))
				for i, child := range children {
					newChildren[i] = cleanNode(child)
				}
				newMeta[k] = newChildren
				continue
			}
		}
		newMeta[k] = v
	}
	return stringParsing.ParsedNode{
		Raw:      node.Raw,
		Switch:   node.Switch,
		Metadata: newMeta,
	}
}
