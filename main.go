package lc

import (
	"github.com/pt-main/lc/events"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/system"
)

const Version = "0.2.5"

func NewEngine(
	generator_res_type int,
	pipeline []string,
	add_default_events bool,
	parser parsing.ParserInterface,
) *system.Engine {
	e := &system.Engine{
		Scope:     make(system.ScopeType),
		Commands:  make(map[string]system.CommandMeta),
		Generator: *system.NewGenerator(generator_res_type, pipeline),
		Event:     *system.NewEvents(),
		Parser:    parser,
	}
	if add_default_events {
		de := events.DefaultEvents{}
		e.Event.NewEvent(system.ParseEvent, de.ParsingEvent(parser))
		e.Event.NewEvent(system.ParseEvent, de.CallEvent)
	}
	return e
}
