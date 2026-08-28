package speedtest

import (
	"context"
	"testing"

	"github.com/pt-main/lc"
	"github.com/pt-main/lc/engine"
	"github.com/pt-main/lc/engine/core"
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
	e.NewCommandFull(0, func(bei engine.ByteEngineInterface, pb *byteParsing.ParsedBytes) core.ErrorInterface {
		be := bei.(*engine.ByteEngine)
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

Lc version - 1.5.7

ITERATIONS=1_000_000_000
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/tests/speedtest/byte/bench -count=10
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/tests/speedtest/byte/bench
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkByteProcessing-8               1        4636165728 ns/op               215.7 Mops/s         6992 B/op        105 allocs/op
BenchmarkByteProcessing-8               1        4678134135 ns/op               213.8 Mops/s         6976 B/op        104 allocs/op
BenchmarkByteProcessing-8               1        4545117089 ns/op               220.0 Mops/s         6576 B/op        102 allocs/op
BenchmarkByteProcessing-8               1        4521750876 ns/op               221.2 Mops/s         6976 B/op        104 allocs/op
BenchmarkByteProcessing-8               1        5308026598 ns/op               188.4 Mops/s         6976 B/op        104 allocs/op
BenchmarkByteProcessing-8               1        4701634606 ns/op               212.7 Mops/s         6976 B/op        104 allocs/op
BenchmarkByteProcessing-8               1        4964960666 ns/op               201.4 Mops/s         6528 B/op        101 allocs/op
BenchmarkByteProcessing-8               1        4707789927 ns/op               212.4 Mops/s         6528 B/op        101 allocs/op
BenchmarkByteProcessing-8               1        4483261343 ns/op               223.1 Mops/s         6528 B/op        101 allocs/op
BenchmarkByteProcessing-8               1        4644296688 ns/op               215.3 Mops/s         6912 B/op        104 allocs/op
PASS
ok      github.com/pt-main/lc/example/tests/speedtest/byte/bench        47.634s

ITERATIONS=500_000_000
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/tests/speedtest/byte/bench -count=10
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/tests/speedtest/byte/bench
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkByteProcessing-8              1        2416450521 ns/op               206.9 Mops/s         8440 B/op        109 allocs/op
BenchmarkByteProcessing-8              1        2308473392 ns/op               216.6 Mops/s         7688 B/op        103 allocs/op
BenchmarkByteProcessing-8              1        2233277662 ns/op               223.9 Mops/s         7688 B/op        103 allocs/op
BenchmarkByteProcessing-8              1        2194493841 ns/op               227.8 Mops/s         7688 B/op        103 allocs/op
BenchmarkByteProcessing-8              1        2419690868 ns/op               206.6 Mops/s         7688 B/op        103 allocs/op
BenchmarkByteProcessing-8              1        2407318933 ns/op               207.7 Mops/s         8376 B/op        108 allocs/op
BenchmarkByteProcessing-8              1        2760941809 ns/op               181.1 Mops/s         8376 B/op        108 allocs/op
BenchmarkByteProcessing-8              1        2616158593 ns/op               191.1 Mops/s         7688 B/op        103 allocs/op
BenchmarkByteProcessing-8              1        2629497097 ns/op               190.2 Mops/s         7688 B/op        103 allocs/op
BenchmarkByteProcessing-8              1        2482259449 ns/op               201.4 Mops/s         8376 B/op        108 allocs/op
PASS
ok      github.com/pt-main/lc/example/tests/speedtest/byte/bench        24.940s

ITERATIONS=250_000_000
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/tests/speedtest/byte/bench -count=10
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/tests/speedtest/byte/bench
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkByteProcessing-8              1        1155587732 ns/op               216.3 Mops/s         8376 B/op        108 allocs/op
BenchmarkByteProcessing-8              1        1101299266 ns/op               227.0 Mops/s         8376 B/op        108 allocs/op
BenchmarkByteProcessing-8              1        1200872455 ns/op               208.2 Mops/s         8376 B/op        108 allocs/op
BenchmarkByteProcessing-8              1        1187507477 ns/op               210.5 Mops/s         8392 B/op        109 allocs/op
BenchmarkByteProcessing-8              1        1209859392 ns/op               206.6 Mops/s         7688 B/op        103 allocs/op
BenchmarkByteProcessing-8              1        1182012312 ns/op               211.5 Mops/s         8376 B/op        108 allocs/op
BenchmarkByteProcessing-8              1        1182744295 ns/op               211.4 Mops/s         7784 B/op        104 allocs/op
BenchmarkByteProcessing-8              1        1196835605 ns/op               208.9 Mops/s         7688 B/op        103 allocs/op
BenchmarkByteProcessing-8              1        1188541998 ns/op               210.3 Mops/s         7848 B/op        104 allocs/op
BenchmarkByteProcessing-8              1        1139287512 ns/op               219.4 Mops/s         8376 B/op        108 allocs/op
PASS
ok      github.com/pt-main/lc/example/tests/speedtest/byte/bench        12.220s

ITERATIONS=100_000_000
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/tests/speedtest/byte/bench -count=10
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/tests/speedtest/byte/bench
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkByteProcessing-8           6278            192517 ns/op                82.74 Mops/s        6591 B/op         98 allocs/op
BenchmarkByteProcessing-8           6279            160522 ns/op                99.21 Mops/s        6591 B/op         98 allocs/op
BenchmarkByteProcessing-8           5919            182307 ns/op                92.67 Mops/s        6496 B/op         98 allocs/op
BenchmarkByteProcessing-8           6784            161409 ns/op                91.32 Mops/s        6550 B/op         98 allocs/op
BenchmarkByteProcessing-8           6258            162812 ns/op                98.15 Mops/s        6592 B/op         98 allocs/op
BenchmarkByteProcessing-8           6234            163855 ns/op                97.90 Mops/s        6594 B/op         98 allocs/op
BenchmarkByteProcessing-8           5823            181810 ns/op                94.46 Mops/s        6503 B/op         98 allocs/op
BenchmarkByteProcessing-8           4544            226392 ns/op                97.21 Mops/s        6499 B/op         98 allocs/op
BenchmarkByteProcessing-8           5347            198100 ns/op                94.41 Mops/s        6543 B/op         98 allocs/op
BenchmarkByteProcessing-8           5827            171721 ns/op                99.94 Mops/s        6503 B/op         98 allocs/op
PASS
ok      github.com/pt-main/lc/example/tests/speedtest/byte/bench        76.599s
*/
