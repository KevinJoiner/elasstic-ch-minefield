package clickhouse

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ClickHouse/ch-go"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/KevinJoiner/vss-translator/internal/generated/vsstable"
	"github.com/KevinJoiner/vss-translator/internal/model"
	"github.com/KevinJoiner/vss-translator/pkg/convert"
)

// RFC3339Layout     = "2006-01-02T15:04:05Z07:00"
const RFC3339Layout = "%Y-%m-%dT%H:%M:%SZ"

type Service struct {
	conn   driver.Conn
	client *ch.Client
}

func New(ctx context.Context, host string) (*Service, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{host},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open clickhouse connection: %w", err)
	}
	client, err := ch.Dial(ctx, ch.Options{Address: host})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to clickhouse: %w", err)
	}
	return &Service{
		conn:   conn,
		client: client,
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

func (s *Service) InsertVehiclesJSON(ctx context.Context, vehicles []*vsstable.Vehicle) error {
	// var input proto.ColBytes
	jsonData, err := json.Marshal(vehicles)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}
	var builder strings.Builder
	builder.WriteString("INSERT INTO vss FORMAT JSONEachRow ")
	builder.Write(jsonData)
	// cols := []proto.InputColumn{{Name: "", Data: &input}}
	err = s.client.Do(ctx, ch.Query{
		Body: builder.String(),
	})
	if err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
	}
	return nil
}

// InsertVehicles inserts the vehicles into the clickhouse database.
func (s *Service) InsertVehiclesCol(ctx context.Context, vehicles []*vsstable.Vehicle) error {
	batch, err := s.conn.PrepareBatch(ctx, vsstable.InsertStatment)
	if err != nil {
		return fmt.Errorf("failed to prepare batch: %w", err)
	}
	for _, vehicle := range vehicles {
		err := vehicle.AppendTo(batch)
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

// InsertVehiclesColSubAndTime inserts the vehicles into the clickhouse database with only the subject and timestamp.
func (s *Service) InsertVehiclesColSubAndTime(ctx context.Context, vehicles []*vsstable.Vehicle) error {
	statement := fmt.Sprintf("INSERT into vss (%s, %s)", vsstable.FieldVehicleDIMOSubject, vsstable.FieldVehicleDIMOTimestamp)
	batch, err := s.conn.PrepareBatch(ctx, statement)
	if err != nil {
		return fmt.Errorf("failed to prepare batch: %w", err)
	}
	for _, vehicle := range vehicles {
		err = batch.Append(vehicle.VehicleDIMOSubject, vehicle.VehicleDIMOTimestamp)
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

// InsertVehiclesProtoCol inserts the vehicles into the clickhouse database.
func (s *Service) InsertVehiclesProtoCol(ctx context.Context, vehicles []*vsstable.Vehicle) error {
	input := vsstable.GetColumns(vehicles)
	err := s.client.Do(ctx, ch.Query{
		Body:  vsstable.InsertStatmentValues,
		Input: input,
	})
	if err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
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
func (s *Service) GetVehiclesBySubject(ctx context.Context, subject string, limit int) ([]*vsstable.Vehicle, error) {
	statement := fmt.Sprintf("SELECT * FROM vss WHERE %s = ? LIMIT ?", vsstable.FieldVehicleDIMOSubject)
	rows, err := s.conn.Query(ctx, statement, subject, limit)
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

func (s *Service) GetVehiclesBySubjectOrderByTime(ctx context.Context, subject string, count int) ([]*vsstable.Vehicle, error) {
	statement := fmt.Sprintf("SELECT * FROM vss WHERE %s = ? ORDER BY %s ASC LIMIT ?", vsstable.FieldVehicleDIMOSubject, vsstable.FieldVehicleDIMOTimestamp)
	rows, err := s.conn.Query(ctx, statement, subject, count)
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

func (s *Service) GetVehiclesBySubjBrandOrderByOdm(ctx context.Context, subject, model string, count int) ([]*vsstable.Vehicle, error) {
	statement := fmt.Sprintf("SELECT * FROM vss WHERE %s = ? AND %s = ? ORDER BY %s DESC LIMIT ?", vsstable.FieldVehicleDIMOSubject, vsstable.FieldVehicleVehicleIdentificationModel, vsstable.FieldVehiclePowertrainTransmissionTravelledDistance)
	rows, err := s.conn.Query(ctx, statement, subject, model, count)
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

// GetMaxSpeedBySubjectInTimeRange returns the max speed from the clickhouse database by subject in the given time range.
func (s *Service) GetMaxSpeedBySubjectInTimeRange(ctx context.Context, subject string, start, end int64) (float64, error) {
	statement := fmt.Sprintf("SELECT max(%s) FROM vss WHERE %s = ? AND %s >= ? AND %s <= ?", vsstable.FieldVehicleSpeed, vsstable.FieldVehicleDIMOSubject, vsstable.FieldVehicleDIMOTimestamp, vsstable.FieldVehicleDIMOTimestamp)
	rows, err := s.conn.Query(ctx, statement, subject, start, end)
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()
	var maxSpeed float32
	for rows.Next() {
		err := rows.Scan(&maxSpeed)
		if err != nil {
			return 0, fmt.Errorf("failed to scan max speed: %w", err)
		}
	}
	return float64(maxSpeed), nil
}

// GetVehicleAny returns a vehicles from the clickhouse database with the given limit.
func (s *Service) GetVehicleAny(ctx context.Context, limit int) ([]*vsstable.Vehicle, error) {
	statement := fmt.Sprintf("SELECT * FROM vss LIMIT %d", limit)
	rows, err := s.conn.Query(ctx, statement)
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

// TruncateVSSTable truncates the VSS table in ClickHouse.
func (s *Service) TruncateVSSTable(ctx context.Context) error {
	err := s.conn.Exec(ctx, "TRUNCATE TABLE vss")
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	return nil
}
