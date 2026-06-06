package bytecode

type InstructionsGenerator struct {
	OpcodeLen    int
	ArglenLen    int
	ArgscountLen int
	Endianess    int
}

func (ig *InstructionsGenerator) Generate(
	opcode int, args [][]byte,
) []byte {
	u := Utils{}
	res := append(
		append([]byte{}, u.IntToBytes(opcode, ig.OpcodeLen, ig.Endianess)...),
		u.IntToBytes(len(args), ig.ArgscountLen, ig.Endianess)...,
	)
	for _, arg := range args {
		res = append(res, u.IntToBytes(len(arg), ig.ArglenLen, ig.Endianess)...)
	}
	return res
}
