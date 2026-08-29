package lc

import (
	"context"
	"errors"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/tooling/plugin"
)

type stringParser parsing.ParserInterface[string, stringParsing.ParsedNode]
type byteParser parsing.ParserInterface[[]byte, byteParsing.ParsedBytes]

// EngineBuilder is a fluent builder for constructing universal engines.
// It allows to configure pipeline stages, event handling, logging,
// custom parsers, scope variables, and byte order before calling Build().
// Use NewEngineBuilder to create a builder instance.
type EngineBuilder struct {
	engineType       public.EngineType
	pipeline         []string
	addDefaultEvents bool
	logger           *core.Logger
	scope            core.ScopeType
	stringParser     stringParser
	byteParser       byteParser
	endianess        public.EndianType
	colorEnabled     bool
	context          context.Context
	pm               bool
	plugins          []plugin.PluginInterface
	cancel           context.CancelCauseFunc
	resType          public.ResType
}

// NewEngineBuilder creates a new EngineBuilder for the given engine type.
// engineType must be either ByteEngineType or StringEngineType.
// Defaults: pipeline = []string{"main"}, default events enabled,
// endianess = bytecode.LittleEndian, empty scope.
// Example:
//
//	builder := lc.NewEngineBuilder(lc.StringEngineType).
//	            WithPipeline([]string{"pre","main"}).
//	            WithStringParser(myParser)
func NewEngineBuilder(engineType public.EngineType, resType public.ResType) *EngineBuilder {
	return &EngineBuilder{
		engineType:       engineType,
		resType:          resType,
		pipeline:         []string{"main"},
		addDefaultEvents: true,
		endianess:        public.LittleEndian,
		scope:            make(core.ScopeType),
		context:          context.Background(),
	}
}

func (b *EngineBuilder) WithPipeline(pipeline []string) *EngineBuilder {
	b.pipeline = pipeline
	return b
}

func (b *EngineBuilder) WithContext(ctx context.Context) *EngineBuilder {
	b.context, b.cancel = context.WithCancelCause(ctx)
	return b
}

func (b *EngineBuilder) WithDefaultEvents(add bool) *EngineBuilder {
	b.addDefaultEvents = add
	return b
}

func (b *EngineBuilder) WithLogger(logger *core.Logger) *EngineBuilder {
	b.logger = logger
	return b
}

func (b *EngineBuilder) WithScope(scope core.ScopeType) *EngineBuilder {
	for k, v := range scope {
		b.scope[k] = v
	}
	return b
}

func (b *EngineBuilder) WithStringParser(parser stringParser) *EngineBuilder {
	b.stringParser = parser
	return b
}

func (b *EngineBuilder) WithByteParser(parser byteParser) *EngineBuilder {
	b.byteParser = parser
	return b
}

func (b *EngineBuilder) WithEndianess(endianess public.EndianType) *EngineBuilder {
	b.endianess = endianess
	return b
}

func (b *EngineBuilder) WithPlugins(plugins ...plugin.PluginInterface) *EngineBuilder {
	b.pm = true
	b.plugins = append(b.plugins, plugins...)
	return b
}

// Build constructs and returns an EngineUniversal or an error if
// required components are missing (e.g., a string parser for a StringEngine).
// The returned engineUniversal can process strings or bytes depending
// on its type and provides methods to register commands.
func (b *EngineBuilder) Build() (*EngineUniversal, error) {
	var eu *EngineUniversal
	switch b.engineType {
	case public.StringEngineType:
		if b.stringParser == nil {
			return nil, errors.New("string parser is required for StringEngine")
		}
		strEngine := NewStringEngine(
			b.resType,
			b.pipeline,
			b.addDefaultEvents,
			b.stringParser,
			b.context,
		)
		if b.logger != nil {
			strEngine.UEP.Logger = b.logger
		}
		for k, v := range b.scope {
			strEngine.UEP.Scope[k] = v
		}
		eu = &EngineUniversal{
			Plugins:        &plugin.PluginManager{},
			Type:           b.engineType,
			StringEngine:   strEngine,
			opcode_counter: 0,
		}

	case public.ByteEngineType:
		if b.byteParser == nil {
			return nil, errors.New("byte parser is required for ByteEngine")
		}
		byteEngine := NewByteEngine(
			b.resType,
			b.pipeline,
			b.addDefaultEvents,
			b.byteParser,
			b.endianess,
			b.context,
		)
		if b.logger != nil {
			byteEngine.UEP.Logger = b.logger
		}
		for k, v := range b.scope {
			byteEngine.UEP.Scope[k] = v
		}
		eu = &EngineUniversal{
			Plugins:        &plugin.PluginManager{},
			Type:           b.engineType,
			ByteEngine:     byteEngine,
			opcode_counter: 0,
		}

	default:
		return nil, errors.New("EngineBuilder.Build: unknown engine type")
	}
	pm := &plugin.PluginManager{
		Plugins: make(map[string]plugin.PluginInterface),
		Scope:   core.ScopeType{public.PluginsScopeEuPtr: eu},
	}
	uep, _ := eu.GetUEP()
	if b.pm {
		if b.plugins != nil {
			for _, plugin := range b.plugins {
				err := pm.AddPlugin(plugin)
				if err != nil {
					return nil, errors.New("EngineBuilder.Build: " + err.Error())
				}
			}
			for k, v := range uep.Scope {
				pm.Scope[k] = v
			}
		}
	}
	eu.CtxCancelCause = b.cancel
	eu.ended = false
	eu.Plugins = pm
	uep.Scope[public.EuScopePmPtr] = pm
	return eu, nil
}
