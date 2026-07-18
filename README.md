<h1 align="center">Lc — Language Creator & Devkit</h1>
<p align="center">
  <img alt="banner-low" src="https://github.com/user-attachments/assets/cd798499-aaa3-4de7-b5ae-4dbe39503453" />
</p>
<p align="center">
  <a href="https://pkg.go.dev/github.com/pt-main/lc"><img src="https://img.shields.io/badge/Go-Reference-007d9c?logo=go&logoColor=white"></a>
  <a href="https://github.com/pt-main/lc/releases"><img src="https://img.shields.io/github/v/release/pt-main/lc?color=blue"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-Apache_2.0-yellow"></a>
  <a href="https://github.com/pt-main/Lc/wiki"><img src="https://img.shields.io/badge/Project-Wiki-red"></a>
</p>


> ```bash
> go get github.com/pt-main/lc
> ```
> **Lc** is a production-oriented toolkit for building things like language tools, compiler, interpreters or bytecode-driven processors in Go.


It is intentionally straightforward to adopt, while preserving industrial runtime properties:
- explicit execution lifecycle,
- deterministic output assembly,
- context-aware cancellation,
- thread-safe core primitives,
- clear extension contracts for parsers and command handlers, plugins.

Lc does not enforce one grammar style or one VM model.  
Instead, it gives you one runtime surface with two engine backends:
- **String Engine** for text-first processing.
- **Byte Engine** for binary instruction execution.

# Engine model
## String Engine
Input string (code) and process that - edit, execute, generate code, etc.

Default lifecycle:
1. store input in scope;
2. parse input to `[]ParsedNode`;
3. dispatch handlers by `ParsedNode.Switch`;
4. emit output through `UEP.Generator` (if need).

## Byte Engine
Input bytecode and process that. Very fast hotloop (~60m+ ops/s).

Default lifecycle:
1. store input in scope;
2. parse input to `[]ParsedBytes`;
3. decode opcode from `ParsedBytes.Switch` bytes using configured endianness;
4. dispatch opcode handler;
5. advance instruction pointer automatically or manually.

# Quick start
<details> <summary>StringEngine Example</summary>

```go
package main

import (
	"fmt"
	"strings"

	"github.com/pt-main/lc"
	enginepkg "github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
)

func main() {
	parser := &stringParsing.Parser2{}

	engine, err := lc.NewEngineBuilder(public.StringEngineType, public.StringResType).
		WithPipeline([]string{"main"}).
		WithStringParser(parser).
		WithDefaultEvents(true).
		Build()
	if err != nil {
		panic(err)
	}

	err = engine.NewCommandString("print", func(se *enginepkg.StringEngine, node stringParsing.ParsedNode) error {
		args, _ := node.Metadata["args"].(string)
		return se.UEP.Generator.AddString(args, "main")
	}, "append text to output")
	if err != nil {
		panic(err)
	}

	err = engine.ProcessString(strings.Join([]string{
		"print service_start",
		"print service_ready",
	}, "\n"))
	if err != nil {
		panic(err)
	}

	out, err := engine.GetUEP().Generator.GetStringRes("\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
```

</details> <details> <summary>ByteEngine Example</summary>

```go
package main

import (
	"fmt"

	"github.com/pt-main/lc"
	"github.com/pt-main/lc/parsing/byteParsing"
	enginepkg "github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/tooling/bytecode"
	"github.com/pt-main/lc/public"
)

func main() {
	parser := &byteParsing.Parser1{
		Config: byteParsing.Parser1Config{
			GConfig: bytecode.GenerationConfig{
				CommandBytelen:   1,
				ArgscountBytelen: 1,
				ArglenBytelen:    1,
				Endianess:        public.LittleEndian,
			},
			Shifter: bytecode.Shift{},
		},
	}

	engine, err := lc.NewEngineBuilder(public.ByteEngineType, public.ByteResType).
		WithPipeline([]string{"main"}).
		WithByteParser(parser).
		WithDefaultEvents(true).
		Build()
	if err != nil {
		panic(err)
	}

	err = engine.NewCommandByte(1, func(be *enginepkg.ByteEngine, node byteParsing.ParsedBytes) error {
		return be.UEP.Generator.AddBytes(node.Raw, "main")
	}, "mirror instruction bytes", true)
	if err != nil {
		panic(err)
	}

	code := []byte{
		0x01, 0x01, 0x03, 0x61, 0x62, 0x63, // opcode=1, args=1, argLen=3, arg="abc"
	}

	err = engine.ProcessBytes(code)
	if err != nil {
		panic(err)
	}

	out, err := engine.GetUEP().Generator.GetBytesRes()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", out)
}
```
</details>




# Tools and features

## Powerful core and UEP (Universal Engine Params)
Engines core contains all necessary tools for runtime work. UEP contains then.

You can use it like:

```go
engine, _ := lc.NewStringEngine(...)
engine.UEP.Generator.AddString(...)
engine.UEP...
```

Or:

```go
engine, _ := lc.NewEngineBuilder(...).
	[...].
	Build()
engine.GetUEP().Generator.AddString(...)
engine.GetUEP()...
```

### `Events`
Engine arch is event-driven. Events can communicate with `Events.Scope`, work with context (`Events.Context`), call by pipeline. 

Event handlers input `*Events, *EventInput`.

You can override Events by implementing `core.EventsInterface`.

#### Example
```go
events := core.NewEvents(context.Background()) // new manager
events.NewEvent("event1", handler1) // create main handler in "event1" event
events.NewEvent("event1", handler2) // append handler to end of "event1"
events.NewEventBefore("event1", handler3) // append handler to start of "event1"
// "event1" - [handler3, handler1, handler2]
```

### `Generator`
Powerful tool for codegen.

Work with points pipeline for storing code in independent points. Can generate bytes or string.

#### Example

```go
pipeline := []string{"pre", "main"}
generator := core.NewGenerator([result-type], pipeline)
generator.AddStrings([]string{ // add strings to main
	"string1 ",
	"string2.",
}, "main")
generator.AddStrings([]string{ // add strings to pre
	"string3 ",
	"string4. ",
}, "pre")
res := core.GetStringRes(generator, "") // get code
// res = string3 string4. string1 string2.
```

### `Scope`
The Scope is a thread-safe `map[string]interface{}` shared across all event handlers, parsers, and commands. It serves as a runtime context for passing data between pipeline stages.

**Important:** Do not overwrite keys from `public/` package in your custom handlers unless you know exactly what you're doing — they are used by default events.

#### Custom scope usage
```go
engine, _ := lc.NewEngineBuilder(...).
	WithScope(core.ScopeType{
		"tenant_id": "prod-001",
		"env":       "production",
	}).
	Build()

// later, in your command handler:
func myHandler(se *engine.StringEngine, node stringParsing.ParsedNode) error {
	tenant, _ := core.ScopeGet[string](se.UEP.Scope, "tenant_id")
	fmt.Println("Running for tenant:", tenant)
	return nil
}
```

### `Logger`
Structured logger built into UEP. Supports status-based formatting and log level filtering.

#### Example
```go
logger := core.NewLogger("") // uses default format: "[?BE]%s[?RT] [?CN][%v][?RT] [?GN][%s][?RT]\n"
logger.Logging["debug"] = true  // enable debug output
// other logging will be disabled

// in your engine builder:
engine, _ := lc.NewEngineBuilder(...).
	WithLogger(logger).
	Build()

// in your handlers:
func myHandler(se *engine.StringEngine, node stringParsing.ParsedNode) error {
	se.UEP.Logger.PrintLog("debug", "Processing node: "+node.Switch)
	se.UEP.Logger.PrintLog("error", "Error: "+...) // disabled
	...
}
```

#### Custom status format
```go
logger := core.NewLogger("")
logger.Statuses["warn"] = "[?YW]WARN[?RT] [%v] [?RD]%s[?RT]\n" // pt-main/tap color format
logger.PrintLog("warn", "This is a warning")
```

## Plugin System
Lc has a built‑in plugin manager that allows dynamic registration and execution of external logic. Plugins are isolated via their own events and scope.

### Creating a plugin
```go
import "github.com/pt-main/lc/tooling/plugin"

myPlugin := plugin.NewPlugin(
	"my_plugin",          // name
	"init_event",         // event called on init
	"main_event",         // event called on Run()
	"close_event",        // event called on Close()
	"scope_return",   // plugin.Run (or plugin method) event can put output here
)

// Add handlers to plugin events
myPlugin.Events.NewEvent("init_event", func(ev *core.Events, i *EventInput) error {
	ev.Scope["plugin_ready"] = true
	return nil
})

myPlugin.Events.NewEvent("main_event", func(ev *core.Events, i *EventInput) error {
	// i.Input is whatever was passed to plugin.Run()
	return nil
})
```

### Registering and using a plugin
```go
engne, _ := lc.NewEngineBuilder(...).
	WithPlugins(myPlugin). // call "init_event"
	Build()

// Later, call plugin methods:
result, err := engine.Plugins.RunPlugin("my_plugin", "some input") // call "main_event"
```

## Parsers — ready‑to‑use implementations
Lc ships with several parsers for different use cases:

### StringParsing parsers

| Parser | Description | Best for |
|--------|-------------|----------|
| **Lexer** | Token-based lexer with regexp2 rules, supports bracket balancing and prev/next links | Tokenization |
| **Parser1** | Regex-based grammar with line continuation and bracket balancing | DSLs with line-oriented syntax |
| **Parser2** | Simple `command args` line parser | Quick prototyping, shell-like languages |
| **Parser3** | PEG-inspired parser with combinators (Sequence, Choice, Repeat, Optional, Named) | Complex grammars, AST generation |
| **Adapter** | `Parser3` adapter for string engine. |

#### Example: Parser2 (simplest)
```go
parser := &stringParsing.Parser2{}
// Input: "print hello world"
// Output: ParsedNode{Switch: "print", Metadata: {args: "hello world"}}
```

### ByteParsing parsers
| Parser | Description |
|--------|-------------|
| **Parser1** | Binary instruction decoder with configurable field lengths and endianness |

```go
parser := &byteParsing.Parser1{
	Config: byteParsing.Parser1Config{
		GConfig: bytecode.GenerationConfig{
			CommandBytelen:   1,
			ArgscountBytelen: 1,
			ArglenBytelen:    1,
			Endianess:        public.LittleEndian,
		},
		Shifter: bytecode.Shift{},
	},
}
```

## Context support
All `Process*` methods have `WithCtx` variants that accept `context.Context`. This allows:
- Timeout-based cancellation
- Graceful shutdown
- Request-scoped values

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

err := engine.ProcessStringWithCtx(input, ctx)
if errors.Is(err, context.DeadlineExceeded) {
	fmt.Println("Execution timed out")
}
```





# Execution semantics
- Event handlers run in registration order.
- Generator result follows declared pipeline order.
- `Process[*]WithCtx` respects cancellation/deadline.
- Default String dispatch skips unknown commands.
- Default Byte dispatch expects valid opcode/autoshift registration for processed commands.

# Observability
Lc provides core mechanisms for operational visibility:
- thread-safe `core.Logger`,
- event lifecycle hooks (call start/call end),
- centralized runtime scope for contextual metadata,
- structured error wrapping in default event flows.

## License
Apache 2.0 - see `LICENSE`.

By Pt.
