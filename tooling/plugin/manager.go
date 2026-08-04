package plugin

import (
	"fmt"

	"github.com/pt-main/lc/engine/core"
)

// # PluginManager
//
// Plugin manager contains plugins, and scope with flags (for plugins communicating)
//
// # Methods:
//
//	AddPlugin(PluginInterface) // Add plugin to manager. Name will got from plugin.Name(). If plugin with same name was already registred -> error. Calling plugin.Init(Scope, *PluginManager), and return result.
//
//	DeletePlugin(string) // Delete plugin from manager. No return if plugin is not found. Calling plugin.Close(), and return result.
//
//	GetPlugin(string) // Get plugin. Return error if not found.
//
//	RunPlugin(string, any) // Run plugin. Call plugin.Run(input) and return result. Return error if plugin is not found.
//
//	CallPluginMethod(string, string, core.Option...) // Call plugin method. Call plugin.Call(method, opts...) and return result. Return error if plugin not found.
//
//	End() // End plugin lifecycle
type PluginManager struct {
	Plugins map[string]PluginInterface
	Scope   core.ScopeType
	flags   []string // You can work with flags with Tools
}

// Create new plugin manager.
//
// Args:
//
// - scope: scope (of engine, or empty), or nil
func NewPluginManager(scope core.ScopeType) *PluginManager {
	if scope == nil {
		scope = make(core.ScopeType)
	}
	return &PluginManager{
		Plugins: make(map[string]PluginInterface),
		Scope:   scope,
	}
}

// Add plugin to manager. Name will got from plugin.Name(). If plugin with same name was
// already registred -> error.  Calling plugin.Init(Scope, *PluginManager), and return result.
func (pm *PluginManager) AddPlugin(plugin PluginInterface) error {
	name := plugin.Name()
	if _, exists := pm.Plugins[name]; exists {
		return fmt.Errorf("Plugin %s already loaded", name)
	}
	pm.Plugins[name] = plugin
	return plugin.Init(pm.Scope, pm)
}

// Delete plugin from manager. No return if plugin is not found. Calling plugin.Close(), and
// return result.
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

// Get plugin. Return error if not found.
func (pm *PluginManager) GetPlugin(name string) (PluginInterface, error) {
	plugin, ok := pm.Plugins[name]
	if !ok {
		return nil, fmt.Errorf("Plugin %s not found", name)
	}
	return plugin, nil
}

// Run plugin. Call plugin.Run(input) and return result. Return error if plugin is not found.
func (pm *PluginManager) RunPlugin(name string, input any) (any, error) {
	plugin, err := pm.GetPlugin(name)
	if err != nil {
		return nil, err
	}
	return plugin.Run(input)
}

// Call plugin method. Call plugin.Call(method, opts...) and return result. Return error if plugin not found.
func (pm *PluginManager) CallPluginMethod(name, method string, opts ...core.Option) (any, error) {
	plugin, err := pm.GetPlugin(name)
	if err != nil {
		return nil, err
	}
	return plugin.Call(method, opts...)
}

// End plugin lifecycle
func (pm *PluginManager) End() (err error) {
	for plugin := range pm.Plugins {
		err = pm.DeletePlugin(plugin)
		if err != nil {
			return
		}
	}
	return
}
