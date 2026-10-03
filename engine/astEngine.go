package engine

import (
	"slices"
	"sync"

	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/parsing/stringParsing"
	"github.com/pt-main/lc/v2/public"
	"github.com/pt-main/lc/v2/public/errors"
	"github.com/pt-main/lc/v2/tooling/astools"
)

// AstCommandCtxPath is the position of a node inside the tree walk.
type AstCommandCtxPath struct {
	Path  []string
	Nodes []*stringParsing.ParsedNode
}

func MakeAstCommandCtxPath() *AstCommandCtxPath {
	return &AstCommandCtxPath{
		Path:  []string{},
		Nodes: []*stringParsing.ParsedNode{},
	}
}

// AstCommandCtx is the per-command state a handler can inspect and set while
// the tree is walked.
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
		breakif = [][]string{}
	}
	if skipif == nil {
		skipif = [][]string{}
	}
	return &AstCommandCtx{
		Name:            name,
		Path:            MakeAstCommandCtxPath(),
		CurrentChildren: make(map[string][]*stringParsing.ParsedNode),
		BreakIf:         breakif,
		SkipIf:          skipif,
	}
}

func (ctx *AstCommandCtx) FindChildren(command string, pn *stringParsing.ParsedNode, children []string) *AstCommandCtx {
	for _, child := range children {
		ctx.CurrentChildren[child] = append(ctx.CurrentChildren[child], astools.FindChildrenPointers(pn, child)...)
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
	ae.mu.Lock()
	defer ae.mu.Unlock()
	if _, ok := ae.AstCommandCtx[command]; !ok {
		ae.AstCommandCtx[command] = AstMakeCommandCtx(command, nil, nil)
	}
	return ae.AstCommandCtx[command]
}

// Process runs the pipeline for a string input. It stores the input in
// scope[public.StringEngineScopeInput], calls the string parse event to fill
// the parsed nodes, then walks them. Any error stops execution.
//
// Err errors.AstEngineProcessError1 | errors.AstEngineProcessError2.
// (cause from 'CallEvents')
func (ae *AstEngine) Process(input string) core.ErrorInterface {
	ae.UEP.Scope[public.StringEngineScopeInput] = input
	err := ae.UEP.Event.CallEvents(&core.EventInput{
		Input: ae,
	}, public.StringParseEvent, false)
	if err != nil {
		return core.Wrap(errors.AstEngineProcessError1, err, "%s", core.GetRealErrorReverse(err))
	}
	parsed, err := core.ScopeGet[[]stringParsing.ParsedNode](ae.UEP.Scope, public.StringEngineScopeParsed)
	if err != nil {
		return core.Wrap(errors.CorePackageSystemError, err, "Process:ScopeGet: %s", core.GetRealErrorReverse(err))
	}
	err = ae.Work(parsed)
	if err != nil {
		return core.Wrap(errors.AstEngineProcessError2, err, "%s", core.GetRealErrorReverse(err))
	}
	return nil
}

func (ae *AstEngine) Work(parsed []stringParsing.ParsedNode) core.ErrorInterface {
	for i := range parsed {
		if err := ae.WorkIter(&parsed[i]); err != nil {
			return err
		}
	}
	return nil
}

func (ae *AstEngine) HasCommand(isMainCmd bool, command string) core.ErrorInterface {
	_, has := ae.GetCommand(command)
	if has {
		return nil
	}
	if isMainCmd {
		if ae.CanMainNodeBeUnknown {
			return nil
		}
		return core.Err(errors.AstEngineUnknown, "Unregistered node: %v", command).
			WithMeta(core.EMK(0, "string"), command)
	}
	if ae.CanBeUnknown {
		return nil
	}
	return core.Err(errors.AstEngineUnknown, "Unregistered node: %v", command).
		WithMeta(core.EMK(0, "string"), command)
}

func (ae *AstEngine) WorkIter(node *stringParsing.ParsedNode) core.ErrorInterface {
	sw := node.Switch
	if err := ae.HasCommand(false, sw); err != nil {
		return err
	}
	meta, known := ae.GetCommand(sw)
	if !known {
		return nil
	}
	if err := meta.Handler(ae, node); err != nil {
		return err
	}
	ctx := ae.GetCommandCtx(sw)
	ctx.Name = node.Switch
	ae.UEP.Event.CallEvents(&core.EventInput{
		Input: ctx,
	}, public.AstCommandCallEvent, true)
	work := true
	var parent *stringParsing.ParsedNode
	// The documented meaning of these rules is subtree wide, so a node is
	// matched when the rule path is a prefix of the current path.
	hasPathPrefix := func(path []string, paths [][]string) bool {
		for _, p := range paths {
			if len(p) <= len(path) && slices.Equal(p, path[:len(p)]) {
				return true
			}
		}
		return false
	}
	walkErr := astools.WalkWithPath(node, func(pn *stringParsing.ParsedNode, path []string) error {
		// The node itself is already handled above; the walk covers descendants.
		if pn == node {
			return nil
		}
		if hasPathPrefix(path, ctx.BreakIf) {
			work = false
		}
		if hasPathPrefix(path, ctx.SkipIf) || !work {
			return nil
		}
		innerSw := pn.Switch
		if err := ae.HasCommand(false, innerSw); err != nil {
			return err
		}
		innerMeta, innerKnown := ae.GetCommand(innerSw)
		if !innerKnown {
			return nil
		}
		innerCtx := ae.GetCommandCtx(innerSw)
		innerCtx.Name = innerSw
		innerCtx.Parent = parent
		ae.UEP.Event.CallEvents(&core.EventInput{
			Input: innerCtx,
		}, public.AstCommandCallEvent, true)
		if err := innerMeta.Handler(ae, pn); err != nil {
			return err
		}
		parent = pn
		return nil
	})
	if walkErr != nil {
		return core.Wrap(errors.AstEngineHandlerError, walkErr, "%s", core.GetRealErrorReverse(walkErr))
	}
	return nil
}

func (ae *AstEngine) NewCommandFull(cmd_switch string, handler astCommandType, doc string) {
	ae.mu.Lock()
	defer ae.mu.Unlock()
	ae.Commands[cmd_switch] = astCommandMeta{
		Handler: handler,
		Doc:     doc,
	}
}

// NewCommand registers a command for the EngineInterface. o.Input string = doc
func (ae *AstEngine) NewCommand(cmd_switch string, handler astCommandType, o *core.SimpleInput) error {
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

func (ae *AstEngine) GetCommands() map[string]astCommandMeta {
	ae.mu.RLock()
	defer ae.mu.RUnlock()
	res := make(map[string]astCommandMeta, len(ae.Commands))
	for k, v := range ae.Commands {
		res[k] = v
	}
	return res
}

func (ae *AstEngine) GetCommand(cmd_switch string) (astCommandMeta, bool) {
	ae.mu.RLock()
	defer ae.mu.RUnlock()
	cmd, ok := ae.Commands[cmd_switch]
	return cmd, ok
}

func (ae *AstEngine) GetUep() *core.UniversalEngineParams {
	return ae.UEP
}

func (ae *AstEngine) GetParser() stringParser {
	return ae.Parser
}
