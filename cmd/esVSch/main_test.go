package main

import (
	"context"
	"flag"
	"fmt"
	"testing"
	"time"

	"github.com/KevinJoiner/vss-translator/internal/generated/vsstable"
	"github.com/KevinJoiner/vss-translator/pkg/convert"
)

const (
	testSubject  = "vehicle-0"
	testModel    = "model-0"
	testVehicles = 10000
)

var (
	// one month time range
	testStartTime = time.Date(2019, 12, 1, 0, 0, 0, 0, time.UTC)
	testEndTime   = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	vehicleFile   = ""
)

func TestMain(m *testing.M) {
	flag.StringVar(&vehicleFile, "vehicles", "./vehicles.json.gz", "File containing the vehicles to insert")
	flag.Parse()
	m.Run()
}

func BenchmarkClickHouseVsElasstic(b *testing.B) {
	esModule, chModule, chService, err := startContainers()
	if err != nil {
		b.Fatalf("failed to start containers: %v", err)
	}
	defer esModule.Cleanup()
	defer chModule.Cleanup()

	vehicles, err := getVehicles(testVehicles, vehicleFile)
	if err != nil {
		b.Fatalf("failed to create vehicles: %v", err)
	}
	vssVehicles := make([]*vsstable.Vehicle, len(vehicles))
	for i := range vssVehicles {
		vssVehicles[i] = convert.ESStatusToCHVehicle(vehicles[i])
	}
	lastDoc := 0
	const maxInsertCount = 1024
	for insertionCount := 1; insertionCount <= maxInsertCount; insertionCount *= 2 {
		name := fmt.Sprintf("Insert%dVehicles", insertionCount)
		b.Run(name, func(b *testing.B) {
			b.Run("DB=ClickHouseStruct", func(b *testing.B) {
				insertionBenchCHStruct(b, chService, vssVehicles[lastDoc:insertionCount])
			})
			// b.Run("DB=ClickHouseProtoCol", func(b *testing.B) {
			// 	insertionBenchCHProtoCol(b, chService, vssVehicles[lastDoc:insertionCount])
			// })
			// b.Run("DB=ClickHouseCol", func(b *testing.B) {
			// 	insertionBenchCHCol(b, chService, vssVehicles[lastDoc:insertionCount])
			// })
			// b.Run("DB=ClickHouseJSON", func(b *testing.B) {
			// 	insertionBenchCHJSON(b, chService, vssVehicles[lastDoc:insertionCount])
			// })
			// b.Run("DB=ClickHouseSubjectAndTime", func(b *testing.B) {
			// 	insertionBenchCHColSubNTime(b, chService, vssVehicles[lastDoc:insertionCount])
			// })
			// b.Run("DB=Elastic", func(b *testing.B) {
			// 	insertionBenchEs(b, esModule.Client, vehicles[lastDoc:insertionCount])
			// })
			b.Run("DB=ElasticVSS", func(b *testing.B) {
				insertionBenchEsVSS(b, esModule.Client, vssVehicles[lastDoc:insertionCount])
			})
		})
		lastDoc = insertionCount
	}
	const maxSelectCount = 10000
	fmt.Printf("Inserting %d vehicles into Elastic\n", len(vehicles))
	err = insertVehiclesIntoElaticLimited(context.TODO(), esModule.Client, vehicles)
	if err != nil {
		b.Fatalf("failed to insert vehicles into elastic: %v", err)
	}

	fmt.Printf("Inserting %d vehicles into ClickHouse\n", len(vssVehicles))
	err = chService.InsertVehicles(context.TODO(), vssVehicles)
	if err != nil {
		b.Fatalf("failed to insert vehicles into clickhouse: %v", err)
	}

	b.Run("SelectMaxSpeedBySubjectInTimeRange", func(b *testing.B) {
		b.Run("DB=ClickHouse", func(b *testing.B) {
			selectBenchCHMaxSpeedBySubjectInTimeRange(b, chService, testSubject, testStartTime.Unix(), testEndTime.Unix())
		})
		b.Run("DB=Elastic", func(b *testing.B) {
			selectBenchEsMaxSpeedBySubjectInTimeRange(b, esModule.Client, testSubject, testStartTime.UnixMilli(), testEndTime.UnixMilli())
		})
	})
	// Run Seleccts with different Limits
	for selectCount := 100; selectCount <= maxSelectCount; selectCount *= 10 {
		name := fmt.Sprintf("SelectAny%dVehicles", selectCount)
		b.Run(name, func(b *testing.B) {
			b.Run("DB=ClickHouse", func(b *testing.B) {
				selectBenchCH(b, chService, selectCount)
			})
			b.Run("DB=Elastic", func(b *testing.B) {
				selectBenchEs(b, esModule.Client, selectCount)
			})
		})
		name = fmt.Sprintf("SelectSubject%dVehicles", selectCount)
		b.Run(name, func(b *testing.B) {
			b.Run("DB=ClickHouse", func(b *testing.B) {
				selectBenchCHSubject(b, chService, testSubject, selectCount)
			})
			b.Run("DB=Elastic", func(b *testing.B) {
				selectBenchEsSubject(b, esModule.Client, testSubject, selectCount)
			})
		})
		name = fmt.Sprintf("SelectSubjectOrderByTime%dVehicles", selectCount)
		b.Run(name, func(b *testing.B) {
			b.Run("DB=ClickHouse", func(b *testing.B) {
				selectBenchCHSubjectOrderByTime(b, chService, testSubject, selectCount)
			})
			b.Run("DB=Elastic", func(b *testing.B) {
				selectBenchEsSubjectOrderByTime(b, esModule.Client, testSubject, selectCount)
			})
		})
		name = fmt.Sprintf("SelectSubjectAndBrandOrderOdmeter%dVehicles", selectCount)
		b.Run(name, func(b *testing.B) {
			b.Run("DB=ClickHouse", func(b *testing.B) {
				selectBenchCHSubjectAndBrandOrderOdmeter(b, chService, testSubject, testModel, selectCount)
			})
			b.Run("DB=Elastic", func(b *testing.B) {
				selectBenchEsSubjectAndBrandOrderOdmeter(b, esModule.Client, testSubject, testModel, selectCount)
			})
		})
	}
}
