package events

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/stringParsing"
)

func (de *DefaultEvents) StringParsingEvent(_e interface{}, events *core.Events) error {
	e, ok := _e.(*engine.StringEngine)
	if !ok {
		return fmt.Errorf("Can't get byte engine: invalid input")
	}
	var okExit bool = false
	e.UEP.Logger.PrintLog("event", "start parsing event")
	defer e.UEP.Logger.PrintLog("event", fmt.Sprintf("end parsing event: [ok: %v]", okExit))
	input, ok := e.UEP.Scope[engine.StringEngineScopeInput].(string)
	if !ok {
		return errors.New("No input in scope or invalid input")
	}
	e.UEP.Logger.PrintLog("event", fmt.Sprintf("start parsing: [parser: %s]", e.Parser))
	nodes, err := e.Parser.Parse(input, &parsing.ParseOption{UEP: e.UEP})
	if err != nil {
		return err
	}
	e.UEP.Logger.PrintLog("event", "end parsing")
	okExit = true
	e.UEP.Scope[engine.StringEngineScopeParsed] = nodes
	return nil
}

func (de *DefaultEvents) StringCallEvent(_e interface{}, events *core.Events) error {
	e, ok := _e.(*engine.StringEngine)
	if !ok {
		return fmt.Errorf("Can't get byte engine: invalid input")
	}
	_parsed, _ := e.UEP.Scope[engine.StringEngineScopeParsed]
	parsed, ok := _parsed.([]stringParsing.ParsedNode)
	if !ok {
		return errors.New("Can't start call event. Invalid type of parsed result.")
	}
	var okExit bool = false
	e.UEP.Logger.PrintLog("event", fmt.Sprintf("start call event: [parsed: %v]", parsed))
	defer e.UEP.Logger.PrintLog("event", fmt.Sprintf("end call event: [ok: %v]", okExit))
	var err error = nil
	for _, node := range parsed {
		ctx := e.UEP.GetContext()
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		cmd_switch := node.Switch
		handler, ok := e.Commands[cmd_switch]
		if ok {
			err = handler.Handler(e, node)
		}
		if err != nil {
			err = errors.New("Handler error: " + err.Error())
		}
		if err != nil {
			raw := node.Raw
			return fmt.Errorf("[?RD]Error at:[?RT]\n[?BBK]    |[?RT]%v\n[?RD]Error:[?RT]\n[?RD]->[?RT]    %v",
				strings.ReplaceAll(raw, "\n", "\n[?BBK]    |[?RT]"),
				strings.ReplaceAll(err.Error(), "\n", "\n[?RD]->[?RT]    "),
			)
		}
	}
	okExit = true
	return nil
}
