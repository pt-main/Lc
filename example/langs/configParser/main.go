// Build a small config language with parser3: nested blocks, a list of values
// and an ActionExpr that validates them while the parse is still running.

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/parsing/stringParsing"
	"github.com/pt-main/lc/v2/parsing/stringParsing/parser3"
)

func lexer() *stringParsing.Lexer {
	rules := []stringParsing.LexerRule{
		{Type: "LBRACE", Pattern: regexp2.MustCompile(`\{`, 0)},
		{Type: "RBRACE", Pattern: regexp2.MustCompile(`\}`, 0)},
		{Type: "ASSIGN", Pattern: regexp2.MustCompile(`=`, 0)},
		{Type: "COMMA", Pattern: regexp2.MustCompile(`,`, 0)},
		{Type: "NUMBER", Pattern: regexp2.MustCompile(`\d+`, 0)},
		{Type: "IDENT", Pattern: regexp2.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`, 0)},
		{Type: "WS", Pattern: regexp2.MustCompile(`\s+`, 0)},
	}
	return stringParsing.NewLexer(rules, nil)
}

func grammar() parser3.Grammar {
	numberValue := parser3.ActionExpr{
		Expr: parser3.TokenExpr{TokenType: "NUMBER"},
		Action: func(nodes []stringParsing.ParsedNode) (stringParsing.ParsedNode, core.ErrorInterface) {
			n, err := strconv.Atoi(nodes[0].Raw)
			if err != nil {
				return stringParsing.ParsedNode{}, &parser3.AdapterError{Msg: "bad number: " + nodes[0].Raw}
			}
			if n > 100 {
				return stringParsing.ParsedNode{}, &parser3.AdapterError{
					Msg: fmt.Sprintf("value %d is out of range (max 100)", n),
				}
			}
			return stringParsing.ParsedNode{
				Switch:   "number",
				Raw:      nodes[0].Raw,
				Metadata: map[string]interface{}{"value": n},
			}, nil
		},
	}

	return parser3.Grammar{
		"config": {Name: "config", Expr: parser3.NodeExpr{NodeType: "config", Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
			parser3.RepeatExpr{Expr: parser3.NamedExpr{RuleName: "entry"}},
		}}}},
		"entry": {Name: "entry", Expr: parser3.ChoiceExpr{Alternatives: []parser3.Expr{
			parser3.NamedExpr{RuleName: "pair"},
			parser3.NamedExpr{RuleName: "block"},
		}}},
		"pair": {Name: "pair", Expr: parser3.NodeExpr{NodeType: "pair", Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
			parser3.TokenExpr{TokenType: "IDENT"},
			parser3.TokenExpr{TokenType: "ASSIGN"},
			parser3.NamedExpr{RuleName: "value"},
		}}}},
		"block": {Name: "block", Expr: parser3.NodeExpr{NodeType: "block", Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
			parser3.TokenExpr{TokenType: "IDENT"},
			parser3.TokenExpr{TokenType: "LBRACE"},
			parser3.NamedExpr{RuleName: "config"},
			parser3.TokenExpr{TokenType: "RBRACE"},
		}}}},
		"value": {Name: "value", Expr: parser3.ChoiceExpr{Alternatives: []parser3.Expr{
			numberValue,
			parser3.TokenExpr{TokenType: "IDENT"},
		}}},
	}
}

func main() {
	p := parser3.NewParser(lexer(), grammar(), "config", []string{"WS"})

	source := "name = 10\nlimits {\n  cpu = 4\n  mem = 32\n}"
	fmt.Println("source:")
	fmt.Println(source)

	nodes, err := p.Parse(source)
	if err != nil {
		fmt.Println("\nparse failed:")
		fmt.Println(parser3.FormatError(err, false))
		return
	}

	fmt.Println("\nAST:")
	printTree(nodes[0], 0)

	fmt.Println("\nvalues collected by the action:")
	collectNumbers(nodes[0])

	fmt.Println("\nrejected input (value over the limit):")
	if _, err := p.Parse("name = 500"); err != nil {
		fmt.Println(parser3.FormatError(err, false))
	}
}

func printTree(node stringParsing.ParsedNode, indent int) {
	fmt.Printf("%s[%s] %q\n", strings.Repeat("  ", indent), node.Switch, node.Raw)
	children, cerr := parser3.ChildrenOf(node)
	if cerr != nil {
		return
	}
	for _, child := range children {
		printTree(child, indent+1)
	}
}

func collectNumbers(node stringParsing.ParsedNode) {
	if node.Switch == "number" {
		if v, ok := node.Metadata["value"].(int); ok {
			fmt.Printf("  %d\n", v)
		}
	}
	children, cerr := parser3.ChildrenOf(node)
	if cerr != nil {
		return
	}
	for _, child := range children {
		collectNumbers(child)
	}
}
