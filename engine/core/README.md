# engine/core

The core primitives every Lc engine is built from. Nothing here knows about
strings or bytecode; it is all generic or container-level.

| File | Contents |
|---|---|
| `universalEngineParams.go` | `UniversalEngineParams` (UEP), the object every handler receives |
| `events.go` | `Events`, the event manager, and `EventsTools` |
| `generator.go` | `Generator`, ordered output collection |
| `scope.go` | `ScopeType` and the typed, lock-aware scope helpers |
| `logger.go` | `Logger`, structured diagnostics |
| `errors.go` | `Error`, `ErrorInterface` and the error constructors |
| `types.go` | `CommandType`, `EventType`, `Option`, `EventInput` |

## UEP

`UniversalEngineParams` bundles everything a handler needs. It is embedded in
both engine types, so `engine.UEP` works everywhere.

```go
type UniversalEngineParams struct {
	Generator *Generator
	Event     EventsInterface
	Scope     ScopeType
	Logger    LoggerInterface
	Context   context.Context
}
```

`GetContext()` returns `context.Background()` when no context is set, so
handlers can call it without a nil check.

## Events

An event is an ordered list of handlers under a name. `CallEvents` runs them in
registration order and stops at the first error.

```go
events := core.NewEvents(context.Background())

events.NewEvent("cmd", handler1)        // becomes the core event
events.NewEvent("cmd", handler2)        // appended after it
events.NewEventBefore("cmd", handler3)  // inserted before it

events.CallEvents(&core.EventInput{}, "cmd", false)
```

The first handler registered for an event is its **core event**, and the
position is tracked separately from the list. `EventsTools` uses that position to
replace exactly the core handler:

```go
et := core.EventsTools{Events: events}

old, _ := et.GetCoreEvent("cmd")             // read it
et.ChangeCoreEvent("cmd", replacement)       // swap it
```

This is what `ExtensibleCLPlugin` uses to wrap the engine call loop without
copying it: the plugin installs its own core event and keeps the original to
restore on `Close()`.

`canWorkWithoutHandler` controls whether a missing event is an error. Engine
pipelines pass `true` for events that are allowed to be absent.

`SetProperty("debug", true)` turns on the `CallEventsStartEvent` and
`CallEventsEndEvent` wrapping, which traces event dispatch. It is off by default
and nothing in the framework sets it; without it, those two events are never
called. The handlers behind them are registered by `NewUniversalEngineParams` and
write through the logger, so they need the "event" level enabled in the logger to
produce output:

```go
events.SetProperty("debug", true)
logger.Logging["event"] = true
```

Thread safety: `Events` guards its maps with a `sync.RWMutex`, but
`Scope()` returns the map itself and is not a synchronized accessor.

## Generator

Collects fragments into named points and emits them in `Pipeline` order.

```go
g := core.NewGenerator(public.StringResType, []string{"pre", "main"})
g.AddString("header", "main")
g.AddString("body", "main")

res, _ := core.GetStringRes(g, "")
```

The result type is fixed at construction. Mixing types returns
`GeneratorGenerationTypeError` rather than panicking:

- `StringResType` accepts `AddString` / `AddStrings` and yields
  `GetStringRes` / `GetStringArrRes`;
- `ByteResType` accepts `AddBytes` and yields `GetBytesRes`.

`Pipeline` is a public field and can be reassigned between calls to change the
emission order without regenerating anything. A point named in `Pipeline` that
has never been written to is an error, since its position in the output is
undefined.

## Scope

A plain `map[string]interface{}`, kept as a map so that existing code writing
`scope["k"] = v` keeps compiling.

```go
val, err := core.ScopeGet[int](scope, "count")
err = core.Err("CMD", "not set").WithMeta(core.EMK(0, "string"), "count")
val, err = core.GetMetaValue[int](err, 0, "string")
```

`ScopeGet` is type-checked: a wrong type is reported as `ScopeGetError` rather
than returning zero silently.

The `*Synced` variants take a per-map lock, looked up by map pointer:

```go
core.ScopeSetSynced(scope, "key", val)
val, _ := core.ScopeGetSynced[int](scope, "key")
core.ScopeDeleteSynced(scope, "key")
```

Use them when one goroutine writes while another reads. Direct map access is not
protected, even though the surrounding types are thread-safe.

## Logger

Structured logging with per-status formats and level filtering.

```go
l := core.NewLogger("")
l.Logging["debug"] = true
l.PrintLog("debug", "message")
```

`PrintLog` writes to stdout only when `Logging[status]` is true, and always
appends to the internal slice. The format is `fmt.Sprintf` with three
placeholders: status, timestamp, message.

```go
l := core.NewLogger("STATUS:[?RD]%v [?RT]Time:[%v] [?RT]Text:[?v][?RT]")
l.Statuses["warn"] = "WARN [%v] %s\n"
l.SetStatusForm("info", "INFO [%v] %s\n")
```

`MaxLogLength` bounds the retained slice; it drops the oldest entry when
exceeded. A value of `-1` (the default) means unbounded. `GetLog()` joins the
retained entries with newlines.

## Errors

`Error` carries a code, a message, metadata and an optional cause. It satisfies
both `error` and `ErrorInterface`.

```go
func handle(node *ParsedNode) core.ErrorInterface {
	if bad(node) {
		return core.Err("MY_CODE", "bad node %v", node.Switch).
			WithMeta(core.EMK(0, "string"), node.Switch)
	}
	return nil
}
```

`Wrap` attaches a cause, which is how a layer adds context without losing the
original error:

```go
return core.Wrap("MY_CODE", err, "while processing %v", name)
```

`Format()` renders the whole chain with indentation, printing metadata inline:

```text
String:PROCESS_ERR2: EVENT_ERROR: Event handler failed
  Caused by:
    |MY_CODE: bad node
    |  Meta:
    |    0_META:string: broken
```

`GetErr` unwraps one level, converting a plain `error` cause into a `WrappedError`
so the chain stays inspectable. `GetRealErrorReverse` renders the chain
innermost-last, which reads better in logs where the root cause matters most.

Metadata keys are generated by `EMK(n, valType)` and read back by
`GetMetaValue[T]`, so the key and the expected type cannot drift apart.

`ErrExit` is a sentinel: returning it from a handler stops processing without
being treated as a failure. `example/langs/configLang` uses it for an early exit
command.

## Types

```go
type CommandType[EI, N any] func(EI, *N) ErrorInterface
type CommandMeta[EI, N any] struct {
	Handler CommandType[EI, N]
	Doc     string
}
type EventType func(*Events, *EventInput) ErrorInterface
```

A command is a function from the engine interface and a node pointer to an
error. `CommandMeta` pairs it with a doc string, which is what
`NewCommandString` takes as its last argument.

`EventInput` is an alias of `SimpleInput`, carrying an `Option` (flags and a
scope) and an arbitrary `Input`.
