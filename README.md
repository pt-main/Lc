<h1 align="center">Lc — Language Creator & Devkit</h1>
<p align="center">
  <img alt="banner-low" src="https://github.com/user-attachments/assets/cd798499-aaa3-4de7-b5ae-4dbe39503453" />
</p>
<p align="center">
  <a href="https://pkg.go.dev/github.com/pt-main/lc"><img src="https://img.shields.io/badge/Go-Reference-007d9c?logo=go&logoColor=white"></a>
  <a href="https://goreportcard.com/report/github.com/pt-main/lc"><img src="https://goreportcard.com/badge/github.com/pt-main/lc"></a>
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
Input bytecode and process that.

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
## Powerfull core and UEP (Universal Engine Params)
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

Event handlers input `interface{}, *Events`.

#### Example
```go
events := core.NewEvents(context.Background()) // new manager
events.NewEvent("event1", handler1) // create main handler in "event1" event
events.NewEvent("event1", handler2) // append handler to end of "event1"
events.NewEventBefore("event1", handler3) // append handler to start of "event1"
// "event1" - [handler3, handler1, handler2]
```

### `Generator`
Powerfull tool for codegen. 

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
}, "main")
res := core.GetStringRes(generator, "") // get code
// res = string3 string4. string1 string2.
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
