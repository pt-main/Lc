package events

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/tap/color"
)

func (de *DefaultEvents) StringParsingEvent(_e interface{}, events *core.Events) error {
	e, ok := _e.(*engine.StringEngine)
	if !ok {
		return fmt.Errorf("Can't get byte engine: invalid input")
	}
	var okExit bool = false
	e.UEP.Logger.PrintLog("event", "start parsing event")
	defer e.UEP.Logger.PrintLog("event", fmt.Sprintf("end parsing event: [ok: %v]", okExit))
	input, ok := e.UEP.Scope[public.StringEngineScopeInput].(string)
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
	e.UEP.Scope[public.StringEngineScopeParsed] = nodes
	return nil
}

func (de *DefaultEvents) StringCallEvent(_e interface{}, events *core.Events) (err error) {
	e, ok := _e.(*engine.StringEngine)
	if !ok {
		return fmt.Errorf("Can't get byte engine: invalid input")
	}
	_parsed, _ := e.UEP.Scope[public.StringEngineScopeParsed]
	parsed, ok := _parsed.([]stringParsing.ParsedNode)
	if !ok {
		return errors.New("Can't start call event. Invalid type of parsed result.")
	}
	var okExit bool = false
	e.UEP.Logger.PrintLog("event", fmt.Sprintf("start call event: [parsed: %v]", parsed))
	defer e.UEP.Logger.PrintLog("event", fmt.Sprintf("end call event: [ok: %v]", okExit))
	var raw string
	for _, node := range parsed {
		ctx := e.UEP.GetContext()
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		raw = node.Raw
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("[?BRD]Panic recovered[?RT]: \n%v", r)
			}
		}()
		cmd_switch := node.Switch
		handler, ok := e.Commands[cmd_switch]
		if ok {
			err = handler.Handler(e, node)
		}
		if err != nil {
			err = errors.New("[?BRD]Handler error[?BRD]: \n" + err.Error())
		}
	}
	if err != nil {
		return errors.New(color.Set(fmt.Sprintf("[?RD]Error at:[?RT]\n[?BBK]    |[?RT]%v\n[?RD]Error:[?RT]\n[?RD]->[?RT]    %v",
			strings.ReplaceAll(raw, "\n", "\n    [?BBK]|[?RT]"),
			strings.ReplaceAll(err.Error(), "\n", "\n[?RD]->[?RT]    "),
		)))
	}
	okExit = true
	return nil
}
