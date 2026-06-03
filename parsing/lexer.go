package parsing

import (
	"errors"

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

func (lp *Lexer) Parse(code string) ([]ParsedNode, error) {
	var nodes []ParsedNode
	pos := 0
	length := len(code)

	for pos < length {
		matched := false
		for _, rule := range lp.rules {
			m, err := rule.Pattern.FindStringMatch(code[pos:])
			if err != nil {
				return nil, err
			}
			if m != nil && m.Index == 0 {
				tokenValue := code[pos : pos+m.Length]
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
			return nil, errors.New("lexer error: unexpected sequence at position " + string(rune(pos)))
		}
	}
	return nodes, nil
}
