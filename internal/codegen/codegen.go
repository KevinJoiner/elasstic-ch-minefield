package codegen

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"golang.org/x/tools/imports"
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

type TemplateData struct {
	Signals     []*SignalInfo
	DataSignals []*SignalInfo
}

func GenerateCode(signals []*SignalInfo, translations *Translations, outputDir string) (err error) {
	// Sort the signals by name
	slices.SortStableFunc(signals, func(a, b *SignalInfo) int {
		return strings.Compare(a.Name, b.Name)
	})

	dataSignals := []*SignalInfo{}
	for _, signal := range signals {
		trans, ok := translations.FromVspecName[signal.CHName]
		if !ok || signal.DataType == "" {
			continue
		}
		if trans.ClickHouseType != "" {
			signal.BaseCHType = trans.ClickHouseType
		}
		if trans.GoType != "" {
			signal.BaseGoType = trans.GoType
		}
		if trans.IsArray != nil {
			signal.IsArray = *trans.IsArray
		}

		dataSignals = append(dataSignals, signal)
	}
	data := TemplateData{
		Signals:     signals,
		DataSignals: dataSignals,
	}
	// Create a new template and parse the struct template
	tmpl, err := template.New("structTemplate").Funcs(template.FuncMap{
		"sqlFileName": func() string { return clickhouseFileName },
	}).Parse(structTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	// Create a new ClickHouse table template
	clickhouseTableTmpl, err := template.New("clickhouseTableTemplate").Funcs(template.FuncMap{
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

	var outBuf bytes.Buffer
	if err := tmpl.Execute(&outBuf, &data); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}
	formatted, err := imports.Process("", outBuf.Bytes(), &imports.Options{
		AllErrors: true, Comments: true, TabIndent: true, TabWidth: 4,
	})
	if err != nil {
		return fmt.Errorf("error formatting source: %w", err)
	}
	_, err = goOutputFile.Write(formatted)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
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

	err = clickhouseTableTmpl.Execute(clickhouseTableOutputFile, &data)
	if err != nil {
		return fmt.Errorf("error executing ClickHouse table template: %w", err)
	}

	return nil
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

func escapeDesc(desc string) string {
	return strings.ReplaceAll(desc, `'`, `\'`)
}
