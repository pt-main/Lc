// Working with engine.core.Generator.

package main

import (
	"fmt"

	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/public"
)

// A byte generator only accepts bytes, and emits them in pipeline order.
func byteGenerator() {
	fmt.Println("=== byte generator ===")
	g := core.NewGenerator(public.ByteResType, []string{"main"})
	g.AddBytes([]byte{0, 1, 2, 3, 4}, "main")
	g.AddBytes([]byte{5, 6, 7, 8, 9}, "main")
	fmt.Println(g.GetBytesRes())
}

// A string generator accepts strings and can return them either as a slice
// or joined with a separator.
func stringGenerator() {
	fmt.Println("=== string generator ===")
	g := core.NewGenerator(public.StringResType, []string{"main"})
	g.AddString("test", "main")
	g.AddString("test2", "main")
	fmt.Println(g.GetStringArrRes())
	fmt.Println(core.GetStringRes(g, ", "))
}

// Pipeline is a public field, so the emission order can be changed between
// calls without regenerating anything.
func pipelineReordering() {
	fmt.Println("=== pipeline reordered ===")
	g := core.NewGenerator(public.StringResType, nil)
	g.AddString("test", "1")
	g.AddString("test2", "2")

	g.Pipeline = []string{"1", "2"}
	fmt.Println(core.GetStringRes(g, ", "))

	g.Pipeline = []string{"2", "1"}
	fmt.Println(core.GetStringRes(g, ", "))
}

func main() {
	byteGenerator()
	stringGenerator()
	pipelineReordering()
}

/*
=== byte generator ===
[0 1 2 3 4 5 6 7 8 9] <nil>
=== string generator ===
[test test2] <nil>
test, test2 <nil>
=== pipeline reordered ===
test, test2 <nil>
test2, test <nil>
*/
