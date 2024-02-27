package main

import (
	"context"
	"fmt"
	"testing"

	vsstable "github.com/KevinJoiner/vss-translator/internal/generated/vss"
	"github.com/KevinJoiner/vss-translator/internal/model"
	"github.com/KevinJoiner/vss-translator/pkg/convert"
	"github.com/KevinJoiner/vss-translator/pkg/service/clickhouse"
	"github.com/elastic/go-elasticsearch/v8"
)

var testVehicles = 1025

func BenchmarkClickHouseVsElasstic(b *testing.B) {
	esModule, chModule, chService, err := startContainers()
	if err != nil {
		b.Fatalf("failed to start containers: %v", err)
	}
	defer esModule.Cleanup()
	defer chModule.Cleanup()

	vehicles, err := createVehicle(testVehicles)
	if err != nil {
		b.Fatalf("failed to create vehicles: %v", err)
	}
	const maxInsertCount = 1024
	for insertionCount := 1; insertionCount <= maxInsertCount; insertionCount *= 2 {
		name := fmt.Sprintf("Insert%dVehicles", insertionCount)
		b.Run(name, func(b *testing.B) {
			b.Run("DB=ClickHouse", func(b *testing.B) {
				// Test the insertion of 100 vehicles into clickhouse
				vssVehicles := make([]*vsstable.Vehicle, insertionCount)
				for i := range vssVehicles {
					vssVehicles[i] = convert.ESStatusToCHVehicle(vehicles[i])
				}
				insertionBenchCH(b, chService, vssVehicles)
			})
			b.Run("DB=Elastic", func(b *testing.B) {
				// Test the insertion of 100 vehicles into elastic
				insertionBenchEs(b, esModule.Client, vehicles[:insertionCount])
			})
		})
	}
}

func insertionBenchCH(b *testing.B, service *clickhouse.Service, vehicles []*vsstable.Vehicle) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := service.InsertVehicles(context.Background(), vehicles)
		if err != nil {
			b.Fatalf("failed to insert vehicles: %v", err)
		}
	}
}

func insertionBenchEs(b *testing.B, client *elasticsearch.TypedClient, vehicles []*model.DataPoint) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := insertVehiclesIntoElatic(context.Background(), client, vehicles)
		if err != nil {
			b.Fatalf("failed to insert vehicles: %v", err)
		}
	}
}
