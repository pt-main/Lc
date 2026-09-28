# tooling/plugin

Система плагинов: интерфейс, базовая реализация, менеджер и небольшой набор
хелперов для плагинов, которым нужно знать друг о друге.

| Файл | Содержимое |
|---|---|
| `interface.go` | `PluginInterface` |
| `realization.go` | `Plugin`, базовая реализация на событиях |
| `manager.go` | `PluginManager`, хранит плагины и их общий scope |
| `tools.go` | `Tools`, флаги и проверки наличия плагинов |

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

`Name` должно быть константным и уникальным. `Init` получает scope менеджера и
сам менеджер - именно так плагин достает до движка: движок кладет указатель на
себя в этот scope под ключом `public.PluginsScopeEuPtr`, а указатель на
менеджер - под `public.EuScopePmPtr`.

Реализовать интерфейс напрямую можно. На практике `Plugin` покрывает обычный
случай, потому что четыре метода выше соответствуют стандартному жизненному циклу.

## Plugin

Плагин - это контейнер событий с тремя именованными событиями и двумя ключами
результата.

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

| Событие | Кто вызывает | Назначение |
|---|---|---|
| `InitEvent` | `AddPlugin` | Однократная настройка, получает менеджер |
| `MainEvent` | `Run` | Основная работа, получает вход вызывающего |
| `CloseEvent` | `Close` | Завершение |
| любое другое | `Call(name)` | Именованные точки входа |

`Run` и `Call` возвращают то, что обработчики положили в ключ scope с
результатом. Оба читают `ScopeRunResultKey`:

```go
p.Events.NewEvent("main_event", func(ev *core.Events, i *core.EventInput) core.ErrorInterface {
	ev.Scope()[p.ScopeRunResultKey] = compute(i.Input)
	return nil
})

res, err := p.Run("input")
```

Обратите внимание: `ScopeCallResultKey` сохраняется в структуре, но не
читается - `Call` тоже возвращает значение из `ScopeRunResultKey`. Плагин,
который возвращает разные значения из `Run` и `Call`, должен писать оба в ключ
run или читать результат самостоятельно.

`Init` вызывается с менеджером в `EventInput.Input`, так что обработчик может
добраться до него и оттуда, и из scope. Базовая реализация ничего не
изолирует: плагин может напрямую читать и писать scope движка.

Имена событий хранятся в структуре (`InitEvent`, `MainEvent`, `CloseEvent`,
`ScopeRunResultKey`, `ScopeCallResultKey`) и публичны, поэтому их можно читать
или переприсваивать после создания.

## PluginManager

```go
pm := plugin.NewPluginManager(scope)
```

| Метод | Поведение |
|---|---|
| `AddPlugin` | Регистрирует по `Name()`, отклоняет дубликаты, затем вызывает `Init` |
| `DeletePlugin` | Вызывает `Close`, затем удаляет. Отсутствующее имя не ошибка |
| `GetPlugin` | Поиск по имени |
| `RunPlugin` | `GetPlugin` + `Run` |
| `CallPluginMethod` | `GetPlugin` + `Call` |
| `End` | Закрывает и удаляет все плагины |

`AddPlugin` регистрирует плагин до вызова `Init`, поэтому плагин, упавший на
`Init`, уже будет в карте. Убрать его должен вызывающий.

Менеджер доступен из движка как `engine.Plugins`, а `engine.End()` вызывает
`PluginManager.End()`, который по очереди закрывает все плагины. Если один
`Close` падает, остальные не закрываются.

Именно поэтому важен порядок регистрации. `profiler.Init` возвращает ошибку,
если `extensiblePlugin` еще не зарегистрирован, так что плагин расширения должен
идти первым.

## Tools

```go
t := &plugin.Tools{Pm: pm}
```

| Метод | Для чего |
|---|---|
| `HasFlag` / `SetFlag` | Приватные строковые флаги менеджера для связи плагинов |
| `IsPluginInstalled` | Проверка наличия другого плагина по имени |

`IsPluginInstalled` - так `profiler` проверяет свою зависимость.

## Регистрация на движке

На этапе сборки, через `WithPlugins`:

```go
engine, err := lc.NewEngineBuilder(public.StringEngineType, public.StringResType).
	WithStringParser(parser).
	WithPlugins(myPlugin).
	Build()
```

`Build` вызывает `Init` для каждого плагина. Если какой-то `Init` падает,
`Build` возвращает ошибку.

После сборки - так требуется плагинам, которым нужен указатель на движок:

```go
err := engine.Plugins.AddPlugin(extensiblePlugin.New(engine))
```

`extensiblePlugin.New` принимает движок, поэтому его нельзя передать в
`WithPlugins` - на тот момент движка еще не существует.
