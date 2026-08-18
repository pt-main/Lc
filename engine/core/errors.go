package core

import (
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
	var b strings.Builder
	b.WriteString(string(e.Code))
	b.WriteString(": ")
	b.WriteString(e.Msg)
	return b.String()
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

	b.WriteString("----+")
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

// Error Meta Key n
func EMK(n int, valType string) errors.ErrorMetaType {
	return errors.ErrorMetaType(strconv.Itoa(n) + "_META:" + valType)
}

func GetMetaValue[T any](in error, n int, valType string) (res T, err error) {
	newE, ok := in.(*Error)
	if !ok {
		err = Err(errors.CorePackageSystemError, "Invalid input: error is not lc *Error")
		return
	}
	key := EMK(n, valType)
	val, ok := newE.Meta[key]
	if !ok {
		err = Err(errors.CorePackageSystemError, "Key not found: %v", key)
		return
	}
	res, ok = val.(T)
	if !ok {
		err = Err(errors.CorePackageSystemError, "Invalid type: meta type and generic type is different")
		return
	}
	return
}

func GetRealError(err error) string {
	if err != nil {
		errText := err.Error()
		if ce, ok := err.(ErrorInterface); ok {
			errText = ce.Format()
		}
		return errText
	}
	return ""
}

var ErrExit = Err(errors.ErrExit, "")
