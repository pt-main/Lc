package bytecode

import "github.com/pt-main/lc/public"

// GenerationConfig describes the fixed field layout of one instruction.
type GenerationConfig struct {
	CommandBytelen   int
	ArglenBytelen    int
	ArgscountBytelen int
	Endianness       public.EndianType
}

type InstructionsGenerator struct {
	Config GenerationConfig
}

// Generate encodes one instruction in the layout byteParsing.Parser1 reads:
// opcode, argument count, then the length and data of every argument.
//
// It panics on an empty argument, which cannot be encoded in that layout.
func (ig *InstructionsGenerator) Generate(opcode int, args [][]byte) []byte {
	u := Utils{}
	res := append(
		append([]byte{}, u.IntToBytes(opcode, ig.Config.CommandBytelen, ig.Config.Endianness)...),
		u.IntToBytes(len(args), ig.Config.ArgscountBytelen, ig.Config.Endianness)...,
	)
	for _, arg := range args {
		if len(arg) == 0 {
			panic("Argument length cannot be zero")
		}
		res = append(res, u.IntToBytes(len(arg), ig.Config.ArglenBytelen, ig.Config.Endianness)...)
		res = append(res, arg...)
	}
	return res
}
