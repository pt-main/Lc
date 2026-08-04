// Work with engine.core.Logger

package main

import (
	"fmt"
	"strings"

	"github.com/pt-main/lc/engine/core"
)

func test1() { // Simple example
	fmt.Println("=== test1 ===")
	l := core.NewLogger("") // logger with default status form

	// setup logging
	l.Logging["info"] = true

	// log will print if l.Logging["info"] is true
	l.PrintLog("info", "test info log")   // printing
	l.PrintLog("debug", "test debug log") // not printing
}

func test2() { // Hard example
	fmt.Println("=== test2 ===")
	// change default status form
	// 3 values input, colored output
	// using tap.colors for coloring text
	l := core.NewLogger("Status:[?RD]%v [?RT]Time:[?BE][%v] [?RT]Text:[?GN][%v][?RT]")

	l.MaxLogLength = 4 // set the saving logs limit to 4
	l.Logging["info"] = true

	// add 12 logs
	l.PrintLog("info", "first info log")
	for range 10 {
		l.PrintLog("debug", "test debug log")
	}
	l.PrintLog("info", "last info log")

	// print only 4 logs
	fmt.Println("\n[" + strings.Join(l.Log, "]\n[") + "]\n") // print all logs
	// or just
	fmt.Println(l.GetLog())
}

func main() {
	test1()
	test2()
}

/*
=== test1 ===
info [2026-08-04 19:55:38.651721 +0000 UTC] [test info log]

=== test2 ===
Status:info Time:[2026-08-04 19:55:38.651797 +0000 UTC] Text:[first info log]
Status:info Time:[2026-08-04 19:55:38.651977 +0000 UTC] Text:[last info log]

[Status:debug Time:[2026-08-04 19:55:38.651937 +0000 UTC] Text:[test debug log]]
[Status:debug Time:[2026-08-04 19:55:38.65195 +0000 UTC] Text:[test debug log]]
[Status:debug Time:[2026-08-04 19:55:38.651963 +0000 UTC] Text:[test debug log]]
[Status:info Time:[2026-08-04 19:55:38.651977 +0000 UTC] Text:[last info log]]

Status:debug Time:[2026-08-04 19:55:38.651937 +0000 UTC] Text:[test debug log]
Status:debug Time:[2026-08-04 19:55:38.65195 +0000 UTC] Text:[test debug log]
Status:debug Time:[2026-08-04 19:55:38.651963 +0000 UTC] Text:[test debug log]
Status:info Time:[2026-08-04 19:55:38.651977 +0000 UTC] Text:[last info log]
*/
