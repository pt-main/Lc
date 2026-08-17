package parser3

import (
	"fmt"
	"strings"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/stringParsing"
)

type Expr interface {
	Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface)
}

type Rule struct {
	Name string
	Expr Expr
}

type Grammar map[string]Rule

type TokenExpr struct {
	TokenType string
}

func (t TokenExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	tok, err := p.Expect(t.TokenType)
	if err != nil {
		return nil, &GrammarError{
			Code:  "TokenExpr",
			Msg:   fmt.Sprintf("expected token '%s'", t.TokenType),
			Cause: err,
		}
	}
	return []stringParsing.ParsedNode{tok}, nil
}

type SequenceExpr struct {
	Exprs []Expr
}

func (s SequenceExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	var children []stringParsing.ParsedNode
	for i, e := range s.Exprs {
		nodes, err := e.Parse(p)
		if err != nil {
			return nil, &GrammarError{
				Code:  "SequenceExpr",
				Msg:   fmt.Sprintf("element %d/%d failed", i+1, len(s.Exprs)),
				Cause: err,
			}
		}
		children = append(children, nodes...)
	}
	return children, nil
}

type ChoiceExpr struct {
	Alternatives []Expr
}

func (c ChoiceExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	savedPos := p.pos
	for _, alt := range c.Alternatives {
		nodes, err := alt.Parse(p)
		if err == nil {
			return nodes, nil
		}
		p.pos = savedPos
	}
	altNames := make([]string, len(c.Alternatives))
	for i, alt := range c.Alternatives {
		altNames[i] = fmt.Sprintf("%T", alt)
	}
	return nil, &GrammarError{
		Code: "ChoiceExpr",
		Msg:  fmt.Sprintf("no alternative matched at %s (%d alternatives tried: %v)", tokenPos(p.tokens, p.pos), len(c.Alternatives), altNames),
	}
}

type RepeatExpr struct {
	Expr Expr
	Min  int
}

func (r RepeatExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	var all []stringParsing.ParsedNode
	for {
		savedPos := p.pos
		nodes, err := r.Expr.Parse(p)
		if err != nil {
			p.pos = savedPos
			break
		}
		// Guard against infinite loop on zero-width match.
		if p.pos == savedPos {
			p.pos = savedPos
			break
		}
		all = append(all, nodes...)
	}
	if len(all) < r.Min {
		return nil, &GrammarError{
			Code: "RepeatExpr",
			Msg:  fmt.Sprintf("expected at least %d repetition(s), got %d at %s", r.Min, len(all), tokenPos(p.tokens, p.pos)),
		}
	}
	return all, nil
}

type OptionalExpr struct {
	Expr Expr
}

func (o OptionalExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	savedPos := p.pos
	nodes, err := o.Expr.Parse(p)
	if err != nil {
		p.pos = savedPos
		return []stringParsing.ParsedNode{}, nil
	}
	return nodes, nil
}

type NamedExpr struct {
	RuleName string
}

func (n NamedExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	rule, ok := p.grammar[n.RuleName]
	if !ok {
		keys := make([]string, 0, len(p.grammar))
		for k := range p.grammar {
			keys = append(keys, k)
		}
		return nil, &GrammarError{
			Code: "NamedExpr",
			Msg:  fmt.Sprintf("undefined rule '%s' (grammar has %d rule(s): %v)", n.RuleName, len(p.grammar), keys),
		}
	}
	return rule.Expr.Parse(p)
}

type NodeExpr struct {
	NodeType string
	Expr     Expr
}

func (n NodeExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	children, err := n.Expr.Parse(p)
	if err != nil {
		return nil, &GrammarError{
			Code:  "NodeExpr",
			Msg:   fmt.Sprintf("building node '%s'", n.NodeType),
			Cause: err,
		}
	}
	var b strings.Builder
	for _, child := range children {
		b.WriteString(child.Raw)
	}
	node := stringParsing.ParsedNode{
		Switch: n.NodeType,
		Raw:    b.String(),
		Metadata: map[string]interface{}{
			"children": children,
		},
	}
	return []stringParsing.ParsedNode{node}, nil
}

type ActionExpr struct {
	Expr   Expr
	Action func([]stringParsing.ParsedNode) (stringParsing.ParsedNode, core.ErrorInterface)
}

func (a ActionExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	children, err := a.Expr.Parse(p)
	if err != nil {
		return nil, &GrammarError{
			Code:  "ActionExpr",
			Msg:   "sub-expression failed",
			Cause: err,
		}
	}
	node, err := a.Action(children)
	if err != nil {
		return nil, &GrammarError{
			Code:  "ActionExpr",
			Msg:   fmt.Sprintf("user action returned error on %d node(s)", len(children)),
			Cause: err,
		}
	}
	return []stringParsing.ParsedNode{node}, nil
}

type NotExpr struct {
	Expr Expr
}

func (n NotExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	savedPos := p.pos
	_, err := n.Expr.Parse(p)
	p.pos = savedPos
	if err == nil {
		return nil, &GrammarError{
			Code: "NotExpr",
			Msg:  fmt.Sprintf("unexpected match at %s (expression should not match here)", tokenPos(p.tokens, p.pos)),
		}
	}
	return []stringParsing.ParsedNode{}, nil
}

type AndExpr struct {
	Expr Expr
}

func (a AndExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	savedPos := p.pos
	_, err := a.Expr.Parse(p)
	p.pos = savedPos
	if err != nil {
		return nil, &GrammarError{
			Code:  "AndExpr",
			Msg:   fmt.Sprintf("expression did not match at %s", tokenPos(p.tokens, p.pos)),
			Cause: err,
		}
	}
	return []stringParsing.ParsedNode{}, nil
}

type PeekExpr struct {
	TokenType string
}

func (p PeekExpr) Parse(prs *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	tok, err := prs.Peek()
	if err != nil {
		return nil, &GrammarError{
			Code:  "PeekExpr",
			Msg:   fmt.Sprintf("peeking for '%s'", p.TokenType),
			Cause: err,
		}
	}
	if tok.Switch != p.TokenType {
		return nil, &GrammarError{
			Code: "PeekExpr",
			Msg:  fmt.Sprintf("expected '%s', got '%s' (raw: %q) at %s", p.TokenType, tok.Switch, tok.Raw, tokenPos(prs.tokens, prs.pos)),
		}
	}
	return []stringParsing.ParsedNode{}, nil
}

type SeparatedRepeatExpr struct {
	Element Expr
	Sep     string
	Min     int
	Max     int
}

func (s SeparatedRepeatExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	var all []stringParsing.ParsedNode
	count := 0
	firstPos := p.pos
	firstNodes, err := s.Element.Parse(p)
	if err != nil {
		if s.Min == 0 {
			return []stringParsing.ParsedNode{}, nil
		}
		p.pos = firstPos
		return nil, &GrammarError{
			Code:  "SeparatedRepeatExpr",
			Msg:   fmt.Sprintf("expected at least %d element(s), got none at %s (separator: '%s')", s.Min, tokenPos(p.tokens, p.pos), s.Sep),
			Cause: err,
		}
	}
	all = append(all, firstNodes...)
	count++
	for {
		if s.Max > 0 && count >= s.Max {
			break
		}
		savedPos := p.pos
		_, err := p.Expect(s.Sep)
		if err != nil {
			p.pos = savedPos
			break
		}
		nodes, err := s.Element.Parse(p)
		if err != nil {
			p.pos = savedPos
			break
		}
		all = append(all, nodes...)
		count++
	}
	if count < s.Min {
		return nil, &GrammarError{
			Code: "SeparatedRepeatExpr",
			Msg:  fmt.Sprintf("expected at least %d elements, got %d at %s (separator: '%s')", s.Min, count, tokenPos(p.tokens, p.pos), s.Sep),
		}
	}
	return all, nil
}

type Associativity int

const (
	LeftAssoc Associativity = iota
	RightAssoc
	NonAssoc
)

type InfixInfo struct {
	Precedence int
	Assoc      Associativity
}

type PrattExpr struct {
	Atom     Expr
	Prefixes map[string]Expr
	Infixes  map[string]InfixInfo
}

const maxPrefixPrecedence = 100

func (p *PrattExpr) Parse(prs *Parser) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	node, err := p.parseExpression(prs, 0)
	if err != nil {
		return nil, &GrammarError{
			Code:  "PrattExpr",
			Msg:   "expression parsing failed",
			Cause: err,
		}
	}
	return []stringParsing.ParsedNode{node}, nil
}

func (p *PrattExpr) parseExpression(prs *Parser, minPrec int) (stringParsing.ParsedNode, core.ErrorInterface) {
	var leftNode stringParsing.ParsedNode
	tok, err := prs.Peek()
	if err == nil {
		if _, ok := p.Prefixes[tok.Switch]; ok {
			_, err := prs.Expect(tok.Switch)
			if err != nil {
				return stringParsing.ParsedNode{}, &GrammarError{
					Code:  "PrattExpr",
					Msg:   fmt.Sprintf("prefix operator '%s'", tok.Switch),
					Cause: err,
				}
			}
			rightNode, err := p.parseExpression(prs, maxPrefixPrecedence)
			if err != nil {
				return stringParsing.ParsedNode{}, &GrammarError{
					Code:  "PrattExpr",
					Msg:   fmt.Sprintf("prefix operator '%s' right operand", tok.Switch),
					Cause: err,
				}
			}
			var b strings.Builder
			b.WriteString(tok.Raw)
			b.WriteString(rightNode.Raw)
			leftNode = stringParsing.ParsedNode{
				Switch: "PrefixOp",
				Raw:    b.String(),
				Metadata: map[string]interface{}{
					"operator": tok.Switch,
					"operand":  rightNode,
				},
			}
		} else {
			atomNodes, err := p.Atom.Parse(prs)
			if err != nil {
				return stringParsing.ParsedNode{}, &GrammarError{
					Code:  "PrattExpr",
					Msg:   fmt.Sprintf("atom at %s", tokenPos(prs.tokens, prs.pos)),
					Cause: err,
				}
			}
			if len(atomNodes) == 0 {
				return stringParsing.ParsedNode{}, &GrammarError{
					Code: "PrattExpr",
					Msg:  fmt.Sprintf("atom returned empty at %s", tokenPos(prs.tokens, prs.pos)),
				}
			}
			if len(atomNodes) == 1 {
				leftNode = atomNodes[0]
			} else {
				var b strings.Builder
				for _, n := range atomNodes {
					b.WriteString(n.Raw)
				}
				leftNode = stringParsing.ParsedNode{
					Switch: "Sequence",
					Raw:    b.String(),
					Metadata: map[string]interface{}{
						"children": atomNodes,
					},
				}
			}
		}
	} else {
		return stringParsing.ParsedNode{}, &GrammarError{
			Code: "PrattExpr",
			Msg:  fmt.Sprintf("unexpected end of input at %s (need atom or prefix operator)", tokenPos(prs.tokens, prs.pos)),
		}
	}
	for {
		nextTok, err := prs.Peek()
		if err != nil {
			break
		}
		infix, ok := p.Infixes[nextTok.Switch]
		if !ok {
			break
		}
		if infix.Precedence < minPrec {
			break
		}
		opTok, err := prs.Expect(nextTok.Switch)
		if err != nil {
			return stringParsing.ParsedNode{}, &GrammarError{
				Code:  "PrattExpr",
				Msg:   fmt.Sprintf("infix operator '%s' at %s", nextTok.Switch, tokenPos(prs.tokens, prs.pos)),
				Cause: err,
			}
		}
		nextMinPrec := infix.Precedence
		if infix.Assoc == LeftAssoc {
			nextMinPrec = infix.Precedence + 1
		}
		rightNode, err := p.parseExpression(prs, nextMinPrec)
		if err != nil {
			return stringParsing.ParsedNode{}, &GrammarError{
				Code:  "PrattExpr",
				Msg:   fmt.Sprintf("infix operator '%s' right operand (minPrec=%d)", nextTok.Switch, nextMinPrec),
				Cause: err,
			}
		}
		var b strings.Builder
		b.WriteString(leftNode.Raw)
		b.WriteString(opTok.Raw)
		b.WriteString(rightNode.Raw)
		leftNode = stringParsing.ParsedNode{
			Switch: "BinaryOp",
			Raw:    b.String(),
			Metadata: map[string]interface{}{
				"operator": opTok.Switch,
				"left":     leftNode,
				"right":    rightNode,
			},
		}
	}
	return leftNode, nil
}
