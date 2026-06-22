package stringParsing

import (
	"testing"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
)

func TestLexer_Parse(t *testing.T) {
	rules := []LexerRule{
		{Type: "WHITESPACE", Pattern: regexp2.MustCompile(`\s+`, 0)},
		{Type: "COMMENT", Pattern: regexp2.MustCompile(`(?s)/\*\s*@(.+?)@\*/`, 0)},
		{Type: "COMMENT", Pattern: regexp2.MustCompile(`//@.+?$`, 0)},
		{Type: "IDENT", Pattern: regexp2.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]+`, 0)},
		{Type: "NUMBER", Pattern: regexp2.MustCompile(`[0-9]+(?:\.[0-9]+)?`, 0)},
		{Type: "STRING", Pattern: regexp2.MustCompile(`"(?:[^"\\]|\\.)*"`, 0)},
		{Type: "LBRACE", Pattern: regexp2.MustCompile(`\{`, 0)},
		{Type: "RBRACE", Pattern: regexp2.MustCompile(`\}`, 0)},
		{Type: "LPAREN", Pattern: regexp2.MustCompile(`\(`, 0)},
		{Type: "RPAREN", Pattern: regexp2.MustCompile(`\)`, 0)},
		{Type: "COMMA", Pattern: regexp2.MustCompile(`,`, 0)},
		{Type: "EQ", Pattern: regexp2.MustCompile(`=`, 0)},
	}

	lexer := NewLexer(rules, nil)
	nodes, err := lexer.Parse(`//@ test
1.0 test_param1 = "\" string \""`,
		&parsing.ParseOption{UEP: &core.UniversalEngineParams{Logger: core.NewLogger("")}})
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}

	types := []string{"COMMENT", "NUMBER", "WHITESPACE", "IDENT", "WHITESPACE", "EQ", "WHITESPACE", "STRING"}
	for idx, node := range nodes {
		if types[idx] != node.Switch {
			t.Fatalf("Mismatched types: must be %v, got %v", types[idx], node.Switch)
		}
	}
}
