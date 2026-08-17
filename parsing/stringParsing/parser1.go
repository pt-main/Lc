package stringParsing

import (
	"regexp"
	"strings"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
)

type GrammarRule struct {
	Type    string
	Pattern *regexp.Regexp
}

type Parser1Config struct {
	UseLineContinuation bool
	UseBracketBalance   bool
	SkipEmptyLines      bool
	Brackets            []string
	TrimBlocksSpace     bool
}

// Parser1 implements a regex‑based grammar parser with line continuation
// and bracket balancing support.
type Parser1 struct {
	grammar     []GrammarRule
	config      Parser1Config
	openToClose map[rune]rune
	closeToOpen map[rune]rune
}

// NewParser1 constructs a Parser1 with given grammar rules and configuration.
func NewParser1(rules []GrammarRule, config Parser1Config) *Parser1 {
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
	return &Parser1{
		grammar:     rules,
		config:      config,
		openToClose: openToClose,
		closeToOpen: closeToOpen,
	}
}

// Parse splits the input into logical blocks and applies grammar rules.
//
// Err errors.ParsingError:
//   - If no rule matches a block.
//     Meta: EMK(0, "string") – the block that failed to match.
func (p *Parser1) Parse(code string, opts ...*parsing.ParseOption) ([]ParsedNode, core.ErrorInterface) {
	log := func(text string) {
		text = "\n" + text
		if len(opts) > 0 {
			logger := opts[0].UEP.Logger
			if logger != nil {
				logger.PrintLog(public.LogParsing, text)
			}
		}
	}
	log("start parsing code " + code)
	lines := strings.Split(code, "\n")
	var result []ParsedNode

	var blockLines []string
	bracketStack := []rune{}

	flush := func(block string) core.ErrorInterface {
		node, err := p.matchGrammar(block)
		if err != nil {
			return core.Wrap(errors.ParsingError, err, "Failed to match grammar for block").
				WithMeta(core.EMK(0, "string"), block)
		}
		if node.Raw != "" {
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

	return addPrevNextNodes(result), nil
}

func (p *Parser1) matchGrammar(block string) (ParsedNode, core.ErrorInterface) {
	absolutely_raw := block
	if p.config.TrimBlocksSpace {
		block = strings.TrimSpace(block)
	}
	if block == "" && p.config.SkipEmptyLines {
		return ParsedNode{Raw: ""}, nil
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
			meta["__raw"] = absolutely_raw
			return ParsedNode{
				Raw:      block,
				Switch:   rule.Type,
				Metadata: meta,
			}, nil
		}
	}
	return ParsedNode{}, core.Err(errors.ParsingError, "Syntax error: no rule matches block: %q", block).
		WithMeta(core.EMK(0, "string"), block)
}

func (p *Parser1) String() string {
	return "lc/parsing/stringParsing/Parser1"
}
