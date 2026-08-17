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

func test1(ITERATIONS int) int {
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
	handler := func(be engine.ByteEngineInterface, pb *byteParsing.ParsedBytes) core.ErrorInterface {
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
	handler := func(bei engine.ByteEngineInterface, pb *byteParsing.ParsedBytes) core.ErrorInterface {
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

func test(testf func(int) int, Start int) {
	iters := []int{}
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

func main() {
	fmt.Println("Lc version -", lc.Version)
	fmt.Println("===== Test1 =====")
	test(test1, 10)
	fmt.Println("===== Test2 =====")
	test(test2, 32)
}

/*
Lc version - 1.5.1
===== Test1 =====
Test1: Iters: 10000000, Time: 83.631101ms, 119047619 mOps/s
Test1: Iters: 10000000, Time: 80.806105ms, 123456790 mOps/s
Test1: Iters: 10000000, Time: 83.966236ms, 119047619 mOps/s
Test1: Iters: 10000000, Time: 84.77732ms, 117647058 mOps/s
Test1: Iters: 10000000, Time: 92.00846ms, 108695652 mOps/s
Test1: Iters: 10000000, Time: 81.311348ms, 123456790 mOps/s
Test1: Iters: 10000000, Time: 81.962906ms, 121951219 mOps/s
Test1: Iters: 10000000, Time: 80.21774ms, 125000000 mOps/s
Test1: Iters: 10000000, Time: 87.345851ms, 114942528 mOps/s
Test1: Iters: 10000000, Time: 82.67824ms, 120481927 mOps/s
Test1: Iters: 100000000, Time: 802.337396ms, 124688279 mOps/s
Test1: Iters: 100000000, Time: 847.627388ms, 117924528 mOps/s
Test1: Iters: 100000000, Time: 798.604082ms, 125156445 mOps/s
Test1: Iters: 100000000, Time: 788.823594ms, 126742712 mOps/s
Test1: Iters: 100000000, Time: 762.808279ms, 131061598 mOps/s
Test1: Iters: 200000000, Time: 1.682078851s, 118906064 mOps/s
Test1: Iters: 200000000, Time: 1.645749681s, 121506682 mOps/s
Test1: Iters: 400000000, Time: 9.133057429s, 43797218 mOps/s
Medium: 116861707, Median: 121506682,
Iters [43797218 108695652 114942528 117647058 117924528 118906064 119047619 119047619 120481927 121506682 121951219 123456790 123456790 124688279 125000000 125156445 126742712 131061598]
===== Test2 =====
Test2: Iters: 10000000, Time: 74.136485ms, 135135135 mOps/s
Test2: Iters: 10000000, Time: 80.690947ms, 123456790 mOps/s
Test2: Iters: 10000000, Time: 74.163965ms, 135135135 mOps/s
Test2: Iters: 10000000, Time: 109.85331ms, 90909090 mOps/s
Test2: Iters: 10000000, Time: 76.635534ms, 129870129 mOps/s
Test2: Iters: 10000000, Time: 66.212814ms, 151515151 mOps/s
Test2: Iters: 10000000, Time: 76.931153ms, 129870129 mOps/s
Test2: Iters: 10000000, Time: 66.131764ms, 151515151 mOps/s
Test2: Iters: 10000000, Time: 63.494247ms, 158730158 mOps/s
Test2: Iters: 10000000, Time: 66.28764ms, 151515151 mOps/s
Test2: Iters: 10000000, Time: 67.429024ms, 149253731 mOps/s
Test2: Iters: 10000000, Time: 67.209301ms, 149253731 mOps/s
Test2: Iters: 10000000, Time: 64.580462ms, 153846153 mOps/s
Test2: Iters: 10000000, Time: 72.29076ms, 138888888 mOps/s
Test2: Iters: 10000000, Time: 65.606012ms, 151515151 mOps/s
Test2: Iters: 10000000, Time: 64.321543ms, 156250000 mOps/s
Test2: Iters: 10000000, Time: 64.266137ms, 156250000 mOps/s
Test2: Iters: 10000000, Time: 69.527309ms, 142857142 mOps/s
Test2: Iters: 10000000, Time: 75.758207ms, 131578947 mOps/s
Test2: Iters: 10000000, Time: 65.858495ms, 151515151 mOps/s
Test2: Iters: 10000000, Time: 65.835576ms, 151515151 mOps/s
Test2: Iters: 10000000, Time: 62.492821ms, 161290322 mOps/s
Test2: Iters: 10000000, Time: 69.924138ms, 142857142 mOps/s
Test2: Iters: 10000000, Time: 64.052846ms, 156250000 mOps/s
Test2: Iters: 10000000, Time: 65.010355ms, 153846153 mOps/s
Test2: Iters: 10000000, Time: 66.219458ms, 151515151 mOps/s
Test2: Iters: 10000000, Time: 65.548766ms, 151515151 mOps/s
Test2: Iters: 10000000, Time: 70.480818ms, 142857142 mOps/s
Test2: Iters: 10000000, Time: 62.270297ms, 161290322 mOps/s
Test2: Iters: 10000000, Time: 62.410387ms, 161290322 mOps/s
Test2: Iters: 10000000, Time: 61.298797ms, 163934426 mOps/s
Test2: Iters: 10000000, Time: 69.284619ms, 144927536 mOps/s
Test2: Iters: 100000000, Time: 643.836767ms, 155279503 mOps/s
Test2: Iters: 100000000, Time: 685.867418ms, 145772594 mOps/s
Test2: Iters: 100000000, Time: 653.605758ms, 152905198 mOps/s
Test2: Iters: 100000000, Time: 634.22513ms, 157728706 mOps/s
Test2: Iters: 100000000, Time: 629.747093ms, 158730158 mOps/s
Test2: Iters: 100000000, Time: 636.79118ms, 156985871 mOps/s
Test2: Iters: 100000000, Time: 631.620296ms, 158227848 mOps/s
Test2: Iters: 100000000, Time: 637.004552ms, 156985871 mOps/s
Test2: Iters: 100000000, Time: 639.77815ms, 156250000 mOps/s
Test2: Iters: 100000000, Time: 627.149309ms, 159489633 mOps/s
Test2: Iters: 100000000, Time: 630.189662ms, 158730158 mOps/s
Test2: Iters: 100000000, Time: 640.183614ms, 156250000 mOps/s
Test2: Iters: 100000000, Time: 671.145516ms, 149031296 mOps/s
Test2: Iters: 100000000, Time: 673.820605ms, 148367952 mOps/s
Test2: Iters: 100000000, Time: 751.899969ms, 132978723 mOps/s
Test2: Iters: 100000000, Time: 928.532697ms, 107642626 mOps/s
Test2: Iters: 200000000, Time: 1.309047815s, 152788388 mOps/s
Test2: Iters: 200000000, Time: 1.417237526s, 141143260 mOps/s
Test2: Iters: 200000000, Time: 1.3435233s, 148809523 mOps/s
Test2: Iters: 200000000, Time: 1.329056673s, 150489089 mOps/s
Test2: Iters: 200000000, Time: 1.239711748s, 161290322 mOps/s
Test2: Iters: 200000000, Time: 1.234268487s, 162074554 mOps/s
Test2: Iters: 200000000, Time: 1.264810986s, 158102766 mOps/s
Test2: Iters: 200000000, Time: 1.236752963s, 161681487 mOps/s
Test2: Iters: 400000000, Time: 2.558859178s, 156311059 mOps/s
Test2: Iters: 400000000, Time: 2.614027526s, 153022188 mOps/s
Test2: Iters: 400000000, Time: 2.528229248s, 158227848 mOps/s
Test2: Iters: 400000000, Time: 2.580307925s, 155038759 mOps/s
Test2: Iters: 800000000, Time: 5.04125604s, 158698670 mOps/s
Test2: Iters: 800000000, Time: 5.089182014s, 157201807 mOps/s
Medium: 149486864, Median: 152905198,
Iters [90909090 107642626 123456790 129870129 129870129 131578947 132978723 135135135 135135135 138888888 141143260 142857142 142857142 142857142 144927536 145772594 148367952 148809523 149031296 149253731 149253731 150489089 151515151 151515151 151515151 151515151 151515151 151515151 151515151 151515151 152788388 152905198 153022188 153846153 153846153 155038759 155279503 156250000 156250000 156250000 156250000 156250000 156311059 156985871 156985871 157201807 157728706 158102766 158227848 158227848 158698670 158730158 158730158 158730158 159489633 161290322 161290322 161290322 161290322 161681487 162074554 163934426]

*/
