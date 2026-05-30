package lc

import "github.com/pt-main/lc/system"

const Version = "0.1.1"

func NewEngine(generator_res_type int, pipeline []string) *system.Engine {
	return &system.Engine{
		Scope:     make(system.ScopeType),
		Commands:  make(map[string]system.CommandMeta),
		Generator: *system.NewGenerator(generator_res_type, pipeline),
		Event:     *system.NewEvents(),
	}
}
