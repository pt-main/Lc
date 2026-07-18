package extensiblePlugin

import (
	"fmt"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/engine/events"
)

func (ep *ExtensibleCLPlugin) StringCallLoopEvent(ev *core.Events, i *core.EventInput) (err error) {
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
	ep.Events.CallEvents(nil, CLEPreEvent, true)
	sd, err := core.ScopeGet[SCLEData](ep.Events.Scope(), CLEScopeData)
	if err != nil {
		return err
	}
	cld = sd.Input
	idx = sd.Idx
	parsed = sd.Parsed
	pLen = sd.PLen
	ctx = sd.Ctx
	e = sd.E
	for *idx < pLen && *idx >= 0 {
		ep.Events.CallEvents(nil, CLEInPreEvent, true)
		err = ep.de.StringCallEventIteration(parsed, idx, ev, ctx, e)
		if err != nil {
			return
		}
		ep.Events.CallEvents(nil, CLEInPostEvent, true)
	}
	ep.Events.CallEvents(nil, CLEInPostEvent, true)
	return
}

func (ep *ExtensibleCLPlugin) ByteCallHotLoopEvent(ev *core.Events, i *core.EventInput) (err error) {
	hld, ok := i.Input.(events.ByteCLDType)
	if !ok {
		return fmt.Errorf("Invalid event input: can't get hotloop data")
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
	ep.Events.CallEvents(nil, CLEPreEvent, true)
	sd, err := core.ScopeGet[BCLEData](ep.Events.Scope(), CLEScopeData)
	if err != nil {
		return err
	}
	hld = sd.Input
	idx = sd.Idx
	parsed = sd.Parsed
	p2len = sd.PLen
	ctx = sd.Ctx
	e = sd.E
	for *idx < p2len && *idx >= 0 {
		ep.Events.CallEvents(nil, CLEInPreEvent, true)
		if ctx.Err() != nil {
			err = ctx.Err()
			break
		}
		err = ep.de.ByteCallEventIteration(
			idx, &parsed[*idx], e,
		)
		if err != nil {
			break
		}
		ep.Events.CallEvents(nil, CLEInPostEvent, true)
	}
	ep.Events.CallEvents(nil, CLEPostEvent, true)
	return
}
