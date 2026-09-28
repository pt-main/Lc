package lc

import (
	"context"
	"errors"

	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/parsing"
	"github.com/pt-main/lc/v2/parsing/byteParsing"
	"github.com/pt-main/lc/v2/parsing/stringParsing"
	"github.com/pt-main/lc/v2/public"
	"github.com/pt-main/lc/v2/tooling/plugin"
)

type (
	stringParser parsing.ParserInterface[string, stringParsing.ParsedNode]
	byteParser   parsing.ParserInterface[[]byte, byteParsing.ParsedBytes]
)

// EngineBuilder is a fluent builder for universal engines. It configures
// pipeline stages, event handling, logging, custom parsers, scope variables
// and byte order before Build produces the engine.
type EngineBuilder struct {
	engineType       public.EngineType
	resType          public.ResType
	pipeline         []string
	addDefaultEvents bool
	logger           *core.Logger
	scope            core.ScopeType
	stringParser     stringParser
	byteParser       byteParser
	endianness       public.EndianType
	hasPlugins       bool
	plugins          []plugin.PluginInterface
	context          context.Context
	cancel           context.CancelCauseFunc
}

// NewEngineBuilder creates a builder for the given engine and result type.
// Defaults: pipeline = []string{"main"}, default events enabled,
// endianness = public.LittleEndian, empty scope.
func NewEngineBuilder(engineType public.EngineType, resType public.ResType) *EngineBuilder {
	return &EngineBuilder{
		engineType:       engineType,
		resType:          resType,
		pipeline:         []string{"main"},
		addDefaultEvents: true,
		endianness:       public.LittleEndian,
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

func (b *EngineBuilder) WithEndianness(endianness public.EndianType) *EngineBuilder {
	b.endianness = endianness
	return b
}

func (b *EngineBuilder) WithPlugins(plugins ...plugin.PluginInterface) *EngineBuilder {
	b.hasPlugins = true
	b.plugins = append(b.plugins, plugins...)
	return b
}

// Build constructs the EngineUniversal, or returns an error when a required
// component is missing, such as a parser for the selected engine type.
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
			Plugins:      &plugin.PluginManager{},
			Type:         b.engineType,
			StringEngine: strEngine,
			Context:      b.context,
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
			b.endianness,
			b.context,
		)
		if b.logger != nil {
			byteEngine.UEP.Logger = b.logger
		}
		for k, v := range b.scope {
			byteEngine.UEP.Scope[k] = v
		}
		eu = &EngineUniversal{
			Plugins:    &plugin.PluginManager{},
			Type:       b.engineType,
			ByteEngine: byteEngine,
			Context:    b.context,
		}

	default:
		return nil, errors.New("EngineBuilder.Build: unknown engine type")
	}
	pm := &plugin.PluginManager{
		Plugins: make(map[string]plugin.PluginInterface),
		Scope:   core.ScopeType{public.PluginsScopeEuPtr: eu},
	}
	uep, _ := eu.GetUEP()
	// The manager has to be published before the plugins initialise, otherwise
	// Init sees a nil eu.Plugins and a scope without EuScopePmPtr.
	eu.CtxCancelCause = b.cancel
	eu.ended.Store(false)
	eu.Plugins = pm
	uep.Scope[public.EuScopePmPtr] = pm
	if !b.hasPlugins {
		return eu, nil
	}
	for k, v := range uep.Scope {
		pm.Scope[k] = v
	}
	for _, p := range b.plugins {
		if err := pm.AddPlugin(p); err != nil {
			// Plugins already added hold engine resources, so a failed
			// build must not leave them running.
			_ = pm.End()
			if b.cancel != nil {
				b.cancel(errors.New("EngineBuilder.Build: " + err.Error()))
			}
			return nil, errors.New("EngineBuilder.Build: " + err.Error())
		}
	}
	return eu, nil
}
