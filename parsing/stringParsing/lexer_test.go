package stringParsing

import (
	"testing"

	"github.com/dlclark/regexp2"
)

func TestLexer_Parse(t *testing.T) {
	rules := []LexerRule{
		{
			Type:    "NUMBER",
			Pattern: regexp2.MustCompile(`\d+`, 0),
		},
		{
			Type:    "IDENT",
			Pattern: regexp2.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`, 0),
		},
		{
			Type:    "WHITESPACE",
			Pattern: regexp2.MustCompile(`\s+`, 0),
		},
	}

	lexer := NewLexer(rules, nil)
	nodes, err := lexer.Parse("123 abc 456", nil)
	if err != nil {
		t.Fatalf("Error: %s", err)
		return
	}

	types := []string{"NUMBER", "WHITESPACE", "IDENT", "WHITESPACE", "NUMBER"}
	for idx, node := range nodes {
		if types[idx] != node.Switch {
			t.Fatalf("Mismatched types: must be %v, got %v", types[idx], node.Switch)
		}
	}
}
