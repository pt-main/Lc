<h1 align="center">Lc - Language Creator and Devkit</h1>
<p align="center">
  <img alt="lc-banner" src="https://github.com/user-attachments/assets/8fa74598-5cee-403e-a9dc-417e86d22dcd" />
</p>
<p align="center">
  <a href="https://pkg.go.dev/github.com/pt-main/lc"><img src="https://img.shields.io/badge/Go-Reference-007d9c?logo=go&logoColor=white" alt="Go Reference"></a>
  <a href="https://github.com/pt-main/lc/releases"><img src="https://img.shields.io/github/v/release/pt-main/lc?color=blue" alt="Latest release"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-Apache_2.0-yellow" alt="License Apache 2.0"></a>
  <a href="https://github.com/pt-main/lc/wiki"><img src="https://img.shields.io/badge/Project-Wiki-red" alt="Project Wiki"></a>
</p>
<p align="center">
  <a href="README.md">English</a> | <a href="README-ru.md">Русский</a>
</p>

**Lc** is a Go framework for building language runtimes: interpreters, custom
VMs, DSLs, configuration languages and bytecode-driven processors.

The framework supplies the parts that every language runtime needs and that are
tedious to write correctly: a parser layer, an event pipeline, command dispatch,
deterministic output assembly, plugins, logging and context handling. The
language itself - grammar, commands, semantics - is written by the user.

```bash
go get github.com/pt-main/lc
```

## Table of Contents

- [Scope](#scope)
- [Architecture](#architecture)
- [Extension points](#extension-points)
- [Quick start](#quick-start)
- [Engines](#engines)
- [Core concepts](#core-concepts)
- [Parsers](#parsers)
- [Plugins](#plugins)
- [Debugging and tooling](#debugging-and-tooling)
- [Context and cancellation](#context-and-cancellation)
- [Performance](#performance)
- [Examples](#examples)
- [License](#license)

## Scope

What Lc is, stated plainly.

**In scope.** The execution layer of a language: turning input into nodes,
dispatching nodes to handlers, running a bytecode loop, collecting output,
supporting plugins, and doing all of it with cancellation and error typing.

**Out of scope.** Lexer and grammar authoring tools, type checking, optimization,
compilation to native code, and standard library design. Lc parses text into a
flat or nested node structure and executes it; it does not compile to machine code
and does not ship a language.

**Not included on purpose.** There is no built-in sandbox or resource limiter.
The framework does not attempt to make untrusted bytecode safe. Anything that
processes untrusted input needs its own limits, usually via context timeouts.

A runtime built on Lc typically looks like this: a lexer or grammar produces
nodes, a command table maps node names to Go functions, and the engine walks the
input and calls them. Lc provides the walking, the dispatch, the output
collection and the lifecycle; the node structure and the command semantics belong
to the language.

## Architecture

Execution is a pipeline of events. The engine stores input, fires a parse event,
fires a call event, and handlers do the work. Nothing in that sequence is
hardcoded: the parser is supplied by the user, the event manager is an interface,
and the call loop itself can be replaced.

```text
input
  |
  v
[ Parser ]    ->  []ParsedNode / []ParsedBytes
  |
  v
[ Events ]    ->  dispatch by node name or opcode
  |
  v
[ Handlers ]  ->  UEP.Generator, UEP.Scope, UEP.Logger
```

`UEP` (`UniversalEngineParams`) is the single object a handler receives. It holds
the `Generator` for output, the `Events` manager, the `Scope` for shared state,
the `Logger` and the active `Context`. Because every handler sees the same UEP,
any component can be replaced without changing user code.

## Extension points

There are five interfaces, and all five are replaceable.

| Interface | Replaced for | Default |
|---|---|---|
| `parsing.ParserInterface[I, P]` | Reading input and producing nodes | `stringParsing`, `byteParsing` |
| `core.EventsInterface` | The event manager and call loop | `core.Events` |
| `core.LoggerInterface` | Diagnostic output | `core.Logger` |
| `core.ErrorInterface` | Error values | `core.Error` |
| `plugin.PluginInterface` | External logic with its own events | `plugin.Plugin` |

The most important one is the parser. `ParserInterface` is generic in both the
input and the node type, and the engine interface is generic in the command key:

```go
type ParserInterface[I any, P any] interface {
	Parse(I, ...*ParseOption) ([]P, core.ErrorInterface)
	String() string
}
```

**The practical extension point is a custom parser, not a custom node type.**
The engine constructors are typed to the bundled node types - `NewStringEngine`
takes a `ParserInterface[string, ParsedNode]` and `NewByteEngine` takes a
`ParserInterface[[]byte, ParsedBytes]` - so a parser that emits its own node
struct cannot be connected to the engine. Writing a custom parser means
implementing one of these interfaces and emitting `ParsedNode` or
`ParsedBytes`.

That is enough to cover a large amount of real work, because `ParsedNode` is a
deliberately loose shape:

```go
type ParsedNode struct {
	Raw      string    // exact source text
	Switch   string    // token type or command name
	Metadata ScopeType // named groups, positions, user data
}
```

A custom parser decides what `Switch` means and what goes into `Metadata`, so
arbitrary language concepts are carried as metadata on the standard node:

```go
type UpperParser struct{}

func (UpperParser) Parse(code string, opts ...*parsing.ParseOption) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	res := make([]stringParsing.ParsedNode, 0, len(code))
	for i, r := range code {
		if r < 'a' || r > 'z' {
			continue
		}
		res = append(res, stringParsing.ParsedNode{
			Switch:   "letter",
			Raw:      string(r),
			Metadata: core.ScopeType{"pos": i},
		})
	}
	return res, nil
}

func (UpperParser) String() string { return "UpperParser" }
```

It connects to the engine like any bundled parser, and the handler is an
ordinary function:

```go
engine, _ := lc.NewEngineBuilder(public.StringEngineType, public.StringResType).
	WithStringParser(UpperParser{}).
	Build()

engine.NewCommandString("letter", func(se enginepkg.StringEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
	pos, _ := node.Metadata["pos"].(int)
	return se.GetUep().Generator.AddString(fmt.Sprintf("[%d]%s", pos, node.Raw), "main")
}, "echo letter with position")

engine.ProcessString("hello")
// output: [0]h [1]e [2]l [3]l [4]o
```

What cannot be replaced: the node types the constructors accept, and therefore
the command handler signature. A runtime that needs its own node struct would
need its own engine type; that is a deliberate boundary, not a missing feature.

## Quick start

### String engine

Parses `command args` lines and writes generated text into a `Generator` point.

```go
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pt-main/lc"
	enginepkg "github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
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

	err = engine.NewCommandString("log", func(se enginepkg.StringEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
		args, _ := node.Metadata["args"].(string)
		return se.GetUep().Generator.AddString(fmt.Sprintf("Log [%v]: %v",
			time.Now().Format(time.Stamp), args), "main")
	}, "append log with timestamp")
	if err != nil {
		panic(err)
	}

	err = engine.ProcessString(strings.Join([]string{
		"log service_start",
		"log service_ready",
	}, "\n"))
	if err != nil {
		panic(err)
	}

	uep, _ := engine.GetUEP()
	out, err := core.GetStringRes(uep.Generator, "\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
```

```bash
go run ./example/readme/string
```

```text
Log [Aug  7 18:44:40]: service_start
Log [Aug  7 18:44:40]: service_ready
```

### Byte engine

Decodes a simple instruction format and dispatches opcode handlers.

```go
package main

import (
	"fmt"

	"github.com/pt-main/lc"
	enginepkg "github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/tooling/bytecode"
)

func main() {
	// instruction {
	//     [bytes : cmd] [bytes : argscount] [bytes : arglen]  [bytes arglen : arg],
	//                                       [bytes : arglen2] [bytes arglen2 : arg2]...
	// }
	parser := &byteParsing.Parser1{
		Config: byteParsing.Parser1Config{
			GConfig: bytecode.GenerationConfig{
				CommandBytelen:   1,
				ArgscountBytelen: 1,
				ArglenBytelen:    2,
				Endianness:        public.LittleEndian,
			},
			Shifter: bytecode.Shift{},
		},
	}

	engine, err := lc.NewEngineBuilder(public.ByteEngineType, public.StringResType).
		WithPipeline([]string{"main"}).
		WithByteParser(parser).
		WithDefaultEvents(true).
		Build()
	if err != nil {
		panic(err)
	}

	err = engine.NewCommandByte(1, func(be enginepkg.ByteEngineInterface, node *byteParsing.ParsedBytes) core.ErrorInterface {
		for _, arg := range node.Args {
			if err := be.GetUep().Generator.AddString(string(arg), "main"); err != nil {
				return err
			}
		}
		return nil
	}, "add to output instruction", true)
	if err != nil {
		panic(err)
	}

	code := []byte{
		0x01,              // opcode=1
		0x01,              // argsCount=1
		0x03, 0x00,        // arglen=3 (little endian, 2 bytes)
		0x61, 0x62, 0x63,  // args="abc" (3 bytes)
	}

	err = engine.ProcessBytes(code)
	if err != nil {
		panic(err)
	}

	uep, _ := engine.GetUEP()
	out, err := core.GetStringRes(uep.Generator, "")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", out)
}
```

```bash
go run ./example/readme/byte
```

```text
abc
```

## Engines

### String engine lifecycle

Takes text (source code, commands, config) and does something with it: edit it,
execute it, translate it, or generate code from it.

1. Store the input in scope under `public.StringEngineScopeInput`.
2. Parse it into `[]stringParsing.ParsedNode` (`public.StringParseEvent`).
3. Dispatch a handler by `ParsedNode.Switch` (`public.StringCallEvent`).
4. Emit output through `UEP.Generator` if needed.

### Byte engine lifecycle

Takes bytecode and executes it. Built for hot loops: parsed instructions are
converted into a compact call attribute before dispatch.

1. Store the input in scope under `public.ByteEngineScopeInput`.
2. Parse it into `[]byteParsing.ParsedBytes` (`public.ByteParseEvent`).
3. Convert to a compact call attribute for the hot loop.
4. Dispatch the opcode handler.
5. Advance the instruction pointer automatically or manually.

### AST engine

`AstEngine` is not a third backend. It implements the same
`StringEngineInterface` and reuses the same string parse event, but instead of
dispatching only top-level nodes it walks the parsed tree and calls handlers for
nested nodes as well, with path control.

```go
astEngine := lc.NewAstEngine(
	public.StringResType,      // result type
	[]string{"main"},          // pipeline
	true,                      // default events
	adapter,                   // parser
	context.Background(),
	true,                      // CanBeUnknown
	false,                     // CanMainNodeBeUnknown
)
```

Each command can declare path rules:

```go
ctx := astEngine.GetCommandCtx("block")
ctx.BreakIf = [][]string{{"block", "inline"}} // stop walking at this path
ctx.SkipIf = [][]string{{"block", "comment"}} // skip this subtree
```

`CanBeUnknown` decides whether an unregistered node type is an error or is
silently ignored, which is what makes the engine usable on real input where parts
of the tree are not interesting to the runtime.

It is built with `NewAstEngine` rather than the builder on purpose: it does not
produce an `EngineUniversal`, it is a dispatcher layered over the string parse
event, and it is not interchangeable with the string or byte engine.

## Core concepts

### UEP

`UniversalEngineParams` bundles the `Generator`, the `Events` manager, the
`Scope`, the `Logger` and the active `Context`.

```go
engine, _ := lc.NewEngineBuilder(...).
	WithScope(core.ScopeType{"env": "production"}).
	Build()

uep, _ := engine.GetUEP()
uep.Generator.AddString("hello", "main")
```

### Events

Execution is event-driven. A handler receives `(*core.Events, *core.EventInput)`.

```go
events := core.NewEvents(context.Background())
events.NewEvent("event1", handler1)        // first handler of "event1"
events.NewEvent("event1", handler2)        // appended to the end
events.NewEventBefore("event1", handler3)  // inserted at the start
// "event1" -> [handler3, handler1, handler2]
```

The first handler of an event is its **core event**. `core.EventsTools` replaces
just that handler and leaves everything registered around it intact, which is how
a plugin wraps the call loop instead of copying it:

```go
et := core.EventsTools{Events: events}

old, _ := et.GetCoreEvent(public.ByteCallHotloopEvent)
et.ChangeCoreEvent(public.ByteCallHotloopEvent, myHandler)
```

Default events used by the engines:

| Engine | Events |
|---|---|
| String | `StringParseEvent`, `StringCallEvent`, `StringCallCallLoopEvent` |
| Byte | `ByteParseEvent`, `ByteCallEvent`, `ByteCallHotloopEvent` |

The event manager itself can be replaced by implementing `core.EventsInterface`.

### Generator

Collects code fragments into named pipeline points and emits them in a fixed
order. Thread-safe, and works with strings or bytes depending on the result type.

```go
pipeline := []string{"pre", "main"}
generator := core.NewGenerator(public.StringResType, pipeline)
generator.AddStrings([]string{"string1 ", "string2."}, "main")
generator.AddStrings([]string{"string3 ", "string4. "}, "pre")
res, _ := core.GetStringRes(generator, "")
// res = "string3 string4. string1 string2."
```

`Generator.Pipeline` is public, so emission order can be changed at runtime
without regenerating anything.

### Scope

A thread-safe `map[string]interface{}` shared by event handlers, parsers and
commands, used to pass data between pipeline stages.

```go
engine, _ := lc.NewEngineBuilder(...).
	WithScope(core.ScopeType{
		"tenant_id": "prod-001",
		"env":       "production",
	}).
	Build()

// later, in a command handler:
tenant, _ := core.ScopeGet[string](se.GetUep().Scope, "tenant_id")
```

`core.ScopeGetSynced` and `core.ScopeSetSynced` take the scope lock, for scopes
written from one goroutine and read from another.

**Note:** the keys declared in the `public` package are used by the default
events. Overwriting them in a custom handler changes engine behaviour.

### Logger

Thread-safe structured logger stored in the UEP, with per-status formats and
level filtering.

```go
logger := core.NewLogger("")   // default format: "%s [%v] [%s]\n"
logger.Logging["debug"] = true  // only debug output is printed
logger.PrintLog("debug", "processing node: "+node.Switch)
```

Custom status format:

```go
logger := core.NewLogger("")
logger.Statuses["warn"] = "WARN [%v] %s\n"
logger.PrintLog("warn", "This is a warning")
```

Attach it with `WithLogger(logger)`. `MaxLogLength` bounds the retained log, and
`GetLog()` returns it as a single string.

### Errors

Errors returned by handlers are `core.ErrorInterface`, carrying a package code, a
message, metadata and a cause chain.

```go
err := engine.ProcessString("...")

var lerr *core.Error
if goerrors.As(err, &lerr) {
	fmt.Println(lerr.GetCode(), lerr.GetMsg())
}
```

`Format()` renders the whole chain with indentation, which is what the
calculator example uses to print an evaluation failure. Error codes are declared
in `public/errors`.

## Parsers

### String parsers

| Parser | Behaviour | Typical use |
|---|---|---|
| `Lexer` | Token-based lexer on `regexp2` rules, bracket balance, prev/next links | Tokenization, grammar input |
| `Parser1` | Regex grammar with line continuation and bracket balancing | Line-oriented DSLs, config formats |
| `Parser2` | Simple `command args` line parser | Prototyping, shell-like languages |
| `Parser3` | Recursive descent with combinators, actions and operator precedence | Real grammars, AST generation |

`Parser2` is the simplest:

```go
parser := &stringParsing.Parser2{}
// Input: "print hello world"
// Output: ParsedNode{Switch: "print", Metadata: {"args": "hello world"}}
```

`Parser3` provides 13 combinators (`TokenExpr`, `SequenceExpr`, `ChoiceExpr`,
`RepeatExpr`, `SeparatedRepeatExpr`, `OptionalExpr`, `AndExpr`, `NotExpr`,
`PeekExpr`, `NamedExpr`, `NodeExpr`, `ActionExpr`, `PrattExpr`), rule memoization
and typed errors with positions. Full reference:
[parsing/stringParsing/parser3](parsing/stringParsing/parser3).

### Byte parser

`byteParsing.Parser1` decodes binary instructions with configurable field lengths
and endianness.

```go
parser := &byteParsing.Parser1{
	Config: byteParsing.Parser1Config{
		GConfig: bytecode.GenerationConfig{
			CommandBytelen:   1,
			ArgscountBytelen: 1,
			ArglenBytelen:    1,
			Endianness:        public.LittleEndian,
		},
		Shifter: bytecode.Shift{},
	},
}
```

## Plugins

A plugin is a bundle of events with its own scope, registered through the plugin
manager. Plugins are not isolated: they can access the engine and the manager.

```go
import "github.com/pt-main/lc/tooling/plugin"

myPlugin := plugin.NewPlugin(
	"my_plugin",  // name
	"init_event", // called on Init
	"main_event", // called on Run
	"close_event",
	"run_result", // scope keys where Run/Call store their result
	"call_result",
	context.Background(),
)

myPlugin.Events.NewEvent("main_event", func(ev *core.Events, i *core.EventInput) core.ErrorInterface {
	// i.Input is whatever was passed to Run
	return nil
})
```

Registration and use:

```go
engine, _ := lc.NewEngineBuilder(...).
	WithPlugins(myPlugin). // calls "init_event"
	Build()

result, err := engine.Plugins.RunPlugin("my_plugin", "some input") // calls "main_event"
```

`CallPluginMethod` runs a named event and returns whatever the handler stored in
the result scope key:

```go
res, err := engine.Plugins.CallPluginMethod("profiler", "report")
```

`engine.End()` closes the engine and calls `Close()` on every plugin.

## Debugging and tooling

- `plugin` - plugin manager and the `PluginInterface` base implementation.
- `bytecode` - int/float to bytecode and back, plus the shifter used by the byte
  parser.
- `astools` - AST helpers: `GetChildren`, `FindChild`, `FindChildren`, `Walk`,
  `WalkWithPath`.
- `debugging/extensiblePlugin` - replaces the default call loop with a hookable
  one. This is the foundation the profiler is built on.
- `debugging/profiler` - per-command call counts and timings. Requires
  `ExtensibleCLPlugin`.

`ExtensibleCLPlugin` needs a reference to the engine, so it is added after
`Build()`:

```go
engine, _ := lc.NewEngineBuilder(...).
	WithStringParser(parser).
	Build()

err := engine.Plugins.AddPlugin(extensiblePlugin.New(engine))
if err != nil {
	return err
}

err = engine.Plugins.AddPlugin(profiler.New())
if err != nil {
	return err
}

report, err := engine.Plugins.CallPluginMethod("profiler", "report")
```

The profiler reports per-command numbers:

```text
Profiler report (0.00 sec total):

  String commands:
  String calls: 12 (total time: 32.848µs)
    keyval: count=9, total=26.352µs, avg=2.928µs, min=1.163µs, max=9.907µs
    section: count=2, total=3.652µs, avg=1.826µs, min=1.456µs, max=2.196µs
```

## Context and cancellation

Every `Process*` method has a `WithCtx` variant accepting a `context.Context`,
which covers timeouts, graceful shutdown and request-scoped values.

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

err := engine.ProcessStringWithCtx(input, ctx)
if goerrors.Is(err, context.DeadlineExceeded) {
	fmt.Println("Execution timed out")
}
```

A context can also be supplied at build time with `WithContext(ctx)`. Calling
`engine.End()` cancels it with a cause.

**Cancellation is cooperative.** The context is not enforced by an interrupt: a
deadline does not stop a running handler on its own. The handler observes the
context through the UEP and is expected to return when it sees it:

```go
select {
case <-se.GetUep().GetContext().Done():
	return core.Err("CMD", "cancelled")
default:
	// continue
}
```

Without that check a long-running handler runs to completion even after the
deadline has passed.

## Performance

The byte engine is built for hot loops. The raw hot loop reaches roughly
170-200M ops/s on an `i7-4770HQ`. See
[example/tests/speedtest](example/tests/speedtest) for the full benchmark and
instructions on reproducing it.

## Examples

All examples are in [example/](example).

### `readme/`

The two snippets from this README, minimal and self-contained.

```bash
go run ./example/readme/string
go run ./example/readme/byte
```

### `langs/`

Complete languages, each running on the engine.

| Example | Shows |
|---|---|
| `langs/calculator` | Float expressions through `Parser3` with an engine command, separate parse and eval error reporting |
| `langs/math` | Integer language with variables, assignment and scope-backed evaluation |
| `langs/configLang` | ini-like config to JSON with the profiler plugin attached |

```bash
$ go run ./example/langs/calculator '(2+3)**4'
Lc version - 1.6.0
Result: 625

$ go run ./example/langs/calculator '(2-3*4)+(5*2-1)*2'
Lc version - 1.6.0
Result: 8

$ go run ./example/langs/math 'a := 5
b := a * 2
a + b'
15
```

`langs/calculator` separates the two error kinds. A malformed expression reports
a parser error with a position:

```bash
$ go run ./example/langs/calculator '2 +'
Lc version - 1.6.0
Parse error:
 parser3/Expect: expected "MUL", got "PLUS" (raw: "+") at idx=2 start=2-3
```

A well-formed expression that fails during evaluation reports the wrapped error
chain instead:

```bash
$ go run ./example/langs/calculator '1/0'
Lc version - 1.6.0
Eval error:
 String:PROCESS_ERR2: EVENT_ERROR: Event handler failed
  Caused by:
    |DEFAULT_EVENTS:CONTEXT_ERROR: Error at line: "1/0"
```

`langs/configLang` prints a profiler report next to the translated JSON:

```bash
$ go run ./example/langs/configLang
Lc version - 1.6.0
Profiler report (0.00 sec total):

  String commands:
  String calls: 12 (total time: 32.848µs)
    keyval: count=9, total=26.352µs, avg=2.928µs, min=1.163µs, max=9.907µs
    section: count=2, total=3.652µs, avg=1.826µs, min=1.456µs, max=2.196µs
```

### `tests/`

| Example | Shows |
|---|---|
| `tests/parser3Test` | The parser on its own: AST as text and JSON, Pratt precedence, error formatting. Does not use the engine |
| `tests/speedtest` | Byte engine benchmarks |

```bash
go run ./example/tests/parser3Test
go test -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ ./example/tests/speedtest/byte/bench
```

### `packages/`

Core components in isolation, for when the behaviour of a single building block
is what matters.

| Example | Shows |
|---|---|
| `packages/engine/core/events` | Event basics, core events, `EventsTools`, scope sharing |
| `packages/engine/core/generator` | String and byte generation, pipeline order changed at runtime |
| `packages/engine/core/logger` | Level filtering, custom formats, log trimming |
| `packages/engine/core/other` | `Scope` read and write |

`packages/engine/engines/string` and `packages/engine/engines/byte` are
placeholders and do not do anything yet.

## Changelog

All notable changes are listed in [docs/changelog.md](docs/changelog.md), also
available in [Russian](docs/changelog-ru.md).

## License

Apache 2.0 - see [LICENSE](LICENSE).

By Pt.
