# tooling

Independent helpers around the engine. Each package is usable on its own.

| Package | Purpose | README |
|---|---|---|
| `plugin` | Plugin manager, `PluginInterface` and the event-based `Plugin` | [README](plugin/README.md) |
| `debugging` | `extensiblePlugin` and `profiler` | [README](debugging/README.md) |
| `bytecode` | int/float to bytecode and back, instruction shifter | |
| `astools` | Helpers for the `Parser3` AST | |

## bytecode

`bytecode.GenerationConfig` describes a binary instruction layout, and
`bytecode.Utils` converts values:

```go
u := &bytecode.Utils{}
b := u.IntToBytes(1234, 4, public.LittleEndian)
n := u.BytesToInt(b, public.LittleEndian)
```

Floating point has the same shape, including a range-mapping variant
(`Float64ToBytesRange`) that scales a value into a fixed-width field. This is the
package to reach for when producing input for `byteParsing.Parser1`, since both
sides must agree on field lengths and endianness.

`Shift` moves the instruction pointer through a byte slice, which is what the
byte parser uses to walk instructions.

## astools

Helpers for walking the tree produced by `Parser3`:

| Function | Returns |
|---|---|
| `GetChildren` | Child nodes of a node |
| `FindChild` | First child matching a name |
| `FindChildren` | All children matching a name |
| `FindChildIndex` | Index of the first matching child |
| `GetChildAt` | Child at a position |
| `GetTokenValue` | Token value from node metadata |
| `Walk` | Depth-first traversal |
| `WalkWithPath` | Traversal that also passes the path from the root |

`WalkWithPath` is what `AstEngine` uses for `BreakIf` and `SkipIf`: the path
identifies where in the tree the current node sits, and the command context
decides whether to stop or skip.

Every function tolerates a nil node and returns a zero value rather than
panicking, which keeps handler code free of nil checks when a parser may produce
partial trees.
