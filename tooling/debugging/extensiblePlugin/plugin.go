package extensiblePlugin

import (
	"fmt"

	"github.com/pt-main/lc"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/engine/events"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/tooling/plugin"
)

const Name = "extensible call loop"

// # ExtensibleCLPlugin
//
// Replace standart calloops and add calloops with event hooks.
// System plugin.
type ExtensibleCLPlugin struct {
	de     events.DefaultEvents
	Eu     *lc.EngineUniversal
	Events core.EventsInterface
	ETools core.EventsTools
	WasE   core.EventType
}

func New(eu *lc.EngineUniversal) *ExtensibleCLPlugin {
	uep, err := eu.GetUEP()
	if err != nil {
		panic("Can't add extensible plugin: " + err.Error())
	}
	e := uep.Event
	return &ExtensibleCLPlugin{
		de:     events.DefaultEvents{},
		Eu:     eu,
		Events: e,
		ETools: core.EventsTools{
			Events: e,
		},
		WasE: nil,
	}
}

func (ep *ExtensibleCLPlugin) changeEvents(val bool) (string, error) {
	euType := ep.Eu.Type
	var name string
	var event core.EventType
	switch euType {
	case public.StringEngineType:
		name = public.StringCallCalloopEvent
	case public.ByteEngineType:
		name = public.ByteCallHotloopEvent
	}
	switch val {
	case true:
		var err error
		ep.WasE, err = ep.ETools.GetCoreEvent(name)
		if err != nil {
			return "", err
		}
		switch euType {
		case public.StringEngineType:
			event = ep.StringCallLoopEvent
		case public.ByteEngineType:
			event = ep.ByteCallHotLoopEvent
		}
	default:
		event = ep.WasE
	}
	err := ep.ETools.ChangeCoreEvent(name, event)
	if err != nil {
		return "", err
	}
	return name, nil
}

func (ep *ExtensibleCLPlugin) Init(scope core.ScopeType, pm *plugin.PluginManager) error {
	eu, ok := scope[public.PluginsScopeEuPtr].(*lc.EngineUniversal)
	if !ok {
		return fmt.Errorf("Bad scope: can't find EngineUniversal. Plugins didn't load.")
	}
	ep.Eu = eu
	uep, err := eu.GetUEP()
	if err != nil {
		return err
	}
	_, err = ep.changeEvents(true)
	if err != nil {
		return err
	}
	ep.ETools = core.EventsTools{
		Events: uep.Event,
	}
	(&plugin.Tools{Pm: pm}).SetFlag(ECLFlag)
	return nil
}

func (ep *ExtensibleCLPlugin) Name() string { return Name }

func (ep *ExtensibleCLPlugin) Close() error {
	_, err := ep.changeEvents(false)
	if err != nil {
		return err
	}
	return nil
}

func (ep *ExtensibleCLPlugin) Call(string, ...core.Option) (o any, e error) {
	return
}

func (ep *ExtensibleCLPlugin) Run(input any) (o any, e error) {
	return
}
