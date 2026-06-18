package plugin

import (
	"errors"
	"fmt"

	"github.com/pt-main/lc/engine/core"
)

type Plugin struct {
	Events         *core.Events
	ScopeResultKey string

	InitEvent  string
	CloseEvent string
	MainEvent  string

	Name string
}

func (p *Plugin) InitPlugin(Scope core.ScopeType, pm *PluginManager) error {
	for key, val := range Scope {
		p.Events.Scope[key] = val
	}
	return p.Events.CallEvents(pm, p.InitEvent, true)
}

func (p *Plugin) ClosePlugin() error {
	return p.Events.CallEvents(p, p.CloseEvent, false)
}

func (p *Plugin) RunPlugin(input any) (any, error) {
	err := p.Events.CallEvents(input, p.MainEvent, true)
	return p.Events.Scope[p.ScopeResultKey], err
}

type PluginManager struct {
	Plugins map[string]*Plugin
	Scope   core.ScopeType
}

func (pm *PluginManager) AddPlugin(plugin *Plugin) error {
	pm.Plugins[plugin.Name] = plugin
	return plugin.InitPlugin(pm.Scope, pm)
}

func (pm *PluginManager) DeletePlugin(name string) {
	plugins := map[string]*Plugin{}
	for _name, plugin := range pm.Plugins {
		if _name != name {
			plugins[_name] = plugin
		} else {
			plugin.ClosePlugin()
		}
	}
	pm.Plugins = plugins
}

func (pm *PluginManager) GetPlugin(name string) (*Plugin, error) {
	plugin, ok := pm.Plugins[name]
	if !ok {
		return nil, errors.New("Can't find plugin: " + name)
	}
	return plugin, nil
}

func (pm *PluginManager) CallPlugin(name string, input any) (any, error) {
	plugin, err := pm.GetPlugin(name)
	if err != nil {
		return nil, fmt.Errorf("Call plugin '%v' error: %v", name, err)
	}
	return plugin.RunPlugin(input)
}
