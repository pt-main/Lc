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
macbook@MacBook-Pro lc % go run ./example/speedtest/simple
Test1: Iters: 300000000, Time: 2.446358945s, 122649223 mOps/s
Test2: Iters: 10000000, Time: 76.311314ms, 131578947 mOps/s
Test2: Iters: 10000000, Time: 66.759756ms, 149253731 mOps/s
Test2: Iters: 10000000, Time: 72.029353ms, 138888888 mOps/s
Test2: Iters: 10000000, Time: 71.847038ms, 138888888 mOps/s
Test2: Iters: 10000000, Time: 76.29681ms, 131578947 mOps/s
Test2: Iters: 10000000, Time: 69.096247ms, 144927536 mOps/s
Test2: Iters: 10000000, Time: 71.735267ms, 138888888 mOps/s
Test2: Iters: 10000000, Time: 72.585652ms, 136986301 mOps/s
Test2: Iters: 10000000, Time: 70.953707ms, 140845070 mOps/s
Test2: Iters: 10000000, Time: 67.725707ms, 147058823 mOps/s
Test2: Iters: 10000000, Time: 67.901685ms, 147058823 mOps/s
Test2: Iters: 10000000, Time: 68.041727ms, 147058823 mOps/s
Test2: Iters: 10000000, Time: 71.975185ms, 138888888 mOps/s
Test2: Iters: 10000000, Time: 69.4756ms, 144927536 mOps/s
Test2: Iters: 10000000, Time: 69.540027ms, 142857142 mOps/s
Test2: Iters: 10000000, Time: 67.524449ms, 147058823 mOps/s
Test2: Iters: 10000000, Time: 70.798522ms, 140845070 mOps/s
Test2: Iters: 10000000, Time: 67.910101ms, 147058823 mOps/s
Test2: Iters: 10000000, Time: 69.006502ms, 144927536 mOps/s
Test2: Iters: 10000000, Time: 77.495368ms, 129870129 mOps/s
Test2: Iters: 10000000, Time: 67.567186ms, 147058823 mOps/s
Test2: Iters: 10000000, Time: 68.06376ms, 147058823 mOps/s
Test2: Iters: 100000000, Time: 663.759594ms, 150602409 mOps/s
Test2: Iters: 100000000, Time: 634.795679ms, 157480314 mOps/s
Test2: Iters: 100000000, Time: 669.849435ms, 149253731 mOps/s
Test2: Iters: 100000000, Time: 706.638528ms, 141442715 mOps/s
Test2: Iters: 100000000, Time: 706.734595ms, 141442715 mOps/s
Test2: Iters: 100000000, Time: 701.25373ms, 142653352 mOps/s
Test2: Iters: 100000000, Time: 693.388277ms, 144300144 mOps/s
Test2: Iters: 100000000, Time: 660.085522ms, 151515151 mOps/s
Test2: Iters: 100000000, Time: 661.084127ms, 151285930 mOps/s
Test2: Iters: 100000000, Time: 681.472074ms, 146842878 mOps/s
Test2: Iters: 100000000, Time: 665.36999ms, 150375939 mOps/s
Test2: Iters: 200000000, Time: 1.32977461s, 150375939 mOps/s
Test2: Iters: 200000000, Time: 1.338498563s, 149476831 mOps/s
Test2: Iters: 200000000, Time: 1.379985022s, 144927536 mOps/s
Test2: Iters: 200000000, Time: 1.362704474s, 146735143 mOps/s
Test2: Iters: 200000000, Time: 1.254650044s, 159362549 mOps/s
Test2: Iters: 400000000, Time: 2.572008703s, 155520995 mOps/s
Test2: Iters: 400000000, Time: 2.702153071s, 148038490 mOps/s
Test2: Iters: 800000000, Time: 4.966114275s, 161095449 mOps/s
Medium: 144974825, Median: 146842878,
Iters [122649223 129870129 131578947 131578947 136986301 138888888 138888888 138888888 138888888 140845070 140845070 141442715 141442715 142653352 142857142 144300144 144927536 144927536 144927536 144927536 146735143 146842878 147058823 147058823 147058823 147058823 147058823 147058823 147058823 148038490 149253731 149253731 149476831 150375939 150375939 150602409 151285930 151515151 155520995 157480314 159362549 161095449]
macbook@MacBook-Pro lc %
*/
