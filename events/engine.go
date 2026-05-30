package events

import (
	"errors"
	"strings"

	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/system"
)

type DefaultEvents struct{}

func (de *DefaultEvents) ParsingEvent(parser parsing.Parser) system.EventType {
	return func(e *system.Engine) error {
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

func (de *DefaultEvents) CallEvent(e *system.Engine) error {
	_parsed, _ := e.Scope["parsed_[]ParsedNode"]
	parsed, ok := _parsed.([]parsing.ParsedNode)
	if !ok {
		return errors.New("Cant start call event. Invalid type of parsed result.")
	}
	for _, node := range parsed {
		cmd_switch := node.Switch
		handler := e.Commands[cmd_switch].Handler
		err := handler(e, node)
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
