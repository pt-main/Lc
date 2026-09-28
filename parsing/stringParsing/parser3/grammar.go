package parser3

import (
	"fmt"
	"strings"

	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/parsing/stringParsing"
)

// ParsedNode metadata keys produced by the parser.
const (
	MetaChildren = "children"
	MetaOperator = "operator"
	MetaLeft     = "left"
	MetaRight    = "right"
	MetaOperand  = "operand"
)

// Expr is one grammar construct. Every construct must consume tokens or fail;
// an Expr that matches without consuming and can succeed forever makes the
// enclosing repeat loop stop.
type Expr interface {
	Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface)
}

// Rule is a named grammar production.
type Rule struct {
	Name string
	Expr Expr
}

// Grammar maps a rule name to its production.
type Grammar map[string]Rule

func joinRaw(nodes []stringParsing.ParsedNode) string {
	var b strings.Builder
	for i := range nodes {
		b.WriteString(nodes[i].Raw)
	}
	return b.String()
}

// TokenExpr consumes one token of the given type.
type TokenExpr struct {
	TokenType string
}

func (t TokenExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	tok, err := p.Expect(t.TokenType)
	if err != nil {
		return nil, err
	}
	return []stringParsing.ParsedNode{tok}, nil
}

// SequenceExpr matches each element in order.
type SequenceExpr struct {
	Exprs []Expr
}

func (s SequenceExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	children := make([]stringParsing.ParsedNode, 0, len(s.Exprs))
	for _, e := range s.Exprs {
		if e == nil {
			return nil, &GrammarError{Phase: "SequenceExpr", Msg: "sequence has a nil element"}
		}
		nodes, err := e.Parse(p)
		if err != nil {
			return nil, err
		}
		children = append(children, nodes...)
	}
	return children, nil
}

// ChoiceExpr tries each alternative in order and returns the first success.
// On failure it reports the token types actually present at the position the
// alternatives reached furthest, which is far more useful than "nothing
// matched".
type ChoiceExpr struct {
	Alternatives []Expr
}

func (c ChoiceExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	if len(c.Alternatives) == 0 {
		return nil, &GrammarError{Phase: "ChoiceExpr", Msg: "choice has no alternatives"}
	}

	startPos := p.pos
	bestPos := startPos - 1
	var bestErr core.ErrorInterface

	for _, alt := range c.Alternatives {
		if alt == nil {
			return nil, &GrammarError{Phase: "ChoiceExpr", Msg: "choice has a nil alternative"}
		}
		p.pos = startPos
		nodes, err := alt.Parse(p)
		if err == nil {
			return nodes, nil
		}
		if fatal(err) {
			return nil, err
		}
		if p.pos > bestPos {
			bestPos = p.pos
			bestErr = err
		}
		p.note(p.pos, err)
	}

	p.pos = startPos
	return nil, p.noAlternative(startPos, bestPos, bestErr, len(c.Alternatives))
}

// noAlternative builds the failure for a choice that matched nothing.
// The alternative that got furthest is usually the one the author meant, so
// its error is preferred over a generic "nothing matched".
func (p *Parser) noAlternative(startPos, bestPos int, bestErr core.ErrorInterface, tried int) core.ErrorInterface {
	if bestErr != nil {
		return bestErr
	}

	found := p.tokenTypesAt(startPos)
	return &ParseError{
		Phase:    PhasePeek,
		TokenIdx: startPos,
		TokenPos: tokenPos(p.tokens, startPos),
		Got:      p.tokenTypeAt(startPos),
		Raw:      p.tokenRawAt(startPos),
		Found:    found,
		Msg:      fmt.Sprintf("no alternative matched (%d tried)", tried),
	}
}

// tokenTypeAt returns the significant token type at pos, or "" at EOF.
func (p *Parser) tokenTypeAt(pos int) string {
	tok, ok := p.tokenAt(pos)
	if !ok {
		return ""
	}
	return tok.Switch
}

func (p *Parser) tokenRawAt(pos int) string {
	tok, ok := p.tokenAt(pos)
	if !ok {
		return ""
	}
	return tok.Raw
}

// tokenAt returns the first non-ignored token at or after pos.
func (p *Parser) tokenAt(pos int) (stringParsing.ParsedNode, bool) {
	for i := pos; i < len(p.tokens); i++ {
		if !p.ignore[p.tokens[i].Switch] {
			return p.tokens[i], true
		}
	}
	return stringParsing.ParsedNode{}, false
}

// tokenTypesAt lists the significant token types at pos, most frequent first.
// It is used to build "found: NUMBER, PLUS" hints for error messages.
func (p *Parser) tokenTypesAt(pos int) []string {
	tok, ok := p.tokenAt(pos)
	if !ok {
		return nil
	}
	out := []string{tok.Switch}
	limit := 8
	if limit > len(p.tokens)-pos {
		limit = len(p.tokens) - pos
	}
	for i := pos + 1; i < pos+limit; i++ {
		t, ok := p.tokenAt(i)
		if !ok {
			break
		}
		if !containsString(out, t.Switch) {
			out = append(out, t.Switch)
		}
	}
	return out
}

func containsString(list []string, s string) bool {
	for i := range list {
		if list[i] == s {
			return true
		}
	}
	return false
}

// RepeatExpr repeats Expr between Min and Max times (Max<=0 means unlimited).
// It no longer treats "matched zero times" as an error when Min is 0, and it
// stops cleanly when an iteration makes no progress.
type RepeatExpr struct {
	Expr Expr
	Min  int
	Max  int
}

func (r RepeatExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	var all []stringParsing.ParsedNode
	count := 0

	for {
		if r.Max > 0 && count >= r.Max {
			break
		}
		savedPos := p.pos
		nodes, err := r.Expr.Parse(p)
		if err != nil {
			// A rejected value or a broken grammar is not a normal end of the
			// repetition: repeating cannot fix it, so it must surface.
			if fatal(err) {
				return nil, err
			}
			// p.pos is rolled back below, so an error found past savedPos
			// belongs to an abandoned branch and must not become deepErr.
			if p.pos == savedPos {
				p.note(p.pos, err)
			}
			p.pos = savedPos
			break
		}
		if p.pos == savedPos {
			p.pos = savedPos
			break
		}
		all = append(all, nodes...)
		count++
	}

	if count < r.Min {
		return nil, &ParseError{
			Phase:    "RepeatExpr",
			Expected: fmt.Sprintf("at least %d repetition(s)", r.Min),
			TokenIdx: p.pos,
			TokenPos: tokenPos(p.tokens, p.pos),
			Got:      p.tokenTypeAt(p.pos),
			Found:    p.tokenTypesAt(p.pos),
			Msg:      fmt.Sprintf("repetition stopped after %d", count),
		}
	}
	return all, nil
}

// OptionalExpr matches Expr at most once and never fails.
type OptionalExpr struct {
	Expr Expr
}

func (o OptionalExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	savedPos := p.pos
	nodes, err := o.Expr.Parse(p)
	if err != nil || p.pos == savedPos {
		// The branch is abandoned, so only a failure at the start position
		// describes where the parser actually is.
		if p.pos == savedPos {
			p.note(p.pos, err)
		}
		p.pos = savedPos
		return nil, nil
	}
	return nodes, nil
}

// NamedExpr refers to another rule. Results are memoized per position and
// direct left recursion is rejected instead of hanging the process.
type NamedExpr struct {
	RuleName string
}

func (n NamedExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	rule, ok := p.grammar[n.RuleName]
	if !ok {
		return nil, &GrammarError{
			Phase: "NamedExpr",
			Msg:   fmt.Sprintf("undefined rule %q (defined: %v)", n.RuleName, p.ruleNames()),
		}
	}
	if rule.Expr == nil {
		return nil, &GrammarError{
			Phase: "NamedExpr",
			Msg:   fmt.Sprintf("rule %q has a nil expression", n.RuleName),
		}
	}

	key := memoKey{rule: n.RuleName, pos: p.pos}
	if entry, ok := p.memo[key]; ok {
		// Replaying a cached rule must also replay the tokens it consumed.
		p.pos = entry.end
		return entry.nodes, nil
	}

	if p.activeRules[key] {
		return nil, &GrammarError{
			Phase: "NamedExpr",
			Msg:   fmt.Sprintf("rule %q is left-recursive at %s; add an optional base case", n.RuleName, tokenPos(p.tokens, p.pos)),
		}
	}
	if p.depth >= maxRuleDepth {
		return nil, &GrammarError{
			Phase: "NamedExpr",
			Msg:   fmt.Sprintf("rule nesting exceeds %d levels at %q", maxRuleDepth, n.RuleName),
		}
	}

	p.activeRules[key] = true
	p.depth++
	nodes, err := rule.Expr.Parse(p)
	p.depth--
	delete(p.activeRules, key)

	if err == nil {
		p.memo[key] = memoEntry{nodes: nodes, end: p.pos}
	}
	return nodes, err
}

// NodeExpr wraps the result of Expr into a node of type NodeType.
type NodeExpr struct {
	NodeType string
	Expr     Expr
}

func (n NodeExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	children, err := n.Expr.Parse(p)
	if err != nil {
		return nil, err
	}
	node := stringParsing.ParsedNode{
		Switch:   n.NodeType,
		Raw:      joinRaw(children),
		Metadata: map[string]interface{}{MetaChildren: children},
	}
	return []stringParsing.ParsedNode{node}, nil
}

// ActionExpr runs Action on the nodes produced by Expr. If Action is nil the
// child nodes are returned unchanged.
type ActionExpr struct {
	Expr   Expr
	Action func([]stringParsing.ParsedNode) (stringParsing.ParsedNode, core.ErrorInterface)
}

func (a ActionExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	children, err := a.Expr.Parse(p)
	if err != nil {
		return nil, err
	}
	if a.Action == nil {
		return children, nil
	}
	node, err := a.Action(children)
	if err != nil {
		return nil, &AdapterError{Msg: "user action failed", Cause: err}
	}
	return []stringParsing.ParsedNode{node}, nil
}

// NotExpr succeeds when Expr does not match, consuming nothing.
type NotExpr struct {
	Expr Expr
}

func (n NotExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	savedPos := p.pos
	_, err := n.Expr.Parse(p)
	p.pos = savedPos
	if err == nil {
		return nil, &ParseError{
			Phase:    "NotExpr",
			TokenIdx: p.pos,
			TokenPos: tokenPos(p.tokens, p.pos),
			Got:      p.tokenTypeAt(p.pos),
			Raw:      p.tokenRawAt(p.pos),
			Msg:      "token was not expected here",
		}
	}
	return nil, nil
}

// AndExpr succeeds when Expr matches, but consumes nothing.
type AndExpr struct {
	Expr Expr
}

func (a AndExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	savedPos := p.pos
	_, err := a.Expr.Parse(p)
	p.pos = savedPos
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// PeekExpr succeeds when the next significant token has TokenType, and
// consumes nothing.
type PeekExpr struct {
	TokenType string
}

func (pk PeekExpr) Parse(prs *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	tok, err := prs.Peek()
	if err != nil {
		return nil, &ParseError{
			Phase:    PhasePeek,
			Expected: pk.TokenType,
			TokenIdx: prs.pos,
			TokenPos: tokenPos(prs.tokens, prs.pos),
			Cause:    err,
		}
	}
	if tok.Switch != pk.TokenType {
		return nil, &ParseError{
			Phase:    PhasePeek,
			Expected: pk.TokenType,
			Got:      tok.Switch,
			Raw:      tok.Raw,
			TokenIdx: prs.pos,
			TokenPos: tokenPos(prs.tokens, prs.pos),
		}
	}
	return nil, nil
}

// SeparatedRepeatExpr matches Element (Sep Element)* with Min and Max bounds.
// A trailing separator without an element after it is always an error, not a
// silently accepted match.
type SeparatedRepeatExpr struct {
	Element Expr
	Sep     string
	Min     int
	Max     int
}

func (s SeparatedRepeatExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	var all []stringParsing.ParsedNode
	count := 0

	for {
		if s.Max > 0 && count >= s.Max {
			break
		}
		iterPos := p.pos
		if count > 0 {
			if _, err := p.Expect(s.Sep); err != nil {
				p.pos = iterPos
				break
			}
		}
		nodes, err := s.Element.Parse(p)
		if err != nil {
			if fatal(err) {
				return nil, err
			}
			if count == 0 && s.Min > 0 {
				return nil, &ParseError{
					Phase:    "SeparatedRepeatExpr",
					Expected: fmt.Sprintf("at least %d element(s) separated by %q", s.Min, s.Sep),
					TokenIdx: p.pos,
					TokenPos: tokenPos(p.tokens, p.pos),
					Got:      p.tokenTypeAt(p.pos),
					Found:    p.tokenTypesAt(p.pos),
					Cause:    err,
				}
			}
			p.pos = iterPos
			break
		}
		if p.pos == iterPos {
			p.pos = iterPos
			break
		}
		all = append(all, nodes...)
		count++
	}

	if count < s.Min {
		return nil, &ParseError{
			Phase:    "SeparatedRepeatExpr",
			Expected: fmt.Sprintf("at least %d element(s) separated by %q", s.Min, s.Sep),
			TokenIdx: p.pos,
			TokenPos: tokenPos(p.tokens, p.pos),
			Got:      p.tokenTypeAt(p.pos),
			Found:    p.tokenTypesAt(p.pos),
			Msg:      fmt.Sprintf("got %d", count),
		}
	}
	return all, nil
}

// Associativity selects how an infix operator groups equal-precedence operands.
type Associativity int

const (
	LeftAssoc Associativity = iota
	RightAssoc
	NonAssoc
)

// InfixInfo is the precedence and associativity of an infix operator.
type InfixInfo struct {
	Precedence int
	Assoc      Associativity
}

// PrattExpr builds an expression tree honouring operator precedence.
// Prefixes bind tighter than any infix operator.
type PrattExpr struct {
	Atom     Expr
	Prefixes map[string]Expr
	Infixes  map[string]InfixInfo
}

const maxPrattDepth = 512

func (p *PrattExpr) Parse(prs *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	if p.Atom == nil {
		return nil, &GrammarError{Phase: "PrattExpr", Msg: "PrattExpr has no atom expression"}
	}
	node, err := p.parseExpression(prs, 0, 0, p.prefixPrecedence())
	if err != nil {
		return nil, err
	}
	return []stringParsing.ParsedNode{node}, nil
}

// prefixPrecedence returns a floor above every infix the table declares, so a
// prefix operand never swallows a following operator regardless of how large
// an InfixInfo.Precedence the author chose.
func (p *PrattExpr) prefixPrecedence() int {
	highest := 0
	for _, infix := range p.Infixes {
		if infix.Precedence > highest {
			highest = infix.Precedence
		}
	}
	return highest + 1
}

func (p *PrattExpr) parseExpression(prs *Parser, minPrec, depth, prefixPrec int) (stringParsing.ParsedNode, core.ErrorInterface) {
	if depth > maxPrattDepth {
		return stringParsing.ParsedNode{}, &GrammarError{
			Phase: "PrattExpr",
			Msg:   fmt.Sprintf("expression nesting exceeds %d levels", maxPrattDepth),
		}
	}

	left, err := p.parsePrefix(prs, depth, prefixPrec)
	if err != nil {
		return stringParsing.ParsedNode{}, err
	}

	// blockedAt is the precedence of a non-associative operator that was just
	// consumed. Another operator of the same precedence here is a syntax
	// error, so "a < b < c" is rejected instead of silently grouped.
	blockedAt := -1

	for {
		next, perr := prs.Peek()
		if perr != nil {
			return left, nil
		}
		infix, ok := p.Infixes[next.Switch]
		if !ok || infix.Precedence < minPrec {
			return left, nil
		}

		if infix.Assoc == NonAssoc && infix.Precedence == blockedAt {
			return stringParsing.ParsedNode{}, &ParseError{
				Phase:    "PrattExpr",
				Expected: "end of expression",
				Got:      next.Switch,
				Raw:      next.Raw,
				TokenIdx: prs.pos,
				TokenPos: tokenPos(prs.tokens, prs.pos),
				Msg:      "operator is non-associative and cannot repeat here",
			}
		}

		opTok, eerr := prs.Expect(next.Switch)
		if eerr != nil {
			return stringParsing.ParsedNode{}, eerr
		}

		nextMinPrec := infix.Precedence
		blockedAt = -1
		switch infix.Assoc {
		case LeftAssoc:
			nextMinPrec = infix.Precedence + 1
		case NonAssoc:
			nextMinPrec = infix.Precedence + 1
			blockedAt = infix.Precedence
		}

		right, rerr := p.parseExpression(prs, nextMinPrec, depth+1, prefixPrec)
		if rerr != nil {
			return stringParsing.ParsedNode{}, rerr
		}

		left = stringParsing.ParsedNode{
			Switch: "BinaryOp",
			Raw:    left.Raw + opTok.Raw + right.Raw,
			Metadata: map[string]interface{}{
				MetaOperator: opTok.Switch,
				MetaLeft:     left,
				MetaRight:    right,
			},
		}
	}
}

func (p *PrattExpr) parsePrefix(prs *Parser, depth, prefixPrec int) (stringParsing.ParsedNode, core.ErrorInterface) {
	tok, perr := prs.Peek()
	if perr != nil {
		return stringParsing.ParsedNode{}, &ParseError{
			Phase:    PhaseEOF,
			TokenIdx: prs.pos,
			TokenPos: tokenPos(prs.tokens, prs.pos),
			Msg:      "expected an operand or prefix operator",
		}
	}

	if _, ok := p.Prefixes[tok.Switch]; ok {
		if _, err := prs.Expect(tok.Switch); err != nil {
			return stringParsing.ParsedNode{}, err
		}
		operand, err := p.parseExpression(prs, prefixPrec, depth+1, prefixPrec)
		if err != nil {
			return stringParsing.ParsedNode{}, err
		}
		return stringParsing.ParsedNode{
			Switch: "PrefixOp",
			Raw:    tok.Raw + operand.Raw,
			Metadata: map[string]interface{}{
				MetaOperator: tok.Switch,
				MetaOperand:  operand,
			},
		}, nil
	}

	atomNodes, err := p.Atom.Parse(prs)
	if err != nil {
		return stringParsing.ParsedNode{}, err
	}
	switch len(atomNodes) {
	case 0:
		return stringParsing.ParsedNode{}, &ParseError{
			Phase:    "PrattExpr",
			TokenIdx: prs.pos,
			TokenPos: tokenPos(prs.tokens, prs.pos),
			Got:      prs.tokenTypeAt(prs.pos),
			Msg:      "atom matched no tokens",
		}
	case 1:
		return atomNodes[0], nil
	default:
		return stringParsing.ParsedNode{
			Switch:   "Sequence",
			Raw:      joinRaw(atomNodes),
			Metadata: map[string]interface{}{MetaChildren: atomNodes},
		}, nil
	}
}
