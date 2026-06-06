package bytecode

type InstructionsGenrator struct {
	OpcodeLen    int
	ArglenLen    int
	ArgscountLen int
	Endianess    int
}

func (ig *InstructionsGenrator) Generate(
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
