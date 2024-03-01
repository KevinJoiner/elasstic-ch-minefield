// TODO combine this with the field name generation.
// Update break up this file into logical parts
// Bench the insertion time of elastic vs clickhouse
// create a convert function from new essstatus to new struct
package main

import (
	_ "embed"
	"flag"
	"fmt"

	"github.com/KevinJoiner/vss-translator/internal/codegen"
)

func main() {
	// Command-line flags
	outputDir := flag.String("output", ".", "Output directory for the generated Go file")
	filePath := flag.String("csv", "./vspec.csv", "Path to the vspec CSV file")
	tansPath := flag.String("translations", "./translations.json", "Path to the translations JSON file")
	flag.Parse()

	// Read CSV file
	signals, err := codegen.ReadSignalsCSV(*filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Read translation file
	translations, err := codegen.ReadTranslationJSON(*tansPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = codegen.GenerateCode(signals, translations, *outputDir)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
