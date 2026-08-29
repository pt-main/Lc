package engine

import (
	"slices"
	"sync"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/public/errors"
	"github.com/pt-main/lc/tooling/astools"
)

type AstCommandCtxPath struct {
	Path  []string
	Nodes []*stringParsing.ParsedNode
}

func MakeAstCommandCtxPath() *AstCommandCtxPath {
	return &AstCommandCtxPath{
		Path:  make([]string, 0),
		Nodes: make([]*stringParsing.ParsedNode, 0),
	}
}

type AstCommandCtx struct {
	Name            string
	Path            *AstCommandCtxPath
	Parent          *stringParsing.ParsedNode
	CurrentChildren map[string][]*stringParsing.ParsedNode
	BreakIf         [][]string
	SkipIf          [][]string
}

func AstMakeCommandCtx(name string, breakif, skipif [][]string) *AstCommandCtx {
	if breakif == nil {
		breakif = make([][]string, 0)
	}
	if skipif == nil {
		breakif = make([][]string, 0)
	}
	return &AstCommandCtx{
		Name:            name,
		Path:            MakeAstCommandCtxPath(),
		CurrentChildren: make(map[string][]*stringParsing.ParsedNode),
		BreakIf:         breakif,
		SkipIf:          skipif,
		Parent:          nil,
	}
}

func (ctx *AstCommandCtx) FindChildren(command string, pn *stringParsing.ParsedNode, children []string) *AstCommandCtx {
	for _, child := range children {
		for _, findedChild := range astools.FindChildren(pn, child) {
			ctx.CurrentChildren[child] = append(ctx.CurrentChildren[child], &findedChild)
		}
	}
	return ctx
}

type astCommandMeta = core.CommandMeta[AstEngineInterface, stringParsing.ParsedNode]
type astCommandType = core.CommandType[AstEngineInterface, stringParsing.ParsedNode]

type AstEngine struct {
	AstCommandCtx        map[string]*AstCommandCtx
	UEP                  *core.UniversalEngineParams
	Parser               stringParser
	Commands             map[string]astCommandMeta
	CanBeUnknown         bool
	CanMainNodeBeUnknown bool
	mu                   sync.RWMutex
}

func (ae *AstEngine) GetCommandCtx(command string) *AstCommandCtx {
	if _, ok := ae.AstCommandCtx[command]; ok {
		ae.AstCommandCtx[command] = AstMakeCommandCtx(command, nil, nil)
	}
	return ae.AstCommandCtx[command]
}

// Process executes the compilation pipeline for a string input.
// It stores the input in scope[public.StringEngineScopeInput], then calls the
// AstParseEvent (to parse into []ParsedNode) and AstCallEvent
// (to dispatch commands). Any error stops execution.
//
// Err errors.AstEngineProcessError1 | errors.AstEngineProcessError2.
// (cause from 'CallEvents')
func (ae *AstEngine) Process(input string) core.ErrorInterface {
	ae.UEP.Scope[public.StringEngineScopeInput] = input
	err := ae.UEP.Event.CallEvents(&core.EventInput{
		Input: ae,
	}, public.StringParseEvent, false)
	if err != nil {
		return core.Wrap(errors.AstEngineProcessError1, err, core.GetRealErrorReverse(err))
	}
	parsed, err := core.ScopeGet[[]stringParsing.ParsedNode](ae.UEP.Scope, "")
	if err != nil {
		return core.Wrap(errors.CorePackageSystemError, err, "Process:ScopeGet: "+core.GetRealErrorReverse(err))
	}
	err = ae.Work(parsed)
	if err != nil {
		return core.Wrap(errors.AstEngineProcessError2, err, core.GetRealErrorReverse(err))
	}
	return nil
}

func (ae *AstEngine) Work(parsed []stringParsing.ParsedNode) core.ErrorInterface {
	for _, node := range parsed {
		if err := ae.WorkIter(&node); err != nil {
			return err
		}
	}
	return nil
}

func (ae *AstEngine) HasCommand(isMainCmd bool, command string) core.ErrorInterface {
	_, has := ae.Commands[command]
	if !has && ae.CanBeUnknown && !isMainCmd {
		return core.Err("skip", "skip")
	} else if !has && !ae.CanBeUnknown || isMainCmd && !has && !ae.CanMainNodeBeUnknown {
		return core.Err(errors.AstEngineUnknown, "Unregistred node: %v", command).
			WithMeta(core.EMK(0, "string"), command)
	}
	return nil
}

func (ae *AstEngine) WorkIter(node *stringParsing.ParsedNode) core.ErrorInterface {
	sw := node.Switch
	if err := ae.HasCommand(false, sw); err != nil {
		return err
	}
	err := ae.Commands[sw].Handler(ae, node)
	if err != nil {
		return err
	}
	ctx := ae.GetCommandCtx(sw)
	ctx.Name = node.Switch
	ae.UEP.Event.CallEvents(&core.EventInput{
		Input: ctx,
	}, public.AstCommandCallEvent, true)
	work := true
	var parrent *stringParsing.ParsedNode
	containsPath := func(path []string, paths [][]string) bool {
		for _, p := range paths {
			if slices.Equal(p, path) {
				return true
			}
		}
		return false
	}
	err_ := astools.WalkWithPath(node, func(pn *stringParsing.ParsedNode, path []string) error {
		if containsPath(path, ctx.BreakIf) {
			work = false
		}
		if containsPath(path, ctx.SkipIf) || !work {
			return nil
		}
		innersw := pn.Switch
		if err := ae.HasCommand(false, innersw); err != nil {
			return err
		}
		innerctx := ae.GetCommandCtx(innersw)
		innerctx.Name = innersw
		innerctx.Parent = parrent
		ae.UEP.Event.CallEvents(&core.EventInput{
			Input: innerctx,
		}, public.AstCommandCallEvent, true)
		err := ae.Commands[innersw].Handler(ae, pn)
		if err != nil {
			return err
		}
		parrent = pn
		return nil
	})
	if err_ != nil {
		return core.Wrap(errors.AstEngineHandlerError, err, core.GetRealErrorReverse(err))
	}
	return nil
}

func (ae *AstEngine) NewCommandFull(cmd_switch string,
	handler astCommandType,
	doc string) {
	ae.mu.Lock()
	defer ae.mu.Unlock()
	ae.Commands[cmd_switch] = astCommandMeta{
		Handler: handler,
		Doc:     doc,
	}
}

// For interface. o.Input string = doc
func (ae *AstEngine) NewCommand(cmd_switch string,
	handler astCommandType,
	o *core.SimpleInput) error {
	ae.mu.Lock()
	defer ae.mu.Unlock()
	doc, ok := o.Input.(string)
	if !ok {
		return core.Err(errors.CorePackageSystemError, "Invalid input: 'o.Input' must be string")
	}
	ae.Commands[cmd_switch] = astCommandMeta{
		Handler: handler,
		Doc:     doc,
	}
	return nil
}

// For interface
func (ae *AstEngine) GetCommands() map[string]astCommandMeta {
	return ae.Commands
}

// For interface
func (ae *AstEngine) GetUep() *core.UniversalEngineParams {
	return ae.UEP
}

// For interface
func (ae *AstEngine) GetParser() stringParser {
	return ae.Parser
}
