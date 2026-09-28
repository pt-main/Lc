# Examples

All examples run with `go run ./example/<path>`.

## `langs/`

Complete languages running on the engine.

- `langs/calculator` - float expressions. Lexer, `Parser3` grammar through an
  engine command, separate reporting for parse and evaluation errors.
- `langs/math` - integer language with variables, assignment and scope-backed
  evaluation.
- `langs/listLang` - the largest example: a full interpreter with functions,
  closures, recursion, lists, `if` / `while` / `for`, `break` / `continue` /
  `return` and higher order builtins. Shows a Pratt grammar, the engine
  generator as output and the profiler plugin.
- `langs/configLang` - ini-like config to JSON translator with the profiler
  plugin attached.

```bash
$ go run ./example/langs/calculator '(2+3)**4'
Lc version - 2.0.0
Result: 625

$ go run ./example/langs/math 'a := 5
b := a * 2
a + b'
15

$ go run ./example/langs/listLang 'fn fact(n) { if n <= 1 { return 1 } return n * fact(n - 1) }
print(fact(10))'
Lc version - 2.0.0
3628800
Profiler report (0.01 sec total):
  ...

$ go run ./example/langs/configLang
Lc version - 2.0.0
Profiler report (0.00 sec total):
  ...
{
  "database": { ... },
  "global": { ... },
  "server": { ... }
}
```

## `readme/`

The two minimal examples from the root README.

- `readme/string` - minimal string engine.
- `readme/byte` - minimal byte engine.

## `tests/`

- `tests/parser3Test` - the parser on its own: AST as text and JSON, Pratt
  precedence, error formatting. Does not use the engine.
- `tests/speedtest` - byte engine benchmarks.
  - `speedtest/byte/bench` - `go test -bench` benchmark.
  - `speedtest/byte/tests` - standalone timing runs.

```bash
$ go run ./example/tests/parser3Test

$ go test -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ ./example/tests/speedtest/byte/bench
```

## `packages/`

Core components in isolation, for when the behaviour of a single building block
is what matters.

- `packages/engine/core/events` - event basics, core events, `EventsTools`, scope
  sharing.
- `packages/engine/core/generator` - string and byte generation, pipeline order
  changed at runtime.
- `packages/engine/core/logger` - level filtering, custom status formats, log
  trimming.
- `packages/engine/core/other` - `Scope` read and write.
- `packages/engine/engines/string`, `packages/engine/engines/byte` - placeholders,
  they do nothing yet.
