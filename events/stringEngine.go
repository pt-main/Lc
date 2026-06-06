package events

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/stringParsing"
)

func (de *DefaultEvents) StringParsingEvent(_e interface{}) error {
	e, ok := _e.(*engine.StringEngine)
	if !ok {
		return fmt.Errorf("Can't get byte engine: invalid iput")
	}
	input, ok := e.UEP.Scope["input_string"].(string)
	if !ok {
		return errors.New("No input in scope or invalid input")
	}
	nodes, err := e.Parser.Parse(input)
	if err != nil {
		return err
	}
	e.UEP.Scope["parsed_[]ParsedNode"] = nodes
	return nil
}

func (de *DefaultEvents) StringCallEvent(_e interface{}) error {
	e, ok := _e.(*engine.StringEngine)
	if !ok {
		return fmt.Errorf("Can't get byte engine: invalid iput")
	}
	_parsed, _ := e.UEP.Scope["parsed_[]ParsedNode"]
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
