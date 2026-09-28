package lc

import (
	"context"
	"testing"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
)

type nullParser struct{}

func (nullParser) Parse(code string, o ...*parsing.ParseOption) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	return nil, nil
}

func (nullParser) String() string { return "nullParser" }

func buildEngine(t *testing.T) *EngineUniversal {
	t.Helper()
	eng, err := NewEngineBuilder(public.StringEngineType, public.StringResType).
		WithStringParser(nullParser{}).Build()
	if err != nil {
		t.Fatalf("build failed: %v", err)
	}
	return eng
}

func TestEndClosesTheLifecycle(t *testing.T) {
	eng := buildEngine(t)
	if err := eng.CheckEnded(); err != nil {
		t.Fatalf("a fresh engine must not be ended: %v", err)
	}
	if err := eng.End(); err != nil {
		t.Fatalf("End() returned %v, want nil", err)
	}
	if err := eng.CheckEnded(); err == nil {
		t.Error("CheckEnded must report the engine as ended after End()")
	}
	if err := eng.ProcessString("x"); err == nil {
		t.Error("processing after End() must be refused")
	}
}

func TestEndIsIdempotent(t *testing.T) {
	eng := buildEngine(t)
	if err := eng.End(); err != nil {
		t.Fatalf("first End(): %v", err)
	}
	if err := eng.End(); err == nil {
		t.Error("a second End() must report that the lifecycle already ended")
	}
}

func TestEndWithoutCancelCauseDoesNotPanic(t *testing.T) {
	eng := buildEngine(t)
	eng.CtxCancelCause = nil
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("End panicked without a cancel cause: %v", r)
		}
	}()
	if err := eng.End(); err != nil {
		t.Fatalf("End() = %v, want nil", err)
	}
}

func TestEndCancelsTheContext(t *testing.T) {
	eng, err := NewEngineBuilder(public.StringEngineType, public.StringResType).
		WithStringParser(nullParser{}).
		WithContext(context.Background()).
		Build()
	if err != nil {
		t.Fatalf("build failed: %v", err)
	}
	if eng.CtxCancelCause == nil {
		t.Fatal("WithContext must install a cancel cause")
	}
	if err := eng.End(); err != nil {
		t.Fatalf("End() = %v, want nil", err)
	}
	if eng.Context.Err() == nil {
		t.Error("End() must cancel the engine context")
	}
}
