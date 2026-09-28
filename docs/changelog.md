# Changelog

All notable changes to **Lc** (a Go framework for building language runtimes).

Format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
versioning follows [SemVer](https://semver.org/).

[Russian version](changelog-ru.md)

## [2.0.0] - 2026-09-28

A release about consistency, correctness and documentation. The architecture did
not change: all package paths and `go.mod` stayed the same, no renames. The bulk
of the work is a full pass over the public API, thread safety, a rewrite of
`parser3` and the lexer, and the migration of documentation to bilingual READMEs.

### Breaking

- Global typo fixes across the public API: `Endianess` -> `Endianness`
  (`EngineBuilder.WithEndianness`, `public.ByteEngineScopeEndianness`,
  `bytecode.GenerationConfig.Endianness`), `Calloop` -> `CallLoop`
  (`public.StringCallCallLoopEvent`, `public.AstCallCallLoopEvent`),
  `Loger` -> `Logger` (`core.LoggerInterface`), `Instaled` -> `Installed`
  (`plugin.Tools.IsPluginInstalled`).
- String values of constants changed: `ByteEngineScopeEndianness`
  (`"ENDIANNESS int"`), `ByteEngineScopeHotloopCtxCheckPeriod`
  (`"CTX_CHECK_PERIOD int"`), call loop event names (`"STRINGCALLLOOP"`,
  `"AST:CALLLOOP"`), default events error codes (`"SYSTEM@DEFAULT_EVENTS"`,
  `"DEFAULT_EVENTS:..."`), `extensiblePlugin` hooks (`"CallLoopE ...`).
- The `errors.DefaultEventsCallErrorContex` error code was removed; its usage
  moved to `DefaultEventsCallErrorContexted`.
- `engine.EngineInterface` gains the method `GetCommand(CmdT)`, so third-party
  implementations no longer satisfy the contract.
- `parser3.ParseError.Code` and `parser3.GrammarError.Code` were renamed to
  `Phase`; creating such errors via a literal with `Code:` no longer compiles.
  `ParseError` also gains a `Found []string` field listing the token types at
  the point of failure.
- `parser3.GrammarError.GetCode()` now always returns the `GrammarErrCode`
  constant instead of the phase code.
- `parser3.Adapter.Parse` no longer wraps the parse error in
  `*AdapterError`, it returns the original parser error. `parser3.AsParseError`
  and `parser3.AsGrammarError` were added for type checks.
- The output format of `parser3.FormatError(err, false)` changed from the
  single-line `"a -> b -> c"` to a multi-line form with
  `"  caused by: "` continuations.
- `plugin.NewPlugin` was fixed: `ScopeRunResultKey` no longer receives the
  value of `scopeCallResultKey`, and `Plugin.Call` reads `ScopeCallResultKey`.
  Plugins that wrote the result to the run key must write to the call key.
- `EngineUniversal.End()` is now idempotent (a repeated call returns
  `CorePackageLcLifecycleError`), always cancels the context when
  `CtxCancelCause` is set, releases scope guards, and returns the plugin error
  directly instead of through string formatting.
- `byteParsing.Parser1.Parse` returns `nil` instead of a partially parsed slice
  on error, and rejects a negative `argscount`, which previously arose from sign
  extension inside `BytesToInt`.
- `stringParsing.Lexer.Parse` no longer panics on a `nil` or empty
  `*parsing.ParseOption` - previously `opts[0].UEP.Logger` was dereferenced
  without a check.
- `EngineBuilder.WithColors()` and the `colorEnable` argument were removed
  (leftovers of the `tap` dependency dropped in v1.6.0).

### Fixed

- **Data races eliminated:** event maps (`GetEvents`, `CoreEvents`,
  `SetEvents`, `ReplaceEvent`, `SetProperty`), `GetCommands` and
  `GetCommandCtx` of all three engines, plus the profiler `enabled` flag are
  now mutex-protected or return copies. `ByteCallEvent` no longer reads
  `e.Commands` and `e.AutoBytecodeIndexShift` without locking.
- `NewCommandByte`: the auto index shift is now applied only when
  `autoBytecodeIdxShift == true` (the flag used to be set unconditionally, so
  the `false` value was silently ignored), and the auto-operation counter with
  an explicit opcode grows as `max(opcode+1, counter)`, so an auto-operation
  can no longer emit an already taken opcode.
- `EngineBuilder.Build` publishes the plugin manager and the engine scope
  **before** calling `Init` on plugins (otherwise `Init` saw a nil
  `eu.Plugins`), and on an `AddPlugin` failure it calls `pm.End()` and cancels
  the context instead of leaking half-initialised plugins.
- `End()` no longer panics on a repeated call and correctly releases scope
  guards before resetting the engines.
- `ByteCallHotLoopEvent` checks the context on the first iteration, so short
  programs are no longer run without reacting to cancellation.
- `AstMakeCommandCtx` contained a copy-paste bug: with `nil` in `skipif` the
  `breakif` value was overwritten. Fixed.
- `WalkWithPath` added the parent's name instead of the child's and wrote paths
  into a shared buffer, so neighbouring node paths overwrote each other.
- Child node pointers in `astools` now really point into the tree
  (`&children[i]`) instead of copies of the loop variable.
- `AstEngine.Process` read the parsed nodes from the empty scope key `""`
  instead of `public.StringEngineScopeParsed`.
- `bytecode`: encoding `float64` via `math.Round` instead of truncating
  `uint64(x * max)`, plus fixing the [-1,1] range encoding to
  `(value+1)/2*max`, make `BytesToFloat64(Float64ToBytes(x))` an exact inverse.
- `astools.GetTokenValue` read the key `"__value"`, which is what the lexer
  actually writes (it used to read the non-existent key `"value"`).
- `GetCommandCtx` creates a missing context (the condition was inverted in
  v1.6.0).
- `HasCommand` no longer returns the bogus error `"skip"`.
- `BreakIf` / `SkipIf` path matching moved from exact equality to prefix
  comparison, so `BreakIf = [][]string{{"block","inline"}}` now really prunes the
  subtree.
- `WorkIter` skips the root node so the root handler is not called twice.
- `core.GetErr` returns `nil` for a `nil` input and `nil` cause instead of
  panicking.

### Added

- **Thread-safe scope access:** `core.ScopeGetSynced`, `core.ScopeSetSynced`,
  `core.ScopeDeleteSynced`, `core.ReleaseScopeGuard` and the `scopeGuards`
  guard registry storing the map itself, so the address cannot be reused after
  garbage collection. `ScopeType` intentionally stays a plain map, so
  `scope[k] = v` still compiles.
- `GetCommand` - a point lookup of a command without copying the whole command
  map - in `StringEngine`, `ByteEngine` and `AstEngine`, plus
  `ByteEngine.GetAutoBytecodeIndexShift() map[int]bool`.
- `engine/events.AstParsingEvent` - an analogue of `StringParsingEvent` for
  `engine.AstEngine`; `NewAstEngine` with `addDefaultEvents=true` now really
  registers it on `public.StringParseEvent` (in v1.6.0 the argument was accepted
  and ignored).
- `parser3`: phase constants `PhaseLexer`, `PhaseStart`, `PhaseEnd`,
  `PhaseExpect`, `PhasePeek`, `PhaseEOF` and meta keys `MetaChildren`,
  `MetaOperator`, `MetaLeft`, `MetaRight`, `MetaOperand`; functions
  `AsParseError`, `AsGrammarError`, `ChildrenOf`.
- **`parser3` speed and robustness:** rule nesting limit `maxRuleDepth = 256`,
  left-recursion detection, memoization of `NamedExpr` results by
  (rule, position) pair, tracking of the furthest error for backtracking,
  `fatal(err)` so grammar errors are not swallowed, and
  `maxPrattDepth = 512` for `PrattExpr`.
- **The lexer moved from quadratic to linear complexity:**
  `regexp2.FindRunesMatchStartingAt` over a shared rune slice is used instead of
  `FindStringMatch` over a fresh copy of the tail at every position;
  zero-length and unanchored matches are rejected so a rule cannot loop.
- `core.Events.GetEvents` and `CoreEvents` return copies; `SetEvents` stores a
  copy; private `Events.isDebug()`.
- 14 new test files (~1960 lines), `go test ./...` fully green:
  `end_test.go`, `astEngine_test.go`, `engine/astEngine_test.go`,
  `engine/core/errors_events_test.go`, `parsing/stringParsing/parser3/parser3_test.go`
  (45 tests), `tooling/astools/main_test.go`, `tooling/bytecode/utils_test.go`,
  `tooling/plugin/realization_test.go`, `parsing/stringParsing/lexer_perf_test.go`,
  `parsing/stringParsing/bracket_test.go`, `parsing/stringParsing/parser1_option_test.go`.
- New examples: a full `example/langs/listLang` language interpreter
  (closures with lexical scoping, `if`/`else`, `while`, `for..in`,
  `break`/`continue`/`return`, built-in `print`, `len`, `str`, `int`, `map`,
  `filter`, `reduce`), `example/langs/configParser` with `ActionExpr` for
  validating values right during parsing, and also
  `example/packages/engine/engines/ast` and
  `example/packages/tooling/plugin`.
- A `.gitignore` file (excludes `AGENTS.md`, `.gocache/`, `.gocache2/`).
- Support for unsigned conversions in `bytecode.Utils`: `unsignedMaxValue`,
  `bytesToUint`, `uintToBytes` for correct handling of signed sizes.

### Changed

- All packages are now covered by ordinary English Go doc comments;
  `GODOC`-style markers, em dash, en dash, non-breaking spaces, typographic
  quotes and ellipses were removed from the code.
- `engine/core/events.go`, `errors.go`, `generator.go`,
  `universalEngineParams.go`, `engine/astEngine.go`, `engine/byteEngine.go`,
  `engine/stringEngine.go`, `engine/events/*` were reworked: names unified
  (`event` -> `eventHandler`, `val` -> `value`, `snake_case` parameters ->
  `camelCase`), shadowed variables eliminated.
- `tooling/bytecode/utils.go` fully rewritten (216 lines changed):
  duplication removed, tests added for int/float64/big-endian conversions.
- `parsing/stringParsing/parser3`: `grammar.go` (+628), `parser.go` (+235),
  `errors.go` (+125), `formatter.go` (+173), new `helpers.go`.
- `EngineUniversal.ProcessString` / `ProcessBytes` no longer replace the
  configured engine context with `context.Background()`.
- `astools.Walk` rewritten from recursion to an explicit iterative DFS with no
  depth limit.

### Docs

- `README.md` fully rewritten (737 lines): sections Scope, Architecture,
  Extension points, Engines, Core concepts, Plugins, Debugging and tooling,
  Context and cancellation, Performance, Examples. It explains that
  `AstEngine` is not a third backend but implements `StringEngineInterface`,
  and describes a practical limitation: the engine constructors are typed to
  `ParsedNode` / `ParsedBytes`.
- A full Russian translation `README-ru.md` was added (791 lines).
- 15 generated gomarkdoc `GODOC.md` files were deleted from packages and 12
  from examples; in their place bilingual READMEs were added for
  `engine/core/`, `parsing/stringParsing/parser3/`, `tooling/plugin/`,
  `tooling/debugging/`, and `parsing/README.md`,
  `parsing/stringParsing/README.md`, `tooling/README.md`, `example/README.md`
  were rewritten. `tooling/debugging/README.md` honestly describes the
  profiler's overhead.
- The `backup/lexer.five` and `scan_results.txt` artifacts were removed.
- Examples ported to the new API: `example/langs/calculator` and
  `example/langs/math` now work through
  `lc.NewEngineBuilder(...).WithStringParser(&parser3.Adapter{...})` and report
  errors via `parser3.AsParseError`.

## [1.6.0] - 2026-08-29

> **Warning:** the v1.6.0 tag does not compile. The `colorEnable` parameter was
> removed from `NewStringEngine` / `NewByteEngine`, but `builder.go` in this
> commit still passes it. The error is fixed in the following commit
> `2481e0d "Fixes"`, which is part of v2.0.0. Use v2.0.0.

### Added

- **AST engine** (`engine/astEngine.go`, 222 lines): `engine.AstEngine` with
  the fields `AstCommandCtx`, `UEP`, `Parser`, `Commands`, `CanBeUnknown`,
  `CanMainNodeBeUnknown`; types `AstCommandCtx` (`Name`, `Path`, `Parent`,
  `CurrentChildren`, `BreakIf`, `SkipIf`) and `AstCommandCtxPath`
  (`Path []string`, `Nodes []*ParsedNode`); constructors
  `MakeAstCommandCtxPath`, `AstMakeCommandCtx`; methods `Process`, `Work`,
  `WorkIter`, `HasCommand`, `GetCommandCtx`, `NewCommandFull`, `NewCommand`,
  `GetCommands`, `GetUep`, `GetParser`.
- `engine.AstEngineInterface` - an alias for `StringEngineInterface` (not a
  new type).
- Constructor
  `lc.NewAstEngine(resType, pipeline, addDefaultEvents, parser, context, canNodeBeUnknown, canMainNodeBeUnknown)`.
- Error codes `AstEngineProcessError1/2`, `AstEngineHandlerError`,
  `AstEngineUnknown`, `EngineLifecycleEnd`.
- Events `public.AstCallEvent`, `public.AstCallCallLoopEvent`,
  `public.AstCommandCallEvent`.
- `core.LoggerInterface` - a logger interface making the logger swappable.
- `Logger.SetStatusForm(status, form)` - programmatic status format setting.
- `astools.WalkWithPath(node, fn)` - tree traversal passing the path to a node.

### Breaking

- The `colorEnable bool` parameter was **removed** from `NewStringEngine` and
  `NewByteEngine`.
- The `github.com/pt-main/tap` v1.4.7 dependency was removed from `go.mod`;
  only `dlclark/regexp2 v1.12.0` remains. All colour handling was removed
  from `main.go`, `engine/core/logger.go` and `parser3/formatter.go`.
- Event name strings became prefixed with the engine name:
  `"INPUT string->..."` -> `"STRING:INPUT ..."`, `"BYTE:INPUT ..."`,
  `"STRINGCALLOP ..."` and others.
- `UniversalEngineParams.Logger` changes type from `*Logger` to the
  `LoggerInterface` interface.
- `parser3.FormatError(err, useColors bool)` ignores the `useColors` argument
  and always returns the chain without colour.

### Changed

- `StringEngine.Process` and `AstEngine.Process` use
  `core.GetRealErrorReverse(err)` when wrapping parse and call errors so
  messages read from the root cause.
- The default logger format changed from
  `"[?BE]%s[?RT] [?CN][%v][?RT] [?GN][%s][?RT]\n"` to `"%s [%v] [%s]\n"`.
- `astools.Walk` rewritten from recursion to an iterative DFS in the same
  preorder, but with no depth limit.
- `plugin.NewPlugin` takes a `context context.Context` parameter.
- `CallLoopData.Engine` and `CLEData.E` change from `*E` to `E`.

### Fixed

- `ErrExit` handling moved into `ByteCallEventIteration` (`*idx = -1`), which
  fixes exiting the byte engine's hot loop.
- The lock order in `Events.NewEventBefore` was fixed; the method now returns
  an error if the event is missing.
- `StringCallEvent` uses `core.ScopeGet` instead of an unchecked type assertion
  from scope.
- The `backup/lexer.five` artifact was removed.

## [1.5.8] - 2026-08-28

### Added

- New flag `public.ByteEngineScopeHotloopCtxCheckPeriod` for configuring how
  often the context is checked in the byte engine's hot loop.
- `CallLoopData.Other` - a new field for passing extra data to hooks.
- New example `example/langs/math` - a calculator on parser3 + PrattExpr.

### Fixed

- `runtime.GC()` calls were removed from the hot loop (they were the cause of
  a performance drop from ~160m to ~120m ops/s).
- The auto index shift (`*idx++`) now happens **before** the handler call, and
  `*idx` is rolled back on error - the pointer no longer skips a faulty opcode.
- Fixed a wrong scope key used to read the context check period (`BytecodeIdx`
  was read instead of the new key).
- Fixed double event registration in the speedtest benchmark.

### Changed

- The context check in the hot loop runs once every N iterations (255 by
  default) instead of every iteration.
- The loop exit condition uses `uint` - a trick to handle `idx == -1` without a
  double check.

## [1.5.7] - 2026-08-25

### Breaking

- `EngineUniversal.ProcessString`, `ProcessStringWithCtx`, `ProcessBytes`,
  `ProcessBytesWithCtx`, `CheckEnded` and all `Process` methods of the engines
  return `core.ErrorInterface` instead of `error`.
- New error codes `public/errors.CorePackageLcError` (`"SYSTEM@LC"`) and
  `public/errors.CorePackageLcLifecycleError` (`"SYSTEM@LC:LIFECYCLE"`).

### Added

- `core.GetRealErrorReverse` - formats the error chain from root to wrapper.
- `core.ErrorInterface` extended with `GetCode()`, `GetMsg()`, `GetMeta()`,
  `Unwrap()` for proper `errors.Is` / `errors.As` support.
- parser3 error types (`ParseError`, `GrammarError`, `AdapterError`) were
  brought to the same contract.
- Exported parser3 error codes: `ParseErrCode`, `GrammarErrCode`,
  `AdapterErrCode`.

### Fixed

- **HOTFIX (v1.5.7-f):** in `ByteEngine.Process` the second error handler used
  `err1` instead of `err2` - byte call errors showed the parse error instead of
  the real cause.

## [1.5.6] - 2026-08-24

### Added

- `InstructionsGenerator.Generate` panics on a zero-length argument, with an
  explicit warning in the documentation.

### Fixed

- Restored the zero-length argument check in `byteParsing.Parser1.Parse` (it was
  removed in v1.5.5, now returns a structured `"Zero argument length"` error
  with `EMK(0,"int")` metadata).
- `byteParsing.Parser1.Parse` saves and restores `p.Config.Shifter.Idx` before
  and after the loop - the parser can be reused and used concurrently.

## [1.5.5] - 2026-08-22

### Fixed

- **Critical:** `byteParsing.Parser1.Parse` modified the shared
  `p.Config.Shifter.Idx` pointer. The value and the pointer itself are now saved
  before the loop and restored in `defer`, so `Parser1` can be reused and used
  concurrently without clobbering the caller's index.
- The logging check in `Parser1.Parse` is protected against a `nil`
  `*ParseOption` pointer.

### Changed

- `parser3.Adapter.Parse` returns `core.ErrorInterface`, completing conformance
  with `ParserInterface`.

## [1.5.4] - 2026-08-19

### Added

- `core.GetErr` - unwraps one level of wrapper; if the cause is not an
  `ErrorInterface`, an `Error{Code: WrappedError}` is synthesised.
- Error code `public/errors.WrappedError`.

### Changed

- `Error.Format` no longer appends the `"----+"` separator.

## [1.5.3] - 2026-08-18

### Added

- `core.ErrorInterface` extended with `GetCode()`, `GetMsg()`, `GetMeta()` and
  `Unwrap()`, allowing `errors.Is` / `errors.As` across the whole lc error
  chain.
- parser3 error types (`ParseError`, `GrammarError`, `AdapterError`) were
  brought to the same contract.
- Exported parser3 error codes: `ParseErrCode`, `GrammarErrCode`,
  `AdapterErrCode`.
- `parser3.Parser.Parse` and `Expect` return `core.ErrorInterface`, completing
  conformance with `ParserInterface`.

## [1.5.2] - 2026-08-18

A major release for structured errors and documentation.

### Added

- Package `public/errors` systematising error codes by domain: `engines.go`,
  `events.go`, `generator.go`, `main.go`, `others.go` (about 25
  `ErrorCodeType` constants).
- Type `ErrorCodeType` and `ErrorMetaType`, function `EMK(n, valType)` for
  generating metadata keys.
- File `engine/core/errors.go`: `core.Error{Code, Msg, Meta, Cause}` with
  methods `Error()`, `Format()`, `WithMeta`; constructors `core.Err`,
  `core.Wrap`; helpers `GetMetaValue[T]`, `GetRealError`, `core.ErrExit`.
- parser3 error types: `ParseError`, `GrammarError`, `AdapterError` with the
  fields `TokenIdx`, `TokenPos`, `Code`, `Expected`, `Got`, `Raw`, plus
  `Unwrap` and `Format` methods.
- Error formatters `parser3.FormatError(err, useColors)` and
  `FormatErrorPretty`; hardcoded ANSI sequences were removed from the code.
- `core.Events.SetProperty(name, value)` with support for the `"debug"`
  property: with debug off, the start and end events of each call are not
  created, removing the overhead.
- 27 gomarkdoc-generated `GODOC.md` files for all public packages.

### Breaking

- Almost all public APIs switch from `error` to `core.ErrorInterface`:
  `core.ScopeGet`, all `Generator` methods, `NewUniversalEngineParams`, all
  `EventsInterface` methods, `parsing.ParserInterface.Parse`, all
  `bytecode.Shift` methods and all `DefaultEvents` handlers.
- `core.EventType` and `core.CommandType` return `ErrorInterface`.
- `public.ErrExit` was removed, replaced by `core.ErrExit` and the code
  `public/errors.ErrExit`.
- `engine/core/config.go` was renamed to `engine/core/types.go`.

### Fixed

- **Critical:** inverted `canBeUnknown` check in
  `StringCallEventIteration` - registered commands were skipped and
  unregistered ones were called. Fixed: unknown commands are now skipped only
  when `StringEngineScopeCanBeUnknown = true`, otherwise
  `DefaultEventsCallErrorUnknown` is returned.
- `EventsTools.ChangeCoreEvent` with `idx == 0` replaced the whole handler list
  instead of a single entry; `idx < 0` now returns an error.
- `parser3.engineAdapter` accepts `children` of any named slice type via
  `reflect`, not only the exact `[]ParsedNode`.
- `byteParsing.Parser1` returns structured parse errors with position, command
  and argument number metadata.
- `Generator.GetBytesRes` / `GetStringArrRes` attach `EMK(0, "string")`
  metadata with the missing pipeline point.
- Improved lexer error messages: line and column position, context snippet,
  correct handling of unclosed and extra brackets.
- Every lexer token now always carries `__pos`, `__start` and `__end` in its
  metadata.

### Performance

- Lexer: precomputed `ruleGroups` indices and `openByByte` / `closeByByte`
  buckets for brackets eliminate linear map traversals at every position.

## [1.5.0-pre2] - 2026-08-08

### Added

- `astools.FindChildIndex(node, switchName) int` - the index of the first child
  with the given name, or -1.

### Fixed

- **Critical:** `parser3.ChoiceExpr.Parse` stopped trying alternatives after
  the first failure. The position is now reset and the search continues.
- `parser3.Parser.Parse` passes `opts ...*parsing.ParseOption` to the lexer
  (the options used to be lost).
- `parser3.Adapter.Parse` correctly forwards options and produces coloured
  errors.

### Changed

- parser3 diagnostics reworked: every grammar expression returns a detailed
  message with the token position, the expected and the received value, the
  number of alternatives and a list of unselected tokens.
- README gained blocks with the output of `go run ./example/readme/string`
  and `.../byte`.

## [1.5.0-pre] - 2026-08-04

A pre-release of 1.5.0, the largest architectural release before v2.0.0.

### Added

- `engine.EngineInterface` - a common interface for both engine types (string
  and byte). Typed aliases: `StringEngineInterface`, `ByteEngineInterface`.
- `tooling/astools` - AST traversal utilities: `GetChildren`, `FindChild`,
  `FindChildren`, `GetTokenValue`, `GetChildAt`, `Walk`.
- `tooling/debugging/profiler` - a performance profiler with the metrics
  `Count`, `TotalTime`, `MinTime`, `MaxTime`, `Avg()` and the methods
  `Report`, `Reset`, `Enable`, `Disable`.
- `extensiblePlugin.New` and a full `Close()` to restore the original handlers;
  the plugin registers the `ECLFlag` flag.
- `core.Events.ReplaceEvent(name)` - full event replacement.
- New scope entry `public.StringEngineScopeCanBeUnknown` - controls the
  behaviour for unknown commands.
- Examples `example/langs/calculator`, `example/langs/configLang`,
  `example/readme/string`, `example/readme/byte` and the
  `example/packages/**` set.
- Godoc comments for `Lexer`, `Parser1`, `Parser2`, `PluginManager`,
  `plugin.Plugin`, `ProfilerPlugin`, `ExtensibleCLPlugin`.

### Breaking

- **The main change:** all command handlers now take
  `engine.StringEngineInterface` / `engine.ByteEngineInterface` instead of
  `*engine.StringEngine` / `*engine.ByteEngine`. This changes the signature of
  every handler: `func(e *StringEngine, n *ParsedNode) error` ->
  `func(e StringEngineInterface, n *ParsedNode) error`.
- `EngineUniversal.StringEngine` and `.ByteEngine` become interface-typed
  fields; the accessors `GetUep()`, `GetParser()`, `GetCommands()` were added.
- `core.EventInput` was renamed to `core.SimpleInput` (the alias is kept).
- `NewCommand` takes a `*core.SimpleInput` and returns `error`; the previous
  signature is kept in `NewCommandFull`.
- The flag `engine.AutoshiftNewCommandFlag = "autoShift"` controls the auto
  index shift in the byte engine.
- `plugin.Tools.Pm` changed from a value to a pointer; `PluginManager.Flags`
  became the private field `flags`, accessible only through `Tools`.
- `CallLoopData.Engine` and `CLEData.E` change from `*E` to `E`.

### Fixed

- `ErrExit` handling moved into `ByteCallEventIteration` (`*idx = -1`), which
  fixes exiting the byte engine's hot loop.
- The lock order in `Events.NewEventBefore` was fixed; the method returns an
  error if the event is not found.
- `StringCallEvent` uses `core.ScopeGet` instead of an unchecked type assertion
  from scope.
- Unknown commands in the string engine now produce an error instead of being
  silently skipped.

## [1.4.6] - 2026-07-30

### Fixed

- `ErrExit` handling moved from `StringCallLoopEvent` to
  `StringCallEventIteration` - setting `*idx = -1` for immediate loop
  termination.

### Changed

- Multibyte bracket support in the lexer: `Brackets` changed from `[]string`
  to `[][2]string` (explicit open/close pairs).
- Internal bracket maps changed from `map[rune]rune` to `map[string]string`.

## [1.4.5] - 2026-07-29

The version is marked as 1.4.5 in the repository; the `v1.4.4` tag points to
the previous commit.

### Added

- `public.ErrExit` - handlers can return it for a controlled exit from the call
  loop; the loops recognise it via `errors.Is` and terminate cleanly, without
  tracing the error.
- The instruction index is published to scope via
  `public.StringEngineScopeInstrIdx`.
- `public.ByteEngineScopeHotloopCtxCheckPeriod` - cancellation check
  frequency setting.

### Fixed

- **Critical:** the context check in the byte engine's hot loop sat inside the
  `iter&4095 == 0` condition, so for short programs cancellation was never
  checked at all. `ctx.Err()` is now checked at the start of every iteration.
- `StringCallLoopEvent`: the exit condition
  `for *idx < pLen && *idx >= 0` was simplified to `for *idx < pLen` with an
  explicit `*idx < 0` check inside, since a negative index is used as the exit
  signal.

## [1.4.3] - 2026-07-28

### Breaking

- `parser3.Adapter.ParseFlat` was renamed to `parser3.Adapter.Parse` to match
  `parsing.ParserInterface`.

### Fixed

- `StringCallEvent`: the panic message used the wrong specifier `%e` instead of
  `%v`, so the panic value was printed incorrectly.
- `StringCallEvent`: the log said "end parsing event" instead of "end call
  event".
- Fixed the context marker in the error output: `[?BBK]    |` -> `[?BBK]>    |`.

## [1.4.2] - 2026-07-21

### Added

- `EngineUniversal.End()` - explicit end of the engine lifecycle.
- The `EngineUniversal.CtxCancelCause context.CancelCauseFunc` field and the
  `WithContext(ctx)` builder method, giving proper cancellation with a cause.
- `EngineUniversal.CheckEnded()` - state check before every operation; called
  in `ProcessString`, `ProcessBytes`, `GetUEP`, `NewCommandByte`,
  `NewCommandString`.

### Changed

- **Breaking:** `GetUEP()` now returns
  `(*core.UniversalEngineParams, error)` instead of a plain pointer.
- `End()` cancels the context, handles panics and returns the plugin error
  directly instead of through string formatting.

## [1.4.1] - 2026-07-19

### Added

- Performance measurement infrastructure: `example/speedtest/tests/main.go`
  (the test set), `example/speedtest/byte_test.go` (the benchmark); speedtest
  was converted from a plain program to `testing.B`.

### Changed

- `builder.WithEndianess` takes `public.EndianType` instead of `int` (together
  with the `NewEngineBuilder` typing introduced in 1.3.2).

## [1.4.0] - 2026-07-17

### Added

- Package `tooling/debugging/extensiblePlugin` - a plugin extending the command
  processing loop: the hooks `CLEPreEvent`, `CLEInPreEvent`, `CLEInPostEvent`
  (called once around the loop) and traversal data `SCLEData` / `BCLEData` in
  scope.
- `tooling/plugin/tools.go` with the `Tools` type - access to the plugin
  manager flags.
- `core.EventsInterface` - an interface for the event system.
- `public/logging.go` with the logging status names: `LogEvents`, `LogParsing`,
  `LogVerbose`.
- `core.Events.NewEventBefore` to register a handler at the start of the list;
  `EventsTools` with `ChangeCoreEvent`, `GetCoreEvent`, `GetCoreEventIdx`.
- The `public.ByteCallHotloopEvent` event and `public.CLEScopeData`: a separate
  hot loop for the byte engine.

### Changed

- Events renamed: the strings were brought to the form
  `"INPUT string->PARSED []ParsedNode"`, `"call(PARSED []ParsedNode)"`,
  `"CALLOP call(PARSED []ParsedNode)"`, `"HOTLOOP call(PARSED []ParsedBytes)"`.
- The logging system was extended: `MaxLogLength` in the logger.

## [1.3.9] - 2026-07-12

### Added

- `example/speedtest/byte.go` - a byte engine measurement program.
- Type `ByteCallAttr` (`RawIdx`, `RawNode`, `Abis`, `Handler`) - precomputed
  command attributes for the hot loop.

### Changed

- `ByteCallEventIteration` was reworked: the signature
  `(idx *int, parsed ByteCallAttr, e *engine.ByteEngine) error` instead of
  parsing the node inside every iteration. Opcode parsing, handler lookup and
  the `AutoBytecodeIndexShift` read were moved out of the loop, and a ready
  `Handler` is passed to the handler.
- `ByteParsingEvent` passes a `*parsing.ParseOption` to the parser instead of
  calling `Parse(input)` without options.

## [1.3.7] - 2026-07-06

### Fixed

- `ByteCallEvent` no longer dereferences the `idx` pointer in `defer` before
  it is initialised: `last_cmd_switch` and `idx` were moved above the `defer`,
  so the deferred error handler is now safe.
- `ByteCallEventIteration` returns an error instead of `nil`, so the handler
  error is no longer lost.

## [1.3.6] - 2026-06-29

### Fixed

- **Critical:** in `StringCallEvent` a wrong `break` stopped the loop after the
  first command. Now all commands of the line are executed, and the loop is
  interrupted only on error.
- `recover` was moved from the loop body (where `defer` accumulated on every
  iteration) into a single `defer` for the whole event.
- `StringCallEvent` and `ByteCallEvent` print the error position (the source
  line and the command index) directly in the deferred handler, so a panic in a
  handler also yields a meaningful message.

## [1.3.5] - 2026-06-24

### Added

- `parsing/stringParsing/parser3/engineAdapter.go` - a parser3 adapter to
  `parsing.ParserInterface`.
- `tooling/plugin/realization.go` - the `PluginInterface` implementation.

### Changed

- `tooling/plugin/core.go` was renamed to `tooling/plugin/manager.go`.
- The lexer was reworked: -196 lines, bracket handling simplified.
- `Parser1`, `Parser2` and the parser3 grammar moved to `*parsing.ParseOption`.

## [1.3.3] - 2026-06-23

### Fixed

- `StringCallEvent`: `defer func() { recover() }` is no longer created on every
  loop iteration (memory growth and wasted work), `recover` was moved into a
  single deferred handler for the whole event.
- A handler error and a panic now interrupt the command loop with a clear
  message instead of silently continuing.

### Changed

- `byteParsing` uses the shared `parsing.ParserInterface` type.

## [1.3.2] - 2026-06-23

A major release for typed scope, enum constants and the plugin system.

### Added

- Package `public` - constants and types shared by the engines.
- `public/types.go`: `ResType` (`ByteResType`, `StringResType`), `EndianType`
  (`BigEndian`, `LittleEndian`), `EngineType` (`ByteEngineType`,
  `StringEngineType`).
- `public/scope.go`: all engine, event and plugin scope strings are collected
  into constants (`ByteEngineScopeParsed`, `StringEngineScopeParsed`,
  `EventsScopeCallName`, `PluginsScopeEuPtr` and others).
- `public/events.go` - event name constants.
- `engine/core/scope.go` with the generic helper `core.ScopeGet[T]` for
  type-safe scope reads.
- `tooling/plugin/interface.go` with `PluginInterface` and the full
  `Name/Init/Close/Call/Run` contract.

### Changed

- **Breaking:** `NewEngineBuilder(engineType int)` ->
  `NewEngineBuilder(engineType public.EngineType, resType public.ResType)`; the
  generator result type became an explicit builder parameter.
- **Breaking:** `WithEndianess(endianess int)` ->
  `WithEndianess(endianess public.EndianType)`.
- **Breaking:** `WithPluginManager(plugins ...*plugin.Plugin)` ->
  `WithPlugins(plugins ...plugin.PluginInterface)`.
- The default event handlers read the input and the parsed nodes via
  `core.ScopeGet` instead of a direct type assertion from scope.
- `ByteCallEventIteration` was extracted into a separate method - the bytecode
  iteration step was moved out.

## [1.2.0] - 2026-06-20

### Added

- **New package `parsing/stringParsing/parser3`** - a generative parser with a
  built-in AST: the types `Grammar`, `Rule`, `Expr` (interface), `TokenExpr`,
  `SequenceExpr`, `ChoiceExpr`, `RepeatExpr`, `OptionalExpr`, `NamedExpr`,
  `NodeExpr`; the constructor
  `NewParser(lexer, grammar, startRule, ignoreTypes)`; the methods `Expect`,
  `Peek`, `skipIgnored`, `Errorf`.
- `parsing/main.go`: the `ParseOption{UEP, Flags, Other}` type and the generic
  `ParserInterface[I, P] { Parse(I, ...*ParseOption) ([]P, error); String() string }`.
- `example/parser3.go` - a calculator that outputs the tree as JSON.
- `core.Events.NewEventBefore` - handler registration at the start of the
  event list.

### Changed

- **Breaking:** the `Parse(I, ...interface{})` signature was replaced with
  `Parse(I, ...*ParseOption)`.
- **Breaking:** the lexer became block-oriented: `LexerConfig.UseLineContinuation`,
  `SkipEmptyLines`, `TrimBlocksSpace` appeared, and tokens get `__block_index`
  metadata.
- **Breaking:** `tooling/plugin` was reworked around interfaces: `Plugin.Name`
  went from a field to a method, the `NewPlugin(name, initEvent, mainEvent, closeEvent)`
  constructor was added, the `InitPlugin` / `ClosePlugin` / `RunPlugin` methods
  were renamed to `Init` / `Close` / `Run`, and `PluginInterface` appeared.
- **Breaking:** `PluginManager.Plugins` is now `map[string]PluginInterface`;
  `DeletePlugin` and `GetPlugin` return an error if the plugin is not found.

### Fixed

- Parse logging now goes through `*parsing.ParseOption` (previously
  `core.UniversalEngineParams` was passed by value, so the logger inside the
  parser never fired).

### Removed

- The `example/configLang1` examples were removed, `example/configLang2`
  remains as a reference.

## [1.1.5] - 2026-06-19

### Added

- `stringParsing.LexerConfig` with `UseBracketBalance` and `Brackets []string`;
  `NewLexer(rules, config *LexerConfig)` takes a second argument.
- Bracket balancing in the lexer: the `isBracketBalanced` method, the result is
  written to the token metadata as `__bracket_balanced`.
- The `String() string` method in `parsing.ParserInterface`; the `Parser1`,
  `Parser2` and `byteParsing.Parser1` implementations return their paths.

### Changed

- `stringParsing.ParsedNode.Metadata` was moved from
  `map[string]interface{}` to `core.ScopeType`.
- `Lexer.Parse` accepts `i ...interface{}`.

## [1.1.2] - 2026-06-18

### Added

- **Package reshuffle:** `stringParsing/` -> `parsing/stringParsing/`,
  `byteParsing/` -> `parsing/byteParsing/`, `events/` -> `engine/events/`.
- `parsing/interface.go` with the generic `parsing.ParserInterface[I, P]`
  (`Parse(I, ...interface{}) ([]P, error)`); the duplicate interfaces in
  `byteParsing` and `stringParsing` were removed.
- **New package `tooling/plugin`:** the `Plugin` type, the `PluginManager` with
  `AddPlugin`, `DeletePlugin`, `GetPlugin`, `CallPlugin`, the `Init` / `Close` /
  `Main` events; the builder gains
  `WithPluginManager(plugins ...*plugin.Plugin)`.
- `EngineUniversal` gains the `Plugins` field; the constants
  `EuPtrPluginsScope` and `PluginManagerEuScope = "LC-PM"`.
- Examples `example/configLang1` (own parser) and `example/configLang2`
  (on `stringParsing.Parser1`).
- Apache 2.0 license, GitHub issue templates, `example/README.md`.

### Breaking

- **Dynamic library loading was removed:** the `plugin.Open` mechanism, the
  `NewPluginFunction` type, the `EngineUniversal.LoadPluginFromFile(path)`
  method, the `pluginFileSymbolName` constant and the `plugin` import were
  deleted. Plugins are now registered programmatically only, through
  `PluginManager`.
- The old `ParserInterface` in `byteParsing` and `stringParsing` was removed in
  favour of the shared `parsing.ParserInterface`.

### Fixed

- `ByteCallEvent` now returns an error on an unknown opcode instead of
  silently skipping it.
- `NewLogger` initialises the `Logging` map, without which any logger access
  panicked.
- `Logger.Logging map[string]bool`: output to stdout only for the allowed
  statuses, `Log` writes strings without ANSI sequences.
- Fixed a typo in the scope key `BYECODE_IDX` -> `BYTECODE_IDX`.
- `EngineBuilder.Build` adds the `EngineBuilder.Build:` prefix to error
  messages.

## [1.0.0] - 2026-06-17

The first stable release.

### Added

- **Two engines:** `StringEngine` (text commands, scope keys `INPUT string`
  and `PARSED []ParsedNode`) and `ByteEngine` (opcode commands, keys
  `ENDIANESS int`, `BYTECODE_IDX *int`, `INPUT []byte`, plus
  `AddToBytecodeIdx`, `SetBytecodeIdx`, `GetBytecodeIdx` and the
  `AutoBytecodeIndexShift` flag).
- **`UniversalEngineParams`** - a shared container for handlers: `Generator`,
  `Event`, `Scope`, `Logger`, `Context`; the `GetContext()` method.
- **Event system** - `Events` based on `orderedmap` with the methods
  `NewEvent`, `GetEvents`, `CallEvents`.
- **Package `events`** with `DefaultEvents`: `StringParsingEvent`,
  `StringCallEvent`, `ByteParsingEvent`, `ByteCallEvent`.
- **`Generator`** - accumulation of code at named pipeline points with the
  methods `AddString`, `AddBytes`, `GetStringArrRes`, `GetBytesRes`.
- **Parsers:** `stringParsing` (`Parser1` on grammar rules, `Parser2`
  line-based, `Lexer` on regular expressions with bracket balancing,
  `ParsedNode` with metadata and `__prev` / `__next` links), and `byteParsing`
  (`Parser1` with `Parser1Config`, `ShiftStruct`, `ParsedBytes`).
- **`tooling/bytecode`** - `Utils` for converting int and float64 to bytes and
  back respecting byte order, `InstructionsGenerator`, `GenerationConfig`, the
  constants `BigEndian` / `LittleEndian`.
- **`Logger`** - a thread-safe logger with the methods `PrintLog`, `GetLog`,
  `Statuses` and the `NewLogger` constructor.
- **`EngineBuilder`** - the Builder pattern with the methods `WithPipeline`,
  `WithContext`, `WithDefaultEvents`, `WithLogger`, `WithScope`, `WithColors`,
  `WithStringParser`, `WithByteParser`, `WithEndianess`, `Build`.
- **`EngineUniversal`** - the wrapper with `ProcessString`, `ProcessBytes`,
  `ProcessStringWithCtx`, `ProcessBytesWithCtx`, `GetUEP`, `NewCommandString`,
  `NewCommandByte` (automatic opcode assignment when `-1` is passed).
- Cancellation via `context.Context` is checked in the command dispatch loop of
  all engines.
- Tests: `byteParsing/parser1_test.go`, `stringParsing/lexer_test.go`,
  `engine/core/generator_test.go`.
- Dependencies: `iancoleman/orderedmap v0.3.0`, `dlclark/regexp2 v1.12.0`,
  `pt-main/tap v1.1.1`.

### Changed

- **Breaking:** `core.CommandType` and `CommandMeta` became generic
  (`CommandType[E, N any] func(*E, N) error`).
- **Breaking:** `core.EventType` is now `func(interface{}, *Events) error`.
- **Breaking:** `GetUEP()` returns `*core.UniversalEngineParams`, the `UEP`
  field in the engines became a pointer.
- **Breaking:** the `another engineAnother` field was removed from
  `EngineUniversal`, the opcode counter became `opcode_counter`.
- Coloured output via `pt-main/tap` (`color.Set`, `color.ColorEnabled`), colour
  codes `[?RD]` and `[?YW]` in error messages.

### Fixed

- `StringEngine` and `ByteEngine` are protected by `sync.RWMutex` - commands
  can be registered and read concurrently.
- `NewUniversalEngineParams` returns an error on `nil` parameters.
- Fixed `logE`, which always panicked on an incorrect `error` type.
- Removed redundant output in the byte call event.

### Removed

- The `engine/converts.go` file with the `Converts` type - replaced by generic
  types.

## [0.10.1] - 2026-06-14

### Changed

- `tooling/bytecode`: the `InstructionsGenerator` fields (`OpcodeLen`,
  `ArglenLen`, `ArgscountLen`, `Endianess`) were replaced by a single
  `bytecode.GenerationConfig` struct.
- `byteParsing.Parser1Config` no longer stores lengths and byte order
  separately: instead of them the `GConfig bytecode.GenerationConfig` field was
  introduced (the `Shifter` field is kept).

## [0.9.15] - 2026-06-09

An intermediate release with no public API changes: edits in
`engine/byteEngine.go`, `engine/core/events.go`, `events/byteEngine.go` and
`tooling/bytecode/utils.go`.

## [0.9.12] - 2026-06-08

### Added

- `float64` conversion in `tooling/bytecode`: `Utils.Float64ToBytes`,
  `BytesToFloat64`, `Float64ToBytesRange`, `BytesToFloat64Range`, plus the
  `Float64ToBytesBigEndian` / `Float64ToBytesLittleEndian` and
  `BytesToFloat64BigEndian` / `BytesToFloat64LittleEndian` variants.
- Shift methods `ShiftFloat64Error`, `ShiftFloat64Panic`,
  `ShiftFloat64RangeError`, `ShiftFloat64RangePanic`.
- `ByteEngine.GetBytecodeIdx() (*int, error)`.
- Documentation for `NewCommandByte` stating that the handler is responsible
  for advancing the bytecode index.

### Changed

- `SetBytecodeIdx` now stores a **pointer** `&n` in scope instead of the value,
  to match the behaviour of `AddToBytecodeIdx`.

## [0.9.11] - 2026-06-08

### Added

- Byte engine scope constants: `ByteEngineScopeEndianess`,
  `ByteEngineScopeBytecodeIdx`, `ByteEngineScopeInput`.
- The methods `ByteEngine.AddToBytecodeIdx(n int)` and `SetBytecodeIdx(n int)`.

### Changed

- Event strings lost the underscores:
  `"input string->parsed []ParsedNode"` and others.
- `ByteCallEvent` iterates the bytecode by the index from scope instead of
  through `range`, so a step-by-step traversal appeared with the possibility
  for a handler to append commands.

## [0.9.9] - 2026-06-08

### Changed

- `bytecode.Shift.code` was exported as `bytecode.Shift.Code`.
- `Utils.ShiftStruct(code, idx)` was replaced by the constructor
  `bytecode.NewShift(code []byte, idx *int) *Shift`.

### Fixed

- `byteParsing.Parser1.Parse` assigns `p.Config.Shifter.Code = code`.

## [0.9.8] - 2026-06-08

### Changed

- The type `engineUniversal` was exported as `engine.EngineUniversal`, the
  `Build()` signature changed to `(*EngineUniversal, error)`.
- `byteParsing.Parser1Config` gained the `Shifter bytecode.Shift` field: the
  shift was moved outside, `Utils.ShiftStruct` is no longer used in the parser.

## [0.9.7] - 2026-06-07

### Changed

- Bulk godoc documentation for `EngineBuilder`, `EngineUniversal`, `ByteEngine`,
  `Lexer`, `LexerRule`, `NewLexer`, `Parser1`, `NewParser1`, `ParsedNode`,
  `ParsedBytes`, `Events`, `NewEvents`, `ProcessString`, `ProcessBytes`,
  `NewCommandByte`, `NewCommandString`.
- `Lexer.Parse` stores the rest of the source in `__raw`, not the token itself.

## [0.9.1] - 2026-06-06

### Added

- The root files `builder.go` and `engine.go` (package `lc`).
- `lc.NewEngineBuilder(engineType int)` with the fluent methods `WithPipeline`,
  `WithDefaultEvents`, `WithLogger`, `WithScope`, `WithStringParser`,
  `WithByteParser`, `WithEndianess` and `Build()`.
- The constants `lc.ByteEngineType` and `lc.StringEngineType`.
- The type `engineUniversal` with the methods `ProcessString`, `ProcessBytes`,
  `GetUEP`, `NewCommandByte`, `NewCommandString`; `opcode = -1` means an
  auto-increment.

### Changed

- `CallEventsEvent` was replaced by the pair `CallEventsStartEvent` /
  `CallEventsEndEvent`; `Events.CallEvents` writes `call_name` and
  `call_error` to scope.
- The lexer writes the metadata keys `__raw` and `__value`.
- `Parser2.Parse` returns `addPrevNextNodes(result)`.

## [0.8.8] - 2026-06-06

A pre-release with a major internal package restructure.

### Added

- Package `tooling/bytecode` (based on the former `system/utils.go`): the
  `Utils` type, `IntToBytes` / `BytesToInt` conversion with variants for both
  byte orders, the constants `BigEndian` / `LittleEndian`, `ShiftStruct` with
  `ShiftError` and `ShiftPanic`, and `InstructionsGenerator` with the
  `Generate` method.
- `engine/core/logger.go` with the `Logger` type and the methods
  `GetStatusForm`, `PrintLog`, `GetLog`.
- `engine/core/universalEngineParams.go` with `UniversalEngineParams` and the
  `NewUniversalEngineParams` constructor.
- `stringParsing/utils.go` with `addPrevNextNodes` - linking nodes via the
  `__prev` / `__next` metadata.

### Changed

- `system/config.go`, `system/events.go`, `system/generator.go` moved to
  `system/core/`.
- `StringEngine` and `ByteEngine` no longer store `Scope`, `Generator` and
  `Event` separately - everything is collected in the `UEP` field.
- `Process` errors are wrapped with the event name.

## [0.6.3p] - 2026-06-05

### Added

- `Converts.ConvertByteCommandTypeArgs(args []interface{})` - a converter for
  byte command handler arguments.

## [0.6.3] - 2026-06-05

### Changed

- The event constants `ParseEvent` / `CallEvent` were renamed to
  `StringParseEvent` / `StringCallEvent`; `ByteParseEvent` and `ByteCallEvent`
  were added.
- `Events.CallEvents` got the `canWorkWithoutHandler bool` parameter and returns
  an error instead of a silent `nil`.

### Fixed

- `ConvertStringCommandTypeArgs` takes the node from `args[1]`, not `args[0]`.

## [0.6.2] - 2026-06-05

The first release with byte engine support.

### Added

- Package `byteParsing`: `ParserInterface` with the method
  `Parse(code []byte) ([]ParsedBytes, error)`, the type `ParsedBytes`
  (`Switch`, `Raw`, `Args`, `Metadata`), `Parser1Config` (`CommandBytelen`,
  `ArglenBytelen`, `ArgscountBytelen`), `Parser1` and `Utils`.
- `system/byteEngine.go` with `ByteEngine` and `events/byteEngine.go` with the
  byte events.
- `system/converts.go` with the `Converts` type.

### Changed

- The `parsing` package was renamed to `stringParsing`, `system/engine.go` to
  `system/stringEngine.go`, `events/engine.go` to `events/stringEngine.go`.
- `CommandType` became `func([]interface{}) error`, `EventType` became
  `func(interface{}) error` (both used to be bound to `*Engine`).

## [0.2.5] - 2026-06-03

### Changed

- **Breaking:** the lexer was moved from the `regexp` engine to
  `github.com/dlclark/regexp2` v1.12.0 with lookaround and backreference
  support.
- **Breaking:** `LexerRule.Pattern` is now `*regexp2.Regexp`, named group
  extraction was rewritten to `GetGroupNames()` / `GroupByName`.

### Renamed

- `parsing.LexerParser` -> `parsing.Lexer`,
  `parsing.NewLexerParser` -> `parsing.NewLexer`.
- `parsing.NewParser` -> `parsing.NewParser1`.

## [0.2.1] - 2026-06-02

### Added

- The lexer `parsing.LexerParser` and the `parsing.NewLexerParser` constructor;
  the type `LexerRule{Type string, Pattern *regexp.Regexp}`.
- The `ParserConfig.SkipEmptyLines` and `TrimBlocksSpace` fields.

### Fixed

- `Parser1.matchGrammar` skips an empty block only with `SkipEmptyLines`, and
  `TrimSpace` applies only with `TrimBlocksSpace`; `__raw` stores the original,
  untrimmed line.
- In `DefaultEvents.CallEvent` the `err := errors.New("")` stub was replaced
  with `var err error = nil`.

## [0.2.0] - 2026-06-02

### Fixed

- `DefaultEvents.CallEvent` no longer fails on an unknown command: a check for
  the presence of the entry in `e.Commands` was added.

## [0.1.5] - 2026-05-30

### Added

- `parsing.ParserInterface` with the method
  `Parse(code string) ([]ParsedNode, error)`.
- The line-based parser `parsing.Parser2`, writing the metadata keys `command`,
  `args`, `__raw`.

### Changed

- **Breaking:** the `NewEngine` signature was extended with the
  `add_default_events bool` and `parser parsing.ParserInterface` parameters.
- `system.Engine` gained the `Parser` field.
- `parsing.Parser` was renamed to `parsing.Parser1`.

### Fixed

- `NewGenerator` initialises `code[point]` for every pipeline element.

## [0.1.1] - 2026-05-30

The first public release.

### Added

- The root package `lc` with the `Version` constant and
  `NewEngine(generator_res_type int, pipeline []string) *system.Engine`.
- Basic types: `system.Engine`, `system.ScopeType`, `system.CommandType`,
  `system.CommandMeta`, `system.EventType`, `system.Events` based on
  `orderedmap`, `system.Generator`, `parsing.ParsedNode`,
  `events.DefaultEvents` with `ParsingEvent` and `CallEvent`.
- `NewEvents`, `NewEvent`, `CallEvents`, `GetEvents`; `AddString`,
  `AddStrings`, `AddBytes`, `GetBytesRes`, `GetStringArrRes`, `GetStringRes`
  on the generator.

### Changed

- `NewEngine` was moved from the `system` package to the root package `lc`.

---

## Change types

- **Added** - new functionality.
- **Changed** - changes in existing functionality.
- **Deprecated** - soon to be removed functionality.
- **Removed** - removed functionality.
- **Fixed** - bug fixes.
- **Security** - vulnerability fixes.
- **Breaking** - API breaking changes (called out in v1.5.0, v1.5.2, v1.5.7,
  v1.6.0, v2.0.0).
- **Performance** - performance optimisations.

## Links

- [GitHub](https://github.com/pt-main/lc)
- [Go Reference](https://pkg.go.dev/github.com/pt-main/lc)
- [Project Wiki](https://github.com/pt-main/lc/wiki)
