package parser3

import (
	"strings"
	"testing"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/stringParsing"
)

func newCalcLexer() *stringParsing.Lexer {
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
	return stringParsing.NewLexer(rules, &stringParsing.LexerConfig{})
}

func calcGrammar() Grammar {
	return Grammar{
		"expr": {Name: "expr", Expr: NodeExpr{NodeType: "expr", Expr: SequenceExpr{Exprs: []Expr{
			NamedExpr{RuleName: "term"},
			RepeatExpr{Expr: SequenceExpr{Exprs: []Expr{
				ChoiceExpr{Alternatives: []Expr{
					TokenExpr{TokenType: "PLUS"},
					TokenExpr{TokenType: "MINUS"},
				}},
				NamedExpr{RuleName: "term"},
			}}},
		}}}},
		"term": {Name: "term", Expr: NodeExpr{NodeType: "term", Expr: SequenceExpr{Exprs: []Expr{
			NamedExpr{RuleName: "factor"},
			RepeatExpr{Expr: SequenceExpr{Exprs: []Expr{
				ChoiceExpr{Alternatives: []Expr{
					TokenExpr{TokenType: "STAR"},
					TokenExpr{TokenType: "SLASH"},
				}},
				NamedExpr{RuleName: "factor"},
			}}},
		}}}},
		"factor": {Name: "factor", Expr: NodeExpr{NodeType: "factor", Expr: ChoiceExpr{Alternatives: []Expr{
			TokenExpr{TokenType: "NUMBER"},
			SequenceExpr{Exprs: []Expr{
				TokenExpr{TokenType: "LPAREN"},
				NamedExpr{RuleName: "expr"},
				TokenExpr{TokenType: "RPAREN"},
			}},
		}}}},
	}
}

func newCalcParser() *Parser {
	return NewParser(newCalcLexer(), calcGrammar(), "expr", []string{"WHITESPACE"})
}

func TestParseExpression(t *testing.T) {
	p := newCalcParser()
	nodes, err := p.Parse("3 + 5 * (2 - 1)")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(nodes) != 1 {
		t.Fatalf("expected 1 root node, got %d", len(nodes))
	}
	if nodes[0].Switch != "expr" {
		t.Errorf("root switch = %q, want %q", nodes[0].Switch, "expr")
	}
	if nodes[0].Raw != "3+5*(2-1)" {
		t.Errorf("root raw = %q, want %q", nodes[0].Raw, "3+5*(2-1)")
	}
}

func TestParseSingleNumber(t *testing.T) {
	p := newCalcParser()
	nodes, err := p.Parse("42")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if nodes[0].Raw != "42" {
		t.Errorf("raw = %q, want %q", nodes[0].Raw, "42")
	}
}

func TestParseMissingOperand(t *testing.T) {
	p := newCalcParser()
	_, err := p.Parse("3 +")
	if err == nil {
		t.Fatal("expected an error for a trailing operator")
	}
	if !strings.Contains(err.Error(), "RPAREN") && !strings.Contains(err.Error(), "term") {
		t.Logf("error message: %v", err)
	}
	if err.GetCode() != ParseErrCode {
		t.Errorf("GetCode() = %q, want %q", err.GetCode(), ParseErrCode)
	}
}

func TestParseTrailingTokens(t *testing.T) {
	p := newCalcParser()
	_, err := p.Parse("3 (4)")
	if err == nil {
		t.Fatal("expected an error for unconsumed input")
	}
	pe, ok := AsParseError(err)
	if !ok {
		t.Fatalf("error is not a *ParseError: %T", err)
	}
	if pe.Phase == "" {
		t.Error("a parse error should name the phase that failed")
	}
	if !strings.Contains(pe.Error(), "LPAREN") {
		t.Errorf("error should name the offending token: %v", pe)
	}
}

func TestErrorPointsAtDeepestFailure(t *testing.T) {
	p := newCalcParser()
	_, err := p.Parse("3 + * 4")
	if err == nil {
		t.Fatal("expected an error")
	}
	pe, ok := AsParseError(err)
	if !ok {
		t.Fatalf("error is not a *ParseError: %T", err)
	}
	if pe.Got != "STAR" {
		t.Errorf("the deepest failure should be reported, got Got = %q in %v", pe.Got, pe)
	}
}

func TestLexerError(t *testing.T) {
	p := newCalcParser()
	_, err := p.Parse("3 + $")
	if err == nil {
		t.Fatal("expected a lexer error")
	}
	pe, ok := AsParseError(err)
	if !ok {
		t.Fatalf("error is not a *ParseError: %T", err)
	}
	if pe.Phase != PhaseLexer {
		t.Errorf("phase = %q, want %q", pe.Phase, PhaseLexer)
	}
}

func TestMissingStartRule(t *testing.T) {
	p := NewParser(newCalcLexer(), calcGrammar(), "nope", []string{"WHITESPACE"})
	_, err := p.Parse("1")
	if err == nil {
		t.Fatal("expected an error for an unknown start rule")
	}
	if err.GetCode() != GrammarErrCode {
		t.Errorf("GetCode() = %q, want %q", err.GetCode(), GrammarErrCode)
	}
}

func TestUndefinedRule(t *testing.T) {
	g := Grammar{"start": {Name: "start", Expr: NamedExpr{RuleName: "missing"}}}
	p := NewParser(newCalcLexer(), g, "start", nil)
	_, err := p.Parse("1")
	if err == nil {
		t.Fatal("expected an error for an undefined rule")
	}
	if !strings.Contains(err.Error(), "missing") {
		t.Errorf("error should name the missing rule: %v", err)
	}
}

func TestLeftRecursionDetected(t *testing.T) {
	g := Grammar{
		"start": {Name: "start", Expr: SequenceExpr{Exprs: []Expr{
			TokenExpr{TokenType: "NUMBER"},
			OptionalExpr{Expr: NamedExpr{RuleName: "start"}},
		}}},
	}
	p := NewParser(newCalcLexer(), g, "start", nil)
	_, err := p.Parse("1")
	if err != nil {
		t.Fatalf("right recursion must still work: %v", err)
	}

	bad := Grammar{
		"start": {Name: "start", Expr: SequenceExpr{Exprs: []Expr{
			NamedExpr{RuleName: "start"},
			TokenExpr{TokenType: "NUMBER"},
		}}},
	}
	p2 := NewParser(newCalcLexer(), bad, "start", nil)
	if _, err := p2.Parse("1"); err == nil {
		t.Fatal("expected left recursion to be reported, not to hang")
	}
}

func TestGrammarErrorCodeIsStable(t *testing.T) {
	ge := &GrammarError{Phase: "NamedExpr", Msg: "boom"}
	if ge.GetCode() != GrammarErrCode {
		t.Errorf("GetCode() = %q, want %q", ge.GetCode(), GrammarErrCode)
	}
	if ge.Code() != "NamedExpr" {
		t.Errorf("Code() = %q, want %q", ge.Code(), "NamedExpr")
	}
}

func TestFormatErrorNoColors(t *testing.T) {
	err := &ParseError{Phase: PhaseExpect, Expected: "PLUS", Got: "STAR", Raw: "*", TokenPos: "idx=3 start=2-3"}
	out := FormatError(err, false)
	if strings.Contains(out, "\x1b[") {
		t.Errorf("plain output must not contain ANSI codes: %q", out)
	}
	for _, want := range []string{"parser3/Expect", "PLUS", "STAR", "idx=3"} {
		if !strings.Contains(out, want) {
			t.Errorf("output %q should contain %q", out, want)
		}
	}
}

func TestFormatErrorWithColors(t *testing.T) {
	err := &ParseError{Phase: PhaseExpect, Expected: "PLUS", Got: "STAR", Raw: "*"}
	out := FormatErrorPretty(err)
	if !strings.Contains(out, "\x1b[") {
		t.Errorf("pretty output should contain ANSI codes: %q", out)
	}
}

func TestFormatErrorChain(t *testing.T) {
	inner := &ParseError{Phase: PhaseExpect, Expected: "RPAREN", Got: "NUMBER", Raw: "2"}
	outer := &GrammarError{Phase: "ChoiceExpr", Msg: "no alternative matched", Cause: inner}
	out := FormatError(outer, false)
	if !strings.Contains(out, "ChoiceExpr") || !strings.Contains(out, "RPAREN") {
		t.Errorf("chain output should mention both errors: %q", out)
	}
	if !strings.Contains(out, "caused by") {
		t.Errorf("chain output should mark the cause: %q", out)
	}
}

func TestFormatErrorNil(t *testing.T) {
	if got := FormatError(nil, true); got != "" {
		t.Errorf("FormatError(nil) = %q, want empty", got)
	}
}

func TestParseErrorMeta(t *testing.T) {
	err := &ParseError{Phase: PhasePeek, Expected: "LPAREN", Got: "NUMBER", Raw: "7", TokenIdx: 2, TokenPos: "idx=2"}
	meta := err.GetMeta()
	if meta["Expected"] != "LPAREN" {
		t.Errorf("meta Expected = %v, want LPAREN", meta["Expected"])
	}
	if meta["Code"] != PhasePeek {
		t.Errorf("meta Code = %v, want %q", meta["Code"], PhasePeek)
	}
}

func TestActionExpr(t *testing.T) {
	called := false
	g := Grammar{
		"start": {Name: "start", Expr: ActionExpr{
			Expr: TokenExpr{TokenType: "NUMBER"},
			Action: func(nodes []stringParsing.ParsedNode) (stringParsing.ParsedNode, core.ErrorInterface) {
				called = true
				return stringParsing.ParsedNode{
					Switch:   "doubled",
					Raw:      nodes[0].Raw + nodes[0].Raw,
					Metadata: map[string]interface{}{MetaChildren: nodes},
				}, nil
			},
		}},
	}
	p := NewParser(newCalcLexer(), g, "start", nil)
	nodes, err := p.Parse("21")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !called {
		t.Error("action was not called")
	}
	if nodes[0].Raw != "2121" {
		t.Errorf("raw = %q, want %q", nodes[0].Raw, "2121")
	}
}

func TestActionExprNilAction(t *testing.T) {
	g := Grammar{
		"start": {Name: "start", Expr: ActionExpr{Expr: TokenExpr{TokenType: "NUMBER"}}},
	}
	p := NewParser(newCalcLexer(), g, "start", nil)
	nodes, err := p.Parse("5")
	if err != nil {
		t.Fatalf("a nil action must pass nodes through: %v", err)
	}
	child := onlyChild(t, nodes)
	if child.Switch != "NUMBER" {
		t.Errorf("child = %q, want NUMBER", child.Switch)
	}
}

func TestPeekAndNotConsume(t *testing.T) {
	g := Grammar{
		"start": {Name: "start", Expr: SequenceExpr{Exprs: []Expr{
			PeekExpr{TokenType: "NUMBER"},
			TokenExpr{TokenType: "NUMBER"},
		}}},
	}
	p := NewParser(newCalcLexer(), g, "start", nil)
	if _, err := p.Parse("7"); err != nil {
		t.Fatalf("PeekExpr must not consume: %v", err)
	}
}

func TestAndNotExpr(t *testing.T) {
	andG := Grammar{
		"start": {Name: "start", Expr: SequenceExpr{Exprs: []Expr{
			AndExpr{Expr: TokenExpr{TokenType: "NUMBER"}},
			TokenExpr{TokenType: "NUMBER"},
		}}},
	}
	p := NewParser(newCalcLexer(), andG, "start", nil)
	if _, err := p.Parse("7"); err != nil {
		t.Fatalf("AndExpr must not consume: %v", err)
	}

	notG := Grammar{
		"start": {Name: "start", Expr: SequenceExpr{Exprs: []Expr{
			NotExpr{Expr: TokenExpr{TokenType: "PLUS"}},
			TokenExpr{TokenType: "NUMBER"},
		}}},
	}
	p2 := NewParser(newCalcLexer(), notG, "start", nil)
	if _, err := p2.Parse("7"); err != nil {
		t.Fatalf("NotExpr must succeed on a non-match: %v", err)
	}
	if _, err := p2.Parse("+"); err == nil {
		t.Fatal("NotExpr must fail when the expression does match")
	}
}

func TestSeparatedRepeatExpr(t *testing.T) {
	g := Grammar{
		"start": {Name: "start", Expr: SeparatedRepeatExpr{
			Element: TokenExpr{TokenType: "NUMBER"},
			Sep:     "COMMA",
			Min:     2,
		}},
	}
	rules := []stringParsing.LexerRule{
		{Type: "NUMBER", Pattern: regexp2.MustCompile(`\d+`, 0)},
		{Type: "COMMA", Pattern: regexp2.MustCompile(`,`, 0)},
	}
	p := NewParser(stringParsing.NewLexer(rules, nil), g, "start", nil)

	nodes, err := p.Parse("1,2,3")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	children, cerr := ChildrenOf(nodes[0])
	if cerr != nil {
		t.Fatalf("unexpected error: %v", cerr)
	}
	if len(children) != 3 {
		t.Errorf("expected 3 elements, got %d", len(children))
	}

	if _, err := p.Parse("1"); err == nil {
		t.Fatal("a single element must violate Min=2")
	}
	if _, err := p.Parse("1,"); err == nil {
		t.Fatal("a trailing separator must be an error, not a silent match")
	}
}

// digitLexer matches one digit per token, so repetition counts are observable.
func digitLexer() *stringParsing.Lexer {
	rules := []stringParsing.LexerRule{
		{Type: "DIGIT", Pattern: regexp2.MustCompile(`\d`, 0)},
	}
	return stringParsing.NewLexer(rules, nil)
}

// onlyChild unwraps the single node produced by a start rule.
func onlyChild(t *testing.T, nodes []stringParsing.ParsedNode) stringParsing.ParsedNode {
	t.Helper()
	children, err := ChildrenOf(nodes[0])
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(children) != 1 {
		t.Fatalf("expected exactly 1 child, got %d", len(children))
	}
	return children[0]
}

func TestRepeatExprMin(t *testing.T) {
	g := Grammar{
		"start": {Name: "start", Expr: RepeatExpr{Expr: TokenExpr{TokenType: "DIGIT"}, Min: 2}},
	}
	p := NewParser(digitLexer(), g, "start", nil)
	if _, err := p.Parse("1"); err == nil {
		t.Fatal("a single digit must violate Min=2")
	}
	if _, err := p.Parse("12"); err != nil {
		t.Fatalf("two digits should satisfy Min=2: %v", err)
	}
}

func TestRepeatExprMax(t *testing.T) {
	g := Grammar{
		"start": {Name: "start", Expr: RepeatExpr{
			Expr: TokenExpr{TokenType: "DIGIT"},
			Max:  2,
		}},
	}
	p := NewParser(digitLexer(), g, "start", nil)
	if _, err := p.Parse("123"); err == nil {
		t.Fatal("Max=2 must reject a third repetition")
	}
}

func TestPrattExprPrecedence(t *testing.T) {
	pratt := &PrattExpr{
		Atom:     TokenExpr{TokenType: "NUMBER"},
		Prefixes: map[string]Expr{"MINUS": TokenExpr{TokenType: "MINUS"}},
		Infixes: map[string]InfixInfo{
			"PLUS":  {Precedence: 1, Assoc: LeftAssoc},
			"STAR":  {Precedence: 2, Assoc: LeftAssoc},
			"SLASH": {Precedence: 2, Assoc: LeftAssoc},
		},
	}
	g := Grammar{"start": {Name: "start", Expr: pratt}}
	p := NewParser(newCalcLexer(), g, "start", []string{"WHITESPACE"})

	nodes, err := p.Parse("1 + 2 * 3")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	root := onlyChild(t, nodes)
	if root.Switch != "BinaryOp" {
		t.Fatalf("root = %q, want BinaryOp", root.Switch)
	}
	meta := root.Metadata
	if meta[MetaOperator] != "PLUS" {
		t.Fatalf("top operator = %v, want PLUS", meta[MetaOperator])
	}
	left := meta[MetaLeft].(stringParsing.ParsedNode)
	right := meta[MetaRight].(stringParsing.ParsedNode)
	if right.Metadata[MetaOperator] != "STAR" {
		t.Errorf("* must bind tighter than +: right = %+v", right)
	}
	if left.Switch != "NUMBER" || left.Raw != "1" {
		t.Errorf("left = %+v, want NUMBER 1", left)
	}
}

func TestPrattExprPrefix(t *testing.T) {
	pratt := &PrattExpr{
		Atom:     TokenExpr{TokenType: "NUMBER"},
		Prefixes: map[string]Expr{"MINUS": TokenExpr{TokenType: "MINUS"}},
		Infixes:  map[string]InfixInfo{"STAR": {Precedence: 2, Assoc: LeftAssoc}},
	}
	g := Grammar{"start": {Name: "start", Expr: pratt}}
	p := NewParser(newCalcLexer(), g, "start", nil)

	nodes, err := p.Parse("-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if onlyChild(t, nodes).Switch != "PrefixOp" {
		t.Errorf("root = %q, want PrefixOp", onlyChild(t, nodes).Switch)
	}
}

func TestPrattExprLeftAssoc(t *testing.T) {
	pratt := &PrattExpr{
		Atom:    TokenExpr{TokenType: "NUMBER"},
		Infixes: map[string]InfixInfo{"MINUS": {Precedence: 1, Assoc: LeftAssoc}},
	}
	g := Grammar{"start": {Name: "start", Expr: pratt}}
	p := NewParser(newCalcLexer(), g, "start", nil)

	nodes, err := p.Parse("1-2-3")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	root := onlyChild(t, nodes)
	left, ok := root.Metadata[MetaLeft].(stringParsing.ParsedNode)
	if !ok {
		t.Fatalf("root has no %s metadata: %+v", MetaLeft, root)
	}
	if left.Metadata[MetaOperator] != "MINUS" {
		t.Errorf("left-associative parsing must nest on the left, got %+v", root)
	}
}

func TestPrattExprRightAssoc(t *testing.T) {
	pratt := &PrattExpr{
		Atom:    TokenExpr{TokenType: "DIGIT"},
		Infixes: map[string]InfixInfo{"POWER": {Precedence: 1, Assoc: RightAssoc}},
	}
	rules := []stringParsing.LexerRule{
		{Type: "DIGIT", Pattern: regexp2.MustCompile(`\d`, 0)},
		{Type: "POWER", Pattern: regexp2.MustCompile(`\^`, 0)},
	}
	p := NewParser(stringParsing.NewLexer(rules, nil), pratt2Grammar(pratt), "start", nil)

	nodes, err := p.Parse("2^3^4")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	root := onlyChild(t, nodes)
	right, ok := root.Metadata[MetaRight].(stringParsing.ParsedNode)
	if !ok {
		t.Fatalf("root has no %s: %+v", MetaRight, root)
	}
	if right.Metadata[MetaOperator] != "POWER" {
		t.Errorf("right-associative parsing must nest on the right, got %+v", root)
	}
}

func TestPrattExprNonAssoc(t *testing.T) {
	pratt := &PrattExpr{
		Atom:    TokenExpr{TokenType: "DIGIT"},
		Infixes: map[string]InfixInfo{"CMP": {Precedence: 1, Assoc: NonAssoc}},
	}
	rules := []stringParsing.LexerRule{
		{Type: "DIGIT", Pattern: regexp2.MustCompile(`\d`, 0)},
		{Type: "CMP", Pattern: regexp2.MustCompile(`=`, 0)},
	}
	p := NewParser(stringParsing.NewLexer(rules, nil), pratt2Grammar(pratt), "start", nil)

	if _, err := p.Parse("1=2"); err != nil {
		t.Fatalf("a single non-associative operator must parse: %v", err)
	}
	if _, err := p.Parse("1=2=3"); err == nil {
		t.Fatal("NonAssoc must reject a repeated operator at the same precedence")
	}
}

// pratt2Grammar wraps a PrattExpr in a minimal grammar.
func pratt2Grammar(pratt *PrattExpr) Grammar {
	return Grammar{"start": {Name: "start", Expr: pratt}}
}

func TestAdapter(t *testing.T) {
	a := &Adapter{Parser: newCalcParser()}
	children, err := a.Parse("1 + 2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(children) == 0 {
		t.Fatal("adapter returned no children")
	}
}

func TestAdapterNoParser(t *testing.T) {
	a := &Adapter{}
	if _, err := a.Parse("1"); err == nil {
		t.Fatal("expected an error when the adapter has no parser")
	}
}

func TestAdapterMissingChildren(t *testing.T) {
	if _, err := ChildrenOf(stringParsing.ParsedNode{Switch: "bare", Raw: "1"}); err == nil {
		t.Fatal("expected an error when children metadata is missing")
	}
}

func TestAdapterWrongChildrenType(t *testing.T) {
	node := stringParsing.ParsedNode{
		Switch:   "bad",
		Raw:      "1",
		Metadata: map[string]interface{}{MetaChildren: "not a slice"},
	}
	if _, err := ChildrenOf(node); err == nil {
		t.Fatal("expected an error when children metadata is not a node list")
	}
}

func TestMemoizationKeepsResults(t *testing.T) {
	g := Grammar{
		"start": {Name: "start", Expr: ChoiceExpr{Alternatives: []Expr{
			SequenceExpr{Exprs: []Expr{
				NamedExpr{RuleName: "item"},
				NamedExpr{RuleName: "item"},
			}},
			NamedExpr{RuleName: "item"},
		}}},
		"item": {Name: "item", Expr: TokenExpr{TokenType: "DIGIT"}},
	}
	p := NewParser(digitLexer(), g, "start", nil)

	nodes, err := p.Parse("7")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := nodes[0].Raw; got != "7" {
		t.Errorf("raw = %q, want 7", got)
	}

	nodes, err = p.Parse("78")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := nodes[0].Raw; got != "78" {
		t.Errorf("raw = %q, want 78", got)
	}
}

func TestMemoizationIsPerPosition(t *testing.T) {
	g := Grammar{
		"start": {Name: "start", Expr: SequenceExpr{Exprs: []Expr{
			NamedExpr{RuleName: "item"},
			TokenExpr{TokenType: "COMMA"},
			NamedExpr{RuleName: "item"},
		}}},
		"item": {Name: "item", Expr: TokenExpr{TokenType: "DIGIT"}},
	}
	rules := []stringParsing.LexerRule{
		{Type: "DIGIT", Pattern: regexp2.MustCompile(`\d`, 0)},
		{Type: "COMMA", Pattern: regexp2.MustCompile(`,`, 0)},
	}
	p := NewParser(stringParsing.NewLexer(rules, nil), g, "start", nil)

	nodes, err := p.Parse("1,2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := nodes[0].Raw; got != "1,2" {
		t.Errorf("raw = %q, want 1,2", got)
	}
}

func TestParserReuseIsSafe(t *testing.T) {
	p := newCalcParser()
	for i, code := range []string{"1+2", "3*4", "(5-6)", "7"} {
		if _, err := p.Parse(code); err != nil {
			t.Fatalf("call %d with %q failed: %v", i, code, err)
		}
	}
}

func TestGrammarErrorIsNotMaskedByChoice(t *testing.T) {
	recursive := Grammar{"start": {Name: "start", Expr: ChoiceExpr{Alternatives: []Expr{
		NamedExpr{RuleName: "start"},
		TokenExpr{TokenType: "NUMBER"},
	}}}}
	p := NewParser(newCalcLexer(), recursive, "start", nil)
	if _, err := p.Parse("1"); err == nil {
		t.Fatal("left recursion must be reported, not silently accepted")
	} else if _, ok := AsGrammarError(err); !ok {
		t.Errorf("error should be a *GrammarError, got %T: %v", err, err)
	}
}

func TestGrammarErrorFromUndefinedRuleIsNotMasked(t *testing.T) {
	g := Grammar{"start": {Name: "start", Expr: ChoiceExpr{Alternatives: []Expr{
		NamedExpr{RuleName: "missing"},
		TokenExpr{TokenType: "NUMBER"},
	}}}}
	p := NewParser(newCalcLexer(), g, "start", nil)
	_, err := p.Parse("1")
	if err == nil {
		t.Fatal("an undefined rule must be reported even when a later alternative matches")
	}
	if _, ok := AsGrammarError(err); !ok {
		t.Errorf("error should be a *GrammarError, got %T: %v", err, err)
	}
}

func TestActionRejectionIsNotRetriedByChoice(t *testing.T) {
	g := Grammar{"start": {Name: "start", Expr: ChoiceExpr{Alternatives: []Expr{
		ActionExpr{
			Expr: TokenExpr{TokenType: "DIGIT"},
			Action: func(nodes []stringParsing.ParsedNode) (stringParsing.ParsedNode, core.ErrorInterface) {
				if nodes[0].Raw == "1" {
					return stringParsing.ParsedNode{}, &AdapterError{Msg: "digit 1 is reserved"}
				}
				return nodes[0], nil
			},
		},
		TokenExpr{TokenType: "DIGIT"},
	}}}}
	p := NewParser(digitLexer(), g, "start", nil)
	nodes, err := p.Parse("1")
	if err == nil {
		t.Fatalf("a rejected input must not be accepted by a later alternative, got %+v", nodes)
	}
	if !strings.Contains(err.Error(), "reserved") {
		t.Errorf("the action error must survive: %v", err)
	}
}

func TestNilExprReportsGrammarError(t *testing.T) {
	rules := map[string]Expr{
		"nilRule":     nil,
		"nilSequence": SequenceExpr{Exprs: []Expr{nil}},
		"nilChoice":   ChoiceExpr{Alternatives: []Expr{nil}},
	}
	for name, expr := range rules {
		g := Grammar{"start": {Name: "start", Expr: expr}}
		p := NewParser(newCalcLexer(), g, "start", nil)
		_, err := p.Parse("1")
		if err == nil {
			t.Errorf("%s: expected a grammar error for a nil expression", name)
			continue
		}
		if _, ok := AsGrammarError(err); !ok {
			t.Errorf("%s: error should be a *GrammarError, got %T: %v", name, err, err)
		}
	}
}

func TestStaleErrorFromRolledBackOptional(t *testing.T) {
	rules := []stringParsing.LexerRule{
		{Type: "DIGIT", Pattern: regexp2.MustCompile(`\d`, 0)},
		{Type: "COMMA", Pattern: regexp2.MustCompile(`,`, 0)},
		{Type: "WS", Pattern: regexp2.MustCompile(`\s+`, 0)},
	}
	g := Grammar{"start": {Name: "start", Expr: OptionalExpr{Expr: SequenceExpr{Exprs: []Expr{
		TokenExpr{TokenType: "DIGIT"},
		TokenExpr{TokenType: "COMMA"},
	}}}}}
	p := NewParser(stringParsing.NewLexer(rules, nil), g, "start", []string{"WS"})
	_, err := p.Parse("5 5")
	if err == nil {
		t.Fatal("expected an error for unconsumed input")
	}
	pe, ok := AsParseError(err)
	if !ok {
		t.Fatalf("error is not a *ParseError: %T", err)
	}
	if pe.Phase != PhaseEnd {
		t.Errorf("phase = %q, want %q: %v", pe.Phase, PhaseEnd, pe)
	}
	if pe.Got == "WS" {
		t.Errorf("an ignored token must never be blamed: %v", pe)
	}
}

func TestChildrenOfAcceptsPointerSlice(t *testing.T) {
	leaf := &stringParsing.ParsedNode{Switch: "DIGIT", Raw: "1"}
	node := stringParsing.ParsedNode{
		Switch:   "n",
		Raw:      "1",
		Metadata: map[string]interface{}{MetaChildren: []*stringParsing.ParsedNode{leaf}},
	}
	children, err := ChildrenOf(node)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(children) != 1 || children[0].Switch != "DIGIT" || children[0].Raw != "1" {
		t.Errorf("children = %+v, want one DIGIT 1", children)
	}
}

func TestFormatErrorKeepsMsg(t *testing.T) {
	err := &ParseError{Phase: "RepeatExpr", Expected: "at least 1 element(s)", Msg: "repetition stopped after 0"}
	if !strings.Contains(err.Format(), "repetition stopped after 0") {
		t.Errorf("Format() drops Msg: %q", err.Format())
	}
	if !strings.Contains(err.Error(), "repetition stopped after 0") {
		t.Errorf("Error() drops Msg: %q", err.Error())
	}
}

func TestPrattPrefixBindsTighterThanAnyInfix(t *testing.T) {
	rules := []stringParsing.LexerRule{
		{Type: "DIGIT", Pattern: regexp2.MustCompile(`\d`, 0)},
		{Type: "MINUS", Pattern: regexp2.MustCompile(`-`, 0)},
		{Type: "HIGH", Pattern: regexp2.MustCompile(`@`, 0)},
	}
	for _, precedence := range []int{1, 99, 100, 200} {
		pratt := &PrattExpr{
			Atom:     TokenExpr{TokenType: "DIGIT"},
			Prefixes: map[string]Expr{"MINUS": TokenExpr{TokenType: "MINUS"}},
			Infixes:  map[string]InfixInfo{"HIGH": {Precedence: precedence, Assoc: LeftAssoc}},
		}
		g := Grammar{"start": {Name: "start", Expr: pratt}}
		p := NewParser(stringParsing.NewLexer(rules, nil), g, "start", nil)

		nodes, err := p.Parse("-1@2")
		if err != nil {
			t.Fatalf("precedence %d: unexpected error: %v", precedence, err)
		}
		root := onlyChild(t, nodes)
		if root.Switch != "BinaryOp" {
			t.Errorf("precedence %d: root = %q, want BinaryOp with the prefix on the left", precedence, root.Switch)
			continue
		}
		left, ok := root.Metadata[MetaLeft].(stringParsing.ParsedNode)
		if !ok || left.Switch != "PrefixOp" {
			t.Errorf("precedence %d: the prefix must bind tighter than the infix, got left %+v", precedence, left)
		}
	}
}

func TestLexerErrorKeepsPosition(t *testing.T) {
	g := Grammar{"start": {Name: "start", Expr: TokenExpr{TokenType: "NUMBER"}}}
	p := NewParser(newCalcLexer(), g, "start", []string{"WHITESPACE"})
	_, err := p.Parse("12 $ 34")
	if err == nil {
		t.Fatal("expected a lexer error")
	}
	pe, ok := AsParseError(err)
	if !ok {
		t.Fatalf("error is not a *ParseError: %T", err)
	}
	if pe.Phase != PhaseLexer {
		t.Errorf("phase = %q, want %q", pe.Phase, PhaseLexer)
	}
	if pe.TokenPos == "" {
		t.Errorf("the lexer position must survive wrapping, got %q: %v", pe.TokenPos, pe)
	}
	if !strings.Contains(pe.TokenPos, "col 4") {
		t.Errorf("TokenPos = %q, want the column of the offending character", pe.TokenPos)
	}
	if pe.GetMsg() == "" {
		t.Error("a lexer error must still report a message")
	}
}

func TestActionRejectionSurvivesRepeat(t *testing.T) {
	g := Grammar{"start": {Name: "start", Expr: RepeatExpr{Expr: ActionExpr{
		Expr: TokenExpr{TokenType: "DIGIT"},
		Action: func(nodes []stringParsing.ParsedNode) (stringParsing.ParsedNode, core.ErrorInterface) {
			if nodes[0].Raw == "9" {
				return stringParsing.ParsedNode{}, &AdapterError{Msg: "digit 9 is reserved"}
			}
			return nodes[0], nil
		},
	}}}}
	p := NewParser(digitLexer(), g, "start", nil)
	_, err := p.Parse("129")
	if err == nil {
		t.Fatal("a rejection inside a repetition must abort the parse, not end the repetition")
	}
	if !strings.Contains(err.Error(), "reserved") {
		t.Errorf("the action error must survive the repetition, got %v", err)
	}

	if _, err := p.Parse("123"); err != nil {
		t.Errorf("valid digits must still parse: %v", err)
	}
}

func TestActionRejectionSurvivesSeparatedRepeat(t *testing.T) {
	rules := []stringParsing.LexerRule{
		{Type: "DIGIT", Pattern: regexp2.MustCompile(`\d`, 0)},
		{Type: "COMMA", Pattern: regexp2.MustCompile(`,`, 0)},
	}
	g := Grammar{"start": {Name: "start", Expr: SeparatedRepeatExpr{
		Element: ActionExpr{
			Expr: TokenExpr{TokenType: "DIGIT"},
			Action: func(nodes []stringParsing.ParsedNode) (stringParsing.ParsedNode, core.ErrorInterface) {
				if nodes[0].Raw == "9" {
					return stringParsing.ParsedNode{}, &AdapterError{Msg: "digit 9 is reserved"}
				}
				return nodes[0], nil
			},
		},
		Sep: "COMMA",
	}}}
	p := NewParser(stringParsing.NewLexer(rules, nil), g, "start", nil)
	_, err := p.Parse("1,2,9")
	if err == nil {
		t.Fatal("a rejection inside a separated repetition must abort the parse")
	}
	if !strings.Contains(err.Error(), "reserved") {
		t.Errorf("the action error must survive, got %v", err)
	}
}

func BenchmarkParseExpression(b *testing.B) {
	p := newCalcParser()
	code := "1 + 2 * (3 - 4) / 5 + 6 * 7"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := p.Parse(code); err != nil {
			b.Fatal(err)
		}
	}
}

// ambiguousGrammar reuses the same rules from several alternatives, which is
// the shape that makes a packrat-style memo worthwhile.
func ambiguousGrammar() Grammar {
	return Grammar{
		"start": {Name: "start", Expr: SequenceExpr{Exprs: []Expr{
			NamedExpr{RuleName: "pair"},
			RepeatExpr{Expr: SequenceExpr{Exprs: []Expr{
				TokenExpr{TokenType: "COMMA"},
				NamedExpr{RuleName: "pair"},
			}}},
		}}},
		"pair": {Name: "pair", Expr: ChoiceExpr{Alternatives: []Expr{
			SequenceExpr{Exprs: []Expr{
				NamedExpr{RuleName: "atom"},
				NamedExpr{RuleName: "atom"},
			}},
			NamedExpr{RuleName: "atom"},
		}}},
		"atom": {Name: "atom", Expr: TokenExpr{TokenType: "DIGIT"}},
	}
}

func listLexer() *stringParsing.Lexer {
	rules := []stringParsing.LexerRule{
		{Type: "DIGIT", Pattern: regexp2.MustCompile(`\d`, 0)},
		{Type: "COMMA", Pattern: regexp2.MustCompile(`,`, 0)},
	}
	return stringParsing.NewLexer(rules, nil)
}

func BenchmarkAmbiguousGrammar(b *testing.B) {
	p := NewParser(listLexer(), ambiguousGrammar(), "start", nil)
	code := strings.Repeat("1,2,", 12) + "1"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := p.Parse(code); err != nil {
			b.Fatal(err)
		}
	}
}
