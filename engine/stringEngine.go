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
type stringCommandMeta = core.CommandMeta[StringEngineInterface, stringParsing.ParsedNode]

// StringEngine is the core for text‑based languages. It holds command
// definitions, a parser, and universal engine parameters (UEP) that include
// generator, events, scope, and logger. The Process method drives compilation.
type StringEngine struct {
	Commands map[string]stringCommandMeta
	Parser   stringParser
	mu       sync.RWMutex
	UEP      *core.UniversalEngineParams
}

// Process executes the compilation pipeline for a string input.
// It stores the input in scope[public.StringEngineScopeInput], then calls the
// StringParseEvent (to parse into []ParsedNode) and StringCallEvent
// (to dispatch commands). Any error stops execution.
//
// Err errors.StringEngineProcessError1 | errors.StringEngineProcessError2.
// (cause from 'CallEvents')
func (e *StringEngine) Process(input string) core.ErrorInterface {
	e.UEP.Scope[public.StringEngineScopeInput] = input
	err1 := e.UEP.Event.CallEvents(&core.EventInput{
		Input: e,
	}, public.StringParseEvent, false)
	if err1 != nil {
		return core.Wrap(errors.StringEngineProcessError1, err1, core.GetRealErrorReverse(err1))
	}
	err2 := e.UEP.Event.CallEvents(&core.EventInput{
		Input: e,
	}, public.StringCallEvent, false)
	if err2 != nil {
		return core.Wrap(errors.StringEngineProcessError2, err2, core.GetRealErrorReverse(err2))
	}
	return nil
}

func (e *StringEngine) NewCommandFull(cmd_switch string,
	handler core.CommandType[StringEngineInterface, stringParsing.ParsedNode], doc string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.Commands[cmd_switch] = stringCommandMeta{
		Handler: handler,
		Doc:     doc,
	}
}

func (e *StringEngine) GetParser() stringParser {
	return e.Parser
}

// For interface. o.Input string = doc
func (e *StringEngine) NewCommand(cmd_switch string,
	handler core.CommandType[StringEngineInterface, stringParsing.ParsedNode],
	o *core.SimpleInput) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	doc, ok := o.Input.(string)
	if !ok {
		return core.Err(errors.CorePackageSystemError, "Invalid input: 'o.Input' must be string")
	}
	e.Commands[cmd_switch] = stringCommandMeta{
		Handler: handler,
		Doc:     doc,
	}
	return nil
}

// For interface
func (e *StringEngine) GetCommands() map[string]stringCommandMeta {
	return e.Commands
}

// For interface
func (e *StringEngine) GetUep() *core.UniversalEngineParams {
	return e.UEP
}
