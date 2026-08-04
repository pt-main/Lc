package stringParsing

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/public"
)

// LexerRule defines a single token type and its regular expression pattern.
type LexerRule struct {
	Type    string
	Pattern *regexp2.Regexp
}

// LexerConfig holds configuration options for the lexer.
type LexerConfig struct {
	UseBracketBalance bool        // Enable bracket balancing
	Brackets          [][2]string // List of bracket pairs, e.g. [][2]string{{"(", ")"}, {"<!--", "-->"}, {"begin", "end"}}
}

// # Lexer
//
// Lexer converts a source string into a sequence of ParsedNode objects. Works like Parser1, but
// find tokens in full text, not in current line.
type Lexer struct {
	rules       []LexerRule
	config      LexerConfig
	openToClose map[string]string
	closeToOpen map[string]string
}

// NewLexer creates a lexer with the given rule set and optional configuration.
func NewLexer(rules []LexerRule, config *LexerConfig) *Lexer {
	cfg := LexerConfig{}
	if config != nil {
		cfg = *config
	}

	openToClose := make(map[string]string)
	closeToOpen := make(map[string]string)
	for _, pair := range cfg.Brackets {
		if len(pair) == 2 {
			open, close := pair[0], pair[1]
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
// Supports multi‑character opening/closing strings.
func (l *Lexer) isBracketBalanced(text string) bool {
	if !l.config.UseBracketBalance || len(l.openToClose) == 0 {
		return true
	}

	stack := []string{}
	i := 0
	n := len(text)

	for i < n {
		matched := false

		// Try to match an opening bracket
		for open := range l.openToClose {
			if strings.HasPrefix(text[i:], open) {
				stack = append(stack, open)
				i += len(open)
				matched = true
				break
			}
		}
		if matched {
			continue
		}

		// Try to match a closing bracket
		for close, open := range l.closeToOpen {
			if strings.HasPrefix(text[i:], close) {
				if len(stack) == 0 || stack[len(stack)-1] != open {
					return false
				}
				stack = stack[:len(stack)-1]
				i += len(close)
				matched = true
				break
			}
		}
		if matched {
			continue
		}

		// No bracket matched – skip one UTF‑8 character
		_, size := utf8.DecodeRuneInString(text[i:])
		i += size
	}

	return len(stack) == 0
}

// Parse scans the entire input string and returns a slice of ParsedNode.
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
