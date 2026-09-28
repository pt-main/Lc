package engine

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
)

type fakeParser struct{}

func (fakeParser) Parse(code string, o ...*parsing.ParseOption) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	return []stringParsing.ParsedNode{{Switch: "cmd", Raw: code}}, nil
}

func (fakeParser) String() string { return "fakeParser" }

func newTestAstEngine() *AstEngine {
	return &AstEngine{
		UEP:                  mustUEP(),
		Parser:               fakeParser{},
		Commands:             make(map[string]astCommandMeta),
		AstCommandCtx:        make(map[string]*AstCommandCtx),
		CanBeUnknown:         false,
		CanMainNodeBeUnknown: false,
	}
}

func mustUEP() *core.UniversalEngineParams {
	uep, _ := core.NewUniversalEngineParams(
		core.NewGenerator(public.StringResType, []string{"main"}),
		core.NewEvents(context.Background()),
		core.ScopeType{},
		core.NewLogger(""),
		context.Background(),
	)
	return uep
}

func node(sw string, children ...stringParsing.ParsedNode) stringParsing.ParsedNode {
	return stringParsing.ParsedNode{
		Switch:   sw,
		Raw:      sw,
		Metadata: map[string]interface{}{"children": children},
	}
}

func TestGetCommandCtxCreatesMissingEntry(t *testing.T) {
	ae := newTestAstEngine()
	ctx := ae.GetCommandCtx("cmd")
	if ctx == nil {
		t.Fatal("GetCommandCtx returned nil for an unknown command")
	}
	if ctx.Name != "cmd" {
		t.Errorf("ctx.Name = %q, want %q", ctx.Name, "cmd")
	}
	if again := ae.GetCommandCtx("cmd"); again != ctx {
		t.Error("GetCommandCtx must reuse the context it created")
	}
}

func TestAstMakeCommandCtxSkipIf(t *testing.T) {
	ctx := AstMakeCommandCtx("cmd", [][]string{{"a"}}, nil)
	if len(ctx.SkipIf) != 0 {
		t.Errorf("SkipIf = %v, want empty when nil is passed", ctx.SkipIf)
	}
	if len(ctx.BreakIf) != 1 {
		t.Errorf("BreakIf = %v, want the given value", ctx.BreakIf)
	}

	ctx2 := AstMakeCommandCtx("cmd", nil, [][]string{{"b"}})
	if len(ctx2.BreakIf) != 0 {
		t.Errorf("BreakIf = %v, want empty when nil is passed", ctx2.BreakIf)
	}
	if len(ctx2.SkipIf) != 1 {
		t.Errorf("SkipIf = %v, want the given value", ctx2.SkipIf)
	}
}

func TestHasCommandUnknownNodePolicy(t *testing.T) {
	cases := []struct {
		name           string
		canBeUnknown   bool
		canMainUnknown bool
		isMain         bool
		wantErr        bool
	}{
		{"unknown child is refused", false, false, false, true},
		{"unknown child is skipped", true, false, false, false},
		{"unknown main is refused", false, false, true, true},
		{"unknown main is allowed", false, true, true, false},
		{"unknown main is allowed by canBeUnknown alone", true, false, true, true},
	}
	for _, tc := range cases {
		ae := newTestAstEngine()
		ae.CanBeUnknown = tc.canBeUnknown
		ae.CanMainNodeBeUnknown = tc.canMainUnknown
		err := ae.HasCommand(tc.isMain, "nope")
		if tc.wantErr && err == nil {
			t.Errorf("%s: expected an error, got nil", tc.name)
		}
		if !tc.wantErr && err != nil {
			t.Errorf("%s: expected the node to be accepted, got %v", tc.name, err)
		}
	}
}

func TestHasCommandKnownNodeNeverFails(t *testing.T) {
	ae := newTestAstEngine()
	ae.NewCommandFull("known", func(AstEngineInterface, *stringParsing.ParsedNode) core.ErrorInterface { return nil }, "doc")
	if err := ae.HasCommand(false, "known"); err != nil {
		t.Errorf("a registered command must be accepted, got %v", err)
	}
	if err := ae.HasCommand(true, "known"); err != nil {
		t.Errorf("a registered main command must be accepted, got %v", err)
	}
}

func TestWorkSkipsUnknownNodesWhenAllowed(t *testing.T) {
	ae := newTestAstEngine()
	ae.CanBeUnknown = true
	ae.NewCommandFull("known", func(AstEngineInterface, *stringParsing.ParsedNode) core.ErrorInterface { return nil }, "doc")

	tree := node("root",
		node("known"),
		node("mystery"),
		node("known"),
	)
	if err := ae.Work([]stringParsing.ParsedNode{tree}); err != nil {
		t.Fatalf("an unknown node must be skipped, not abort the walk: %v", err)
	}
}

func TestGetCommandsReturnsACopy(t *testing.T) {
	ae := newTestAstEngine()
	ae.NewCommandFull("a", func(AstEngineInterface, *stringParsing.ParsedNode) core.ErrorInterface { return nil }, "doc")

	snapshot := ae.GetCommands()
	snapshot["injected"] = astCommandMeta{}
	delete(snapshot, "a")

	if _, ok := ae.GetCommands()["injected"]; ok {
		t.Error("writing to the returned map must not change the engine state")
	}
	if _, ok := ae.GetCommands()["a"]; !ok {
		t.Error("deleting from the returned map must not change the engine state")
	}
}

func TestGetCommandsIsSafeUnderConcurrency(t *testing.T) {
	ae := newTestAstEngine()
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			ae.NewCommandFull(fmt.Sprintf("c%d", i), func(AstEngineInterface, *stringParsing.ParsedNode) core.ErrorInterface { return nil }, "doc")
		}(i)
		go func() {
			defer wg.Done()
			for range ae.GetCommands() {
			}
		}()
	}
	wg.Wait()
}

func TestWorkVisitsNestedNodesWithoutPanic(t *testing.T) {
	ae := newTestAstEngine()
	var visited []string
	for _, name := range []string{"root", "child", "grand"} {
		current := name
		ae.NewCommandFull(current, func(AstEngineInterface, *stringParsing.ParsedNode) core.ErrorInterface {
			visited = append(visited, current)
			return nil
		}, "doc")
	}

	tree := node("root",
		node("child", node("grand")),
		node("child"),
	)
	if err := ae.Work([]stringParsing.ParsedNode{tree}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, want := range []string{"root", "child", "grand"} {
		found := false
		for _, got := range visited {
			if got == want {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("command %q was never handled, visited %v", want, visited)
		}
	}
}
