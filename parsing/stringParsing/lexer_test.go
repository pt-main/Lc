package stringParsing

import (
	"testing"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc/parsing"
)

func TestLexerWithBlocks(t *testing.T) {
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
			Type:    "LINES",
			Pattern: regexp2.MustCompile(`(?s).+?`, 0),
		},
		{
			Type:    "WHITESPACE",
			Pattern: regexp2.MustCompile(`\s+`, 0),
		},
	}

	config := &LexerConfig{
		UseLineContinuation: true,
		UseBracketBalance:   true,
		Brackets:            []string{"()", "[]"},
		SkipEmptyLines:      true,
		TrimBlocksSpace:     true,
	}

	lexer := NewLexer(rules, config)

	code := `123 abc 456
(continued
line)
[another
block]`
	nodes, err := lexer.Parse(code, &parsing.ParseOption{})
	if err != nil {
		t.Fatalf("Error: %s", err)
	}

	// Expect flat token list with block_index metadata
	for _, n := range nodes {
		t.Logf("Token: %v, block: %v", n.Switch, n.Metadata["__block_index"])
	}
}
