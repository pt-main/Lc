package parser3

import (
	"github.com/pt-main/lc/parsing/stringParsing"
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
		return nil, err
	}
	return []stringParsing.ParsedNode{tok}, nil
}

type SequenceExpr struct {
	Exprs []Expr
}

func (s SequenceExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, error) {
	var children []stringParsing.ParsedNode
	for _, e := range s.Exprs {
		nodes, err := e.Parse(p)
		if err != nil {
			return nil, err
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
	for _, alt := range c.Alternatives {
		nodes, err := alt.Parse(p)
		if err == nil {
			return nodes, nil
		}
		p.pos = savedPos
	}
	return nil, p.Errorf("no alternative matched")
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
		return nil, p.Errorf("repeat: expected at least %d, got %d", r.Min, len(all))
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
		return nil, p.Errorf("undefined rule: %s", n.RuleName)
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
		return nil, err
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
