# parser3

Рекурсивный спуск на основе грамматики для лексера строк Lc.

В отличие от `parser1` (блочный, на регулярных выражениях) и `parser2`
(построчный), `parser3` принимает структурированную грамматику и строит
настоящее дерево синтаксиса. Это тот парсер, который нужен, когда у входа есть
вложенная структура: выражения, конфиги с блоками, маленькие языки.

```go
import "github.com/pt-main/lc/v2/parsing/stringParsing/parser3"
```

## Как это работает

```
код -> stringParsing.Lexer -> []ParsedNode -> parser3.Parser -> AST
```

Лексер превращает текст в плоский поток токенов. Парсер идет по этому потоку в
соответствии с вашей грамматикой и возвращает один корневой узел, в
метаданных которого лежат дочерние узлы.

## Быстрый старт

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

`err` имеет тип `core.ErrorInterface`, см. раздел [Ошибки](#ошибки).

## Конструкции грамматики

| Конструкция | Назначение |
|---|---|
| `TokenExpr{TokenType}` | Consume один токен заданного типа |
| `SequenceExpr{Exprs}` | Match все элементы по порядку |
| `ChoiceExpr{Alternatives}` | Побеждает первый подошедший вариант |
| `RepeatExpr{Expr, Min, Max}` | Повтор; `Max <= 0` означает без ограничения |
| `SeparatedRepeatExpr{Element, Sep, Min, Max}` | `elem (sep elem)*` |
| `OptionalExpr{Expr}` | Match один раз или ни разу |
| `AndExpr{Expr}` | Опережающая проверка: match, но ничего не consume |
| `NotExpr{Expr}` | Отрицающая проверка: ничего не consume |
| `PeekExpr{TokenType}` | Проверить тип следующего токена, ничего не consume |
| `NamedExpr{RuleName}` | Вызвать другое правило |
| `NodeExpr{NodeType, Expr}` | Обернуть результат в узел |
| `ActionExpr{Expr, Action}` | Обработатьmatched-узлы своим кодом |
| `PrattExpr{Atom, Prefixes, Infixes}` | Разбор выражений по приоритетам операторов |

### Правила

`Grammar` это `map[string]Rule`. У `Rule` есть `Name` (по ключу карты ищет
`NamedExpr`) и `Expr`. Третий аргумент `NewParser` - имя стартового правила.

Рекурсия поддерживается и ожидается. Правило может вызвать само себя в *другой*
позиции - именно так работают вложенные скобки:

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

Прямая левая рекурсия (`expr -> expr ...` без consume токена) отклоняется с
понятной ошибкой, вместо того чтобы повесить процесс.

### Игнорируемые токены

Типы токенов, переданные четвертым аргументом в `NewParser`, пропускаются везде,
где парсер смотрит на следующий токен, поэтому пробелы не нужно описывать в
грамматике. Передайте `nil` или пустой срез, если ничего не игнорировать.

## Приоритеты операторов

`PrattExpr` разбирает выражения методом подъема по приоритетам. `Infixes`
сопоставляет тип токена с приоритетом и ассоциативностью; больший `Precedence`
связывает сильнее.

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

`1 + 2 * 3` превращается в `(1 PLUS (2 STAR 3))`. Ассоциативность бывает
`LeftAssoc`, `RightAssoc` или `NonAssoc`. Префиксные операторы связывают
сильнее любого инфиксного.

В дереве используются такие ключи метаданных:

| Ключ | Узел | Значение |
|---|---|---|
| `operator` | `BinaryOp` | тип инфиксного токена |
| `left`, `right` | `BinaryOp` | поддеревья операндов |
| `operator` | `PrefixOp` | тип префиксного токена |
| `operand` | `PrefixOp` | поддерево операнда |

## Действия

`ActionExpr` выполняет Go-код над matched-узлами, для приведения значений или
проверок. Ошибка из действия прерывает разбор с `AdapterError`.

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

## Вид результата

`Parse` возвращает срез ровно из одного корневого узла с именем стартового
правила. Его `Raw` - склейка `Raw` дочерних узлов, а в метаданных `children`
лежат сопоставленные подузлы. Читайте их через константу `MetaChildren` или
функцию `ChildrenOf`:

```go
children, err := parser3.ChildrenOf(nodes[0])
```

Каждый узел - это `stringParsing.ParsedNode`:

| Поле | Значение |
|---|---|
| `Switch` | тип токена для листьев, тип узла для внутренних |
| `Raw` | точный текст исходника, покрытый узлом |
| `Metadata` | `__raw`, `__value`, `__start`, `__end`, именованные группы, `children` |

## Ошибки

Ошибки типизированы и несут позицию, поэтому их можно и показать человеку, и
разобрать программно.

| Тип | `GetCode()` | Когда возникает |
|---|---|---|
| `ParseError` | `parser3` | вход не соответствует грамматике |
| `GrammarError` | `parser3/grammar` | сама грамматика написана неверно |
| `AdapterError` | `parser3/adapter` | неверная форма AST или ошибка в действии |

`ParseError` хранит `Phase` (какая конструкция не сработала: `Lexer`,
`Expect`, `Peek`, `EOF`, `End`, ...), `Expected`, `Got`, `Raw`, `TokenIdx`,
`TokenPos` и `Found` (типы токенов, реально встретившиеся в точке сбоя).

`GrammarError.Phase` называет конструкцию грамматики (`NamedExpr`,
`ChoiceExpr`, `RepeatExpr`, ...), а `Msg` объясняет, что не так. `GetCode()`
всегда возвращает стабильный код пакета, а не фазу, чтобы по кодам ошибок можно
было надежно матчить.

Чтобы получить конкретный тип, используйте `AsParseError` / `AsGrammarError`,
либо `errors.As` для всей цепочки.

### Форматирование

```go
parser3.FormatError(err, false)   // обычный текст, для логов и JSON
parser3.FormatErrorPretty(err)    // ANSI-цвета, для терминала
```

Обе функции печатают всю цепочку ошибок, вложенные причины с отступом под
`caused by`:

```
parser3/Expect: expected "NUMBER", got "STAR" (raw: "*") at idx=4 start=4-5
```

Если одна ветка грамматики зашла дальше другой, сообщается самая глубокая
ошибка, поэтому текст указывает на настоящую причину, а не на первый
неподошедший токен.

## Производительность

Результат каждого именованного правила запоминается на позиции токена, поэтому
правило, достижимое из нескольких вариантов, разбирается один раз, а не по
разу на каждый путь. Ограниченная глубина рекурсии не дает патологической
грамматике исчерпать стек.

`Parser` небезопасен для конкурентного использования: используйте его
последовательно или создавайте отдельный на горутину. `Parse` сбрасывает все
внутреннее состояние, поэтому одним `Parser` можно пользоваться многократно.

## Адаптер движка

`Adapter` оборачивает `Parser` для движка, выполняя разбор и сразу возвращая
дочерние узлы стартового правила:

```go
a := &parser3.Adapter{Parser: p}
children, err := a.Parse("1 + 2")
```

## Пример

`example/tests/parser3Test` разбирает выражение калькулятора, печатает AST в
виде текста и JSON, показывает приоритеты Pratt и вывод ошибок для нескольких
некорректных входов.
