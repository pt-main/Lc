package system

import "github.com/pt-main/lc/byteParsing"

type ByteEngine struct {
	Scope     ScopeType
	Commands  map[int]CommandMeta
	Generator *Generator
	Event     *Events
	Parser    byteParsing.ParserInterface
}

func (e *ByteEngine) Process(input []byte) error {
	e.Scope["input_[]byte"] = input
	err1 := e.Event.CallEvents(e, ByteParseEvent, false)
	if err1 != nil {
		return err1
	}
	err2 := e.Event.CallEvents(e, ByteCallEvent, false)
	if err2 != nil {
		return err2
	}
	return nil
}

func (e *ByteEngine) NewCommand(cmd_switch int, handler CommandType, doc string) {
	e.Commands[cmd_switch] = CommandMeta{
		Handler: handler,
		Doc:     doc,
	}
}
