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

func (de *DefaultEvents) StringParsingEvent(events *core.Events, i *core.EventInput) error {
	e, ok := i.Input.(*engine.StringEngine)
	if !ok {
		return fmt.Errorf("Can't get byte engine: invalid input")
	}
	var okExit bool = false
	e.UEP.Logger.PrintLog("event", "start parsing event")
	defer func() {
		e.UEP.Logger.PrintLog("event", fmt.Sprintf("end parsing event: [ok: %v]", okExit))
	}()
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

func (de *DefaultEvents) StringCallEvent(events *core.Events, i *core.EventInput) (err error) {
	var raw string
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("[?BRD]Panic recovered[?RT]: \n%v", r)
		}
		if err != nil {
			err = errors.New(color.Set(fmt.Sprintf("[?RD]Error at:[?RT]\n[?BBK]    |[?RT]%v\n[?RD]Error:[?RT]\n[?RD]->[?RT]    %v",
				strings.ReplaceAll(raw, "\n", "\n    [?BBK]|[?RT]"),
				strings.ReplaceAll(err.Error(), "\n", "\n[?RD]->[?RT]    "),
			)))
		}
	}()
	raw = "SYSTEM: initializing event..."
	e, ok := i.Input.(*engine.StringEngine)
	if !ok {
		err = fmt.Errorf("Can't get byte engine: invalid input")
		return
	}
	_parsed, _ := e.UEP.Scope[public.StringEngineScopeParsed]
	parsed, ok := _parsed.([]stringParsing.ParsedNode)
	if !ok {
		err = errors.New("Can't start call event. Invalid type of parsed result.")
		return
	}
	var okExit bool = false
	e.UEP.Logger.PrintLog("event", fmt.Sprintf("start call event: [parsed: %v]", parsed))
	defer func() {
		e.UEP.Logger.PrintLog("event", fmt.Sprintf("end parsing event: [ok: %v]", okExit))
	}()
	raw = "SYSTEM: start call cycle..."
	for _, node := range parsed {
		ctx := e.UEP.GetContext()
		select {
		case <-ctx.Done():
			err = ctx.Err()
			return
		default:
		}
		e.UEP.Logger.PrintLog("event", fmt.Sprintf("process (in call event): [node: %v]", node))
		raw = node.Raw
		cmd_switch := node.Switch
		handler, ok := e.Commands[cmd_switch]
		if ok {
			err = handler.Handler(e, node)
		}
		if err != nil {
			err = errors.New("[?BRD]Handler error[?BRD]: \n" + err.Error())
			break
		}
	}
	if err != nil {
		return
	}
	okExit = true
	return nil
}
