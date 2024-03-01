package codegen

import (
	"encoding/csv"
	"os"
)

// SignalInfo represents the structure of each row in the CSV file.

func ReadSignalsCSV(filePath string) ([]*SignalInfo, error) {
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

	var signals []*SignalInfo
	for i := 1; i < len(records); i++ {
		record := records[i]
		signals = append(signals, NewSignal(record))
	}

	return signals, nil
}
