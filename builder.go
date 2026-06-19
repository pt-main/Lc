package lc

import (
	"context"
	"errors"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/tooling/bytecode"
	"github.com/pt-main/lc/tooling/plugin"
)

const (
	EuPtrPluginsScope = "EuPtr *EngineUniversal"
)

type stringParser parsing.ParserInterface[string, stringParsing.ParsedNode]
type byteParser parsing.ParserInterface[[]byte, byteParsing.ParsedBytes]

// EngineBuilder is a fluent builder for constructing universal engines.
// It allows to configure pipeline stages, event handling, logging,
// custom parsers, scope variables, and byte order before calling Build().
// Use NewEngineBuilder to create a builder instance.
type EngineBuilder struct {
	engineType       int
	pipeline         []string
	addDefaultEvents bool
	logger           *core.Logger
	scope            core.ScopeType
	stringParser     stringParser
	byteParser       byteParser
	endianess        int
	colorEnabled     bool
	context          context.Context
	pm               bool
	plugins          []*plugin.Plugin
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
func NewEngineBuilder(engineType int) *EngineBuilder {
	return &EngineBuilder{
		engineType:       engineType,
		pipeline:         []string{"main"},
		addDefaultEvents: true,
		endianess:        bytecode.LittleEndian,
		scope:            make(core.ScopeType),
		context:          context.Background(),
	}
}

func (b *EngineBuilder) WithPipeline(pipeline []string) *EngineBuilder {
	b.pipeline = pipeline
	return b
}

func (b *EngineBuilder) WithContext(context context.Context) *EngineBuilder {
	b.context = context
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

func (b *EngineBuilder) WithColors() *EngineBuilder {
	b.colorEnabled = true
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

func (b *EngineBuilder) WithEndianess(endianess int) *EngineBuilder {
	b.endianess = endianess
	return b
}

func (b *EngineBuilder) WithPluginManager(plugins ...*plugin.Plugin) *EngineBuilder {
	b.pm = true
	b.plugins = plugins
	return b
}

// Build constructs and returns an EngineUniversal or an error if
// required components are missing (e.g., a string parser for a StringEngine).
// The returned engineUniversal can process strings or bytes depending
// on its type and provides methods to register commands.
func (b *EngineBuilder) Build() (*EngineUniversal, error) {
	var eu *EngineUniversal
	switch b.engineType {
	case StringEngineType:
		if b.stringParser == nil {
			return nil, errors.New("string parser is required for StringEngine")
		}
		strEngine := NewStringEngine(
			core.StringResType,
			b.pipeline,
			b.addDefaultEvents,
			b.stringParser,
			b.colorEnabled,
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
			Type:           StringEngineType,
			StringEngine:   strEngine,
			opcode_counter: 0,
		}

	case ByteEngineType:
		if b.byteParser == nil {
			return nil, errors.New("byte parser is required for ByteEngine")
		}
		byteEngine := NewByteEngine(
			core.ByteResType,
			b.pipeline,
			b.addDefaultEvents,
			b.byteParser,
			b.endianess,
			b.colorEnabled,
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
			Type:           ByteEngineType,
			ByteEngine:     byteEngine,
			opcode_counter: 0,
		}

	default:
		return nil, errors.New("EngineBuilder.Build: unknown engine type")
	}
	pm := &plugin.PluginManager{
		Plugins: make(map[string]plugin.PluginInterface),
		Scope:   core.ScopeType{EuPtrPluginsScope: eu},
	}
	if b.pm {
		if b.plugins != nil {
			for _, plugin := range b.plugins {
				err := pm.AddPlugin(plugin)
				if err != nil {
					return nil, errors.New("EngineBuilder.Build: " + err.Error())
				}
			}
			for k, v := range eu.GetUEP().Scope {
				pm.Scope[k] = v
			}
		}
	}
	eu.Plugins = pm
	eu.GetUEP().Scope[PluginManagerEuScope] = pm
	return eu, nil
}
