# tooling/debugging

Two system plugins that make the engine's call loop observable and hookable.

| Package | Role |
|---|---|
| `extensiblePlugin` | Replaces the default call loop with one that fires hooks around each command |
| `profiler` | Builds on those hooks to collect per-command timings |

`example/langs/configLang` uses both.

## extensiblePlugin

The engine's call loop is a single event handler, which normally makes it
awkward to observe: to time a command you would have to reimplement dispatch.
This plugin replaces that one handler, so everything else - the events, the
commands, the parser - stays as it was.

```go
engine, _ := lc.NewEngineBuilder(...).WithStringParser(parser).Build()

err := engine.Plugins.AddPlugin(extensiblePlugin.New(engine))
```

`New` takes the engine, so this plugin cannot go through `WithPlugins` at build
time: the engine does not exist yet then. On `Init` it looks up the current core
event of the call loop (`StringCallCallLoopEvent` for a string engine,
`ByteCallHotloopEvent` for a byte engine), stores it, and installs its own in its
place. `Close` restores the original.

```go
name := public.StringCallCallLoopEvent // or ByteCallHotloopEvent
```

### Hooks

The replacement loop fires four events: two around the whole loop, two around
each command.

| Constant | Fires |
|---|---|
| `CLEPreEvent` | Once, before the loop starts |
| `CLEInPreEvent` | Before each command handler |
| `CLEInPostEvent` | After each command handler returns |
| `CLEPostEvent` | Once, after the loop finishes |

All four are called with `canWorkWithoutHandler = true`, so registering nothing
costs nothing at runtime. The return value of the hook events is discarded: a
hook that returns an error does not stop the loop. Only a failing command handler
stops it, wrapped with `ExtensiblePluginError` metadata.

The profiler uses `CLEInPreEvent` and `CLEInPostEvent`, which is why it can time
individual commands.

### Data

Before the hooks, the loop stores the current position in scope under
`CLEScopeData`:

| Engine | Type |
|---|---|
| String | `SCLEData` |
| Byte | `BCLEData` |

Both are `CLEData[I, P, E]`:

```go
type CLEData[I, P, E any] struct {
	Input  I      // raw input of this run
	Idx    *int   // current position, shared with the engine
	Parsed []P    // parsed nodes
	PLen   int    // length of Parsed
	Ctx    context.Context
	E      E      // the engine itself
}
```

This is the same `*int` the engine uses for its instruction pointer, so reading
it identifies exactly which command is about to run. `SCLEData` exposes
`Parsed` as `[]ParsedNode` and `E` as `StringEngineInterface`; `BCLEData` exposes
`[]ByteCallAttr` and `ByteEngineInterface`.

The loop writes this data, then reads it back. A hook handler that replaces the
value in scope therefore changes what the loop will use next - the loop picks up
whatever is in `CLEScopeData` after `CLEPreEvent` returns, including a modified
`Idx` or `Parsed`. That is the supported way to influence dispatch from a hook.

```go
uep.Event.NewEvent(extensiblePlugin.CLEPreEvent, func(ev *core.Events, i *core.EventInput) core.ErrorInterface {
	data, _ := core.ScopeGet[extensiblePlugin.SCLEData](ev.Scope(), extensiblePlugin.CLEScopeData)
	// shrink the node list to stop the loop early
	data.PLen = 0
	ev.Scope()[extensiblePlugin.CLEScopeData] = data
	return nil
})
```

### Example

A hook that rejects a command at runtime:

```go
ext := extensiblePlugin.New(engine)
engine.Plugins.AddPlugin(ext)

uep, _ := engine.GetUEP()
uep.Event.NewEvent(extensiblePlugin.CLEInPreEvent, func(ev *core.Events, i *core.EventInput) core.ErrorInterface {
	data, err := core.ScopeGet[extensiblePlugin.SCLEData](ev.Scope(), extensiblePlugin.CLEScopeData)
	if err != nil {
		return nil
	}
	idx := data.Idx
	if idx != nil && *idx >= 0 && *idx < len(data.Parsed) {
		// data.Parsed[*idx].Switch is the command about to run
	}
	return nil
})
```

## profiler

Counts and times every command. It registers handlers on the two hook events and
does nothing else, which is why it requires `extensiblePlugin` to be installed
first.

```go
err := engine.Plugins.AddPlugin(profiler.New())

report, _ := engine.Plugins.CallPluginMethod("profiler", "report")
```

Order matters: `profiler.Init` calls `IsPluginInstalled(extensiblePlugin.Name)`
and returns an error if the extension plugin is missing, so it must be added
second.

### Methods

`Call` dispatches on the method name:

| Method | Effect |
|---|---|
| `report` | Returns the report string |
| `reset` | Clears metrics and restarts the total timer |
| `enable` | Resumes collecting |
| `disable` | Stops collecting, keeps existing metrics |

`Run` is a no-op; the profiler is used through `CallPluginMethod`.

### Report

```text
Profiler report (0.00 sec total):

  String commands:
  String calls: 12 (total time: 32.848µs)
    keyval: count=9, total=26.352µs, avg=2.928µs, min=1.163µs, max=9.907µs
    section: count=2, total=3.652µs, avg=1.826µs, min=1.456µs, max=2.196µs
```

Byte metrics are grouped by opcode, string metrics by command name; both are
sorted for stable output. A section appears only for commands that were actually
called.

The total time is measured from `New()`, not from the first command, so it
includes the time before the engine ran.

### Overhead

Every command now passes through two extra event calls and a `time.Now()` on
each side. That is acceptable for development and for measuring relative
weights, but it is not free, and timings on a hot bytecode loop will be
dominated by the instrumentation rather than by the handler. Measure with the
profiler attached, then confirm with the benchmarks in
`example/tests/speedtest`.
