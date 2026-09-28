package plugin

import "github.com/pt-main/lc/v2/engine/core"

type PluginInterface interface {
	Name() string
	Init(scope core.ScopeType, pm *PluginManager) error
	Close() error
	Call(string, ...core.Option) (any, error)
	Run(input any) (any, error)
}
