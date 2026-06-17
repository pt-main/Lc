package stringParsing

import (
	"fmt"

	"github.com/dlclark/regexp2"
)

// LexerRule defines a single token type and its regular expression pattern.
// Type – symbolic name of the token (e.g., "NUMBER", "IDENT", "STRING").
// Pattern – compiled *regexp2.Regexp (supports groups and repetition).
// Example:
//
//	rule := LexerRule{
//	    Type:    "IDENT",
//	    Pattern: regexp2.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`, 0),
//	}
type LexerRule struct {
	Type    string
	Pattern *regexp2.Regexp
}

// Lexer converts a source string into a sequence of ParsedNode objects.
// It applies a set of LexerRule in order, picking the first matching rule
// at each position. It supports capturing named groups from the regexp,
// which become part of Metadata.
type Lexer struct {
	rules []LexerRule
}

// NewLexer creates a lexer with the given rule set.
// Rules are tried in the order they are provided.
// Typically used as the first stage in a string parser.
// Example:
//
//	lex := NewLexer([]LexerRule{...})
func NewLexer(rules []LexerRule) *Lexer {
	return &Lexer{rules: rules}
}

func snippet(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

// Parse scans the entire input string and returns a slice of ParsedNode.
// Each node's Switch field contains the token type, Raw contains the exact
// matched substring, and Metadata includes all named groups from the regexp
// plus the token value ("__value").
// If no rule matches at some position, an error is returned with a snippet.
// The result is automatically enriched with __prev/__next links via
// addPrevNextNodes.
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
					"__raw":   tokenValue,
					"__value": tokenValue,
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
					Raw:      tokenValue,
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
