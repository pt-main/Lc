package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pt-main/lc/parsing/stringParsing"
)

type configParser struct {
	tokens []stringParsing.ParsedNode
	pos    int
	path   []string
}

func (p *configParser) Parse(code string, _ ...interface{}) ([]stringParsing.ParsedNode, error) {

	lexer := newLexer()
	tokens, err := lexer.Parse(code, nil)
	if err != nil {
		return nil, err
	}

	var filtered []stringParsing.ParsedNode
	for _, t := range tokens {
		if t.Switch != "COMMENT" && t.Switch != "WHITESPACE" {
			filtered = append(filtered, t)
		}
	}
	p.tokens = filtered
	p.pos = 0
	p.path = []string{}

	var result []stringParsing.ParsedNode

	for p.pos < len(p.tokens) {
		nodes, err := p.parseTopLevel()
		if err != nil {
			return nil, err
		}
		result = append(result, nodes...)
	}
	return result, nil
}

func (p *configParser) parseTopLevel() ([]stringParsing.ParsedNode, error) {
	tok := p.peek()
	if tok == nil {
		return nil, nil
	}
	if tok.Switch == "IDENT" {

		next := p.peekNext()
		if next != nil && next.Switch == "LBRACE" {
			return p.parseBlock()
		}
		return p.parseAssign()
	}

	return nil, fmt.Errorf("unexpected token %q at position %d", tok.Raw, p.pos)
}

func (p *configParser) parseBlock() ([]stringParsing.ParsedNode, error) {

	nameTok := p.consume()
	if nameTok == nil || nameTok.Switch != "IDENT" {
		return nil, fmt.Errorf("expected block name")
	}

	brace := p.consume()
	if brace == nil || brace.Switch != "LBRACE" {
		return nil, fmt.Errorf("expected '{'")
	}

	p.path = append(p.path, nameTok.Raw)

	var nodes []stringParsing.ParsedNode

	for p.pos < len(p.tokens) && p.peek().Switch != "RBRACE" {
		inner, err := p.parseTopLevel()
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, inner...)
	}

	closeBrace := p.consume()
	if closeBrace == nil || closeBrace.Switch != "RBRACE" {
		return nil, fmt.Errorf("expected '}'")
	}

	p.path = p.path[:len(p.path)-1]
	return nodes, nil
}

func (p *configParser) parseAssign() ([]stringParsing.ParsedNode, error) {

	keyTok := p.consume()
	if keyTok == nil || keyTok.Switch != "IDENT" {
		return nil, fmt.Errorf("expected key name")
	}

	assign := p.consume()
	if assign == nil || assign.Switch != "ASSIGN" {
		return nil, fmt.Errorf("expected '='")
	}

	valTok := p.consume()
	if valTok == nil {
		return nil, fmt.Errorf("expected value")
	}
	var value interface{}
	switch valTok.Switch {
	case "STRING":

		s := valTok.Raw
		if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
			value = s[1 : len(s)-1]
		} else {
			value = s
		}
	case "NUMBER":
		if strings.Contains(valTok.Raw, ".") {
			f, _ := strconv.ParseFloat(valTok.Raw, 64)
			value = f
		} else {
			i, _ := strconv.Atoi(valTok.Raw)
			value = i
		}
	case "BOOL":
		value = valTok.Raw == "true"
	default:
		return nil, fmt.Errorf("unexpected value type %q", valTok.Switch)
	}

	fullKey := strings.Join(append(p.path, keyTok.Raw), ".")

	node := stringParsing.ParsedNode{
		Switch: "assign",
		Raw:    keyTok.Raw,
		Metadata: map[string]interface{}{
			"key":   fullKey,
			"value": value,
		},
	}
	return []stringParsing.ParsedNode{node}, nil
}

func (p *configParser) peek() *stringParsing.ParsedNode {
	if p.pos >= len(p.tokens) {
		return nil
	}
	return &p.tokens[p.pos]
}

func (p *configParser) peekNext() *stringParsing.ParsedNode {
	if p.pos+1 >= len(p.tokens) {
		return nil
	}
	return &p.tokens[p.pos+1]
}

func (p *configParser) consume() *stringParsing.ParsedNode {
	if p.pos >= len(p.tokens) {
		return nil
	}
	tok := &p.tokens[p.pos]
	p.pos++
	return tok
}

func (p *configParser) String() string
