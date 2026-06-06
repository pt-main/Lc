package lc

import (
	"github.com/pt-main/lc/byteParsing"
	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/events"
	"github.com/pt-main/lc/stringParsing"
)

const Version = "0.9.1"

func NewStringEngine(
	generator_res_type int,
	pipeline []string,
	add_default_events bool,
	parser stringParsing.ParserInterface,
) *engine.StringEngine {
	e := core.NewEvents()
	if add_default_events {
		de := events.DefaultEvents{}
		e.NewEvent(core.StringParseEvent, de.StringParsingEvent)
		e.NewEvent(core.StringCallEvent, de.StringCallEvent)
	}
	return &engine.StringEngine{
		UEP: core.UniversalEngineParams{
			Scope:     make(core.ScopeType),
			Generator: core.NewGenerator(generator_res_type, pipeline),
			Event:     e,
		},
		Commands: make(map[string]core.CommandMeta),
		Parser:   parser,
	}
}

func NewByteEngine(
	generator_res_type int,
	pipeline []string,
	add_default_events bool,
	parser byteParsing.ParserInterface,
	endianess int,
) *engine.ByteEngine {
	e := core.NewEvents()
	if add_default_events {
		de := events.DefaultEvents{}
		e.NewEvent(core.ByteParseEvent, de.ByteParsingEvent)
		e.NewEvent(core.ByteCallEvent, de.ByteCallEvent)
	}
	return &engine.ByteEngine{
		UEP: core.UniversalEngineParams{
			Scope:     core.ScopeType{"endianess": endianess},
			Generator: core.NewGenerator(generator_res_type, pipeline),
			Event:     e,
		},
		Commands: make(map[int]core.CommandMeta),
		Parser:   parser,
	}
}
