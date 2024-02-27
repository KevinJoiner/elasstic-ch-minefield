package containers

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/ClickHouse/ch-go"
	"github.com/testcontainers/testcontainers-go"
	chmodule "github.com/testcontainers/testcontainers-go/modules/clickhouse"
)

const (
	vspecProtofileName = "vspec.proto"
	initscriptName     = "init-db.sh"
)

type ClickHouse struct {
	Container *chmodule.ClickHouseContainer
	cleanup   func()
}

func (c *ClickHouse) Cleanup() {
	c.cleanup()
}

// NewClickHouseContainer creates a new clickhouse container and a new clickhouse service with the client and the logger.
func NewClickHouseContainer(ctx context.Context) (*ClickHouse, error) {
	chContainer, cleanup, err := CreateClickHouseContainer(ctx)
	if err != nil {
		cleanup()
		return nil, fmt.Errorf("failed to create ClickHouse container: %w", err)
	}

	return &ClickHouse{
		Container: chContainer,
		cleanup:   cleanup,
	}, nil
}

// CreateClickHouseContainer function starts and clickhouse container.
func CreateClickHouseContainer(ctx context.Context) (*chmodule.ClickHouseContainer, func(), error) {
	// create tmp file for init script

	clickHouseContainer, err := chmodule.RunContainer(ctx,
		testcontainers.WithImage("clickhouse/clickhouse-server:23.3.8.21-alpine"),
		chmodule.WithDatabase(ch.DefaultDatabase),
		chmodule.WithUsername(ch.DefaultUser),
		chmodule.WithPassword(""),
	)
	if err != nil {
		return clickHouseContainer, func() {}, fmt.Errorf("failed to run clickhouse container: %w", err)
	}
	cleanup := func() {
		if err := clickHouseContainer.Terminate(ctx); err != nil {
			fmt.Printf("Could not terminate Clickhouse container: %v\n", err)
		}
	}
	return clickHouseContainer, cleanup, nil
}
