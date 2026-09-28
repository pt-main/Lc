package parser3

import (
	goerr "errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/pt-main/lc/v2/public/errors"
)

// Error codes returned by GetCode.
const (
	ParseErrCode   = "parser3"
	GrammarErrCode = "parser3/grammar"
	AdapterErrCode = "parser3/adapter"
)

// Phase names used in ParseError.Phase. They describe which part of the parse
// produced the failure and are what a grammar author sees in the output.
const (
	PhaseLexer  = "Lexer"
	PhaseStart  = "Start"
	PhaseEnd    = "End"
	PhaseExpect = "Expect"
	PhasePeek   = "Peek"
	PhaseEOF    = "EOF"
)

// ParseError is a parsing failure with enough context for both plain logging
// and rich CLI output.
type ParseError struct {
	// Where it happened.
	TokenIdx int    // index in the token stream
	TokenPos string // human-readable position, e.g. "line 3, col 5"

	// What was expected against what was found.
	Phase    string // grammar phase: "Expect", "Peek", "Lexer", "End", ...
	Expected string // expected token type, rule or value
	Got      string // actual token type or value
	Raw      string // raw text of the offending token
	Msg      string // free-form message, used when Expected and Got do not fit

	// Found lists the token types present at the failure position, best
	// candidates first. Only ChoiceExpr fills it.
	Found []string

	// Cause is the underlying error: a lexer error or a failed user action.
	Cause error
}

// Code returns the grammar phase the parse failed in.
func (e *ParseError) Code() string { return e.Phase }

func (e *ParseError) Error() string {
	var b strings.Builder
	b.WriteString(ParseErrCode)
	if e.Phase != "" {
		b.WriteString("/")
		b.WriteString(e.Phase)
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
		parts++
	}
	if len(e.Found) > 0 {
		b.WriteString(fmt.Sprintf(" (found: %s)", strings.Join(e.Found, ", ")))
	}
	if e.Msg != "" {
		if parts > 0 {
			b.WriteString(" - ")
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
	if parts == 0 && (e.Cause == nil || e.Msg == "") {
		b.WriteString("parse error")
	}
	return b.String()
}

// Unwrap returns the underlying error for errors.Is and errors.As.
func (e *ParseError) Unwrap() error { return e.Cause }

func (e *ParseError) Format() string {
	return FormatError(e, false)
}

func (e *ParseError) GetCode() string {
	return ParseErrCode
}

func (e *ParseError) GetMsg() string {
	if e.Msg != "" {
		return e.Msg
	}
	return e.Error()
}

func (e *ParseError) GetMeta() map[errors.ErrorMetaType]interface{} {
	found := ""
	if len(e.Found) > 0 {
		found = strings.Join(e.Found, ",")
	}
	return map[errors.ErrorMetaType]interface{}{
		"TokenIdx": e.TokenIdx,
		"TokenPos": e.TokenPos,
		"Code":     e.Phase,
		"Expected": e.Expected,
		"Raw":      e.Raw,
		"Got":      e.Got,
		"Found":    found,
	}
}

// GrammarError is raised when the grammar itself is wrong: an undefined rule,
// a missing start rule, or a construct that consumed nothing.
type GrammarError struct {
	// Phase names the grammar construct that failed, such as "NamedExpr" or
	// "ChoiceExpr".
	Phase string
	Msg   string // human-readable description
	Cause error
}

// Code returns the grammar phase the error came from.
func (e *GrammarError) Code() string { return e.Phase }

func (e *GrammarError) Error() string {
	var b strings.Builder
	b.WriteString(GrammarErrCode)
	if e.Phase != "" {
		b.WriteString("/")
		b.WriteString(e.Phase)
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

func (e *GrammarError) GetMsg() string {
	if e.Msg != "" {
		return e.Msg
	}
	return e.Error()
}

func (e *GrammarError) GetMeta() map[errors.ErrorMetaType]interface{} {
	return map[errors.ErrorMetaType]interface{}{
		"Code": e.Phase,
	}
}

func (e *GrammarError) GetCode() string {
	return GrammarErrCode
}

// AdapterError is raised by the engine adapter when the AST shape is wrong.
type AdapterError struct {
	Msg   string
	Cause error
}

func (e *AdapterError) Error() string {
	if e.Cause != nil {
		return AdapterErrCode + ": " + e.Msg + ": " + e.Cause.Error()
	}
	return AdapterErrCode + ": " + e.Msg
}

func (e *AdapterError) Unwrap() error { return e.Cause }

func (e *AdapterError) Format() string {
	return FormatError(e, false)
}

func (e *AdapterError) GetCode() string {
	return AdapterErrCode
}

func (e *AdapterError) GetMsg() string {
	if e.Msg != "" {
		return e.Msg
	}
	return e.Error()
}

func (e *AdapterError) GetMeta() map[errors.ErrorMetaType]interface{} {
	return nil
}

// AsParseError reports the outermost ParseError of the chain, if any.
func AsParseError(err error) (*ParseError, bool) {
	var pe *ParseError
	if goerr.As(err, &pe) {
		return pe, true
	}
	return nil, false
}

// AsGrammarError reports the outermost GrammarError of the chain, if any.
func AsGrammarError(err error) (*GrammarError, bool) {
	var ge *GrammarError
	if goerr.As(err, &ge) {
		return ge, true
	}
	return nil, false
}
