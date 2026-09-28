# engine/core

Базовые примитивы, из которых построен любой движок Lc. Здесь ничего не знает
ни о строках, ни о байткоде - все либо обобщено, либо является контейнером.

| Файл | Содержимое |
|---|---|
| `universalEngineParams.go` | `UniversalEngineParams` (UEP) - объект, который получает любой обработчик |
| `events.go` | `Events`, менеджер событий, и `EventsTools` |
| `generator.go` | `Generator`, упорядоченная сборка вывода |
| `scope.go` | `ScopeType` и типизированные хелперы scope с блокировкой |
| `logger.go` | `Logger`, структурированная диагностика |
| `errors.go` | `Error`, `ErrorInterface` и конструкторы ошибок |
| `types.go` | `CommandType`, `EventType`, `Option`, `EventInput` |

## UEP

`UniversalEngineParams` собирает все, что нужно обработчику. Он встроен в оба
типа движков, поэтому `engine.UEP` работает везде.

```go
type UniversalEngineParams struct {
	Generator *Generator
	Event     EventsInterface
	Scope     ScopeType
	Logger    LoggerInterface
	Context   context.Context
}
```

`GetContext()` возвращает `context.Background()`, если context не задан, поэтому
обработчику не нужна проверка на nil.

## Events

Событие - это упорядоченный список обработчиков под именем. `CallEvents`
вызывает их в порядке регистрации и останавливается на первой ошибке.

```go
events := core.NewEvents(context.Background())

events.NewEvent("cmd", handler1)        // становится core event
events.NewEvent("cmd", handler2)        // добавляется после него
events.NewEventBefore("cmd", handler3)  // вставляется перед ним

events.CallEvents(&core.EventInput{}, "cmd", false)
```

Первый зарегистрированный обработчик события - его **core event**, и его позиция
отслеживается отдельно от самого списка. `EventsTools` использует эту позицию,
чтобы заменить ровно core-обработчик:

```go
et := core.EventsTools{Events: events}

old, _ := et.GetCoreEvent("cmd")             // прочитать
et.ChangeCoreEvent("cmd", replacement)       // заменить
```

Именно этим пользуется `ExtensibleCLPlugin`, чтобы обернуть цикл вызова движка
без его копирования: плагин ставит свой core event и сохраняет исходный, чтобы
восстановить его в `Close()`.

Параметр `canWorkWithoutHandler` определяет, является ли отсутствующее событие
ошибкой. Пайплайны движков передают `true` для событий, которые могут
отсутствовать.

`SetProperty("debug", true)` включает оборачивание в `CallEventsStartEvent` и
`CallEventsEndEvent`, что трассирует диспетчеризацию событий. По умолчанию
выключено, и сам фреймворк его не включает - без этого эти два события не
вызываются. Обработчики за ними регистрирует `NewUniversalEngineParams`, и они
пишут через логгер, поэтому для вывода нужен включенный уровень "event":

```go
events.SetProperty("debug", true)
logger.Logging["event"] = true
```

Потокобезопасность: `Events` защищает свои карты `sync.RWMutex`, но `Scope()`
возвращает саму карту и не является синхронизированным доступом.

## Generator

Собирает фрагменты в именованные точки и отдает их в порядке `Pipeline`.

```go
g := core.NewGenerator(public.StringResType, []string{"pre", "main"})
g.AddString("header", "main")
g.AddString("body", "main")

res, _ := core.GetStringRes(g, "")
```

Тип результата фиксируется при создании. Смешивание типов возвращает
`GeneratorGenerationTypeError`, а не панику:

- `StringResType` принимает `AddString` / `AddStrings` и отдает
  `GetStringRes` / `GetStringArrRes`;
- `ByteResType` принимает `AddBytes` и отдает `GetBytesRes`.

`Pipeline` - публичное поле, его можно переприсвоить между вызовами, чтобы
изменить порядок выдачи, ничего не генерируя заново. Точка из `Pipeline`, в
которую ни разу ничего не писали, считается ошибкой, потому что ее позиция в
выводе не определена.

## Scope

Обычный `map[string]interface{}`, оставленный мапой, чтобы существующий код вида
`scope["k"] = v` продолжал компилироваться.

```go
val, err := core.ScopeGet[int](scope, "count")
err = core.Err("CMD", "not set").WithMeta(core.EMK(0, "string"), "count")
val, err = core.GetMetaValue[int](err, 0, "string")
```

`ScopeGet` проверяет тип: несовпадение типа сообщается как `ScopeGetError`, а не
молча возвращает ноль.

Варианты с суффиксом `Synced` берут блокировку конкретной карты, найденной по
указателю на мапу:

```go
core.ScopeSetSynced(scope, "key", val)
val, _ := core.ScopeGetSynced[int](scope, "key")
core.ScopeDeleteSynced(scope, "key")
```

Их нужно использовать, когда одну карту пишет одна горутина, а читает другая.
Прямой доступ к мапе не защищен, даже хотя окружающие типы потокобезопасны.

## Logger

Структурированное логирование с форматами на каждый статус и фильтрацией по
уровням.

```go
l := core.NewLogger("")
l.Logging["debug"] = true
l.PrintLog("debug", "message")
```

`PrintLog` пишет в stdout только если `Logging[status]` истинно, и при этом
всегда добавляет запись во внутренний срез. Формат - `fmt.Sprintf` с тремя
подстановками: статус, время, сообщение.

```go
l := core.NewLogger("STATUS:[?RD]%v [?RT]Time:[%v] [?RT]Text:[?v][?RT]")
l.Statuses["warn"] = "WARN [%v] %s\n"
l.SetStatusForm("info", "INFO [%v] %s\n")
```

`MaxLogLength` ограничивает размер хранимого среза: при превышении отбрасывается
самая старая запись. Значение `-1` (по умолчанию) означает без ограничений.
`GetLog()` склеивает сохраненные записи переводами строк.

## Ошибки

`Error` несет код, сообщение, метаданные и необязательную причину. Он
удовлетворяет и `error`, и `ErrorInterface`.

```go
func handle(node *ParsedNode) core.ErrorInterface {
	if bad(node) {
		return core.Err("MY_CODE", "bad node %v", node.Switch).
			WithMeta(core.EMK(0, "string"), node.Switch)
	}
	return nil
}
```

`Wrap` добавляет причину - так слой дополняет контекст, не теряя исходную
ошибку:

```go
return core.Wrap("MY_CODE", err, "while processing %v", name)
```

`Format()` выводит всю цепочку с отступами, печатая метаданные по ходу:

```text
String:PROCESS_ERR2: EVENT_ERROR: Event handler failed
  Caused by:
    |MY_CODE: bad node
    |  Meta:
    |    |0_META:string: broken
```

`GetErr` разворачивает один уровень, превращая обычную причину `error` в
`WrappedError`, чтобы цепочка оставалась разбираемой. `GetRealErrorReverse`
выводит цепочку от внутренней ошибки к внешней, что читается лучше в логах, где
важна корневая причина.

Ключи метаданных создаются через `EMK(n, valType)` и читаются обратно через
`GetMetaValue[T]`, поэтому ключ и ожидаемый тип не могут разойтись.

`ErrExit` - это сигнал: возврат его из обработчика останавливает выполнение, не
считаясь ошибкой. `example/langs/configLang` использует его для команды раннего
выхода.

## Типы

```go
type CommandType[EI, N any] func(EI, *N) ErrorInterface
type CommandMeta[EI, N any] struct {
	Handler CommandType[EI, N]
	Doc     string
}
type EventType func(*Events, *EventInput) ErrorInterface
```

Команда - это функция от интерфейса движка и указателя на узел к ошибке.
`CommandMeta` сопоставляет ее со строкой документации, которую и принимает
последним аргументом `NewCommandString`.

`EventInput` - это псевдоним `SimpleInput`, который несет `Option` (флаги и
scope) и произвольный `Input`.
