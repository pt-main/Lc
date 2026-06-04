package system

import "github.com/pt-main/lc/stringParsing"

type StringEngine struct {
	Scope     ScopeType
	Commands  map[string]CommandMeta
	Generator *Generator
	Event     *Events
	Parser    stringParsing.ParserInterface
}

func (e *StringEngine) Process(input string) error {
	e.Scope["input_string"] = input
	err1 := e.Event.CallEvents(e, StringParseEvent, false)
	if err1 != nil {
		return err1
	}
	err2 := e.Event.CallEvents(e, StringCallEvent, false)
	if err2 != nil {
		return err2
	}
	return nil
}

func (e *StringEngine) NewCommand(cmd_switch string, handler CommandType, doc string) {
	e.Commands[cmd_switch] = CommandMeta{
		Handler: handler,
		Doc:     doc,
	}
}
