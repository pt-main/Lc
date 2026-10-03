package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/pt-main/lc/v2"
	"github.com/pt-main/lc/v2/parsing/stringParsing"
	"github.com/pt-main/lc/v2/parsing/stringParsing/parser3"
)

func lexer() (*stringParsing.Lexer, error) {
	rules := []stringParsing.LexerRule{
		{Type: "NUMBER", Pattern: `\d+(\.\d+)?`},
		{Type: "PLUS", Pattern: `\+`},
		{Type: "MINUS", Pattern: `-`},
		{Type: "STAR", Pattern: `\*`},
		{Type: "SLASH", Pattern: `/`},
		{Type: "LPAREN", Pattern: `\(`},
		{Type: "RPAREN", Pattern: `\)`},
		{Type: "WHITESPACE", Pattern: `\s+`},
	}
	return stringParsing.NewLexer(rules, &stringParsing.LexerConfig{UseBracketBalance: false})
}

func calcGrammar() parser3.Grammar {
	return parser3.Grammar{
		"expr": {Name: "expr", Expr: parser3.NodeExpr{NodeType: "expr", Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
			parser3.NamedExpr{RuleName: "term"},
			parser3.RepeatExpr{Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
				parser3.ChoiceExpr{Alternatives: []parser3.Expr{
					parser3.TokenExpr{TokenType: "PLUS"},
					parser3.TokenExpr{TokenType: "MINUS"},
				}},
				parser3.NamedExpr{RuleName: "term"},
			}}},
		}}}},
		"term": {Name: "term", Expr: parser3.NodeExpr{NodeType: "term", Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
			parser3.NamedExpr{RuleName: "factor"},
			parser3.RepeatExpr{Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
				parser3.ChoiceExpr{Alternatives: []parser3.Expr{
					parser3.TokenExpr{TokenType: "STAR"},
					parser3.TokenExpr{TokenType: "SLASH"},
				}},
				parser3.NamedExpr{RuleName: "factor"},
			}}},
		}}}},
		"factor": {Name: "factor", Expr: parser3.NodeExpr{NodeType: "factor", Expr: parser3.ChoiceExpr{Alternatives: []parser3.Expr{
			parser3.TokenExpr{TokenType: "NUMBER"},
			parser3.SequenceExpr{Exprs: []parser3.Expr{
				parser3.TokenExpr{TokenType: "LPAREN"},
				parser3.NamedExpr{RuleName: "expr"},
				parser3.TokenExpr{TokenType: "RPAREN"},
			}},
		}}}},
	}
}

func prattGrammar() parser3.Grammar {
	return parser3.Grammar{
		"start": {Name: "start", Expr: &parser3.PrattExpr{
			Atom: parser3.NodeExpr{
				NodeType: "atom",
				Expr: parser3.ChoiceExpr{Alternatives: []parser3.Expr{
					parser3.TokenExpr{TokenType: "NUMBER"},
					parser3.SequenceExpr{Exprs: []parser3.Expr{
						parser3.TokenExpr{TokenType: "LPAREN"},
						&parser3.PrattExpr{
							Atom:     parser3.TokenExpr{TokenType: "NUMBER"},
							Infixes:  prattInfixes(),
							Prefixes: prattPrefixes(),
						},
						parser3.TokenExpr{TokenType: "RPAREN"},
					}},
				}},
			},
			Prefixes: prattPrefixes(),
			Infixes:  prattInfixes(),
		}},
	}
}

func prattInfixes() map[string]parser3.InfixInfo {
	return map[string]parser3.InfixInfo{
		"PLUS":  {Precedence: 1, Assoc: parser3.LeftAssoc},
		"MINUS": {Precedence: 1, Assoc: parser3.LeftAssoc},
		"STAR":  {Precedence: 2, Assoc: parser3.LeftAssoc},
		"SLASH": {Precedence: 2, Assoc: parser3.LeftAssoc},
	}
}

func prattPrefixes() map[string]parser3.Expr {
	return map[string]parser3.Expr{
		"MINUS": parser3.TokenExpr{TokenType: "MINUS"},
	}
}

func ignore() []string { return []string{"WHITESPACE"} }

func main() {
	fmt.Println("Lc version -", lc.Version)

	expression := "3 + 5 * (2 - 1)"
	lex, lerr := lexer()
	if lerr != nil {
		fail(lerr)
		return
	}
	parser := parser3.NewParser(lex, calcGrammar(), "expr", ignore())

	nodes, err := parser.Parse(expression)
	if err != nil {
		fail(err)
	}
	root := nodes[0]

	fmt.Printf("Raw expr: %q\n\n", expression)
	fmt.Println("AST (as text):")
	printAST(root, 0)

	data, jerr := json.MarshalIndent(cleanNode(root), "", "  ")
	if jerr != nil {
		fmt.Fprintln(os.Stderr, "json error:", jerr)
		os.Exit(1)
	}
	fmt.Println("\nAST (JSON):")
	fmt.Println(string(data))

	precedenceDemo()
	errorDemo()
}

// precedenceDemo shows PrattExpr picking * over +, and unary minus binding
// tighter than any infix operator.
func precedenceDemo() {
	fmt.Println("\n=== Pratt precedence ===")
	lex, lerr := lexer()
	if lerr != nil {
		fail(lerr)
		return
	}
	parser := parser3.NewParser(lex, prattGrammar(), "start", ignore())
	for _, code := range []string{"1 + 2 * 3", "1 * 2 + 3", "-1 + 2", "1 - 2 - 3"} {
		nodes, err := parser.Parse(code)
		if err != nil {
			fail(err)
		}
		children, cerr := parser3.ChildrenOf(nodes[0])
		if cerr != nil {
			fail(cerr)
		}
		fmt.Printf("%-10s -> %s\n", code, shape(children[0]))
	}
}

// shape renders the operator skeleton of a Pratt tree.
func shape(n stringParsing.ParsedNode) string {
	switch n.Switch {
	case "BinaryOp":
		return fmt.Sprintf("(%s %s %s)",
			shape(n.Metadata[parser3.MetaLeft].(stringParsing.ParsedNode)),
			n.Metadata[parser3.MetaOperator],
			shape(n.Metadata[parser3.MetaRight].(stringParsing.ParsedNode)))
	case "PrefixOp":
		return fmt.Sprintf("(%s %s)",
			n.Metadata[parser3.MetaOperator],
			shape(n.Metadata[parser3.MetaOperand].(stringParsing.ParsedNode)))
	default:
		return n.Raw
	}
}

// errorDemo shows what a failure looks like in plain text and in color.
func errorDemo() {
	fmt.Println("\n=== Errors ===")
	lex, lerr := lexer()
	if lerr != nil {
		fail(lerr)
		return
	}
	parser := parser3.NewParser(lex, calcGrammar(), "expr", ignore())

	for _, code := range []string{"3 + * 4", "3 + 4)", "3 +"} {
		_, err := parser.Parse(code)
		if err == nil {
			continue
		}
		fmt.Printf("%q\n%s\n\n", code, parser3.FormatError(err, false))
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, parser3.FormatErrorPretty(err))
	os.Exit(1)
}

func printAST(node stringParsing.ParsedNode, indent int) {
	prefix := strings.Repeat("  ", indent)
	fmt.Printf("%s[%s] Raw: %q\n", prefix, node.Switch, node.Raw)
	children, err := parser3.ChildrenOf(node)
	if err != nil {
		return
	}
	for _, child := range children {
		printAST(child, indent+1)
	}
}

func cleanNode(node stringParsing.ParsedNode) stringParsing.ParsedNode {
	meta := make(map[string]interface{})
	for k, v := range node.Metadata {
		switch k {
		case "__prev", "__next":
			continue
		case parser3.MetaChildren:
			children, ok := v.([]stringParsing.ParsedNode)
			if !ok {
				continue
			}
			cleaned := make([]stringParsing.ParsedNode, len(children))
			for i, child := range children {
				cleaned[i] = cleanNode(child)
			}
			meta[k] = cleaned
		default:
			if sub, ok := v.(stringParsing.ParsedNode); ok {
				meta[k] = cleanNode(sub)
				continue
			}
			meta[k] = v
		}
	}
	return stringParsing.ParsedNode{Raw: node.Raw, Switch: node.Switch, Metadata: meta}
}
