package core

import (
	goerr "errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/pt-main/lc/public/errors"
)

type Error struct {
	Code  errors.ErrorCodeType
	Msg   string
	Meta  map[errors.ErrorMetaType]interface{}
	Cause error
}

type ErrorInterface interface {
	Error() string
	Format() string
	GetCode() string
	GetMsg() string
	GetMeta() map[errors.ErrorMetaType]interface{}
	Unwrap() error
}

func (e *Error) Error() string {
	return string(e.Code) + ": " + e.Msg
}

func (e *Error) GetCode() string {
	return string(e.Code)
}

func (e *Error) Unwrap() error {
	return e.Cause
}

func (e *Error) GetMsg() string {
	return e.Msg
}

func (e *Error) GetMeta() map[errors.ErrorMetaType]interface{} {
	return e.Meta
}

func (e *Error) Format() string {
	var b strings.Builder
	e.writeFull(&b, "")
	return b.String()
}

func (e *Error) writeFull(b *strings.Builder, indent string) {
	b.WriteString(indent)
	b.WriteString(e.Error())
	b.WriteString("\n")

	tab := "    |"

	if len(e.Meta) > 0 {
		b.WriteString(indent)
		b.WriteString("  Meta:\n")
		for k, v := range e.Meta {
			b.WriteString(indent)
			b.WriteString(tab)
			b.WriteString(string(k))
			b.WriteString(": ")
			b.WriteString(fmt.Sprintf("%v", v))
			b.WriteString("\n")
		}
	}

	if e.Cause != nil {
		b.WriteString(indent)
		b.WriteString("  Caused by:\n")
		if ce, ok := e.Cause.(*Error); ok {
			ce.writeFull(b, indent+tab)
		} else {
			causeText := strings.ReplaceAll(e.Cause.Error(), "\n", "\n"+indent+tab)
			b.WriteString(indent)
			b.WriteString(tab)
			b.WriteString(causeText)
			b.WriteString("\n")
		}
	}
}

func Err(code errors.ErrorCodeType, format string, args ...interface{}) *Error {
	return &Error{Code: code, Msg: fmt.Sprintf(format, args...)}
}

func Wrap(code errors.ErrorCodeType, cause error, format string, args ...interface{}) *Error {
	return &Error{Code: code, Msg: fmt.Sprintf(format, args...), Cause: cause}
}

func (e *Error) WithMeta(key errors.ErrorMetaType, value interface{}) *Error {
	if e.Meta == nil {
		e.Meta = make(map[errors.ErrorMetaType]interface{})
	}
	e.Meta[key] = value
	return e
}

// EMK builds a metadata key from an index and the expected value type, so
// the key and the type cannot drift apart.
func EMK(n int, valType string) errors.ErrorMetaType {
	return errors.ErrorMetaType(strconv.Itoa(n) + "_META:" + valType)
}

// GetMetaValue reads back a value stored under EMK(n, valType), checking the
// declared type against the one the caller expects.
//
// Err errors.CorePackageSystemError: wrong error type, missing key or
// mismatching value type.
func GetMetaValue[T any](in error, n int, valType string) (T, error) {
	var res T
	newE, ok := in.(*Error)
	if !ok {
		return res, Err(errors.CorePackageSystemError, "Invalid input: error is not lc *Error")
	}
	key := EMK(n, valType)
	val, ok := newE.Meta[key]
	if !ok {
		return res, Err(errors.CorePackageSystemError, "Key not found: %v", key)
	}
	res, ok = val.(T)
	if !ok {
		return res, Err(errors.CorePackageSystemError, "Invalid type: meta type and generic type is different")
	}
	return res, nil
}

// GetRealError renders the whole cause chain, or an empty string for nil.
func GetRealError(err error) string {
	if err == nil {
		return ""
	}
	if ce, ok := err.(ErrorInterface); ok {
		return ce.Format()
	}
	return err.Error()
}

// GetErr unwraps one level. A cause that is a plain error rather than an
// ErrorInterface is turned into an Error, so the chain stays inspectable.
func GetErr(ei ErrorInterface) ErrorInterface {
	if ei == nil {
		return nil
	}
	inner := ei.Unwrap()
	if inner == nil {
		return nil
	}
	if typed, ok := inner.(ErrorInterface); ok {
		return typed
	}
	return &Error{
		Code: errors.WrappedError,
		Msg:  inner.Error(),
		Meta: make(map[errors.ErrorMetaType]interface{}),
	}
}

var ErrExit = Err(errors.ErrExit, "")

// GetRealErrorReverse renders the chain of plain errors from the innermost
// one outwards, then the chain of a formatted error. That order reads better
// in logs, where the root cause matters most.
func GetRealErrorReverse(err error) string {
	if err == nil {
		return ""
	}
	if ce, ok := err.(ErrorInterface); ok {
		return ce.Format()
	}

	var parts []string
	for cur := err; cur != nil; cur = goerr.Unwrap(cur) {
		if ce, ok := cur.(ErrorInterface); ok {
			formatted := ce.Format()
			if len(parts) == 0 {
				return formatted
			}
			return strings.Join(reverse(parts), ": ") + ": " + formatted
		}
		parts = append(parts, cur.Error())
	}
	if len(parts) > 0 {
		return strings.Join(reverse(parts), ": ")
	}
	return err.Error()
}

func reverse(s []string) []string {
	res := make([]string, len(s))
	for i, v := range s {
		res[len(s)-1-i] = v
	}
	return res
}
