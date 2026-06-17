package lc

import (
	"context"
	"errors"

	"github.com/pt-main/lc/byteParsing"
	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/stringParsing"
)

const (
	ByteEngineType = iota
	StringEngineType
)

type EngineUniversal struct {
	Type           int
	StringEngine   *engine.StringEngine
	ByteEngine     *engine.ByteEngine
	opcode_counter int
	Context        context.Context
}

func (e *EngineUniversal) ProcessStringWithCtx(input string, ctx context.Context) error {
	if e.Type != StringEngineType {
		return errors.New("can't process string in byte engine")
	}
	e.GetUEP().Context = ctx
	return e.StringEngine.Process(input)
}

func (e *EngineUniversal) ProcessBytesWithCtx(input []byte, ctx context.Context) error {
	if e.Type != ByteEngineType {
		return errors.New("can't process bytes in string engine")
	}
	e.GetUEP().Context = ctx
	return e.ByteEngine.Process(input)
}

// ProcessString feeds a string input into the engine.
// It works only for engines of type StringEngineType; otherwise returns an error.
// Internally triggers the parse and call events, executing registered handlers.
func (e *EngineUniversal) ProcessString(input string) error {
	return e.ProcessStringWithCtx(input, context.Background())
}

// ProcessBytes feeds a byte slice into the engine (ByteEngineType only).
// The input is passed via scope under key "input_[]byte", then parsed and processed.
func (e *EngineUniversal) ProcessBytes(input []byte) error {
	return e.ProcessBytesWithCtx(input, context.Background())
}

func (e *EngineUniversal) GetUEP() *core.UniversalEngineParams {
	if e.Type == StringEngineType {
		return e.StringEngine.UEP
	}
	return e.ByteEngine.UEP
}

// NewCommandByte registers a bytecode command identified by an opcode.
// If opcode == -1, the engine automatically assigns the next available opcode.
// handler receives (*ByteEngine, ParsedBytes).
func (e *EngineUniversal) NewCommandByte(
	opcode int, handler core.CommandType[engine.ByteEngine, byteParsing.ParsedBytes], name string,
	autoByecodeIdxShift bool,
) error {
	if e.Type != ByteEngineType {
		return errors.New("can't add byte command to string engine")
	}

	finalOpcode := opcode
	if opcode == -1 {
		finalOpcode = e.opcode_counter
		e.opcode_counter++
	} else {
		e.opcode_counter = max(opcode, e.opcode_counter)
	}

	e.ByteEngine.NewCommand(finalOpcode, handler, name, autoByecodeIdxShift)
	return nil
}

// NewCommandString registers a text-based command in a StringEngine.
// cmdSwitch is the command name (e.g., "print"). handler must have signature
// func([]interface{}) error where arguments are (*StringEngine, ParsedNode).
// doc is an optional documentation string.
func (e *EngineUniversal) NewCommandString(
	cmdSwitch string, handler core.CommandType[engine.StringEngine, stringParsing.ParsedNode], doc string,
) error {
	if e.Type != StringEngineType {
		return errors.New("can't add string command to byte engine")
	}
	e.StringEngine.NewCommand(cmdSwitch, handler, doc)
	return nil
}
