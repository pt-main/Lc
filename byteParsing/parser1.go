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
	shift := func(length int) ([]byte, error) {
		if idx+length > len(code) {
			return nil, errors.New("unexpected end of data")
		}
		res := code[idx : idx+length]
		idx += length
		return res, nil
	}
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
