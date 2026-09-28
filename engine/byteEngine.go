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
type byteCommandType = core.CommandType[ByteEngineInterface, byteParsing.ParsedBytes]
type byteCommandMeta = core.CommandMeta[ByteEngineInterface, byteParsing.ParsedBytes]

// ByteEngine handles binary inputs. Commands are indexed by integer opcodes,
// and Process triggers ByteParseEvent followed by ByteCallEvent.
type ByteEngine struct {
	Commands               map[int]byteCommandMeta
	Parser                 byteParser
	AutoBytecodeIndexShift map[int]bool
	UEP                    *core.UniversalEngineParams
	mu                     sync.RWMutex
}

// Process parses the byte slice and invokes the registered bytecode handlers.
//
// Err errors.ByteEngineProcessError1 | errors.ByteEngineProcessError2.
// (cause from 'CallEvents')
func (e *ByteEngine) Process(input []byte) core.ErrorInterface {
	e.UEP.Scope[public.ByteEngineScopeInput] = input
	if err := e.UEP.Event.CallEvents(&core.EventInput{
		Input: e,
	}, public.ByteParseEvent, false); err != nil {
		return core.Wrap(errors.ByteEngineProcessError1, err, "%s", core.GetRealErrorReverse(err))
	}
	if err := e.UEP.Event.CallEvents(&core.EventInput{
		Input: e,
	}, public.ByteCallEvent, false); err != nil {
		return core.Wrap(errors.ByteEngineProcessError2, err, "%s", core.GetRealErrorReverse(err))
	}
	return nil
}

// NewCommandFull registers a command directly. A handler registered with
// autoBytecodeIndexShift false has to move the instruction index itself,
// through AddToBytecodeIdx or SetBytecodeIdx.
func (e *ByteEngine) NewCommandFull(
	opcode int, handler byteCommandType, name string, autoBytecodeIndexShift bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.Commands[opcode] = byteCommandMeta{
		Handler: handler,
		Doc:     name,
	}
	e.AutoBytecodeIndexShift[opcode] = autoBytecodeIndexShift
}

// NewCommand registers a command for the EngineInterface.
// o.Input string = name, o.Option.Flags[AutoshiftNewCommandFlag] = autoBytecodeIndexShift
//
// Err errors.CorePackageSystemError.
func (e *ByteEngine) NewCommand(opcode int, handler byteCommandType, o *core.SimpleInput) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	name, ok := o.Input.(string)
	if !ok {
		return core.Err(errors.CorePackageSystemError, "Invalid input: 'o.Input' must be string")
	}
	e.Commands[opcode] = byteCommandMeta{
		Handler: handler,
		Doc:     name,
	}
	e.AutoBytecodeIndexShift[opcode] = o.Option.HasFlag(AutoshiftNewCommandFlag)
	return nil
}

func (e *ByteEngine) GetCommands() map[int]byteCommandMeta {
	e.mu.RLock()
	defer e.mu.RUnlock()
	res := make(map[int]byteCommandMeta, len(e.Commands))
	for k, v := range e.Commands {
		res[k] = v
	}
	return res
}

func (e *ByteEngine) GetAutoBytecodeIndexShift() map[int]bool {
	e.mu.RLock()
	defer e.mu.RUnlock()
	res := make(map[int]bool, len(e.AutoBytecodeIndexShift))
	for k, v := range e.AutoBytecodeIndexShift {
		res[k] = v
	}
	return res
}

func (e *ByteEngine) GetCommand(opcode int) (byteCommandMeta, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	cmd, ok := e.Commands[opcode]
	return cmd, ok
}

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
	raw, ok := e.UEP.Scope[public.ByteEngineScopeBytecodeIdx]
	if !ok {
		return nil, core.Err(errors.CorePackageSystemError, "Can't get bytecode index: invalid scope")
	}
	idx, ok := raw.(*int)
	if !ok {
		return nil, core.Err(errors.CorePackageSystemError, "Can't get bytecode index: invalid interface in scope")
	}
	return idx, nil
}
