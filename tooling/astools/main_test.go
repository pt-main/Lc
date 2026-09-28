package astools

import (
	"testing"

	"github.com/pt-main/lc/parsing/stringParsing"
)

func nodeWithChildren(sw string, children ...stringParsing.ParsedNode) *stringParsing.ParsedNode {
	return &stringParsing.ParsedNode{
		Switch:   sw,
		Raw:      sw,
		Metadata: map[string]interface{}{"children": children},
	}
}

func deref(n *stringParsing.ParsedNode) stringParsing.ParsedNode { return *n }

func sampleTree() *stringParsing.ParsedNode {
	return nodeWithChildren("root",
		deref(nodeWithChildren("a", deref(nodeWithChildren("b")))),
		deref(nodeWithChildren("c")),
	)
}

func TestFindChildPointsIntoTheTree(t *testing.T) {
	root := sampleTree()
	first := FindChild(root, "a")
	second := FindChild(root, "a")
	if first == nil || second == nil {
		t.Fatal("FindChild returned nil for a present child")
	}
	if first != second {
		t.Errorf("FindChild must return the same node both times, got %p and %p", first, second)
	}
	first.Raw = "MUTATED"
	if got := FindChild(root, "a").Raw; got != "MUTATED" {
		t.Errorf("a write through the returned pointer must reach the tree, got %q", got)
	}
}

func TestGetChildAtPointsIntoTheTree(t *testing.T) {
	root := sampleTree()
	child := GetChildAt(root, 0)
	if child == nil {
		t.Fatal("GetChildAt returned nil for index 0")
	}
	child.Raw = "MUTATED"
	if got := GetChildAt(root, 0).Raw; got != "MUTATED" {
		t.Errorf("a write through GetChildAt must reach the tree, got %q", got)
	}
}

func TestWalkWriteReachesTheTree(t *testing.T) {
	root := sampleTree()
	err := Walk(root, func(n *stringParsing.ParsedNode) error {
		if n.Switch == "b" {
			n.Raw = "REWRITTEN"
		}
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	b := FindChild(FindChild(root, "a"), "b")
	if b == nil || b.Raw != "REWRITTEN" {
		t.Errorf("Walk must expose the real nodes, got %+v", b)
	}
}

func TestWalkWithPathBuildsRealPaths(t *testing.T) {
	root := sampleTree()
	want := map[string][]string{
		"root": {"root"},
		"a":    {"root", "a"},
		"b":    {"root", "a", "b"},
		"c":    {"root", "c"},
	}
	seen := map[string][]string{}
	err := WalkWithPath(root, func(n *stringParsing.ParsedNode, path []string) error {
		seen[n.Switch] = append([]string{}, path...)
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for sw, expected := range want {
		got, ok := seen[sw]
		if !ok {
			t.Errorf("node %q was never visited", sw)
			continue
		}
		if len(got) != len(expected) {
			t.Errorf("node %q path = %v, want %v", sw, got, expected)
			continue
		}
		for i := range got {
			if got[i] != expected[i] {
				t.Errorf("node %q path = %v, want %v", sw, got, expected)
				break
			}
		}
	}
}

func TestWalkWithPathDoesNotAliasSiblings(t *testing.T) {
	root := nodeWithChildren("root",
		stringParsing.ParsedNode{Switch: "a", Raw: "a", Metadata: map[string]interface{}{
			"children": []stringParsing.ParsedNode{{Switch: "x"}, {Switch: "y"}},
		}},
	)
	paths := map[string][]string{}
	err := WalkWithPath(root, func(n *stringParsing.ParsedNode, path []string) error {
		paths[n.Switch] = append([]string{}, path...)
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := paths["x"]; len(got) != 3 || got[2] != "x" {
		t.Errorf("x path = %v, want [root a x]", got)
	}
	if got := paths["y"]; len(got) != 3 || got[2] != "y" {
		t.Errorf("y path = %v, want [root a y]", got)
	}
}

func TestFindChildOnMissingChild(t *testing.T) {
	if got := FindChild(sampleTree(), "nope"); got != nil {
		t.Errorf("FindChild = %+v, want nil", got)
	}
	if got := GetChildAt(sampleTree(), 5); got != nil {
		t.Errorf("GetChildAt out of range = %+v, want nil", got)
	}
}

// GetTokenValue read Metadata["value"], but the lexer stores the matched text
// under "__value", so the branch could never fire.
func TestGetTokenValueReadsTheLexerKey(t *testing.T) {
	n := &stringParsing.ParsedNode{
		Switch:   "NUMBER",
		Metadata: map[string]interface{}{"__value": "42"},
	}
	if got := GetTokenValue(n); got != "42" {
		t.Errorf("GetTokenValue = %q, want 42", got)
	}
	if got := GetTokenValue(&stringParsing.ParsedNode{Raw: "raw wins"}); got != "raw wins" {
		t.Errorf("GetTokenValue = %q, want the raw text", got)
	}
	if got := GetTokenValue(nil); got != "" {
		t.Errorf("GetTokenValue(nil) = %q, want empty", got)
	}
}
