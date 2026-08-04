// Working with engine.core.Generator

package main

import (
	"fmt"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/public"
)

func test1() { // Byte generator
	fmt.Println("=== test1 ===")
	g := core.NewGenerator(public.ByteResType, []string{"main"})
	// can't add string, only bytes
	g.AddBytes([]byte{0, 1, 2, 3, 4}, "main")
	g.AddBytes([]byte{5, 6, 7, 8, 9}, "main")
	fmt.Println(g.GetBytesRes()) // [0 1 2 3 4 5 6 7 8 9]
}

func test2() { // String generator
	fmt.Println("=== test2 ===")
	g := core.NewGenerator(public.StringResType, []string{"main"})
	// can't add bytes, only string
	g.AddString("test", "main")
	g.AddString("test2", "main")
	fmt.Println(g.GetStringArrRes())        // [test test2]
	fmt.Println(core.GetStringRes(g, ", ")) // test, test2
}

func test3() { // non-linear generation
	fmt.Println("=== test3 ===")
	g := core.NewGenerator(public.StringResType, nil)
	g.AddString("test", "1")
	g.AddString("test2", "2")

	g.Pipeline = []string{"1", "2"}
	fmt.Println(core.GetStringRes(g, ", ")) // test, test2

	g.Pipeline = []string{"2", "1"}
	fmt.Println(core.GetStringRes(g, ", ")) // test2, test
}

func main() {
	test1()
	test2()
	test3()
}

/*
=== test1 ===
[0 1 2 3 4 5 6 7 8 9] <nil>
=== test2 ===
[test test2] <nil>
test, test2 <nil>
=== test3 ===
test, test2 <nil>
test2, test <nil>
*/
