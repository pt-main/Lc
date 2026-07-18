package plugin

import (
	"github.com/pt-main/lc/engine/core"
)

type Plugin struct {
	Events             *core.Events
	ScopeRunResultKey  string
	ScopeCallResultKey string

	InitEvent  string
	CloseEvent string
	MainEvent  string

	name string
}

func NewPlugin(
	name, initEvent,
	mainEvent, closeEvent,
	scopeRunResultKey,
	scopeCallResultKey string,
) *Plugin {
	return &Plugin{
		Events:             core.NewEvents(nil),
		name:               name,
		InitEvent:          initEvent,
		MainEvent:          mainEvent,
		CloseEvent:         closeEvent,
		ScopeRunResultKey:  scopeCallResultKey,
		ScopeCallResultKey: scopeCallResultKey,
	}
}

func (p *Plugin) Name() string {
	return p.name
}

func (p *Plugin) Init(scope core.ScopeType, pm *PluginManager) error {
	return p.Events.CallEvents(&core.EventInput{
		Input: pm,
		Option: &core.Option{
			Scope: scope,
		},
	}, p.InitEvent, true)
}

func (p *Plugin) Close() error {
	return p.Events.CallEvents(&core.EventInput{
		Input: p,
	}, p.CloseEvent, true)
}

func (p *Plugin) Run(input any) (any, error) {
	err := p.Events.CallEvents(&core.EventInput{
		Input: input,
	}, p.MainEvent, true)
	if err != nil {
		return nil, err
	}
	res, _ := core.ScopeGet[any](p.Events.Scope(), p.ScopeRunResultKey)
	return res, nil
}

func (p *Plugin) Call(name string, opts ...core.Option) (any, error) {
	err := p.Events.CallEvents(&core.EventInput{
		Input: opts,
	}, name, false)
	if err != nil {
		return nil, err
	}
	res, _ := core.ScopeGet[any](p.Events.Scope(), p.ScopeRunResultKey)
	return res, nil
}
