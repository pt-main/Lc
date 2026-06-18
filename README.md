# Lc - "Language creator" and language and dsl devkit
# <img alt="banner-low" src="https://github.com/user-attachments/assets/cd798499-aaa3-4de7-b5ae-4dbe39503453" />


[![Go Reference](https://pkg.go.dev/badge/github.com/pt-main/lc.svg)](https://pkg.go.dev/github.com/pt-main/lc)
[![Go Report Card](https://goreportcard.com/badge/github.com/pt-main/lc)](https://goreportcard.com/report/github.com/pt-main/lc)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache2.0-yellow.svg)](https://opensource.org/licenses/Apache2.0)
[![Release](https://img.shields.io/github/v/release/pt-main/lc)](https://github.com/pt-main/lc/releases)

```bash
go get github.com/pt-main/lc
```

**Lc** is a production-oriented toolkit for building language runtimes, compiler-like execution pipelines, command interpreters, and bytecode-driven processors in Go.

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

## Engine model
### String Engine
Best for command languages, script-like syntaxes, config runtimes, and text transformation flows.

Default lifecycle:
1. store input in scope;
2. parse input to `[]ParsedNode`;
3. dispatch handlers by `ParsedNode.Switch`;
4. emit output through `UEP.Generator` (if need).

### Byte Engine
Best for opcode streams, compact protocol runtimes, and binary execution loops.

Default lifecycle:
1. store input in scope (`INPUT []byte`);
2. parse input to `[]ParsedBytes`;
3. decode opcode from `ParsedBytes.Switch` bytes using configured endianness;
4. dispatch opcode handler;
5. advance instruction pointer automatically or manually.

## Quick start - String Engine
```go
package main

import (
	"fmt"
	"strings"

	"github.com/pt-main/lc"
	enginepkg "github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/stringParsing"
)

func main() {
	parser := &stringParsing.Parser2{}

	engine, err := lc.NewEngineBuilder(lc.StringEngineType).
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

## Quick start - Byte Engine
```go
package main

import (
	"fmt"

	"github.com/pt-main/lc"
	"github.com/pt-main/lc/byteParsing"
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
				Endianess:        bytecode.LittleEndian,
			},
			Shifter: bytecode.Shift{},
		},
	}

	engine, err := lc.NewEngineBuilder(lc.ByteEngineType).
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

## Industrial setup pattern
For production embedding, typical builder configuration includes:
- explicit pipeline points (e.g. `bootstrap`, `main`, `finalize`);
- shared service context;
- structured logger instance;
- preloaded scope values (tenant/env/runtime metadata);
- custom events for audit/metrics/tracing.

Example:
```go
engine, err := lc.NewEngineBuilder(lc.StringEngineType).
	WithPipeline([]string{"bootstrap", "main", "finalize"}).
	WithContext(ctx).
	WithDefaultEvents(true).
	WithLogger(logger).
	WithScope(scope).
	WithStringParser(parser).
	Build()
```

## Built-in parser modules
### `stringParsing`
- `Parser1`:
  - regex grammar matching,
  - named capture-group metadata,
  - line continuation support,
  - bracket-balance support.
- `Parser2`: compact line-oriented `command args` parser.
- `Lexer`: ordered regexp2 tokenizer with `__prev` / `__next` node links.

### `byteParsing`
- `Parser1`: binary parser with configurable field lengths and endianness.
- `ParsedBytes` model with `Switch`, `Args`, `Raw`, and `Metadata`.

## Public runtime API overview
### Processing
- `ProcessString(input string) error`
- `ProcessStringWithCtx(input string, ctx context.Context) error`
- `ProcessBytes(input []byte) error`
- `ProcessBytesWithCtx(input []byte, ctx context.Context) error`

### Command registration
- `NewCommandString(cmdSwitch string, handler ..., doc string) error`
- `NewCommandByte(opcode int, handler ..., name string, autoBytecodeIdxShift bool) error`

### Byte instruction pointer controls
- `AddToBytecodeIdx(n int)`
- `SetBytecodeIdx(n int)`
- `GetBytecodeIdx() (*int, error)`

## Execution semantics
- Event handlers run in registration order.
- Generator result follows declared pipeline order.
- `Process*WithCtx` respects cancellation/deadline.
- Default String dispatch skips unknown commands.
- Default Byte dispatch expects valid opcode/autoshift registration for processed commands.

## Reserved scope keys used by default events
String keys:
- `INPUT string`
- `PARSED []ParsedNode`

Byte keys:
- `INPUT []byte`
- `PARSED []ParsedBytes`
- `ENDIANESS int`
- `BYECODE_IDX *int`

Treat these names as reserved runtime contract keys.

## Observability
Lc provides core mechanisms for operational visibility:
- thread-safe `core.Logger`,
- event lifecycle hooks (call start/call end),
- centralized runtime scope for contextual metadata,
- structured error wrapping in default event flows.

## Repository structure
- `builder.go` - fluent builder.
- `engine.go` - universal runtime wrapper and command registration.
- `main.go` - public constructors for String/Byte engines.
- `engine/` - concrete engine implementations.
- `engine/core/` - generator, events, logger, shared params.
- `engine/events/` - default parse/call handlers.
- `parsing/` - basic parsing.
- `parsing/stringParsing/` - text parser ecosystem.
- `parsing/byteParsing/` - byte parser ecosystem.
- `tooling/bytecode/` - byte conversion, shift helpers and instructions generation.
- `tooling/plugins/` - plugins core.
- `example/` - example languages and other examples.

## License
Apache 2.0 - see `LICENSE`.

By Pt.
