package lc

import (
	"context"

	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/engine/events"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/tap/color"
)

const Version = "1.2.0"

// NewStringEngine creates a ready-to-use string-based engine.
// Parameters:
//
//	generator_res_type – core.StringResType (usually) for text generation.
//	pipeline – ordered list of generation points (e.g., []string{"pre","main"}).
//	add_default_events – if true, registers standard parsing and call events.
//	parser – an implementation parser.ParserInterface.
//
// Returns a StringEngine with empty command map and initialized UEP.
func NewStringEngine(
	generator_res_type int,
	pipeline []string,
	add_default_events bool,
	parser stringParser,
	colorEnable bool,
	context context.Context,
) *engine.StringEngine {
	color.ColorEnabled = colorEnable
	e := core.NewEvents(context)
	if add_default_events {
		de := events.DefaultEvents{}
		e.NewEvent(core.StringParseEvent, de.StringParsingEvent)
		e.NewEvent(core.StringCallEvent, de.StringCallEvent)
	}
	uep, _ := core.NewUniversalEngineParams(core.NewGenerator(generator_res_type, pipeline),
		e, make(core.ScopeType), core.NewLogger(""), context)
	return &engine.StringEngine{
		UEP:      uep,
		Commands: make(map[string]core.CommandMeta[engine.StringEngine, stringParsing.ParsedNode]),
		Parser:   parser,
	}
}

// NewByteEngine creates a byte-oriented engine for binary formats or bytecode.
//
// The endianess parameter (e.g., bytecode.LittleEndian) is stored in scope.
//
// It registers default events when add_default_events is true.
//
// The parser must implement paraing.ParserInterface.
func NewByteEngine(
	generator_res_type int,
	pipeline []string,
	add_default_events bool,
	parser byteParser,
	endianess int,
	colorEnable bool,
	context context.Context,
) *engine.ByteEngine {
	color.ColorEnabled = colorEnable
	idx := 0
	e := core.NewEvents(context)
	if add_default_events {
		de := events.DefaultEvents{}
		e.NewEvent(core.ByteParseEvent, de.ByteParsingEvent)
		e.NewEvent(core.ByteCallEvent, de.ByteCallEvent)
	}
	uep, _ := core.NewUniversalEngineParams(core.NewGenerator(
		generator_res_type, pipeline,
	), e, core.ScopeType{
		engine.ByteEngineScopeEndianess:   endianess,
		engine.ByteEngineScopeBytecodeIdx: &idx,
	}, core.NewLogger(""), context)
	return &engine.ByteEngine{
		UEP:                    uep,
		AutoBytecodeIndexShift: make(map[int]bool),
		Commands:               make(map[int]core.CommandMeta[engine.ByteEngine, byteParsing.ParsedBytes]),
		Parser:                 parser,
	}
}
