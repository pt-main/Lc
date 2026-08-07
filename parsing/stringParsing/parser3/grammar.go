package parser3

import (
	"fmt"

	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/tap/color"
)

type Expr interface {
	Parse(p *Parser) ([]stringParsing.ParsedNode, error)
}

type Rule struct {
	Name string
	Expr Expr
}

type Grammar map[string]Rule

type TokenExpr struct {
	TokenType string
}

func (t TokenExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, error) {
	tok, err := p.Expect(t.TokenType)
	if err != nil {
		return nil, fmt.Errorf(
			"%s TokenExpr([?YW]%s[?RT]): %w",
			color.Set("[?RD]grammar[?RT]"),
			t.TokenType,
			err,
		)
	}
	return []stringParsing.ParsedNode{tok}, nil
}

type SequenceExpr struct {
	Exprs []Expr
}

func (s SequenceExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, error) {
	var children []stringParsing.ParsedNode
	for i, e := range s.Exprs {
		nodes, err := e.Parse(p)
		if err != nil {
			return nil, fmt.Errorf(
				"%s SequenceExpr element %s/%d: %w",
				color.Set("[?RD]grammar[?RT]"),
				color.Set(fmt.Sprintf("[?YW]%d[?RT]", i+1)),
				len(s.Exprs),
				err,
			)
		}
		children = append(children, nodes...)
	}
	return children, nil
}

type ChoiceExpr struct {
	Alternatives []Expr
}

func (c ChoiceExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, error) {
	savedPos := p.pos
	for i, alt := range c.Alternatives {
		nodes, err := alt.Parse(p)
		if err == nil {
			return nodes, nil
		}
		p.pos = savedPos
		// Продолжаем пробовать альтернативы, не возвращаем ошибку сразу
		_ = i // для возможной диагностики
	}
	return nil, fmt.Errorf(
		"%s ChoiceExpr: no alternative matched at position %s (%d alternatives tried)",
		color.Set("[?RD]grammar[?RT]"),
		color.Set(fmt.Sprintf("[?YW]%d[?RT]", p.pos)),
		len(c.Alternatives),
	)
}

type RepeatExpr struct {
	Expr Expr
	Min  int
}

func (r RepeatExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, error) {
	var all []stringParsing.ParsedNode
	for {
		savedPos := p.pos
		nodes, err := r.Expr.Parse(p)
		if err != nil {
			p.pos = savedPos
			break
		}
		all = append(all, nodes...)
	}
	if len(all) < r.Min {
		return nil, fmt.Errorf(
			"%s RepeatExpr: expected at least %s repetition(s), got %s at position %s",
			color.Set("[?RD]grammar[?RT]"),
			color.Set(fmt.Sprintf("[?YW]%d[?RT]", r.Min)),
			color.Set(fmt.Sprintf("[?YW]%d[?RT]", len(all))),
			color.Set(fmt.Sprintf("[?YW]%d[?RT]", p.pos)),
		)
	}
	return all, nil
}

type OptionalExpr struct {
	Expr Expr
}

func (o OptionalExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, error) {
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

func (n NamedExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, error) {
	rule, ok := p.grammar[n.RuleName]
	if !ok {
		return nil, fmt.Errorf(
			"%s NamedExpr: undefined rule [?YW]'%s'[?RT] (grammar has %d rule(s))",
			color.Set("[?RD]grammar[?RT]"),
			n.RuleName,
			len(p.grammar),
		)
	}
	return rule.Expr.Parse(p)
}

type NodeExpr struct {
	NodeType string
	Expr     Expr
}

func (n NodeExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, error) {
	children, err := n.Expr.Parse(p)
	if err != nil {
		return nil, fmt.Errorf(
			"%s NodeExpr([?YW]%s[?RT]): %w",
			color.Set("[?RD]grammar[?RT]"),
			n.NodeType,
			err,
		)
	}
	var raw string
	for _, child := range children {
		raw += child.Raw
	}
	node := stringParsing.ParsedNode{
		Switch: n.NodeType,
		Raw:    raw,
		Metadata: map[string]interface{}{
			"children": children,
		},
	}
	return []stringParsing.ParsedNode{node}, nil
}

type ActionExpr struct {
	Expr   Expr
	Action func([]stringParsing.ParsedNode) (stringParsing.ParsedNode, error)
}

func (a ActionExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, error) {
	children, err := a.Expr.Parse(p)
	if err != nil {
		return nil, fmt.Errorf(
			"%s ActionExpr: sub-expression failed: %w",
			color.Set("[?RD]grammar[?RT]"),
			err,
		)
	}
	node, err := a.Action(children)
	if err != nil {
		return nil, fmt.Errorf(
			"%s ActionExpr: user action returned error on %d node(s): %w",
			color.Set("[?RD]grammar[?RT]"),
			len(children),
			err,
		)
	}
	return []stringParsing.ParsedNode{node}, nil
}

type NotExpr struct {
	Expr Expr
}

func (n NotExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, error) {
	savedPos := p.pos
	_, err := n.Expr.Parse(p)
	p.pos = savedPos
	if err == nil {
		return nil, fmt.Errorf(
			"%s NotExpr: unexpected match at position %s (expression should not match here)",
			color.Set("[?RD]grammar[?RT]"),
			color.Set(fmt.Sprintf("[?YW]%d[?RT]", p.pos)),
		)
	}
	return []stringParsing.ParsedNode{}, nil
}

type AndExpr struct {
	Expr Expr
}

func (a AndExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, error) {
	savedPos := p.pos
	_, err := a.Expr.Parse(p)
	p.pos = savedPos
	if err != nil {
		return nil, fmt.Errorf(
			"%s AndExpr: expression did not match at position %s: %w",
			color.Set("[?RD]grammar[?RT]"),
			color.Set(fmt.Sprintf("[?YW]%d[?RT]", p.pos)),
			err,
		)
	}
	return []stringParsing.ParsedNode{}, nil
}

type PeekExpr struct {
	TokenType string
}

func (p PeekExpr) Parse(prs *Parser) ([]stringParsing.ParsedNode, error) {
	tok, err := prs.Peek()
	if err != nil {
		return nil, fmt.Errorf(
			"%s PeekExpr([?YW]%s[?RT]): %w",
			color.Set("[?RD]grammar[?RT]"),
			p.TokenType,
			err,
		)
	}
	if tok.Switch != p.TokenType {
		return nil, fmt.Errorf(
			"%s PeekExpr: expected [?YW]'%s'[?RT], got [?YW]'%s'[?RT] (raw: %q) at position %s",
			color.Set("[?RD]grammar[?RT]"),
			p.TokenType,
			tok.Switch,
			tok.Raw,
			color.Set(fmt.Sprintf("[?YW]%d[?RT]", prs.pos)),
		)
	}
	return []stringParsing.ParsedNode{}, nil
}

type SeparatedRepeatExpr struct {
	Element Expr
	Sep     string
	Min     int
	Max     int
}

func (s SeparatedRepeatExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, error) {
	var all []stringParsing.ParsedNode
	count := 0
	firstPos := p.pos
	firstNodes, err := s.Element.Parse(p)
	if err != nil {
		if s.Min == 0 {
			return []stringParsing.ParsedNode{}, nil
		}
		p.pos = firstPos
		return nil, fmt.Errorf(
			"%s SeparatedRepeatExpr: expected at least %s element(s), got none at position %s (separator: [?YW]'%s'[?RT]): %w",
			color.Set("[?RD]grammar[?RT]"),
			color.Set(fmt.Sprintf("[?YW]%d[?RT]", s.Min)),
			color.Set(fmt.Sprintf("[?YW]%d[?RT]", p.pos)),
			s.Sep,
			err,
		)
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
		return nil, fmt.Errorf(
			"%s SeparatedRepeatExpr: expected at least %s elements, got %s at position %s (separator: [?YW]'%s'[?RT])",
			color.Set("[?RD]grammar[?RT]"),
			color.Set(fmt.Sprintf("[?YW]%d[?RT]", s.Min)),
			color.Set(fmt.Sprintf("[?YW]%d[?RT]", count)),
			color.Set(fmt.Sprintf("[?YW]%d[?RT]", p.pos)),
			s.Sep,
		)
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

func (p *PrattExpr) Parse(prs *Parser) ([]stringParsing.ParsedNode, error) {
	node, err := p.parseExpression(prs, 0)
	if err != nil {
		return nil, fmt.Errorf(
			"%s PrattExpr: %w",
			color.Set("[?RD]grammar[?RT]"),
			err,
		)
	}
	return []stringParsing.ParsedNode{node}, nil
}

func (p *PrattExpr) parseExpression(prs *Parser, minPrec int) (stringParsing.ParsedNode, error) {
	var leftNode stringParsing.ParsedNode
	tok, err := prs.Peek()
	if err == nil {
		if _, ok := p.Prefixes[tok.Switch]; ok {
			_, err := prs.Expect(tok.Switch)
			if err != nil {
				return stringParsing.ParsedNode{}, fmt.Errorf(
					"%s PrattExpr prefix op [?YW]'%s'[?RT]: %w",
					color.Set("[?RD]grammar[?RT]"),
					tok.Switch,
					err,
				)
			}
			rightNode, err := p.parseExpression(prs, 100)
			if err != nil {
				return stringParsing.ParsedNode{}, fmt.Errorf(
					"%s PrattExpr prefix op [?YW]'%s'[?RT] right operand: %w",
					color.Set("[?RD]grammar[?RT]"),
					tok.Switch,
					err,
				)
			}
			leftNode = stringParsing.ParsedNode{
				Switch: "PrefixOp",
				Raw:    tok.Raw + rightNode.Raw,
				Metadata: map[string]interface{}{
					"operator": tok.Switch,
					"operand":  rightNode,
				},
			}
		} else {
			atomNodes, err := p.Atom.Parse(prs)
			if err != nil {
				return stringParsing.ParsedNode{}, fmt.Errorf(
					"%s PrattExpr atom at position %s: %w",
					color.Set("[?RD]grammar[?RT]"),
					color.Set(fmt.Sprintf("[?YW]%d[?RT]", prs.pos)),
					err,
				)
			}
			if len(atomNodes) == 0 {
				return stringParsing.ParsedNode{}, fmt.Errorf(
					"%s PrattExpr: atom returned empty at position %s",
					color.Set("[?RD]grammar[?RT]"),
					color.Set(fmt.Sprintf("[?YW]%d[?RT]", prs.pos)),
				)
			}
			if len(atomNodes) == 1 {
				leftNode = atomNodes[0]
			} else {
				var raw string
				for _, n := range atomNodes {
					raw += n.Raw
				}
				leftNode = stringParsing.ParsedNode{
					Switch: "Sequence",
					Raw:    raw,
					Metadata: map[string]interface{}{
						"children": atomNodes,
					},
				}
			}
		}
	} else {
		return stringParsing.ParsedNode{}, fmt.Errorf(
			"%s PrattExpr: unexpected end of input at position %s (need atom or prefix operator)",
			color.Set("[?RD]grammar[?RT]"),
			color.Set(fmt.Sprintf("[?YW]%d[?RT]", prs.pos)),
		)
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
			return stringParsing.ParsedNode{}, fmt.Errorf(
				"%s PrattExpr infix op [?YW]'%s'[?RT] at position %s: %w",
				color.Set("[?RD]grammar[?RT]"),
				nextTok.Switch,
				color.Set(fmt.Sprintf("[?YW]%d[?RT]", prs.pos)),
				err,
			)
		}
		nextMinPrec := infix.Precedence
		if infix.Assoc == LeftAssoc {
			nextMinPrec = infix.Precedence + 1
		}
		rightNode, err := p.parseExpression(prs, nextMinPrec)
		if err != nil {
			return stringParsing.ParsedNode{}, fmt.Errorf(
				"%s PrattExpr infix op [?YW]'%s'[?RT] right operand (minPrec=%s): %w",
				color.Set("[?RD]grammar[?RT]"),
				nextTok.Switch,
				color.Set(fmt.Sprintf("[?YW]%d[?RT]", nextMinPrec)),
				err,
			)
		}
		leftNode = stringParsing.ParsedNode{
			Switch: "BinaryOp",
			Raw:    leftNode.Raw + opTok.Raw + rightNode.Raw,
			Metadata: map[string]interface{}{
				"operator": opTok.Switch,
				"left":     leftNode,
				"right":    rightNode,
			},
		}
	}
	return leftNode, nil
}
