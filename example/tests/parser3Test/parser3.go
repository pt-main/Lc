package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/parsing/stringParsing/parser3"
)

func main() {
	fmt.Println("Lc version -", lc.Version)

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
	jsonData, err2 := json.MarshalIndent(cleanRoot, "", "  ")
	if err2 != nil {
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

/*
Lc version - 1.5.1
Raw expr: "3 + 5 * (2 - 1)"

AST (as text):
[expr] Raw: "3+5*(2-1)"
  [expr] Raw: "3+5*(2-1)"
    [term] Raw: "3"
      [factor] Raw: "3"
        [NUMBER] Raw: "3"
    [PLUS] Raw: "+"
    [term] Raw: "5*(2-1)"
      [factor] Raw: "5"
        [NUMBER] Raw: "5"
      [STAR] Raw: "*"
      [factor] Raw: "(2-1)"
        [LPAREN] Raw: "("
        [expr] Raw: "2-1"
          [term] Raw: "2"
            [factor] Raw: "2"
              [NUMBER] Raw: "2"
          [MINUS] Raw: "-"
          [term] Raw: "1"
            [factor] Raw: "1"
              [NUMBER] Raw: "1"
        [RPAREN] Raw: ")"



AST (JSON):
{
  "Raw": "3+5*(2-1)",
  "Switch": "expr",
  "Metadata": {
    "children": [
      {
        "Raw": "3+5*(2-1)",
        "Switch": "expr",
        "Metadata": {
          "children": [
            {
              "Raw": "3",
              "Switch": "term",
              "Metadata": {
                "children": [
                  {
                    "Raw": "3",
                    "Switch": "factor",
                    "Metadata": {
                      "children": [
                        {
                          "Raw": "3",
                          "Switch": "NUMBER",
                          "Metadata": {
                            "__raw": "3",
                            "__value": "3"
                          }
                        }
                      ]
                    }
                  }
                ]
              }
            },
            {
              "Raw": "+",
              "Switch": "PLUS",
              "Metadata": {
                "__raw": "+",
                "__value": "+"
              }
            },
            {
              "Raw": "5*(2-1)",
              "Switch": "term",
              "Metadata": {
                "children": [
                  {
                    "Raw": "5",
                    "Switch": "factor",
                    "Metadata": {
                      "children": [
                        {
                          "Raw": "5",
                          "Switch": "NUMBER",
                          "Metadata": {
                            "__raw": "5",
                            "__value": "5"
                          }
                        }
                      ]
                    }
                  },
                  {
                    "Raw": "*",
                    "Switch": "STAR",
                    "Metadata": {
                      "__raw": "*",
                      "__value": "*"
                    }
                  },
                  {
                    "Raw": "(2-1)",
                    "Switch": "factor",
                    "Metadata": {
                      "children": [
                        {
                          "Raw": "(",
                          "Switch": "LPAREN",
                          "Metadata": {
                            "__raw": "(",
                            "__value": "("
                          }
                        },
                        {
                          "Raw": "2-1",
                          "Switch": "expr",
                          "Metadata": {
                            "children": [
                              {
                                "Raw": "2",
                                "Switch": "term",
                                "Metadata": {
                                  "children": [
                                    {
                                      "Raw": "2",
                                      "Switch": "factor",
                                      "Metadata": {
                                        "children": [
                                          {
                                            "Raw": "2",
                                            "Switch": "NUMBER",
                                            "Metadata": {
                                              "__raw": "2",
                                              "__value": "2"
                                            }
                                          }
                                        ]
                                      }
                                    }
                                  ]
                                }
                              },
                              {
                                "Raw": "-",
                                "Switch": "MINUS",
                                "Metadata": {
                                  "__raw": "-",
                                  "__value": "-"
                                }
                              },
                              {
                                "Raw": "1",
                                "Switch": "term",
                                "Metadata": {
                                  "children": [
                                    {
                                      "Raw": "1",
                                      "Switch": "factor",
                                      "Metadata": {
                                        "children": [
                                          {
                                            "Raw": "1",
                                            "Switch": "NUMBER",
                                            "Metadata": {
                                              "__raw": "1",
                                              "__value": "1"
                                            }
                                          }
                                        ]
                                      }
                                    }
                                  ]
                                }
                              }
                            ]
                          }
                        },
                        {
                          "Raw": ")",
                          "Switch": "RPAREN",
                          "Metadata": {
                            "__raw": ")",
                            "__value": ")"
                          }
                        }
                      ]
                    }
                  }
                ]
              }
            }
          ]
        }
      }
    ]
  }
}
*/
