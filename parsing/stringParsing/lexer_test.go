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
		{Type: "BLOCK", Pattern: regexp2.MustCompile(`(?s)\s*begin{(.*)?}end`, 0)},
		{Type: "COMMENT", Pattern: regexp2.MustCompile(`(?s)/\*\s*@(.+?)@\*/`, 0)},
		{Type: "COMMENT", Pattern: regexp2.MustCompile(`//@.*`, 0)},
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

	lexer := NewLexer(rules, &LexerConfig{
		UseBracketBalance: true,
		Brackets:          [][2]string{{"begin{", "}end"}},
	})
	nodes, err := lexer.Parse(`//@ test
1.0 test_param1 = "\" string \""
begin{
	test block
}end`,
		&parsing.ParseOption{UEP: &core.UniversalEngineParams{Logger: core.NewLogger("")}})
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	types := []string{
		"COMMENT", "WHITESPACE", "NUMBER", "WHITESPACE", "IDENT",
		"WHITESPACE", "EQ", "WHITESPACE", "STRING", "WHITESPACE",
		"BLOCK",
	}
	if len(nodes) < len(types) {
		t.Fatalf("Too few tokens: got %d, want at least %d", len(nodes), len(types))
	}
	for idx, expected := range types {
		if nodes[idx].Switch != expected {
			t.Fatalf("Mismatch at index %d: expected %q, got %q", idx, expected, nodes[idx].Switch)
		}
	}
}
