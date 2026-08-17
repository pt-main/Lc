package extensiblePlugin

import (
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/engine/events"
	"github.com/pt-main/lc/public/errors"
)

// StringCallLoopEvent is the main loop for string-based command execution.
// It wraps the standard events with pre‑ and post‑iteration hooks.
//
// Err errors.ExtensiblePluginError:
//   - If core.ScopeGet fails to retrieve SCLEData.
//   - If StringCallEventIteration fails.
//   - Meta: EMK(0, "string") – the raw line that caused the error (if available).
func (ep *ExtensibleCLPlugin) StringCallLoopEvent(ev *core.Events, i *core.EventInput) (err core.ErrorInterface) {
	cld := i.Input.(events.StringCLDType)
	idx := cld.Idx
	parsed := cld.Parsed
	pLen := len(parsed)
	ctx := cld.Ctx
	e := cld.Engine
	ep.Events.Scope()[CLEScopeData] = SCLEData{
		Input:  cld,
		Idx:    idx,
		Parsed: parsed,
		PLen:   pLen,
		Ctx:    ctx,
		E:      e,
	}
	ep.Events.CallEvents(nil, CLEPreEvent, true) // ignore error (canWorkWithoutHandler)
	sd, err := core.ScopeGet[SCLEData](ep.Events.Scope(), CLEScopeData)
	if err != nil {
		return core.Wrap(errors.ExtensiblePluginError, err, "Failed to retrieve SCLEData from scope").
			WithMeta(core.EMK(0, "string"), "SCLEData")
	}
	cld = sd.Input
	idx = sd.Idx
	parsed = sd.Parsed
	pLen = sd.PLen
	ctx = sd.Ctx
	e = sd.E
	for *idx < pLen && *idx >= 0 {
		ep.Events.CallEvents(nil, CLEInPreEvent, true) // ignore error
		err = ep.de.StringCallEventIteration(parsed, idx, ev, ctx, e)
		if err != nil {
			// Обогащаем ошибку, если это core.Error
			if ce, ok := err.(*core.Error); ok {
				ce.WithMeta(core.EMK(0, "string"), "StringCallLoopEvent")
			} else {
				err = core.Wrap(errors.ExtensiblePluginError, err, "StringCallEventIteration failed").
					WithMeta(core.EMK(0, "string"), "StringCallLoopEvent")
			}
			return err
		}
		ep.Events.CallEvents(nil, CLEInPostEvent, true) // ignore error
	}
	ep.Events.CallEvents(nil, CLEPostEvent, true) // ignore error
	return nil
}

// ByteCallHotLoopEvent is the main loop for bytecode execution.
// It wraps the standard byte‑call loop with pre‑ and post‑iteration hooks.
//
// Err errors.ExtensiblePluginError:
//   - If the input is not of type ByteCLDType.
//     Meta: EMK(0, "string") – the actual type (if available).
//   - If core.ScopeGet fails to retrieve BCLEData.
//   - If ByteCallEventIteration fails.
//   - Meta: EMK(0, "int") – bytecode index where the error occurred.
func (ep *ExtensibleCLPlugin) ByteCallHotLoopEvent(ev *core.Events, i *core.EventInput) (err core.ErrorInterface) {
	hld, ok := i.Input.(events.ByteCLDType)
	if !ok {
		return core.Err(errors.ExtensiblePluginError, "Invalid event input: expected ByteCLDType").
			WithMeta(core.EMK(0, "string"), "ByteCLDType")
	}
	idx := hld.Idx
	ctx := hld.Ctx
	parsed := hld.Parsed
	p2len := len(parsed)
	e := hld.Engine
	ep.Events.Scope()[CLEScopeData] = BCLEData{
		Input:  hld,
		Idx:    idx,
		Parsed: parsed,
		PLen:   p2len,
		Ctx:    ctx,
		E:      e,
	}
	ep.Events.CallEvents(nil, CLEPreEvent, true) // ignore error
	sd, err := core.ScopeGet[BCLEData](ep.Events.Scope(), CLEScopeData)
	if err != nil {
		return core.Wrap(errors.ExtensiblePluginError, err, "Failed to retrieve BCLEData from scope").
			WithMeta(core.EMK(0, "string"), "BCLEData")
	}
	hld = sd.Input
	idx = sd.Idx
	parsed = sd.Parsed
	p2len = sd.PLen
	ctx = sd.Ctx
	e = sd.E
	for *idx < p2len && *idx >= 0 {
		ep.Events.CallEvents(nil, CLEInPreEvent, true) // ignore error
		if ctx.Err() != nil {
			err = core.Wrap(errors.ExtensiblePluginError, ctx.Err(), "Context cancelled during loop").
				WithMeta(core.EMK(0, "int"), *idx)
			break
		}
		err = ep.de.ByteCallEventIteration(idx, &parsed[*idx], e)
		if err != nil {
			// Обогащаем ошибку
			if ce, ok := err.(*core.Error); ok {
				ce.WithMeta(core.EMK(0, "int"), *idx)
			} else {
				err = core.Wrap(errors.ExtensiblePluginError, err, "ByteCallEventIteration failed at index %d", *idx).
					WithMeta(core.EMK(0, "int"), *idx)
			}
			break
		}
		ep.Events.CallEvents(nil, CLEInPostEvent, true) // ignore error
	}
	ep.Events.CallEvents(nil, CLEPostEvent, true) // ignore error
	return
}
