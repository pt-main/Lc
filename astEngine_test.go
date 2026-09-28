package lc

import (
	"context"
	"testing"

	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
)

type noOpByteParser struct{}

func (noOpByteParser) Parse(code []byte, o ...*parsing.ParseOption) ([]byteParsing.ParsedBytes, core.ErrorInterface) {
	return nil, nil
}

func (noOpByteParser) String() string { return "noOpByteParser" }

type astTreeParser struct{}

func (astTreeParser) Parse(code string, o ...*parsing.ParseOption) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	return []stringParsing.ParsedNode{{
		Switch: "root",
		Raw:    code,
		Metadata: map[string]interface{}{"children": []stringParsing.ParsedNode{
			{Switch: "word", Raw: "alpha", Metadata: map[string]interface{}{}},
			{Switch: "word", Raw: "beta", Metadata: map[string]interface{}{}},
		}},
	}}, nil
}

func (astTreeParser) String() string { return "astTreeParser" }

// NewAstEngine has to register the parse event, otherwise Process reads the
// parsed nodes from an empty scope key and every call fails.
func TestAstEngineProcessRunsThePipeline(t *testing.T) {
	ae := NewAstEngine(public.StringResType, []string{"main"}, true,
		astTreeParser{}, context.Background(), false, false)

	var seen []string
	add := func(name string) {
		ae.NewCommandFull(name, func(_ engine.AstEngineInterface, n *stringParsing.ParsedNode) core.ErrorInterface {
			seen = append(seen, name+":"+n.Raw)
			return nil
		}, "doc")
	}
	add("root")
	add("word")

	if err := ae.Process("in"); err != nil {
		t.Fatalf("Process returned %v", err)
	}
	want := []string{"root:in", "word:alpha", "word:beta"}
	if len(seen) != len(want) {
		t.Fatalf("handled %v, want %v", seen, want)
	}
	for i := range want {
		if seen[i] != want[i] {
			t.Errorf("handler %d = %q, want %q", i, seen[i], want[i])
		}
	}
}

// The root handler must run once: a walk that starts at the root itself
// would run it a second time.
func TestAstEngineHandlesRootOnce(t *testing.T) {
	ae := NewAstEngine(public.StringResType, []string{"main"}, true,
		astTreeParser{}, context.Background(), false, false)

	rootCalls := 0
	ae.NewCommandFull("root", func(_ engine.AstEngineInterface, _ *stringParsing.ParsedNode) core.ErrorInterface {
		rootCalls++
		return nil
	}, "doc")
	ae.NewCommandFull("word", func(_ engine.AstEngineInterface, _ *stringParsing.ParsedNode) core.ErrorInterface {
		return nil
	}, "doc")

	if err := ae.Process("in"); err != nil {
		t.Fatalf("Process returned %v", err)
	}
	if rootCalls != 1 {
		t.Errorf("root handler ran %d times, want 1", rootCalls)
	}
}

func TestAstEngineUnknownNodePolicy(t *testing.T) {
	ae := NewAstEngine(public.StringResType, []string{"main"}, true,
		astTreeParser{}, context.Background(), false, false)
	ae.NewCommandFull("root", func(_ engine.AstEngineInterface, _ *stringParsing.ParsedNode) core.ErrorInterface {
		return nil
	}, "doc")

	// "word" is not registered and unknown nodes are not tolerated
	if err := ae.Process("in"); err == nil {
		t.Error("an unregistered node must be reported when CanBeUnknown is false")
	}
}

// The automatic opcode counter has to start above every explicitly taken
// opcode, otherwise the next automatic one repeats one already in use.
func TestAutoOpcodeDoesNotCollide(t *testing.T) {
	eu, err := NewEngineBuilder(public.ByteEngineType, public.StringResType).
		WithByteParser(&noOpByteParser{}).Build()
	if err != nil {
		t.Fatalf("build failed: %v", err)
	}
	handler := core.CommandType[engine.ByteEngineInterface, byteParsing.ParsedBytes](
		func(_ engine.ByteEngineInterface, _ *byteParsing.ParsedBytes) core.ErrorInterface {
			return nil
		})
	if err := eu.NewCommandByte(1, handler, "explicit", false); err != nil {
		t.Fatal(err)
	}
	var auto []int
	for i := 0; i < 3; i++ {
		before := len(eu.ByteEngine.GetCommands())
		if err := eu.NewCommandByte(-1, handler, "auto", false); err != nil {
			t.Fatal(err)
		}
		after := eu.ByteEngine.GetCommands()
		if len(after) != before+1 {
			t.Fatalf("automatic command %d reused an opcode: %v", i, after)
		}
		auto = append(auto, len(after))
	}
}
