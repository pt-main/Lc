# Speed tests

## byte/bench/byte_test.go

Bytecode raw hotloop speed with one instruction without auto bytecode instruction index shift. 

<details> <summary>Result</summary>

```log
ITERATIONS=1_000_000_000:
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/speedtest/byte/bench -count=10
Lc version - 1.5.1
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/speedtest/byte/bench
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
ok      github.com/pt-main/lc/example/speedtest/byte/bench   63.261s

ITERATIONS=500_000_000:
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/speedtest/byte/bench -count=10
Lc version - 1.5.1
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/speedtest/byte/bench
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
ok      github.com/pt-main/lc/example/speedtest/byte/bench   28.963s

ITERATIONS=300_000_000:
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/speedtest/byte/bench -count=10
Lc version - 1.5.1
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/speedtest/byte/bench
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
ok      github.com/pt-main/lc/example/speedtest/byte/bench   18.120s


ITERATIONS=100_000_000:
macbook@MacBook-Pro lc % go test -test.fullpath=true -benchmem -run=^$ -bench ^BenchmarkByteProcessing$ github.com/pt-main/lc/example/speedtest/byte/bench -count=10
Lc version - 1.5.1
goos: darwin
goarch: amd64
pkg: github.com/pt-main/lc/example/speedtest/byte/bench
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
ok      github.com/pt-main/lc/example/speedtest/byte/bench   92.990s
```

</details>

## byte/tests/main.go

Bytecode raw hotloop speed with small handler - 

```go
func(be *engine.ByteEngine, pb *byteParsing.ParsedBytes) error {
    allIters += 1
    if allIters >= (ITERATIONS - 1) {
        timeE = time.Now()
        *idx = -1
    }
    return nil
}
```

- `Test1` - test with pre generated instructions
- `Test2` - test with one instruction without auto bytecode instruction index shift. 

<details> <summary>Result</summary>

```log
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
```

</details>