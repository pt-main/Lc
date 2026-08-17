package stringParsing

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
)

// LexerRule defines a single token type and its regular expression pattern.
type LexerRule struct {
	Type    string
	Pattern *regexp2.Regexp
}

// LexerConfig holds configuration options for the lexer.
type LexerConfig struct {
	UseBracketBalance bool
	Brackets          [][2]string
}

// Lexer converts a source string into a sequence of ParsedNode objects.
type Lexer struct {
	rules       []LexerRule
	config      LexerConfig
	openToClose map[string]string
	closeToOpen map[string]string
	ruleGroups  [][]string
	openByByte  map[byte][]string
	closeByByte map[byte][]string
}

// NewLexer creates a lexer with the given rule set and optional configuration.
func NewLexer(rules []LexerRule, config *LexerConfig) *Lexer {
	cfg := LexerConfig{}
	if config != nil {
		cfg = *config
	}

	openToClose := make(map[string]string)
	closeToOpen := make(map[string]string)
	openByByte := make(map[byte][]string)
	closeByByte := make(map[byte][]string)

	for _, pair := range cfg.Brackets {
		if len(pair) != 2 {
			continue
		}
		open, close := pair[0], pair[1]
		openToClose[open] = close
		closeToOpen[close] = open
		if len(open) > 0 {
			openByByte[open[0]] = append(openByByte[open[0]], open)
		}
		if len(close) > 0 {
			closeByByte[close[0]] = append(closeByByte[close[0]], close)
		}
	}

	for b := range openByByte {
		sort.Slice(openByByte[b], func(i, j int) bool {
			return len(openByByte[b][i]) > len(openByByte[b][j])
		})
	}
	for b := range closeByByte {
		sort.Slice(closeByByte[b], func(i, j int) bool {
			return len(closeByByte[b][i]) > len(closeByByte[b][j])
		})
	}

	ruleGroups := make([][]string, len(rules))
	for i, rule := range rules {
		ruleGroups[i] = rule.Pattern.GetGroupNames()
	}

	return &Lexer{
		rules:       rules,
		config:      cfg,
		openToClose: openToClose,
		closeToOpen: closeToOpen,
		ruleGroups:  ruleGroups,
		openByByte:  openByByte,
		closeByByte: closeByByte,
	}
}

func snippet(s string, maxRunes int) string {
	runes := []rune(s)
	if len(runes) <= maxRunes {
		return s
	}
	return string(runes[:maxRunes]) + "..."
}

func snippetFromRunes(runes []rune, pos, maxRunes int) string {
	end := pos + maxRunes
	if end >= len(runes) {
		return string(runes[pos:])
	}
	return string(runes[pos:end]) + "..."
}

func posToLineCol(code string, pos int) (line, col int) {
	line, col = 1, 1
	runeIdx := 0
	for _, r := range code {
		if runeIdx == pos {
			return
		}
		if r == '\n' {
			line++
			col = 1
		} else {
			col++
		}
		runeIdx++
	}
	return
}

// isBracketBalanced checks if the accumulated text has balanced brackets.
func (l *Lexer) isBracketBalanced(text string) bool {
	if !l.config.UseBracketBalance || len(l.openToClose) == 0 {
		return true
	}
	stack := make([]string, 0, 8)
	i := 0
	n := len(text)

	for i < n {
		matched := false
		if candidates, ok := l.openByByte[text[i]]; ok {
			for _, open := range candidates {
				if strings.HasPrefix(text[i:], open) {
					stack = append(stack, open)
					i += len(open)
					matched = true
					break
				}
			}
		}
		if matched {
			continue
		}
		if candidates, ok := l.closeByByte[text[i]]; ok {
			for _, close := range candidates {
				if strings.HasPrefix(text[i:], close) {
					if len(stack) == 0 {
						return false
					}
					last := stack[len(stack)-1]
					if open, ok := l.closeToOpen[close]; !ok || last != open {
						return false
					}
					stack = stack[:len(stack)-1]
					i += len(close)
					matched = true
					break
				}
			}
		}
		if matched {
			continue
		}
		_, size := utf8.DecodeRuneInString(text[i:])
		i += size
	}
	return len(stack) == 0
}

// bracketBalanceError returns a detailed error if brackets are unbalanced.
func (l *Lexer) bracketBalanceError(text string) *core.Error {
	if !l.config.UseBracketBalance || len(l.openToClose) == 0 {
		return nil
	}
	stack := make([]string, 0, 8)
	i := 0
	n := len(text)

	for i < n {
		matched := false
		if candidates, ok := l.openByByte[text[i]]; ok {
			for _, open := range candidates {
				if strings.HasPrefix(text[i:], open) {
					stack = append(stack, open)
					i += len(open)
					matched = true
					break
				}
			}
		}
		if matched {
			continue
		}
		if candidates, ok := l.closeByByte[text[i]]; ok {
			for _, close := range candidates {
				if strings.HasPrefix(text[i:], close) {
					if len(stack) == 0 {
						return core.Err(errors.ParsingError, "Unexpected closing bracket %q at byte %d", close, i).
							WithMeta(core.EMK(0, "string"), close).
							WithMeta(core.EMK(1, "int"), i)
					}
					last := stack[len(stack)-1]
					if open, ok := l.closeToOpen[close]; !ok || last != open {
						return core.Err(errors.ParsingError, "Mismatched bracket: expected closing %q for %q, got %q at byte %d",
							l.openToClose[last], last, close, i).
							WithMeta(core.EMK(0, "string"), close).
							WithMeta(core.EMK(1, "string"), last).
							WithMeta(core.EMK(2, "int"), i)
					}
					stack = stack[:len(stack)-1]
					i += len(close)
					matched = true
					break
				}
			}
		}
		if matched {
			continue
		}
		_, size := utf8.DecodeRuneInString(text[i:])
		i += size
	}
	if len(stack) > 0 {
		unclosed := stack[len(stack)-1]
		return core.Err(errors.ParsingError, "Unclosed bracket %q", unclosed).
			WithMeta(core.EMK(0, "string"), unclosed)
	}
	return nil
}

// Parse scans the entire input string and returns a slice of ParsedNode.
//
// Err errors.ParsingError:
//   - If bracket balancing is enabled and the input has unbalanced brackets.
//     Meta: EMK(0, "string") – the whole input code.
//   - If a regexp rule fails to match.
//     Meta: EMK(0, "string") – rule type, EMK(1, "string") – substring being matched.
//   - If no rule matches at the current position.
//     Meta: EMK(0, "int") – line number, EMK(1, "int") – column number,
//     EMK(2, "string") – context snippet.
func (lp *Lexer) Parse(code string, opts ...*parsing.ParseOption) ([]ParsedNode, core.ErrorInterface) {
	var log func(string)
	if len(opts) > 0 && opts[0].UEP != nil && opts[0].UEP.Logger != nil {
		logger := opts[0].UEP.Logger
		log = func(text string) {
			logger.PrintLog(public.LogParsing, "\\n"+text)
		}
	} else {
		log = func(string) {}
	}
	log("start parsing code [" + code + "]")

	if lp.config.UseBracketBalance {
		if err := lp.bracketBalanceError(code); err != nil {
			return nil, core.Wrap(errors.ParsingError, err, "Bracket balance error").
				WithMeta(core.EMK(0, "string"), code)
		}
	}

	var nodes []ParsedNode
	runes := []rune(code)
	pos := 0
	length := len(runes)

	for pos < length {
		log(fmt.Sprintf("pos %v, length %v", pos, length))
		matched := false
		subStr := string(runes[pos:])

		for ruleIdx, rule := range lp.rules {
			m, err := rule.Pattern.FindStringMatch(subStr)
			if err != nil {
				return nil, core.Wrap(errors.ParsingError, err, "Regexp error for rule %q", rule.Type).
					WithMeta(core.EMK(0, "string"), rule.Type).
					WithMeta(core.EMK(1, "string"), subStr)
			}
			log(fmt.Sprintf("rule %v, pos %v, substrLen %v", rule, pos, len(subStr)))
			if m != nil && m.Index == 0 {
				tokenRunes := runes[pos : pos+m.Length]
				tokenValue := string(tokenRunes)
				startPos := pos
				endPos := pos + m.Length

				meta := map[string]interface{}{
					"__raw":   tokenValue,
					"__value": tokenValue,
					"__pos":   startPos,
					"__start": startPos,
					"__end":   endPos,
				}

				groupNames := lp.ruleGroups[ruleIdx]
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
			line, col := posToLineCol(code, pos)
			context := snippetFromRunes(runes, pos, 20)
			return nil, core.Err(errors.ParsingError, "Unexpected sequence near %q at line %d, col %d", context, line, col).
				WithMeta(core.EMK(0, "int"), line).
				WithMeta(core.EMK(1, "int"), col).
				WithMeta(core.EMK(2, "string"), context)
		}
	}
	return addPrevNextNodes(nodes), nil
}

func (l *Lexer) String() string {
	return "lc/parsing/stringParsing/Lexer"
}
