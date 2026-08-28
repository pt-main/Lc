package events

import (
	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
	"github.com/pt-main/lc/tooling/bytecode"
)

type ByteCLDType CallLoopData[ByteCallAttr, engine.ByteEngineInterface]

// Err errors.DefaultEventsSystemError.
// With meta: EMK(0, "string") - expected type.
// Cause from core.ScopeGet, e.Parser.Parse.
func (de *DefaultEvents) ByteParsingEvent(events *core.Events, i *core.EventInput) core.ErrorInterface {
	e, ok := i.Input.(*engine.ByteEngine)
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Invalid input: expected *engine.ByteEngine").
			WithMeta(core.EMK(0, "string"), "*engine.ByteEngine")
	}
	input, err := core.ScopeGet[[]byte](e.UEP.Scope, public.ByteEngineScopeInput)
	if err != nil {
		return core.Wrap(errors.DefaultEventsSystemError, err, "Cannot get input from scope")
	}
	nodes, err := e.Parser.Parse(input, &parsing.ParseOption{UEP: e.UEP})
	if err != nil {
		return core.Wrap(errors.DefaultEventsSystemError, err, "Parser failed")
	}
	e.UEP.Scope[public.ByteEngineScopeParsed] = nodes
	return nil
}

type ByteCallAttr struct {
	RawNode *byteParsing.ParsedBytes
	Abis    bool
	Handler core.CommandType[engine.ByteEngineInterface, byteParsing.ParsedBytes]
}

// Err errors.DefaultEventsCallErrorCmdNotFound.
// With meta: EMK(0, "int") - opcode.
func (de *DefaultEvents) ByteCallPreprocess(
	parsed []byteParsing.ParsedBytes, endianess public.EndianType,
	u bytecode.Utils, abis map[int]bool,
	cmds map[int]core.CommandMeta[engine.ByteEngineInterface, byteParsing.ParsedBytes],
) ([]ByteCallAttr, core.ErrorInterface) {
	res := make([]ByteCallAttr, 0, len(parsed))
	for _, node := range parsed {
		cmdSwitch := u.BytesToInt(node.Switch, endianess)
		handler, ok := cmds[cmdSwitch]
		if !ok {
			return nil, core.Err(errors.DefaultEventsCallErrorCmdNotFound, "Opcode %d not registered", cmdSwitch).
				WithMeta(core.EMK(0, "int"), cmdSwitch)
		}
		autoshift, ok := abis[cmdSwitch]
		if !ok {
			return nil, core.Err(errors.DefaultEventsCallErrorCmdNotFound, "Autoshift config missing for opcode %d", cmdSwitch).
				WithMeta(core.EMK(0, "int"), cmdSwitch)
		}
		res = append(res, ByteCallAttr{
			RawNode: &node,
			Handler: handler.Handler,
			Abis:    autoshift,
		})
	}
	return res, nil
}

// Err errors.DefaultEventsSystemError.
// Err errors.DefaultEventsPanicError.
// Err errors.DefaultEventsCallErrorContexted.
// With meta: EMK(0, "int") - cmd, EMK(1, "int") - bcIdx, EMK(2, "string") - pb.
func (de *DefaultEvents) ByteCallEvent(events *core.Events, i *core.EventInput) (err core.ErrorInterface) {
	var idx *int
	var parsed2 []ByteCallAttr
	var lastCmd int
	defer func() {
		if r := recover(); r != nil {
			err = core.Err(errors.DefaultEventsPanicError, "Panic recovered: %v", r)
		}
		if err != nil {
			if err.Error() == core.ErrExit.Error() {
				if idx != nil {
					*idx = -1
				}
				return
			}
			idxV := 0
			if idx != nil {
				idxV = *idx
			}
			err = core.Wrap(errors.DefaultEventsCallErrorContexted, err,
				"Error at cmd=%v, bcIdx=%v", lastCmd, idxV).
				WithMeta(core.EMK(0, "int"), lastCmd).
				WithMeta(core.EMK(1, "int"), idxV)
		}
	}()
	e, ok := i.Input.(*engine.ByteEngine)
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Invalid input: expected *engine.ByteEngine").
			WithMeta(core.EMK(0, "string"), "*engine.ByteEngine")
	}
	_parsed, ok := e.GetUep().Scope[public.ByteEngineScopeParsed]
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Parsed data not found in scope")
	}
	parsed, ok := _parsed.([]byteParsing.ParsedBytes)
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Parsed data has wrong type")
	}
	u := bytecode.Utils{}
	endianess, ok := e.GetUep().Scope[public.ByteEngineScopeEndianess].(public.EndianType)
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Invalid endianess in scope")
	}
	idx, err = core.ScopeGet[*int](e.GetUep().Scope, public.ByteEngineScopeBytecodeIdx)
	if err != nil {
		return core.Wrap(errors.DefaultEventsSystemError, err, "Cannot get bytecode index")
	}
	ctx := e.GetUep().GetContext()
	cmds := e.Commands
	abis := e.AutoBytecodeIndexShift

	parsed2, err = de.ByteCallPreprocess(parsed, endianess, u, abis, cmds)
	if err != nil {
		return core.Wrap(errors.DefaultEventsSystemError, err, "Preprocessing failed")
	}

	if err = events.CallEvents(&core.EventInput{Input: ByteCLDType{
		Ctx: ctx, Parsed: parsed2, Engine: e, Idx: idx, Other: &parsed,
	}}, public.ByteCallHotloopEvent, false); err != nil {
		return core.Wrap(errors.DefaultEventsCallErrorContexted, err, "Hot-loop event failed")
	}
	return nil
}

// Err errors.DefaultEventsCallErrorContex.
// Err errors.DefaultEventsCallErrorContexted.
func (de *DefaultEvents) ByteCallHotLoopEvent(events *core.Events, i *core.EventInput) (err core.ErrorInterface) {
	hld, ok := i.Input.(ByteCLDType)
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Invalid event input: expected ByteCLDType").
			WithMeta(core.EMK(0, "string"), "ByteCLDType")
	}
	idx := hld.Idx
	ctx := hld.Ctx
	parsed := hld.Parsed
	p2len := len(parsed)
	e := hld.Engine
	iter := 0
	var checkInterval int
	checkInterval, err = core.ScopeGet[int](e.GetUep().Scope, public.ByteEngineScopeHotloopCtxCheckPeriod)
	if err != nil {
		checkInterval = 255 // 2^8-1
	}
	for {
		iter++
		if iter&int(checkInterval) == 0 {
			if ctx.Err() != nil {
				return core.Wrap(errors.DefaultEventsCallErrorContex, ctx.Err(), "Context cancelled (at %v iter)", iter)
			}
		}
		idxN := *idx
		if uint(idxN) >= uint(p2len) {
			break
		}
		node := &parsed[idxN]
		//go:inline
		if err = de.ByteCallEventIteration(idx, node, e); err != nil {
			if node.Abis == true {
				(*idx)--
			}
			return core.Wrap(errors.DefaultEventsCallErrorHandler, err, "Handler failed")
		}
	}
	return nil
}

// Err errors.DefaultEventsCallErrorHandler.
func (de *DefaultEvents) ByteCallEventIteration(
	idx *int,
	parsed *ByteCallAttr, e engine.ByteEngineInterface,
) core.ErrorInterface {
	if parsed.Abis {
		*idx++
	}
	//go:inline
	return parsed.Handler(e, parsed.RawNode)
}
