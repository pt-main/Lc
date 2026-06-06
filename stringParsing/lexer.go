package stringParsing

import (
	"fmt"

	"github.com/dlclark/regexp2"
)

type LexerRule struct {
	Type    string
	Pattern *regexp2.Regexp
}

type Lexer struct {
	rules []LexerRule
}

func NewLexer(rules []LexerRule) *Lexer {
	return &Lexer{rules: rules}
}

func snippet(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

func (lp *Lexer) Parse(code string) ([]ParsedNode, error) {
	var nodes []ParsedNode
	runes := []rune(code)
	pos := 0
	length := len(runes)

	for pos < length {
		matched := false
		for _, rule := range lp.rules {
			subStr := string(runes[pos:])
			m, err := rule.Pattern.FindStringMatch(subStr)
			if err != nil {
				return nil, err
			}
			if m != nil && m.Index == 0 {
				tokenRunes := runes[pos : pos+m.Length]
				tokenValue := string(tokenRunes)

				meta := map[string]interface{}{
					"__raw": tokenValue,
					"value": tokenValue,
				}

				groupNames := rule.Pattern.GetGroupNames()
				for _, name := range groupNames {
					if name != "0" {
						grp := m.GroupByName(name)
						if grp != nil {
							meta[name] = grp.String()
						}
					}
				}
				nodes = append(nodes, ParsedNode{
					Parsed:   []string{tokenValue},
					Switch:   rule.Type,
					Metadata: meta,
				})
				pos += m.Length
				matched = true
				break
			}
		}
		if !matched {
			return nil, fmt.Errorf("lexer error: unexpected sequence near '%s' position %v", snippet(string(runes[pos:]), 20), pos)
		}
	}
	return addPrevNextNodes(nodes), nil
}
