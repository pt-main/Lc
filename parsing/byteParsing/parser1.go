package byteParsing

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/tooling/bytecode"
)

type Parser1Config struct {
	GConfig bytecode.GenerationConfig
	Shifter bytecode.Shift
}

// Parser1 decodes a binary stream according to a fixed‑length field layout.
// Each bytecode instruction consists of:
//
//	command (CommandBytelen bytes)
//	argscount (ArgscountBytelen bytes)
//	for each argument: arglen (ArglenBytelen bytes) followed by arg data.
//
// Endianess (Little/BigEndian) is used to decode integer fields.
type Parser1 struct {
	Config Parser1Config
}

// Parse reads the byte slice and returns a slice of ParsedBytes.
// Each ParsedBytes contains the raw command bytes, the raw arguments,
// and the original slice of the whole instruction. The ShiftStruct utility
// is used internally for safe bounds checking. Returns error on malformed data.
func (p *Parser1) Parse(code []byte, opts ...*parsing.ParseOption) (result []ParsedBytes, err error) {
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
	log := func(text string) {
		text = "\n" + text
		if len(opts) > 0 {
			logger := opts[0].UEP.Logger
			if logger != nil {
				logger.PrintLog("parsing", text)
			}
		}
	}
	log("start parsing code " + fmt.Sprintf("%s", code))
	u := bytecode.Utils{}
	_Idx := p.Config.Shifter.Idx
	p.Config.Shifter.Idx = idx
	p.Config.Shifter.Code = code
	shift := p.Config.Shifter.ShiftError
	for *idx < len(code) {
		idx_start := idx
		var command []byte
		command, err = shift(p.Config.GConfig.CommandBytelen)
		if err != nil {
			return
		}
		var argscountBytes []byte
		argscountBytes, err = shift(p.Config.GConfig.ArgscountBytelen)
		if err != nil {
			return
		}
		argscount := u.BytesToInt(argscountBytes, p.Config.GConfig.Endianess)
		args := [][]byte{}
		log(fmt.Sprintf("cmd %v, argscount %v", command, argscount))
		for range argscount {
			var arglenBytes []byte
			arglenBytes, err = shift(p.Config.GConfig.ArglenBytelen)
			if err != nil {
				return
			}
			arglen := u.BytesToInt(arglenBytes, p.Config.GConfig.Endianess)
			if arglen == 0 {
				return nil, errors.New("Can't form args with 0 argument length")
			}
			log(fmt.Sprintf("arglen %v", arglen))
			var arg []byte
			arg, err = shift(arglen)
			if err != nil {
				return
			}
			log(fmt.Sprintf("arg %v", arglen))
			args = append(args, arg)
		}
		raw := code[*idx_start:*idx]
		result = append(result, ParsedBytes{
			Switch:   command,
			Args:     args,
			Raw:      raw,
			Metadata: make(map[string]interface{}),
		})
	}
	p.Config.Shifter.Idx = _Idx
	return result, nil
}

func (p *Parser1) String() string {
	return "lc/parsing/byteParsing/Parser1"
}
