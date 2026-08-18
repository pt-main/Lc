package parser3

import (
	"fmt"
	"strings"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/stringParsing"
)

// tokenPos returns a human-readable position string for the token at idx.
// If the lexer injected __start/__end metadata, it includes that.
func tokenPos(tokens []stringParsing.ParsedNode, idx int) string {
	if idx < 0 || idx >= len(tokens) {
		return "EOF"
	}
	tok := tokens[idx]
	if start, ok := tok.Metadata["__start"].(int); ok {
		if end, ok2 := tok.Metadata["__end"].(int); ok2 {
			return fmt.Sprintf("idx=%d start=%d-%d", idx, start, end)
		}
		return fmt.Sprintf("idx=%d start=%d", idx, start)
	}
	return fmt.Sprintf("idx=%d", idx)
}

type Parser struct {
	lexer     *stringParsing.Lexer
	grammar   Grammar
	startRule string
	ignore    map[string]bool

	tokens []stringParsing.ParsedNode
	pos    int
}

func NewParser(lexer *stringParsing.Lexer, grammar Grammar, startRule string, ignoreTypes []string) *Parser {
	ignore := make(map[string]bool, len(ignoreTypes))
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

func (p *Parser) Parse(code string, opts ...*parsing.ParseOption) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	tokens, err := p.lexer.Parse(code, opts...)
	if err != nil {
		return nil, &ParseError{
			Code:  "Lexer",
			Cause: err,
		}
	}
	p.tokens = tokens
	p.pos = 0

	startRule, ok := p.grammar[p.startRule]
	if !ok {
		keys := make([]string, 0, len(p.grammar))
		for k := range p.grammar {
			keys = append(keys, k)
		}
		return nil, &GrammarError{
			Code: "StartRule",
			Msg:  fmt.Sprintf("start rule '%s' not found in grammar (available: %v)", p.startRule, keys),
		}
	}

	children, err := startRule.Expr.Parse(p)
	if err != nil {
		return nil, &ParseError{
			Code:     "Parse",
			TokenIdx: p.pos,
			TokenPos: tokenPos(p.tokens, p.pos),
			Cause:    err,
		}
	}

	p.SkipIgnored()
	if p.pos < len(p.tokens) {
		remaining := p.tokens[p.pos:]
		types := make([]string, len(remaining))
		for i, t := range remaining {
			types[i] = t.Switch
		}
		return nil, &ParseError{
			Code:     "UnexpectedToken",
			TokenIdx: p.pos,
			TokenPos: tokenPos(p.tokens, p.pos),
			Got:      p.tokens[p.pos].Switch,
			Raw:      p.tokens[p.pos].Raw,
			Msg:      fmt.Sprintf("parsing complete — %d token(s) unconsumed (%v)", len(remaining), types),
		}
	}

	var b strings.Builder
	for _, child := range children {
		b.WriteString(child.Raw)
	}
	root := stringParsing.ParsedNode{
		Switch: p.startRule,
		Raw:    b.String(),
		Metadata: map[string]interface{}{
			"children": children,
		},
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
		return stringParsing.ParsedNode{}, &ParseError{
			Code:     "NextToken",
			TokenIdx: p.pos,
			TokenPos: tokenPos(p.tokens, p.pos),
			Msg:      "unexpected EOF",
		}
	}
	tok := p.tokens[p.pos]
	p.pos++
	return tok, nil
}

func (p *Parser) Expect(tokenType string) (stringParsing.ParsedNode, core.ErrorInterface) {
	tok, err := p.NextToken()
	if err != nil {
		return stringParsing.ParsedNode{}, &ParseError{
			Code:     "Expect",
			Expected: tokenType,
			TokenIdx: p.pos,
			TokenPos: tokenPos(p.tokens, p.pos),
			Cause:    err,
		}
	}
	if tok.Switch != tokenType {
		return stringParsing.ParsedNode{}, &ParseError{
			Code:     "Expect",
			Expected: tokenType,
			Got:      tok.Switch,
			Raw:      tok.Raw,
			TokenIdx: p.pos - 1,
			TokenPos: tokenPos(p.tokens, p.pos-1),
		}
	}
	return tok, nil
}

func (p *Parser) Peek() (stringParsing.ParsedNode, error) {
	p.SkipIgnored()
	if p.pos >= len(p.tokens) {
		return stringParsing.ParsedNode{}, &ParseError{
			Code:     "Peek",
			TokenIdx: p.pos,
			TokenPos: tokenPos(p.tokens, p.pos),
			Msg:      "EOF",
		}
	}
	return p.tokens[p.pos], nil
}

func (p *Parser) String() string {
	return "lc/parsing/stringParsing/parser3/Parser"
}
