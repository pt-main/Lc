package byteParsing

import (
	"reflect"
	"testing"

	"github.com/pt-main/lc/tooling/bytecode"
)

func TestParser1_Parse(t *testing.T) {
	config := Parser1Config{
		GConfig: bytecode.GenerationConfig{
			CommandBytelen:   1,
			ArgscountBytelen: 1,
			ArglenBytelen:    1,
			Endianess:        bytecode.LittleEndian,
		},
		Shifter: bytecode.Shift{},
	}
	parser := &Parser1{Config: config}

	code := []byte{0x01, 0x01, 0x03, 0x61, 0x62, 0x63}
	nodes, err := parser.Parse(code)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(nodes) != 1 {
		t.Fatalf("expected 1 node, got %d", len(nodes))
	}
	node := nodes[0]
	if !reflect.DeepEqual(node.Switch, []byte{0x01}) {
		t.Errorf("switch = %v, want [0x01]", node.Switch)
	}
	if len(node.Args) != 1 {
		t.Fatalf("expected 1 arg, got %d", len(node.Args))
	}
	if !reflect.DeepEqual(node.Args[0], []byte{0x61, 0x62, 0x63}) {
		t.Errorf("arg = %v, want [0x61 0x62 0x63]", node.Args[0])
	}
}

func TestParser1_Parse_Error(t *testing.T) {
	config := Parser1Config{
		GConfig: bytecode.GenerationConfig{
			CommandBytelen:   1,
			ArgscountBytelen: 1,
			ArglenBytelen:    1,
			Endianess:        bytecode.LittleEndian,
		},
		Shifter: bytecode.Shift{},
	}
	parser := &Parser1{Config: config}
	code := []byte{0x01, 0x01, 0x03}
	_, err := parser.Parse(code)
	if err == nil {
		t.Error("expected error, got nil")
	}
}
