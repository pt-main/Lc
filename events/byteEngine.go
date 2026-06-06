package events

import (
	"errors"
	"fmt"

	"github.com/pt-main/lc/byteParsing"
	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/tooling/bytecode"
)

func (de *DefaultEvents) ByteParsingEvent(_e interface{}) error {
	e, ok := _e.(*engine.ByteEngine)
	if !ok {
		return fmt.Errorf("Can't get byte engine: invalid iput")
	}
	input, ok := e.UEP.Scope["input_[]byte"].([]byte)
	if !ok {
		return errors.New("no input in scope")
	}
	nodes, err := e.Parser.Parse(input)
	if err != nil {
		return err
	}
	e.UEP.Scope["parsed_[]ParsedBytes"] = nodes
	return nil
}

func (de *DefaultEvents) ByteCallEvent(_e interface{}) error {
	e, ok := _e.(*engine.ByteEngine)
	if !ok {
		return fmt.Errorf("Can't get byte engine: invalid iput")
	}
	_parsed, _ := e.UEP.Scope["parsed_[]ParsedBytes"]
	parsed, ok := _parsed.([]byteParsing.ParsedBytes)
	if !ok {
		return errors.New("Can't start call event. Invalid type of parsed result.")
	}
	u := bytecode.Utils{}
	endianess, ok := e.UEP.Scope["endianess"].(int)
	if !ok {
		return fmt.Errorf("Can't get endianess: not declarated in scope")
	}
	for _, node := range parsed {
		cmd_switch := u.BytesToInt(node.Switch, endianess)
		handler, ok := e.Commands[cmd_switch]
		var err error = nil
		if ok {
			err = handler.Handler([]interface{}{e, node})
		}
		if err != nil {
			return errors.New("Handler error at: " + err.Error())
		}
	}
	return nil
}
