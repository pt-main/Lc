package engine

import (
	"sync"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
)

const AutoshiftNewCommandFlag = "autoShift"

type byteParser = parsing.ParserInterface[[]byte, byteParsing.ParsedBytes]
type byteCmdType = core.CommandType[ByteEngineInterface, byteParsing.ParsedBytes]
type byteCommandMeta = core.CommandMeta[ByteEngineInterface, byteParsing.ParsedBytes]

// ByteEngine handles binary inputs. Commands are indexed by integer opcodes.
// It uses a byte parser to decode raw bytes into ParsedBytes structures.
// The Process method triggers ByteParseEvent and ByteCallEvent in order.
type ByteEngine struct {
	Commands               map[int]byteCommandMeta
	Parser                 byteParser
	AutoBytecodeIndexShift map[int]bool
	UEP                    *core.UniversalEngineParams
	mu                     sync.RWMutex
}

// Process transforms a byte slice by parsing it and invoking the registered
// bytecode handlers.
//
// Err errors.ByteEngineProcessError1 | errors.ByteEngineProcessError2.
// (cause from 'CallEvents')
func (e *ByteEngine) Process(input []byte) core.ErrorInterface {
	e.UEP.Scope[public.ByteEngineScopeInput] = input
	err1 := e.UEP.Event.CallEvents(&core.EventInput{
		Input: e,
	}, public.ByteParseEvent, false)
	if err1 != nil {
		return core.Wrap(errors.ByteEngineProcessError1, err1, core.GetRealError(err1))
	}
	err2 := e.UEP.Event.CallEvents(&core.EventInput{
		Input: e,
	}, public.ByteCallEvent, false)
	if err2 != nil {
		return core.Wrap(errors.ByteEngineProcessError2, err1, core.GetRealError(err1))
	}
	return nil
}

// Your handler MUST shift bytecode index if autoBytecodeIndexShift false!
//
// Usually it's like a:
//
//	AddToBytecodeIdx(1) // next instruction
//
// Or:
//
//	SetBytecodeIdx(10) // jump
//	AddToBytecodeIdx(-1) // prev instruction
func (e *ByteEngine) NewCommandFull(
	cmd_switch int, handler core.CommandType[ByteEngineInterface, byteParsing.ParsedBytes],
	name string, autoBytecodeIndexShift bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.Commands[cmd_switch] = core.CommandMeta[ByteEngineInterface, byteParsing.ParsedBytes]{
		Handler: handler,
		Doc:     name,
	}
	e.AutoBytecodeIndexShift[cmd_switch] = autoBytecodeIndexShift
}

// For interface. o.Option.Flags[AutoshiftNewCommandFlag] = autoBytecodeIndexShift, o.Input string = name
//
// Err errors.CorePackageSystemError.
func (e *ByteEngine) NewCommand(
	cmd_switch int, handler byteCmdType,
	o *core.SimpleInput) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	name, ok := o.Input.(string)
	if !ok {
		return core.Err(errors.CorePackageSystemError, "Invalid input: 'o.Input' must be string")
	}
	autoBytecodeIndexShift := o.Option.HasFlag(AutoshiftNewCommandFlag)
	e.Commands[cmd_switch] = core.CommandMeta[ByteEngineInterface, byteParsing.ParsedBytes]{
		Handler: handler,
		Doc:     name,
	}
	e.AutoBytecodeIndexShift[cmd_switch] = autoBytecodeIndexShift
	return nil
}

// For interface
func (e *ByteEngine) GetCommands() map[int]byteCommandMeta {
	return e.Commands
}

// For interface
func (e *ByteEngine) GetUep() *core.UniversalEngineParams {
	return e.UEP
}

func (e *ByteEngine) GetParser() byteParser {
	return e.Parser
}

func (e *ByteEngine) AddToBytecodeIdx(n int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	*e.UEP.Scope[public.ByteEngineScopeBytecodeIdx].(*int) += n
}

func (e *ByteEngine) SetBytecodeIdx(n int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	*e.UEP.Scope[public.ByteEngineScopeBytecodeIdx].(*int) = n
}

func (e *ByteEngine) GetBytecodeIdx() (*int, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	_idx, ok := e.UEP.Scope[public.ByteEngineScopeBytecodeIdx]
	if !ok {
		return nil, core.Err(errors.CorePackageSystemError, "Can't get bytecode index: invalid scope")
	}
	idx, ok := _idx.(*int)
	if !ok {
		return nil, core.Err(errors.CorePackageSystemError, "Can't get bytecode index: invalid interface in scope")
	}
	return idx, nil
}
