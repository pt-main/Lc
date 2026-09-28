package engine

import (
	"sync"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
)

type stringParser = parsing.ParserInterface[string, stringParsing.ParsedNode]
type stringCommandType = core.CommandType[StringEngineInterface, stringParsing.ParsedNode]
type stringCommandMeta = core.CommandMeta[StringEngineInterface, stringParsing.ParsedNode]

// StringEngine is the core for text-based languages. Process drives the
// parse and call pipeline over the registered commands.
type StringEngine struct {
	Commands map[string]stringCommandMeta
	Parser   stringParser
	mu       sync.RWMutex
	UEP      *core.UniversalEngineParams
}

// Process stores the input in scope[public.StringEngineScopeInput], parses it
// into []ParsedNode and dispatches the commands. Any error stops execution.
//
// Err errors.StringEngineProcessError1 | errors.StringEngineProcessError2.
// (cause from 'CallEvents')
func (e *StringEngine) Process(input string) core.ErrorInterface {
	e.UEP.Scope[public.StringEngineScopeInput] = input
	if err := e.UEP.Event.CallEvents(&core.EventInput{
		Input: e,
	}, public.StringParseEvent, false); err != nil {
		return core.Wrap(errors.StringEngineProcessError1, err, "%s", core.GetRealErrorReverse(err))
	}
	if err := e.UEP.Event.CallEvents(&core.EventInput{
		Input: e,
	}, public.StringCallEvent, false); err != nil {
		return core.Wrap(errors.StringEngineProcessError2, err, "%s", core.GetRealErrorReverse(err))
	}
	return nil
}

func (e *StringEngine) NewCommandFull(cmdSwitch string, handler stringCommandType, doc string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.Commands[cmdSwitch] = stringCommandMeta{
		Handler: handler,
		Doc:     doc,
	}
}

func (e *StringEngine) GetParser() stringParser {
	return e.Parser
}

// NewCommand registers a command for the EngineInterface. o.Input string = doc
func (e *StringEngine) NewCommand(cmdSwitch string, handler stringCommandType, o *core.SimpleInput) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	doc, ok := o.Input.(string)
	if !ok {
		return core.Err(errors.CorePackageSystemError, "Invalid input: 'o.Input' must be string")
	}
	e.Commands[cmdSwitch] = stringCommandMeta{
		Handler: handler,
		Doc:     doc,
	}
	return nil
}

func (e *StringEngine) GetCommands() map[string]stringCommandMeta {
	e.mu.RLock()
	defer e.mu.RUnlock()
	res := make(map[string]stringCommandMeta, len(e.Commands))
	for k, v := range e.Commands {
		res[k] = v
	}
	return res
}

func (e *StringEngine) GetCommand(cmdSwitch string) (stringCommandMeta, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	cmd, ok := e.Commands[cmdSwitch]
	return cmd, ok
}

func (e *StringEngine) GetUep() *core.UniversalEngineParams {
	return e.UEP
}
