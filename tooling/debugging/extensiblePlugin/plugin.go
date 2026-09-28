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

// ExtensibleCLPlugin replaces the standard call loops with hookable ones.
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
	return &ExtensibleCLPlugin{
		de:     events.DefaultEvents{},
		Eu:     eu,
		Events: uep.Event,
		ETools: core.EventsTools{
			Events: uep.Event,
		},
	}
}

// changeEvents swaps the core event of the engine's call loop for the
// hookable one, or restores the original when enable is false.
func (ep *ExtensibleCLPlugin) changeEvents(enable bool) (string, error) {
	var name string
	var event core.EventType
	switch ep.Eu.Type {
	case public.StringEngineType:
		name = public.StringCallCallLoopEvent
	case public.ByteEngineType:
		name = public.ByteCallHotloopEvent
	}
	if enable {
		wasE, err := ep.ETools.GetCoreEvent(name)
		if err != nil {
			return "", err
		}
		ep.WasE = wasE
		switch ep.Eu.Type {
		case public.StringEngineType:
			event = ep.StringCallLoopEvent
		case public.ByteEngineType:
			event = ep.ByteCallHotLoopEvent
		}
	} else {
		event = ep.WasE
	}
	if err := ep.ETools.ChangeCoreEvent(name, event); err != nil {
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
	return err
}

func (ep *ExtensibleCLPlugin) Call(string, ...core.Option) (any, error) {
	return nil, nil
}

func (ep *ExtensibleCLPlugin) Run(any) (any, error) {
	return nil, nil
}
