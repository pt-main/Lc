// Working with engine.core.Scope: the plain map, the typed getter, and the
// synchronized helpers that are safe to use from concurrent event handlers.

package main

import (
	"fmt"
	"sync"

	"github.com/pt-main/lc/v2/engine/core"
)

// ScopeType is a plain map, so a plain write keeps working.
func plainMap() {
	s := core.ScopeType{}
	s["Key"] = 0
	fmt.Println(core.ScopeGet[int](s, "Key"))
}

// ScopeGet is type checked: a missing key and a wrong type are both reported
// instead of silently returning a zero value.
func typedGetters() {
	s := core.ScopeType{}
	if _, err := core.ScopeGet[int](s, "missing"); err != nil {
		fmt.Println("missing key ->", err.GetCode())
	}
	core.ScopeSetSynced(s, "text", "value")
	if _, err := core.ScopeGet[int](s, "text"); err != nil {
		fmt.Println("wrong type  ->", err.GetCode())
	}
	got, err := core.ScopeGet[string](s, "text")
	fmt.Println("right type  ->", got, err)
}

// The *Synced helpers take the lock of the scope they are given, so one
// goroutine can write while another reads.
func concurrentAccess() {
	s := core.ScopeType{}
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			core.ScopeSetSynced(s, fmt.Sprintf("k%d", i), i)
		}(i)
		go func() {
			defer wg.Done()
			_, _ = core.ScopeGetSynced[int](s, "k0")
		}()
	}
	wg.Wait()
	fmt.Println("concurrent scope access: 50 keys,", len(s), "stored")
}

func main() {
	plainMap()
	typedGetters()
	concurrentAccess()
}
