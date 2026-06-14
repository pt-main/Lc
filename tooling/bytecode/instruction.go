package bytecode

type GenerationConfig struct {
	CommandBytelen   int
	ArglenBytelen    int
	ArgscountBytelen int
	Endianess        int
}

type InstructionsGenerator struct {
	Config GenerationConfig
}

func (ig *InstructionsGenerator) Generate(
	opcode int, args [][]byte,
) []byte {
	u := Utils{}
	res := append(
		append([]byte{}, u.IntToBytes(opcode, ig.Config.CommandBytelen, ig.Config.Endianess)...),
		u.IntToBytes(len(args), ig.Config.ArgscountBytelen, ig.Config.Endianess)...,
	)
	for _, arg := range args {
		res = append(res, u.IntToBytes(len(arg), ig.Config.ArglenBytelen, ig.Config.Endianess)...)
	}
	return res
}
