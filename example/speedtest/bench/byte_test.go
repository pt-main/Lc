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
	const ITERATIONS = 1_000_000_000
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
ITERATIONS=1_000_000_000:
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/speedtest/bench -count=10
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/speedtest/bench
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkByteProcessing-8              1        5762814944 ns/op               173.5 Mops/s        16344 B/op        176 allocs/op
BenchmarkByteProcessing-8              1        5751776315 ns/op               173.9 Mops/s        16184 B/op        175 allocs/op
BenchmarkByteProcessing-8              1        5897353936 ns/op               169.6 Mops/s        16952 B/op        184 allocs/op
BenchmarkByteProcessing-8              1        6498543824 ns/op               153.9 Mops/s        16376 B/op        176 allocs/op
BenchmarkByteProcessing-8              1        6598746076 ns/op               151.5 Mops/s        14776 B/op        161 allocs/op
BenchmarkByteProcessing-8              1        7007383883 ns/op               142.7 Mops/s        16264 B/op        175 allocs/op
BenchmarkByteProcessing-8              1        6527556070 ns/op               153.2 Mops/s        16280 B/op        175 allocs/op
BenchmarkByteProcessing-8              1        6253077383 ns/op               159.9 Mops/s        16424 B/op        177 allocs/op
BenchmarkByteProcessing-8              1        6086296262 ns/op               164.3 Mops/s        16248 B/op        175 allocs/op
BenchmarkByteProcessing-8              1        6452127563 ns/op               155.0 Mops/s        16248 B/op        175 allocs/op
PASS
ok      github.com/pt-main/lc/example/speedtest/bench   63.261s
macbook@MacBook-Pro lc %
ITERATIONS=500_000_000:
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/speedtest/bench -count=10
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/speedtest/bench
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkByteProcessing-8              1        2725905999 ns/op               183.4 Mops/s        20904 B/op        174 allocs/op
BenchmarkByteProcessing-8              1        2735958753 ns/op               182.8 Mops/s        15592 B/op        170 allocs/op
BenchmarkByteProcessing-8              1        2721029300 ns/op               183.8 Mops/s        15480 B/op        168 allocs/op
BenchmarkByteProcessing-8              1        2904751570 ns/op               172.1 Mops/s        16408 B/op        177 allocs/op
BenchmarkByteProcessing-8              1        2868240595 ns/op               174.3 Mops/s        16904 B/op        182 allocs/op
BenchmarkByteProcessing-8              1        2987468020 ns/op               167.4 Mops/s        15576 B/op        168 allocs/op
BenchmarkByteProcessing-8              1        2805655377 ns/op               178.2 Mops/s        16344 B/op        176 allocs/op
BenchmarkByteProcessing-8              1        2860468821 ns/op               174.8 Mops/s        15560 B/op        168 allocs/op
BenchmarkByteProcessing-8              1        2881849876 ns/op               173.5 Mops/s        16168 B/op        175 allocs/op
BenchmarkByteProcessing-8              1        3005677711 ns/op               166.4 Mops/s        16952 B/op        182 allocs/op
PASS
ok      github.com/pt-main/lc/example/speedtest/bench   28.963s
macbook@MacBook-Pro lc %
ITERATIONS=300_000_000:
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/speedtest/bench -count=10
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/speedtest/bench
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkByteProcessing-8              1        1788257526 ns/op               167.8 Mops/s        21896 B/op        183 allocs/op
BenchmarkByteProcessing-8              1        1735185620 ns/op               172.9 Mops/s        17048 B/op        182 allocs/op
BenchmarkByteProcessing-8              1        1743715603 ns/op               172.0 Mops/s        16440 B/op        176 allocs/op
BenchmarkByteProcessing-8              1        1698821182 ns/op               176.6 Mops/s        15512 B/op        168 allocs/op
BenchmarkByteProcessing-8              1        1723587268 ns/op               174.1 Mops/s        16376 B/op        177 allocs/op
BenchmarkByteProcessing-8              1        1753887836 ns/op               171.0 Mops/s        14872 B/op        161 allocs/op
BenchmarkByteProcessing-8              1        1743975178 ns/op               172.0 Mops/s        17064 B/op        182 allocs/op
BenchmarkByteProcessing-8              1        1924000551 ns/op               155.9 Mops/s        15576 B/op        168 allocs/op
BenchmarkByteProcessing-8              1        1807119062 ns/op               166.0 Mops/s        16344 B/op        175 allocs/op
BenchmarkByteProcessing-8              1        1769031176 ns/op               169.6 Mops/s        17000 B/op        182 allocs/op
PASS
ok      github.com/pt-main/lc/example/speedtest/bench   18.120s
macbook@MacBook-Pro lc %
ITERATIONS=100_000_000:
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/speedtest/bench -count=10
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/speedtest/bench
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkByteProcessing-8            933           1278002 ns/op                83.87 Mops/s       13039 B/op        158 allocs/op
BenchmarkByteProcessing-8            930           1232047 ns/op                87.28 Mops/s       13038 B/op        158 allocs/op
BenchmarkByteProcessing-8            838           1381281 ns/op                86.39 Mops/s       13099 B/op        158 allocs/op
BenchmarkByteProcessing-8            924           1189648 ns/op                90.97 Mops/s       13053 B/op        158 allocs/op
BenchmarkByteProcessing-8            914           1211689 ns/op                90.29 Mops/s       13026 B/op        157 allocs/op
BenchmarkByteProcessing-8            909           1201342 ns/op                91.57 Mops/s       13058 B/op        158 allocs/op
BenchmarkByteProcessing-8            904           1205359 ns/op                91.77 Mops/s       13062 B/op        158 allocs/op
BenchmarkByteProcessing-8            927           1188666 ns/op                90.75 Mops/s       13054 B/op        158 allocs/op
BenchmarkByteProcessing-8            915           1203756 ns/op                90.79 Mops/s       13034 B/op        158 allocs/op
BenchmarkByteProcessing-8            770           1318173 ns/op                98.52 Mops/s       13203 B/op        158 allocs/op
PASS
ok      github.com/pt-main/lc/example/speedtest/bench   92.990s
macbook@MacBook-Pro lc %
*/
