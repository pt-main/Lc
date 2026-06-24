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

type ActionExpr struct {
	Expr   Expr
	Action func([]stringParsing.ParsedNode) (stringParsing.ParsedNode, error)
}

func (a ActionExpr) Parse(p *Parser) ([]stringParsing.ParsedNode, error) {
	children, err := a.Expr.Parse(p)
	if err != nil {
		return nil, err
	}
	node, err := a.Action(children)
	if err != nil {
		return nil, err
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
		return nil, p.Errorf("not: unexpected match")
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
		return nil, err
	}
	return []stringParsing.ParsedNode{}, nil
}

type PeekExpr struct {
	TokenType string
}

func (p PeekExpr) Parse(prs *Parser) ([]stringParsing.ParsedNode, error) {
	tok, err := prs.Peek()
	if err != nil {
		return nil, err
	}
	if tok.Switch != p.TokenType {
		return nil, prs.Errorf("peek: expected %s, got %s", p.TokenType, tok.Switch)
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
		return nil, p.Errorf("separated repeat: expected at least %d elements", s.Min)
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
		return nil, p.Errorf("separated repeat: expected at least %d elements, got %d", s.Min, count)
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
		return nil, err
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
				return stringParsing.ParsedNode{}, err
			}
			rightNode, err := p.parseExpression(prs, 100)
			if err != nil {
				return stringParsing.ParsedNode{}, err
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
				return stringParsing.ParsedNode{}, err
			}
			if len(atomNodes) == 0 {
				return stringParsing.ParsedNode{}, prs.Errorf("pratt: atom returned empty")
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
		return stringParsing.ParsedNode{}, prs.Errorf("pratt: unexpected end of input")
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
			return stringParsing.ParsedNode{}, err
		}
		nextMinPrec := infix.Precedence
		if infix.Assoc == LeftAssoc {
			nextMinPrec = infix.Precedence + 1
		}
		rightNode, err := p.parseExpression(prs, nextMinPrec)
		if err != nil {
			return stringParsing.ParsedNode{}, err
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
