package stringParsing

import (
	"testing"
	"time"

	"github.com/dlclark/regexp2"
)

func perfLexer() *Lexer {
	return NewLexer([]LexerRule{
		{Type: "NUMBER", Pattern: regexp2.MustCompile(`\d+(\.\d+)?`, 0)},
		{Type: "IDENT", Pattern: regexp2.MustCompile(`[a-z]+`, 0)},
		{Type: "WS", Pattern: regexp2.MustCompile(`[ \t\r\n]+`, 0)},
	}, &LexerConfig{UseBracketBalance: false})
}

// Parse used to copy the whole remaining input into a fresh string at every
// position, which made it quadratic: 6 KB took 160 ms, 24 KB took 2.9 s and
// 96 KB took 48 s, allocating gigabytes. Matching on the shared rune slice
// keeps it linear.
func TestLexer_ParseIsLinearInInputSize(t *testing.T) {
	unit := "12345 abcde "
	// 60 KB, roughly the size that used to need tens of seconds
	var input string
	for i := 0; i < 5000; i++ {
		input += unit
	}

	start := time.Now()
	nodes, err := perfLexer().Parse(input)
	elapsed := time.Since(start)
	if err != nil {
		t.Fatalf("Parse returned %v", err)
	}
	// 4 tokens per unit: NUMBER, WS, IDENT, WS
	if len(nodes) != 20000 {
		t.Errorf("token count = %d, want 20000", len(nodes))
	}
	if elapsed > 5*time.Second {
		t.Errorf("parsing %d bytes took %v, which is far from linear", len(input), elapsed)
	}
}

// A rule that matches only later in the input must not be accepted at an
// earlier position: the old code matched a copy of the tail, so a match away
// from position 0 was invisible, and the new code has to keep that behaviour.
func TestLexer_MatchMustBeAnchoredAtPosition(t *testing.T) {
	lexer := NewLexer([]LexerRule{
		{Type: "LATE", Pattern: regexp2.MustCompile(`bc`, 0)},
		{Type: "A", Pattern: regexp2.MustCompile(`a`, 0)},
	}, nil)

	nodes, err := lexer.Parse("abc")
	if err != nil {
		t.Fatalf("Parse returned %v", err)
	}
	if nodes[0].Switch != "A" || nodes[0].Raw != "a" {
		t.Errorf("first token = %s(%q), want A(a)", nodes[0].Switch, nodes[0].Raw)
	}
	// A must win at position 0 even though LATE also matches later in the
	// input, and the remaining "bc" is then taken by LATE.
	if len(nodes) != 2 {
		t.Fatalf("token count = %d, want 2 (a, bc)", len(nodes))
	}
	if nodes[1].Switch != "LATE" || nodes[1].Raw != "bc" {
		t.Errorf("last token = %s(%q), want LATE(bc)", nodes[1].Switch, nodes[1].Raw)
	}
}

// The lexer must not depend on a logger being wired up, and must not crash
// when ParseOption carries no UEP.
func TestLexer_ParseOptionWithoutUEP(t *testing.T) {
	for _, code := range []string{"1 2 3", "abc", ""} {
		if _, err := perfLexer().Parse(code, nil); err != nil {
			t.Errorf("Parse(%q, nil) = %v", code, err)
		}
	}
}
