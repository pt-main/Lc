package stringParsing

import (
	"fmt"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/pt-main/lc/parsing"
)

// LexerRule defines a single token type and its regular expression pattern.
type LexerRule struct {
	Type    string
	Pattern *regexp2.Regexp
}

// LexerConfig holds configuration options for the lexer.
type LexerConfig struct {
	UseBracketBalance   bool     // Enable bracket balancing (blocks wait for balanced brackets)
	Brackets            []string // List of bracket pairs, e.g., []string{"()","[]","{}"}
	UseLineContinuation bool     // Join lines ending with '\'
	SkipEmptyLines      bool     // Ignore empty blocks (after trimming, if enabled)
	TrimBlocksSpace     bool     // Trim whitespace from block before tokenizing
}

// Lexer converts a source string into a sequence of ParsedNode objects.
// It first splits the input into logical blocks (handling line continuation
// and bracket balancing), then applies lexer rules to each block.
type Lexer struct {
	rules               []LexerRule
	config              LexerConfig
	openToClose         map[rune]rune
	closeToOpen         map[rune]rune
	useLineContinuation bool
	skipEmptyLines      bool
	trimBlocksSpace     bool
}

// NewLexer creates a lexer with the given rule set and optional configuration.
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
		rules:               rules,
		config:              cfg,
		openToClose:         openToClose,
		closeToOpen:         closeToOpen,
		useLineContinuation: cfg.UseLineContinuation,
		skipEmptyLines:      cfg.SkipEmptyLines,
		trimBlocksSpace:     cfg.TrimBlocksSpace,
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
				return false
			}
		}
	}
	return len(stack) == 0
}

// tokenizeBlock applies lexer rules to a single block and returns a list of tokens.
func (l *Lexer) tokenizeBlock(block string, blockIndex int) ([]ParsedNode, error) {
	if l.trimBlocksSpace {
		block = strings.TrimSpace(block)
	}
	if block == "" && l.skipEmptyLines {
		return nil, nil
	}

	var nodes []ParsedNode
	runes := []rune(block)
	pos := 0
	length := len(runes)

	for pos < length {
		matched := false
		for _, rule := range l.rules {
			subStr := string(runes[pos:])
			m, err := rule.Pattern.FindStringMatch(subStr)
			if err != nil {
				return nil, err
			}
			if m != nil && m.Index == 0 {
				tokenRunes := runes[pos : pos+m.Length]
				tokenValue := string(tokenRunes)

				meta := map[string]interface{}{
					"__raw":         tokenValue,
					"__value":       tokenValue,
					"__block_index": blockIndex,
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

				if l.config.UseBracketBalance {
					meta["__bracket_balanced"] = l.isBracketBalanced(tokenValue)
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
			return nil, fmt.Errorf("lexer error: unexpected sequence near '%s' position %v",
				snippet(string(runes[pos:]), 20), pos)
		}
	}
	return nodes, nil
}

// Parse scans the entire input string, splits into blocks, and tokenizes each block.
// Returns a flat list of all tokens, enriched with __prev/__next links.
func (l *Lexer) Parse(code string, opts ...*parsing.ParseOption) ([]ParsedNode, error) {
	log := func(text string) {
		text = "\n" + text
		if len(opts) > 0 {
			uep := opts[0].UEP
			if uep != nil && uep.Logger != nil {
				uep.Logger.PrintLog("parsing", text)
			}
		}
	}
	log("start parsing code [" + code + "]")

	lines := strings.Split(code, "\n")
	var allTokens []ParsedNode
	var blockLines []string
	bracketStack := []rune{}
	blockIndex := 0

	flush := func(block string) error {
		tokens, err := l.tokenizeBlock(block, blockIndex)
		if err != nil {
			return err
		}
		if tokens != nil {
			allTokens = append(allTokens, tokens...)
			blockIndex++
		}
		return nil
	}

	for _, rawLine := range lines {
		line := strings.TrimRight(rawLine, " \t")

		if line == "" && len(bracketStack) == 0 && !l.config.UseBracketBalance {
			// Empty line outside a bracket context – treat as block separator only if not balancing
			// In Parser1, they skip empty lines only if SkipEmptyLines is true and no brackets.
			// We'll follow similar logic: if we are not balancing, an empty line can be a block end.
			// But we need to flush the current block if any.
			if len(blockLines) > 0 {
				block := strings.Join(blockLines, "\n")
				if err := flush(block); err != nil {
					return nil, err
				}
				blockLines = nil
				bracketStack = []rune{}
			}
			continue
		}

		continues := l.useLineContinuation && strings.HasSuffix(line, "\\")
		if continues {
			line = strings.TrimSuffix(line, "\\")
		}

		blockLines = append(blockLines, line)

		if l.config.UseBracketBalance {
			for _, ch := range line {
				if _, ok := l.openToClose[ch]; ok {
					bracketStack = append(bracketStack, ch)
				} else if closeOpen, ok := l.closeToOpen[ch]; ok {
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
		if l.config.UseBracketBalance && len(bracketStack) > 0 {
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

	return addPrevNextNodes(allTokens), nil
}

func (l *Lexer) String() string {
	return "lc/parsing/stringParsing/Lexer"
}
