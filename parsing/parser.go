package parsing

import (
	"errors"
	"regexp"
	"strings"
)

type GrammarRule struct {
	Type    string
	Pattern *regexp.Regexp
}

type ParserConfig struct {
	UseLineContinuation bool
	UseBracketBalance   bool
	Brackets            []string
}

type Parser struct {
	grammar     []GrammarRule
	config      ParserConfig
	openToClose map[rune]rune
	closeToOpen map[rune]rune
}

func NewParser(rules []GrammarRule, config ParserConfig) *Parser {
	openToClose := make(map[rune]rune)
	closeToOpen := make(map[rune]rune)
	for _, pair := range config.Brackets {
		runes := []rune(pair)
		if len(runes) == 2 {
			open, close := runes[0], runes[1]
			openToClose[open] = close
			closeToOpen[close] = open
		}
	}
	return &Parser{
		grammar:     rules,
		config:      config,
		openToClose: openToClose,
		closeToOpen: closeToOpen,
	}
}

func (p *Parser) Parse(code string) ([]ParsedNode, error) {
	lines := strings.Split(code, "\n")
	var result []ParsedNode

	var blockLines []string
	bracketStack := []rune{}

	flush := func(block string) error {
		node, err := p.matchGrammar(block)
		if err != nil {
			return err
		}
		if node.Parsed != nil {
			result = append(result, node)
		}
		return nil
	}

	for _, rawLine := range lines {
		line := strings.TrimRight(rawLine, " \t")

		if line == "" && len(bracketStack) == 0 && !p.config.UseBracketBalance {
			continue
		}

		continues := p.config.UseLineContinuation && strings.HasSuffix(line, "\\")
		if continues {
			line = strings.TrimSuffix(line, "\\")
		}

		blockLines = append(blockLines, line)

		if p.config.UseBracketBalance {
			for _, ch := range line {
				if _, ok := p.openToClose[ch]; ok {
					bracketStack = append(bracketStack, ch)
				} else if closeOpen, ok := p.closeToOpen[ch]; ok {
					if len(bracketStack) > 0 && bracketStack[len(bracketStack)-1] == closeOpen {
						bracketStack = bracketStack[:len(bracketStack)-1]
					}
				}
			}
		}

		blockComplete := true
		if continues {
			blockComplete = false
		}
		if p.config.UseBracketBalance && len(bracketStack) > 0 {
			blockComplete = false
		}

		if blockComplete {
			block := strings.Join(blockLines, "\n")
			if err := flush(block); err != nil {
				return nil, err
			}
			blockLines = nil
			bracketStack = []rune{}
		}
	}

	if len(blockLines) > 0 {
		block := strings.Join(blockLines, "\n")
		if err := flush(block); err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (p *Parser) matchGrammar(block string) (ParsedNode, error) {
	block = strings.TrimSpace(block)
	if block == "" {
		return ParsedNode{Parsed: nil}, nil
	}
	for _, rule := range p.grammar {
		if rule.Pattern.MatchString(block) {
			meta := make(map[string]interface{})

			matches := rule.Pattern.FindStringSubmatch(block)
			names := rule.Pattern.SubexpNames()
			for i, name := range names {
				if i != 0 && name != "" && i < len(matches) {
					meta[name] = matches[i]
				}
			}

			meta["__raw"] = block

			return ParsedNode{
				Parsed:   []string{block},
				Switch:   rule.Type,
				Metadata: meta,
			}, nil
		}
	}
	return ParsedNode{}, errors.New("syntax error: no rule matches block: " + block)
}
