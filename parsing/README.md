# Parsing

The parser layer. Every parser implements `parsing.ParserInterface` and is
supplied to an engine by the user.

```go
type ParserInterface[I any, P any] interface {
	Parse(I, ...*ParseOption) ([]P, core.ErrorInterface)
	String() string
}
```

| Package | What it contains |
|---|---|
| `stringParsing/` | Text parsers: `Lexer`, `Parser1`, `Parser2`, `Parser3` |
| `byteParsing/` | Binary instruction decoder: `Parser1` |
| `stringParsing/parser3/` | Grammar-driven recursive descent, see its own README |

`ParseOption` carries the engine UEP and a `core.Option` into `Parse`, which
lets a parser read engine state. All bundled parsers accept it; whether it is
used is up to the parser.

## stringParsing

| Parser | Behaviour | Typical use |
|---|---|---|
| `Lexer` | Token-based lexer on `regexp2` rules, optional bracket balance, prev/next links on nodes | Tokenization, input for a grammar |
| `Parser1` | Regex grammar with line continuation, bracket balancing, block trimming | Line-oriented DSLs, config formats |
| `Parser2` | Simple `command args` line parser | Prototyping, shell-like languages |
| `Parser3` | Recursive descent with combinators, actions and operator precedence | Real grammars, AST generation |
| `Adapter` | `Parser3` wrapper for the string engine | Feeding a grammar into the engine |

`Parser3` has its own documentation in
[stringParsing/parser3](stringParsing/parser3), with a Russian translation in
[README-ru.md](stringParsing/parser3/README-ru.md).

## byteParsing

`Parser1` decodes binary instructions where each field has a configurable
length, in a configurable endianness. Field layout:

```text
[cmd] [argscount] ([arglen] [arg])...
```

`tooling/bytecode` provides the matching encoder, so input can be produced and
consumed by the same layout.

## Choosing

`Parser2` for anything shell-like, `Parser1` for line-oriented formats,
`Parser3` when the grammar has real nesting, and a custom parser when none of
them fit. The node type is fixed by the engine constructor, so a custom parser
emits `ParsedNode` or `ParsedBytes` - see the extension points section of the
root README.
