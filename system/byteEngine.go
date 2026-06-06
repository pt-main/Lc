package system

import (
	"fmt"

	"github.com/pt-main/lc/byteParsing"
	"github.com/pt-main/lc/system/core"
)

type ByteEngine struct {
	Commands map[int]core.CommandMeta
	Parser   byteParsing.ParserInterface
	UEP      core.UniversalEngineParams
}

func (e *ByteEngine) Process(input []byte) error {
	e.UEP.Scope["input_[]byte"] = input
	err1 := e.UEP.Event.CallEvents(e, core.ByteParseEvent, false)
	if err1 != nil {
		return fmt.Errorf("Calling 'ByteParseEvent' event error: %v", err1.Error())
	}
	err2 := e.UEP.Event.CallEvents(e, core.ByteCallEvent, false)
	if err2 != nil {
		return fmt.Errorf("Calling 'ByteCallEvent' event error: %v", err2.Error())
	}
	return nil
}

func (e *ByteEngine) NewCommand(cmd_switch int, handler core.CommandType, doc string) {
	e.Commands[cmd_switch] = core.CommandMeta{
		Handler: handler,
		Doc:     doc,
	}
}
