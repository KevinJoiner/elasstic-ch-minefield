package containers

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/testcontainers/testcontainers-go"
	esmodule "github.com/testcontainers/testcontainers-go/modules/elasticsearch"
)

const (
	// StatusIndex is the name of the index for standard status data.
	StatusIndex = "test-index"

	// VSSIndex is the name of the index for VSS data.
	VSSIndex = "vss-index"
)

//go:embed config-files/vss_mapping.json
var vssMapping []byte

//go:embed config-files/status_mapping.json
var statusMapping []byte

type Elastic struct {
	Client    *elasticsearch.TypedClient
	Container *esmodule.ElasticsearchContainer
	cleanup   func()
}

func (e *Elastic) Cleanup() {
	e.cleanup()
}

// NewElasticContainer creates a new elastic container and a new elastic service with the client and the logger.
func NewElasticContainer(ctx context.Context) (*Elastic, error) {
	esContaianer, cleanup, err := CreateElasticContainer(ctx)
	if err != nil {
		cleanup()
		return nil, fmt.Errorf("failed to create Elasticsearch container: %w", err)
	}

	esConfig := elasticsearch.Config{
		Addresses: []string{esContaianer.Settings.Address},
		Username:  "elastic",
		Password:  esContaianer.Settings.Password,
		CACert:    esContaianer.Settings.CACert,
	}
	es8Client, err := elasticsearch.NewTypedClient(esConfig)
	if err != nil {
		cleanup()
		return nil, err
	}

	return &Elastic{
		Client:    es8Client,
		Container: esContaianer,
		cleanup:   cleanup,
	}, nil

}

// CreateElasticContainer function starts and elastic container.
func CreateElasticContainer(ctx context.Context) (*esmodule.ElasticsearchContainer, func(), error) {
	esContainer, err := esmodule.RunContainer(
		ctx,
		testcontainers.WithImage("docker.elastic.co/elasticsearch/elasticsearch:8.3.0"),
		esmodule.WithPassword("testpassword"),
	)
	cleanup := func() {
		if err := esContainer.Terminate(ctx); err != nil {
			fmt.Printf("Could not terminate Elasticsearch container: %v\n", err)
		}
	}
	if err != nil {
		return nil, cleanup, fmt.Errorf("Could not start Elasticsearch container: %v", err)
	}

	return esContainer, cleanup, nil
}

// AddVSSMapping inserts a mapping into Elasticsearch using a typed client.
func AddVSSMapping(ctx context.Context, client *elasticsearch.TypedClient) error {
	_, err := client.Indices.Create(VSSIndex).WaitForActiveShards("1").Do(ctx)
	if err != nil {
		return fmt.Errorf("failed to create index: %w", err)
	}
	_, err = client.Indices.PutMapping(VSSIndex).Raw(bytes.NewReader(vssMapping)).Do(ctx)
	if err != nil {
		return fmt.Errorf("failed to insert mapping: %w", err)
	}

	return nil
}

// AddStatusMapping inserts a mapping into Elasticsearch using a typed client.
func AddStatusMapping(ctx context.Context, client *elasticsearch.TypedClient) error {
	_, err := client.Indices.Create(StatusIndex).WaitForActiveShards("1").Do(ctx)
	if err != nil {
		return fmt.Errorf("failed to create index: %w", err)
	}
	_, err = client.Indices.PutMapping(StatusIndex).Raw(bytes.NewReader(statusMapping)).Do(ctx)
	if err != nil {
		return fmt.Errorf("failed to insert mapping: %w", err)
	}

	return nil
}

// TruncateStatusIndex truncates the status index.
func TruncateStatusIndex(ctx context.Context, client *elasticsearch.TypedClient) error {
	_, err := client.Indices.Delete(StatusIndex).Do(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete index: %w", err)
	}
	return AddStatusMapping(ctx, client)
}
