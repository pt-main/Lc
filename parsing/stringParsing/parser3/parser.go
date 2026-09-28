package parser3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/stringParsing"
)

// Token position metadata keys injected by the lexer.
const (
	metaStart = "__start"
	metaEnd   = "__end"
)

// tokenPos returns a human-readable position string for the token at idx.
func tokenPos(tokens []stringParsing.ParsedNode, idx int) string {
	if idx < 0 || idx >= len(tokens) {
		return "EOF"
	}
	tok := tokens[idx]
	start, hasStart := toInt(tok.Metadata[metaStart])
	end, hasEnd := toInt(tok.Metadata[metaEnd])
	switch {
	case hasStart && hasEnd:
		return "idx=" + strconv.Itoa(idx) + " start=" + strconv.Itoa(start) + "-" + strconv.Itoa(end)
	case hasStart:
		return "idx=" + strconv.Itoa(idx) + " start=" + strconv.Itoa(start)
	default:
		return "idx=" + strconv.Itoa(idx)
	}
}

func toInt(v interface{}) (int, bool) {
	switch n := v.(type) {
	case int:
		return n, true
	case int32:
		return int(n), true
	case int64:
		return int(n), true
	default:
		return 0, false
	}
}

// lexerPosition digs the line and column out of a lexer error so the wrapper
// can point at the offending place instead of reporting an empty position.
func lexerPosition(err error) (int, int, bool) {
	ce, ok := err.(core.ErrorInterface)
	if !ok {
		return 0, 0, false
	}
	meta := ce.GetMeta()
	if meta == nil {
		return 0, 0, false
	}
	line, okLine := toInt(meta[core.EMK(0, "int")])
	col, okCol := toInt(meta[core.EMK(1, "int")])
	return line, col, okLine && okCol
}

// Parser is a recursive-descent parser driven by a Grammar.
// It is not safe for concurrent use: reuse a Parser sequentially or build
// one per goroutine.
type Parser struct {
	lexer     *stringParsing.Lexer
	grammar   Grammar
	startRule string
	ignore    map[string]bool

	tokens []stringParsing.ParsedNode
	pos    int

	// depth bounds rule nesting so a pathological grammar cannot exhaust
	// the goroutine stack.
	depth int

	// activeRules detects a rule that is already being parsed at the same
	// position higher in the stack; entering it again would never terminate.
	// A rule may re-enter at a different position - that is how nested
	// constructs like parenthesised expressions work.
	activeRules map[memoKey]bool

	// memo caches NamedExpr results so a rule referenced from several
	// alternatives is parsed once per position instead of once per path.
	memo map[memoKey]memoEntry

	// deepPos/deepErr remember the furthest a speculative branch ever got.
	// A repeat or optional that stops early is normal, so the failure is
	// kept aside; if the overall parse then leaves input unconsumed, this
	// is the error the author actually wants to see.
	deepPos int
	deepErr core.ErrorInterface
}

// note records an error seen during a speculative branch if it got further
// than anything seen before.
func (p *Parser) note(pos int, err core.ErrorInterface) {
	if err != nil && pos > p.deepPos {
		p.deepPos = pos
		p.deepErr = err
	}
}

// fatal reports whether err describes a broken grammar or a rejected value
// rather than a mismatch of the input. Such an error must not be swallowed by
// backtracking: the alternative that produced it did not merely fail to match,
// it refused the input on purpose.
func fatal(err core.ErrorInterface) bool {
	switch err.(type) {
	case *GrammarError, *AdapterError:
		return true
	default:
		return false
	}
}

type memoKey struct {
	rule string
	pos  int
}

type memoEntry struct {
	nodes []stringParsing.ParsedNode
	end   int // token index just past the rule
}

// maxRuleDepth bounds rule nesting; deep grammars are a configuration error,
// not something to recurse into until the stack blows up.
const maxRuleDepth = 256

// NewParser creates a parser bound to a lexer, a grammar and a start rule.
// Token types listed in ignoreTypes are skipped wherever the parser looks
// for the next token.
func NewParser(lexer *stringParsing.Lexer, grammar Grammar, startRule string, ignoreTypes []string) *Parser {
	ignore := make(map[string]bool, len(ignoreTypes))
	for _, t := range ignoreTypes {
		ignore[t] = true
	}
	return &Parser{
		lexer:       lexer,
		grammar:     grammar,
		startRule:   startRule,
		ignore:      ignore,
		activeRules: make(map[memoKey]bool),
	}
}

// Parse turns code into a single root node named after the start rule.
//
// Err parser3: if the lexer fails, the start rule is missing from the grammar,
// the grammar is recursive, or the input does not match the start rule.
func (p *Parser) Parse(code string, opts ...*parsing.ParseOption) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	if p.lexer == nil {
		return nil, &GrammarError{Phase: "Lexer", Msg: "parser has no lexer"}
	}

	tokens, err := p.lexer.Parse(code, opts...)
	if err != nil {
		wrapped := &ParseError{Phase: PhaseLexer, Cause: err}
		if line, col, ok := lexerPosition(err); ok {
			wrapped.TokenPos = "line " + strconv.Itoa(line) + ", col " + strconv.Itoa(col)
			wrapped.Msg = "the lexer could not tokenize the input"
		}
		return nil, wrapped
	}

	p.tokens = tokens
	p.pos = 0
	p.depth = 0
	p.activeRules = make(map[memoKey]bool, len(p.grammar))
	p.memo = make(map[memoKey]memoEntry, len(p.grammar))
	p.deepPos = 0
	p.deepErr = nil

	rule, ok := p.grammar[p.startRule]
	if !ok {
		return nil, &GrammarError{
			Phase: "StartRule",
			Msg:   fmt.Sprintf("start rule %q is not defined (available: %v)", p.startRule, p.ruleNames()),
		}
	}
	if rule.Expr == nil {
		return nil, &GrammarError{
			Phase: "StartRule",
			Msg:   fmt.Sprintf("start rule %q has a nil expression", p.startRule),
		}
	}

	children, err := rule.Expr.Parse(p)
	if err != nil {
		return nil, p.wrap(err)
	}

	p.SkipIgnored()
	if p.pos < len(p.tokens) {
		// A branch that got further than the final position usually points at
		// the real mistake, so prefer it over "input not fully consumed".
		if p.deepErr != nil && p.deepPos > p.pos {
			return nil, p.deepErr
		}
		return nil, &ParseError{
			Phase:    PhaseEnd,
			TokenIdx: p.pos,
			TokenPos: tokenPos(p.tokens, p.pos),
			Got:      p.tokens[p.pos].Switch,
			Raw:      p.tokens[p.pos].Raw,
			Found:    p.tokenTypesAt(p.pos),
			Msg:      fmt.Sprintf("input not fully consumed from %q", p.tokens[p.pos].Raw),
		}
	}

	var b strings.Builder
	for _, child := range children {
		b.WriteString(child.Raw)
	}
	root := stringParsing.ParsedNode{
		Switch:   p.startRule,
		Raw:      b.String(),
		Metadata: map[string]interface{}{MetaChildren: children},
	}
	return []stringParsing.ParsedNode{root}, nil
}

// wrap turns an inner expression failure into a top-level ParseError that
// keeps the position where parsing actually stopped.
func (p *Parser) wrap(err core.ErrorInterface) core.ErrorInterface {
	if pe, ok := AsParseError(err); ok {
		if pe.TokenIdx == 0 && pe.TokenPos == "" {
			pe.TokenIdx = p.pos
			pe.TokenPos = tokenPos(p.tokens, p.pos)
		}
		return pe
	}
	return &ParseError{
		Phase:    PhaseStart,
		TokenIdx: p.pos,
		TokenPos: tokenPos(p.tokens, p.pos),
		Cause:    err,
	}
}

// ruleNames returns the defined rule names in a stable order so error
// messages do not change between runs.
func (p *Parser) ruleNames() []string {
	names := make([]string, 0, len(p.grammar))
	for k := range p.grammar {
		names = append(names, k)
	}
	sortStrings(names)
	return names
}

// SkipIgnored advances past every ignored token at the current position.
func (p *Parser) SkipIgnored() {
	for p.pos < len(p.tokens) && p.ignore[p.tokens[p.pos].Switch] {
		p.pos++
	}
}

// NextToken returns the next significant token and consumes it.
func (p *Parser) NextToken() (stringParsing.ParsedNode, error) {
	p.SkipIgnored()
	if p.pos >= len(p.tokens) {
		return stringParsing.ParsedNode{}, &ParseError{
			Phase:    PhaseEOF,
			TokenIdx: p.pos,
			TokenPos: tokenPos(p.tokens, p.pos),
			Msg:      "unexpected end of input",
		}
	}
	tok := p.tokens[p.pos]
	p.pos++
	return tok, nil
}

// Expect consumes the next token and fails unless it has the given type.
func (p *Parser) Expect(tokenType string) (stringParsing.ParsedNode, core.ErrorInterface) {
	tok, err := p.NextToken()
	if err != nil {
		return stringParsing.ParsedNode{}, &ParseError{
			Phase:    PhaseExpect,
			Expected: tokenType,
			TokenIdx: p.pos,
			TokenPos: tokenPos(p.tokens, p.pos),
			Cause:    err,
		}
	}
	if tok.Switch != tokenType {
		return stringParsing.ParsedNode{}, &ParseError{
			Phase:    PhaseExpect,
			Expected: tokenType,
			Got:      tok.Switch,
			Raw:      tok.Raw,
			TokenIdx: p.pos - 1,
			TokenPos: tokenPos(p.tokens, p.pos-1),
		}
	}
	return tok, nil
}

// Peek returns the next significant token without consuming it.
func (p *Parser) Peek() (stringParsing.ParsedNode, error) {
	p.SkipIgnored()
	if p.pos >= len(p.tokens) {
		return stringParsing.ParsedNode{}, &ParseError{
			Phase:    PhaseEOF,
			TokenIdx: p.pos,
			TokenPos: tokenPos(p.tokens, p.pos),
			Msg:      "no more tokens",
		}
	}
	return p.tokens[p.pos], nil
}

func (p *Parser) String() string {
	return "lc/parsing/stringParsing/parser3/Parser"
}
