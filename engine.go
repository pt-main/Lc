package lc

import (
	"context"
	"errors"
	"fmt"

	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
	lcplugin "github.com/pt-main/lc/tooling/plugin"
)

type EngineUniversal struct {
	Plugins        *lcplugin.PluginManager
	Type           public.EngineType
	StringEngine   engine.EngineInterface[string, string, stringParsing.ParsedNode]
	ByteEngine     engine.EngineInterface[int, []byte, byteParsing.ParsedBytes]
	opcode_counter int
	Context        context.Context
	CtxCancelCause context.CancelCauseFunc
	ended          bool
}

func (e *EngineUniversal) ProcessStringWithCtx(input string, ctx context.Context) error {
	if err := e.CheckEnded(); err != nil {
		return err
	}
	if e.Type != public.StringEngineType {
		return errors.New("Can't process string in byte engine")
	}
	uep, _ := e.GetUEP()
	uep.Context = ctx
	return e.StringEngine.Process(input)
}

func (e *EngineUniversal) ProcessBytesWithCtx(input []byte, ctx context.Context) error {
	if err := e.CheckEnded(); err != nil {
		return err
	}
	if e.Type != public.ByteEngineType {
		return errors.New("Can't process bytes in string engine")
	}
	uep, _ := e.GetUEP()
	uep.Context = ctx
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

func (e *EngineUniversal) GetUEP() (*core.UniversalEngineParams, error) {
	if err := e.CheckEnded(); err != nil {
		return nil, err
	}
	if e.Type == public.StringEngineType {
		return e.StringEngine.GetUep(), nil
	}
	return e.ByteEngine.GetUep(), nil
}

// NewCommandByte registers a bytecode command identified by an opcode.
// If opcode == -1, the engine automatically assigns the next available opcode.
// handler receives (*ByteEngine, ParsedBytes).
func (e *EngineUniversal) NewCommandByte(
	opcode int, handler core.CommandType[engine.ByteEngineInterface, byteParsing.ParsedBytes], name string,
	autoByecodeIdxShift bool,
) error {
	if err := e.CheckEnded(); err != nil {
		return err
	}
	if e.Type != public.ByteEngineType {
		return errors.New("Can't add byte command to string engine")
	}
	finalOpcode := opcode
	if opcode == -1 {
		finalOpcode = e.opcode_counter
		e.opcode_counter++
	} else {
		e.opcode_counter = max(opcode, e.opcode_counter)
	}

	e.ByteEngine.NewCommand(finalOpcode, handler, &core.SimpleInput{
		Input:  name,
		Option: &core.Option{Flags: []string{engine.AutoshiftNewCommandFlag}},
	})
	return nil
}

// NewCommandString registers a text-based command in a StringEngine.
// cmdSwitch is the command name (e.g., "print"). handler must have signature
// func([]interface{}) error where arguments are (*StringEngine, ParsedNode).
// doc is an optional documentation string.
func (e *EngineUniversal) NewCommandString(
	cmdSwitch string, handler core.CommandType[engine.StringEngineInterface, stringParsing.ParsedNode], doc string,
) error {
	if err := e.CheckEnded(); err != nil {
		return err
	}
	if e.Type != public.StringEngineType {
		return errors.New("Can't add string command to byte engine")
	}
	e.StringEngine.NewCommand(cmdSwitch, handler, &core.SimpleInput{
		Input: doc,
	})
	return nil
}

// End - function for stop engines lifecycle.
func (e *EngineUniversal) End() (err error) {
	err = e.CheckEnded()
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("[%v] Panic recovered: %v. ", err, r)
		}
	}()
	if e.Plugins != nil {
		err = fmt.Errorf("[%v] Plugin error: %v. ", err, e.Plugins.End())
	}
	if e.Context.Err() != nil {
		e.CtxCancelCause(fmt.Errorf("EngineUniversal: lifecycle end."))
	}
	e.ByteEngine = nil
	e.StringEngine = nil
	return
}

func (e *EngineUniversal) CheckEnded() (err error) {
	if e.ended {
		err = fmt.Errorf("EngineUniversal: lifecycle ended.")
	}
	return err
}
