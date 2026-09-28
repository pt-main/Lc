package lc

import (
	"context"

	"github.com/pt-main/lc/v2/engine"
	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/engine/events"
	"github.com/pt-main/lc/v2/parsing/byteParsing"
	"github.com/pt-main/lc/v2/parsing/stringParsing"
	"github.com/pt-main/lc/v2/public"
)

const Version = "2.0.0"

// NewStringEngine creates a ready-to-use string-based engine with an empty
// command map and an initialized UEP. addDefaultEvents registers the standard
// parsing and call events.
func NewStringEngine(
	generatorResType public.ResType,
	pipeline []string,
	addDefaultEvents bool,
	parser stringParser,
	ctx context.Context,
) *engine.StringEngine {
	e := core.NewEvents(ctx)
	if addDefaultEvents {
		de := events.DefaultEvents{}
		e.NewEvent(public.StringParseEvent, de.StringParsingEvent)
		e.NewEvent(public.StringCallEvent, de.StringCallEvent)
		e.NewEvent(public.StringCallCallLoopEvent, de.StringCallLoopEvent)
	}
	uep, _ := core.NewUniversalEngineParams(core.NewGenerator(generatorResType, pipeline),
		e, make(core.ScopeType), core.NewLogger(""), ctx)
	return &engine.StringEngine{
		UEP:      uep,
		Commands: make(map[string]core.CommandMeta[engine.StringEngineInterface, stringParsing.ParsedNode]),
		Parser:   parser,
	}
}

// NewByteEngine creates a byte-oriented engine for binary formats or bytecode.
// The endianness is stored in scope, and addDefaultEvents registers the
// standard parsing and call events.
func NewByteEngine(
	generatorResType public.ResType,
	pipeline []string,
	addDefaultEvents bool,
	parser byteParser,
	endianness public.EndianType,
	ctx context.Context,
) *engine.ByteEngine {
	idx := 0
	e := core.NewEvents(ctx)
	if addDefaultEvents {
		de := events.DefaultEvents{}
		e.NewEvent(public.ByteParseEvent, de.ByteParsingEvent)
		e.NewEvent(public.ByteCallEvent, de.ByteCallEvent)
		e.NewEvent(public.ByteCallHotloopEvent, de.ByteCallHotLoopEvent)
	}
	uep, _ := core.NewUniversalEngineParams(core.NewGenerator(
		generatorResType, pipeline,
	), e, core.ScopeType{
		public.ByteEngineScopeEndianness:  endianness,
		public.ByteEngineScopeBytecodeIdx: &idx,
	}, core.NewLogger(""), ctx)
	return &engine.ByteEngine{
		UEP:                    uep,
		AutoBytecodeIndexShift: make(map[int]bool),
		Commands:               make(map[int]core.CommandMeta[engine.ByteEngineInterface, byteParsing.ParsedBytes]),
		Parser:                 parser,
	}
}

// NewAstEngine creates an AST engine that dispatches over the whole parsed
// tree. canNodeBeUnknown and canMainNodeBeUnknown decide whether an
// unregistered node type is an error or is silently skipped.
func NewAstEngine(
	generatorResType public.ResType,
	pipeline []string,
	addDefaultEvents bool,
	parser stringParser,
	ctx context.Context,
	canNodeBeUnknown,
	canMainNodeBeUnknown bool,
) *engine.AstEngine {
	e := core.NewEvents(ctx)
	if addDefaultEvents {
		de := events.DefaultEvents{}
		e.NewEvent(public.StringParseEvent, de.AstParsingEvent)
	}
	uep, _ := core.NewUniversalEngineParams(core.NewGenerator(
		generatorResType, pipeline,
	), e, core.ScopeType{}, core.NewLogger(""), ctx)
	return &engine.AstEngine{
		UEP:                  uep,
		Parser:               parser,
		Commands:             make(map[string]core.CommandMeta[engine.EngineInterface[string, string, stringParsing.ParsedNode], stringParsing.ParsedNode]),
		CanBeUnknown:         canNodeBeUnknown,
		CanMainNodeBeUnknown: canMainNodeBeUnknown,
		AstCommandCtx:        make(map[string]*engine.AstCommandCtx),
	}
}
