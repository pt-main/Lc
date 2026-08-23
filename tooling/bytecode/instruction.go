package bytecode

import "github.com/pt-main/lc/public"

type GenerationConfig struct {
	CommandBytelen   int
	ArglenBytelen    int
	ArgscountBytelen int
	Endianess        public.EndianType
}

type InstructionsGenerator struct {
	Config GenerationConfig
}

// Generate instruction.
// byteParsing.Parser1 specification
//
// !!! PANIC IF ANY ARGLEN IS ZERO !!!
func (ig *InstructionsGenerator) Generate(
	opcode int, args [][]byte,
) []byte {
	u := Utils{}
	res := append(
		append([]byte{}, u.IntToBytes(opcode, ig.Config.CommandBytelen, ig.Config.Endianess)...),
		u.IntToBytes(len(args), ig.Config.ArgscountBytelen, ig.Config.Endianess)...,
	)
	for _, arg := range args {
		if len(arg) == 0 {
			panic("Argument length cannot be zero")
		}
		res = append(res, u.IntToBytes(len(arg), ig.Config.ArglenBytelen, ig.Config.Endianess)...)
		res = append(res, arg...)
	}
	return res
}
