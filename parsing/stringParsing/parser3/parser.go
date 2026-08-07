package parser3

import (
	"fmt"

	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/tap/color"
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
		return nil, fmt.Errorf(
			"%s lexer error: %w",
			color.Set("[?RD]parser3[?RT]"),
			err,
		)
	}
	p.tokens = tokens
	p.pos = 0

	startRule, ok := p.grammar[p.startRule]
	if !ok {
		return nil, fmt.Errorf(
			"%s start rule [?YW]'%s'[?RT] not found in grammar (available: %d rules)",
			color.Set("[?RD]parser3[?RT]"),
			p.startRule,
			len(p.grammar),
		)
	}

	children, err := startRule.Expr.Parse(p)
	if err != nil {
		return nil, fmt.Errorf(
			"%s parse failed at token %d/%d: %w",
			color.Set("[?RD]parser3[?RT]"),
			p.pos,
			len(p.tokens),
			err,
		)
	}

	p.SkipIgnored()
	if p.pos < len(p.tokens) {
		remaining := p.tokens[p.pos:]
		types := make([]string, 0, len(remaining))
		for _, t := range remaining {
			types = append(types, t.Switch)
		}
		return nil, fmt.Errorf(
			"%s unexpected token at position %s: got %s, parsing complete — %d token(s) unconsumed (%v)",
			color.Set("[?RD]parser3[?RT]"),
			color.Set(fmt.Sprintf("[?YW]%d[?RT]", p.pos)),
			color.Set(fmt.Sprintf("[?YW]%s[?RT]", p.tokens[p.pos].Switch)),
			len(remaining),
			types,
		)
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

func (p *Parser) SkipIgnored() {
	for p.pos < len(p.tokens) && p.ignore[p.tokens[p.pos].Switch] {
		p.pos++
	}
}

func (p *Parser) NextToken() (stringParsing.ParsedNode, error) {
	p.SkipIgnored()
	if p.pos >= len(p.tokens) {
		return stringParsing.ParsedNode{}, fmt.Errorf(
			"%s unexpected EOF at position %s (expected more tokens)",
			color.Set("[?RD]parser3[?RT]"),
			color.Set(fmt.Sprintf("[?YW]%d[?RT]", p.pos)),
		)
	}
	tok := p.tokens[p.pos]
	p.pos++
	return tok, nil
}

func (p *Parser) Expect(tokenType string) (stringParsing.ParsedNode, error) {
	tok, err := p.NextToken()
	if err != nil {
		return stringParsing.ParsedNode{}, fmt.Errorf(
			"%s expected token [?YW]'%s'[?RT], but %w",
			color.Set("[?RD]parser3[?RT]"),
			tokenType,
			err,
		)
	}
	if tok.Switch != tokenType {
		return stringParsing.ParsedNode{}, fmt.Errorf(
			"%s expected token [?YW]'%s'[?RT], got [?YW]'%s'[?RT] (raw: %q) at position %s",
			color.Set("[?RD]parser3[?RT]"),
			tokenType,
			tok.Switch,
			tok.Raw,
			color.Set(fmt.Sprintf("[?YW]%d[?RT]", p.pos-1)),
		)
	}
	return tok, nil
}

func (p *Parser) Peek() (stringParsing.ParsedNode, error) {
	p.SkipIgnored()
	if p.pos >= len(p.tokens) {
		return stringParsing.ParsedNode{}, fmt.Errorf(
			"%s EOF at position %s (no more tokens to peek)",
			color.Set("[?RD]parser3[?RT]"),
			color.Set(fmt.Sprintf("[?YW]%d[?RT]", p.pos)),
		)
	}
	return p.tokens[p.pos], nil
}

func (p *Parser) Errorf(format string, args ...interface{}) error {
	return fmt.Errorf(
		"%s %s",
		color.Set("[?RD]parser3[?RT]"),
		fmt.Sprintf(format, args...),
	)
}

func (p *Parser) String() string {
	return "lc/parsing/stringParsing/parser3/Parser"
}
