package lc

import (
	"github.com/pt-main/lc/byteParsing"
	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/events"
	"github.com/pt-main/lc/stringParsing"
)

const Version = "0.9.8"

// NewStringEngine creates a ready-to-use string-based engine.
// Parameters:
//
//	generator_res_type – core.StringResType (usually) for text generation.
//	pipeline – ordered list of generation points (e.g., []string{"pre","main"}).
//	add_default_events – if true, registers standard parsing and call events.
//	parser – an implementation of stringParsing.ParserInterface.
//
// Returns a StringEngine with empty command map and initialized UEP.
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

// NewByteEngine creates a byte-oriented engine for binary formats or bytecode.
// The endianess parameter (e.g., bytecode.LittleEndian) is stored in scope.
// It registers default events when add_default_events is true.
// The parser must implement byteParsing.ParserInterface.
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
