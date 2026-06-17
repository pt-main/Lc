package events

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pt-main/lc/byteParsing"
	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/tooling/bytecode"
)

const ByteEngineScopeParsed = "PARSED []ParsedBytes"

func (de *DefaultEvents) ByteParsingEvent(_e interface{}, events *core.Events) error {
	e, ok := _e.(*engine.ByteEngine)
	if !ok {
		return fmt.Errorf("Can't get byte engine: invalid iput")
	}
	input, ok := e.UEP.Scope[engine.ByteEngineScopeInput].([]byte)
	if !ok {
		return errors.New("No input in scope")
	}
	nodes, err := e.Parser.Parse(input)
	if err != nil {
		return err
	}
	e.UEP.Scope[ByteEngineScopeParsed] = nodes
	return nil
}

func (de *DefaultEvents) ByteCallEvent(_e interface{}, events *core.Events) error {
	e, ok := _e.(*engine.ByteEngine)
	if !ok {
		return fmt.Errorf("Can't get byte engine: invalid input")
	}
	_parsed, _ := e.UEP.Scope[ByteEngineScopeParsed]
	parsed, ok := _parsed.([]byteParsing.ParsedBytes)
	if !ok {
		return errors.New("Can't start call event. Invalid type of parsed result.")
	}
	u := bytecode.Utils{}
	endianess, ok := e.UEP.Scope[engine.ByteEngineScopeEndianess].(int)
	if !ok {
		return fmt.Errorf("Can't get endianess: not declarated in scope or invalid value")
	}
	idx, ok := e.UEP.Scope[engine.ByteEngineScopeBytecodeIdx].(*int)
	if !ok {
		return errors.New("Can't get bytecode index: not declarated in scope or invalid value")
	}
	var err error = nil
	var cmd_switch int
	for *(idx) < len(parsed) && *(idx) >= 0 {
		ctx := e.UEP.GetContext()
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		node := parsed[*(idx)]
		cmd_switch = u.BytesToInt(node.Switch, endianess)
		handler, ok := e.Commands[cmd_switch]

		if ok {
			err = handler.Handler(e, node)
		}
		if err != nil {
			err = errors.New("[?BRD]Handler error:[?RT]\n" + err.Error())
			break
		}
		val, ok := e.AutoBytecodeIndexShift[cmd_switch]
		if !ok {
			err = errors.New("Command not found in autoshift config.")
			break
		}
		if val {
			*idx += 1
		}
	}
	if err != nil {
		return fmt.Errorf(
			"[?RD]Error at [[?RT]cmd:[?YW]%v[?RT], bcIdx:[?YW]%v[?RT][?RD]]:[?RT] \n[?RD]->[?RT]    %v",
			cmd_switch, *idx, strings.ReplaceAll(err.Error(), "\n", "\n[?RD]->[?RT]    "),
		)
	}
	return nil
}
