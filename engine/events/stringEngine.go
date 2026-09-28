package events

import (
	"context"

	"github.com/pt-main/lc/v2/engine"
	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/parsing"
	"github.com/pt-main/lc/v2/parsing/stringParsing"
	"github.com/pt-main/lc/v2/public"
	"github.com/pt-main/lc/v2/public/errors"
)

type StringCLDType CallLoopData[stringParsing.ParsedNode, engine.StringEngineInterface]

// Err errors.DefaultEventsSystemError. Cause from core.ScopeGet, e.Parser.Parse.
func (de *DefaultEvents) StringParsingEvent(_ *core.Events, i *core.EventInput) core.ErrorInterface {
	e, ok := i.Input.(*engine.StringEngine)
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Invalid input: expected *engine.StringEngine").
			WithMeta(core.EMK(0, "string"), "*engine.StringEngine")
	}
	return de.parseStringInput(e.UEP, e.Parser)
}

// AstParsingEvent is the *engine.AstEngine counterpart of StringParsingEvent:
// it reads the same input key and writes the same parsed-nodes key, so
// AstEngine.Process can pick the result up in Process.
//
// Err errors.DefaultEventsSystemError. Cause from core.ScopeGet, e.Parser.Parse.
func (de *DefaultEvents) AstParsingEvent(_ *core.Events, i *core.EventInput) core.ErrorInterface {
	e, ok := i.Input.(*engine.AstEngine)
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Invalid input: expected *engine.AstEngine").
			WithMeta(core.EMK(0, "string"), "*engine.AstEngine")
	}
	return de.parseStringInput(e.UEP, e.Parser)
}

func (de *DefaultEvents) parseStringInput(uep *core.UniversalEngineParams, parser parsing.ParserInterface[string, stringParsing.ParsedNode]) core.ErrorInterface {
	input, ok := uep.Scope[public.StringEngineScopeInput].(string)
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Input not found or invalid type in scope")
	}
	nodes, err := parser.Parse(input, &parsing.ParseOption{UEP: uep})
	if err != nil {
		return core.Wrap(errors.DefaultEventsSystemError, err, "Parser failed")
	}
	uep.Scope[public.StringEngineScopeParsed] = nodes
	return nil
}

// Err errors.DefaultEventsSystemError. Err errors.DefaultEventsPanicError.
// Err errors.DefaultEventsCallErrorContexted. With meta: EMK(0, "string") - raw line.
func (de *DefaultEvents) StringCallEvent(events *core.Events, i *core.EventInput) (err core.ErrorInterface) {
	events.Scope()[public.EventsScopeDERawLine] = "[NIL]"
	defer func() {
		if r := recover(); r != nil {
			err = core.Err(errors.DefaultEventsPanicError, "Panic recovered: %v", r)
		}
		if err == nil || err.Error() == core.ErrExit.Error() {
			return
		}
		raw, _ := core.ScopeGet[string](events.Scope(), public.EventsScopeDERawLine)
		if raw == "" {
			raw = "[NIL]"
		}
		err = core.Wrap(errors.DefaultEventsCallErrorContexted, err,
			"Error at line: %q", raw).
			WithMeta(core.EMK(0, "string"), raw)
	}()
	e, ok := i.Input.(*engine.StringEngine)
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Invalid input: expected *engine.StringEngine").
			WithMeta(core.EMK(0, "string"), "*engine.StringEngine")
	}
	parsed, err := core.ScopeGet[[]stringParsing.ParsedNode](e.UEP.Scope, public.StringEngineScopeParsed)
	if err != nil {
		return core.Wrap(errors.DefaultEventsSystemError, err, "Cannot get parsed nodes")
	}
	idx := 0
	cld := StringCLDType{
		Ctx:    e.UEP.GetContext(),
		Parsed: parsed,
		Engine: e,
		Idx:    &idx,
	}
	e.UEP.Scope[public.StringEngineScopeInstrIdx] = cld.Idx
	if err = events.CallEvents(&core.EventInput{Input: cld}, public.StringCallCallLoopEvent, false); err != nil {
		return core.Wrap(errors.DefaultEventsCallErrorContexted, err, "Call loop failed")
	}
	return nil
}

// Err errors.DefaultEventsCallErrorContexted.
func (de *DefaultEvents) StringCallLoopEvent(events *core.Events, i *core.EventInput) (err core.ErrorInterface) {
	cld := i.Input.(StringCLDType)
	idx := cld.Idx
	parsed := cld.Parsed
	pLen := len(parsed)
	ctx := cld.Ctx
	e := cld.Engine
	for *idx < pLen {
		if ctx.Err() != nil {
			return core.Wrap(errors.DefaultEventsCallErrorContexted, ctx.Err(), "Context cancelled")
		}
		if err = de.StringCallEventIteration(parsed, idx, events, ctx, e); err != nil {
			return core.Wrap(errors.DefaultEventsCallErrorContexted, err, "Iteration failed")
		}
		if *idx < 0 {
			return nil
		}
	}
	return nil
}

// Err errors.DefaultEventsCallErrorContexted. Err errors.DefaultEventsCallErrorUnknown.
// Err errors.DefaultEventsCallErrorHandler.
func (de *DefaultEvents) StringCallEventIteration(
	parsed []stringParsing.ParsedNode,
	idx *int,
	events *core.Events,
	ctx context.Context,
	e engine.StringEngineInterface,
) core.ErrorInterface {
	node := &parsed[*idx]
	canBeUnknown, err := core.ScopeGet[bool](e.GetUep().Scope, public.StringEngineScopeCanBeUnknown)
	if err != nil {
		canBeUnknown = true
	}
	events.Scope()[public.EventsScopeDERawLine] = node.Raw
	if ctx.Err() != nil {
		return core.Wrap(errors.DefaultEventsCallErrorContexted, ctx.Err(), "Context cancelled")
	}
	// A per-command lookup: GetCommands copies the whole map, which would
	// otherwise be paid on every parsed node.
	handler, ok := e.GetCommand(node.Switch)
	if ok {
		err = handler.Handler(e, node)
	} else if !canBeUnknown {
		return core.Err(errors.DefaultEventsCallErrorUnknown, "Unknown command: %s", node.Switch).
			WithMeta(core.EMK(0, "string"), node.Switch)
	}
	if err != nil {
		if err.Error() == core.ErrExit.Error() {
			*idx = -1
			return nil
		}
		return core.Wrap(errors.DefaultEventsCallErrorHandler, err, "Handler failed for %s", node.Switch).
			WithMeta(core.EMK(0, "string"), node.Switch)
	}
	(*idx)++
	return nil
}
