package plugin

import (
	"fmt"

	"github.com/pt-main/lc/engine/core"
)

type PluginInterface interface {
	Name() string
	Init(scope core.ScopeType, pm *PluginManager) error
	Close() error
	Run(input any) (any, error)
}

type Plugin struct {
	Events         *core.Events
	ScopeResultKey string

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
	return p.Events.Scope[p.ScopeResultKey], err
}

type PluginManager struct {
	Plugins map[string]PluginInterface
	Scope   core.ScopeType
}

func NewPluginManager(scope core.ScopeType) *PluginManager {
	if scope == nil {
		scope = make(core.ScopeType)
	}
	return &PluginManager{
		Plugins: make(map[string]PluginInterface),
		Scope:   scope,
	}
}

func (pm *PluginManager) AddPlugin(plugin PluginInterface) error {
	name := plugin.Name()
	if _, exists := pm.Plugins[name]; exists {
		return fmt.Errorf("plugin %s already loaded", name)
	}
	pm.Plugins[name] = plugin
	return plugin.Init(pm.Scope, pm)
}

func (pm *PluginManager) DeletePlugin(name string) error {
	if plugin, exists := pm.Plugins[name]; exists {
		err := plugin.Close()
		if err != nil {
			return err
		}
		delete(pm.Plugins, name)
	}
	return nil
}

func (pm *PluginManager) GetPlugin(name string) (PluginInterface, error) {
	plugin, ok := pm.Plugins[name]
	if !ok {
		return nil, fmt.Errorf("plugin %s not found", name)
	}
	return plugin, nil
}

func (pm *PluginManager) CallPlugin(name string, input any) (any, error) {
	plugin, err := pm.GetPlugin(name)
	if err != nil {
		return nil, err
	}
	return plugin.Run(input)
}
