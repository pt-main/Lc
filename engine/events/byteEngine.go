package events

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/tooling/bytecode"
)

func (de *DefaultEvents) ByteParsingEvent(events *core.Events, i *core.EventInput) error {
	e, ok := i.Input.(*engine.ByteEngine)
	if !ok {
		return fmt.Errorf("Can't get byte engine: invalid input")
	}
	input, err := core.ScopeGet[[]byte](e.UEP.Scope, public.ByteEngineScopeInput)
	if err != nil {
		return errors.New("Can't get input: " + err.Error())
	}
	nodes, err := e.Parser.Parse(input, &parsing.ParseOption{UEP: e.UEP})
	if err != nil {
		return err
	}
	e.UEP.Scope[public.ByteEngineScopeParsed] = nodes
	return nil
}

type ByteCallAttr struct {
	RawIdx  int
	RawNode byteParsing.ParsedBytes
	Abis    bool
	Handler core.CommandType[engine.ByteEngine, byteParsing.ParsedBytes]
}

func (de *DefaultEvents) ByteCallEventIteration(
	idx *int,
	parsed ByteCallAttr, e *engine.ByteEngine,
) error {
	err := parsed.Handler(e, parsed.RawNode)
	if err != nil {
		return fmt.Errorf("[?BRD]Handler error (at %v)[?RT]:\n  %w", parsed.RawIdx, err)
	}
	if parsed.Abis {
		*idx += 1
	}
	return nil
}

func (de *DefaultEvents) ByteCallPreprocess(
	parsed []byteParsing.ParsedBytes, endianess public.EndianType,
	u bytecode.Utils, abis map[int]bool,
	cmds map[int]core.CommandMeta[engine.ByteEngine, byteParsing.ParsedBytes],
) ([]ByteCallAttr, error) {
	res := []ByteCallAttr{}
	for idx, node := range parsed {
		cmd_switch := u.BytesToInt(node.Switch, endianess)
		handler, ok := cmds[cmd_switch]
		if !ok {
			return nil, errors.New("Can't find bytecode op: " + strconv.Itoa(cmd_switch))
		}
		autoshift, ok := abis[cmd_switch]
		if !ok {
			return nil, errors.New("Command not found in autoshift config.")
		}
		res = append(res, ByteCallAttr{
			RawIdx:  idx,
			RawNode: node,
			Handler: handler.Handler,
			Abis:    autoshift,
		})
	}
	return res, nil
}

func (de *DefaultEvents) ByteCallEvent(events *core.Events, i *core.EventInput) (err error) {
	var last_cmd_switch int
	var idx *int
	var parsed []byteParsing.ParsedBytes
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("[?BRD]Panic recovered[?RT]: \n%v", r)
		}
		if err != nil {
			err = fmt.Errorf(
				"[?RD]Error at [[?RT]cmd:[?YW]%v[?RT], bcIdx:[?YW]%v[?RT], pb:[?YW]%v[?RT][?RD]]:[?RT] \n[?RD]->[?RT]    %v",
				last_cmd_switch, *idx, parsed[*idx], strings.ReplaceAll(err.Error(), "\n", "\n[?RD]->[?RT]    "),
			)
		}
	}()
	e, ok := i.Input.(*engine.ByteEngine)
	if !ok {
		return fmt.Errorf("Can't get byte engine: invalid input")
	}
	_parsed, _ := e.UEP.Scope[public.ByteEngineScopeParsed]
	parsed, ok = _parsed.([]byteParsing.ParsedBytes)
	if !ok {
		return errors.New("Can't start call event. Invalid type of parsed result.")
	}
	u := bytecode.Utils{}
	endianess, ok := e.UEP.Scope[public.ByteEngineScopeEndianess].(public.EndianType)
	if !ok {
		return fmt.Errorf("Can't get endianess: not declarated in scope or invalid value")
	}
	idx, err = core.ScopeGet[*int](e.UEP.Scope, public.ByteEngineScopeBytecodeIdx)
	if err != nil {
		return fmt.Errorf("Can't get bytecode index: \n  %w", err)
	}

	ctx := e.UEP.GetContext()
	cmds := e.Commands
	abis := e.AutoBytecodeIndexShift
	parsed2, err := de.ByteCallPreprocess(parsed, endianess, u, abis, cmds)
	p2len := len(parsed2)
	for *idx < p2len && *idx >= 0 {
		if ctx.Err() != nil {
			err = ctx.Err()
			break
		}
		err = de.ByteCallEventIteration(
			idx, parsed2[*idx], e,
		)
		if err != nil {
			break
		}
	}
	return
}

// ----> ~60 m/s on intel core-i7 2.2ghz
