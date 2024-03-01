package main

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/KevinJoiner/vss-translator/internal/generated/vsstable"
	"github.com/KevinJoiner/vss-translator/internal/model"
	"github.com/KevinJoiner/vss-translator/pkg/containers"
	"github.com/KevinJoiner/vss-translator/pkg/convert"
	"github.com/KevinJoiner/vss-translator/pkg/fake"
	"github.com/KevinJoiner/vss-translator/pkg/service/clickhouse"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const numberOfVehicles = 1000000

func main() {
	var vehicleFile string
	flag.StringVar(&vehicleFile, "vehicles", "./vehicles.json.gz", "File containing the vehicles to insert")
	flag.Parse()

	// start elastic container
	esModule, chModule, chService, err := startContainers()
	if err != nil {
		fmt.Printf("failed to start containers: %v", err)
		return
	}
	defer esModule.Cleanup()
	defer chModule.Cleanup()
	fmt.Printf("Containers started\n")

	start := time.Now()
	fmt.Println("Creating Vehicles")
	vehicles, err := getVehicles(numberOfVehicles, vehicleFile)
	chVehicles := make([]*vsstable.Vehicle, len(vehicles))
	for i := range chVehicles {
		chVehicles[i] = convert.ESStatusToCHVehicle(vehicles[i])
	}
	elapsed := time.Since(start)
	fmt.Printf("Creating %d vehicles took: %v\n", numberOfVehicles, elapsed)

	fmt.Println("Inserting Vehicles into ClickHouse")
	start = time.Now()
	err = chService.InsertVehicles(context.Background(), chVehicles)
	if err != nil {
		fmt.Printf("failed to insert vehicles into clickhouse: %v", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("Inserting %d vehicles into ClickHouse took: %v\n", numberOfVehicles, elapsed)

	fmt.Println("Inserting Vehicles into Elastic")
	start = time.Now()
	err = insertVehiclesIntoElaticLimited(context.Background(), esModule.Client, vehicles)
	if err != nil {
		fmt.Printf("failed to insert vehicles into elastic: %v", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("Inserting %d vehicles into Elastic took: %v\n", numberOfVehicles, elapsed)
	for _ = range 100 {
		_, _ = chService.GetVehiclesBySubjectOrderByTime(context.Background(), "vehicle-0", 100)
	}
	fmt.Println("Waiting for ctrl+c")
	// wait for exit signal to terminate the containers
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("Cya Later!")
}

func startContainers() (*containers.Elastic, *containers.ClickHouse, *clickhouse.Service, error) {
	elastic, err := containers.NewElasticContainer(context.TODO())
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create Elastic container: %w", err)
	}
	containers.AddStatusMapping(context.Background(), elastic.Client)

	chModule, err := containers.NewClickHouseContainer(context.TODO())
	if err != nil {
		elastic.Cleanup()
		return nil, nil, nil, fmt.Errorf("failed to create ClickHouse container: %w", err)
	}
	host, err := chModule.Container.ConnectionHost(context.TODO())
	if err != nil {
		chModule.Cleanup()
		elastic.Cleanup()
		return nil, nil, nil, fmt.Errorf("failed to get ClickHouse host: %w", err)
	}

	chService, err := clickhouse.New(context.TODO(), host)
	if err != nil {
		chModule.Cleanup()
		elastic.Cleanup()
		return nil, nil, nil, fmt.Errorf("failed to create ClickHouse service: %w", err)
	}
	err = chService.CreateVSSTable(context.Background())
	if err != nil {
		chService.Close()
		chModule.Cleanup()
		elastic.Cleanup()
		return nil, nil, nil, fmt.Errorf("failed to create VSS table in Clickhou: %w", err)
	}
	return elastic, chModule, chService, nil
}

// creatVehicle creates a new fake elatic status
func getVehicles(numOfVehicles int, vehicleFile string) ([]*model.DataPoint, error) {
	// create a thousand fake vehicles
	vehicles := make([]*model.DataPoint, numOfVehicles)
	if _, err := os.Stat(vehicleFile); err == nil {
		vehicles, err := loadVehiclesFromFile(vehicleFile)
		if err != nil {
			return nil, fmt.Errorf("failed to load vehicles from file: %w", err)
		}
		return vehicles, nil
	} else {
		fake.Setup()
		for i := range numOfVehicles {
			vehicle, err := fake.ElaticStatus()
			if err != nil {
				return nil, fmt.Errorf("failed to create fake vehicle: %w", err)
			}
			vehicle.Subject = fmt.Sprintf("vehicle-%d", i%10)
			vehicle.Data.Model = fmt.Sprintf("model-%d", i%100)
			vehicles[i] = vehicle
		}
	}

	return vehicles, nil
}

// insertVehiclesIntoElatic inserts the vehicles into elatic.
func insertVehiclesIntoElatic(ctx context.Context, client *elasticsearch.TypedClient, vehicles []*model.DataPoint) error {
	// insert the vehicles into elatic
	bulk := client.Bulk().Index(containers.StatusIndex)
	for _, vehicle := range vehicles {
		err := bulk.CreateOp(types.CreateOperation{}, vehicle)
		if err != nil {
			return fmt.Errorf("failed to add vehicle create op: %w", err)
		}
	}
	_, err := bulk.Do(ctx)
	if err != nil {
		return fmt.Errorf("failed to bulk insert vehicles into elatic: %w", err)
	}
	return nil
}

// insertVehiclesIntoElatic inserts the vehicles into elatic.
func insertVehiclesIntoElaticVSS(ctx context.Context, client *elasticsearch.TypedClient, vehicles []*vsstable.Vehicle) error {
	// insert the vehicles into elatic
	bulk := client.Bulk().Index(containers.VSSIndex)
	for _, vehicle := range vehicles {
		err := bulk.CreateOp(types.CreateOperation{}, vehicle)
		if err != nil {
			return fmt.Errorf("failed to add vehicle create op: %w", err)
		}
	}
	_, err := bulk.Do(ctx)
	if err != nil {
		return fmt.Errorf("failed to bulk insert vehicles into elatic: %w", err)
	}
	return nil
}

// insertVehiclesIntoElatic inserts the vehicles into elatic.
func insertVehiclesIntoElaticLimited(ctx context.Context, client *elasticsearch.TypedClient, vehicles []*model.DataPoint) error {
	// insert the vehicles into elatic
	bulk := client.Bulk().Index(containers.StatusIndex)
	for i, vehicle := range vehicles {
		err := bulk.CreateOp(types.CreateOperation{}, vehicle)
		if err != nil {
			return fmt.Errorf("failed to add vehicle create op: %w", err)
		}
		if i%1000 == 0 && i != 0 {
			resp, err := bulk.Do(ctx)
			if err != nil {
				return fmt.Errorf("failed to bulk insert vehicles into elatic: %w", err)
			}
			if resp.Errors {
				return fmt.Errorf("bulk insert had errors: %v", resp.Items)
			}
			bulk = client.Bulk().Index(containers.StatusIndex)
		}
	}
	_, err := bulk.Do(ctx)
	if err != nil {
		return fmt.Errorf("failed to bulk insert vehicles into elatic: %w", err)
	}
	return nil
}

func loadVehiclesFromFile(filePath string) ([]*model.DataPoint, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	vehicles := make([]*model.DataPoint, 0)
	// decompress the file before decoding
	gz, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gz.Close()
	decoder := json.NewDecoder(gz)
	err = decoder.Decode(&vehicles)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}
	return vehicles, nil
}
