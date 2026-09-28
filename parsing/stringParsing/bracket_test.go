package stringParsing

import "testing"

// isBracketBalanced and bracketBalanceError must never disagree. The case
// with a multi-byte word also pins down that the balance scan is rune aware.
func TestBracketChecksAgree(t *testing.T) {
	lex := NewLexer(nil, &LexerConfig{
		UseBracketBalance: true,
		Brackets:          [][2]string{{"(", ")"}, {"[", "]"}, {"{", "}"}},
	})
	cases := []string{
		"(a)", "(a", "a)", "[(a)]", "([a)]", "{[()]}", "", "()[]{}",
		"([)]", "(((", "a(b)c", "unicode (текст) ok",
	}
	for _, c := range cases {
		bal := lex.isBracketBalanced(c)
		err := lex.bracketBalanceError(c)
		if (err == nil) != bal {
			t.Errorf("%q: isBracketBalanced=%v but bracketBalanceError=%v", c, bal, err)
		}
	}
}
