package stringParsing

import (
	"regexp"
	"testing"

	"github.com/pt-main/lc/parsing"
)

// Parse dereferenced opts[0].UEP without checking it, so a caller that passed
// an empty option (or no option at all) crashed instead of parsing.
func TestParser1_ParseOptionIsOptional(t *testing.T) {
	p := NewParser1([]GrammarRule{
		{Type: "CMD", Pattern: regexp.MustCompile(`^cmd .*$`)},
	}, Parser1Config{})

	cases := map[string][]*parsing.ParseOption{
		"empty option": {&parsing.ParseOption{}},
		"nil option":   {nil},
		"no option":    nil,
	}
	for name, opts := range cases {
		nodes, err := p.Parse("cmd a", opts...)
		if err != nil {
			t.Errorf("%s: Parse returned %v", name, err)
			continue
		}
		if len(nodes) != 1 || nodes[0].Switch != "CMD" {
			t.Errorf("%s: got %d nodes, want one CMD node", name, len(nodes))
		}
	}
}

func TestParser2_ParseOptionIsOptional(t *testing.T) {
	p := &Parser2{}
	nodes, err := p.Parse("cmd a", nil)
	if err != nil {
		t.Fatalf("Parse returned %v", err)
	}
	if len(nodes) != 1 || nodes[0].Switch != "cmd" {
		t.Errorf("got %d nodes, want one cmd node", len(nodes))
	}
}
