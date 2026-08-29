// Work with engine.core.Scope

package main

import (
	"fmt"

	"github.com/pt-main/lc/engine/core"
)

func test1() {
	s := core.ScopeType{} // just map[string]any
	s["Key"] = 0
	fmt.Println(core.ScopeGet[int](s, "Key")) // 0
}

func main() {
	test1()
}
