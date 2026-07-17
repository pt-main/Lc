package main

import (
	"context"
	"fmt"
	"time"

	"github.com/pt-main/lc"
	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/tooling/bytecode"
)

func test() {
	ITERATIONS := 100_000_000
	_idx := 0
	end := public.BigEndian
	gc := byteParsing.Parser1Config{
		GConfig: bytecode.GenerationConfig{
			CommandBytelen: 1, ArgscountBytelen: 1,
			ArglenBytelen: 1, Endianess: end,
		}, Shifter: *bytecode.NewShift(make([]byte, 0), &_idx),
	}
	e := lc.NewByteEngine(
		0, nil, true, &byteParsing.Parser1{Config: gc},
		end, true, context.Background(),
	)
	iteration := 0
	var sT time.Time
	var eT time.Time
	e.NewCommand(0, func(be *engine.ByteEngine, pb byteParsing.ParsedBytes) error {
		if iteration == 0 {
			sT = time.Now()
		}
		iteration += 1
		if iteration+1 == ITERATIONS {
			eT = time.Now()
			be.SetBytecodeIdx(-1)
		}
		return nil
	}, "nil", false)
	fmt.Println("start")
	e.Process([]byte{0, 0})
	fmt.Println("end", eT.Sub(sT))
}

func main() {
	test()
}
