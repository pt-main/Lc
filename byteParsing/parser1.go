package byteParsing

import (
	"errors"

	"github.com/pt-main/lc/tooling/bytecode"
)

type Parser1Config struct {
	CommandBytelen   int
	ArglenBytelen    int
	ArgscountBytelen int
	Endianess        int
	Shifter          bytecode.Shift
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
func (p *Parser1) Parse(code []byte) ([]ParsedBytes, error) {
	u := bytecode.Utils{}
	result := []ParsedBytes{}
	idx := 0
	shift := p.Config.Shifter.ShiftError
	for idx < len(code) {
		idx_start := idx
		command, err := shift(p.Config.CommandBytelen)
		if err != nil {
			return nil, err
		}
		argscountBytes, err := shift(p.Config.ArgscountBytelen)
		if err != nil {
			return nil, err
		}
		argscount := u.BytesToInt(argscountBytes, p.Config.Endianess)
		args := [][]byte{}
		for range argscount {
			arglenBytes, err := shift(p.Config.ArglenBytelen)
			if err != nil {
				return nil, err
			}
			arglen := u.BytesToInt(arglenBytes, p.Config.Endianess)
			if arglen == 0 {
				return nil, errors.New("Can't form args with 0 argument length")
			}
			arg, err := shift(arglen)
			if err != nil {
				return nil, err
			}
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
	return result, nil
}
