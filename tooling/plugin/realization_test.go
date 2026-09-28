package plugin

import (
	"testing"

	"github.com/pt-main/lc/v2/engine/core"
)

func TestNewPluginKeepsResultKeysApart(t *testing.T) {
	p := NewPlugin("n", "init", "main", "close", "RUN_KEY", "CALL_KEY", nil)
	if p.ScopeRunResultKey != "RUN_KEY" {
		t.Errorf("ScopeRunResultKey = %q, want %q", p.ScopeRunResultKey, "RUN_KEY")
	}
	if p.ScopeCallResultKey != "CALL_KEY" {
		t.Errorf("ScopeCallResultKey = %q, want %q", p.ScopeCallResultKey, "CALL_KEY")
	}
}

func TestPluginCallReturnsTheCallResult(t *testing.T) {
	p := NewPlugin("n", "init", "main", "close", "RUN_KEY", "CALL_KEY", nil)
	p.Events.NewEvent("main", func(ev *core.Events, _ *core.EventInput) core.ErrorInterface {
		core.ScopeSetSynced(ev.Scope(), "RUN_KEY", "run output")
		return nil
	})
	p.Events.NewEvent("shout", func(ev *core.Events, _ *core.EventInput) core.ErrorInterface {
		core.ScopeSetSynced(ev.Scope(), "CALL_KEY", "call output")
		return nil
	})

	if _, err := p.Run("input"); err != nil {
		t.Fatalf("Run: %v", err)
	}
	res, err := p.Call("shout")
	if err != nil {
		t.Fatalf("Call: %v", err)
	}
	if res != "call output" {
		t.Errorf("Call() = %v, want %q (it must not read the run key)", res, "call output")
	}
}
