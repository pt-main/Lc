package events

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
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
	nodes, err := e.Parser.Parse(input)
	if err != nil {
		return err
	}
	e.UEP.Scope[public.ByteEngineScopeParsed] = nodes
	return nil
}

func (de *DefaultEvents) ByteCallEventIteration(
	parsed []byteParsing.ParsedBytes, idx *int,
	u bytecode.Utils, e *engine.ByteEngine,
	endianess public.EndianType,
) (int, error) {
	node := parsed[*(idx)]
	cmd_switch := u.BytesToInt(node.Switch, endianess)
	handler, ok := e.Commands[cmd_switch]
	if !ok {
		return 0, errors.New("Can't find bytecode op: " + strconv.Itoa(cmd_switch))
	}
	err := handler.Handler(e, node)
	if err != nil {
		return 0, errors.New("[?BRD]Handler error:[?RT]\n  " + err.Error())
	}
	val, ok := e.AutoBytecodeIndexShift[cmd_switch]
	if !ok {
		return 0, errors.New("Command not found in autoshift config.")
	}
	if val {
		*idx += 1
	}
	return cmd_switch, nil
}

func (de *DefaultEvents) ByteCallEvent(events *core.Events, i *core.EventInput) (err error) {
	var last_cmd_switch int
	var idx *int
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("[?BRD]Panic recovered[?RT]: \n%v", r)
		}
		if err != nil {
			err = fmt.Errorf(
				"[?RD]Error at [[?RT]cmd:[?YW]%v[?RT], bcIdx:[?YW]%v[?RT][?RD]]:[?RT] \n[?RD]->[?RT]    %v",
				last_cmd_switch, *idx, strings.ReplaceAll(err.Error(), "\n", "\n[?RD]->[?RT]    "),
			)
		}
	}()
	e, ok := i.Input.(*engine.ByteEngine)
	if !ok {
		return fmt.Errorf("Can't get byte engine: invalid input")
	}
	_parsed, _ := e.UEP.Scope[public.ByteEngineScopeParsed]
	parsed, ok := _parsed.([]byteParsing.ParsedBytes)
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
		return fmt.Errorf("Can't get bytecode index: \n  %v", err)
	}

	for *(idx) < len(parsed) && *(idx) >= 0 {
		ctx := e.UEP.GetContext()
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		last_cmd_switch, err = de.ByteCallEventIteration(parsed, idx, u, e, endianess)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return
	}
	return nil
}
