package engine

import (
	"errors"
	"fmt"

	"github.com/pt-main/lc/byteParsing"
	"github.com/pt-main/lc/engine/core"
)

const (
	ByteEngineScopeEndianess   = "ENDIANESS int"
	ByteEngineScopeBytecodeIdx = "BYECODE_IDX *int"
	ByteEngineScopeInput       = "INPUT []byte"
)

// ByteEngine handles binary inputs. Commands are indexed by integer opcodes.
// It uses a byte parser to decode raw bytes into ParsedBytes structures.
// The Process method triggers ByteParseEvent and ByteCallEvent in order.
type ByteEngine struct {
	Commands map[int]core.CommandMeta
	Parser   byteParsing.ParserInterface
	UEP      core.UniversalEngineParams
}

// Process transforms a byte slice by parsing it and invoking the registered
// bytecode handlers. The parsed result is stored in scope["parsed_[]ParsedBytes"].
func (e *ByteEngine) Process(input []byte) error {
	e.UEP.Scope[ByteEngineScopeInput] = input
	err1 := e.UEP.Event.CallEvents(e, core.ByteParseEvent, false)
	if err1 != nil {
		return fmt.Errorf("Calling 'ByteParseEvent' event error: %v", err1.Error())
	}
	err2 := e.UEP.Event.CallEvents(e, core.ByteCallEvent, false)
	if err2 != nil {
		return fmt.Errorf("Calling 'ByteCallEvent' event error: %v", err2.Error())
	}
	return nil
}

// Your handler MUST shift bytecode index!
//
// Usually it's like a:
//
//	AddToBytecodeIdx(1)
//
// Or:
//
//	SetBytecodeIdx(10) // jump
//
// * if index not shifted your program will be looped
func (e *ByteEngine) NewCommand(cmd_switch int, handler core.CommandType, doc string) {
	e.Commands[cmd_switch] = core.CommandMeta{
		Handler: handler,
		Doc:     doc,
	}
}

func (e *ByteEngine) AddToBytecodeIdx(n int) {
	e.UEP.Scope[ByteEngineScopeBytecodeIdx] = *e.UEP.Scope[ByteEngineScopeBytecodeIdx].(*int) + n
}

func (e *ByteEngine) SetBytecodeIdx(n int) {
	e.UEP.Scope[ByteEngineScopeBytecodeIdx] = &n
}

func (e *ByteEngine) GetBytecodeIdx() (*int, error) {
	_idx, ok := e.UEP.Scope[ByteEngineScopeBytecodeIdx]
	if !ok {
		return nil, errors.New("Can't get bytecode index: invalid scope")
	}
	idx, ok := _idx.(*int)
	if !ok {
		return nil, errors.New("Can't get bytecode index: invalid interface in scope")
	}
	return idx, nil
}
