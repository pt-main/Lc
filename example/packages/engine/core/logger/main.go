// Working with engine.core.Logger.

package main

import (
	"fmt"
	"strings"

	"github.com/pt-main/lc/v2/engine/core"
)

// PrintLog writes to stdout only for the statuses enabled in Logging, but
// always keeps the line in the internal log.
func levelFiltering() {
	fmt.Println("=== level filtering ===")
	l := core.NewLogger("")
	l.Logging["info"] = true

	l.PrintLog("info", "test info log")
	l.PrintLog("debug", "test debug log")
}

// MaxLogLength bounds the retained log: once exceeded, the oldest line is
// dropped.
func logTrimming() {
	fmt.Println("=== log trimming ===")
	// The format takes three values: status, time and text
	l := core.NewLogger("Status:[?RD]%v [?RT]Time:[?BE][%v] [?RT]Text:[?GN][%v][?RT]")

	l.MaxLogLength = 4
	l.Logging["info"] = true

	// 12 lines are written, only the last 4 are kept
	l.PrintLog("info", "first info log")
	for range 10 {
		l.PrintLog("debug", "test debug log")
	}
	l.PrintLog("info", "last info log")

	fmt.Println("\n[" + strings.Join(l.Log, "]\n[") + "]\n")
	fmt.Println(l.GetLog())
}

func main() {
	levelFiltering()
	logTrimming()
}

/*
=== level filtering ===
info [2026-08-04 19:55:38.651721 +0000 UTC] [test info log]

=== log trimming ===
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
