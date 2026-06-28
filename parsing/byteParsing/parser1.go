package byteParsing

import (
	"errors"
	"fmt"

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
func (p *Parser1) Parse(code []byte, opts ...*parsing.ParseOption) ([]ParsedBytes, error) {
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
	result := []ParsedBytes{}
	_Idx := p.Config.Shifter.Idx
	idx := 0
	p.Config.Shifter.Idx = &idx
	p.Config.Shifter.Code = code
	shift := p.Config.Shifter.ShiftError
	for idx < len(code) {
		idx_start := idx
		command, err := shift(p.Config.GConfig.CommandBytelen)
		if err != nil {
			return nil, err
		}
		argscountBytes, err := shift(p.Config.GConfig.ArgscountBytelen)
		if err != nil {
			return nil, err
		}
		argscount := u.BytesToInt(argscountBytes, p.Config.GConfig.Endianess)
		args := [][]byte{}
		log(fmt.Sprintf("cmd %v, argscount %v", command, argscount))
		for range argscount {
			arglenBytes, err := shift(p.Config.GConfig.ArglenBytelen)
			if err != nil {
				return nil, err
			}
			arglen := u.BytesToInt(arglenBytes, p.Config.GConfig.Endianess)
			if arglen == 0 {
				return nil, errors.New("Can't form args with 0 argument length")
			}
			log(fmt.Sprintf("arglen %v", arglen))
			arg, err := shift(arglen)
			if err != nil {
				return nil, err
			}
			log(fmt.Sprintf("arg %v", arglen))
			args = append(args, arg)
		}
		raw := code[idx_start:idx]
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
