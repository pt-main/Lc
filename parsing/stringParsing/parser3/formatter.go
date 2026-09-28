package parser3

import (
	goerr "errors"
	"strconv"
	"strings"
)

// Minimal subset of ANSI SGR codes used by the formatter.
const (
	ansiReset  = "\x1b[0m"
	ansiBold   = "\x1b[1m"
	ansiDim    = "\x1b[2m"
	ansiRed    = "\x1b[31m"
	ansiGreen  = "\x1b[32m"
	ansiYellow = "\x1b[33m"
	ansiCyan   = "\x1b[36m"
)

type palette struct {
	enabled bool
	kind    string
	phase   string
	expect  string
	got     string
	pos     string
}

func newPalette(useColors bool) palette {
	return palette{
		enabled: useColors,
		kind:    ansiBold + ansiRed,
		phase:   ansiCyan,
		expect:  ansiGreen,
		got:     ansiYellow,
		pos:     ansiDim,
	}
}

func (p palette) paint(s, code string) string {
	if !p.enabled || code == "" || s == "" {
		return s
	}
	return code + s + ansiReset
}

// FormatError renders a parser3 error as a multi-line report: every error of
// the chain gets its own line, the inner ones indented under "caused by".
//
// Pass useColors=false for plain-text logs, JSON APIs, or file output.
func FormatError(err error, useColors bool) string {
	if err == nil {
		return ""
	}
	pal := newPalette(useColors)

	var b strings.Builder
	for i, e := range collectChain(err, pal) {
		if i > 0 {
			b.WriteString("\n  caused by: ")
		}
		b.WriteString(e.text)
	}
	return b.String()
}

// FormatErrorPretty is a convenience wrapper for CLI output (colors enabled).
func FormatErrorPretty(err error) string {
	return FormatError(err, true)
}

type errLink struct {
	text string
}

// collectChain unwraps the error chain into rendered lines, outermost first.
// The depth is bounded so a cyclic Cause cannot hang the formatter.
func collectChain(err error, pal palette) []errLink {
	out := make([]errLink, 0, 4)
	for i := 0; err != nil && i < 32; i++ {
		switch e := err.(type) {
		case *ParseError:
			out = append(out, errLink{formatParseError(e, pal)})
			err = e.Cause
		case *GrammarError:
			out = append(out, errLink{formatGrammarError(e, pal)})
			err = e.Cause
		case *AdapterError:
			out = append(out, errLink{formatAdapterError(e, pal)})
			err = e.Cause
		default:
			var pe *ParseError
			var ge *GrammarError
			var ae *AdapterError
			switch {
			case goerr.As(err, &pe):
				out = append(out, errLink{formatParseError(pe, pal)})
				err = pe.Cause
			case goerr.As(err, &ge):
				out = append(out, errLink{formatGrammarError(ge, pal)})
				err = ge.Cause
			case goerr.As(err, &ae):
				out = append(out, errLink{formatAdapterError(ae, pal)})
				err = ae.Cause
			default:
				out = append(out, errLink{pal.paint(err.Error(), ansiDim)})
				return out
			}
		}
	}
	return out
}

func formatParseError(e *ParseError, pal palette) string {
	var b strings.Builder
	b.WriteString(pal.paint("parser3", pal.kind))
	if e.Phase != "" {
		b.WriteString(pal.paint("/"+e.Phase, pal.phase))
	}
	b.WriteString(": ")

	switch {
	case e.Expected != "":
		b.WriteString("expected " + pal.paint(strconv.Quote(e.Expected), pal.expect))
		if e.Got != "" {
			b.WriteString(", got " + pal.paint(strconv.Quote(e.Got), pal.got))
		}
	case e.Got != "":
		b.WriteString("got " + pal.paint(strconv.Quote(e.Got), pal.got))
	}

	if e.Raw != "" && e.Raw != e.Got {
		b.WriteString(" (raw: " + strconv.Quote(e.Raw) + ")")
	}
	if len(e.Found) > 0 {
		b.WriteString(" (found: " + strings.Join(e.Found, ", ") + ")")
	}
	if e.Msg != "" {
		if e.Expected != "" || e.Got != "" || e.Raw != "" {
			b.WriteString(" - ")
		}
		b.WriteString(e.Msg)
	}
	if e.TokenPos != "" {
		b.WriteString(" at " + pal.paint(e.TokenPos, pal.pos))
	}
	return b.String()
}

func formatGrammarError(e *GrammarError, pal palette) string {
	var b strings.Builder
	b.WriteString(pal.paint("grammar", pal.kind))
	if e.Phase != "" {
		b.WriteString(pal.paint("/"+e.Phase, pal.phase))
	}
	b.WriteString(": ")
	b.WriteString(e.Msg)
	return b.String()
}

func formatAdapterError(e *AdapterError, pal palette) string {
	return pal.paint("adapter", pal.kind) + ": " + e.Msg
}
