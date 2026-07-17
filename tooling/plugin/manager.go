package plugin

import (
	"fmt"

	"github.com/pt-main/lc/engine/core"
)

type PluginManager struct {
	Plugins map[string]PluginInterface
	Scope   core.ScopeType
	Flags   []string
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
		return fmt.Errorf("Plugin %s already loaded", name)
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
		return nil, fmt.Errorf("Plugin %s not found", name)
	}
	return plugin, nil
}

func (pm *PluginManager) RunPlugin(name string, input any) (any, error) {
	plugin, err := pm.GetPlugin(name)
	if err != nil {
		return nil, err
	}
	return plugin.Run(input)
}

func (pm *PluginManager) CallPluginMethod(name, method string, opts ...core.Option) (any, error) {
	plugin, err := pm.GetPlugin(name)
	if err != nil {
		return nil, err
	}
	return plugin.Call(method, opts...)
}
