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

func NewPlugin(name, initEvent, mainEvent, closeEvent string) *Plugin {
	return &Plugin{
		Events:     core.NewEvents(nil),
		name:       name,
		InitEvent:  initEvent,
		MainEvent:  mainEvent,
		CloseEvent: closeEvent,
	}
}

func (p *Plugin) Name() string {
	return p.name
}

func (p *Plugin) Init(scope core.ScopeType, pm *PluginManager) error {
	for key, val := range scope {
		p.Events.Scope[key] = val
	}
	return p.Events.CallEvents(pm, p.InitEvent, true)
}

func (p *Plugin) Close() error {
	return p.Events.CallEvents(p, p.CloseEvent, false)
}

func (p *Plugin) Run(input any) (any, error) {
	err := p.Events.CallEvents(input, p.MainEvent, true)
	if err != nil {
		return nil, err
	}
	return p.Events.Scope[p.ScopeRunResultKey], err
}

func (p *Plugin) Call(name string, opts ...core.Option) (any, error) {
	err := p.Events.CallEvents(opts, name, false)
	if err != nil {
		return nil, err
	}
	res := p.Events.Scope[p.ScopeRunResultKey]
	p.Events.Scope[p.ScopeRunResultKey] = nil
	return res, nil
}
