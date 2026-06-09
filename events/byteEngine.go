package events

import (
	"errors"
	"fmt"

	"github.com/pt-main/lc/byteParsing"
	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/tooling/bytecode"
)

const ByteEngineScopeParsed = "parsed []ParsedBytes"

func (de *DefaultEvents) ByteParsingEvent(_e interface{}) error {
	e, ok := _e.(*engine.ByteEngine)
	if !ok {
		return fmt.Errorf("Can't get byte engine: invalid iput")
	}
	input, ok := e.UEP.Scope[engine.ByteEngineScopeInput].([]byte)
	if !ok {
		return errors.New("No input in scope")
	}
	nodes, err := e.Parser.Parse(input)
	if err != nil {
		return err
	}
	e.UEP.Scope[ByteEngineScopeParsed] = nodes
	return nil
}

func (de *DefaultEvents) ByteCallEvent(_e interface{}) error {
	e, ok := _e.(*engine.ByteEngine)
	if !ok {
		return fmt.Errorf("Can't get byte engine: invalid input")
	}
	_parsed, _ := e.UEP.Scope[ByteEngineScopeParsed]
	parsed, ok := _parsed.([]byteParsing.ParsedBytes)
	if !ok {
		return errors.New("Can't start call event. Invalid type of parsed result.")
	}
	u := bytecode.Utils{}
	endianess, ok := e.UEP.Scope[engine.ByteEngineScopeEndianess].(int)
	if !ok {
		return fmt.Errorf("Can't get endianess: not declarated in scope or invalid value")
	}
	idx, ok := e.UEP.Scope[engine.ByteEngineScopeBytecodeIdx].(*int)
	if !ok {
		return errors.New("Can't get bytecode index: not declarated in scope or invalid value")
	}
	for *(idx) < len(parsed) && *(idx) >= 0 {
		println(*idx)
		node := parsed[*(idx)]
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
