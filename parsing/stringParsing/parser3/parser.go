package parser3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/parsing"
	"github.com/pt-main/lc/v2/parsing/stringParsing"
)

// Token position metadata keys injected by the lexer.
const (
	metaStart = "__start"
	metaEnd   = "__end"
)

// posAt returns a human-readable position string for the token at idx.
func (p *Parser) posAt(idx int) string {
	if idx < 0 || idx >= len(p.tokens) {
		return "EOF"
	}
	start, hasStart := toInt(p.tokens[idx].Metadata[metaStart])
	if !hasStart {
		return "idx=" + strconv.Itoa(idx)
	}
	end, hasEnd := toInt(p.tokens[idx].Metadata[metaEnd])
	if !hasEnd {
		return "idx=" + strconv.Itoa(idx) + " start=" + strconv.Itoa(start)
	}
	var b strings.Builder
	b.Grow(len("idx= start=-"))
	b.WriteString("idx=")
	b.WriteString(strconv.Itoa(idx))
	b.WriteString(" start=")
	b.WriteString(strconv.Itoa(start))
	b.WriteByte('-')
	b.WriteString(strconv.Itoa(end))
	return b.String()
}

// ignoredAt reports whether the token at idx is of an ignored type. The answer
// is precomputed per token so the hot paths never hash a token type.
func (p *Parser) ignoredAt(idx int) bool {
	return idx < len(p.ignored) && p.ignored[idx]
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

	// ignored marks the tokens the parser must step over. It is derived once
	// per parse so the hot loops never hash a token type.
	ignored []bool

	// depth bounds rule nesting so a pathological grammar cannot exhaust
	// the goroutine stack.
	depth int

	// activeRules detects a rule that is already being parsed at the same
	// position higher in the stack; entering it again would never terminate.
	// A rule may re-enter at a different position - that is how nested
	// constructs like parenthesised expressions work.
	// Keys are rule ids packed with the token index, so no string hashing
	// happens on the recursive path.
	activeRules map[uint64]bool

	// memo caches NamedExpr results so a rule referenced from several
	// alternatives is parsed once per position instead of once per path.
	memo map[uint64]memoEntry

	// deepPos/deepErr remember the furthest a speculative branch ever got.
	// A repeat or optional that stops early is normal, so the failure is
	// kept aside; if the overall parse then leaves input unconsumed, this
	// is the error the author actually wants to see.
	deepPos int
	deepErr core.ErrorInterface

	// ruleIDs maps a rule name to its index in ruleOrder, giving every rule a
	// small integer usable as a map key.
	ruleIDs   map[string]int
	ruleOrder []string
}

// memoKey packs a rule id and a token index into one comparable key.
func memoKeyFor(ruleID, pos int) uint64 {
	return uint64(uint32(ruleID))<<32 | uint64(uint32(pos))
}

// resetCaches empties the memo and recursion tables while keeping their
// allocated buckets, so a reused parser does not reallocate them per parse.
// clear leaves a nil map nil, hence the explicit rebuild.
func (p *Parser) resetCaches() {
	if p.memo == nil {
		p.memo = make(map[uint64]memoEntry, len(p.grammar))
	} else {
		clear(p.memo)
	}
	if p.activeRules == nil {
		p.activeRules = make(map[uint64]bool, len(p.grammar))
	} else {
		clear(p.activeRules)
	}
}

// prepare builds the per-token lookup tables. It reuses the existing slices
// so a reused parser does not reallocate them on every Parse call.
func (p *Parser) prepare(tokens []stringParsing.ParsedNode) {
	if cap(p.ignored) < len(tokens) {
		p.ignored = make([]bool, len(tokens))
	} else {
		p.ignored = p.ignored[:len(tokens)]
	}
	if len(p.ignore) == 0 {
		clear(p.ignored)
		return
	}
	for i := range tokens {
		p.ignored[i] = p.ignore[tokens[i].Switch]
	}
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
	ruleIDs := make(map[string]int, len(grammar))
	ruleOrder := make([]string, 0, len(grammar))
	for name := range grammar {
		ruleIDs[name] = len(ruleOrder)
		ruleOrder = append(ruleOrder, name)
	}
	return &Parser{
		lexer:       lexer,
		grammar:     grammar,
		startRule:   startRule,
		ignore:      ignore,
		ruleIDs:     ruleIDs,
		ruleOrder:   ruleOrder,
		activeRules: make(map[uint64]bool),
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
	p.prepare(tokens)
	p.resetCaches()
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
			TokenPos: p.posAt(p.pos),
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
			pe.TokenPos = p.posAt(p.pos)
		}
		return pe
	}
	return &ParseError{
		Phase:    PhaseStart,
		TokenIdx: p.pos,
		TokenPos: p.posAt(p.pos),
		Cause:    err,
	}
}

// ruleNames returns the defined rule names in a stable order so error
// messages do not change between runs.
func (p *Parser) ruleNames() []string {
	if len(p.ruleOrder) == len(p.grammar) {
		names := make([]string, len(p.ruleOrder))
		copy(names, p.ruleOrder)
		sortStrings(names)
		return names
	}
	names := make([]string, 0, len(p.grammar))
	for k := range p.grammar {
		names = append(names, k)
	}
	sortStrings(names)
	return names
}

// SkipIgnored advances past every ignored token at the current position.
func (p *Parser) SkipIgnored() {
	for p.pos < len(p.ignored) && p.ignored[p.pos] {
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
			TokenPos: p.posAt(p.pos),
			Msg:      "unexpected end of input",
		}
	}
	tok := p.tokens[p.pos]
	p.pos++
	return tok, nil
}

// Expect consumes the next token and fails unless it has the given type.
//
// The common case is a successful match, so the token type is compared before
// any error is built: a mismatch during backtracking discards the error, and
// formatting one there costs an allocation per failed token.
func (p *Parser) Expect(tokenType string) (stringParsing.ParsedNode, core.ErrorInterface) {
	p.SkipIgnored()
	if p.pos < len(p.tokens) && p.tokens[p.pos].Switch == tokenType {
		tok := p.tokens[p.pos]
		p.pos++
		return tok, nil
	}
	return stringParsing.ParsedNode{}, p.expectFailed(tokenType)
}

// expectFailed builds the error for a token that did not match, after
// SkipIgnored has already advanced past any ignored tokens.
func (p *Parser) expectFailed(tokenType string) core.ErrorInterface {
	if p.pos >= len(p.tokens) {
		return &ParseError{
			Phase:    PhaseExpect,
			Expected: tokenType,
			TokenIdx: p.pos,
			TokenPos: p.posAt(p.pos),
			Cause: &ParseError{
				Phase:    PhaseEOF,
				TokenIdx: p.pos,
				TokenPos: p.posAt(p.pos),
				Msg:      "unexpected end of input",
			},
		}
	}
	tok := p.tokens[p.pos]
	p.pos++
	return &ParseError{
		Phase:    PhaseExpect,
		Expected: tokenType,
		Got:      tok.Switch,
		Raw:      tok.Raw,
		TokenIdx: p.pos - 1,
		TokenPos: p.posAt(p.pos - 1),
	}
}

// Peek returns the next significant token without consuming it.
func (p *Parser) Peek() (stringParsing.ParsedNode, error) {
	p.SkipIgnored()
	if p.pos >= len(p.tokens) {
		return stringParsing.ParsedNode{}, &ParseError{
			Phase:    PhaseEOF,
			TokenIdx: p.pos,
			TokenPos: p.posAt(p.pos),
			Msg:      "no more tokens",
		}
	}
	return p.tokens[p.pos], nil
}

func (p *Parser) String() string {
	return "lc/parsing/stringParsing/parser3/Parser"
}
