package events

import (
	"context"
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

type StringCLDType CallLoopData[stringParsing.ParsedNode, engine.StringEngine]

func (de *DefaultEvents) StringParsingEvent(events *core.Events, i *core.EventInput) error {
	e, ok := i.Input.(*engine.StringEngine)
	if !ok {
		return fmt.Errorf("Can't get byte engine: invalid input")
	}
	var okExit bool = false
	e.UEP.Logger.PrintLog(public.LogEvents, "start parsing event")
	defer func() {
		e.UEP.Logger.PrintLog(public.LogEvents, fmt.Sprintf("end parsing event: [ok: %v]", okExit))
	}()
	input, ok := e.UEP.Scope[public.StringEngineScopeInput].(string)
	if !ok {
		return errors.New("No input in scope or invalid input")
	}
	e.UEP.Logger.PrintLog(public.LogEvents, fmt.Sprintf("start parsing: [parser: %s]", e.Parser))
	nodes, err := e.Parser.Parse(input, &parsing.ParseOption{UEP: e.UEP})
	if err != nil {
		return err
	}
	e.UEP.Logger.PrintLog(public.LogEvents, "end parsing")
	okExit = true
	e.UEP.Scope[public.StringEngineScopeParsed] = nodes
	return nil
}

func (de *DefaultEvents) StringCallEvent(events *core.Events, i *core.EventInput) (err error) {
	var status string
	events.Scope()[public.EventsScopeDERawLine] = "[NIL]"
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("[?BRD]Panic recovered[?RT]: \n%v", r)
		}
		if err != nil {
			raw := "[NIL]"
			if res, err1 := core.ScopeGet[string](events.Scope(), public.EventsScopeDERawLine); err1 == nil {
				raw = res
			}
			text := fmt.Sprintf("[?RD]Error[?RT]: [?YW]Status [?RT][%v]:\n"+
				"[?RD]Error at:[?RT]\n[?BBK]    |[?RT]%v\n[?RD]Error:[?RT]\n[?RD]->[?RT]    %v",
				status, strings.ReplaceAll(raw, "\n", "\n[?BBK]>    |[?RT]"),
				strings.ReplaceAll(err.Error(), "\n", "\n[?RD]->[?RT]    "))
			err = errors.New(color.Set(text))
		}
	}()
	status = "SYSTEM: initializing event..."
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
	e.UEP.Logger.PrintLog(public.LogEvents, fmt.Sprintf("start call event: [parsed: %v]", parsed))
	defer func() {
		e.UEP.Logger.PrintLog(public.LogEvents, fmt.Sprintf("end parsing event: [ok: %v]", okExit))
	}()
	status = "SYSTEM: call cycle..."
	_idx := 0
	cld := StringCLDType{
		Ctx:    e.UEP.GetContext(),
		Parsed: parsed,
		Engine: e,
		Idx:    &_idx,
	}
	err = events.CallEvents(&core.EventInput{Input: cld}, public.StringCallCalloopEvent, false)
	if err != nil {
		return
	}
	okExit = true
	return nil
}

func (de *DefaultEvents) StringCallEventIteration(parsed []stringParsing.ParsedNode, idx *int,
	events *core.Events, ctx context.Context, e *engine.StringEngine) (err error) {
	node := parsed[*idx]
	events.Scope()[public.EventsScopeDERawLine] = node.Raw
	if err = ctx.Err(); err != nil {
		return
	}
	e.UEP.Logger.PrintLog(public.LogEvents, fmt.Sprintf("process (in call event): [node: %v]", node))
	cmd_switch := node.Switch
	handler, ok := e.Commands[cmd_switch]
	if ok {
		err = handler.Handler(e, &node)
	}
	if err != nil {
		return errors.New("[?BRD]Handler error[?BRD]: \n" + err.Error())
	}
	*idx += 1
	return
}

func (de *DefaultEvents) StringCallLoopEvent(events *core.Events, i *core.EventInput) (err error) {
	cld := i.Input.(StringCLDType)
	idx := cld.Idx
	parsed := cld.Parsed
	pLen := len(parsed)
	ctx := cld.Ctx
	e := cld.Engine
	for *idx < pLen && *idx >= 0 {
		err = de.StringCallEventIteration(parsed, idx, events, ctx, e)
		if err != nil {
			return
		}
	}
	return
}
