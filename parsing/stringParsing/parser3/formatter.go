package parser3

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pt-main/tap/color"
)

// FormatError renders a parser3 error with optional ANSI color codes.
// Pass useColors=false for plain-text logs, JSON APIs, or file output.
func FormatError(err error, useColors bool) string {
	if err == nil {
		return ""
	}

	// Build a chain of structured errors.
	var parts []string
	walk(err, &parts)

	if !useColors {
		return strings.Join(parts, " -> ")
	}

	// Colorize each part.
	colored := make([]string, len(parts))
	for i, p := range parts {
		if i == len(parts)-1 {
			// The root cause — highlight in red.
			colored[i] = color.Set(fmt.Sprintf("[?RD]%s[?RT]", p))
		} else {
			colored[i] = color.Set(fmt.Sprintf("[?YW]%s[?RT]", p))
		}
	}
	return strings.Join(colored, " -> ")
}

// FormatErrorPretty is a convenience wrapper for CLI output (colors enabled).
func FormatErrorPretty(err error) string {
	return FormatError(err, true)
}

// walk unwraps the error chain and collects human-readable messages.
func walk(err error, out *[]string) {
	if err == nil {
		return
	}

	// Try structured types first for rich output.
	var pe *ParseError
	if errors.As(err, &pe) {
		*out = append(*out, formatParseError(pe))
		walk(pe.Cause, out)
		return
	}

	var ge *GrammarError
	if errors.As(err, &ge) {
		*out = append(*out, formatGrammarError(ge))
		walk(ge.Cause, out)
		return
	}

	var ae *AdapterError
	if errors.As(err, &ae) {
		*out = append(*out, formatAdapterError(ae))
		walk(ae.Cause, out)
		return
	}

	// Fallback for plain errors.
	*out = append(*out, err.Error())
}

func formatParseError(e *ParseError) string {
	var b strings.Builder
	b.WriteString("parser3")
	if e.Code != "" {
		b.WriteString("/")
		b.WriteString(e.Code)
	}
	b.WriteString(": ")
	if e.Expected != "" {
		b.WriteString(fmt.Sprintf("expected '%s'", e.Expected))
		if e.Got != "" {
			b.WriteString(fmt.Sprintf(", got '%s'", e.Got))
		}
	} else if e.Got != "" {
		b.WriteString(fmt.Sprintf("got '%s'", e.Got))
	}
	if e.Raw != "" && e.Raw != e.Got {
		b.WriteString(fmt.Sprintf(" (raw: %q)", e.Raw))
	}
	if e.TokenPos != "" {
		b.WriteString(fmt.Sprintf(" at %s", e.TokenPos))
	}
	return b.String()
}

func formatGrammarError(e *GrammarError) string {
	var b strings.Builder
	b.WriteString("grammar")
	if e.Code != "" {
		b.WriteString("/")
		b.WriteString(e.Code)
	}
	b.WriteString(": ")
	b.WriteString(e.Msg)
	return b.String()
}

func formatAdapterError(e *AdapterError) string {
	return "adapter: " + e.Msg
}
