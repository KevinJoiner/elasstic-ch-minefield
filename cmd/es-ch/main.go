// Package For testing random code.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ClickHouse/ch-go"
	chproto "github.com/ClickHouse/ch-go/proto"
	"github.com/ClickHouse/clickhouse-go/v2"
	vsstable "github.com/KevinJoiner/vss-translator/internal/generated/vss"
	"github.com/KevinJoiner/vss-translator/internal/model"
	"github.com/KevinJoiner/vss-translator/internal/protobuf"
	"github.com/KevinJoiner/vss-translator/internal/protobuf/vss"
	"github.com/KevinJoiner/vss-translator/pkg/collector"
	"github.com/KevinJoiner/vss-translator/pkg/convert"
	"github.com/KevinJoiner/vss-translator/pkg/fake"
	"google.golang.org/protobuf/encoding/protodelim"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var (
	// numberOfVehicles is the number of vehicles
	// numberOfVehicles = 1000000
	numberOfVehicles = 1
)

// main function takes in a directory path containging JSON data of elastic status and converts it to VSS status
func main() {
	// use the flags package to parse the command line arguments
	dir := flag.String("dir", "", "directory containing JSON data of elastic status")
	flag.Parse()
	var vehicles []*model.DataPoint
	var err error
	if dir != nil && *dir != "" {
		vehicles, err = collectVehicles(*dir)
		if err != nil {
			log.Fatalf("failed to collect vehicles: %v", err)
			return
		}
	} else {
		vehicles, err = createVehicle()
		if err != nil {
			log.Fatalf("failed to create vehicles: %v", err)
			return
		}
	}
	var protoVehicles bytes.Buffer
	// jsonVehicles := make([]json.RawMessage, 0, len(vehicles))
	vssVehicles := make([]*vsstable.Vehicle, 0, len(vehicles))
	for _, vehicle := range vehicles {
		vssVehicles = append(vssVehicles, convert.ESStatusToCHVehicle(vehicle))
		// fmt.Printf("writting vehicle %d buffer\n", i)
		// vssStatus := convert.ElasticStatusToVSSStatus(vehicle)
		// err = getProtoDelim(&protoVehicles, vssStatus)
		// if err != nil {
		// 	log.Fatalf("failed to marshal vehicle: %v", err)
		// 	return
		// }
		// vssVehicles = append(vssVehicles, vssStatus)
		// d, err := getProto(vssStatus)
		// if err != nil {
		// 	log.Fatalf("failed to marshal vehicle: %v", err)
		// 	return
		// }
		// err = ch1(vssStatus)
		// if err != nil {
		// 	log.Fatalf("failed to write to clickhouse: %v", err)
		// 	return
		// }

		// jsonData, err := json.Marshal(vssStatus)
		// if err != nil {
		// 	log.Fatalf("failed to marshal vehicle: %v", err)
		// 	return
		// }
		// jsonVehicles = append(jsonVehicles, jsonData)
	}

	err = ch1(vssVehicles)
	if err != nil {
		log.Fatalf("failed to write to clickhouse: %v", err)
		return
	}
	fmt.Printf("writing protos to file\n")
	os.WriteFile("vssALL.bin", protoVehicles.Bytes(), 0644)

	fmt.Printf("writing json to file\n")
	allJSON, err := json.Marshal(vehicles)
	if err != nil {
		log.Fatalf("failed to marshal vehicles: %v", err)
		return
	}
	os.WriteFile("vssALL.json", allJSON, 0644)
}
func getProtoDelim(writer io.Writer, protoMessage proto.Message) error {
	_, err := protodelim.MarshalTo(writer, protoMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal vehicle: %w", err)
	}
	return nil
}
func getProto(protoMessage proto.Message) ([]byte, error) {
	data, err := proto.Marshal(protoMessage)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal vehicle: %w", err)
	}
	return data, nil
}

func getJSON(writer io.Writer, protoMessage any) error {
	err := json.NewEncoder(writer).Encode(protoMessage)
	if err != nil {
		return fmt.Errorf("failed to encode vehicle to JSON: %w", err)
	}
	return nil
}
func getProtoJSON(protoMessage proto.Message) ([]byte, error) {
	data, err := protojson.Marshal(protoMessage)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal vehicle: %w", err)
	}
	return data, nil
}

func ch1(vehicles []*vsstable.Vehicle) error {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"localhost:9000"},
	})
	if err != nil {
		log.Fatal(err)
	}

	batch, err := conn.PrepareBatch(context.Background(), "INSERT INTO vss")
	if err != nil {
		return fmt.Errorf("failed to prepare batch: %w", err)
	}
	for _, vehicle := range vehicles {
		err = batch.AppendStruct(vehicle)
		if err != nil {
			return fmt.Errorf("failed to append vehicles to batch: %w", err)
		}
	}

	// Serialize the Protobuf message to bytes
	err = batch.Send()
	if err != nil {
		return fmt.Errorf("failed to send batch: %w", err)
	}
	err = conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}

	return nil
}

func ch2(data []*protobuf.Vehicle) error {
	ctx := context.Background()
	conn, err := ch.Dial(ctx, ch.Options{Address: "localhost:9000"})
	if err != nil {
		return err
	}
	// Define all columns of table.
	// create vars for each model.Vehicle field that is sued above
	var (
		subject                               = vss.NewCol(vss.FieldSubject)
		timestamp                             = vss.NewCol(vss.FieldTimestamp)
		dimoType                              = vss.NewCol(vss.FieldDIMOType)
		dimoSrc                               = vss.NewCol(vss.FieldDIMOSource)
		dimoVehID                             = vss.NewCol(vss.FieldDIMOVehicleID)
		dimoDefID                             = vss.NewCol(vss.FieldDIMODefinitionID)
		vehIdBrand                            = vss.NewCol(vss.FieldVehicleIdentificationBrand)
		vehIdModel                            = vss.NewCol(vss.FieldVehicleIdentificationModel)
		vehIdYear                             = vss.NewCol(vss.FieldVehicleIdentificationYear)
		chassisAxleRow1WheelLeftTirePressure  = vss.NewCol(vss.FieldChassisAxleRow1WheelLeftTirePressure)
		chassisAxleRow1WheelRightTirePressure = vss.NewCol(vss.FieldChassisAxleRow1WheelRightTirePressure)
		chassisAxleRow2WheelLeftTirePressure  = vss.NewCol(vss.FieldChassisAxleRow2WheelLeftTirePressure)
		chassisAxleRow2WheelRightTirePressure = vss.NewCol(vss.FieldChassisAxleRow2WheelRightTirePressure)
		currLocAltitude                       = vss.NewCol(vss.FieldCurrentLocationAltitude)
		currLocLatitude                       = vss.NewCol(vss.FieldCurrentLocationLatitude)
		currLocLongitude                      = vss.NewCol(vss.FieldCurrentLocationLongitude)
		currLocTimestamp                      = vss.NewCol(vss.FieldCurrentLocationTimestamp)
		extAirTemp                            = vss.NewCol(vss.FieldExteriorAirTemperature)
		lowVoltBatCurrVolt                    = vss.NewCol(vss.FieldLowVoltageBatteryCurrentVoltage)
		obdBaroPressure                       = vss.NewCol(vss.FieldOBDBarometricPressure)
		obdEngineLoad                         = vss.NewCol(vss.FieldOBDEngineLoad)
		obdIntakeTemp                         = vss.NewCol(vss.FieldOBDIntakeTemp)
		obdRunTime                            = vss.NewCol(vss.FieldOBDRunTime)
		ptCEECT                               = vss.NewCol(vss.FieldPowertrainCombustionEngineECT)
		ptCEOilLevel                          = vss.NewCol(vss.FieldPowertrainCombustionEngineEngineOilLevel)
		ptCESpeed                             = vss.NewCol(vss.FieldPowertrainCombustionEngineSpeed)
		ptCETPS                               = vss.NewCol(vss.FieldPowertrainCombustionEngineTPS)
		ptFSALevel                            = vss.NewCol(vss.FieldPowertrainFuelSystemAbsoluteLevel)
		ptFSSFT                               = vss.NewCol(vss.FieldPowertrainFuelSystemSupportedFuelTypes)
		ptRange                               = vss.NewCol(vss.FieldPowertrainRange)
		ptTracCharLimit                       = vss.NewCol(vss.FieldPowertrainTractionBatteryChargingChargeLimit)
		ptTracIsCharging                      = vss.NewCol(vss.FieldPowertrainTractionBatteryChargingIsCharging)
		ptTracGrossCap                        = vss.NewCol(vss.FieldPowertrainTractionBatteryGrossCapacity)
		ptTracSOC                             = vss.NewCol(vss.FieldPowertrainTractionBatteryStateOfChargeCurrent)
		ptTransDist                           = vss.NewCol(vss.FieldPowertrainTransmissionTravelledDistance)
		speed                                 = vss.NewCol(vss.FieldSpeed)
	)
	// Append vehicles
	for _, vehicle := range data {
		vss.MustAppend(vehicle.GetSubject(), subject)
		vss.MustAppend(vehicle.GetTimestamp(), timestamp)
		vss.MustAppend(vehicle.GetDIMO().GetType(), dimoType)
		vss.MustAppend(vehicle.GetDIMO().GetSource(), dimoSrc)
		vss.MustAppend(vehicle.GetDIMO().GetVehicleID(), dimoVehID)
		vss.MustAppend(vehicle.GetDIMO().GetDefinitionID(), dimoDefID)
		vss.MustAppend(vehicle.GetVehicleIdentification().GetBrand(), vehIdBrand)
		vss.MustAppend(vehicle.GetVehicleIdentification().GetModel(), vehIdModel)
		vss.MustAppend(vehicle.GetVehicleIdentification().GetYear(), vehIdYear)
		vss.MustAppend(vehicle.GetChassis().GetAxle().GetRow1().GetWheel().GetLeft().GetTire().GetPressure(), chassisAxleRow1WheelLeftTirePressure)
		vss.MustAppend(vehicle.GetChassis().GetAxle().GetRow1().GetWheel().GetRight().GetTire().GetPressure(), chassisAxleRow1WheelRightTirePressure)
		vss.MustAppend(vehicle.GetChassis().GetAxle().GetRow2().GetWheel().GetLeft().GetTire().GetPressure(), chassisAxleRow2WheelLeftTirePressure)
		vss.MustAppend(vehicle.GetChassis().GetAxle().GetRow2().GetWheel().GetRight().GetTire().GetPressure(), chassisAxleRow2WheelRightTirePressure)
		vss.MustAppend(vehicle.GetCurrentLocation().GetAltitude(), currLocAltitude)
		vss.MustAppend(vehicle.GetCurrentLocation().GetLatitude(), currLocLatitude)
		vss.MustAppend(vehicle.GetCurrentLocation().GetLongitude(), currLocLongitude)
		vss.MustAppend(vehicle.GetCurrentLocation().GetTimestamp(), currLocTimestamp)
		vss.MustAppend(vehicle.GetExterior().GetAirTemperature(), extAirTemp)
		vss.MustAppend(vehicle.GetLowVoltageBattery().GetCurrentVoltage(), lowVoltBatCurrVolt)
		vss.MustAppend(vehicle.GetOBD().GetBarometricPressure(), obdBaroPressure)
		vss.MustAppend(vehicle.GetOBD().GetEngineLoad(), obdEngineLoad)
		vss.MustAppend(vehicle.GetOBD().GetIntakeTemp(), obdIntakeTemp)
		vss.MustAppend(vehicle.GetOBD().GetRunTime(), obdRunTime)
		vss.MustAppend(vehicle.GetPowertrain().GetCombustionEngine().GetECT(), ptCEECT)
		vss.MustAppend(vehicle.GetPowertrain().GetCombustionEngine().GetEngineOilLevel(), ptCEOilLevel)
		vss.MustAppend(vehicle.GetPowertrain().GetCombustionEngine().GetSpeed(), ptCESpeed)
		vss.MustAppend(vehicle.GetPowertrain().GetCombustionEngine().GetTPS(), ptCETPS)
		vss.MustAppend(vehicle.GetPowertrain().GetFuelSystem().GetAbsoluteLevel(), ptFSALevel)
		vss.MustAppend(vehicle.GetPowertrain().GetFuelSystem().GetSupportedFuelTypes(), ptFSSFT)
		vss.MustAppend(vehicle.GetPowertrain().GetRange(), ptRange)
		vss.MustAppend(vehicle.GetPowertrain().GetTractionBattery().GetCharging().GetChargeLimit(), ptTracCharLimit)
		vss.MustAppend(vehicle.GetPowertrain().GetTractionBattery().GetCharging().GetIsCharging(), ptTracIsCharging)
		vss.MustAppend(vehicle.GetPowertrain().GetTractionBattery().GetGrossCapacity(), ptTracGrossCap)
		vss.MustAppend(vehicle.GetPowertrain().GetTractionBattery().GetStateOfCharge().GetCurrent(), ptTracSOC)
		vss.MustAppend(vehicle.GetPowertrain().GetTransmission().GetTravelledDistance(), ptTransDist)
		vss.MustAppend(vehicle.GetSpeed(), speed)
	}

	// Insert single data block.
	input := []chproto.InputColumn{
		{Name: vss.FieldSubject, Data: subject},
		{Name: vss.FieldTimestamp, Data: timestamp},
		{Name: vss.FieldDIMOType, Data: dimoType},
		{Name: vss.FieldDIMOSource, Data: dimoSrc},
		{Name: vss.FieldDIMOVehicleID, Data: dimoVehID},
		{Name: vss.FieldDIMODefinitionID, Data: dimoDefID},
		{Name: vss.FieldVehicleIdentificationBrand, Data: vehIdBrand},
		{Name: vss.FieldVehicleIdentificationModel, Data: vehIdModel},
		{Name: vss.FieldVehicleIdentificationYear, Data: vehIdYear},
		{Name: vss.FieldChassisAxleRow1WheelLeftTirePressure, Data: chassisAxleRow1WheelLeftTirePressure},
		{Name: vss.FieldChassisAxleRow1WheelRightTirePressure, Data: chassisAxleRow1WheelRightTirePressure},
		{Name: vss.FieldChassisAxleRow2WheelLeftTirePressure, Data: chassisAxleRow2WheelLeftTirePressure},
		{Name: vss.FieldChassisAxleRow2WheelRightTirePressure, Data: chassisAxleRow2WheelRightTirePressure},
		{Name: vss.FieldCurrentLocationAltitude, Data: currLocAltitude},
		{Name: vss.FieldCurrentLocationLatitude, Data: currLocLatitude},
		{Name: vss.FieldCurrentLocationLongitude, Data: currLocLongitude},
		{Name: vss.FieldCurrentLocationTimestamp, Data: currLocTimestamp},
		{Name: vss.FieldExteriorAirTemperature, Data: extAirTemp},
		{Name: vss.FieldLowVoltageBatteryCurrentVoltage, Data: lowVoltBatCurrVolt},
		{Name: vss.FieldOBDBarometricPressure, Data: obdBaroPressure},
		{Name: vss.FieldOBDEngineLoad, Data: obdEngineLoad},
		{Name: vss.FieldOBDIntakeTemp, Data: obdIntakeTemp},
		{Name: vss.FieldOBDRunTime, Data: obdRunTime},
		{Name: vss.FieldPowertrainCombustionEngineECT, Data: ptCEECT},
		{Name: vss.FieldPowertrainCombustionEngineEngineOilLevel, Data: ptCEOilLevel},
		{Name: vss.FieldPowertrainCombustionEngineSpeed, Data: ptCESpeed},
		{Name: vss.FieldPowertrainCombustionEngineTPS, Data: ptCETPS},
		{Name: vss.FieldPowertrainFuelSystemAbsoluteLevel, Data: ptFSALevel},
		{Name: vss.FieldPowertrainFuelSystemSupportedFuelTypes, Data: ptFSSFT},
		{Name: vss.FieldPowertrainRange, Data: ptRange},
		{Name: vss.FieldPowertrainTractionBatteryChargingChargeLimit, Data: ptTracCharLimit},
		{Name: vss.FieldPowertrainTractionBatteryChargingIsCharging, Data: ptTracIsCharging},
		{Name: vss.FieldPowertrainTractionBatteryGrossCapacity, Data: ptTracGrossCap},
		{Name: vss.FieldPowertrainTractionBatteryStateOfChargeCurrent, Data: ptTracSOC},
		{Name: vss.FieldPowertrainTransmissionTravelledDistance, Data: ptTransDist},
		{Name: vss.FieldSpeed, Data: speed},
	}

	if err := conn.Do(ctx, ch.Query{
		Body:  "INSERT INTO vss VALUES",
		Input: input,
	}); err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
	}
	return nil
}

// collect all vehicles from the directory
func collectVehicles(dir string) ([]*model.DataPoint, error) {
	// create a new collector
	coll := collector.NewCollector[*model.DataPoint](nil, nil)
	objs, err := coll.FromDir(dir, true)
	if err != nil {
		return nil, fmt.Errorf("failed to collect vehicles from directory: %w", err)
	}
	return objs, nil
}

// creatVehicle creates a new fake elatic status
func createVehicle() ([]*model.DataPoint, error) {
	fake.Setup()
	// create a thousand fake vehicles
	vehicles := make([]*model.DataPoint, numberOfVehicles)
	for i := range numberOfVehicles {
		vehicle, err := fake.ElaticStatus()
		if err != nil {
			return nil, fmt.Errorf("failed to create fake vehicle: %w", err)
		}
		vehicles[i] = vehicle
		// print the vehicle number and the vehicle subject
		fmt.Printf("Vehicle %d: %s\n", i, vehicle.Subject)
	}
	return vehicles, nil
}
