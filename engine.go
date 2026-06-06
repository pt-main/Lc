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

func (e *engineUniversal) ProcessString(input string) error {
	if e.Type != StringEngineType {
		return errors.New("can't process string in byte engine")
	}
	return e.StringEngine.Process(input)
}

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

func (e *engineUniversal) NewCommandString(cmdSwitch string, handler core.CommandType, doc string) error {
	if e.Type != StringEngineType {
		return errors.New("can't add string command to byte engine")
	}
	e.StringEngine.NewCommand(cmdSwitch, handler, doc)
	return nil
}
