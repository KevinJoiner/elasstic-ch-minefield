# minefield
Random assortment of test code for migrating Elastic -> CH nothing here is set in stone.

### cmd/csv2CH
This will convert a given CSV that was generated via vspec to a new CH table and matching go struct. This can be run with `go generate` via generate.go

### cmd/esVsch
This binary will start a generate fake vehicle data start a test elastic container and a test clickhouse container then insert the fake vehicle data into both databases in their respective formats that mimick what is in production. The program will wait for ctrl-c to exit so you can exec into the containers and run any test queries you may have.

#### Test
Run the following for test output
```bash
go install golang.org/x/perf/cmd/benchstat@latest
go test -run='^$' -bench=. -count=10 -cpu 1 > insertionBench.txt
benchstat -col  /DB insertionBench.txt
```
Output:
``` bash
goos: darwin
goarch: arm64
pkg: github.com/KevinJoiner/vss-translator/cmd/esVSch
                                           │  ClickHouse   │                 Elastic                  │
                                           │    sec/op     │     sec/op      vs base                  │
ClickHouseVsElasstic/Insert1Vehicles-14      19.501m ±  2%     1.767m ± 46%    -90.94% (p=0.000 n=10)
ClickHouseVsElasstic/Insert2Vehicles-14      19.865m ±  3%     2.457m ± 14%    -87.63% (p=0.000 n=10)
ClickHouseVsElasstic/Insert4Vehicles-14      20.097m ±  2%     5.003m ± 11%    -75.11% (p=0.000 n=10)
ClickHouseVsElasstic/Insert8Vehicles-14       22.88m ±  4%     10.61m ±  5%    -53.63% (p=0.000 n=10)
ClickHouseVsElasstic/Insert16Vehicles-14      24.75m ±  7%     20.07m ±  5%    -18.92% (p=0.000 n=10)
ClickHouseVsElasstic/Insert32Vehicles-14      27.26m ± 11%     39.77m ± 10%    +45.89% (p=0.000 n=10)
ClickHouseVsElasstic/Insert64Vehicles-14      29.72m ± 13%     84.61m ±  7%   +184.65% (p=0.000 n=10)
ClickHouseVsElasstic/Insert128Vehicles-14     32.88m ±  4%    162.10m ±  8%   +393.03% (p=0.000 n=10)
ClickHouseVsElasstic/Insert256Vehicles-14     37.02m ± 11%    369.29m ± 17%   +897.55% (p=0.000 n=10)
ClickHouseVsElasstic/Insert512Vehicles-14     54.43m ±  7%    615.04m ± 17%  +1030.05% (p=0.000 n=10)
ClickHouseVsElasstic/Insert1024Vehicles-14    80.87m ±  6%   1418.93m ±  7%  +1654.55% (p=0.000 n=10)
geomean                                       30.21m           42.48m          +40.61%

```

### cmd/es-ch 
This a main.go I am using to test various CH client code with the different libraries.

### internal/protbuf
This contains random code from when we were investigating inferring the schema from protobuf.

