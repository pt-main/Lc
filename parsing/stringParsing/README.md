# String Parsing

Text parsers for the string engine. Each implements
`parsing.ParserInterface[string, stringParsing.ParsedNode]`.

| Parser | Behaviour |
|---|---|
| `Lexer` | Token-based lexer on standard library regex rules, optional bracket balance, prev/next links on nodes |
| `Parser1` | Regex grammar with line continuation, bracket balancing, block trimming |
| `Parser2` | Line parser for `command args ...` |
| `Parser3` | Grammar-driven recursive descent, see [parser3](parser3) |
| `Adapter` | `Parser3` wrapper that feeds a grammar into the engine |

`Parser3` has its own README: [English](parser3/README.md) and
[Russian](parser3/README-ru.md).

## Node shape

Every string parser produces `ParsedNode`:

```go
type ParsedNode struct {
	Raw      string    // exact source text
	Switch   string    // token type or command name
	Metadata ScopeType // named groups, __raw, __value, __prev, __next
}
```

`Switch` is the dispatch key: the string engine calls the command registered
under that name. `Metadata` carries whatever the parser found, including regexp
named groups and, for the basic parsers, `__raw` for the full original text and
`__prev` / `__next` links to neighbouring nodes.

## Choosing

`Parser2` when input is one command per line. `Parser1` when lines have internal
structure and multi-line constructs are needed. `Lexer` plus `Parser3` when the
language has real nesting. A custom `ParserInterface` when none of them fit -
the engine accepts any parser that emits `ParsedNode`.
