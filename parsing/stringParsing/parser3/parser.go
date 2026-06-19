package parser3

import (
	"fmt"

	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/stringParsing"
)

type Parser struct {
	lexer     *stringParsing.Lexer
	grammar   Grammar
	startRule string
	ignore    map[string]bool

	tokens []stringParsing.ParsedNode
	pos    int
}

func NewParser(lexer *stringParsing.Lexer, grammar Grammar, startRule string, ignoreTypes []string) *Parser {
	ignore := make(map[string]bool)
	for _, t := range ignoreTypes {
		ignore[t] = true
	}
	return &Parser{
		lexer:     lexer,
		grammar:   grammar,
		startRule: startRule,
		ignore:    ignore,
	}
}

func (p *Parser) Parse(code string, opts ...*parsing.ParseOption) ([]stringParsing.ParsedNode, error) {
	tokens, err := p.lexer.Parse(code, opts...)
	if err != nil {
		return nil, err
	}
	p.tokens = tokens
	p.pos = 0

	startRule, ok := p.grammar[p.startRule]
	if !ok {
		return nil, fmt.Errorf("start rule '%s' not found", p.startRule)
	}

	children, err := startRule.Expr.Parse(p)
	if err != nil {
		return nil, err
	}

	p.skipIgnored()
	if p.pos < len(p.tokens) {
		return nil, fmt.Errorf("unexpected token at position %d: %v", p.pos, p.tokens[p.pos])
	}

	root := stringParsing.ParsedNode{
		Switch: p.startRule,
		Raw:    "",
		Metadata: map[string]interface{}{
			"children": children,
		},
	}
	for _, child := range children {
		root.Raw += child.Raw
	}
	return []stringParsing.ParsedNode{root}, nil
}

func (p *Parser) skipIgnored() {
	for p.pos < len(p.tokens) && p.ignore[p.tokens[p.pos].Switch] {
		p.pos++
	}
}

func (p *Parser) nextToken() (stringParsing.ParsedNode, error) {
	p.skipIgnored()
	if p.pos >= len(p.tokens) {
		return stringParsing.ParsedNode{}, fmt.Errorf("unexpected EOF")
	}
	tok := p.tokens[p.pos]
	p.pos++
	return tok, nil
}

func (p *Parser) Expect(tokenType string) (stringParsing.ParsedNode, error) {
	tok, err := p.nextToken()
	if err != nil {
		return stringParsing.ParsedNode{}, err
	}
	if tok.Switch != tokenType {
		return stringParsing.ParsedNode{}, fmt.Errorf("expected %s, got %s", tokenType, tok.Switch)
	}
	return tok, nil
}

func (p *Parser) Peek() (stringParsing.ParsedNode, error) {
	p.skipIgnored()
	if p.pos >= len(p.tokens) {
		return stringParsing.ParsedNode{}, fmt.Errorf("EOF")
	}
	return p.tokens[p.pos], nil
}

func (p *Parser) Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

func (p *Parser) String() string {
	return "lc/parsing/stringParsing/parser3/Parser"
}
