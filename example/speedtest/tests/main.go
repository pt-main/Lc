package main

import (
	"context"
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/pt-main/lc"
	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/engine/events"
	"github.com/pt-main/lc/parsing/byteParsing"
)

func test(ITERATIONS int) int {
	fmt.Printf("Test1: ")
	_idx := 0
	idx := &_idx
	e := lc.NewByteEngine(
		0, nil, true, &byteParsing.Parser1{},
		0, true, context.Background(),
	) // pseudo engine

	allIters := 0

	parsed := make([]events.ByteCallAttr, 0, ITERATIONS)
	rawNode := byteParsing.ParsedBytes{
		Switch: []byte{0},
		Args:   [][]byte{},
	}
	var timeE time.Time
	handler := func(be *engine.ByteEngine, pb *byteParsing.ParsedBytes) error {
		allIters += 1
		if allIters == (ITERATIONS - 1) {
			timeE = time.Now()
		}
		return nil
	}
	bca := events.ByteCallAttr{
		RawNode: &rawNode,
		Abis:    true,
		Handler: handler,
	}
	for i := 0; i < int(ITERATIONS); i++ {
		parsed = append(parsed, bca)
	}

	de := events.DefaultEvents{}
	i := &core.EventInput{
		Input: events.ByteCLDType{
			Ctx:    e.UEP.Context,
			Parsed: parsed,
			Engine: e,
			Idx:    idx,
		},
	}
	ev := e.UEP.Event.(*core.Events)

	// f, err := os.Create("cpu.prof")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// if err := pprof.StartCPUProfile(f); err != nil {
	// 	panic(err)
	// }

	var timeS time.Time = time.Now()
	de.ByteCallHotLoopEvent(ev, i)
	// pprof.StopCPUProfile()
	time := timeE.Sub(timeS)

	err := de.ByteCallHotLoopEvent(ev, i)
	if err != nil {
		fmt.Println(err)
	}
	mops := int(float64(ITERATIONS) / (math.Round(time.Seconds()*1000) / 1000))
	fmt.Printf("Iters: %v, Time: %v, %v mOps/s\n", ITERATIONS, time, mops)
	return mops
}

func test2(ITERATIONS int) int {
	fmt.Printf("Test2: ")
	_idx := 0
	idx := &_idx
	e := lc.NewByteEngine(
		0, nil, true, &byteParsing.Parser1{},
		0, true, context.Background(),
	) // pseudo engine

	allIters := 0

	rawNode := byteParsing.ParsedBytes{
		Switch: []byte{0},
		Args:   [][]byte{},
	}
	var timeE time.Time
	handler := func(be *engine.ByteEngine, pb *byteParsing.ParsedBytes) error {
		allIters += 1
		if allIters >= (ITERATIONS - 1) {
			timeE = time.Now()
			*idx = -1
		}
		return nil
	}
	bca := events.ByteCallAttr{
		RawNode: &rawNode,
		Abis:    false,
		Handler: handler,
	}

	de := events.DefaultEvents{}
	i := &core.EventInput{
		Input: events.ByteCLDType{
			Ctx:    e.UEP.Context,
			Parsed: []events.ByteCallAttr{bca},
			Engine: e,
			Idx:    idx,
		},
	}
	ev := e.UEP.Event.(*core.Events)

	// f, err := os.Create("cpu.prof")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// if err := pprof.StartCPUProfile(f); err != nil {
	// 	panic(err)
	// }

	var timeS time.Time = time.Now()
	de.ByteCallHotLoopEvent(ev, i)
	// pprof.StopCPUProfile()
	time := timeE.Sub(timeS)

	err := de.ByteCallHotLoopEvent(ev, i)
	if err != nil {
		fmt.Println(err)
	}
	mops := int(float64(ITERATIONS) / (math.Round(time.Seconds()*1000) / 1000))
	fmt.Printf("Iters: %v, Time: %v, %v mOps/s\n", ITERATIONS, time, mops)
	return mops
}

func main() {
	iters := []int{}
	iters = append(iters, test(300_000_000))
	Start := 22
	testf := test2
	for range Start {
		iters = append(iters, testf(10_000_000))
	}
	for range Start / 2 {
		iters = append(iters, testf(100_000_000))
	}
	for range Start / 4 {
		iters = append(iters, testf(200_000_000))
	}
	for range Start / 8 {
		iters = append(iters, testf(400_000_000))
	}
	for range Start / 16 {
		iters = append(iters, testf(800_000_000))
	}
	res := 0
	for _, val := range iters {
		res += val
	}
	res /= len(iters)
	m := int(math.Round(float64(len(iters) / 2)))
	sort.Ints(iters)
	fmt.Printf("Medium: %v, Median: %v, \nIters %v\n", res, iters[m], iters)
}

/*
macbook@MacBook-Pro lc % go run ./example/speedtest/tests/main.go -gcflags=-l
Iters: 1000000, Time: 6.44025ms, 166666666 mOps/s
Iters: 1000000, Time: 7.474175ms, 142857142 mOps/s
Iters: 1000000, Time: 6.23607ms, 166666666 mOps/s
Iters: 1000000, Time: 6.414676ms, 166666666 mOps/s
Iters: 1000000, Time: 6.283959ms, 166666666 mOps/s
Iters: 1000000, Time: 6.411133ms, 166666666 mOps/s
Iters: 1000000, Time: 6.255218ms, 166666666 mOps/s
Iters: 1000000, Time: 6.066866ms, 166666666 mOps/s
Iters: 1000000, Time: 6.730023ms, 142857142 mOps/s
Iters: 1000000, Time: 5.835124ms, 166666666 mOps/s
Iters: 1000000, Time: 6.859053ms, 142857142 mOps/s
Iters: 1000000, Time: 6.408876ms, 166666666 mOps/s
Iters: 1000000, Time: 6.381622ms, 166666666 mOps/s
Iters: 1000000, Time: 7.036407ms, 142857142 mOps/s
Iters: 1000000, Time: 6.644972ms, 142857142 mOps/s
Iters: 1000000, Time: 6.300575ms, 166666666 mOps/s
Iters: 10000000, Time: 60.794027ms, 163934426 mOps/s
Iters: 10000000, Time: 68.917019ms, 144927536 mOps/s
Iters: 10000000, Time: 60.186579ms, 166666666 mOps/s
Iters: 10000000, Time: 65.465579ms, 153846153 mOps/s
Iters: 10000000, Time: 61.775413ms, 161290322 mOps/s
Iters: 10000000, Time: 65.19226ms, 153846153 mOps/s
Iters: 10000000, Time: 72.696387ms, 136986301 mOps/s
Iters: 10000000, Time: 65.504644ms, 151515151 mOps/s
Iters: 100000000, Time: 598.915141ms, 166944908 mOps/s
Iters: 100000000, Time: 1.089113669s, 91827364 mOps/s
Medium: 155361692, Median: 166666666,
Iters [91827364 136986301 142857142 142857142 142857142 142857142 142857142 144927536 151515151 153846153 153846153 161290322 163934426 166666666 166666666 166666666 166666666 166666666 166666666 166666666 166666666 166666666 166666666 166666666 166666666 166944908]
macbook@MacBook-Pro lc %
*/
