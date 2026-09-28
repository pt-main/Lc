# tooling/debugging

Два системных плагина, которые делают цикл вызова движка наблюдаемым и
расширяемым хуками.

| Пакет | Роль |
|---|---|
| `extensiblePlugin` | Заменяет стандартный цикл вызова на такой, который стреляет хуками вокруг каждой команды |
| `profiler` | Строит тайминги по командам на этих хуках |

Оба используются в `example/langs/configLang`.

## extensiblePlugin

Цикл вызова движка - это один обработчик события, что обычно делает его
наблюдение неудобным: чтобы замерить команду, пришлось бы переписать
диспетчеризацию. Этот плагин заменяет ровно этот обработчик, поэтому все
остальное - события, команды, парсер - остается как было.

```go
engine, _ := lc.NewEngineBuilder(...).WithStringParser(parser).Build()

err := engine.Plugins.AddPlugin(extensiblePlugin.New(engine))
```

`New` принимает движок, поэтому этот плагин нельзя передать в `WithPlugins` на
этапе сборки: движка тогда еще нет. При `Init` плагин находит текущий core event
цикла вызова (`StringCallCallLoopEvent` для строкового движка,
`ByteCallHotloopEvent` для байтового), сохраняет его и ставит на его место свой.
`Close` возвращает исходный.

```go
name := public.StringCallCallLoopEvent // или ByteCallHotloopEvent
```

### Хуки

Замещающий цикл стреляет четырьмя событиями: два вокруг всего цикла, два вокруг
каждой команды.

| Константа | Когда стреляет |
|---|---|
| `CLEPreEvent` | Один раз, до начала цикла |
| `CLEInPreEvent` | Перед каждым обработчиком команды |
| `CLEInPostEvent` | После возврата каждого обработчика команды |
| `CLEPostEvent` | Один раз, после конца цикла |

Все четыре вызываются с `canWorkWithoutHandler = true`, поэтому если ничего не
зарегистрировано, они не стоят ничего. Возвращаемое значение хуков
отбрасывается: хук, вернувший ошибку, не останавливает цикл. Цикл останавливает
только упавший обработчик команды, обернутый метаданными
`ExtensiblePluginError`.

Профилер использует `CLEInPreEvent` и `CLEInPostEvent` - поэтому он умеет
мерить команды по отдельности.

### Данные

Перед хуками цикл сохраняет текущую позицию в scope под ключом
`CLEScopeData`:

| Движок | Тип |
|---|---|
| Строковый | `SCLEData` |
| Байтовый | `BCLEData` |

Оба - это `CLEData[I, P, E]`:

```go
type CLEData[I, P, E any] struct {
	Input  I      // сырой вход текущего запуска
	Idx    *int   // текущая позиция, общая с движком
	Parsed []P    // разобранные узлы
	PLen   int    // длина Parsed
	Ctx    context.Context
	E      E      // сам движок
}
```

Это тот же `*int`, который движок использует как указатель инструкции, поэтому
по нему видно, какая именно команда вот-вот выполнится. `SCLEData` отдает
`Parsed` как `[]ParsedNode` и `E` как `StringEngineInterface`; `BCLEData` -
`[]ByteCallAttr` и `ByteEngineInterface`.

Цикл записывает эти данные, а затем читает их обратно. Обработчик хука,
подменивший значение в scope, тем самым меняет то, что цикл использует дальше:
после возврата `CLEPreEvent` цикл берет все, что лежит в `CLEScopeData`,
включая измененные `Idx` или `Parsed`. Это поддерживаемый способ влиять на
диспетчеризацию из хука.

```go
uep.Event.NewEvent(extensiblePlugin.CLEPreEvent, func(ev *core.Events, i *core.EventInput) core.ErrorInterface {
	data, _ := core.ScopeGet[extensiblePlugin.SCLEData](ev.Scope(), extensiblePlugin.CLEScopeData)
	// усечь список узлов, чтобы остановить цикл досрочно
	data.PLen = 1
	ev.Scope()[extensiblePlugin.CLEScopeData] = data
	return nil
})
```

### Пример

Хук, который отклоняет команду во время выполнения:

```go
ext := extensiblePlugin.New(engine)
engine.Plugins.AddPlugin(ext)

uep, _ := engine.GetUEP()
uep.Event.NewEvent(extensiblePlugin.CLEInPreEvent, func(ev *core.Events, i *core.EventInput) core.ErrorInterface {
	data, err := core.ScopeGet[extensiblePlugin.SCLEData](ev.Scope(), extensiblePlugin.CLEScopeData)
	if err != nil {
		return nil
	}
	idx := data.Idx
	if idx != nil && *idx >= 0 && *idx < len(data.Parsed) {
		// data.Parsed[*idx].Switch - команда, которая вот-вот выполнится
	}
	return nil
})
```

## profiler

Считает и замеряет каждую команду. Он лишь регистрирует обработчики на двух
хук-событиях и больше ничего не делает, поэтому требует, чтобы
`extensiblePlugin` был установлен раньше.

```go
err := engine.Plugins.AddPlugin(profiler.New())

report, _ := engine.Plugins.CallPluginMethod("profiler", "report")
```

Порядок важен: `profiler.Init` вызывает `IsPluginInstalled(extensiblePlugin.Name)`
и возвращает ошибку, если плагин расширения отсутствует, поэтому добавлять его
нужно вторым.

### Методы

`Call` разбирает имя метода:

| Метод | Действие |
|---|---|
| `report` | Возвращает строку отчета |
| `reset` | Очищает метрики и перезапускает общий таймер |
| `enable` | Возобновляет сбор |
| `disable` | Останавливает сбор, метрики сохраняются |

`Run` - заглушка, профилер используется через `CallPluginMethod`.

### Отчет

```text
Profiler report (0.00 sec total):

  String commands:
  String calls: 12 (total time: 32.848µs)
    keyval: count=9, total=26.352µs, avg=2.928µs, min=1.163µs, max=9.907µs
    section: count=2, total=3.652µs, avg=1.826µs, min=1.456µs, max=2.196µs
```

Метрики байтов группируются по опкодам, строковые - по именам команд; обе
группы сортируются, чтобы вывод был стабильным. Секция появляется только для
команд, которые действительно вызывались.

Общее время считается от `New()`, а не от первой команды, поэтому включает время
до запуска движка.

### Накладные расходы

Теперь каждая команда проходит через два дополнительных вызова событий и по
одному `time.Now()` с каждой стороны. Для разработки и для сравнения весов это
приемлемо, но не бесплатно, и на горячем цикле байткода тайминги будут
определяться инструментацией, а не обработчиком. Измеряй с подключенным
профилером, а затем подтверждай бенчмарками из
`example/tests/speedtest`.
