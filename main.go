package lc

import (
	"github.com/pt-main/lc/byteParsing"
	"github.com/pt-main/lc/events"
	"github.com/pt-main/lc/stringParsing"
	"github.com/pt-main/lc/system"
)

const Version = "0.6.2"

func NewStringEngine(
	generator_res_type int,
	pipeline []string,
	add_default_events bool,
	parser stringParsing.ParserInterface,
) *system.StringEngine {
	e := &system.StringEngine{
		Scope:     make(system.ScopeType),
		Commands:  make(map[string]system.CommandMeta),
		Generator: *system.NewGenerator(generator_res_type, pipeline),
		Event:     *system.NewEvents(),
		Parser:    parser,
	}
	if add_default_events {
		de := events.DefaultEvents{}
		e.Event.NewEvent(system.ParseEvent, de.StringParsingEvent(parser))
		e.Event.NewEvent(system.ParseEvent, de.StringCallEvent)
	}
	return e
}

func NewByteEngine(
	generator_res_type int,
	pipeline []string,
	add_default_events bool,
	parser byteParsing.ParserInterface,
) *system.ByteEngine {
	return &system.ByteEngine{
		Scope:     make(system.ScopeType),
		Commands:  make(map[int]system.CommandMeta),
		Generator: *system.NewGenerator(generator_res_type, pipeline),
		Event:     *system.NewEvents(),
		Parser:    parser,
	}
}
