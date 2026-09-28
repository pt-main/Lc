// Work with engine.ByteEngine: build bytecode by hand, register opcodes and
// read the decoded arguments in the handler.

package main

import (
	"fmt"

	"github.com/pt-main/lc/v2"
	enginepkg "github.com/pt-main/lc/v2/engine"
	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/parsing/byteParsing"
	"github.com/pt-main/lc/v2/public"
	"github.com/pt-main/lc/v2/tooling/bytecode"
)

func main() {
	// instruction {
	//     [bytes : cmd] [bytes : argscount] [bytes : arglen]  [bytes arglen : arg]...
	// }
	parser := &byteParsing.Parser1{
		Config: byteParsing.Parser1Config{
			GConfig: bytecode.GenerationConfig{
				CommandBytelen:   1,
				ArgscountBytelen: 1,
				ArglenBytelen:    2,
				Endianness:       public.LittleEndian,
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
	}, "print every argument", true)
	if err != nil {
		panic(err)
	}

	err = engine.NewCommandByte(2, func(be enginepkg.ByteEngineInterface, node *byteParsing.ParsedBytes) core.ErrorInterface {
		return be.GetUep().Generator.AddString(fmt.Sprintf("(no args expected, got %d)", len(node.Args)), "main")
	}, "report an unexpected argument list", true)
	if err != nil {
		panic(err)
	}

	code := []byte{
		0x01,       // opcode 1
		0x02,       // argscount 2
		0x03, 0x00, // arglen 3
		0x61, 0x62, 0x63, // "abc"
		0x02, 0x00, // arglen 2
		0x64, 0x65, // "de"
		0x02, // opcode 2
		0x00, // argscount 0
	}

	if err := engine.ProcessBytes(code); err != nil {
		panic(err)
	}

	uep, _ := engine.GetUEP()
	out, err := core.GetStringRes(uep.Generator, "\n")
	if err != nil {
		panic(err)
	}
	fmt.Println(out)

	if err := engine.End(); err != nil {
		panic(err)
	}
	fmt.Println("engine lifecycle: ended")
}
