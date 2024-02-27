package main

import (
	"cmp"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"text/template"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
)

// VSSColumns is a map of column names to their corresponding column types.
// the names are flattened with a dot separator.
var VSSColumns = map[string]column.Interface{}

const templateText = `// Code generated from clickhouse vss table DO NOT EDIT.
package vss

import "github.com/ClickHouse/ch-go/proto"

const (
{{- range .}}
	// Field{{.GOName}} is the name for the {{.ColName}} column in the Clickhouse.
	Field{{.GOName}} string = "{{.ColName}}"
{{- end}}
)
var (
	newColsFunc = map[string]func() proto.Column{
{{- range . -}}
		{{if .IsTuple}}{{- continue -}}{{end}}
		Field{{.GOName}}: MustGetColumn("{{.CHType}}"),
{{- end}}
	}
)
`

// Generate reads the JSON file at filePath and generates a Go file with the structs
func Generate(schemas []Schema, dest string) error {
	// Create a new Go file
	file, err := os.Create(filepath.Clean(dest))
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Execute the template
	tmpl, err := template.New("generateFields").Parse(templateText)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}
	slices.SortStableFunc(schemas, func(a Schema, b Schema) int {
		return cmp.Compare(a.GOName, b.GOName)
	})
	err = tmpl.Execute(file, schemas)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	return nil
}

func getProtoType(dataType string) string {
	// Implement your logic for converting dataType to the corresponding proto type
	// For simplicity, this example assumes that the dataType is directly usable as a proto type.
	return dataType
}
