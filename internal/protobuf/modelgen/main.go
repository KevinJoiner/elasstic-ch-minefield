package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/ClickHouse/ch-go"
	"github.com/ClickHouse/ch-go/proto"
)

var goNameReplacer = strings.NewReplacer(".", "", "_", "")

func main() {
	ctx := context.Background()
	client, err := ch.Dial(ctx, ch.Options{Address: "localhost:9000"})
	if err != nil {
		fmt.Printf("failed to connect to clickhouse: %v\n", err)
		return
	}
	defer client.Close()
	schemas, err := getSchemas(ctx, client)
	if err != nil {
		fmt.Printf("failed to get schemas: %v\n", err)
		return
	}

	err = Generate(schemas, "pkg/vss/fields.go")
	if err != nil {
		fmt.Printf("failed to generate file: %v\n", err)
		return
	}
}

type Schema struct {
	ColName string
	CHType  string
	GOName  string
	IsTuple bool
}

func getSchemas(ctx context.Context, client *ch.Client) ([]Schema, error) {
	var (
		nameData proto.ColStr
		typeData proto.ColStr
	)

	schemas := []Schema{}
	err := client.Do(ctx, ch.Query{
		Body: "DESCRIBE TABLE vss SETTINGS describe_include_subcolumns=1",
		Result: proto.Results{
			{Name: "name", Data: &nameData},
			{Name: "type", Data: &typeData},
			{Name: "default_type", Data: &proto.ColAuto{}},
			{Name: "default_expression", Data: &proto.ColAuto{}},
			{Name: "comment", Data: &proto.ColAuto{}},
			{Name: "codec_expression", Data: &proto.ColAuto{}},
			{Name: "ttl_expression", Data: &proto.ColAuto{}},
			{Name: "is_subcolumn", Data: &proto.ColAuto{}},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	nameData.ForEach(func(i int, name string) error {
		chType := typeData.Row(i)
		sh := Schema{
			ColName: name,
			CHType:  chType,
			GOName:  goNameReplacer.Replace(name),
			IsTuple: strings.HasPrefix(chType, "Tuple"),
		}
		schemas = append(schemas, sh)
		return nil

	})
	return schemas, nil
}
