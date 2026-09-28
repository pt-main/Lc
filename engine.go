package lc

import (
	"context"
	goerr "errors"
	"fmt"
	"sync/atomic"

	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
	lcplugin "github.com/pt-main/lc/tooling/plugin"
)

// EngineUniversal is the engine the builder returns: it holds the concrete
// engine of its type, the plugin manager and the lifecycle state.
type EngineUniversal struct {
	Plugins        *lcplugin.PluginManager
	Type           public.EngineType
	StringEngine   engine.EngineInterface[string, string, stringParsing.ParsedNode]
	ByteEngine     engine.EngineInterface[int, []byte, byteParsing.ParsedBytes]
	opcodeCounter  int
	Context        context.Context
	CtxCancelCause context.CancelCauseFunc
	ended          atomic.Bool
}

func (e *EngineUniversal) ProcessStringWithCtx(input string, ctx context.Context) core.ErrorInterface {
	if err := e.CheckEnded(); err != nil {
		return err
	}
	if e.Type != public.StringEngineType {
		return core.Err(errors.CorePackageLcError, "Can't process string in byte engine")
	}
	uep, _ := e.GetUEP()
	uep.Context = ctx
	return e.StringEngine.Process(input)
}

func (e *EngineUniversal) ProcessBytesWithCtx(input []byte, ctx context.Context) core.ErrorInterface {
	if err := e.CheckEnded(); err != nil {
		return err
	}
	if e.Type != public.ByteEngineType {
		return core.Err(errors.CorePackageLcError, "Can't process bytes in string engine")
	}
	uep, _ := e.GetUEP()
	uep.Context = ctx
	return e.ByteEngine.Process(input)
}

// ProcessString feeds a string input into the engine.
// It works only for engines of type StringEngineType; otherwise returns a core.ErrorInterface.
// Internally triggers the parse and call events, executing registered handlers.
func (e *EngineUniversal) ProcessString(input string) core.ErrorInterface {
	if err := e.CheckEnded(); err != nil {
		return err
	}
	if e.Type != public.StringEngineType {
		return core.Err(errors.CorePackageLcError, "Can't process string in byte engine")
	}
	return e.StringEngine.Process(input)
}

// ProcessBytes feeds a byte slice into the engine (ByteEngineType only).
// The input is passed via scope under key "input_[]byte", then parsed and processed.
func (e *EngineUniversal) ProcessBytes(input []byte) core.ErrorInterface {
	if err := e.CheckEnded(); err != nil {
		return err
	}
	if e.Type != public.ByteEngineType {
		return core.Err(errors.CorePackageLcError, "Can't process bytes in string engine")
	}
	return e.ByteEngine.Process(input)
}

func (e *EngineUniversal) GetUEP() (*core.UniversalEngineParams, error) {
	if err := e.CheckEnded(); err != nil {
		return nil, err
	}
	if e.Type == public.StringEngineType {
		return e.StringEngine.GetUep(), nil
	}
	return e.ByteEngine.GetUep(), nil
}

// NewCommandByte registers a bytecode command under the given opcode. An
// opcode of -1 makes the engine assign the next free one.
func (e *EngineUniversal) NewCommandByte(
	opcode int, handler core.CommandType[engine.ByteEngineInterface, byteParsing.ParsedBytes], name string,
	autoBytecodeIdxShift bool,
) error {
	if err := e.CheckEnded(); err != nil {
		return err
	}
	if e.Type != public.ByteEngineType {
		return goerr.New("Can't add byte command to string engine")
	}
	finalOpcode := opcode
	if opcode == -1 {
		finalOpcode = e.opcodeCounter
		e.opcodeCounter++
	} else {
		// An explicit opcode occupies its own value, so the next automatic
		// one has to start above it; max() alone would hand out this opcode
		// again and collide.
		e.opcodeCounter = max(opcode+1, e.opcodeCounter)
	}

	// The flag is only set when the caller asked for it: a command that shifts
	// the bytecode index itself must keep full control of the cursor.
	opt := &core.Option{}
	if autoBytecodeIdxShift {
		opt.Flags = []string{engine.AutoshiftNewCommandFlag}
	}
	return e.ByteEngine.NewCommand(finalOpcode, handler, &core.SimpleInput{
		Input:  name,
		Option: opt,
	})
}

// NewCommandString registers a text-based command under the given name, with
// an optional documentation string.
func (e *EngineUniversal) NewCommandString(
	cmdSwitch string, handler core.CommandType[engine.StringEngineInterface, stringParsing.ParsedNode], doc string,
) error {
	if err := e.CheckEnded(); err != nil {
		return err
	}
	if e.Type != public.StringEngineType {
		return goerr.New("Can't add string command to byte engine")
	}
	return e.StringEngine.NewCommand(cmdSwitch, handler, &core.SimpleInput{
		Input: doc,
	})
}

// End stops the engine lifecycle: it cancels the context, releases the scope
// guards, drops the engines and closes every plugin.
func (e *EngineUniversal) End() (err error) {
	if e.ended.Swap(true) {
		return core.Err(errors.CorePackageLcLifecycleError, "EngineUniversal: lifecycle ended.")
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("EngineUniversal: End panic recovered: %v", r)
		}
	}()

	if e.CtxCancelCause != nil {
		e.CtxCancelCause(goerr.New("EngineUniversal: lifecycle end."))
	}

	// Release the scope guards before dropping the engines, so the registry
	// does not keep the scope maps alive for the life of the process.
	if e.StringEngine != nil {
		if uep := e.StringEngine.GetUep(); uep != nil {
			core.ReleaseScopeGuard(uep.Scope)
		}
	}
	if e.ByteEngine != nil {
		if uep := e.ByteEngine.GetUep(); uep != nil {
			core.ReleaseScopeGuard(uep.Scope)
		}
	}

	e.ByteEngine = nil
	e.StringEngine = nil

	if e.Plugins != nil {
		if perr := e.Plugins.End(); perr != nil {
			return perr
		}
	}
	return nil
}

// CheckEnded reports the lifecycle error when End has already been called.
func (e *EngineUniversal) CheckEnded() core.ErrorInterface {
	if e.ended.Load() {
		return core.Err(errors.CorePackageLcLifecycleError, "EngineUniversal: lifecycle ended.")
	}
	return nil
}
