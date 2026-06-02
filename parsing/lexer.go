package parsing

import (
	"errors"
	"regexp"
)

type LexerRule struct {
	Type    string
	Pattern *regexp.Regexp
}

type LexerParser struct {
	rules []LexerRule
}

func NewLexerParser(rules []LexerRule) *LexerParser {
	return &LexerParser{rules: rules}
}

func (lp *LexerParser) Parse(code string) ([]ParsedNode, error) {
	var nodes []ParsedNode
	pos := 0
	length := len(code)

	for pos < length {
		matched := false
		for _, rule := range lp.rules {
			loc := rule.Pattern.FindStringIndex(code[pos:])
			if loc != nil && loc[0] == 0 {
				tokenValue := code[pos : pos+loc[1]]
				meta := map[string]interface{}{
					"__raw": tokenValue,
					"value": tokenValue,
				}

				matches := rule.Pattern.FindStringSubmatch(code[pos:])
				if matches != nil {
					names := rule.Pattern.SubexpNames()
					for i, name := range names {
						if name != "" && i < len(matches) {
							meta[name] = matches[i]
						}
					}
				}
				nodes = append(nodes, ParsedNode{
					Parsed:   []string{tokenValue},
					Switch:   rule.Type,
					Metadata: meta,
				})
				pos += loc[1]
				matched = true
				break
			}
		}
		if !matched {
			return nil, errors.New("lexer error: unexpected sequence at position " + string(rune(pos)))
		}
	}
	return nodes, nil
}
