package byteParsing

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pt-main/lc/parsing"
	"github.com/pt-main/lc/public"
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
	var last_cmd_switch []byte = []byte{0}
	var raw []byte
	_idx := 0
	var idx *int = &_idx
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("[?BRD]Panic recovered[?RT]: \n%v", r)
		}
		if err != nil {
			err = fmt.Errorf(
				"[?RD]Parsing error at [[?RT]cmd:[?YW]%v[?RT], bcIdx:[?YW]%v[?RT][?RD]]\n"+
					"Raw: [[?YW]%v[?RT][?RD]][?RT]:[?RT] \n[?RD]->[?RT]    %v",
				fmt.Sprint(raw), fmt.Sprint(last_cmd_switch), *idx,
				strings.ReplaceAll(err.Error(), "\n", "\n[?RD]->[?RT]    "),
			)
		}
	}()
	log := func(text string) {
		if len(opts) > 0 {
			logger := opts[0].UEP.Logger
			if logger != nil {
				logger.PrintLog(public.LogParsing, text)
			}
		}
	}
	log(fmt.Sprintf("=========== START ==========="))
	log(fmt.Sprintf("start parsing code: '%v'", code))
	log(fmt.Sprintf("config: %v", p.Config))
	u := bytecode.Utils{}
	_Idx := p.Config.Shifter.Idx
	p.Config.Shifter.Idx = idx
	p.Config.Shifter.Code = code
	shift := p.Config.Shifter.ShiftError
	for *idx < len(code) {
		idx_start := *idx
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
		last_cmd_switch = command
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
			log(fmt.Sprintf("arglen, args %v; %v", arglen, args))
			args = append(args, arg)
		}
		raw = code[idx_start:*idx]
		result = append(result, ParsedBytes{
			Switch:   command,
			Args:     args,
			Raw:      raw,
			Metadata: make(map[string]interface{}),
		})
	}
	p.Config.Shifter.Idx = _Idx
	log(fmt.Sprintf("end parsing code:\n %v", result))
	log(fmt.Sprintf("=========== END ==========="))
	return result, nil
}

func (p *Parser1) String() string {
	return "lc/parsing/byteParsing/Parser1"
}
