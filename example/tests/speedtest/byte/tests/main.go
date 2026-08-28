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

	parsed := make([]events.ByteCallAttr, 0, ITERATIONS)
	rawNode := byteParsing.ParsedBytes{
		Switch: []byte{0},
		Args:   [][]byte{},
	}
	handler := func(be engine.ByteEngineInterface, pb *byteParsing.ParsedBytes) core.ErrorInterface {
		return nil
	}
	bca := events.ByteCallAttr{
		RawNode: &rawNode,
		Handler: handler,
		Abis:    true,
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
	err := de.ByteCallHotLoopEvent(ev, i)
	timeE := time.Now()
	// pprof.StopCPUProfile()
	time := timeE.Sub(timeS)
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
	err := de.ByteCallHotLoopEvent(ev, i)
	// pprof.StopCPUProfile()
	time := timeE.Sub(timeS)
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
		iters = append(iters, testf(5_000_000))
	}
	for range Start / 2 {
		iters = append(iters, testf(10_000_000))
	}
	for range Start / 4 {
		iters = append(iters, testf(20_000_000))
	}
	for range Start / 8 {
		iters = append(iters, testf(40_000_000))
	}
	for range Start / 16 {
		iters = append(iters, testf(80_000_000))
	}
	for range Start / 32 {
		iters = append(iters, testf(160_000_000))
	}
	for range Start / 64 {
		iters = append(iters, testf(320_000_000))
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
	test(test1, 32)
	fmt.Println("===== Test2 =====")
	test(test2, 32)
}

/*
macbook@MacBook-Pro lc % go run ./example/tests/speedtest/byte/tests -gcflags="-m -m" -ldflags="-s -w"
Lc version - 1.5.7
===== Test1 =====
Test1: Iters: 5000000, Time: 19.331003ms, 263157894 mOps/s
Test1: Iters: 5000000, Time: 18.733908ms, 263157894 mOps/s
Test1: Iters: 5000000, Time: 22.980312ms, 217391304 mOps/s
Test1: Iters: 5000000, Time: 19.558461ms, 250000000 mOps/s
Test1: Iters: 5000000, Time: 19.739296ms, 250000000 mOps/s
Test1: Iters: 5000000, Time: 21.091147ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 21.742407ms, 227272727 mOps/s
Test1: Iters: 5000000, Time: 19.765059ms, 250000000 mOps/s
Test1: Iters: 5000000, Time: 21.57979ms, 227272727 mOps/s
Test1: Iters: 5000000, Time: 21.565381ms, 227272727 mOps/s
Test1: Iters: 5000000, Time: 22.473135ms, 227272727 mOps/s
Test1: Iters: 5000000, Time: 22.237547ms, 227272727 mOps/s
Test1: Iters: 5000000, Time: 23.38193ms, 217391304 mOps/s
Test1: Iters: 5000000, Time: 22.048882ms, 227272727 mOps/s
Test1: Iters: 5000000, Time: 20.128706ms, 250000000 mOps/s
Test1: Iters: 5000000, Time: 21.16066ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 19.281225ms, 263157894 mOps/s
Test1: Iters: 5000000, Time: 20.663264ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 20.035794ms, 250000000 mOps/s
Test1: Iters: 5000000, Time: 23.065067ms, 217391304 mOps/s
Test1: Iters: 5000000, Time: 22.874167ms, 217391304 mOps/s
Test1: Iters: 5000000, Time: 20.898845ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 20.811976ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 21.192832ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 21.092074ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 20.848486ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 20.688949ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 19.763027ms, 250000000 mOps/s
Test1: Iters: 5000000, Time: 21.939585ms, 227272727 mOps/s
Test1: Iters: 5000000, Time: 20.649626ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 20.758666ms, 238095238 mOps/s
Test1: Iters: 5000000, Time: 20.853927ms, 238095238 mOps/s
Test1: Iters: 10000000, Time: 39.078747ms, 256410256 mOps/s
Test1: Iters: 10000000, Time: 40.295967ms, 250000000 mOps/s
Test1: Iters: 10000000, Time: 38.578624ms, 256410256 mOps/s
Test1: Iters: 10000000, Time: 37.681705ms, 263157894 mOps/s
Test1: Iters: 10000000, Time: 38.255899ms, 263157894 mOps/s
Test1: Iters: 10000000, Time: 40.260515ms, 250000000 mOps/s
Test1: Iters: 10000000, Time: 40.59357ms, 243902439 mOps/s
Test1: Iters: 10000000, Time: 41.570728ms, 238095238 mOps/s
Test1: Iters: 10000000, Time: 38.745396ms, 256410256 mOps/s
Test1: Iters: 10000000, Time: 39.545718ms, 250000000 mOps/s
Test1: Iters: 10000000, Time: 38.340595ms, 263157894 mOps/s
Test1: Iters: 10000000, Time: 37.275479ms, 270270270 mOps/s
Test1: Iters: 10000000, Time: 39.588624ms, 250000000 mOps/s
Test1: Iters: 10000000, Time: 40.150805ms, 250000000 mOps/s
Test1: Iters: 10000000, Time: 37.195498ms, 270270270 mOps/s
Test1: Iters: 10000000, Time: 40.036955ms, 250000000 mOps/s
Test1: Iters: 20000000, Time: 74.349315ms, 270270270 mOps/s
Test1: Iters: 20000000, Time: 74.60037ms, 266666666 mOps/s
Test1: Iters: 20000000, Time: 73.616943ms, 270270270 mOps/s
Test1: Iters: 20000000, Time: 71.842844ms, 277777777 mOps/s
Test1: Iters: 20000000, Time: 76.208526ms, 263157894 mOps/s
Test1: Iters: 20000000, Time: 70.059533ms, 285714285 mOps/s
Test1: Iters: 20000000, Time: 75.241165ms, 266666666 mOps/s
Test1: Iters: 20000000, Time: 71.8311ms, 277777777 mOps/s
Test1: Iters: 40000000, Time: 148.648065ms, 268456375 mOps/s
Test1: Iters: 40000000, Time: 145.840353ms, 273972602 mOps/s
Test1: Iters: 40000000, Time: 146.652388ms, 272108843 mOps/s
Test1: Iters: 40000000, Time: 146.301707ms, 273972602 mOps/s
Test1: Iters: 80000000, Time: 285.775098ms, 279720279 mOps/s
Test1: Iters: 80000000, Time: 308.652823ms, 258899676 mOps/s
Test1: Iters: 160000000, Time: 867.274031ms, 184544405 mOps/s
Medium: 248862061, Median: 250000000,
Iters [184544405 217391304 217391304 217391304 217391304 227272727 227272727 227272727 227272727 227272727 227272727 227272727 238095238 238095238 238095238 238095238 238095238 238095238 238095238 238095238 238095238 238095238 238095238 238095238 238095238 243902439 250000000 250000000 250000000 250000000 250000000 250000000 250000000 250000000 250000000 250000000 250000000 250000000 256410256 256410256 256410256 258899676 263157894 263157894 263157894 263157894 263157894 263157894 263157894 266666666 266666666 268456375 270270270 270270270 270270270 270270270 272108843 273972602 273972602 277777777 277777777 279720279 285714285]
===== Test2 =====
Test2: Iters: 5000000, Time: 23.743946ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 23.778253ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 24.995652ms, 200000000 mOps/s
Test2: Iters: 5000000, Time: 23.941669ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 24.6286ms, 200000000 mOps/s
Test2: Iters: 5000000, Time: 24.779641ms, 200000000 mOps/s
Test2: Iters: 5000000, Time: 25.993257ms, 192307692 mOps/s
Test2: Iters: 5000000, Time: 25.318152ms, 200000000 mOps/s
Test2: Iters: 5000000, Time: 23.632493ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 23.744656ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 23.530285ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 22.8118ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 23.455578ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 23.230671ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 24.748188ms, 200000000 mOps/s
Test2: Iters: 5000000, Time: 23.377032ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 24.238652ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 24.465862ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 24.719291ms, 200000000 mOps/s
Test2: Iters: 5000000, Time: 23.440745ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 23.847918ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 24.911544ms, 200000000 mOps/s
Test2: Iters: 5000000, Time: 23.74668ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 22.658085ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 22.688899ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 23.432506ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 22.990193ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 22.863689ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 25.492215ms, 200000000 mOps/s
Test2: Iters: 5000000, Time: 23.930831ms, 208333333 mOps/s
Test2: Iters: 5000000, Time: 22.686733ms, 217391304 mOps/s
Test2: Iters: 5000000, Time: 23.156547ms, 217391304 mOps/s
Test2: Iters: 10000000, Time: 45.848181ms, 217391304 mOps/s
Test2: Iters: 10000000, Time: 45.559144ms, 217391304 mOps/s
Test2: Iters: 10000000, Time: 45.96151ms, 217391304 mOps/s
Test2: Iters: 10000000, Time: 45.77535ms, 217391304 mOps/s
Test2: Iters: 10000000, Time: 45.637576ms, 217391304 mOps/s
Test2: Iters: 10000000, Time: 45.141271ms, 222222222 mOps/s
Test2: Iters: 10000000, Time: 45.407469ms, 222222222 mOps/s
Test2: Iters: 10000000, Time: 44.985172ms, 222222222 mOps/s
Test2: Iters: 10000000, Time: 44.891841ms, 222222222 mOps/s
Test2: Iters: 10000000, Time: 47.209454ms, 212765957 mOps/s
Test2: Iters: 10000000, Time: 45.353728ms, 222222222 mOps/s
Test2: Iters: 10000000, Time: 45.751145ms, 217391304 mOps/s
Test2: Iters: 10000000, Time: 45.442977ms, 222222222 mOps/s
Test2: Iters: 10000000, Time: 47.030275ms, 212765957 mOps/s
Test2: Iters: 10000000, Time: 45.153691ms, 222222222 mOps/s
Test2: Iters: 10000000, Time: 47.250445ms, 212765957 mOps/s
Test2: Iters: 20000000, Time: 90.743824ms, 219780219 mOps/s
Test2: Iters: 20000000, Time: 98.412486ms, 204081632 mOps/s
Test2: Iters: 20000000, Time: 93.030863ms, 215053763 mOps/s
Test2: Iters: 20000000, Time: 90.437058ms, 222222222 mOps/s
Test2: Iters: 20000000, Time: 97.165843ms, 206185567 mOps/s
Test2: Iters: 20000000, Time: 91.614653ms, 217391304 mOps/s
Test2: Iters: 20000000, Time: 93.531025ms, 212765957 mOps/s
Test2: Iters: 20000000, Time: 88.639409ms, 224719101 mOps/s
Test2: Iters: 40000000, Time: 179.760469ms, 222222222 mOps/s
Test2: Iters: 40000000, Time: 185.233176ms, 216216216 mOps/s
Test2: Iters: 40000000, Time: 183.163349ms, 218579234 mOps/s
Test2: Iters: 40000000, Time: 220.352058ms, 181818181 mOps/s
Test2: Iters: 80000000, Time: 373.641796ms, 213903743 mOps/s
Test2: Iters: 80000000, Time: 376.967661ms, 212201591 mOps/s
Test2: Iters: 160000000, Time: 723.861096ms, 220994475 mOps/s
Medium: 212682645, Median: 217391304,
Iters [181818181 192307692 200000000 200000000 200000000 200000000 200000000 200000000 200000000 200000000 204081632 206185567 208333333 208333333 208333333 208333333 208333333 208333333 208333333 208333333 208333333 208333333 208333333 212201591 212765957 212765957 212765957 212765957 213903743 215053763 216216216 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 217391304 218579234 219780219 220994475 222222222 222222222 222222222 222222222 222222222 222222222 222222222 222222222 222222222 224719101]
*/
