package engine

import (
	"fmt"
	"sync"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/tap/color"
)

// StringEngine is the core for text‑based languages. It holds command
// definitions, a parser, and universal engine parameters (UEP) that include
// generator, events, scope, and logger. The Process method drives compilation.
type StringEngine struct {
	Commands map[string]core.CommandMeta[StringEngine, stringParsing.ParsedNode]
	Parser   parsing.ParserInterface[string, stringParsing.ParsedNode]
	UEP      *core.UniversalEngineParams
	mu       sync.RWMutex
}

// Process executes the compilation pipeline for a string input.
// It stores the input in scope["input_string"], then calls the
// StringParseEvent (to parse into []ParsedNode) and StringCallEvent
// (to dispatch commands). Any error stops execution.
func (e *StringEngine) Process(input string) error {
	e.UEP.Scope[public.StringEngineScopeInput] = input
	err1 := e.UEP.Event.CallEvents(e, public.StringParseEvent, false)
	if err1 != nil {
		return fmt.Errorf(
			color.Set("[?RD]Process error (1)[?RT]: \n%v"),
			color.Set(err1.Error()))
	}
	err2 := e.UEP.Event.CallEvents(e, public.StringCallEvent, false)
	if err2 != nil {
		return fmt.Errorf(
			color.Set("[?RD]Process error (2)[?RT]: \n%v"),
			color.Set(err2.Error()))
	}
	return nil
}

func (e *StringEngine) NewCommand(cmd_switch string,
	handler core.CommandType[StringEngine, stringParsing.ParsedNode], doc string) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.Commands[cmd_switch] = core.CommandMeta[StringEngine, stringParsing.ParsedNode]{
		Handler: handler,
		Doc:     doc,
	}
}
