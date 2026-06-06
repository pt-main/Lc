package lc

import (
	"errors"

	"github.com/pt-main/lc/byteParsing"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/stringParsing"
	"github.com/pt-main/lc/tooling/bytecode"
)

type EngineBuilder struct {
	engineType       int
	pipeline         []string
	addDefaultEvents bool
	logger           *core.Logger
	scope            core.ScopeType
	stringParser     stringParsing.ParserInterface
	byteParser       byteParsing.ParserInterface
	endianess        int
}

func NewEngineBuilder(engineType int) *EngineBuilder {
	return &EngineBuilder{
		engineType:       engineType,
		pipeline:         []string{"main"},
		addDefaultEvents: true,
		endianess:        bytecode.LittleEndian,
		scope:            make(core.ScopeType),
	}
}

func (b *EngineBuilder) WithPipeline(pipeline []string) *EngineBuilder {
	b.pipeline = pipeline
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

func (b *EngineBuilder) WithStringParser(parser stringParsing.ParserInterface) *EngineBuilder {
	b.stringParser = parser
	return b
}

func (b *EngineBuilder) WithByteParser(parser byteParsing.ParserInterface) *EngineBuilder {
	b.byteParser = parser
	return b
}

func (b *EngineBuilder) WithEndianess(endianess int) *EngineBuilder {
	b.endianess = endianess
	return b
}

func (b *EngineBuilder) Build() (*engineUniversal, error) {
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
		)
		if b.logger != nil {
			strEngine.UEP.Logger = b.logger
		}
		for k, v := range b.scope {
			strEngine.UEP.Scope[k] = v
		}
		return &engineUniversal{
			Type:         StringEngineType,
			StringEngine: strEngine,
			another:      engineAnother{opcode_counter: 0},
		}, nil

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
		)
		if b.logger != nil {
			byteEngine.UEP.Logger = b.logger
		}
		for k, v := range b.scope {
			byteEngine.UEP.Scope[k] = v
		}
		return &engineUniversal{
			Type:       ByteEngineType,
			ByteEngine: byteEngine,
			another:    engineAnother{opcode_counter: 0},
		}, nil

	default:
		return nil, errors.New("unknown engine type")
	}
}
