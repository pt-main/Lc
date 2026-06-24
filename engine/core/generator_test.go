package core

import (
	"reflect"
	"testing"

	"github.com/pt-main/lc/public"
)

func TestGenerator_AddAndGetString(t *testing.T) {
	gen := NewGenerator(public.StringResType, []string{"pre", "main"})

	err := gen.AddString("hello", "pre")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	err = gen.AddStrings([]string{" ", "world"}, "main")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	res, err := GetStringRes(gen, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "hello world"
	if res != expected {
		t.Errorf("got %q, want %q", res, expected)
	}
}

func TestGenerator_AddBytes(t *testing.T) {
	gen := NewGenerator(public.ByteResType, []string{"main"})
	err := gen.AddBytes([]byte{0x01, 0x02}, "main")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	res, err := gen.GetBytesRes()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []byte{0x01, 0x02}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("got %v, want %v", res, expected)
	}
}

func TestGenerator_WrongType(t *testing.T) {
	gen := NewGenerator(public.StringResType, []string{"main"})
	err := gen.AddBytes([]byte{0x01}, "main")
	if err == nil {
		t.Error("expected error, got nil")
	}
}
