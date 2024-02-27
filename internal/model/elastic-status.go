package model

import "github.com/ClickHouse/ch-go/proto"

type DataPoint struct {
	Data    ElaticStatus `json:"data"`
	Subject string       `fake:"{randomstring:[subject1,subject2,subject3,subject4]}"    json:"subject"`
	Time    string       `fake:"{timeRangeStr:2010,2020}"                                json:"time"`
	Type    string       `fake:"{randomstring:[zone.dimo.device.status.update]}"         json:"type"`
	Source  string       `fake:"{randomstring:[dimo/integration/1, dimo/integration/2]}" json:"source"`
}

// ElaticStatus represents the structure of the vehicle.
type ElaticStatus struct {
	Timestamp            string  `fake:"{timeRangeStr:2010, 2020}"                            json:"timestamp"`
	FuelType             string  `fake:"{carfueltype}"                                        json:"fuelType"`
	Make                 string  `fake:"{carmaker}"                                           json:"make"`
	Model                string  `fake:"{carmodel}"                                           json:"model"`
	VIN                  string  `fake:"-"                                                    json:"vin"`
	VehicleID            string  `fake:"randomstring:{[vehicle1,vehicle2,vehicle3,vehicle4]}" json:"vehicleId"`
	DefinitionID         string  `json:"definitionId"`
	Altitude             float64 `fake:"{float32range:-100, 10000}"                           json:"altitude"`
	Latitude             float64 `fake:"{latitude}"                                           json:"latitude"`
	Longitude            float64 `fake:"{longitude}"                                          json:"longitude"`
	AmbientTemp          float64 `fake:"{float32range:-40, 120}"                              json:"ambientTemp"`
	BatteryVoltage       float64 `fake:"{float32range:0, 15}"                                 json:"batteryVoltage"`
	BarometricPressure   float64 `fake:"{float32range:0, 30}"                                 json:"barometricPressure"`
	EngineLoad           float64 `fake:"{float32range:0, 100}"                                json:"engineLoad"`
	IntakeTemp           float64 `fake:"{float32range:-40, 120}"                              json:"intakeTemp"`
	RunTime              float64 `fake:"{float32range:0, 10000}"                              json:"runTime"`
	CoolantTemp          float64 `fake:"{float32range:-40, 250}"                              json:"coolantTemp"`
	Oil                  float64 `fake:"{float32range:0, 100}"                                json:"oil"`
	EngineSpeed          float64 `fake:"{float32range:500, 1500}"                             json:"engineSpeed"`
	ThrottlePosition     float64 `fake:"{float32range:0, 3}"                                  json:"throttlePosition"`
	FuelPercentRemaining float64 `fake:"{float32range:0, 100}"                                json:"fuelPercentRemaining"`
	Range                float64 `fake:"{float32range:0, 1000}"                               json:"range"`
	ChargeLimit          float64 `fake:"{float32range:0, 100}"                                json:"chargeLimit"`
	Charging             bool    `fake:"{bool}"                                               json:"charging"`
	BatteryCapacity      float64 `fake:"{float32range:0, 100}"                                json:"batteryCapacity"`
	SOC                  float64 `fake:"{float32range:0, 1}"                                  json:"soc"`
	Speed                float64 `fake:"{float32range:0, 150}"                                json:"speed"`
	Odometer             float64 `fake:"{float32range:0, 100000}"                             json:"odometer"`
	Year                 float64 `fake:"{intrange:2007, 2022}"                                json:"year"`
	Tires                Tires   `json:"tires"`
}
type Tires struct {
	FrontLeft  float64 `fake:"{float32range:0, 50}" json:"frontLeft"`
	FrontRight float64 `fake:"{float32range:0, 50}" json:"frontRight"`
	BackLeft   float64 `fake:"{float32range:0, 50}" json:"backLeft"`
	BackRight  float64 `fake:"{float32range:0, 50}" json:"backRight"`
}

// VehicleCol is a representation of a vehicle with its fields redefined as cols
type VehicleCol struct {
	Subject   proto.ColStr
	Timestamp proto.ColStr
}
