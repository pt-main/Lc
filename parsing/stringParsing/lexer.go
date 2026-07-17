package stringParsing

import (
	"fmt"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/public"
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

// LexerConfig holds configuration options for the lexer.
type LexerConfig struct {
	UseBracketBalance bool     // Enable bracket balancing
	Brackets          []string // List of bracket pairs, e.g., []string{"()","[]","{}"}
}

// Lexer converts a source string into a sequence of ParsedNode objects.
// It applies a set of LexerRule in order, picking the first matching rule
// at each position. It supports capturing named groups from the regexp,
// which become part of Metadata.
type Lexer struct {
	rules       []LexerRule
	config      LexerConfig
	openToClose map[rune]rune
	closeToOpen map[rune]rune
}

// NewLexer creates a lexer with the given rule set and optional configuration.
// Rules are tried in the order they are provided.
// Typically used as the first stage in a string parser.
// Example:
//
//	lex := NewLexer([]LexerRule{...}, LexerConfig{})
func NewLexer(rules []LexerRule, config *LexerConfig) *Lexer {
	cfg := LexerConfig{}
	if config != nil {
		cfg = *config
	}

	openToClose := make(map[rune]rune)
	closeToOpen := make(map[rune]rune)
	for _, pair := range cfg.Brackets {
		runes := []rune(pair)
		if len(runes) == 2 {
			open, close := runes[0], runes[1]
			openToClose[open] = close
			closeToOpen[close] = open
		}
	}

	return &Lexer{
		rules:       rules,
		config:      cfg,
		openToClose: openToClose,
		closeToOpen: closeToOpen,
	}
}

func snippet(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

// isBracketBalanced checks if the accumulated text has balanced brackets.
func (l *Lexer) isBracketBalanced(text string) bool {
	if !l.config.UseBracketBalance {
		return true
	}

	stack := []rune{}
	for _, ch := range text {
		if _, ok := l.openToClose[ch]; ok {
			stack = append(stack, ch)
		} else if closeOpen, ok := l.closeToOpen[ch]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == closeOpen {
				stack = stack[:len(stack)-1]
			} else {
				return false // unmatched closing bracket
			}
		}
	}
	return len(stack) == 0
}

// Parse scans the entire input string and returns a slice of ParsedNode.
// Each node's Switch field contains the token type, Raw contains the exact
// matched substring, and Metadata includes all named groups from the regexp
// plus the token value ("__value").
// If no rule matches at some position, an error is returned with a snippet.
// The result is automatically enriched with __prev/__next links via
// addPrevNextNodes.
func (lp *Lexer) Parse(code string, opts ...*parsing.ParseOption) ([]ParsedNode, error) {
	log := func(text string) {
		text = "\n" + text
		if len(opts) > 0 {
			logger := opts[0].UEP.Logger
			if logger != nil {
				logger.PrintLog(public.LogParsing, text)
			}
		}
	}
	log("start parsing code [" + code + "]")

	if lp.config.UseBracketBalance && !lp.isBracketBalanced(code) {
		return nil, fmt.Errorf("lexer error: unbalanced brackets in input")
	}

	var nodes []ParsedNode
	runes := []rune(code)
	pos := 0
	length := len(runes)

	for pos < length {
		log(fmt.Sprintf("pos %v, length %v", pos, length))
		matched := false
		for _, rule := range lp.rules {
			subStr := string(runes[pos:])
			m, err := rule.Pattern.FindStringMatch(subStr)
			if err != nil {
				return nil, err
			}
			log(fmt.Sprintf("rule %v, pos %v, substrLen %v", rule, pos, len(subStr)))
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

				// Add bracket balance info to metadata if enabled
				if lp.config.UseBracketBalance {
					meta["__bracket_balanced"] = lp.isBracketBalanced(tokenValue)
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

func (l *Lexer) String() string {
	return "lc/parsing/stringParsing/Lexer"
}
