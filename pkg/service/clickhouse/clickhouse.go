package clickhouse

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	vsstable "github.com/KevinJoiner/vss-translator/internal/generated/vss"
	"github.com/KevinJoiner/vss-translator/internal/model"
	"github.com/KevinJoiner/vss-translator/pkg/convert"
)

type Service struct {
	conn driver.Conn
}

func New(host string) (*Service, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{host},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open clickhouse connection: %w", err)
	}
	return &Service{
		conn: conn,
	}, nil
}

// InsertVehiclesFromES inserts the vehicles into the clickhouse database.
func (s *Service) InsertVehiclesFromES(ctx context.Context, vehicles []*model.DataPoint) error {
	batch, err := s.conn.PrepareBatch(ctx, "INSERT into vss")
	if err != nil {
		return fmt.Errorf("failed to prepare batch: %w", err)
	}
	for _, vehicle := range vehicles {
		err := batch.AppendStruct(convert.ESStatusToCHVehicle(vehicle))
		if err != nil {
			return fmt.Errorf("failed to append vehicle to batch: %w", err)
		}
	}
	err = batch.Send()
	if err != nil {
		return fmt.Errorf("failed to send batch: %w", err)
	}
	return nil
}

// InsertVehicles inserts the vehicles into the clickhouse database.
func (s *Service) InsertVehicles(ctx context.Context, vehicles []*vsstable.Vehicle) error {
	batch, err := s.conn.PrepareBatch(ctx, "INSERT into vss")
	if err != nil {
		return fmt.Errorf("failed to prepare batch: %w", err)
	}
	for _, vehicle := range vehicles {
		err := batch.AppendStruct(vehicle)
		if err != nil {
			return fmt.Errorf("failed to append vehicle to batch: %w", err)
		}
	}
	err = batch.Send()
	if err != nil {
		return fmt.Errorf("failed to send batch: %w", err)
	}
	return nil
}

// Close closes the clickhouse connection.
func (s *Service) Close() {
	s.conn.Close()
}

// GetVehiclesAll returns the vehicles from the clickhouse database.
func (s *Service) GetVehiclesAll(ctx context.Context) ([]*vsstable.Vehicle, error) {
	rows, err := s.conn.Query(ctx, "SELECT * FROM vss")
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()
	vehicles := []*vsstable.Vehicle{}
	for rows.Next() {
		var vehicle vsstable.Vehicle
		err := rows.ScanStruct(&vehicle)
		if err != nil {
			return nil, fmt.Errorf("failed to scan vehicle: %w", err)
		}
		vehicles = append(vehicles, &vehicle)
	}
	return vehicles, nil
}

// GetVehiclesBySubject returns the vehicles from the clickhouse database by subject.
func (s *Service) GetVehiclesBySubject(ctx context.Context, subject string) ([]*vsstable.Vehicle, error) {
	statement := fmt.Sprintf("SELECT * FROM vss WHERE %s = ?", vsstable.FieldVehicleDIMOSubject)
	rows, err := s.conn.Query(ctx, statement, subject)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()
	vehicles := []*vsstable.Vehicle{}
	for rows.Next() {
		var vehicle vsstable.Vehicle
		err := rows.ScanStruct(&vehicle)
		if err != nil {
			return nil, fmt.Errorf("failed to scan vehicle: %w", err)
		}
		vehicles = append(vehicles, &vehicle)
	}
	return vehicles, nil
}

// GetVehiclesAsJSON returns the vehicles from the clickhouse database as json.
func (s *Service) GetVehiclesAsJSON(ctx context.Context) ([]byte, error) {
	rows, err := s.conn.Query(ctx, "SELECT * FROM vss FORMAT JSONEachRow")
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	var list []json.RawMessage
	for rows.Next() {
		var raw json.RawMessage
		err := rows.Scan(&raw)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		list = append(list, raw)
	}
	retData, err := json.Marshal(list)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}
	return retData, nil
}

// CreateVSSTable VSS table in ClickHouse.
func (s *Service) CreateVSSTable(ctx context.Context) error {
	err := s.conn.Exec(ctx, vsstable.VSSTableCreateQuery)
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	return nil
}
