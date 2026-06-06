package engine

import (
	"fmt"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/stringParsing"
)

type StringEngine struct {
	Commands map[string]core.CommandMeta
	Parser   stringParsing.ParserInterface
	UEP      core.UniversalEngineParams
}

func (e *StringEngine) Process(input string) error {
	e.UEP.Scope["input_string"] = input
	err1 := e.UEP.Event.CallEvents(e, core.StringParseEvent, false)
	if err1 != nil {
		return fmt.Errorf("Calling 'StringParseEvent' event error: %v", err1.Error())
	}
	err2 := e.UEP.Event.CallEvents(e, core.StringCallEvent, false)
	if err2 != nil {
		return fmt.Errorf("Calling 'StringCallEvent' event error: %v", err2.Error())
	}
	return nil
}

func (e *StringEngine) NewCommand(cmd_switch string, handler core.CommandType, doc string) {
	e.Commands[cmd_switch] = core.CommandMeta{
		Handler: handler,
		Doc:     doc,
	}
}
