package parser3

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseError is a structured parsing error.
// It carries enough context for both plain logging and rich CLI output.
type ParseError struct {
	// Where it happened.
	TokenIdx int    // index in the token stream
	TokenPos string // human-readable position, e.g. "line 3, col 5" or "start=9-11"

	// What was expected vs what we got.
	Code     string // operation name: "Expect", "Peek", "ChoiceExpr", "PrattExpr", etc.
	Expected string // expected token type / rule / value
	Got      string // actual token type / value
	Raw      string // raw text of the offending token
	Msg      string // free-form message (used when Expected/Got don't fit)

	// Cause holds the underlying error (lexer error, user action error, etc.)
	Cause error
}

func (e *ParseError) Error() string {
	var b strings.Builder
	b.WriteString("parser3")
	if e.Code != "" {
		b.WriteString("/")
		b.WriteString(e.Code)
	}
	b.WriteString(": ")

	parts := 0
	if e.Expected != "" {
		b.WriteString("expected ")
		b.WriteString(strconv.Quote(e.Expected))
		parts++
	}
	if e.Got != "" {
		if parts > 0 {
			b.WriteString(", got ")
		} else {
			b.WriteString("got ")
		}
		b.WriteString(strconv.Quote(e.Got))
		parts++
	}
	if e.Raw != "" && e.Raw != e.Got {
		b.WriteString(fmt.Sprintf(" (raw: %s)", strconv.Quote(e.Raw)))
	}
	if e.Msg != "" {
		if parts > 0 {
			b.WriteString(" — ")
		}
		b.WriteString(e.Msg)
		parts++
	}
	if e.TokenPos != "" {
		b.WriteString(" at ")
		b.WriteString(e.TokenPos)
	}
	if e.Cause != nil {
		b.WriteString(": ")
		b.WriteString(e.Cause.Error())
	}
	if parts == 0 && e.Cause == nil {
		b.WriteString("parse error")
	}
	return b.String()
}

// Unwrap returns the underlying error for errors.Is / errors.As.
func (e *ParseError) Unwrap() error { return e.Cause }

func (e *ParseError) Format() string {
	return FormatError(e, false)
}

// GrammarError is raised when the grammar itself is misconfigured
// (undefined rule, missing start rule, etc.).
type GrammarError struct {
	Code  string // e.g. "NamedExpr", "ChoiceExpr"
	Msg   string // human-readable description
	Cause error
}

func (e *GrammarError) Error() string {
	var b strings.Builder
	b.WriteString("parser3/grammar")
	if e.Code != "" {
		b.WriteString("/")
		b.WriteString(e.Code)
	}
	b.WriteString(": ")
	b.WriteString(e.Msg)
	if e.Cause != nil {
		b.WriteString(": ")
		b.WriteString(e.Cause.Error())
	}
	return b.String()
}

func (e *GrammarError) Unwrap() error { return e.Cause }

func (e *GrammarError) Format() string {
	return FormatError(e, false)
}

// AdapterError is raised by the engine adapter when the AST shape is wrong.
type AdapterError struct {
	Msg   string
	Cause error
}

func (e *AdapterError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("parser3/adapter: %s: %v", e.Msg, e.Cause)
	}
	return "parser3/adapter: " + e.Msg
}

func (e *AdapterError) Unwrap() error { return e.Cause }

func (e *AdapterError) Format() string {
	return FormatError(e, false)
}
