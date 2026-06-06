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
}

type Parser1 struct {
	Config Parser1Config
}

func (p *Parser1) Parse(code []byte) ([]ParsedBytes, error) {
	u := bytecode.Utils{}
	result := []ParsedBytes{}
	idx := 0
	shiftStruct := u.ShiftStruct(code, &idx)
	shift := shiftStruct.ShiftError
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
