package main

import (
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/KevinJoiner/vss-translator/internal/model"
	"github.com/KevinJoiner/vss-translator/pkg/fake"
)

func main() {
	var numOfVehicles int
	var outputfile string
	var nozip bool
	flag.IntVar(&numOfVehicles, "vehicles", 100, "Number of vehicles to generate")
	flag.StringVar(&outputfile, "output", "vehicles.json.gz", "Output file for the generated vehicles")
	flag.BoolVar(&nozip, "nozip", false, "Do not compress the output file")
	flag.Parse()
	fake.Setup()
	// create a thousand fake vehicles
	start := time.Now()
	vehicles := make([]*model.DataPoint, numOfVehicles)
	for i := range numOfVehicles {
		vehicle, err := fake.ElaticStatus()
		if err != nil {
			fmt.Printf("failed to create fake vehicle: %v\n", err)
			return
		}
		vehicle.Subject = fmt.Sprintf("vehicle-%d", i%10)
		vehicle.Data.Model = fmt.Sprintf("model-%d", i%100)
		vehicles[i] = vehicle
	}
	elapsed := time.Since(start)
	fmt.Printf("Creating %d vehicles took: %v\n", numOfVehicles, elapsed)
	start = time.Now()
	// write the vehicles to a file
	err := writeVehiclesToFile(outputfile, vehicles, nozip)
	if err != nil {
		fmt.Printf("failed to write vehicles to file: %v\n", err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("Writing %d vehicles to file took: %v\n", numOfVehicles, elapsed)

	start = time.Now()
	// // load the vehicles from the file
	// loadedVehicles, err := loadVehiclesFromFile(outputfile)
	// if err != nil {
	// 	fmt.Printf("failed to load vehicles from file: %v\n", err)
	// 	return
	// }
	// elapsed = time.Since(start)
	// fmt.Printf("Loading %d vehicles from file took: %v\n", numOfVehicles, elapsed)
	// fmt.Printf("Loaded %d vehicles from file\n", len(loadedVehicles))

}

func writeVehiclesToFile(filePath string, vehicles []*model.DataPoint, nozip bool) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	// compress the vehicles before writing to file
	if nozip {
		encoder := json.NewEncoder(file)
		err = encoder.Encode(vehicles)
		if err != nil {
			return fmt.Errorf("failed to encode json: %w", err)
		}
		return nil
	}
	gz := gzip.NewWriter(file)
	defer gz.Close()
	encoder := json.NewEncoder(gz)
	err = encoder.Encode(vehicles)
	if err != nil {
		return fmt.Errorf("failed to encode json: %w", err)
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
