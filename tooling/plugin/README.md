# tooling/plugin

The plugin system: an interface, a base implementation, a manager and a small
helper set for plugins that need to know about each other.

| File | Contents |
|---|---|
| `interface.go` | `PluginInterface` |
| `realization.go` | `Plugin`, the event-based base implementation |
| `manager.go` | `PluginManager`, holds plugins and their shared scope |
| `tools.go` | `Tools`, flags and plugin-presence checks |

## PluginInterface

```go
type PluginInterface interface {
	Name() string
	Init(scope core.ScopeType, pm *PluginManager) error
	Close() error
	Call(string, ...core.Option) (any, error)
	Run(input any) (any, error)
}
```

`Name` must be constant and unique. `Init` receives the manager scope and the
manager itself, which is how a plugin reaches the engine: the engine puts a
pointer to itself under `public.PluginsScopeEuPtr` in that scope, and a pointer
to the manager under `public.EuScopePmPtr`.

Implementing the interface directly is possible. In practice `Plugin` covers the
common case, because the four methods above map onto a standard lifecycle.

## Plugin

A plugin is an event container with three named events and two result keys.

```go
p := plugin.NewPlugin(
	"my_plugin",
	"init_event",
	"main_event",
	"close_event",
	"run_result",
	"call_result",
	context.Background(),
)
```

| Event | Called by | Purpose |
|---|---|---|
| `InitEvent` | `AddPlugin` | One-time setup, receives the manager |
| `MainEvent` | `Run` | Main work, receives the caller's input |
| `CloseEvent` | `Close` | Cleanup |
| any other | `Call(name)` | Named entry points |

`Run` and `Call` return whatever the handlers stored in the result scope key.
Both read `ScopeRunResultKey`:

```go
p.Events.NewEvent("main_event", func(ev *core.Events, i *core.EventInput) core.ErrorInterface {
	ev.Scope()[p.ScopeRunResultKey] = compute(i.Input)
	return nil
})

res, err := p.Run("input")
```

Note that `ScopeCallResultKey` is stored on the struct but not read: `Call`
returns the value under `ScopeRunResultKey` as well. A plugin that returns
different values from `Run` and `Call` must write both into the run key, or read
it back itself.

`Init` is called with the manager as `EventInput.Input`, so a handler can reach
it from there as well as from the scope. The base implementation does not isolate
anything: a plugin can read and write the engine scope directly.

The event names are stored on the struct (`InitEvent`, `MainEvent`,
`CloseEvent`, `ScopeRunResultKey`, `ScopeCallResultKey`) and are public, so they
can be read or reassigned after construction.

## PluginManager

```go
pm := plugin.NewPluginManager(scope)
```

| Method | Behaviour |
|---|---|
| `AddPlugin` | Registers by `Name()`, rejects duplicates, then calls `Init` |
| `DeletePlugin` | Calls `Close`, then removes. Missing name is not an error |
| `GetPlugin` | Lookup by name |
| `RunPlugin` | `GetPlugin` + `Run` |
| `CallPluginMethod` | `GetPlugin` + `Call` |
| `End` | Closes and removes every plugin |

`AddPlugin` registers the plugin before calling `Init`, so a plugin that fails
`Init` is already present in the map. Removing it is the caller's job.

The manager is reachable from the engine at `engine.Plugins`, and
`engine.End()` calls `PluginManager.End()`, which closes every plugin in turn.
If one `Close` fails, the remaining plugins are not closed.

`AddPlugin` is what makes ordering matter. `profiler.Init` returns an error
unless `extensiblePlugin` is already registered, so the extension plugin has to
go first.

## Tools

```go
t := &plugin.Tools{Pm: pm}
```

| Method | Use |
|---|---|
| `HasFlag` / `SetFlag` | Private string flags on the manager, for plugin coordination |
| `IsPluginInstalled` | Check whether another plugin is present, by name |

`IsPluginInstalled` is how `profiler` verifies its dependency.

## Registering on an engine

At build time, with `WithPlugins`:

```go
engine, err := lc.NewEngineBuilder(public.StringEngineType, public.StringResType).
	WithStringParser(parser).
	WithPlugins(myPlugin).
	Build()
```

`Build` calls `Init` for each plugin. If any `Init` fails, `Build` returns an
error.

After build, which is required for plugins that need the engine pointer:

```go
err := engine.Plugins.AddPlugin(extensiblePlugin.New(engine))
```

`extensiblePlugin.New` takes the engine, so it cannot be passed to `WithPlugins`
- the engine does not exist yet at that point.
