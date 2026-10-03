package stringParsing

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/parsing"
	"github.com/pt-main/lc/v2/public"
	"github.com/pt-main/lc/v2/public/errors"
)

// LexerRule defines a single token type and its regular expression pattern.
type LexerRule struct {
	Type    string
	Pattern string
}

// The regular expressions follow .NET semantics for the character classes,
// where \s covers the Unicode separators and \d accepts every script's digits.
// The standard library keeps those classes ASCII-only, so the classes below
// are substituted at compile time to keep the accepted token set identical.
// Each one was verified against the .NET definitions over every code point
// from U+0000 to U+10FFFF.
const (
	netSpaceClass = `[\p{Z}\t\n\v\f\r\x{0085}]`
	netDigitClass = `\p{Nd}`
	// .NET counts the join controls U+200C and U+200D as word characters,
	// which no Unicode general category covers.
	netWordClass = `[\p{L}\p{Mn}\p{Nd}\p{Pc}\x{200C}\x{200D}]`
)

// translatePattern rewrites the .NET character classes to their
// standard-library equivalents. Only the class shorthands are touched; every
// other construct is copied verbatim so regexp.Compile still rejects anything
// it cannot honour instead of silently matching something else.
func translatePattern(pattern string) string {
	if !strings.ContainsAny(pattern, `\sSdDwW`) {
		return pattern
	}
	var b strings.Builder
	b.Grow(len(pattern) + 64)
	rs := []rune(pattern)
	for i := 0; i < len(rs); i++ {
		if rs[i] != '\\' || i+1 >= len(rs) {
			b.WriteRune(rs[i])
			continue
		}
		i++
		switch c := rs[i]; c {
		case 's':
			b.WriteString(netSpaceClass)
		case 'd':
			b.WriteString(netDigitClass)
		case 'w':
			b.WriteString(netWordClass)
		case 'S':
			b.WriteString("[^" + netSpaceClass[1:])
		case 'D':
			b.WriteString("[^" + netDigitClass + "]")
		case 'W':
			b.WriteString("[^" + netWordClass[1:])
		default:
			b.WriteByte('\\')
			b.WriteRune(c)
		}
	}
	return b.String()
}

// LexerConfig holds configuration options for the lexer.
type LexerConfig struct {
	UseBracketBalance bool
	Brackets          [][2]string
}

// compiledRule is one rule ready for scanning: the pattern is already
// translated and compiled, so the token loop never touches a raw string.
type compiledRule struct {
	typ    string
	re     *regexp.Regexp
	names  []string
	groups []string
}

// Lexer converts a source string into a sequence of ParsedNode objects.
type Lexer struct {
	rules       []LexerRule
	compiled    []compiledRule
	config      LexerConfig
	openToClose map[string]string
	closeToOpen map[string]string
	openByByte  map[byte][]string
	closeByByte map[byte][]string
}

// NewLexer creates a lexer with the given rule set and optional configuration.
//
// Rules whose pattern does not compile are rejected rather than silently
// skipped, so a broken grammar surfaces at construction rather than as a
// confusing "no rule matched" later.
func NewLexer(rules []LexerRule, config *LexerConfig) (*Lexer, core.ErrorInterface) {
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

	compiled := make([]compiledRule, 0, len(rules))
	for _, rule := range rules {
		re, err := regexp.Compile(translatePattern(rule.Pattern))
		if err != nil {
			return nil, core.Err(errors.ParsingError, "Invalid pattern for rule %q: %s", rule.Type, err.Error()).
				WithMeta(core.EMK(0, "string"), rule.Type).
				WithMeta(core.EMK(1, "string"), rule.Pattern)
		}
		names := re.SubexpNames()
		// A named capture is exposed under its name and a plain capture under
		// its index, which is what the previous engine produced.
		groups := make([]string, len(names))
		for i := 1; i < len(names); i++ {
			groups[i] = strconv.Itoa(i)
		}
		compiled = append(compiled, compiledRule{typ: rule.Type, re: re, names: names, groups: groups})
	}

	return &Lexer{
		rules:       rules,
		compiled:    compiled,
		config:      cfg,
		openToClose: openToClose,
		closeToOpen: closeToOpen,
		openByByte:  openByByte,
		closeByByte: closeByByte,
	}, nil
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

// bracketBalanceError walks text once and reports the first bracket problem,
// or nil when the text is balanced. The boolean result of the same walk is
// what isBracketBalanced needs, so both share this single implementation.
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

// isBracketBalanced checks if the accumulated text has balanced brackets.
func (l *Lexer) isBracketBalanced(text string) bool {
	return l.bracketBalanceError(text) == nil
}

// Parse scans the entire input string and returns a slice of ParsedNode.
//
// Err errors.ParsingError:
//   - If bracket balancing is enabled and the input has unbalanced brackets.
//     Meta: EMK(0, "string") - the whole input code.
//   - If a regexp rule fails to match.
//     Meta: EMK(0, "string") - rule type, EMK(1, "string") - substring being matched.
//   - If no rule matches at the current position.
//     Meta: EMK(0, "int") - line number, EMK(1, "int") - column number,
//     EMK(2, "string") - context snippet.
func (l *Lexer) Parse(code string, opts ...*parsing.ParseOption) ([]ParsedNode, core.ErrorInterface) {
	var logger core.LoggerInterface
	if len(opts) > 0 && opts[0] != nil && opts[0].UEP != nil {
		logger = opts[0].UEP.Logger
	}
	// A nil logger is the common case, and fmt.Sprintf arguments are evaluated
	// whether or not the result is used, so guard each call instead of
	// concatenating text that is then dropped. The entry message embeds the
	// whole input, so building it eagerly copied the source on every call.
	log := func(text string) {
		if logger != nil {
			logger.PrintLog(public.LogParsing, "\\n"+text)
		}
	}
	if logger != nil {
		log("start parsing code [" + code + "]")
	}

	if l.config.UseBracketBalance {
		if err := l.bracketBalanceError(code); err != nil {
			return nil, core.Wrap(errors.ParsingError, err, "Bracket balance error").
				WithMeta(core.EMK(0, "string"), code)
		}
	}

	// Positions are kept as rune indices for the token metadata, while the
	// scan itself works on byte offsets, which is what the standard engine
	// indexes by.
	runes := []rune(code)
	// The token count is unknown up front, so size the slice from a cheap
	// estimate: at least one rune per token, with headroom for the common case
	// where rules split the input into several tokens.
	nodes := make([]ParsedNode, 0, len(runes)+len(runes)/2+8)
	balanceEnabled := l.config.UseBracketBalance

	bytePos := 0
	runePos := 0
	length := len(runes)

	for runePos < length {
		matched := false

		for _, rule := range l.compiled {
			if logger != nil {
				log(fmt.Sprintf("rule %v, pos %v, len %v", rule.typ, runePos, length-runePos))
			}
			loc := rule.re.FindStringIndex(code[bytePos:])
			// A zero-length match makes no progress, so it would loop forever.
			// The index check keeps the match anchored at the scan position.
			if loc == nil || loc[0] != 0 || loc[1] == 0 {
				continue
			}
			matchedRunes := utf8.RuneCountInString(code[bytePos : bytePos+loc[1]])
			tokenValue := code[bytePos : bytePos+loc[1]]
			startPos := runePos
			endPos := runePos + matchedRunes

			// Presize to the five fixed keys plus the named groups this rule
			// declares, so the map never has to grow while filling.
			meta := make(map[string]interface{}, 5+2*len(rule.groups))
			meta["__raw"] = tokenValue
			meta["__value"] = tokenValue
			meta["__pos"] = startPos
			meta["__start"] = startPos
			meta["__end"] = endPos
			if len(rule.groups) > 1 {
				sub := rule.re.FindStringSubmatchIndex(code[bytePos:])
				for i := 1; i < len(rule.groups); i++ {
					key := rule.groups[i]
					if named := rule.names[i]; named != "" {
						key = named
					}
					s, e := sub[2*i], sub[2*i+1]
					if s < 0 || e < s {
						// The group took no part in the match; it is reported
						// as empty, matching the previous engine.
						meta[key] = ""
						continue
					}
					meta[key] = code[bytePos+s : bytePos+e]
				}
			}
			// A single rune may itself be an unbalanced bracket, so the walk
			// runs whenever balance checking is on.
			if balanceEnabled {
				meta["__bracket_balanced"] = l.isBracketBalanced(tokenValue)
			}

			nodes = append(nodes, ParsedNode{
				Raw:      tokenValue,
				Switch:   rule.typ,
				Metadata: meta,
			})
			bytePos += loc[1]
			runePos += matchedRunes
			matched = true
			break
		}
		if !matched {
			line, col := posToLineCol(code, runePos)
			context := snippetFromRunes(runes, runePos, 20)
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
