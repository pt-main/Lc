package speedtest

import (
	"context"
	"testing"

	"github.com/pt-main/lc"
	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/parsing/byteParsing"
	"github.com/pt-main/lc/public"
	"github.com/pt-main/lc/tooling/bytecode"
)

func BenchmarkByteProcessing(b *testing.B) {
	const ITERATIONS = 100_000_000
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
	var iteration int
	e.NewCommand(0, func(be *engine.ByteEngine, pb *byteParsing.ParsedBytes) error {
		iteration += 1
		if iteration+1 == ITERATIONS {
			be.SetBytecodeIdx(-1)
		}
		return nil
	}, "nil", false)

	b.ResetTimer()
	for range b.N {
		iteration = 0
		e.Process([]byte{0, 0})
		b.ReportMetric(float64(ITERATIONS)/b.Elapsed().Seconds()/1e6, "Mops/s")
	}
}

/* ----==== RESULTS ====----
ITERATIONS=100_000_000:
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/speedtest -count=10
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/speedtest
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkByteProcessing-8            868           1197091 ns/op                96.24 Mops/s       13080 B/op        158 allocs/op
BenchmarkByteProcessing-8           1224            939444 ns/op                86.97 Mops/s       13298 B/op        158 allocs/op
BenchmarkByteProcessing-8           1033            985341 ns/op                98.25 Mops/s       13184 B/op        158 allocs/op
BenchmarkByteProcessing-8           1220            948301 ns/op                86.44 Mops/s       13305 B/op        158 allocs/op
BenchmarkByteProcessing-8           1078            960905 ns/op                96.54 Mops/s       13181 B/op        158 allocs/op
BenchmarkByteProcessing-8            982           1028498 ns/op                99.01 Mops/s       13238 B/op        158 allocs/op
BenchmarkByteProcessing-8           1123            963167 ns/op                92.45 Mops/s       13129 B/op        158 allocs/op
BenchmarkByteProcessing-8           1094            965980 ns/op                94.63 Mops/s       13134 B/op        158 allocs/op
BenchmarkByteProcessing-8            789           1313552 ns/op                96.49 Mops/s       13145 B/op        157 allocs/op
BenchmarkByteProcessing-8            688           1458637 ns/op                99.65 Mops/s       13004 B/op        157 allocs/op
PASS
ok      github.com/pt-main/lc/example/speedtest 115.830s
*/
