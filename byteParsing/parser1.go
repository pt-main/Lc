package byteParsing

import "errors"

type Parser1Config struct {
	CommandBytelen   int
	ArglenBytelen    int
	ArgscountBytelen int
}

type Parser1 struct {
	Config Parser1Config
}

func (p *Parser1) Parse(code []byte) ([]ParsedBytes, error) {
	u := Utils{}
	result := []ParsedBytes{}
	idx := 0
	shiftStruct := u.ShiftStruct(code)
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
		argscount := u.BytesToInt(argscountBytes)
		args := [][]byte{}
		for range argscount {
			arglenBytes, err := shift(p.Config.ArglenBytelen)
			if err != nil {
				return nil, err
			}
			arglen := u.BytesToInt(arglenBytes)
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
