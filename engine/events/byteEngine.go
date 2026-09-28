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
func (de *DefaultEvents) ByteParsingEvent(_ *core.Events, i *core.EventInput) core.ErrorInterface {
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
	parsed []byteParsing.ParsedBytes, endianness public.EndianType,
	u bytecode.Utils, abis map[int]bool,
	cmds map[int]core.CommandMeta[engine.ByteEngineInterface, byteParsing.ParsedBytes],
) ([]ByteCallAttr, core.ErrorInterface) {
	res := make([]ByteCallAttr, 0, len(parsed))
	for i := range parsed {
		node := &parsed[i]
		cmdSwitch := u.BytesToInt(node.Switch, endianness)
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
			RawNode: node,
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
	var attrs []ByteCallAttr
	var lastCmd int
	defer func() {
		if r := recover(); r != nil {
			err = core.Err(errors.DefaultEventsPanicError, "Panic recovered: %v", r)
		}
		if err == nil {
			return
		}
		if err.Error() == core.ErrExit.Error() {
			if idx != nil {
				*idx = -1
			}
			return
		}
		idxVal := 0
		if idx != nil {
			idxVal = *idx
		}
		// Without an index the hot loop cannot tell which opcode failed, so
		// the last known one is used instead of an unconditional zero.
		cmdVal := lastCmd
		if idx != nil && idxVal >= 0 && idxVal < len(attrs) && attrs[idxVal].RawNode != nil {
			cmdVal = (&bytecode.Utils{}).BytesToInt(attrs[idxVal].RawNode.Switch, public.LittleEndian)
		}
		err = core.Wrap(errors.DefaultEventsCallErrorContexted, err,
			"Error at cmd=%v, bcIdx=%v", cmdVal, idxVal).
			WithMeta(core.EMK(0, "int"), cmdVal).
			WithMeta(core.EMK(1, "int"), idxVal)
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
	endianness, ok := e.GetUep().Scope[public.ByteEngineScopeEndianness].(public.EndianType)
	if !ok {
		return core.Err(errors.DefaultEventsSystemError, "Invalid endianness in scope")
	}
	idx, err = core.ScopeGet[*int](e.GetUep().Scope, public.ByteEngineScopeBytecodeIdx)
	if err != nil {
		return core.Wrap(errors.DefaultEventsSystemError, err, "Cannot get bytecode index")
	}
	ctx := e.GetUep().GetContext()
	// The maps are read without e.mu otherwise, which races with a concurrent
	// NewCommand and can crash with "concurrent map read and map write".
	cmds := e.GetCommands()
	abis := e.GetAutoBytecodeIndexShift()

	attrs, err = de.ByteCallPreprocess(parsed, endianness, u, abis, cmds)
	if err != nil {
		return core.Wrap(errors.DefaultEventsSystemError, err, "Preprocessing failed")
	}

	if err = events.CallEvents(&core.EventInput{Input: ByteCLDType{
		Ctx: ctx, Parsed: attrs, Engine: e, Idx: idx, Other: &parsed,
	}}, public.ByteCallHotloopEvent, false); err != nil {
		return core.Wrap(errors.DefaultEventsCallErrorContexted, err, "Hot-loop event failed")
	}
	return nil
}

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
		// The first iteration is always checked, otherwise a program shorter
		// than checkPeriod would finish without ever seeing a cancelled ctx.
		if iter == 1 || iter&checkInterval == 0 {
			if ctx.Err() != nil {
				return core.Wrap(errors.DefaultEventsCallErrorContexted, ctx.Err(), "Context cancelled (at %v iter)", iter)
			}
		}
		idxN := *idx
		if uint(idxN) >= uint(p2len) {
			break
		}
		node := &parsed[idxN]
		if err = de.ByteCallEventIteration(idx, node, e); err != nil {
			if node.Abis {
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
	return parsed.Handler(e, parsed.RawNode)
}
