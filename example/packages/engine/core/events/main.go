// Working with engine/core/Events

package main

import (
	"context"
	"fmt"

	"github.com/pt-main/lc/engine/core"
)

func test1() { // Basic usage
	fmt.Println("=== test1 ===")
	e := core.NewEvents(context.Background())

	// add event 'test'
	e.NewEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Test event is active")
		return nil
	})

	// call event with empty input and canWorkWithoutHandler=false
	fmt.Println(e.CallEvents(&core.EventInput{}, "test", false)) // "Test event is active\n<nil>"

	// call unregistered event with canWorkWithoutHandler=false
	fmt.Println(e.CallEvents(nil, "unknown", false)) // "Event 'unknown' is not found.""

	// call unregistered event with canWorkWithoutHandler=true
	fmt.Println(e.CallEvents(&core.EventInput{}, "unknown", true)) // <nil>
}

func test2() { // Core events
	fmt.Println("=== test2 ===")
	e := core.NewEvents(context.Background())

	// while event is not registred, it's created as core event
	e.NewEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Test core event is active")
		return nil
	})

	// after creating core event you can append new handlers after core event
	e.NewEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Call's after test core event")
		return nil
	})

	// and you can creating events before core event
	e.NewEventBefore("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Call's before test core event")
		return nil
	})

	// you can add more events after/before core event
	e.NewEventBefore("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("(1) Call's before test core event")
		return nil
	})

	fmt.Println(e.CallEvents(nil, "test", false)) // <nil>

	// And you can replace event
	e.ReplaceEvent("test")
	fmt.Println(e.CallEvents(nil, "test", false)) // "Event 'test' is not found."
}

func test3() { // Core events deeper
	fmt.Println("=== test3 ===")
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

	// Creating EventTools for work with core events
	et := core.EventsTools{
		Events: e,
	}

	// Change only core event
	et.ChangeCoreEvent("test", func(e *core.Events, ei *core.EventInput) core.ErrorInterface {
		fmt.Println("Not test core event")
		return nil
	})

	fmt.Println(e.CallEvents(nil, "test", false))
}

func test4() { // Hard example
	fmt.Println("=== test4 ===")
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
	test1()
	test2()
	test3()
	test4()
}

/*
=== test1 ===
Test event is active
<nil>
Event 'unknown' is not found.
<nil>
=== test2 ===
(1) Call's before test core event
Call's before test core event
Test core event is active
Call's after test core event
<nil>
Event 'test' is not found.
=== test3 ===
Call's before test core event
Not test core event
Call's after test core event
<nil>
=== test4 ===
Before: 1
Call's before test core event. 1 <nil>
Test core event is active. 4 <nil>
Call's after test core event. 2 <nil>
<nil>
After: 3
*/
