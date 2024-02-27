package codegen

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"text/template"
)

const (
	clickhouseFileName = "vss-table.sql"
	goFileName         = "vss-structs.go"
	readAll            = 0755
)

// embedd structTemplate with go:embed
//
//go:embed vehicle.tmpl
var structTemplate string

//go:embed vssTable.tmpl
var clickhouseTableTemplate string

func GenerateCode(signals []SignalInfo, outputDir string) (err error) {
	// Sort the signals by name
	slices.SortStableFunc(signals, func(a, b SignalInfo) int {
		return strings.Compare(a.Name, b.Name)
	})

	// Create a new template and parse the struct template
	tmpl, err := template.New("structTemplate").Funcs(template.FuncMap{
		"goName":      goName,
		"chName":      chName,
		"toGOType":    toGOType,
		"sqlFileName": func() string { return clickhouseFileName },
	}).Parse(structTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	// Create a new ClickHouse table template
	clickhouseTableTmpl, err := template.New("clickhouseTableTemplate").Funcs(template.FuncMap{
		"chName":     chName,
		"toCHType":   toCHType,
		"escapeDesc": escapeDesc,
	}).Parse(clickhouseTableTemplate)
	if err != nil {
		return fmt.Errorf("error parsing ClickHouse table template: %w", err)
	}
	_, err = os.Stat(outputDir)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("error checking output directory: %w", err)
		}
		// create the output directory
		err = os.MkdirAll(outputDir, readAll)
		if err != nil {
			return fmt.Errorf("error creating output directory: %w", err)
		}
	}

	// Execute the template for the signals
	goOutputPath := filepath.Join(outputDir, goFileName)
	goOutputFile, err := os.Create(goOutputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer func() {
		if cerr := goOutputFile.Close(); err == nil && cerr != nil {
			err = cerr
		}
	}()

	err = tmpl.Execute(goOutputFile, signals)
	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	// Execute the ClickHouse table template for the signals
	clickhouseTableOutputPath := filepath.Join(outputDir, clickhouseFileName)
	clickhouseTableOutputFile, err := os.Create(clickhouseTableOutputPath)
	if err != nil {
		return fmt.Errorf("error creating ClickHouse table output file: %w", err)
	}
	defer func() {
		if cerr := clickhouseTableOutputFile.Close(); err == nil && cerr != nil {
			err = cerr
		}
	}()

	err = clickhouseTableTmpl.Execute(clickhouseTableOutputFile, signals)
	if err != nil {
		return fmt.Errorf("error executing ClickHouse table template: %w", err)
	}

	return nil
}

func goName(name string) string {
	// Remove special characters and ensure PascalCase
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	return re.ReplaceAllString(name, "")
}

// toGOType converts vspec type to golang types
func toGOType(dataType string) string {
	// if name is ends with [] move it to the front of the type
	baseType := dataType
	prefix := ""
	if strings.HasSuffix(dataType, "[]") {
		baseType = dataType[:len(dataType)-2]
		prefix = "[]"
	}
	switch baseType {
	case "uint8", "int8", "uint16", "int16", "uint32", "int32", "uint64", "int64", "string":
		return prefix + baseType
	case "boolean":
		return prefix + "bool"
	case "float":
		return prefix + "float32"
	case "double":
		return prefix + "float64"
	default:
		return prefix + "any"
	}
}

// toCHType converts vspec type to clickhouse types.
func toCHType(dataType string) string {
	// if name is ends with [] move it to the front of the type
	baseType := dataType
	// string[] -> Array(String)
	isArray := strings.HasSuffix(dataType, "[]")
	if isArray {
		baseType = dataType[:len(dataType)-2]
	}

	switch baseType {
	case "uint8":
		baseType = "UInt8"
	case "int8":
		baseType = "Int8"
	case "uint16":
		baseType = "UInt16"
	case "int16":
		baseType = "Int16"
	case "uint32":
		baseType = "UInt32"
	case "int32":
		baseType = "Int32"
	case "uint64":
		baseType = "UInt64"
	case "int64":
		baseType = "Int64"
	case "string":
		baseType = "String"
	case "boolean":
		baseType = "Bool"
	case "float":
		baseType = "Float32"
	case "double":
		baseType = "Float64"
	default:
		baseType = "String"
	}
	if isArray {
		baseType = "Array(" + baseType + ")"
	}
	return baseType
}
