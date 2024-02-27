package convert

import (
	"math"

	vsstable "github.com/KevinJoiner/vss-translator/internal/generated/vss"
	"github.com/KevinJoiner/vss-translator/internal/model"
	"github.com/KevinJoiner/vss-translator/internal/protobuf"
)

// EsStatusToProtoVehicle converts an ElaticStatus struct into a protobuf model.Vehicle struct.
func EsStatusToProtoVehicle(dataPoint *model.DataPoint) *protobuf.Vehicle {
	esStatus := dataPoint.Data
	vssStatus := &protobuf.Vehicle{
		Subject:   dataPoint.Subject,
		Timestamp: dataPoint.Time,
		DIMO: &protobuf.VehicleDIMO{
			Type:         dataPoint.Type,
			Source:       dataPoint.Source,
			VehicleID:    esStatus.VehicleID,
			DefinitionID: esStatus.DefinitionID,
		},
		VehicleIdentification: &protobuf.VehicleVehicleIdentification{
			Brand: esStatus.Make,
			Model: esStatus.Model,
			Year:  float64toUint32(esStatus.Year),
		},

		Chassis: &protobuf.VehicleChassis{
			Axle: &protobuf.VehicleChassisAxle{
				Row1: &protobuf.VehicleChassisAxleRow1{
					Wheel: &protobuf.VehicleChassisAxleRow1Wheel{
						Left: &protobuf.VehicleChassisAxleRow1WheelLeft{
							Tire: &protobuf.VehicleChassisAxleRow1WheelLeftTire{
								Pressure: float64toUint32(esStatus.Tires.FrontLeft),
							},
						},
						Right: &protobuf.VehicleChassisAxleRow1WheelRight{
							Tire: &protobuf.VehicleChassisAxleRow1WheelRightTire{
								Pressure: float64toUint32(esStatus.Tires.FrontRight),
							},
						},
					},
				},
				Row2: &protobuf.VehicleChassisAxleRow2{
					Wheel: &protobuf.VehicleChassisAxleRow2Wheel{
						Left: &protobuf.VehicleChassisAxleRow2WheelLeft{
							Tire: &protobuf.VehicleChassisAxleRow2WheelLeftTire{
								Pressure: float64toUint32(esStatus.Tires.BackLeft),
							},
						},
						Right: &protobuf.VehicleChassisAxleRow2WheelRight{
							Tire: &protobuf.VehicleChassisAxleRow2WheelRightTire{
								Pressure: float64toUint32(esStatus.Tires.BackRight),
							},
						},
					},
				},
			},
		},
		CurrentLocation: &protobuf.VehicleCurrentLocation{
			Altitude:  esStatus.Altitude,
			Latitude:  esStatus.Latitude,
			Longitude: esStatus.Longitude,
			Timestamp: esStatus.Timestamp,
		},
		Exterior: &protobuf.VehicleExterior{
			AirTemperature: float64Tofloat32(esStatus.AmbientTemp),
		},
		LowVoltageBattery: &protobuf.VehicleLowVoltageBattery{
			CurrentVoltage: float64Tofloat32(esStatus.BatteryVoltage),
		},
		OBD: &protobuf.VehicleOBD{
			BarometricPressure: float64Tofloat32(esStatus.BarometricPressure),
			EngineLoad:         float64Tofloat32(esStatus.EngineLoad),
			IntakeTemp:         float64Tofloat32(esStatus.IntakeTemp),
			RunTime:            float64Tofloat32(esStatus.RunTime),
		},
		Powertrain: &protobuf.VehiclePowertrain{
			CombustionEngine: &protobuf.VehiclePowertrainCombustionEngine{
				ECT:            float64Tofloat32(esStatus.CoolantTemp),
				EngineOilLevel: convertOilLevelToEnums(esStatus.Oil),
				Speed:          float64toUint32(esStatus.EngineSpeed),
				TPS:            float64toUint32(esStatus.ThrottlePosition),
			},
			FuelSystem: &protobuf.VehiclePowertrainFuelSystem{
				AbsoluteLevel:      float64Tofloat32(esStatus.FuelPercentRemaining),
				SupportedFuelTypes: []string{esStatus.FuelType},
			},
			Range: float64toUint32(esStatus.Range),
			TractionBattery: &protobuf.VehiclePowertrainTractionBattery{
				Charging: &protobuf.VehiclePowertrainTractionBatteryCharging{
					ChargeLimit: float64toUint32(esStatus.ChargeLimit),
					IsCharging:  esStatus.Charging,
				},
				GrossCapacity: float64toUint32(esStatus.BatteryCapacity),
				StateOfCharge: &protobuf.VehiclePowertrainTractionBatteryStateOfCharge{
					Current: float64Tofloat32(esStatus.SOC),
				},
			},
			Transmission: &protobuf.VehiclePowertrainTransmission{
				TravelledDistance: float64Tofloat32(esStatus.Odometer),
			},
		},
		Speed: float64Tofloat32(esStatus.Speed),
	}
	return vssStatus
}

// ESStatusToCHVehicle converts an ElaticStatus struct into a ch table compatible struct.
func ESStatusToCHVehicle(dataPoint *model.DataPoint) *vsstable.Vehicle {
	return &vsstable.Vehicle{
		VehicleDIMOSubject:                                   dataPoint.Subject,
		VehicleDIMOTimestamp:                                 dataPoint.Time,
		VehicleDIMOType:                                      dataPoint.Type,
		VehicleDIMOSource:                                    dataPoint.Source,
		VehicleDIMOVehicleID:                                 dataPoint.Data.VehicleID,
		VehicleDIMODefinitionID:                              dataPoint.Data.DefinitionID,
		VehicleVehicleIdentificationBrand:                    dataPoint.Data.Make,
		VehicleVehicleIdentificationModel:                    dataPoint.Data.Model,
		VehicleVehicleIdentificationYear:                     float64toUint16(dataPoint.Data.Year),
		VehicleChassisAxleRow1WheelLeftTirePressure:          float64toUint16(dataPoint.Data.Tires.FrontLeft),
		VehicleChassisAxleRow1WheelRightTirePressure:         float64toUint16(dataPoint.Data.Tires.FrontRight),
		VehicleChassisAxleRow2WheelLeftTirePressure:          float64toUint16(dataPoint.Data.Tires.BackLeft),
		VehicleChassisAxleRow2WheelRightTirePressure:         float64toUint16(dataPoint.Data.Tires.BackRight),
		VehicleCurrentLocationAltitude:                       dataPoint.Data.Altitude,
		VehicleCurrentLocationLatitude:                       dataPoint.Data.Latitude,
		VehicleCurrentLocationLongitude:                      dataPoint.Data.Longitude,
		VehicleCurrentLocationTimestamp:                      dataPoint.Data.Timestamp,
		VehicleExteriorAirTemperature:                        float64Tofloat32(dataPoint.Data.AmbientTemp),
		VehicleLowVoltageBatteryCurrentVoltage:               float64Tofloat32(dataPoint.Data.BatteryVoltage),
		VehicleOBDBarometricPressure:                         float64Tofloat32(dataPoint.Data.BarometricPressure),
		VehicleOBDEngineLoad:                                 float64Tofloat32(dataPoint.Data.EngineLoad),
		VehicleOBDIntakeTemp:                                 float64Tofloat32(dataPoint.Data.IntakeTemp),
		VehicleOBDRunTime:                                    float64Tofloat32(dataPoint.Data.RunTime),
		VehiclePowertrainCombustionEngineECT:                 float64Tofloat32(dataPoint.Data.CoolantTemp),
		VehiclePowertrainCombustionEngineEngineOilLevel:      convertOilLevelToEnums(dataPoint.Data.Oil),
		VehiclePowertrainCombustionEngineSpeed:               float64toUint16(dataPoint.Data.EngineSpeed),
		VehiclePowertrainCombustionEngineTPS:                 float64toUint8(dataPoint.Data.ThrottlePosition),
		VehiclePowertrainFuelSystemAbsoluteLevel:             float64Tofloat32(dataPoint.Data.FuelPercentRemaining),
		VehiclePowertrainFuelSystemSupportedFuelTypes:        []string{dataPoint.Data.FuelType},
		VehiclePowertrainRange:                               float64toUint32(dataPoint.Data.Range),
		VehiclePowertrainTractionBatteryChargingChargeLimit:  float64toUint8(dataPoint.Data.ChargeLimit),
		VehiclePowertrainTractionBatteryChargingIsCharging:   dataPoint.Data.Charging,
		VehiclePowertrainTractionBatteryGrossCapacity:        float64toUint16(dataPoint.Data.BatteryCapacity),
		VehiclePowertrainTractionBatteryStateOfChargeCurrent: float64Tofloat32(dataPoint.Data.SOC),
		VehiclePowertrainTransmissionTravelledDistance:       float64Tofloat32(dataPoint.Data.Odometer),
		VehicleSpeed: float64Tofloat32(dataPoint.Data.Speed),
	}
}

// float64Tofloat32 converts float64 to float32
func float64Tofloat32(val float64) float32 {
	if val > math.MaxFloat32 {
		return math.MaxFloat32
	}
	return float32(val)
}

// float64toUint32 converts float64 to uint32
func float64toUint32(val float64) uint32 {
	if val > math.MaxUint32 {
		return math.MaxUint32
	}
	if val < 0 {
		return 0
	}
	return uint32(val)
}

// float64toUint16 converts float64 to uint16
func float64toUint16(val float64) uint16 {
	if val > math.MaxUint16 {
		return math.MaxUint16
	}
	if val < 0 {
		return 0
	}
	return uint16(val)
}

// float64toUint8 converts float64 to uint8
func float64toUint8(val float64) uint8 {
	if val > math.MaxUint8 {
		return math.MaxUint8
	}
	if val < 0 {
		return 0
	}
	return uint8(val)
}

func convertOilLevelToEnums(oilLevel float64) string {
	switch {
	case oilLevel < 0.25:
		return "CRITICALLY_LOW"
	case oilLevel < 0.5:
		return "LOW"
	case oilLevel < 0.75:
		return "NORMAL"
	case oilLevel < .99:
		return "HIGH"
	default:
		return "CRITICALLY_HIGH"
	}
}
