package lc

import (
	"github.com/pt-main/lc/byteParsing"
	"github.com/pt-main/lc/events"
	"github.com/pt-main/lc/stringParsing"
	"github.com/pt-main/lc/system"
	"github.com/pt-main/lc/system/core"
)

const Version = "0.8.8"

func NewStringEngine(
	generator_res_type int,
	pipeline []string,
	add_default_events bool,
	parser stringParsing.ParserInterface,
) *system.StringEngine {
	e := core.NewEvents()
	if add_default_events {
		de := events.DefaultEvents{}
		e.NewEvent(core.StringParseEvent, de.StringParsingEvent)
		e.NewEvent(core.StringCallEvent, de.StringCallEvent)
	}
	return &system.StringEngine{
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
) *system.ByteEngine {
	e := core.NewEvents()
	if add_default_events {
		de := events.DefaultEvents{}
		e.NewEvent(core.ByteParseEvent, de.ByteParsingEvent)
		e.NewEvent(core.ByteCallEvent, de.ByteCallEvent)
	}
	return &system.ByteEngine{
		UEP: core.UniversalEngineParams{
			Scope:     core.ScopeType{"endianess": endianess},
			Generator: core.NewGenerator(generator_res_type, pipeline),
			Event:     e,
		},
		Commands: make(map[int]core.CommandMeta),
		Parser:   parser,
	}
}
