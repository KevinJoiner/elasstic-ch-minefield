package codegen

import (
	"encoding/csv"
	"os"
	"strings"
)

// SignalInfo represents the structure of each row in the CSV file.
type SignalInfo struct {
	Name       string
	Type       string
	DataType   string
	Deprecated bool
	Unit       string
	Min        string
	Max        string
	Desc       string
}

const (
	nameCol       = 0
	typeCol       = 1
	dataTypeCol   = 2
	deprecatedCol = 3
	unitCol       = 4
	minCol        = 5
	maxCol        = 6
	descCol       = 7
)

func ReadSignalsCSV(filePath string) ([]SignalInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var signals []SignalInfo
	for i := 1; i < len(records); i++ {
		record := records[i]
		signal := SignalInfo{
			Name:       record[nameCol],
			Type:       record[typeCol],
			DataType:   record[dataTypeCol],
			Deprecated: record[deprecatedCol] == "true", // Assuming "true" or "false" string in the CSV
			Unit:       record[unitCol],
			Min:        record[minCol],
			Max:        record[maxCol],
			Desc:       record[descCol],
		}
		signals = append(signals, signal)
	}

	return signals, nil
}

func escapeDesc(desc string) string {
	return strings.ReplaceAll(desc, `'`, `\'`)
}

func chName(name string) string {
	return strings.ReplaceAll(name, ".", "_")
}
