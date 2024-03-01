package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/KevinJoiner/vss-translator/internal/model"
	"github.com/KevinJoiner/vss-translator/pkg/containers"
	"github.com/KevinJoiner/vss-translator/pkg/service/clickhouse"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

func selectBenchCH(b *testing.B, service *clickhouse.Service, count int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v, err := service.GetVehicleAny(context.Background(), count)
		if err != nil {
			b.Fatalf("failed to select vehicles: %v", err)
		}
		if len(v) != count {
			b.Fatalf("failed to select vehicles: expected %d got %d", count, len(v))
		}
	}

}

func selectBenchCHSubject(b *testing.B, service *clickhouse.Service, subject string, count int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v, err := service.GetVehiclesBySubject(context.Background(), subject, count)
		if err != nil {
			b.Fatalf("failed to select vehicles: %v", err)
		}
		if len(v) != count {
			b.Fatalf("failed to select vehicles: expected %d got %d", count, len(v))
		}
	}
}

func selectBenchCHSubjectOrderByTime(b *testing.B, service *clickhouse.Service, subject string, count int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v, err := service.GetVehiclesBySubjectOrderByTime(context.Background(), subject, count)
		if err != nil {
			b.Fatalf("failed to select vehicles: %v", err)
		}
		if len(v) != count {
			b.Fatalf("failed to select vehicles: expected %d got %d", count, len(v))
		}
	}
}

func selectBenchCHSubjectAndBrandOrderOdmeter(b *testing.B, service *clickhouse.Service, subject, model string, count int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v, err := service.GetVehiclesBySubjBrandOrderByOdm(context.Background(), subject, model, count)
		if err != nil {
			b.Fatalf("failed to select vehicles: %v", err)
		}
		if len(v) != count {
			b.Fatalf("failed to select vehicles: expected %d got %d", count, len(v))
		}
	}
}

func selectBenchCHMaxSpeedBySubjectInTimeRange(b *testing.B, service *clickhouse.Service, subject string, start, end int64) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := service.GetMaxSpeedBySubjectInTimeRange(context.Background(), subject, start, end)
		if err != nil {
			b.Fatalf("failed to select vehicles: %v", err)
		}
	}
}

func selectBenchEs(b *testing.B, client *elasticsearch.TypedClient, count int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v, err := getVehiclesFromElastic(context.Background(), client, count)
		if err != nil {
			b.Fatalf("failed to select vehicles: %v", err)
		}
		if len(v) != count {
			b.Fatalf("failed to select vehicles: expected %d got %d", count, len(v))
		}
	}
}

func selectBenchEsSubject(b *testing.B, client *elasticsearch.TypedClient, subject string, count int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v, err := getVehiclesFromElasticBySubject(context.Background(), client, subject, count)
		if err != nil {
			b.Fatalf("failed to select vehicles: %v", err)
		}
		if len(v) != count {
			b.Fatalf("failed to select vehicles: expected %d got %d", count, len(v))
		}
	}
}

func selectBenchEsSubjectOrderByTime(b *testing.B, client *elasticsearch.TypedClient, subject string, count int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v, err := getVehiclesFromElasticBySubjectOrderByTime(context.Background(), client, subject, count)
		if err != nil {
			b.Fatalf("failed to select vehicles: %v", err)
		}
		if len(v) != count {
			b.Fatalf("failed to select vehicles: expected %d got %d", count, len(v))
		}
	}
}

func selectBenchEsSubjectAndBrandOrderOdmeter(b *testing.B, client *elasticsearch.TypedClient, subject, model string, count int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v, err := getVehiclesFromElasticBySubjectOrderByTime(context.Background(), client, subject, count)
		if err != nil {
			b.Fatalf("failed to select vehicles: %v", err)
		}
		if len(v) != count {
			b.Fatalf("failed to select vehicles: expected %d got %d", count, len(v))
		}
	}
}

func selectBenchEsMaxSpeedBySubjectInTimeRange(b *testing.B, client *elasticsearch.TypedClient, subject string, start, end int64) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := getMaxSpeedBySubjectInTimeRange(context.Background(), client, subject, start, end)
		if err != nil {
			b.Fatalf("failed to select vehicles: %v", err)
		}
	}
}

func getVehiclesFromElastic(ctx context.Context, client *elasticsearch.TypedClient, count int) ([]*model.DataPoint, error) {
	vehicles := make([]*model.DataPoint, 0)
	resp, err := client.Search().Index(containers.StatusIndex).Size(count).Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to search vehicles: %w", err)
	}
	for _, hit := range resp.Hits.Hits {
		var vehicle model.DataPoint
		// if err := json.Unmarshal(hit.Source_, &vehicle); err != nil {
		// 	return nil, fmt.Errorf("failed to unmarshal vehicle: %w", err)
		// }
		_ = hit
		vehicles = append(vehicles, &vehicle)
	}
	return vehicles, nil
}

func getVehiclesFromElasticBySubject(ctx context.Context, client *elasticsearch.TypedClient, subject string, count int) ([]*model.DataPoint, error) {
	vehicles := make([]*model.DataPoint, 0)
	query := types.Query{
		Bool: &types.BoolQuery{
			Filter: []types.Query{
				{Term: map[string]types.TermQuery{"subject": {Value: subject}}},
			},
		},
	}
	resp, err := client.Search().Index(containers.StatusIndex).Query(&query).Size(count).Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to search vehicles: %w", err)
	}
	for _, hit := range resp.Hits.Hits {
		var vehicle model.DataPoint
		// if err := json.Unmarshal(hit.Source_, &vehicle); err != nil {
		// 	return nil, fmt.Errorf("failed to unmarshal vehicle: %w", err)
		// }
		_ = hit
		vehicles = append(vehicles, &vehicle)
	}
	return vehicles, nil
}

func getVehiclesFromElasticBySubjectOrderByTime(ctx context.Context, client *elasticsearch.TypedClient, subject string, count int) ([]*model.DataPoint, error) {
	vehicles := make([]*model.DataPoint, 0)
	query := types.Query{
		Bool: &types.BoolQuery{
			Filter: []types.Query{
				{Term: map[string]types.TermQuery{"subject": {Value: subject}}},
			},
		},
	}
	sort := types.SortCombinations(
		&types.SortOptions{
			SortOptions: map[string]types.FieldSort{"data.timestamp": {Order: &sortorder.Asc}},
		},
	)
	resp, err := client.Search().Index(containers.StatusIndex).Query(&query).Sort(sort).Size(count).Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to search vehicles: %w", err)
	}
	for _, hit := range resp.Hits.Hits {
		var vehicle model.DataPoint
		// if err := json.Unmarshal(hit.Source_, &vehicle); err != nil {
		// 	return nil, fmt.Errorf("failed to unmarshal vehicle: %w", err)
		// }
		_ = hit
		vehicles = append(vehicles, &vehicle)
	}
	return vehicles, nil
}

func getVehiclesFromElasticBySubjectBrandOrderByOdm(ctx context.Context, client *elasticsearch.TypedClient, subject, carModel string, count int) ([]*model.DataPoint, error) {
	vehicles := make([]*model.DataPoint, 0)
	query := types.Query{
		Bool: &types.BoolQuery{
			Filter: []types.Query{
				{Term: map[string]types.TermQuery{"subject": {Value: subject}, "data.model": {Value: carModel}}},
			},
		},
	}
	sort := types.SortCombinations(
		&types.SortOptions{
			SortOptions: map[string]types.FieldSort{"data.odometer": {Order: &sortorder.Desc}},
		},
	)
	resp, err := client.Search().Index(containers.StatusIndex).Query(&query).Sort(sort).Size(count).Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to search vehicles: %w", err)
	}
	for _, hit := range resp.Hits.Hits {
		var vehicle model.DataPoint
		// if err := json.Unmarshal(hit.Source_, &vehicle); err != nil {
		// 	return nil, fmt.Errorf("failed to unmarshal vehicle: %w", err)
		// }
		_ = hit
		vehicles = append(vehicles, &vehicle)
	}
	return vehicles, nil
}

func getMaxSpeedBySubjectInTimeRange(ctx context.Context, client *elasticsearch.TypedClient, subject string, start, end int64) (float64, error) {
	startArg := fmt.Sprintf("%d", start)
	endArg := fmt.Sprintf("%d", end)
	query := types.Query{
		Bool: &types.BoolQuery{
			Filter: []types.Query{
				{Term: map[string]types.TermQuery{"subject": {Value: subject}}},
				{Range: map[string]types.RangeQuery{"data.timestamp": types.DateRangeQuery{Gte: &startArg, Lte: &endArg}}},
			},
		},
	}
	// agg := types.Aggregation{
	// 	Max: map[string]types.MaxAggregation{"max_speed": {Field: "data.speed"}},
	// }
	speed := "data.speed"
	agg := map[string]types.Aggregations{
		"max_speed": {
			Max: &types.MaxAggregation{Field: &speed},
		},
	}
	resp, err := client.Search().Index(containers.StatusIndex).Query(&query).Aggregations(agg).Size(0).Do(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to search vehicles: %w", err)
	}
	maxAgg, ok := (resp.Aggregations["max_speed"]).(*types.MaxAggregate)
	if !ok {
		return 0, fmt.Errorf("failed to get max speed aggregate")
	}
	return float64(maxAgg.Value), nil
}
