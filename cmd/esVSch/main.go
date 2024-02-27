package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/KevinJoiner/vss-translator/internal/model"
	"github.com/KevinJoiner/vss-translator/pkg/containers"
	"github.com/KevinJoiner/vss-translator/pkg/convert"
	"github.com/KevinJoiner/vss-translator/pkg/fake"
	"github.com/elastic/go-elasticsearch/v8"
)

const numberOfVehicles = 50000

func main() {
	// start elastic container
	fmt.Println("Starting Elastic Container")
	elastic, err := containers.NewElasticContainer(context.TODO())
	if err != nil {
		fmt.Printf("failed to create Elastic container: %v", err)
		return
	}
	defer elastic.Cleanup()
	containers.AddStatusMapping(context.Background(), elastic.Client)

	fmt.Println("Starting ClickHouse Container")
	clickhouse, err := containers.NewClickHouseContainer(context.TODO())
	if err != nil {
		fmt.Printf("failed to create ClickHouse container: %v", err)
		return
	}
	defer clickhouse.Cleanup()
	err = containers.CreateVSSTable(context.Background(), clickhouse.Client)
	if err != nil {
		fmt.Printf("failed to create VSS table: %v", err)
		return
	}

	fmt.Println("Containers started")
	start := time.Now()
	fmt.Println("Creating Vehicles")
	vehicles, err := createVehicle()
	if err != nil {
		fmt.Printf("failed to create vehicles: %v", err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("Creating %d vehicles took: %v\n", numberOfVehicles, elapsed)

	fmt.Println("Inserting Vehicles into ClickHouse")
	start = time.Now()
	err = insertVehiclesIntoClickhouse(context.Background(), vehicles, clickhouse.Client)
	if err != nil {
		fmt.Printf("failed to insert vehicles into clickhouse: %v", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("Inserting %d vehicles into ClickHouse took: %v\n", numberOfVehicles, elapsed)

	fmt.Println("Inserting Vehicles into Elastic")
	start = time.Now()
	err = insertVehiclesIntoElatic(context.Background(), elastic.Client, vehicles)
	if err != nil {
		fmt.Printf("failed to insert vehicles into elastic: %v", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("Inserting %d vehicles into Elastic took: %v\n", numberOfVehicles, elapsed)

	fmt.Println("Waiting for ctrl+c")
	// wait for exit signal to terminate the containers
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("Cya Later!")
}

// creatVehicle creates a new fake elatic status
func createVehicle() ([]*model.DataPoint, error) {
	fake.Setup()
	// create a thousand fake vehicles
	vehicles := make([]*model.DataPoint, numberOfVehicles)
	for i := range numberOfVehicles {
		vehicle, err := fake.ElaticStatus()
		if err != nil {
			return nil, fmt.Errorf("failed to create fake vehicle: %w", err)
		}
		vehicles[i] = vehicle
	}
	return vehicles, nil
}

// writeProtoVehicles writes the proto encoded vehicles to a tmp file
func insertVehiclesIntoClickhouse(ctx context.Context, vehicles []*model.DataPoint, chConn driver.Conn) error {
	batch, err := chConn.PrepareBatch(ctx, "INSERT into vss")
	if err != nil {
		return fmt.Errorf("failed to prepare batch: %w", err)
	}
	for _, vehicle := range vehicles {
		err := batch.AppendStruct(convert.ESStatusToCHVehicle(vehicle))
		if err != nil {
			return fmt.Errorf("failed to append vehicle to batch: %w", err)
		}
	}
	err = batch.Send()
	if err != nil {
		return fmt.Errorf("failed to send batch: %w", err)
	}
	return nil
}

// insertVehiclesIntoElatic inserts the vehicles into elatic.
func insertVehiclesIntoElatic(ctx context.Context, client *elasticsearch.TypedClient, vehicles []*model.DataPoint) error {
	// insert the vehicles into elatic
	for _, vehicle := range vehicles {
		_, err := client.Index(containers.StatusIndex).Document(vehicle).Do(ctx)
		if err != nil {
			return fmt.Errorf("failed to insert vehicle: %w", err)
		}
	}
	return nil
}
