// Working with engine/core/Events.

package main

import (
	"context"
	"fmt"

	"github.com/pt-main/lc/engine/core"
)

// Basic usage: register one handler, then call the event.
func basicUsage() {
	fmt.Println("=== basic usage ===")
	e := core.NewEvents(context.Background())

	e.NewEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Test event is active")
		return nil
	})

	fmt.Println(e.CallEvents(&core.EventInput{}, "test", false))

	// canWorkWithoutHandler=false makes a missing event an error
	fmt.Println(e.CallEvents(nil, "unknown", false))

	// canWorkWithoutHandler=true makes a missing event a no-op
	fmt.Println(e.CallEvents(&core.EventInput{}, "unknown", true))
}

// Core events: the first registered handler is the core event. NewEvent
// appends after it, NewEventBefore inserts in front of it.
func coreEventOrdering() {
	fmt.Println("=== core events ===")
	e := core.NewEvents(context.Background())

	e.NewEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Test core event is active")
		return nil
	})

	e.NewEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Call's after test core event")
		return nil
	})

	e.NewEventBefore("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Call's before test core event")
		return nil
	})

	// A second insert before the core event goes in front of the first one
	e.NewEventBefore("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("(1) Call's before test core event")
		return nil
	})

	fmt.Println(e.CallEvents(nil, "test", false))

	// ReplaceEvent drops the whole event, so calling it now fails
	e.ReplaceEvent("test")
	fmt.Println(e.CallEvents(nil, "test", false))
}

// EventsTools changes the core event in place, leaving the handlers
// registered around it untouched.
func coreEventReplacement() {
	fmt.Println("=== core events replaced ===")
	e := core.NewEvents(context.Background())

	e.NewEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Test core event is active")
		return nil
	})

	e.NewEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Call's after test core event")
		return nil
	})

	e.NewEventBefore("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Call's before test core event")
		return nil
	})

	et := core.EventsTools{
		Events: e,
	}
	et.ChangeCoreEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Not test core event")
		return nil
	})

	fmt.Println(e.CallEvents(nil, "test", false))
}

// Handlers of the same event share the scope, so they can pass data along
// without a return value.
func scopeSharing() {
	fmt.Println("=== scope sharing ===")
	e := core.NewEvents(context.Background())

	e.Scope()["numStr"] = "1"

	e.NewEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Print("Test core event is active. ")
		fmt.Println(core.ScopeGet[string](e.Scope(), "numStr"))
		e.Scope()["numStr"] = "2"
		return nil
	})

	e.NewEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Print("Call's after test core event. ")
		fmt.Println(core.ScopeGet[string](e.Scope(), "numStr"))
		e.Scope()["numStr"] = "3"
		return nil
	})

	e.NewEventBefore("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Print("Call's before test core event. ")
		fmt.Println(core.ScopeGet[string](e.Scope(), "numStr"))
		e.Scope()["numStr"] = "4"
		return nil
	})

	numStr, _ := core.ScopeGet[string](e.Scope(), "numStr")
	fmt.Println("Before: " + numStr)
	fmt.Println(e.CallEvents(nil, "test", false))
	numStr, _ = core.ScopeGet[string](e.Scope(), "numStr")
	fmt.Println("After: " + numStr)
}

func main() {
	basicUsage()
	coreEventOrdering()
	coreEventReplacement()
	scopeSharing()
}

/*
=== basic usage ===
Test event is active
<nil>
Event 'unknown' is not found.
<nil>
=== core events ===
(1) Call's before test core event
Call's before test core event
Test core event is active
Call's after test core event
<nil>
Event 'test' is not found.
=== core events replaced ===
Call's before test core event
Not test core event
Call's after test core event
<nil>
=== scope sharing ===
Before: 1
Call's before test core event. 1 <nil>
Test core event is active. 4 <nil>
Call's after test core event. 2 <nil>
<nil>
After: 3
*/
