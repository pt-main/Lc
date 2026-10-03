<h1 align="center">Lc - Language Creator and Devkit</h1>
<p align="center">
  <img alt="lc-banner" src="https://github.com/user-attachments/assets/8fa74598-5cee-403e-a9dc-417e86d22dcd" />
</p>
<p align="center">
  <a href="https://pkg.go.dev/github.com/pt-main/lc/v2"><img src="https://img.shields.io/badge/Go-Reference-007d9c?logo=go&logoColor=white" alt="Go Reference"></a>
  <a href="https://github.com/pt-main/lc/releases"><img src="https://img.shields.io/github/v/release/pt-main/lc?color=blue" alt="Latest release"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-Apache_2.0-yellow" alt="License Apache 2.0"></a>
  <a href="https://github.com/pt-main/lc/wiki"><img src="https://img.shields.io/badge/Project-Wiki-red" alt="Project Wiki"></a>
</p>
<p align="center">
  <a href="README.md">English</a> | <a href="README-ru.md">Русский</a>
</p>

**Lc** - это Go-фреймворк для создания рантаймов языков: интерпретаторов,
собственных VM, DSL, конфигурационных языков и обработчиков байткода.

Фреймворк дает то, что нужно любому рантайму языка и что утомительно писать
правильно: слой парсинга, событийный пайплайн, диспетчеризацию команд,
детерминированную сборку вывода, плагины, логирование и работу с context. Сам
язык - грамматика, команды, семантика - пишет пользователь.

```bash
go get github.com/pt-main/lc/v2
```

## Содержание

- [Область применения](#область-применения)
- [Архитектура](#архитектура)
- [Точки расширения](#точки-расширения)
- [Быстрый старт](#быстрый-старт)
- [Движки](#движки)
- [Основные концепции](#основные-концепции)
- [Парсеры](#парсеры)
- [Плагины](#плагины)
- [Отладка и tooling](#отладка-и-tooling)
- [Context и отмена](#context-и-отмена)
- [Производительность](#производительность)
- [Примеры](#примеры)
- [Лицензия](#лицензия)

## Область применения

Что такое Lc, коротко и по делу.

**Входит в область.** Слой выполнения языка: превращение входа в узлы,
диспетчеризация узлов в обработчики, цикл исполнения байткода, сборка вывода,
поддержка плагинов, отмена выполнения и типизированные ошибки.

**Не входит.** Инструменты для написания лексеров и грамматик, проверка типов,
оптимизация, компиляция в машинный код, проектирование стандартной библиотеки.
Lc разбирает текст в плоскую или вложенную структуру узлов и исполняет ее, но не
компилирует в машинный код и не поставляет язык.

**Осознанно не включено.** Встроенной песочницы и ограничения ресурсов нет.
Фреймворк не делает недоверенный байткод безопасным. Все, что обрабатывает
недоверенный ввод, должно задавать свои пределы, обычно через таймауты context.

Рантайм на Lc обычно устроен так: лексер или грамматика дает узлы, таблица
команд сопоставляет имена узлов с функциями на Go, а движок обходит ввод и
вызывает их. Lc предоставляет обход, диспетчеризацию, сборку вывода и жизненный
цикл, а структура узлов и семантика команд принадлежат языку.

## Архитектура

Выполнение - это пайплайн событий. Движок сохраняет вход, вызывает событие
парсинга, вызывает событие вызова, а обработчики делают работу. В этой
последовательности ничего не захардкожено: парсер поставляется пользователем,
менеджер событий - это интерфейс, а сам цикл вызова можно заменить.

```text
вход
  |
  v
[ Парсер ]  ->  []ParsedNode / []ParsedBytes
  |
  v
[ События ]  ->  диспетчеризация по имени узла или опкоду
  |
  v
[ Обработчики ]  ->  UEP.Generator, UEP.Scope, UEP.Logger
```

`UEP` (`UniversalEngineParams`) - это объект, который получает обработчик. В нем
лежат `Generator` для вывода, менеджер `Events`, `Scope` для общих данных,
`Logger` и активный `Context`. Так как все обработчики видят один и тот же UEP,
любой компонент можно заменить, не трогая код пользователя.

## Точки расширения

Интерфейсов пять, и все они заменяемы.

| Интерфейс | Что заменяет | Реализация по умолчанию |
|---|---|---|
| `parsing.ParserInterface[I, P]` | Чтение входа и получение узлов | `stringParsing`, `byteParsing` |
| `core.EventsInterface` | Менеджер событий и цикл вызова | `core.Events` |
| `core.LoggerInterface` | Диагностический вывод | `core.Logger` |
| `core.ErrorInterface` | Значения ошибок | `core.Error` |
| `plugin.PluginInterface` | Внешняя логика со своими событиями | `plugin.Plugin` |

Главный из них - парсер. `ParserInterface` обобщен и по типу входа, и по типу
узлов, а интерфейс движка обобщен по ключу команды:

```go
type ParserInterface[I any, P any] interface {
	Parse(I, ...*ParseOption) ([]P, core.ErrorInterface)
	String() string
}
```

**Практическая точка расширения - свой парсер, а не свой тип узлов.**
Конструкторы движков привязаны к встроенным типам узлов: `NewStringEngine`
принимает `ParserInterface[string, ParsedNode]`, а `NewByteEngine` -
`ParserInterface[[]byte, ParsedBytes]`, поэтому парсер, выдающий собственную
структуру узла, подключить к движку нельзя. Свой парсер означает реализацию
одного из этих интерфейсов с выдачей `ParsedNode` или `ParsedBytes`.

Этого хватает для значительного объема реальных задач, потому что `ParsedNode`
намеренно свободная структура:

```go
type ParsedNode struct {
	Raw      string    // точный исходный текст
	Switch   string    // тип токена или имя команды
	Metadata ScopeType // именованные группы, позиции, пользовательские данные
}
```

Свой парсер сам решает, что значит `Switch` и что класть в `Metadata`, поэтому
произвольные понятия языка переносятся как метаданные на стандартном узле:

```go
type UpperParser struct{}

func (UpperParser) Parse(code string, opts ...*parsing.ParseOption) ([]stringParsing.ParsedNode, core.ErrorInterface) {
	res := make([]stringParsing.ParsedNode, 0, len(code))
	for i, r := range code {
		if r < 'a' || r > 'z' {
			continue
		}
		res = append(res, stringParsing.ParsedNode{
			Switch:   "letter",
			Raw:      string(r),
			Metadata: core.ScopeType{"pos": i},
		})
	}
	return res, nil
}

func (UpperParser) String() string { return "UpperParser" }
```

Он подключается к движку как любой встроенный парсер, а обработчик - обычная
функция:

```go
engine, _ := lc.NewEngineBuilder(public.StringEngineType, public.StringResType).
	WithStringParser(UpperParser{}).
	Build()

engine.NewCommandString("letter", func(se enginepkg.StringEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
	pos, _ := node.Metadata["pos"].(int)
	return se.GetUep().Generator.AddString(fmt.Sprintf("[%d]%s", pos, node.Raw), "main")
}, "echo letter with position")

engine.ProcessString("hello")
// вывод: [0]h [1]e [2]l [3]l [4]o
```

Что заменить нельзя: типы узлов, которые принимают конструкторы, и, как следствие,
сигнатуру обработчика команды. Рантайму, которому нужна своя структура узла,
понадобится свой тип движка. Это осознанная граница, а не недостающая функция.

## Быстрый старт

### Строковый движок

Разбирает строки вида `команда аргументы` и пишет сгенерированный текст в точку
`Generator`.

```go
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/pt-main/lc/v2"
	enginepkg "github.com/pt-main/lc/v2/engine"
	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/parsing/stringParsing"
	"github.com/pt-main/lc/v2/public"
)

func main() {
	parser := &stringParsing.Parser2{}

	engine, err := lc.NewEngineBuilder(public.StringEngineType, public.StringResType).
		WithPipeline([]string{"main"}).
		WithStringParser(parser).
		WithDefaultEvents(true).
		Build()
	if err != nil {
		panic(err)
	}

	err = engine.NewCommandString("log", func(se enginepkg.StringEngineInterface, node *stringParsing.ParsedNode) core.ErrorInterface {
		args, _ := node.Metadata["args"].(string)
		return se.GetUep().Generator.AddString(fmt.Sprintf("Log [%v]: %v",
			time.Now().Format(time.Stamp), args), "main")
	}, "append log with timestamp")
	if err != nil {
		panic(err)
	}

	err = engine.ProcessString(strings.Join([]string{
		"log service_start",
		"log service_ready",
	}, "\n"))
	if err != nil {
		panic(err)
	}

	uep, _ := engine.GetUEP()
	out, err := core.GetStringRes(uep.Generator, "\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
```

```bash
go run ./example/readme/string
```

```text
Log [Aug  7 18:44:40]: service_start
Log [Aug  7 18:44:40]: service_ready
```

### Байтовый движок

Декодирует простой формат инструкций и вызывает обработчики опкодов.

```go
package main

import (
	"fmt"

	"github.com/pt-main/lc/v2"
	enginepkg "github.com/pt-main/lc/v2/engine"
	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/parsing/byteParsing"
	"github.com/pt-main/lc/v2/public"
	"github.com/pt-main/lc/v2/tooling/bytecode"
)

func main() {
	// instruction {
	//     [bytes : cmd] [bytes : argscount] [bytes : arglen]  [bytes arglen : arg],
	//                                       [bytes : arglen2] [bytes arglen2 : arg2]...
	// }
	parser := &byteParsing.Parser1{
		Config: byteParsing.Parser1Config{
			GConfig: bytecode.GenerationConfig{
				CommandBytelen:   1,
				ArgscountBytelen: 1,
				ArglenBytelen:    2,
				Endianness:        public.LittleEndian,
			},
			Shifter: bytecode.Shift{},
		},
	}

	engine, err := lc.NewEngineBuilder(public.ByteEngineType, public.StringResType).
		WithPipeline([]string{"main"}).
		WithByteParser(parser).
		WithDefaultEvents(true).
		Build()
	if err != nil {
		panic(err)
	}

	err = engine.NewCommandByte(1, func(be enginepkg.ByteEngineInterface, node *byteParsing.ParsedBytes) core.ErrorInterface {
		for _, arg := range node.Args {
			if err := be.GetUep().Generator.AddString(string(arg), "main"); err != nil {
				return err
			}
		}
		return nil
	}, "add to output instruction", true)
	if err != nil {
		panic(err)
	}

	code := []byte{
		0x01,              // opcode=1
		0x01,              // argsCount=1
		0x03, 0x00,        // arglen=3 (little endian, 2 bytes)
		0x61, 0x62, 0x63,  // args="abc" (3 bytes)
	}

	err = engine.ProcessBytes(code)
	if err != nil {
		panic(err)
	}

	uep, _ := engine.GetUEP()
	out, err := core.GetStringRes(uep.Generator, "")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", out)
}
```

```bash
go run ./example/readme/byte
```

```text
abc
```

## Движки

### Жизненный цикл строкового движка

Берет текст (исходный код, команды, конфиг) и делает с ним что-то: редактирует,
выполняет, переводит, генерирует по нему код.

1. Кладет вход в scope под ключом `public.StringEngineScopeInput`.
2. Парсит его в `[]stringParsing.ParsedNode` (`public.StringParseEvent`).
3. Вызывает обработчик по `ParsedNode.Switch` (`public.StringCallEvent`).
4. При необходимости пишет вывод через `UEP.Generator`.

### Жизненный цикл байтового движка

Берет байткод и выполняет его. Сделан под горячие циклы: разобранные инструкции
перед диспетчеризацией сводятся в компактный атрибут вызова.

1. Кладет вход в scope под ключом `public.ByteEngineScopeInput`.
2. Парсит его в `[]byteParsing.ParsedBytes` (`public.ByteParseEvent`).
3. Сводит в компактный атрибут вызова для горячего цикла.
4. Вызывает обработчик опкода.
5. Двигает указатель инструкции автоматически или вручную.

### AST-движок

`AstEngine` - не третий бэкенд. Он реализует тот же `StringEngineInterface` и
использует то же событие парсинга строк, но вместо диспетчеризации только
верхнеуровневых узлов обходит разобранное дерево и вызывает обработчики и для
вложенных узлов, с управлением по путям.

```go
astEngine := lc.NewAstEngine(
	public.StringResType,      // тип результата
	[]string{"main"},          // пайплайн
	true,                      // события по умолчанию
	adapter,                   // парсер
	context.Background(),
	true,                      // CanBeUnknown
	false,                     // CanMainNodeBeUnknown
)
```

У каждой команды можно задать правила по путям:

```go
ctx := astEngine.GetCommandCtx("block")
ctx.BreakIf = [][]string{{"block", "inline"}} // прекратить обход на этом пути
ctx.SkipIf = [][]string{{"block", "comment"}} // пропустить это поддерево
```

`CanBeUnknown` решает, является ли незарегистрированный тип узла ошибкой или
молча игнорируется. Именно это делает движок пригодным для реального ввода, где
части дерева рантайму не интересны.

Он создается через `NewAstEngine`, а не через билдер намеренно: он не дает
`EngineUniversal`, он является диспетчером поверх события парсинга строк, и он не
взаимозаменяем со строковым или байтовым движком.

## Основные концепции

### UEP

`UniversalEngineParams` собирает `Generator`, менеджер `Events`, `Scope`,
`Logger` и активный `Context`.

```go
engine, _ := lc.NewEngineBuilder(...).
	WithScope(core.ScopeType{"env": "production"}).
	Build()

uep, _ := engine.GetUEP()
uep.Generator.AddString("hello", "main")
```

### Events

Выполнение управляется событиями. Обработчик получает
`(*core.Events, *core.EventInput)`.

```go
events := core.NewEvents(context.Background())
events.NewEvent("event1", handler1)        // первый обработчик "event1"
events.NewEvent("event1", handler2)        // добавлен в конец
events.NewEventBefore("event1", handler3)  // вставлен в начало
// "event1" -> [handler3, handler1, handler2]
```

Первый обработчик события - это его **core event**. `core.EventsTools` заменяет
только этот обработчик, не трогая остальных зарегистрированных вокруг него. Именно
так плагин оборачивает цикл вызова, а не копирует его:

```go
et := core.EventsTools{Events: events}

old, _ := et.GetCoreEvent(public.ByteCallHotloopEvent)
et.ChangeCoreEvent(public.ByteCallHotloopEvent, myHandler)
```

События по умолчанию, которые используют движки:

| Движок | События |
|---|---|
| Строковый | `StringParseEvent`, `StringCallEvent`, `StringCallCallLoopEvent` |
| Байтовый | `ByteParseEvent`, `ByteCallEvent`, `ByteCallHotloopEvent` |

Сам менеджер событий заменяется реализацией `core.EventsInterface`.

### Generator

Собирает фрагменты кода в именованные точки пайплайна и отдает их в заданном
порядке. Потокобезопасен, работает со строками или байтами в зависимости от типа
результата.

```go
pipeline := []string{"pre", "main"}
generator := core.NewGenerator(public.StringResType, pipeline)
generator.AddStrings([]string{"string1 ", "string2."}, "main")
generator.AddStrings([]string{"string3 ", "string4. "}, "pre")
res, _ := core.GetStringRes(generator, "")
// res = "string3 string4. string1 string2."
```

`Generator.Pipeline` публичный, поэтому порядок выдачи можно менять во время
работы, ничего не генерируя заново.

### Scope

Потокобезопасный `map[string]interface{}`, общий для обработчиков событий,
парсеров и команд. Используется для передачи данных между этапами пайплайна.

```go
engine, _ := lc.NewEngineBuilder(...).
	WithScope(core.ScopeType{
		"tenant_id": "prod-001",
		"env":       "production",
	}).
	Build()

// позже, в обработчике команды:
tenant, _ := core.ScopeGet[string](se.GetUep().Scope, "tenant_id")
```

`core.ScopeGetSynced` и `core.ScopeSetSynced` берут блокировку scope - для
scope, который пишется из одной горутины, а читается из другой.

**Важно:** ключи, объявленные в пакете `public`, используются событиями по
умолчанию. Их перезапись в своем обработчике меняет поведение движка.

### Logger

Потокобезопасный структурированный логгер в UEP, со своими форматами для
статусов и фильтрацией по уровням.

```go
logger := core.NewLogger("")   // формат по умолчанию: "%s [%v] [%s]\n"
logger.Logging["debug"] = true  // печатается только debug
logger.PrintLog("debug", "processing node: "+node.Switch)
```

Свой формат для статуса:

```go
logger := core.NewLogger("")
logger.Statuses["warn"] = "WARN [%v] %s\n"
logger.PrintLog("warn", "This is a warning")
```

Подключается через `WithLogger(logger)`. `MaxLogLength` ограничивает размер
хранимого лога, `GetLog()` возвращает его одной строкой.

### Ошибки

Ошибки, которые возвращают обработчики, имеют тип `core.ErrorInterface` и несут
код пакета, сообщение, метаданные и цепочку причин.

```go
err := engine.ProcessString("...")

var lerr *core.Error
if goerrors.As(err, &lerr) {
	fmt.Println(lerr.GetCode(), lerr.GetMsg())
}
```

`Format()` выводит всю цепочку с отступами - именно это делает пример
калькулятора, печатая ошибку вычисления. Коды ошибок объявлены в
`public/errors`.

## Парсеры

### Строковые парсеры

| Парсер | Поведение | Типичное применение |
|---|---|---|
| `Lexer` | Токенайзер на правилах стандартной библиотеки, баланс скобок, связи prev/next | Токенизация, вход для грамматики |
| `Parser1` | Грамматика на регулярных выражениях, перенос строк, баланс скобок | Построчные DSL, конфигурационные форматы |
| `Parser2` | Простой построчный парсер `команда аргументы` | Прототипы, языки вроде shell |
| `Parser3` | Рекурсивный спуск с комбинаторами, действиями и приоритетами операторов | Настоящие грамматики, генерация AST |

`Parser2` - самый простой:

```go
parser := &stringParsing.Parser2{}
// Вход: "print hello world"
// Выход: ParsedNode{Switch: "print", Metadata: {"args": "hello world"}}
```

В `Parser3` есть 13 комбинаторов (`TokenExpr`, `SequenceExpr`, `ChoiceExpr`,
`RepeatExpr`, `SeparatedRepeatExpr`, `OptionalExpr`, `AndExpr`, `NotExpr`,
`PeekExpr`, `NamedExpr`, `NodeExpr`, `ActionExpr`, `PrattExpr`), мемоизация
правил и типизированные ошибки с позициями. Полный справочник -
в [parsing/stringParsing/parser3](parsing/stringParsing/parser3).

### Байтовый парсер

`byteParsing.Parser1` декодирует бинарные инструкции с настраиваемой длиной полей
и разрядностью.

```go
parser := &byteParsing.Parser1{
	Config: byteParsing.Parser1Config{
		GConfig: bytecode.GenerationConfig{
			CommandBytelen:   1,
			ArgscountBytelen: 1,
			ArglenBytelen:    1,
			Endianness:        public.LittleEndian,
		},
		Shifter: bytecode.Shift{},
	},
}
```

## Плагины

Плагин - это набор событий со своим scope, регистрируемый через менеджер
плагинов. Плагины не изолированы: у них есть доступ к движку и к менеджеру.

```go
import "github.com/pt-main/lc/v2/tooling/plugin"

myPlugin := plugin.NewPlugin(
	"my_plugin",  // имя
	"init_event", // вызывается при Init
	"main_event", // вызывается при Run
	"close_event",
	"run_result", // ключи scope, куда Run/Call кладут результат
	"call_result",
	context.Background(),
)

myPlugin.Events.NewEvent("main_event", func(ev *core.Events, i *core.EventInput) core.ErrorInterface {
	// i.Input - это то, что передали в Run
	return nil
})
```

Регистрация и вызов:

```go
engine, _ := lc.NewEngineBuilder(...).
	WithPlugins(myPlugin). // вызывает "init_event"
	Build()

result, err := engine.Plugins.RunPlugin("my_plugin", "some input") // вызывает "main_event"
```

`CallPluginMethod` запускает именованное событие и возвращает то, что обработчик
положил в ключ scope с результатом:

```go
res, err := engine.Plugins.CallPluginMethod("profiler", "report")
```

`engine.End()` закрывает движок и вызывает `Close()` у каждого плагина.

## Отладка и tooling

- `plugin` - менеджер плагинов и базовая реализация `PluginInterface`.
- `bytecode` - преобразование int/float в байткод и обратно, а также шифтер,
  который использует байтовый парсер.
- `astools` - хелперы для AST: `GetChildren`, `FindChild`, `FindChildren`,
  `Walk`, `WalkWithPath`.
- `debugging/extensiblePlugin` - заменяет стандартный цикл вызова на такой, в
  который можно добавить хуки. Это основа, на которой построен профилер.
- `debugging/profiler` - счетчики и тайминги по каждой команде, требует
  `ExtensibleCLPlugin`.

`ExtensibleCLPlugin` требует ссылку на движок, поэтому его добавляют после
`Build()`:

```go
engine, _ := lc.NewEngineBuilder(...).
	WithStringParser(parser).
	Build()

err := engine.Plugins.AddPlugin(extensiblePlugin.New(engine))
if err != nil {
	return err
}

err = engine.Plugins.AddPlugin(profiler.New())
if err != nil {
	return err
}

report, err := engine.Plugins.CallPluginMethod("profiler", "report")
```

Профилер выдает числа по каждой команде:

```text
Profiler report (0.00 sec total):

  String commands:
  String calls: 12 (total time: 32.848µs)
    keyval: count=9, total=26.352µs, avg=2.928µs, min=1.163µs, max=9.907µs
    section: count=2, total=3.652µs, avg=1.826µs, min=1.456µs, max=2.196µs
```

## Context и отмена

У каждого метода `Process*` есть вариант `WithCtx`, принимающий
`context.Context`. Это покрывает таймауты, мягкое завершение и значения в рамках
запроса.

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

err := engine.ProcessStringWithCtx(input, ctx)
if goerrors.Is(err, context.DeadlineExceeded) {
	fmt.Println("Execution timed out")
}
```

Context можно передать и на этапе сборки через `WithContext(ctx)`. Вызов
`engine.End()` отменяет его с указанием причины.

**Отмена кооперативная.** Context не применяется прерыванием: дедлайн сам по себе
не останавливает уже выполняющийся обработчик. Обработчик видит context через
UEP и сам решает, когда вернуться:

```go
select {
case <-se.GetUep().GetContext().Done():
	return core.Err("CMD", "cancelled")
default:
	// продолжаем
}
```

Без этой проверки долгий обработчик доработает до конца, даже когда дедлайн уже
прошел.

## Производительность

Байтовый движок сделан под горячие циклы: сырой цикл выдает примерно
170-200 млн операций в секунду на `i7-4770HQ`. Полный бенчмарк и инструкции по
запуску - в [example/tests/speedtest](example/tests/speedtest).

## Примеры

Все примеры лежат в [example/](example).

### `readme/`

Два сниппета из этого README, минимальные и самодостаточные.

```bash
go run ./example/readme/string
go run ./example/readme/byte
```

### `langs/`

Готовые языки, каждый работает на движке.

| Пример | Что показывает |
|---|---|
| `langs/calculator` | Выражения с float через `Parser3` и команду движка, раздельный вывод ошибок парсинга и вычисления |
| `langs/math` | Целочисленный язык с переменными, присваиванием и вычислением через scope |
| `langs/configLang` | Транслятор ini-конфига в JSON с подключенным профилером |

```bash
$ go run ./example/langs/calculator '(2+3)**4'
Lc version - 2.0.0
Result: 625

$ go run ./example/langs/calculator '(2-3*4)+(5*2-1)*2'
Lc version - 2.0.0
Result: 8

$ go run ./example/langs/math 'a := 5
b := a * 2
a + b'
15
```

`langs/calculator` разделяет два вида ошибок. Некорректное выражение дает
ошибку парсера с позицией:

```bash
$ go run ./example/langs/calculator '2 +'
Lc version - 2.0.0
Parse error:
 parser3/Expect: expected "MUL", got "PLUS" (raw: "+") at idx=2 start=2-3
```

Корректное выражение, упавшее при вычислении, дает цепочку обернутой ошибки:

```bash
$ go run ./example/langs/calculator '1/0'
Lc version - 2.0.0
Eval error:
 String:PROCESS_ERR2: EVENT_ERROR: Event handler failed
  Caused by:
    |DEFAULT_EVENTS:CONTEXT_ERROR: Error at line: "1/0"
```

`langs/configLang` печатает отчет профилера рядом с переведенным JSON:

```bash
$ go run ./example/langs/configLang
Lc version - 2.0.0
Profiler report (0.00 sec total):

  String commands:
  String calls: 12 (total time: 32.848µs)
    keyval: count=9, total=26.352µs, avg=2.928µs, min=1.163µs, max=9.907µs
    section: count=2, total=3.652µs, avg=1.826µs, min=1.456µs, max=2.196µs
```

### `tests/`

| Пример | Что показывает |
|---|---|
| `tests/parser3Test` | Сам парсер без движка: AST в виде текста и JSON, приоритеты Pratt, форматирование ошибок |
| `tests/speedtest` | Бенчмарки байтового движка |

```bash
go run ./example/tests/parser3Test
go test -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ ./example/tests/speedtest/byte/bench
```

### `packages/`

Базовые компоненты по отдельности, когда важно поведение одного кирпичика.

| Пример | Что показывает |
|---|---|
| `packages/engine/core/events` | Основы событий, core events, `EventsTools`, обмен через scope |
| `packages/engine/core/generator` | Генерация строк и байт, смена порядка пайплайна на лету |
| `packages/engine/core/logger` | Фильтрация уровней, свои форматы, обрезка лога |
| `packages/engine/core/other` | Чтение и запись в `Scope` |

`packages/engine/engines/string` и `packages/engine/engines/byte` - заглушки,
пока ничего не делают.

## Ченджлог

Все значимые изменения перечислены в [docs/changelog-ru.md](docs/changelog-ru.md),
английская версия - в [docs/changelog.md](docs/changelog.md).

## Лицензия

Apache 2.0 - см. [LICENSE](LICENSE).

By Pt.
