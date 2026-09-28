package core

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/pt-main/lc/public/errors"
)

type errPlain struct{}

func (errPlain) Error() string { return "plain" }

func TestGetErrNilCause(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("GetErr panicked on a nil cause: %v", r)
		}
	}()
	if got := GetErr(Err(errors.ParsingError, "no cause here")); got != nil {
		t.Errorf("GetErr of an error without a cause = %v, want nil", got)
	}
}

func TestGetErrNilInput(t *testing.T) {
	if got := GetErr(nil); got != nil {
		t.Errorf("GetErr(nil) = %v, want nil", got)
	}
}

func TestGetErrKeepsInnerErrorInterface(t *testing.T) {
	inner := Err(errors.ParsingError, "inner").WithMeta(EMK(0, "string"), "kept")
	outer := &Error{Code: errors.ParsingError, Msg: "outer", Cause: inner}

	got := GetErr(outer)
	if got != ErrorInterface(inner) {
		t.Fatalf("GetErr must return the inner error as is, got %T: %v", got, got)
	}
	if value, err := GetMetaValue[string](inner, 0, "string"); err != nil || value != "kept" {
		t.Errorf("meta must survive GetErr, got %q err %v", value, err)
	}
}

func TestGetErrWrapsPlainError(t *testing.T) {
	outer := &Error{Code: errors.ParsingError, Msg: "outer", Cause: errPlain{}}
	got := GetErr(outer)
	if got == nil {
		t.Fatal("GetErr of a non-ErrorInterface cause must still produce an error")
	}
	if got.GetCode() != string(errors.WrappedError) {
		t.Errorf("code = %q, want %q", got.GetCode(), errors.WrappedError)
	}
}

func TestCoreEventsReturnsACopy(t *testing.T) {
	e := NewEvents(context.Background())
	e.NewEvent("a", func(*Events, *EventInput) ErrorInterface { return nil })

	snapshot := e.CoreEvents()
	snapshot["injected"] = 123
	delete(snapshot, "a")

	if _, err := e.GetCoreEventIdx("injected"); err == nil {
		t.Error("writing to the returned map must not change the internal state")
	}
	if _, err := e.GetEvents("a"); err != nil {
		t.Error("deleting from the returned map must not change the internal state")
	}
}

func TestGetEventsReturnsACopy(t *testing.T) {
	e := NewEvents(context.Background())
	h2 := func(*Events, *EventInput) ErrorInterface { return nil }
	e.NewEvent("a", func(*Events, *EventInput) ErrorInterface { return nil })

	events, err := e.GetEvents("a")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	events[0] = h2
	after, _ := e.GetEvents("a")
	events = append(events, h2)
	if got, _ := e.GetEvents("a"); len(got) != 1 {
		t.Errorf("len = %d, want 1: the returned slice must not share its array", len(got))
	}
	if len(after) != 1 {
		t.Errorf("len = %d, want 1", len(after))
	}
}

func TestScopeSyncedHelpers(t *testing.T) {
	scope := make(ScopeType)
	ScopeSetSynced(scope, "k", "v")
	got, err := ScopeGet[string](scope, "k")
	if err != nil || got != "v" {
		t.Errorf("ScopeGet = %q, %v; want v, nil", got, err)
	}
	ScopeDeleteSynced(scope, "k")
	if _, err := ScopeGet[string](scope, "k"); err == nil {
		t.Error("ScopeDeleteSynced must remove the key")
	}
}

func TestScopeSyncedIsSafeUnderConcurrency(t *testing.T) {
	scope := make(ScopeType)
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			ScopeSetSynced(scope, fmt.Sprintf("k%d", i), i)
		}(i)
		go func() {
			defer wg.Done()
			_, _ = ScopeGetSynced[string](scope, "k0")
		}()
	}
	wg.Wait()
}

func TestEventsConcurrentAccess(t *testing.T) {
	e := NewEvents(context.Background())
	e.NewEvent("a", func(*Events, *EventInput) ErrorInterface { return nil })

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(3)
		go func(i int) {
			defer wg.Done()
			e.ReplaceEvent("x")
		}(i)
		go func() {
			defer wg.Done()
			e.GetCoreEventIdx("a")
		}()
		go func() {
			defer wg.Done()
			_ = e.CoreEvents()
		}()
	}
	wg.Wait()
}
