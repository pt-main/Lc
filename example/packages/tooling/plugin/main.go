// Work with the plugin manager: register a plugin, run it, and call one of its
// methods. The run result and the call result use separate scope keys.

package main

import (
	"context"
	"fmt"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/tooling/plugin"
)

func main() {
	pm := plugin.NewPluginManager(nil)

	p := plugin.NewPlugin(
		"greeter",  // name
		"init",     // event called on init
		"main",     // event called on Run()
		"close",    // event called on Close()
		"run_out",  // Run() result key
		"call_out", // Call() result key
		context.Background(),
	)

	fmt.Println("scope keys:")
	fmt.Println("  ScopeRunResultKey :", p.ScopeRunResultKey)
	fmt.Println("  ScopeCallResultKey:", p.ScopeCallResultKey)

	p.Events.NewEvent("init", func(ev *core.Events, _ *core.EventInput) core.ErrorInterface {
		core.ScopeSetSynced(ev.Scope(), "greeting", "hello")
		return nil
	})

	p.Events.NewEvent("main", func(ev *core.Events, i *core.EventInput) core.ErrorInterface {
		greeting, err := core.ScopeGetSynced[string](ev.Scope(), "greeting")
		if err != nil {
			return err
		}
		core.ScopeSetSynced(ev.Scope(), "run_out", greeting+", "+i.Input.(string))
		return nil
	})

	p.Events.NewEvent("shout", func(ev *core.Events, i *core.EventInput) core.ErrorInterface {
		text, _ := core.ScopeGetSynced[string](ev.Scope(), "run_out")
		core.ScopeSetSynced(ev.Scope(), "call_out", text+"!!!")
		return nil
	})

	p.Events.NewEvent("close", func(ev *core.Events, _ *core.EventInput) core.ErrorInterface {
		core.ScopeDeleteSynced(ev.Scope(), "greeting")
		return nil
	})

	if err := pm.AddPlugin(p); err != nil {
		panic(err)
	}
	fmt.Println("\nregistered plugins:", len(pm.Plugins))

	runOut, err := pm.RunPlugin("greeter", "world")
	if err != nil {
		panic(err)
	}
	fmt.Println("Run() ->", runOut)

	callOut, err := pm.CallPluginMethod("greeter", "shout")
	if err != nil {
		panic(err)
	}
	fmt.Println("Call() ->", callOut)

	if err := pm.End(); err != nil {
		panic(err)
	}
	fmt.Println("\nafter End(): registered plugins:", len(pm.Plugins))
}
