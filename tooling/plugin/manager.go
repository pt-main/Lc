package plugin

import (
	"fmt"

	"github.com/pt-main/lc/v2/engine/core"
)

// PluginManager holds the registered plugins and the scope they communicate
// through. The private flags are only reachable through Tools.
type PluginManager struct {
	Plugins map[string]PluginInterface
	Scope   core.ScopeType
	flags   []string
}

// NewPluginManager creates a manager over the given scope, or over a fresh
// empty one when scope is nil.
func NewPluginManager(scope core.ScopeType) *PluginManager {
	if scope == nil {
		scope = make(core.ScopeType)
	}
	return &PluginManager{
		Plugins: make(map[string]PluginInterface),
		Scope:   scope,
	}
}

// AddPlugin registers the plugin under plugin.Name() and initialises it. A
// name that is already taken is an error.
func (pm *PluginManager) AddPlugin(plugin PluginInterface) error {
	name := plugin.Name()
	if _, exists := pm.Plugins[name]; exists {
		return fmt.Errorf("Plugin %s already loaded", name)
	}
	pm.Plugins[name] = plugin
	return plugin.Init(pm.Scope, pm)
}

// DeletePlugin closes the plugin and drops it. A name that is not registered
// is not an error.
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

// GetPlugin looks a plugin up by name.
func (pm *PluginManager) GetPlugin(name string) (PluginInterface, error) {
	plugin, ok := pm.Plugins[name]
	if !ok {
		return nil, fmt.Errorf("Plugin %s not found", name)
	}
	return plugin, nil
}

// RunPlugin calls Run on the named plugin with the given input.
func (pm *PluginManager) RunPlugin(name string, input any) (any, error) {
	plugin, err := pm.GetPlugin(name)
	if err != nil {
		return nil, err
	}
	return plugin.Run(input)
}

// CallPluginMethod calls the named method of the named plugin.
func (pm *PluginManager) CallPluginMethod(name, method string, opts ...core.Option) (any, error) {
	plugin, err := pm.GetPlugin(name)
	if err != nil {
		return nil, err
	}
	return plugin.Call(method, opts...)
}

// End closes and drops every plugin, stopping at the first error.
func (pm *PluginManager) End() error {
	for name := range pm.Plugins {
		if err := pm.DeletePlugin(name); err != nil {
			return err
		}
	}
	return nil
}
