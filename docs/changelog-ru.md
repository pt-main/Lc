# Changelog

Все значимые изменения проекта **Lc** (Go-фреймворк для создания рантаймов языков).

Формат основан на [Keep a Changelog](https://keepachangelog.com/ru/1.1.0/),
версионирование - [SemVer](https://semver.org/lang/ru/).

[Английская версия](changelog.md)

## [2.0.0] - 2026-09-28

Релиз о согласованности, корректности и документации. Архитектура не менялась:
пути всех пакетов и `go.mod` остались прежними, переименований нет. Основная
работа - сплошная вычитка публичного API, потокобезопасность, переработка
`parser3` и лексера, миграция документации на двуязычные README.

### Breaking

- Глобальное исправление опечаток в публичном API: `Endianess` -> `Endianness`
  (`EngineBuilder.WithEndianness`, `public.ByteEngineScopeEndianness`,
  `bytecode.GenerationConfig.Endianness`), `Calloop` -> `CallLoop`
  (`public.StringCallCallLoopEvent`, `public.AstCallCallLoopEvent`),
  `Loger` -> `Logger` (`core.LoggerInterface`), `Instaled` -> `Installed`
  (`plugin.Tools.IsPluginInstalled`).
- Изменены строковые значения констант: `ByteEngineScopeEndianness`
  (`"ENDIANNESS int"`), `ByteEngineScopeHotloopCtxCheckPeriod`
  (`"CTX_CHECK_PERIOD int"`), имена событий цикла (`"STRINGCALLLOOP"`,
  `"AST:CALLLOOP"`), коды ошибок default events (`"SYSTEM@DEFAULT_EVENTS"`,
  `"DEFAULT_EVENTS:..."`), хуков `extensiblePlugin` (`"CallLoopE ...`).
- Удалён код ошибки `errors.DefaultEventsCallErrorContex`, его использование
  перенесено в `DefaultEventsCallErrorContexted`.
- `engine.EngineInterface` получает метод `GetCommand(CmdT)` - сторонние
  реализации интерфейса перестают удовлетворять контракту.
- `parser3.ParseError.Code` и `parser3.GrammarError.Code` переименованы в
  `Phase`; создание таких ошибок через литерал с `Code:` больше не компилируется.
  `ParseError` также получает поле `Found []string` - список типов токенов
  в точке сбоя.
- `parser3.GrammarError.GetCode()` теперь всегда возвращает константу
  `GrammarErrCode` вместо кода фазы.
- `parser3.Adapter.Parse` больше не оборачивает ошибку разбора в
  `*AdapterError` - возвращается исходная ошибка парсера. Для проверки типа
  добавлены `parser3.AsParseError` и `parser3.AsGrammarError`.
- Формат вывода `parser3.FormatError(err, false)` изменён с однострочного
  `"a -> b -> c"` на многострочный с продолжениями `"  caused by: "`.
- `plugin.NewPlugin` исправлен: `ScopeRunResultKey` больше не получает
  значение `scopeCallResultKey`, а `Plugin.Call` читает `ScopeCallResultKey`.
  Плагины, писавшие результат в run-ключ, должны писать в call-ключ.
- `EngineUniversal.End()` стал идемпотентным (повторный вызов возвращает
  `CorePackageLcLifecycleError`), всегда отменяет context при заданном
  `CtxCancelCause`, освобождает scope guard-ы и возвращает ошибку плагинов
  напрямую, а не через строковое форматирование.
- `byteParsing.Parser1.Parse` возвращает `nil` вместо частично разобранного
  среза при ошибке и отклоняет отрицательный `argscount`, который раньше
  возникал из-за знакового расширения в `BytesToInt`.
- `stringParsing.Lexer.Parse` больше не паникует на `nil` или пустом
  `*parsing.ParseOption` - ранее `opts[0].UEP.Logger` разыменовывался без проверки.
- Удалены `EngineBuilder.WithColors()` и аргумент `colorEnable` (наследие
  удалённой в v1.6.0 зависимости `tap`).

### Fixed

- **Устранены гонки данных:** карты событий (`GetEvents`, `CoreEvents`,
  `SetEvents`, `ReplaceEvent`, `SetProperty`), `GetCommands` и `GetCommandCtx`
  всех трёх движков, а также флаг `enabled` профилировщика теперь защищены
  мьютексом или возвращают копии. `ByteCallEvent` больше не читает `e.Commands`
  и `e.AutoBytecodeIndexShift` без блокировки.
- `NewCommandByte`: авто-сдвиг индекса теперь применяется только при
  `autoBytecodeIdxShift == true` (раньше флаг выставлялся всегда, и значение
  `false` молча игнорировалось), а счётчик авто-операций при явном опкоде
  растёт как `max(opcode+1, counter)`, из-за чего авто-операция больше не может
  выдать уже занятый опкод.
- `EngineBuilder.Build` публикует менеджер плагинов и scope движка **до**
  вызова `Init` плагинов (иначе `Init` видел `nil eu.Plugins`), а при сбое
  `AddPlugin` вызывает `pm.End()` и отменяет context вместо утечки
  полуинициализированных плагинов.
- `End()` больше не паникует при повторном вызове и корректно освобождает
  scope guard-ы до сброса движков.
- `ByteCallHotLoopEvent` проверяет контекст на первой итерации, поэтому
  короткие программы больше не выполняются без реакции на отмену.
- `AstMakeCommandCtx` содержал ошибку копипасты: при `nil` в `skipif`
  значение `breakif` затиралось. Исправлено.
- `WalkWithPath` добавлял имя родителя вместо имени ребёнка и писал пути
  в общий буфер, из-за чего пути соседних узлов затирали друг друга.
- Указатели на дочерние узлы в `astools` теперь действительно указывают
  в дерево (`&children[i]`), а не на копии переменной цикла.
- `AstEngine.Process` читал разобранные узлы из пустого ключа scope `""`
  вместо `public.StringEngineScopeParsed`.
- `bytecode`: кодирование `float64` через `math.Round` вместо усечения
  `uint64(x * max)` и исправление кодировки диапазона [-1,1] в `(value+1)/2*max`
  делают `BytesToFloat64(Float64ToBytes(x))` точной обратной операцией.
- `astools.GetTokenValue` читает ключ `"__value"`, который лексер действительно
  пишет (раньше читался несуществующий ключ `"value"`).
- `GetCommandCtx` создаёт отсутствующий контекст (в v1.6.0 условие было
  инвертировано).
- `HasCommand` больше не возвращает фиктивную ошибку `"skip"`.
- Сопоставление путей `BreakIf` / `SkipIf` переведено с точного равенства
  на сравнение по префиксу, благодаря чему `BreakIf = [][]string{{"block","inline"}}`
  действительно отсекает поддерево.
- `WorkIter` пропускает корневой узел, чтобы обработчик корня не
  вызывался дважды.
- `core.GetErr` возвращает `nil` на `nil`-входе и `nil`-причине вместо паники.

### Added

- **Потокобезопасный доступ к scope:** `core.ScopeGetSynced`, `core.ScopeSetSynced`,
  `core.ScopeDeleteSynced`, `core.ReleaseScopeGuard` и реестр guard-ов
  `scopeGuards` с хранением самой карты, чтобы адрес не мог быть переиспользован
  после сборки мусора. `ScopeType` намеренно остаётся обычной картой,
  поэтому `scope[k] = v` продолжает компилироваться.
- `GetCommand` - точечный поиск команды без копирования всей карты команд -
  в `StringEngine`, `ByteEngine` и `AstEngine`, а также
  `ByteEngine.GetAutoBytecodeIndexShift() map[int]bool`.
- `engine/events.AstParsingEvent` - аналог `StringParsingEvent` для
  `engine.AstEngine`; `NewAstEngine` с `addDefaultEvents=true` теперь реально
  регистрирует его на `public.StringParseEvent` (в v1.6.0 аргумент принимался
  и игнорировался).
- `parser3`: константы фаз `PhaseLexer`, `PhaseStart`, `PhaseEnd`, `PhaseExpect`,
  `PhasePeek`, `PhaseEOF` и мета-ключей `MetaChildren`, `MetaOperator`,
  `MetaLeft`, `MetaRight`, `MetaOperand`; функции `AsParseError`,
  `AsGrammarError`, `ChildrenOf`.
- **Скорость и устойчивость `parser3`:** ограничение вложенности правил
  `maxRuleDepth = 256`, обнаружение леворекурсии, мемоизация результатов
  `NamedExpr` по паре (правило, позиция), отслеживание самой дальней ошибки
  для бэктрекинга, `fatal(err)` чтобы ошибки грамматики не глотались, и
  `maxPrattDepth = 512` для `PrattExpr`.
- **Лексер переведён с квадратичной на линейную сложность:**
  используется `regexp2.FindRunesMatchStartingAt` по общему срезу рун вместо
  `FindStringMatch` по свежей копии хвоста на каждой позиции; отклоняются
  совпадения нулевой длины и не привязанные к позиции, чтобы правило
  не могло зациклиться.
- `core.Events.GetEvents` и `CoreEvents` возвращают копии; `SetEvents`
  сохраняет копию; приватный `Events.isDebug()`.
- 14 новых файлов тестов (~1960 строк), `go test ./...` полностью зелёный:
  `end_test.go`, `astEngine_test.go`, `engine/astEngine_test.go`,
  `engine/core/errors_events_test.go`, `parsing/stringParsing/parser3/parser3_test.go`
  (45 тестов), `tooling/astools/main_test.go`, `tooling/bytecode/utils_test.go`,
  `tooling/plugin/realization_test.go`, `parsing/stringParsing/lexer_perf_test.go`,
  `parsing/stringParsing/bracket_test.go`, `parsing/stringParsing/parser1_option_test.go`.
- Новые примеры: полноценный интерпретатор языка `example/langs/listLang`
  (замыкания с лексической областью видимости, `if`/`else`, `while`,
  `for..in`, `break`/`continue`/`return`, встроенные `print`, `len`, `str`,
  `int`, `map`, `filter`, `reduce`), `example/langs/configParser` с
  `ActionExpr` для валидации значений прямо в ходе разбора, а также
  `example/packages/engine/engines/ast` и `example/packages/tooling/plugin`.
- Файл `.gitignore` (исключает `AGENTS.md`, `.gocache/`, `.gocache2/`).
- Поддержка беззнаковых конверсий в `bytecode.Utils`: `unsignedMaxValue`,
  `bytesToUint`, `uintToBytes` для корректной работы со знаковыми размерами.

### Changed

- Все пакеты снабжены обычными Go-документированными комментариями
  на английском языке; из кода убраны символы `GODOC`-стиля, em dash, en dash,
  неразрывные пробелы, типографские кавычки и многоточия.
- Переработаны `engine/core/events.go`, `errors.go`, `generator.go`,
  `universalEngineParams.go`, `engine/astEngine.go`, `engine/byteEngine.go`,
  `engine/stringEngine.go`, `engine/events/*`: унифицированы имена
  (`event` -> `eventHandler`, `val` -> `value`, параметры с `snake_case`
  на `camelCase`), устранены теневые переменные.
- `tooling/bytecode/utils.go` полностью переписан (216 строк изменений):
  устранено дублирование, добавлены тесты на int/float64/big-endian конверсии.
- `parsing/stringParsing/parser3`: `grammar.go` (+628), `parser.go` (+235),
  `errors.go` (+125), `formatter.go` (+173), новый `helpers.go`.
- `EngineUniversal.ProcessString` / `ProcessBytes` больше не заменяют
  настроенный context движка на `context.Background()`.
- `astools.Walk` переписан с рекурсии на явный итеративный DFS без ограничения
  глубины.

### Docs

- `README.md` переписан полностью (737 строк): разделы Scope, Architecture,
  Extension points, Engines, Core concepts, Plugins, Debugging and tooling,
  Context and cancellation, Performance, Examples. В нём разъясняется, что
  `AstEngine` не является третьим бэкендом, а реализует `StringEngineInterface`,
  и описывается практическое ограничение: конструкторы движков типизированы
  `ParsedNode` / `ParsedBytes`.
- Добавлен полный русский перевод `README-ru.md` (791 строка).
- Удалены 15 сгенерированных gomarkdoc файлов `GODOC.md` из пакетов
  и 12 из примеров; вместо них добавлены двуязычные README для
  `engine/core/`, `parsing/stringParsing/parser3/`, `tooling/plugin/`,
  `tooling/debugging/`, а также переписаны `parsing/README.md`,
  `parsing/stringParsing/README.md`, `tooling/README.md`, `example/README.md`.
  `tooling/debugging/README.md` честно описывает накладные расходы профилировщика.
- Удалены артефакты `backup/lexer.five` и `scan_results.txt`.
- Примеры переведены на новый API: `example/langs/calculator` и
  `example/langs/math` работают через `lc.NewEngineBuilder(...).WithStringParser(&parser3.Adapter{...})`
  и сообщают об ошибках через `parser3.AsParseError`.

## [1.6.0] - 2026-08-29

> **Внимание:** тег v1.6.0 не компилируется. Параметр `colorEnable` был удалён
> из `NewStringEngine` / `NewByteEngine`, но `builder.go` в этом коммите всё
> ещё его передаёт. Ошибка исправлена в следующем коммите `2481e0d "Fixes"`,
> который входит в состав v2.0.0. Используйте v2.0.0.

### Added

- **Движок AST** (`engine/astEngine.go`, 222 строки): `engine.AstEngine` с полями
  `AstCommandCtx`, `UEP`, `Parser`, `Commands`, `CanBeUnknown`,
  `CanMainNodeBeUnknown`; типы `AstCommandCtx` (`Name`, `Path`, `Parent`,
  `CurrentChildren`, `BreakIf`, `SkipIf`) и `AstCommandCtxPath`
  (`Path []string`, `Nodes []*ParsedNode`); конструкторы `MakeAstCommandCtxPath`,
  `AstMakeCommandCtx`; методы `Process`, `Work`, `WorkIter`, `HasCommand`,
  `GetCommandCtx`, `NewCommandFull`, `NewCommand`, `GetCommands`, `GetUep`, `GetParser`.
- `engine.AstEngineInterface` - алиас на `StringEngineInterface` (не новый тип).
- Конструктор `lc.NewAstEngine(resType, pipeline, addDefaultEvents, parser, context, canNodeBeUnknown, canMainNodeBeUnknown)`.
- Коды ошибок `AstEngineProcessError1/2`, `AstEngineHandlerError`,
  `AstEngineUnknown`, `EngineLifecycleEnd`.
- События `public.AstCallEvent`, `public.AstCallCalloopEvent`,
  `public.AstCommandCallEvent`.
- `core.LogerInterface` - интерфейс логгера, делающий логгер подменяемым.
- `Logger.SetStatusForm(status, form)` - программная установка формата статуса.
- `astools.WalkWithPath(node, fn)` - обход дерева с передачей пути до узла.

### Breaking

- Из `NewStringEngine` и `NewByteEngine` **удалён параметр** `colorEnable bool`.
- Удалена зависимость `github.com/pt-main/tap` v1.4.7 из `go.mod`;
  остаётся только `dlclark/regexp2 v1.12.0`. Вся работа с цветом из
  `main.go`, `engine/core/logger.go` и `parser3/formatter.go` удалена.
- Строковые значения имён событий стали префиксованными именем движка:
  `"INPUT string->..."` -> `"STRING:INPUT ..."`, `"BYTE:INPUT ..."`,
  `"STRINGCALLOP ..."` и другие.
- `UniversalEngineParams.Logger` меняет тип с `*Logger` на интерфейс `LogerInterface`.
- `parser3.FormatError(err, useColors bool)` игнорирует аргумент `useColors`
  и всегда возвращает цепочку без цвета.

### Changed

- `StringEngine.Process` и `AstEngine.Process` используют
  `core.GetRealErrorReverse(err)` при обёртывании ошибок разбора и вызова,
  чтобы сообщения читались от корневой причины.
- Формат логгера по умолчанию изменён с
  `"[?BE]%s[?RT] [?CN][%v][?RT] [?GN][%s][?RT]\n"` на `"%s [%v] [%s]\n"`.
- `astools.Walk` переписан с рекурсии на итеративный DFS в том же
  предпорядковом порядке, но без ограничения глубины.
- `plugin.NewPlugin` получает параметр `context context.Context`.
- `CallLoopData.Engine` и `CLEData.E` меняются с `*E` на `E`.

### Fixed

- Обработка `ErrExit` перенесена в `ByteCallEventIteration` (`*idx = -1`),
  что чинит выход из горячего цикла байтового движка.
- Порядок блокировок в `Events.NewEventBefore` исправлен; метод теперь
  возвращает ошибку, если событие отсутствует.
- `StringCallEvent` использует `core.ScopeGet` вместо непроверяемого
  приведения типа из scope.
- Удалены артефакты `backup/lexer.five`.


## [1.5.8] - 2026-08-28

### Added

- Новый флаг `public.ByteEngineScopeHotloopCtxCheckPeriod` (`"CTX_CKECK_PERIOD int"`)
  для настройки периодичности проверки контекста в горячем цикле байтового движка.
- `CallLoopData.Other` - новое поле для передачи дополнительных данных в хуки.
- Новый пример `example/langs/math` - калькулятор на parser3 + PrattExpr.

### Fixed

- Убраны вызовы `runtime.GC()` из горячего цикла (были причиной падения
  производительности с ~160m до ~120m ops/s).
- Авто-сдвиг индекса (`*idx++`) теперь выполняется **до** вызова обработчика,
  а `*idx` откатывается при ошибке - указатель не пропускает ошибочный opcode.
- Исправлен неправильный scope-ключ для чтения периода проверки контекста
  (читался `BytecodeIdx` вместо нового ключа).
- Исправлена двойная регистрация события в speedtest-бенчмарке.

### Changed

- Проверка контекста в горячем цикле выполняется раз в N итераций (по умолчанию 255)
  вместо каждой итерации.
- Условие выхода из цикла использует `uint` - трюк для обработки `idx == -1` без
  двойной проверки.

## [1.5.7] - 2026-08-25

### Breaking

- `EngineUniversal.ProcessString`, `ProcessStringWithCtx`, `ProcessBytes`,
  `ProcessBytesWithCtx`, `CheckEnded` и все `Process` методы движков
  возвращают `core.ErrorInterface` вместо `error`.
- Новые коды ошибок `public/errors.CorePackageLcError` (`"SYSTEM@LC"`) и
  `public/errors.CorePackageLcLifecycleError` (`"SYSTEM@LC:LIFECYCLE"`).

### Added

- `core.GetRealErrorReverse` - форматирование цепочки ошибок от корня к оболочке.
- `core.ErrorInterface` расширен методами `GetCode()`, `GetMsg()`, `GetMeta()`,
  `Unwrap()` для полноценной работы с `errors.Is` / `errors.As`.
- Типы ошибок parser3 (`ParseError`, `GrammarError`, `AdapterError`) приведены
  к тому же контракту.
- Экспортированные коды ошибок parser3: `ParseErrCode`, `GrammarErrCode`, `AdapterErrCode`.

### Fixed

- **HOTFIX (v1.5.7-f):** в `ByteEngine.Process` второй обработчик ошибки
  использовал `err1` вместо `err2` - байтовые ошибки вызова показывали
  ошибку парсинга вместо реальной причины.

## [1.5.6] - 2026-08-24

### Added

- `InstructionsGenerator.Generate` паникует при аргументе нулевой длины
  с явным предупреждением в документации.

### Fixed

- Восстановлена проверка нулевой длины аргумента в `byteParsing.Parser1.Parse`
  (была удалена в v1.5.5, теперь возвращает структурированную ошибку
  `"Zero argument length"` с метаданными `EMK(0,"int")`).
- `byteParsing.Parser1.Parse` сохраняет и восстанавливает `p.Config.Shifter.Idx`
  перед и после цикла - парсер можно использовать повторно и параллельно.

## [1.5.5] - 2026-08-22

### Fixed

- **Критично:** `byteParsing.Parser1.Parse` изменял общий указатель
  `p.Config.Shifter.Idx`. Теперь значение и сам указатель сохраняются
  перед циклом и восстанавливаются в `defer`, поэтому `Parser1` можно
  использовать повторно и параллельно, не затирая индекс вызывающего.
- Проверка логирования в `Parser1.Parse` защищена от `nil`-указателя
  `*ParseOption`.

### Changed

- `parser3.Adapter.Parse` возвращает `core.ErrorInterface`, завершая
  соответствие `ParserInterface`.

## [1.5.4] - 2026-08-19

### Added

- `core.GetErr` - разворачивает один уровень обертки; если причина не является
  `ErrorInterface`, синтезируется `Error{Code: WrappedError}`.
- Код ошибки `public/errors.WrappedError`.

### Changed

- `Error.Format` больше не добавляет завершающий разделитель `"----+"`.

## [1.5.3] - 2026-08-18

### Added

- `core.ErrorInterface` расширен методами `GetCode()`, `GetMsg()`,
  `GetMeta()` и `Unwrap()`, что позволяет использовать `errors.Is` / `errors.As`
  по всей цепочке lc-ошибок.
- Типы ошибок parser3 (`ParseError`, `GrammarError`, `AdapterError`)
  приведены к тому же контракту.
- Экспортированные коды ошибок parser3: `ParseErrCode`, `GrammarErrCode`,
  `AdapterErrCode`.
- `parser3.Parser.Parse` и `Expect` возвращают `core.ErrorInterface`,
  завершая соответствие `ParserInterface`.

## [1.5.2] - 2026-08-18

Крупный релиз структурированных ошибок и документации.

### Added

- Пакет `public/errors` с систематизацией кодов ошибок по доменам:
  `engines.go`, `events.go`, `generator.go`, `main.go`, `others.go`
  (около 25 констант `ErrorCodeType`).
- Тип `ErrorCodeType` и `ErrorMetaType`, функция `EMK(n, valType)`
  для генерации ключей метаданных.
- Файл `engine/core/errors.go`: `core.Error{Code, Msg, Meta, Cause}` с
  методами `Error()`, `Format()`, `WithMeta`; конструкторы `core.Err`,
  `core.Wrap`; помощники `GetMetaValue[T]`, `GetRealError`, `core.ErrExit`.
- Типы ошибок parser3: `ParseError`, `GrammarError`, `AdapterError` с
  полями `TokenIdx`, `TokenPos`, `Code`, `Expected`, `Got`, `Raw`,
  а также методами `Unwrap` и `Format`.
- Форматтеры ошибок `parser3.FormatError(err, useColors)` и
  `FormatErrorPretty`; из кода убраны вшитые ANSI-последовательности.
- `core.Events.SetProperty(name, value)` с поддержкой свойства `"debug"`:
  при выключенной отладке события начала и конца каждого вызова
  не создаются, что убирает накладные расходы.
- 27 файлов `GODOC.md`, сгенерированных gomarkdoc, для всех публичных пакетов.

### Breaking

- Почти все публичные API переходят с `error` на `core.ErrorInterface`:
  `core.ScopeGet`, все методы `Generator`, `NewUniversalEngineParams`,
  все методы `EventsInterface`, `parsing.ParserInterface.Parse`,
  все методы `bytecode.Shift` и все обработчики `DefaultEvents`.
- `core.EventType` и `core.CommandType` возвращают `ErrorInterface`.
- `public.ErrExit` удалён, заменён на `core.ErrExit` и код
  `public/errors.ErrExit`.
- `engine/core/config.go` переименован в `engine/core/types.go`.

### Fixed

- **Критично:** инвертированная проверка `canBeUnknown` в
  `StringCallEventIteration` - зарегистрированные команды пропускались,
  а незарегистрированные вызывались. Исправлено: теперь неизвестные
  команды пропускаются только при `StringEngineScopeCanBeUnknown = true`,
  иначе возвращается `DefaultEventsCallErrorUnknown`.
- `EventsTools.ChangeCoreEvent` при `idx == 0` заменял весь список обработчиков
  вместо точечной замены; `idx < 0` теперь возвращает ошибку.
- `parser3.engineAdapter` принимает `children` любого именованного типа среза
  через `reflect`, а не только точного `[]ParsedNode`.
- `byteParsing.Parser1` возвращает структурированные ошибки разбора
  с метаданными позиции, команды и номера аргумента.
- `Generator.GetBytesRes` / `GetStringArrRes` прикрепляют метаданные
  `EMK(0, "string")` с отсутствующей точкой пайплайна.
- Улучшены сообщения об ошибках лексера: позиция строки и столбца,
  контекстный фрагмент, корректная обработка незакрытых и лишних скобок.
- Каждый токен лексера теперь всегда несёт `__pos`, `__start` и `__end`
  в метаданных.

### Performance

- Лексер: предвычисленные индексы `ruleGroups` и бакеты `openByByte` /
  `closeByByte` для скобок устраняют линейные проходы по картам на
  каждой позиции.

## [1.5.0-pre2] - 2026-08-08

### Added

- `astools.FindChildIndex(node, switchName) int` - индекс первого ребёнка
  с заданным именем или -1.

### Fixed

- **Критично:** `parser3.ChoiceExpr.Parse` прекращал перебор после первой
  неудачной альтернативы. Теперь позиция сбрасывается и перебор
  продолжается.
- `parser3.Parser.Parse` передаёт `opts ...*parsing.ParseOption` в лексер
  (ранее опции терялись).
- `parser3.Adapter.Parse` корректно пробрасывает опции и выдаёт
  окрашенные ошибки.

### Changed

- Диагностика parser3 переработана: каждое выражение грамматики возвращает
  подробное сообщение с позицией токена, ожидаемым и полученным значением,
  числом альтернатив и списком невыбраненных токенов.
- В README добавлены блоки с выводом `go run ./example/readme/string`
  и `.../byte`.

## [1.5.0-pre] - 2026-08-04

Предрелизная версия 1.5.0, крупнейший архитектурный релиз до v2.0.0.

### Added

- `engine.EngineInterface` - общий интерфейс для обоих типов движков
  (строковый и байтовый). Типизированные алиасы: `StringEngineInterface`,
  `ByteEngineInterface`.
- `tooling/astools` - утилиты обхода AST: `GetChildren`, `FindChild`,
  `FindChildren`, `GetTokenValue`, `GetChildAt`, `Walk`.
- `tooling/debugging/profiler` - профилировщик производительности с
  метриками `Count`, `TotalTime`, `MinTime`, `MaxTime`, `Avg()` и
  методами `Report`, `Reset`, `Enable`, `Disable`.
- `extensiblePlugin.New` и полноценный `Close()` для восстановления
  исходных обработчиков; плагин регистрирует флаг `ECLFlag`.
- `core.Events.ReplaceEvent(name)` - полная замена события.
- Новая область `public.StringEngineScopeCanBeUnknown` - контроль
  поведения при неизвестных командах.
- Примеры `example/langs/calculator`, `example/langs/configLang`,
  `example/readme/string`, `example/readme/byte` и набор
  `example/packages/**`.
- Godoc-комментарии для `Lexer`, `Parser1`, `Parser2`, `PluginManager`,
  `plugin.Plugin`, `ProfilerPlugin`, `ExtensibleCLPlugin`.

### Breaking

- **Главное изменение:** все обработчики команд теперь принимают
  `engine.StringEngineInterface` / `engine.ByteEngineInterface` вместо
  `*engine.StringEngine` / `*engine.ByteEngine`. Это меняет сигнатуру
  каждого обработчика: `func(e *StringEngine, n *ParsedNode) error` ->
  `func(e StringEngineInterface, n *ParsedNode) error`.
- `EngineUniversal.StringEngine` и `.ByteEngine` становятся интерфейсными
  полями; добавлены аксессоры `GetUep()`, `GetParser()`, `GetCommands()`.
- `core.EventInput` переименован в `core.SimpleInput` (алиас сохранён).
- `NewCommand` принимает `*core.SimpleInput` и возвращает `error`;
  прежняя сигнатура сохранена в `NewCommandFull`.
- Флаг `engine.AutoshiftNewCommandFlag = "autoShift"` управляет
  авто-сдвигом индекса в байтовом движке.
- `plugin.Tools.Pm` изменен с значения на указатель; `PluginManager.Flags`
  стал приватным полем `flags`, доступ - только через `Tools`.
- `CallLoopData.Engine` и `CLEData.E` меняются с `*E` на `E`.

### Fixed

- Обработка `ErrExit` перенесена в `ByteCallEventIteration`
  (`*idx = -1`), что чинит выход из горячего цикла байтового движка.
- Порядок блокировок в `Events.NewEventBefore` исправлен; метод
  возвращает ошибку, если событие не найдено.
- `StringCallEvent` использует `core.ScopeGet` вместо непроверяемого
  приведения типа из scope.
- Неизвестные команды в строковом движке теперь приводят к ошибке
  вместо молчаливого пропуска.

## [1.4.6] - 2026-07-30

### Fixed

- Обработка `ErrExit` перенесена из `StringCallLoopEvent` в
  `StringCallEventIteration` - установка `*idx = -1` для немедленного
  завершения цикла.

### Changed

- Лексер поддерживает многобайтовые скобки: `LexerConfig.Brackets`
  изменён с `[]string` на `[][2]string` (явные пары
  открывающая/закрывающая).
- Внутренние карты скобок изменены с `map[rune]rune`
  на `map[string]string`.

## [1.4.5] - 2026-07-29

В репозитории версия помечена как 1.4.5, тег `v1.4.4` указывает
на предыдущий коммит.

### Added

- `public.ErrExit` - обработчики могут вернуть эту ошибку для
  управляемого выхода из цикла вызовов; циклы распознают её через
  `errors.Is` и завершаются штатно, без трассировки ошибки.
- Индекс инструкции публикуется в scope через
  `public.StringEngineScopeInstrIdx`.
- `public.ByteEngineScopeHotloopCtxCheckPeriod` - настройка частоты
  проверки отмены.

### Fixed

- **Критично:** проверка контекста в горячем цикле байтового движка
  стояла внутри условия `iter&4095 == 0`, из-за чего при коротких
  программах отмена не проверялась вообще. Теперь `ctx.Err()`
  проверяется в начале каждой итерации.
- `StringCallLoopEvent`: условие выхода `for *idx < pLen && *idx >= 0`
  упрощено до `for *idx < pLen` с явной проверкой `*idx < 0` внутри,
  так как отрицательный индекс используется как сигнал выхода.

## [1.4.3] - 2026-07-28

### Breaking

- `parser3.Adapter.ParseFlat` переименован в `parser3.Adapter.Parse`,
  чтобы соответствовать `parsing.ParserInterface`.

### Fixed

- `StringCallEvent`: в сообщении о панике использовался неверный
  спецификатор `%e` вместо `%v`, из-за чего значение паники
  выводилось некорректно.
- `StringCallEvent`: в лог писалось "end parsing event" вместо
  "end call event".
- Исправлен маркер контекста в выводе ошибки: `[?BBK]    |`
  -> `[?BBK]>    |`.

## [1.4.2] - 2026-07-21

### Added

- `EngineUniversal.End()` - явное завершение жизненного цикла движка.
- Поле `EngineUniversal.CtxCancelCause context.CancelCauseFunc` и
  builder-метод `WithContext(ctx)`, дающие корректную отмену
  с указанием причины.
- `EngineUniversal.CheckEnded()` - проверка состояния перед
  каждой операцией; вызывается в `ProcessString`, `ProcessBytes`,
  `GetUEP`, `NewCommandByte`, `NewCommandString`.

### Changed

- **Breaking:** `GetUEP()` теперь возвращает
  `(*core.UniversalEngineParams, error)` вместо простого указателя.
- `End()` отменяет context, обрабатывает паники и возвращает ошибку
  плагинов напрямую, а не через строковое форматирование.

## [1.4.1] - 2026-07-19

### Added

- Инфраструктура замеров производительности:
  `example/speedtest/tests/main.go` (набор тестов),
  `example/speedtest/byte_test.go` (бенчмарк), speedtest переведён
  из обычной программы в `testing.B`.

### Changed

- `builder.WithEndianess` принимает `public.EndianType` вместо `int`
  (вместе с типизацией `NewEngineBuilder`, введённой в 1.3.2).

## [1.4.0] - 2026-07-17

### Added

- Пакет `tooling/debugging/extensiblePlugin` - плагин, расширяющий цикл
  обработки команд: хуки `CLEPreEvent`, `CLEInPreEvent`,
  `CLEInPostEvent` (вызываются по одному разу вокруг цикла)
  и данные обхода `SCLEData` / `BCLEData` в scope.
- `tooling/plugin/tools.go` с типом `Tools` - доступ к флагам
  менеджера плагинов.
- `core.EventsInterface` - интерфейс для системы событий.
- `public/logging.go` с именами статусов логирования: `LogEvents`,
  `LogParsing`, `LogVerbose`.
- `core.Events.NewEventBefore` для регистрации обработчика в начало списка;
  `EventsTools` с `ChangeCoreEvent`, `GetCoreEvent`, `GetCoreEventIdx`.
- Событие `public.ByteCallHotloopEvent` и `public.CLEScopeData`:
  отдельный горячий цикл для байтового движка.

### Changed

- События переименованы: строки приведены к виду
  `"INPUT string->PARSED []ParsedNode"`, `"call(PARSED []ParsedNode)"`,
  `"CALLOP call(PARSED []ParsedNode)"`, `"HOTLOOP call(PARSED []ParsedBytes)"`.
- Расширена система логирования: `MaxLogLength` в логгере.

## [1.3.9] - 2026-07-12

### Added

- `example/speedtest/byte.go` - программа замеров байтового движка.
- Тип `ByteCallAttr` (`RawIdx`, `RawNode`, `Abis`, `Handler`) -
  предрассчитанные атрибуты команды для горячего цикла.

### Changed

- `ByteCallEventIteration` переработан: сигнатура
  `(idx *int, parsed ByteCallAttr, e *engine.ByteEngine) error`
  вместо разбора узла внутри каждой итерации. Разбор opcode,
  поиск обработчика и чтение `AutoBytecodeIndexShift` вынесены
  из цикла, в обработчик передаётся готовый `Handler`.
- `ByteParsingEvent` передаёт `*parsing.ParseOption` в парсер вместо
  вызова `Parse(input)` без опций.

## [1.3.7] - 2026-07-06

### Fixed

- `ByteCallEvent` больше не разыменовывает указатель `idx` в
  `defer` до его инициализации: `last_cmd_switch` и `idx` вынесены
  выше `defer`, отложенный обработчик ошибки теперь безопасен.
- `ByteCallEventIteration` возвращает ошибку вместо
  `nil`, поэтому ошибка обработчика не теряется.

## [1.3.6] - 2026-06-29

### Fixed

- **Критично:** в `StringCallEvent` ошибочный `break` после обработки
  первой команды останавливал цикл. Теперь все команды строки
  выполняются, а цикл прерывается только при ошибке.
- `recover` перенесён из тела цикла (где `defer` накапливался
  на каждой итерации) в единственный `defer` на всё событие.
- `StringCallEvent` и `ByteCallEvent` печатают позицию ошибки
  (исходную строку и индекс команды) прямо в отложенном обработчике,
  поэтому паника в обработчике тоже даёт осмысленное сообщение.

## [1.3.5] - 2026-06-24

### Added

- `parsing/stringParsing/parser3/engineAdapter.go` - адаптер parser3
  к `parsing.ParserInterface`.
- `tooling/plugin/realization.go` - реализация `PluginInterface`.

### Changed

- `tooling/plugin/core.go` переименован в `tooling/plugin/manager.go`.
- Переработан лексер: -196 строк, упрощена работа со скобками.
- `Parser1`, `Parser2` и grammar parser3 переведены на `*parsing.ParseOption`.

## [1.3.3] - 2026-06-23

### Fixed

- `StringCallEvent`: `defer func() { recover() }` больше не создаётся
  на каждой итерации цикла (рост памяти и работа вхолостую),
  `recover` вынесен в один отложенный обработчик на всё событие.
- Ошибка обработчика и паника теперь прерывают цикл команд
  с понятным сообщением вместо молчаливого продолжения.

### Changed

- `byteParsing` использует общий тип `parsing.ParserInterface`.

## [1.3.2] - 2026-06-23

Крупный релиз типизированных scope, enum-констант и системы плагинов.

### Added

- Пакет `public` - общие для движков константы и типы.
- `public/types.go`: `ResType` (`ByteResType`, `StringResType`),
  `EndianType` (`BigEndian`, `LittleEndian`), `EngineType`
  (`ByteEngineType`, `StringEngineType`).
- `public/scope.go`: все строки scope движков, событий и плагинов
  собраны в константы (`ByteEngineScopeParsed`, `StringEngineScopeParsed`,
  `EventsScopeCallName`, `PluginsScopeEuPtr` и другие).
- `public/events.go` - константы имён событий.
- `engine/core/scope.go` с generic-хелпером `core.ScopeGet[T]`
  для типобезопасного чтения scope.
- `tooling/plugin/interface.go` с `PluginInterface` и полным
  контрактом `Name/Init/Close/Call/Run`.

### Changed

- **Breaking:** `NewEngineBuilder(engineType int)` ->
  `NewEngineBuilder(engineType public.EngineType, resType public.ResType)`;
  тип результата генератора стал явным параметром билдера.
- **Breaking:** `WithEndianess(endianess int)` ->
  `WithEndianess(endianess public.EndianType)`.
- **Breaking:** `WithPluginManager(plugins ...*plugin.Plugin)` ->
  `WithPlugins(plugins ...plugin.PluginInterface)`.
- Обработчики default events читают вход и разобранные узлы
  через `core.ScopeGet` вместо прямого приведения типа из scope.
- `ByteCallEventIteration` выделен в отдельный метод -
  вынесен шаг обхода байткода.

## [1.2.0] - 2026-06-20

### Added

- **Новый пакет `parsing/stringParsing/parser3`** - генераторный
  парсер со встроенным AST: типы `Grammar`, `Rule`, `Expr`
  (интерфейс), `TokenExpr`, `SequenceExpr`, `ChoiceExpr`,
  `RepeatExpr`, `OptionalExpr`, `NamedExpr`, `NodeExpr`;
  конструктор `NewParser(lexer, grammar, startRule, ignoreTypes)`;
  методы `Expect`, `Peek`, `skipIgnored`, `Errorf`.
- `parsing/main.go`: тип `ParseOption{UEP, Flags, Other}` и generic
  `ParserInterface[I, P] { Parse(I, ...*ParseOption) ([]P, error); String() string }`.
- `example/parser3.go` - калькулятор с выводом дерева в JSON.
- `core.Events.NewEventBefore` - регистрация обработчика
  в начало списка событий.

### Changed

- **Breaking:** сигнатура `Parse(I, ...interface{})` заменена
  на `Parse(I, ...*ParseOption)`.
- **Breaking:** лексер стал блок-ориентированным: появились
  `LexerConfig.UseLineContinuation`, `SkipEmptyLines`,
  `TrimBlocksSpace`, а токены получают метаданные `__block_index`.
- **Breaking:** `tooling/plugin` переработан на интерфейсы:
  `Plugin.Name` из поля стал методом, добавлен конструктор
  `NewPlugin(name, initEvent, mainEvent, closeEvent)`,
  методы `InitPlugin` / `ClosePlugin` / `RunPlugin` переименованы
  в `Init` / `Close` / `Run`, появился `PluginInterface`.
- **Breaking:** `PluginManager.Plugins` теперь
  `map[string]PluginInterface`; `DeletePlugin` и `GetPlugin`
  возвращают ошибку, если плагин не найден.

### Fixed

- Логирование разбора теперь идёт через `*parsing.ParseOption`
  (раньше `core.UniversalEngineParams` передавался по значению,
  из-за чего логгер внутри парсера не срабатывал).

### Removed

- Примеры `example/configLang1` удалены, `example/configLang2`
  остаётся как эталон.

## [1.1.5] - 2026-06-19

### Added

- `stringParsing.LexerConfig` с `UseBracketBalance` и `Brackets []string`;
  второй аргумент `NewLexer(rules, config *LexerConfig)`.
- Балансировка скобок в лексере: метод `isBracketBalanced`,
  результат пишется в метаданные токена `__bracket_balanced`.
- Метод `String() string` в `parsing.ParserInterface`; реализации
  `Parser1`, `Parser2` и `byteParsing.Parser1` возвращают свои пути.

### Changed

- `stringParsing.ParsedNode.Metadata` переведён с
  `map[string]interface{}` на `core.ScopeType`.
- `Lexer.Parse` принимает `i ...interface{}`.

## [1.1.2] - 2026-06-18

### Added

- **Реорганизация пакетов:** `stringParsing/` -> `parsing/stringParsing/`,
  `byteParsing/` -> `parsing/byteParsing/`, `events/` -> `engine/events/`.
- `parsing/interface.go` с generic `parsing.ParserInterface[I, P]`
  (`Parse(I, ...interface{}) ([]P, error)`); дубликаты интерфейсов
  в `byteParsing` и `stringParsing` удалены.
- **Новый пакет `tooling/plugin`:** тип `Plugin`, менеджер
  `PluginManager` с `AddPlugin`, `DeletePlugin`, `GetPlugin`,
  `CallPlugin`, события `Init` / `Close` / `Main`; билдер получает
  `WithPluginManager(plugins ...*plugin.Plugin)`.
- `EngineUniversal` получает поле `Plugins`; константы
  `EuPtrPluginsScope` и `PluginManagerEuScope = "LC-PM"`.
- Примеры `example/configLang1` (собственный парсер) и
  `example/configLang2` (на `stringParsing.Parser1`).
- Лицензия Apache 2.0, шаблоны issue для GitHub,
  `example/README.md`.

### Breaking

- **Удалена динамическая загрузка библиотек:** удалены
  `plugin.Open`-механизм, тип `NewPluginFunction`, метод
  `EngineUniversal.LoadPluginFromFile(path)`, константа
  `pluginFileSymbolName` и импорт `plugin`. Плагины регистрируются
  только программно через `PluginManager`.
- Старые `ParserInterface` в `byteParsing` и `stringParsing` удалены
  в пользу общего `parsing.ParserInterface`.

### Fixed

- `ByteCallEvent` теперь выдаёт ошибку на неизвестный opcode вместо
  молчаливого пропуска.
- `NewLogger` инициализирует карту `Logging`, без чего обращение
  к логгеру падало.
- `Logger.Logging map[string]bool`: вывод в stdout только для
  разрешённых статусов, в `Log` пишутся строки без ANSI-последовательностей.
- Исправлена опечатка в scope-ключе `BYECODE_IDX` -> `BYTECODE_IDX`.
- `EngineBuilder.Build` добавляет префикс `EngineBuilder.Build:`
  к сообщениям об ошибках.

## [1.0.0] - 2026-06-17

Первый стабильный релиз.

### Added

- **Два движка:** `StringEngine` (текстовые команды, ключи scope
  `INPUT string` и `PARSED []ParsedNode`) и `ByteEngine` (команды
  по опкодам, ключи `ENDIANESS int`, `BYTECODE_IDX *int`,
  `INPUT []byte`, а также `AddToBytecodeIdx`, `SetBytecodeIdx`,
  `GetBytecodeIdx` и флаг `AutoBytecodeIndexShift`).
- **`UniversalEngineParams`** - общий контейнер для обработчиков:
  `Generator`, `Event`, `Scope`, `Logger`, `Context`; метод `GetContext()`.
- **Система событий** - `Events` на основе `orderedmap` с методами
  `NewEvent`, `GetEvents`, `CallEvents`.
- **Пакет `events`** с `DefaultEvents`: `StringParsingEvent`,
  `StringCallEvent`, `ByteParsingEvent`, `ByteCallEvent`.
- **`Generator`** - накопление кода в именованных точках пайплайна
  с методами `AddString`, `AddBytes`, `GetStringArrRes`, `GetBytesRes`.
- **Парсеры:** `stringParsing` (`Parser1` на правилах грамматики,
  `Parser2` построчный, `Lexer` на регулярных выражениях
  с балансировкой скобок, `ParsedNode` с метаданными и связями
  `__prev` / `__next`), а также `byteParsing` (`Parser1` с
  `Parser1Config`, `ShiftStruct`, `ParsedBytes`).
- **`tooling/bytecode`** - `Utils` для преобразования int и float64
  в байты и обратно с учётом порядка байтов, `InstructionsGenerator`,
  `GenerationConfig`, константы `BigEndian` / `LittleEndian`.
- **`Logger`** - потокобезопасный логгер с методами `PrintLog`,
  `GetLog`, `Statuses` и конструктором `NewLogger`.
- **`EngineBuilder`** - паттерн Builder с методами `WithPipeline`,
  `WithContext`, `WithDefaultEvents`, `WithLogger`, `WithScope`,
  `WithColors`, `WithStringParser`, `WithByteParser`, `WithEndianess`,
  `Build`.
- **`EngineUniversal`** - оболочка с `ProcessString`, `ProcessBytes`,
  `ProcessStringWithCtx`, `ProcessBytesWithCtx`, `GetUEP`,
  `NewCommandString`, `NewCommandByte` (автоназначение опкода
  при передаче `-1`).
- Отмена через `context.Context` проверяется в цикле диспетчеризации
  команд всех движков.
- Тесты: `byteParsing/parser1_test.go`, `stringParsing/lexer_test.go`,
  `engine/core/generator_test.go`.
- Зависимости: `iancoleman/orderedmap v0.3.0`, `dlclark/regexp2 v1.12.0`,
  `pt-main/tap v1.1.1`.

### Changed

- **Breaking:** `core.CommandType` и `CommandMeta` стали generic
  (`CommandType[E, N any] func(*E, N) error`).
- **Breaking:** `core.EventType` теперь `func(interface{}, *Events) error`.
- **Breaking:** `GetUEP()` возвращает `*core.UniversalEngineParams`,
  поле `UEP` в движках стало указателем.
- **Breaking:** из `EngineUniversal` удалено поле `another engineAnother`,
  счётчик опкодов стал `opcode_counter`.
- Цветной вывод через `pt-main/tap` (`color.Set`, `color.ColorEnabled`),
  цветные коды `[?RD]` и `[?YW]` в сообщениях об ошибках.

### Fixed

- `StringEngine` и `ByteEngine` защищены `sync.RWMutex` - команды
  можно регистрировать и читать конкурентно.
- `NewUniversalEngineParams` возвращает ошибку на `nil`-параметры.
- Исправлен `logE`, который всегда падал на неверном типе `error`.
- Убран лишний вывод в byte call event.

### Removed

- Файл `engine/converts.go` с типом `Converts` - заменён
  generic-типами.

## [0.10.1] - 2026-06-14

### Changed

- `tooling/bytecode`: поля `InstructionsGenerator` (`OpcodeLen`, `ArglenLen`,
  `ArgscountLen`, `Endianess`) заменены на единую структуру
  `bytecode.GenerationConfig`.
- `byteParsing.Parser1Config` больше не хранит длины и порядок байтов
  по отдельности: вместо них введено поле `GConfig bytecode.GenerationConfig`
  (поле `Shifter` сохранено).

## [0.9.15] - 2026-06-09

Промежуточный выпуск без изменений публичного API: правки в
`engine/byteEngine.go`, `engine/core/events.go`, `events/byteEngine.go`
и `tooling/bytecode/utils.go`.

## [0.9.12] - 2026-06-08

### Added

- Конвертация `float64` в `tooling/bytecode`: `Utils.Float64ToBytes`,
  `BytesToFloat64`, `Float64ToBytesRange`, `BytesToFloat64Range`,
  а также варианты `Float64ToBytesBigEndian` / `Float64ToBytesLittleEndian`
  и `BytesToFloat64BigEndian` / `BytesToFloat64LittleEndian`.
- Методы шифта `ShiftFloat64Error`, `ShiftFloat64Panic`,
  `ShiftFloat64RangeError`, `ShiftFloat64RangePanic`.
- `ByteEngine.GetBytecodeIdx() (*int, error)`.
- Документация к `NewCommandByte` о том, что обработчик обязан сам
  двигать индекс байткода.

### Changed

- `SetBytecodeIdx` теперь кладёт в scope **указатель** `&n`, а не значение,
  чтобы согласовать поведение с `AddToBytecodeIdx`.

## [0.9.11] - 2026-06-08

### Added

- Константы scope байтового движка: `ByteEngineScopeEndianess`,
  `ByteEngineScopeBytecodeIdx`, `ByteEngineScopeInput`.
- Методы `ByteEngine.AddToBytecodeIdx(n int)` и `SetBytecodeIdx(n int)`.

### Changed

- Строки событий потеряли нижние подчёркивания:
  `"input string->parsed []ParsedNode"` и другие.
- `ByteCallEvent` обходит байткод по индексу из scope, а не через `range`,
  - появился пошаговый обход с возможностью досылки команд обработчиком.

## [0.9.9] - 2026-06-08

### Changed

- `bytecode.Shift.code` экспортирован в `bytecode.Shift.Code`.
- `Utils.ShiftStruct(code, idx)` заменён конструктором
  `bytecode.NewShift(code []byte, idx *int) *Shift`.

### Fixed

- `byteParsing.Parser1.Parse` присваивает `p.Config.Shifter.Code = code`.

## [0.9.8] - 2026-06-08

### Changed

- Тип `engineUniversal` экспортирован как `engine.EngineUniversal`,
  сигнатура `Build()` изменена на `(*EngineUniversal, error)`.
- `byteParsing.Parser1Config` получает поле `Shifter bytecode.Shift`:
  шифт вынесен наружу, `Utils.ShiftStruct` в парсере больше не используется.

## [0.9.7] - 2026-06-07

### Changed

- Массовая документация godoc на `EngineBuilder`, `EngineUniversal`,
  `ByteEngine`, `Lexer`, `LexerRule`, `NewLexer`, `Parser1`, `NewParser1`,
  `ParsedNode`, `ParsedBytes`, `Events`, `NewEvents`, `ProcessString`,
  `ProcessBytes`, `NewCommandByte`, `NewCommandString`.
- `Lexer.Parse` кладёт в `__raw` остаток исходника, а не сам токен.

## [0.9.1] - 2026-06-06

### Added

- Корневые файлы `builder.go` и `engine.go` (пакет `lc`).
- `lc.NewEngineBuilder(engineType int)` с флюент-методами `WithPipeline`,
  `WithDefaultEvents`, `WithLogger`, `WithScope`, `WithStringParser`,
  `WithByteParser`, `WithEndianess` и `Build()`.
- Константы `lc.ByteEngineType` и `lc.StringEngineType`.
- Тип `engineUniversal` с методами `ProcessString`, `ProcessBytes`, `GetUEP`,
  `NewCommandByte`, `NewCommandString`; `opcode = -1` означает
  авто-инкремент.

### Changed

- `CallEventsEvent` заменён на пару `CallEventsStartEvent` /
  `CallEventsEndEvent`; `Events.CallEvents` пишет в scope `call_name`
  и `call_error`.
- Лексер кладёт в метаданные ключи `__raw` и `__value`.
- `Parser2.Parse` возвращает `addPrevNextNodes(result)`.

## [0.8.8] - 2026-06-06

Предрелизная версия с крупной перестройкой внутренних пакетов.

### Added

- Пакет `tooling/bytecode` (на основе бывшего `system/utils.go`):
  тип `Utils`, конвертация `IntToBytes` / `BytesToInt` с вариантами
  для обоих порядков байтов, константы `BigEndian` / `LittleEndian`,
  `ShiftStruct` с `ShiftError` и `ShiftPanic`, а также
  `InstructionsGenerator` с методом `Generate`.
- `engine/core/logger.go` с типом `Logger` и методами `GetStatusForm`,
  `PrintLog`, `GetLog`.
- `engine/core/universalEngineParams.go` с `UniversalEngineParams` и
  конструктором `NewUniversalEngineParams`.
- `stringParsing/utils.go` с `addPrevNextNodes` - линковкой узлов
  через метаданные `__prev` / `__next`.

### Changed

- `system/config.go`, `system/events.go`, `system/generator.go`
  переехали в `system/core/`.
- `StringEngine` и `ByteEngine` больше не хранят `Scope`, `Generator`
  и `Event` по отдельности - всё собрано в поле `UEP`.
- Ошибки `Process` оборачиваются с указанием имени события.

## [0.6.3p] - 2026-06-05

### Added

- `Converts.ConvertByteCommandTypeArgs(args []interface{})` -
  конвертер аргументов обработчика байтовой команды.

## [0.6.3] - 2026-06-05

### Changed

- Константы событий `ParseEvent` / `CallEvent` переименованы
  в `StringParseEvent` / `StringCallEvent`; добавлены `ByteParseEvent`
  и `ByteCallEvent`.
- `Events.CallEvents` получает параметр `canWorkWithoutHandler bool`
  и возвращает ошибку вместо молчаливого `nil`.

### Fixed

- `ConvertStringCommandTypeArgs` берёт узел из `args[1]`, а не `args[0]`.

## [0.6.2] - 2026-06-05

Первый релиз с поддержкой байтового движка.

### Added

- Пакет `byteParsing`: `ParserInterface` с методом
  `Parse(code []byte) ([]ParsedBytes, error)`, тип `ParsedBytes`
  (`Switch`, `Raw`, `Args`, `Metadata`), `Parser1Config`
  (`CommandBytelen`, `ArglenBytelen`, `ArgscountBytelen`), `Parser1`
  и `Utils`.
- `system/byteEngine.go` с `ByteEngine` и `events/byteEngine.go`
  с байтовыми событиями.
- `system/converts.go` с типом `Converts`.

### Changed

- Пакет `parsing` переименован в `stringParsing`, `system/engine.go`
  в `system/stringEngine.go`, `events/engine.go` в `events/stringEngine.go`.
- `CommandType` стал `func([]interface{}) error`, `EventType` -
  `func(interface{}) error` (раньше оба были привязаны к `*Engine`).

## [0.2.5] - 2026-06-03

### Changed

- **Breaking:** лексер переведён с движка `regexp` на
  `github.com/dlclark/regexp2` v1.12.0 с поддержкой lookaround
  и обратных ссылок.
- **Breaking:** `LexerRule.Pattern` теперь `*regexp2.Regexp`, извлечение
  именованных групп переписано на `GetGroupNames()` / `GroupByName`.

### Renamed

- `parsing.LexerParser` -> `parsing.Lexer`,
  `parsing.NewLexerParser` -> `parsing.NewLexer`.
- `parsing.NewParser` -> `parsing.NewParser1`.

## [0.2.1] - 2026-06-02

### Added

- Лексер `parsing.LexerParser` и конструктор `parsing.NewLexerParser`;
  тип `LexerRule{Type string, Pattern *regexp.Regexp}`.
- Поля `ParserConfig.SkipEmptyLines` и `TrimBlocksSpace`.

### Fixed

- `Parser1.matchGrammar` пропускает пустой блок только при
  `SkipEmptyLines`, а `TrimSpace` применяет только при
  `TrimBlocksSpace`; в `__raw` хранится исходная, не обрезанная строка.
- В `DefaultEvents.CallEvent` заглушка `err := errors.New("")`
  заменена на `var err error = nil`.

## [0.2.0] - 2026-06-02

### Fixed

- `DefaultEvents.CallEvent` больше не падает на неизвестной команде:
  добавлена проверка наличия записи в `e.Commands`.

## [0.1.5] - 2026-05-30

### Added

- `parsing.ParserInterface` с методом
  `Parse(code string) ([]ParsedNode, error)`.
- Построчный парсер `parsing.Parser2`, пишущий в метаданные ключи
  `command`, `args`, `__raw`.

### Changed

- **Breaking:** сигнатура `NewEngine` расширена параметрами
  `add_default_events bool` и `parser parsing.ParserInterface`.
- `system.Engine` получает поле `Parser`.
- `parsing.Parser` переименован в `parsing.Parser1`.

### Fixed

- `NewGenerator` инициализирует `code[point]` для каждого элемента
  пайплайна.

## [0.1.1] - 2026-05-30

Первый публичный релиз.

### Added

- Корневой пакет `lc` с константой `Version` и
  `NewEngine(generator_res_type int, pipeline []string) *system.Engine`.
- Базовые типы: `system.Engine`, `system.ScopeType`, `system.CommandType`,
  `system.CommandMeta`, `system.EventType`, `system.Events` на основе
  `orderedmap`, `system.Generator`, `parsing.ParsedNode`,
  `events.DefaultEvents` с `ParsingEvent` и `CallEvent`.
- `NewEvents`, `NewEvent`, `CallEvents`, `GetEvents`;
  `AddString`, `AddStrings`, `AddBytes`, `GetBytesRes`, `GetStringArrRes`,
  `GetStringRes` у генератора.

### Changed

- `NewEngine` перенесён из пакета `system` в корневой пакет `lc`.

---

## Типы изменений

- **Added** - новая функциональность.
- **Changed** - изменение существующей функциональности.
- **Deprecated** - устаревшая функциональность.
- **Removed** - удаленная функциональность.
- **Fixed** - исправление ошибок.
- **Security** - исправления уязвимостей.
- **Breaking** - ломающие изменения API (выделены в v1.5.0, v1.5.2, v1.5.7, v1.6.0, v2.0.0).
- **Performance** - оптимизации производительности.

## Ссылки

- [GitHub](https://github.com/pt-main/lc)
- [Go Reference](https://pkg.go.dev/github.com/pt-main/lc)
- [Project Wiki](https://github.com/pt-main/lc/wiki)
