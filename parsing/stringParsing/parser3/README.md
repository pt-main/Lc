# parser3

Grammar-driven recursive-descent parser for the Lc string lexer.

Unlike `parser1` (block/regex based) and `parser2` (line based), `parser3` takes a
structured grammar and builds a real syntax tree. It is the parser to reach for
when input has nested structure: expressions, config with blocks, small languages.

```go
import "github.com/pt-main/lc/v2/parsing/stringParsing/parser3"
```

## Pipeline

```
code -> stringParsing.Lexer -> []ParsedNode -> parser3.Parser -> AST
```

The lexer turns text into a flat token stream. The parser walks that stream with
the grammar you supply and returns a single root node whose `children` metadata
holds the matched sub-nodes.

## Quick start

```go
lexer := stringParsing.NewLexer([]stringParsing.LexerRule{
    {Type: "NUMBER", Pattern: regexp2.MustCompile(`\d+`, 0)},
    {Type: "PLUS",  Pattern: regexp2.MustCompile(`\+`, 0)},
    {Type: "WHITESPACE", Pattern: regexp2.MustCompile(`\s+`, 0)},
}, nil)

grammar := parser3.Grammar{
    "sum": {Name: "sum", Expr: parser3.NodeExpr{
        NodeType: "sum",
        Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
            parser3.NamedExpr{RuleName: "number"},
            parser3.RepeatExpr{Expr: parser3.SequenceExpr{Exprs: []parser3.Expr{
                parser3.TokenExpr{TokenType: "PLUS"},
                parser3.NamedExpr{RuleName: "number"},
            }}},
        }},
    }},
    "number": {Name: "number", Expr: parser3.TokenExpr{TokenType: "NUMBER"}},
}

p := parser3.NewParser(lexer, grammar, "sum", []string{"WHITESPACE"})
nodes, err := p.Parse("1 + 2 + 3")
```

`err` is a `core.ErrorInterface`; see [Errors](#errors) below.

## Grammar constructs

| Construct | Purpose |
|---|---|
| `TokenExpr{TokenType}` | consume one token of a type |
| `SequenceExpr{Exprs}` | match every element in order |
| `ChoiceExpr{Alternatives}` | first alternative that matches wins |
| `RepeatExpr{Expr, Min, Max}` | repeat; `Max <= 0` means unlimited |
| `SeparatedRepeatExpr{Element, Sep, Min, Max}` | `elem (sep elem)*` |
| `OptionalExpr{Expr}` | match once or not at all |
| `AndExpr{Expr}` | lookahead: match, but consume nothing |
| `NotExpr{Expr}` | negative lookahead: consume nothing |
| `PeekExpr{TokenType}` | assert next token type, consume nothing |
| `NamedExpr{RuleName}` | call another rule |
| `NodeExpr{NodeType, Expr}` | wrap the result in a node |
| `ActionExpr{Expr, Action}` | post-process the matched nodes |
| `PrattExpr{Atom, Prefixes, Infixes}` | operator-precedence expression parsing |

### Rules

`Grammar` is `map[string]Rule`. A `Rule` has a `Name` (the map key is what
`NamedExpr` looks up) and an `Expr`. `NewParser` takes the start rule name as
its third argument.

Recursion is supported and expected. A rule may re-enter itself at a *different*
position, which is how nested parentheses work:

```go
"factor": {Expr: ChoiceExpr{Alternatives: []Expr{
    TokenExpr{TokenType: "NUMBER"},
    SequenceExpr{Exprs: []Expr{ // "(" expr ")"
        TokenExpr{TokenType: "LPAREN"},
        NamedExpr{RuleName: "expr"},
        TokenExpr{TokenType: "RPAREN"},
    }},
}}},
```

Direct left recursion (`expr -> expr ...` with no token consumed first) is
rejected with a clear error instead of hanging the process.

### Ignoring tokens

Token types passed as the fourth argument to `NewParser` are skipped wherever
the parser looks for the next token, so whitespace never has to appear in the
grammar. Use `nil` or an empty slice when nothing should be ignored.

## Operator precedence

`PrattExpr` parses expressions by precedence climbing. `Infixes` maps a token
type to a precedence and associativity; a higher `Precedence` binds tighter.

```go
&parser3.PrattExpr{
    Atom: parser3.TokenExpr{TokenType: "NUMBER"},
    Prefixes: map[string]parser3.Expr{
        "MINUS": parser3.TokenExpr{TokenType: "MINUS"},
    },
    Infixes: map[string]parser3.InfixInfo{
        "PLUS":  {Precedence: 1, Assoc: parser3.LeftAssoc},
        "MINUS": {Precedence: 1, Assoc: parser3.LeftAssoc},
        "STAR":  {Precedence: 2, Assoc: parser3.LeftAssoc},
    },
}
```

`1 + 2 * 3` becomes `(1 PLUS (2 STAR 3))`. Associativity is `LeftAssoc`,
`RightAssoc` or `NonAssoc`. Prefix operators bind tighter than any infix
operator.

The resulting tree uses these metadata keys:

| Key | Node | Meaning |
|---|---|---|
| `operator` | `BinaryOp` | infix token type |
| `left`, `right` | `BinaryOp` | operand sub-trees |
| `operator` | `PrefixOp` | prefix token type |
| `operand` | `PrefixOp` | operand sub-tree |

## Actions

`ActionExpr` runs Go code on matched nodes, for value conversion or validation.
Returning an error from the action aborts the parse with an `AdapterError`.

```go
parser3.ActionExpr{
    Expr: parser3.TokenExpr{TokenType: "NUMBER"},
    Action: func(nodes []stringParsing.ParsedNode) (stringParsing.ParsedNode, core.ErrorInterface) {
        n, _ := strconv.Atoi(nodes[0].Raw)
        if n > 100 {
            return stringParsing.ParsedNode{}, core.Err(errors.ParsingError, "value %d out of range", n)
        }
        return stringParsing.ParsedNode{
            Switch:   "int",
            Raw:      nodes[0].Raw,
            Metadata: map[string]interface{}{"value": n},
        }, nil
    },
}
```

## Result shape

`Parse` returns a slice with exactly one root node named after the start rule.
Its `Raw` is the concatenation of the children's `Raw`, and its `children`
metadata holds the matched sub-nodes. Use the `MetaChildren` constant or
`ChildrenOf` to read them:

```go
children, err := parser3.ChildrenOf(nodes[0])
```

Every node is a `stringParsing.ParsedNode`:

| Field | Meaning |
|---|---|
| `Switch` | token type for leaves, node type for inner nodes |
| `Raw` | exact source text covered by the node |
| `Metadata` | `__raw`, `__value`, `__start`, `__end`, named groups, `children` |

## Errors

Failures are typed and carry a position, so they can be rendered for humans or
inspected programmatically.

| Type | `GetCode()` | Raised when |
|---|---|---|
| `ParseError` | `parser3` | the input does not match the grammar |
| `GrammarError` | `parser3/grammar` | the grammar itself is wrong |
| `AdapterError` | `parser3/adapter` | the AST shape is wrong, or a user action failed |

`ParseError` records `Phase` (which construct failed: `Lexer`, `Expect`, `Peek`,
`EOF`, `End`, ...), `Expected`, `Got`, `Raw`, `TokenIdx`, `TokenPos` and `Found`
(token types present at the failure point).

`GrammarError.Phase` names the grammar construct (`NamedExpr`, `ChoiceExpr`,
`RepeatExpr`, ...) and `Msg` explains what is wrong. `GetCode()` always returns
the stable package code, never the phase, so error codes stay usable for
matching.

Inspect a specific type with `AsParseError` / `AsGrammarError`, or use the full
chain with `errors.As`.

### Formatting

```go
parser3.FormatError(err, false)   // plain text, for logs and JSON
parser3.FormatErrorPretty(err)    // ANSI colors, for a terminal
```

Both print the whole error chain, innermost causes indented under `caused by`:

```
parser3/Expect: expected "NUMBER", got "STAR" (raw: "*") at idx=4 start=4-5
```

When a branch of the grammar gets further than another, the deepest failure is
the one reported, so the message points at the real mistake rather than at the
first token that happened to fail.

## Performance

Results of each named rule are memoized per token position, so a rule reachable
from several alternatives is parsed once instead of once per path. Bounded
recursion depth prevents a pathological grammar from exhausting the stack.

`Parser` is not safe for concurrent use: reuse one sequentially, or build one
per goroutine. `Parse` resets all internal state, so a single `Parser` can be
reused across many calls.

## Engine adapter

`Adapter` wraps a `Parser` for the engine, running the parse and returning the
start rule's children directly:

```go
a := &parser3.Adapter{Parser: p}
children, err := a.Parse("1 + 2")
```

## Example

`example/tests/parser3Test` parses a calculator expression, prints the AST as
text and JSON, demonstrates Pratt precedence, and shows the error output for
several malformed inputs.
