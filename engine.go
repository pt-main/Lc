package lc

import (
	"errors"

	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
)

const (
	ByteEngineType = iota
	StringEngineType
)

type engineAnother struct {
	opcode_counter int
}

type engineUniversal struct {
	Type         int
	StringEngine *engine.StringEngine
	ByteEngine   *engine.ByteEngine
	another      engineAnother
}

// ProcessString feeds a string input into the engine.
// It works only for engines of type StringEngineType; otherwise returns an error.
// Internally triggers the parse and call events, executing registered handlers.
func (e *engineUniversal) ProcessString(input string) error {
	if e.Type != StringEngineType {
		return errors.New("can't process string in byte engine")
	}
	return e.StringEngine.Process(input)
}

// ProcessBytes feeds a byte slice into the engine (ByteEngineType only).
// The input is passed via scope under key "input_[]byte", then parsed and processed.
func (e *engineUniversal) ProcessBytes(input []byte) error {
	if e.Type != ByteEngineType {
		return errors.New("can't process bytes in string engine")
	}
	return e.ByteEngine.Process(input)
}

func (e *engineUniversal) GetUEP() core.UniversalEngineParams {
	if e.Type == StringEngineType {
		return e.StringEngine.UEP
	}
	return e.ByteEngine.UEP
}

// NewCommandByte registers a bytecode command identified by an opcode.
// If opcode == -1, the engine automatically assigns the next available opcode.
// handler receives (*ByteEngine, ParsedBytes).
// doc documents the instruction.
func (e *engineUniversal) NewCommandByte(opcode int, handler core.CommandType, doc string) error {
	if e.Type != ByteEngineType {
		return errors.New("can't add byte command to string engine")
	}

	finalOpcode := opcode
	if opcode == -1 {
		finalOpcode = e.another.opcode_counter
		e.another.opcode_counter++
	} else {
		e.another.opcode_counter = max(opcode, e.another.opcode_counter)
	}

	e.ByteEngine.NewCommand(finalOpcode, handler, doc)
	return nil
}

// NewCommandString registers a text-based command in a StringEngine.
// cmdSwitch is the command name (e.g., "print"). handler must have signature
// func([]interface{}) error where arguments are (*StringEngine, ParsedNode).
// doc is an optional documentation string.
func (e *engineUniversal) NewCommandString(cmdSwitch string, handler core.CommandType, doc string) error {
	if e.Type != StringEngineType {
		return errors.New("can't add string command to byte engine")
	}
	e.StringEngine.NewCommand(cmdSwitch, handler, doc)
	return nil
}
