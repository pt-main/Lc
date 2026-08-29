package main

import (
	"fmt"

	"github.com/pt-main/lc"
	enginepkg "github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/tooling/bytecode"
)

func main() {
	// Parsing format:
	// instruction {
	//     [bytes : cmd] [bytes : argscount] [bytes : arglen]  [bytes arglen : arg],
	//                                       [bytes : arglen2] [bytes arglen2 : arg2]...
	// }
	parser := &byteParsing.Parser1{
		Config: byteParsing.Parser1Config{
			GConfig: bytecode.GenerationConfig{
				CommandBytelen:   1,
				ArgscountBytelen: 1,
				ArglenBytelen:    2,
				Endianess:        public.LittleEndian,
			},
			Shifter: bytecode.Shift{},
		},
	}

	engine, err := lc.NewEngineBuilder(public.ByteEngineType, public.StringResType).
		WithPipeline([]string{"main"}).
		WithByteParser(parser).
		WithDefaultEvents(true).
		Build()
	if err != nil {
		panic(err)
	}

	err = engine.NewCommandByte(1, func(be enginepkg.ByteEngineInterface, node *byteParsing.ParsedBytes) core.ErrorInterface {
		for _, arg := range node.Args {
			if err := be.GetUep().Generator.AddString(string(arg), "main"); err != nil {
				return err
			}
		}
		return nil
	}, "add to output instruction", true)
	if err != nil {
		panic(err)
	}

	code := []byte{
		0x01,       // opcode=1
		0x01,       // argsCount=1
		0x03, 0x00, // arglen=3 (little endian, 2 bytes)
		0x61, 0x62, 0x63, // args="abc" (3 bytes)
	}

	err = engine.ProcessBytes(code)
	if err != nil {
		panic(err)
	}

	uep, _ := engine.GetUEP()
	out, err := core.GetStringRes(uep.Generator, "")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", out)
}
