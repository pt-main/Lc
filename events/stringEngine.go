package events

import (
	"errors"
	"strings"

	"github.com/pt-main/lc/stringParsing"
	"github.com/pt-main/lc/system"
)

func (de *DefaultEvents) StringParsingEvent(parser stringParsing.ParserInterface) system.EventType {
	return func(_e interface{}) error {
		e := _e.(*system.StringEngine)
		input, ok := e.Scope["input_string"].(string)
		if !ok {
			return errors.New("no input in scope")
		}
		nodes, err := parser.Parse(input)
		if err != nil {
			return err
		}
		e.Scope["parsed_[]ParsedNode"] = nodes
		return nil
	}
}

func (de *DefaultEvents) StringCallEvent(_e interface{}) error {
	e := _e.(*system.StringEngine)
	_parsed, _ := e.Scope["parsed_[]ParsedNode"]
	parsed, ok := _parsed.([]stringParsing.ParsedNode)
	if !ok {
		return errors.New("Can't start call event. Invalid type of parsed result.")
	}
	for _, node := range parsed {
		cmd_switch := node.Switch
		handler, ok := e.Commands[cmd_switch]
		var err error = nil
		if ok {
			err = handler.Handler([]interface{}{e, node})
		}
		_raw, ok := node.Metadata["__raw"]
		if !ok {
			return errors.New("Can't get raw string of code")
		}
		raw, ok := _raw.(string)
		if !ok {
			return errors.New("Can't get raw string of code")
		}
		if err != nil {
			return errors.New("Handler error at '" + strings.ReplaceAll(raw, "\n", "\\n") + "': " + err.Error())
		}
	}
	return nil
}
