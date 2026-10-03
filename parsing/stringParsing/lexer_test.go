package stringParsing

import (
	"fmt"
	"testing"
	"time"

	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/parsing"
)

// mustLexer builds a lexer for a test and fails the run on invalid rules,
// keeping the call sites free of error plumbing.
func mustLexer(tb testing.TB, rules []LexerRule, config *LexerConfig) *Lexer {
	tb.Helper()
	lex, err := NewLexer(rules, config)
	if err != nil {
		tb.Fatalf("NewLexer: %v", err)
	}
	return lex
}

func TestLexer_Parse(t *testing.T) {
	rules := []LexerRule{
		{Type: "WHITESPACE", Pattern: `\s+`},
		{Type: "BLOCK", Pattern: `(?s)\s*begin\{(.*)?\}end`},
		{Type: "COMMENT", Pattern: `(?s)/\*\s*@(.+?)@\*/`},
		{Type: "COMMENT", Pattern: `//@.*`},
		{Type: "IDENT", Pattern: `[a-zA-Z_][a-zA-Z0-9_]+`},
		{Type: "NUMBER", Pattern: `[0-9]+(?:\.[0-9]+)?`},
		{Type: "STRING", Pattern: `"(?:[^"\\]|\\.)*"`},
		{Type: "LBRACE", Pattern: `\{`},
		{Type: "RBRACE", Pattern: `\}`},
		{Type: "LPAREN", Pattern: `\(`},
		{Type: "RPAREN", Pattern: `\)`},
		{Type: "COMMA", Pattern: `,`},
		{Type: "EQ", Pattern: `=`},
	}

	lexer := mustLexer(t, rules, &LexerConfig{
		UseBracketBalance: true,
		Brackets:          [][2]string{{"begin{", "}end"}},
	})
	nodes, err := lexer.Parse(`//@ test
1.0 test_param1 = "..." string "   "
begin{
	test block
}end`,
		&parsing.ParseOption{UEP: &core.UniversalEngineParams{Logger: core.NewLogger("")}})
	if err != nil {
		t.Fatalf("Error: %s", err)
	}
	types := []string{
		"COMMENT", "WHITESPACE", "NUMBER", "WHITESPACE", "IDENT",
		"WHITESPACE", "EQ", "WHITESPACE", "STRING", "WHITESPACE", "IDENT",
		"WHITESPACE", "STRING", "WHITESPACE", "BLOCK",
	}
	if len(nodes) < len(types) {
		t.Fatalf("Too few tokens: got %d, want at least %d", len(nodes), len(types))
	}
	for idx, expected := range types {
		fmt.Printf("%v '%v' %v : %v\n", idx, nodes[idx].Raw, nodes[idx].Switch, expected)
		if nodes[idx].Switch != expected {
			t.Fatalf("Mismatch at index %d: expected %q, got %q", idx, expected, nodes[idx].Switch)
		}
	}
}

func TestLexer_ZeroLengthMatchDoesNotHang(t *testing.T) {
	rules := []LexerRule{
		{Type: "EMPTY", Pattern: `x*`},
		{Type: "A", Pattern: `a`},
	}
	lexer := mustLexer(t, rules, nil)

	done := make(chan struct{})
	var nodes []ParsedNode
	var err error
	go func() {
		nodes, err = lexer.Parse("abc")
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("a rule that can match the empty string made the lexer spin forever")
	}
	if err == nil {
		t.Fatalf("an empty match must not be accepted as a token, got %d nodes", len(nodes))
	}
}
