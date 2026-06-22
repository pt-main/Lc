package engine

import (
	"errors"
	"fmt"
	"sync"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/tap/color"
)

// ByteEngine handles binary inputs. Commands are indexed by integer opcodes.
// It uses a byte parser to decode raw bytes into ParsedBytes structures.
// The Process method triggers ByteParseEvent and ByteCallEvent in order.
type ByteEngine struct {
	Commands               map[int]core.CommandMeta[ByteEngine, byteParsing.ParsedBytes]
	Parser                 parsing.ParserInterface[[]byte, byteParsing.ParsedBytes]
	AutoBytecodeIndexShift map[int]bool
	UEP                    *core.UniversalEngineParams
	mu                     sync.RWMutex
}

// Process transforms a byte slice by parsing it and invoking the registered
// bytecode handlers. The parsed result is stored in scope["parsed_[]ParsedBytes"].
func (e *ByteEngine) Process(input []byte) error {
	e.UEP.Scope[public.ByteEngineScopeInput] = input
	err1 := e.UEP.Event.CallEvents(e, public.ByteParseEvent, false)
	if err1 != nil {
		return fmt.Errorf(
			color.Set("[?RD]Process error (1)[?RT]: \n%v"),
			color.Set(err1.Error()))
	}
	err2 := e.UEP.Event.CallEvents(e, public.ByteCallEvent, false)
	if err2 != nil {
		return fmt.Errorf(
			color.Set("[?RD]Process error (2)[?RT]: \n%v"),
			color.Set(err2.Error()))
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
func (e *ByteEngine) NewCommand(
	cmd_switch int, handler core.CommandType[ByteEngine, byteParsing.ParsedBytes],
	name string, autoBytecodeIndexShift bool) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.Commands[cmd_switch] = core.CommandMeta[ByteEngine, byteParsing.ParsedBytes]{
		Handler: handler,
		Doc:     name,
	}
	e.AutoBytecodeIndexShift[cmd_switch] = autoBytecodeIndexShift
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
		return nil, errors.New("Can't get bytecode index: invalid scope")
	}
	idx, ok := _idx.(*int)
	if !ok {
		return nil, errors.New("Can't get bytecode index: invalid interface in scope")
	}
	return idx, nil
}
