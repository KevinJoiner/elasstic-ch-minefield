package containers

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/ClickHouse/ch-go"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	vsstable "github.com/KevinJoiner/vss-translator/internal/generated/vss"
	"github.com/testcontainers/testcontainers-go"
	chmodule "github.com/testcontainers/testcontainers-go/modules/clickhouse"
)

const (
	vspecProtofileName = "vspec.proto"
	initscriptName     = "init-db.sh"
)

type ClickHouse struct {
	Client    driver.Conn
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
	host, err := chContainer.ConnectionHost(context.TODO())
	if err != nil {
		cleanup()
		return nil, fmt.Errorf("failed to get connection host: %w", err)
	}

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{host},
	})
	if err != nil {
		cleanup()
		return nil, err
	}

	return &ClickHouse{
		Client:    conn,
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

// CreateVSSTable VSS table in ClickHouse.
func CreateVSSTable(ctx context.Context, conn driver.Conn) error {
	err := conn.Exec(ctx, vsstable.VSSTableCreateQuery)
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	return nil
}
