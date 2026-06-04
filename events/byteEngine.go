package events

import (
	"errors"

	"github.com/pt-main/lc/byteParsing"
	"github.com/pt-main/lc/system"
)

func (de *DefaultEvents) ByteParsingEvent(parser byteParsing.ParserInterface) system.EventType {
	return func(_e interface{}) error {
		e := _e.(*system.ByteEngine)
		input, ok := e.Scope["input_[]byte"].([]byte)
		if !ok {
			return errors.New("no input in scope")
		}
		nodes, err := parser.Parse(input)
		if err != nil {
			return err
		}
		e.Scope["parsed_[]ParsedBytes"] = nodes
		return nil
	}
}

func (de *DefaultEvents) ByteCallEvent(_e interface{}) error {
	e := _e.(*system.ByteEngine)
	_parsed, _ := e.Scope["parsed_[]ParsedBytes"]
	parsed, ok := _parsed.([]byteParsing.ParsedBytes)
	if !ok {
		return errors.New("Can't start call event. Invalid type of parsed result.")
	}
	u := byteParsing.Utils{}
	for _, node := range parsed {
		cmd_switch := u.BytesToInt(node.Switch)
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
