package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/KevinJoiner/vss-translator/internal/generated/vsstable"
	"github.com/KevinJoiner/vss-translator/internal/model"
	"github.com/KevinJoiner/vss-translator/pkg/service/clickhouse"
	"github.com/elastic/go-elasticsearch/v8"
)

func insertionBenchCHCol(b *testing.B, service *clickhouse.Service, vehicles []*vsstable.Vehicle) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := service.InsertVehiclesCol(context.Background(), vehicles)
		if err != nil {
			b.Fatalf("failed to insert vehicles: %v", err)
		}
	}
}
func insertionBenchCHProtoCol(b *testing.B, service *clickhouse.Service, vehicles []*vsstable.Vehicle) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := service.InsertVehiclesProtoCol(context.Background(), vehicles)
		if err != nil {
			b.Fatalf("failed to insert vehicles: %v", err)
		}
	}
}

func insertionBenchCHStruct(b *testing.B, service *clickhouse.Service, vehicles []*vsstable.Vehicle) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := service.InsertVehicles(context.Background(), vehicles)
		if err != nil {
			b.Fatalf("failed to insert vehicles: %v", err)
		}
	}
}

func insertionBenchCHJSON(b *testing.B, service *clickhouse.Service, vehicles []*vsstable.Vehicle) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := service.InsertVehiclesJSON(context.Background(), vehicles)
		if err != nil {
			b.Fatalf("failed to insert vehicles: %v", err)
		}
	}
}

func insertionBenchCHColSubNTime(b *testing.B, service *clickhouse.Service, vehicles []*vsstable.Vehicle) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := service.InsertVehiclesColSubAndTime(context.Background(), vehicles)
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

func insertionBenchEsVSS(b *testing.B, client *elasticsearch.TypedClient, vehicles []*vsstable.Vehicle) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := insertVehiclesIntoElaticVSS(context.Background(), client, vehicles)
		if err != nil {
			b.Fatalf("failed to insert vehicles: %v", err)
		}
	}
}

// elasticTruncateTable truncates the elastic index
func elasticTruncateTable(ctx context.Context, client *elasticsearch.Client) error {
	_, err := client.Indices.Delete([]string{"vss"}, client.Indices.Delete.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("failed to delete index: %w", err)
	}
	return nil
}
