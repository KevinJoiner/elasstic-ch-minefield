// Code generated by "vss-translator" DO NOT EDIT.
package vsstable

import (
	_ "embed"
	"time"

	"github.com/ClickHouse/ch-go/proto"
)

//go:embed vss-table.sql
var VSSTableCreateQuery string

const (
	// FieldVehicleChassisAxleRow1WheelLeftTirePressure Tire pressure in kilo-Pascal.
	FieldVehicleChassisAxleRow1WheelLeftTirePressure = "Vehicle_Chassis_Axle_Row1_Wheel_Left_Tire_Pressure"
	// FieldVehicleChassisAxleRow1WheelRightTirePressure Tire pressure in kilo-Pascal.
	FieldVehicleChassisAxleRow1WheelRightTirePressure = "Vehicle_Chassis_Axle_Row1_Wheel_Right_Tire_Pressure"
	// FieldVehicleChassisAxleRow2WheelLeftTirePressure Tire pressure in kilo-Pascal.
	FieldVehicleChassisAxleRow2WheelLeftTirePressure = "Vehicle_Chassis_Axle_Row2_Wheel_Left_Tire_Pressure"
	// FieldVehicleChassisAxleRow2WheelRightTirePressure Tire pressure in kilo-Pascal.
	FieldVehicleChassisAxleRow2WheelRightTirePressure = "Vehicle_Chassis_Axle_Row2_Wheel_Right_Tire_Pressure"
	// FieldVehicleCurrentLocationAltitude Current altitude relative to WGS 84 reference ellipsoid, as measured at the position of GNSS receiver antenna.
	FieldVehicleCurrentLocationAltitude = "Vehicle_CurrentLocation_Altitude"
	// FieldVehicleCurrentLocationLatitude Current latitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
	FieldVehicleCurrentLocationLatitude = "Vehicle_CurrentLocation_Latitude"
	// FieldVehicleCurrentLocationLongitude Current longitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
	FieldVehicleCurrentLocationLongitude = "Vehicle_CurrentLocation_Longitude"
	// FieldVehicleCurrentLocationTimestamp Timestamp from GNSS system for current location, formatted according to ISO 8601 with UTC time zone.
	FieldVehicleCurrentLocationTimestamp = "Vehicle_CurrentLocation_Timestamp"
	// FieldVehicleDIMODefinitionID ID for the vehicles definition
	FieldVehicleDIMODefinitionID = "Vehicle_DIMO_DefinitionID"
	// FieldVehicleDIMOSource where the data was sourced from
	FieldVehicleDIMOSource = "Vehicle_DIMO_Source"
	// FieldVehicleDIMOSubject subjet of this vehicle data
	FieldVehicleDIMOSubject = "Vehicle_DIMO_Subject"
	// FieldVehicleDIMOTimestamp timestamp of when this data was colllected
	FieldVehicleDIMOTimestamp = "Vehicle_DIMO_Timestamp"
	// FieldVehicleDIMOType type of data collected
	FieldVehicleDIMOType = "Vehicle_DIMO_Type"
	// FieldVehicleDIMOVehicleID unque DIMO ID for the vehicle
	FieldVehicleDIMOVehicleID = "Vehicle_DIMO_VehicleID"
	// FieldVehicleExteriorAirTemperature Air temperature outside the vehicle.
	FieldVehicleExteriorAirTemperature = "Vehicle_Exterior_AirTemperature"
	// FieldVehicleLowVoltageBatteryCurrentVoltage Current Voltage of the low voltage battery.
	FieldVehicleLowVoltageBatteryCurrentVoltage = "Vehicle_LowVoltageBattery_CurrentVoltage"
	// FieldVehicleOBDBarometricPressure PID 33 - Barometric pressure
	FieldVehicleOBDBarometricPressure = "Vehicle_OBD_BarometricPressure"
	// FieldVehicleOBDEngineLoad PID 04 - Engine load in percent - 0 = no load, 100 = full load
	FieldVehicleOBDEngineLoad = "Vehicle_OBD_EngineLoad"
	// FieldVehicleOBDIntakeTemp PID 0F - Intake temperature
	FieldVehicleOBDIntakeTemp = "Vehicle_OBD_IntakeTemp"
	// FieldVehicleOBDRunTime PID 1F - Engine run time
	FieldVehicleOBDRunTime = "Vehicle_OBD_RunTime"
	// FieldVehiclePowertrainCombustionEngineECT Engine coolant temperature.
	FieldVehiclePowertrainCombustionEngineECT = "Vehicle_Powertrain_CombustionEngine_ECT"
	// FieldVehiclePowertrainCombustionEngineEngineOilLevel Engine oil level.
	FieldVehiclePowertrainCombustionEngineEngineOilLevel = "Vehicle_Powertrain_CombustionEngine_EngineOilLevel"
	// FieldVehiclePowertrainCombustionEngineSpeed Engine speed measured as rotations per minute.
	FieldVehiclePowertrainCombustionEngineSpeed = "Vehicle_Powertrain_CombustionEngine_Speed"
	// FieldVehiclePowertrainCombustionEngineTPS Current throttle position.
	FieldVehiclePowertrainCombustionEngineTPS = "Vehicle_Powertrain_CombustionEngine_TPS"
	// FieldVehiclePowertrainFuelSystemAbsoluteLevel Current available fuel in the fuel tank expressed in liters.
	FieldVehiclePowertrainFuelSystemAbsoluteLevel = "Vehicle_Powertrain_FuelSystem_AbsoluteLevel"
	// FieldVehiclePowertrainFuelSystemSupportedFuelTypes High level information of fuel types supported
	FieldVehiclePowertrainFuelSystemSupportedFuelTypes = "Vehicle_Powertrain_FuelSystem_SupportedFuelTypes"
	// FieldVehiclePowertrainRange Remaining range in meters using all energy sources available in the vehicle.
	FieldVehiclePowertrainRange = "Vehicle_Powertrain_Range"
	// FieldVehiclePowertrainTractionBatteryChargingChargeLimit Target charge limit (state of charge) for battery.
	FieldVehiclePowertrainTractionBatteryChargingChargeLimit = "Vehicle_Powertrain_TractionBattery_Charging_ChargeLimit"
	// FieldVehiclePowertrainTractionBatteryChargingIsCharging True if charging is ongoing. Charging is considered to be ongoing if energy is flowing from charger to vehicle.
	FieldVehiclePowertrainTractionBatteryChargingIsCharging = "Vehicle_Powertrain_TractionBattery_Charging_IsCharging"
	// FieldVehiclePowertrainTractionBatteryGrossCapacity Gross capacity of the battery.
	FieldVehiclePowertrainTractionBatteryGrossCapacity = "Vehicle_Powertrain_TractionBattery_GrossCapacity"
	// FieldVehiclePowertrainTractionBatteryStateOfChargeCurrent Physical state of charge of the high voltage battery, relative to net capacity. This is not necessarily the state of charge being displayed to the customer.
	FieldVehiclePowertrainTractionBatteryStateOfChargeCurrent = "Vehicle_Powertrain_TractionBattery_StateOfCharge_Current"
	// FieldVehiclePowertrainTransmissionTravelledDistance Odometer reading, total distance travelled during the lifetime of the transmission.
	FieldVehiclePowertrainTransmissionTravelledDistance = "Vehicle_Powertrain_Transmission_TravelledDistance"
	// FieldVehicleSpeed Vehicle speed.
	FieldVehicleSpeed = "Vehicle_Speed"
	// FieldVehicleVehicleIdentificationBrand Vehicle brand or manufacturer.
	FieldVehicleVehicleIdentificationBrand = "Vehicle_VehicleIdentification_Brand"
	// FieldVehicleVehicleIdentificationModel Vehicle model.
	FieldVehicleVehicleIdentificationModel = "Vehicle_VehicleIdentification_Model"
	// FieldVehicleVehicleIdentificationYear Model year of the vehicle.
	FieldVehicleVehicleIdentificationYear = "Vehicle_VehicleIdentification_Year"
)

type Vehicle struct {
	// VehicleChassisAxleRow1WheelLeftTirePressure Tire pressure in kilo-Pascal.
	VehicleChassisAxleRow1WheelLeftTirePressure uint16 `ch:"Vehicle_Chassis_Axle_Row1_Wheel_Left_Tire_Pressure" json:"Vehicle_Chassis_Axle_Row1_Wheel_Left_Tire_Pressure,omitempty"`
	// VehicleChassisAxleRow1WheelRightTirePressure Tire pressure in kilo-Pascal.
	VehicleChassisAxleRow1WheelRightTirePressure uint16 `ch:"Vehicle_Chassis_Axle_Row1_Wheel_Right_Tire_Pressure" json:"Vehicle_Chassis_Axle_Row1_Wheel_Right_Tire_Pressure,omitempty"`
	// VehicleChassisAxleRow2WheelLeftTirePressure Tire pressure in kilo-Pascal.
	VehicleChassisAxleRow2WheelLeftTirePressure uint16 `ch:"Vehicle_Chassis_Axle_Row2_Wheel_Left_Tire_Pressure" json:"Vehicle_Chassis_Axle_Row2_Wheel_Left_Tire_Pressure,omitempty"`
	// VehicleChassisAxleRow2WheelRightTirePressure Tire pressure in kilo-Pascal.
	VehicleChassisAxleRow2WheelRightTirePressure uint16 `ch:"Vehicle_Chassis_Axle_Row2_Wheel_Right_Tire_Pressure" json:"Vehicle_Chassis_Axle_Row2_Wheel_Right_Tire_Pressure,omitempty"`
	// VehicleCurrentLocationAltitude Current altitude relative to WGS 84 reference ellipsoid, as measured at the position of GNSS receiver antenna.
	VehicleCurrentLocationAltitude float64 `ch:"Vehicle_CurrentLocation_Altitude" json:"Vehicle_CurrentLocation_Altitude,omitempty"`
	// VehicleCurrentLocationLatitude Current latitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
	VehicleCurrentLocationLatitude float64 `ch:"Vehicle_CurrentLocation_Latitude" json:"Vehicle_CurrentLocation_Latitude,omitempty"`
	// VehicleCurrentLocationLongitude Current longitude of vehicle in WGS 84 geodetic coordinates, as measured at the position of GNSS receiver antenna.
	VehicleCurrentLocationLongitude float64 `ch:"Vehicle_CurrentLocation_Longitude" json:"Vehicle_CurrentLocation_Longitude,omitempty"`
	// VehicleCurrentLocationTimestamp Timestamp from GNSS system for current location, formatted according to ISO 8601 with UTC time zone.
	VehicleCurrentLocationTimestamp string `ch:"Vehicle_CurrentLocation_Timestamp" json:"Vehicle_CurrentLocation_Timestamp,omitempty"`
	// VehicleDIMODefinitionID ID for the vehicles definition
	VehicleDIMODefinitionID string `ch:"Vehicle_DIMO_DefinitionID" json:"Vehicle_DIMO_DefinitionID,omitempty"`
	// VehicleDIMOSource where the data was sourced from
	VehicleDIMOSource string `ch:"Vehicle_DIMO_Source" json:"Vehicle_DIMO_Source,omitempty"`
	// VehicleDIMOSubject subjet of this vehicle data
	VehicleDIMOSubject string `ch:"Vehicle_DIMO_Subject" json:"Vehicle_DIMO_Subject,omitempty"`
	// VehicleDIMOTimestamp timestamp of when this data was colllected
	VehicleDIMOTimestamp time.Time `ch:"Vehicle_DIMO_Timestamp" json:"Vehicle_DIMO_Timestamp,omitempty"`
	// VehicleDIMOType type of data collected
	VehicleDIMOType string `ch:"Vehicle_DIMO_Type" json:"Vehicle_DIMO_Type,omitempty"`
	// VehicleDIMOVehicleID unque DIMO ID for the vehicle
	VehicleDIMOVehicleID string `ch:"Vehicle_DIMO_VehicleID" json:"Vehicle_DIMO_VehicleID,omitempty"`
	// VehicleExteriorAirTemperature Air temperature outside the vehicle.
	VehicleExteriorAirTemperature float32 `ch:"Vehicle_Exterior_AirTemperature" json:"Vehicle_Exterior_AirTemperature,omitempty"`
	// VehicleLowVoltageBatteryCurrentVoltage Current Voltage of the low voltage battery.
	VehicleLowVoltageBatteryCurrentVoltage float32 `ch:"Vehicle_LowVoltageBattery_CurrentVoltage" json:"Vehicle_LowVoltageBattery_CurrentVoltage,omitempty"`
	// VehicleOBDBarometricPressure PID 33 - Barometric pressure
	VehicleOBDBarometricPressure float32 `ch:"Vehicle_OBD_BarometricPressure" json:"Vehicle_OBD_BarometricPressure,omitempty"`
	// VehicleOBDEngineLoad PID 04 - Engine load in percent - 0 = no load, 100 = full load
	VehicleOBDEngineLoad float32 `ch:"Vehicle_OBD_EngineLoad" json:"Vehicle_OBD_EngineLoad,omitempty"`
	// VehicleOBDIntakeTemp PID 0F - Intake temperature
	VehicleOBDIntakeTemp float32 `ch:"Vehicle_OBD_IntakeTemp" json:"Vehicle_OBD_IntakeTemp,omitempty"`
	// VehicleOBDRunTime PID 1F - Engine run time
	VehicleOBDRunTime float32 `ch:"Vehicle_OBD_RunTime" json:"Vehicle_OBD_RunTime,omitempty"`
	// VehiclePowertrainCombustionEngineECT Engine coolant temperature.
	VehiclePowertrainCombustionEngineECT float32 `ch:"Vehicle_Powertrain_CombustionEngine_ECT" json:"Vehicle_Powertrain_CombustionEngine_ECT,omitempty"`
	// VehiclePowertrainCombustionEngineEngineOilLevel Engine oil level.
	VehiclePowertrainCombustionEngineEngineOilLevel string `ch:"Vehicle_Powertrain_CombustionEngine_EngineOilLevel" json:"Vehicle_Powertrain_CombustionEngine_EngineOilLevel,omitempty"`
	// VehiclePowertrainCombustionEngineSpeed Engine speed measured as rotations per minute.
	VehiclePowertrainCombustionEngineSpeed uint16 `ch:"Vehicle_Powertrain_CombustionEngine_Speed" json:"Vehicle_Powertrain_CombustionEngine_Speed,omitempty"`
	// VehiclePowertrainCombustionEngineTPS Current throttle position.
	VehiclePowertrainCombustionEngineTPS uint8 `ch:"Vehicle_Powertrain_CombustionEngine_TPS" json:"Vehicle_Powertrain_CombustionEngine_TPS,omitempty"`
	// VehiclePowertrainFuelSystemAbsoluteLevel Current available fuel in the fuel tank expressed in liters.
	VehiclePowertrainFuelSystemAbsoluteLevel float32 `ch:"Vehicle_Powertrain_FuelSystem_AbsoluteLevel" json:"Vehicle_Powertrain_FuelSystem_AbsoluteLevel,omitempty"`
	// VehiclePowertrainFuelSystemSupportedFuelTypes High level information of fuel types supported
	VehiclePowertrainFuelSystemSupportedFuelTypes []string `ch:"Vehicle_Powertrain_FuelSystem_SupportedFuelTypes" json:"Vehicle_Powertrain_FuelSystem_SupportedFuelTypes,omitempty"`
	// VehiclePowertrainRange Remaining range in meters using all energy sources available in the vehicle.
	VehiclePowertrainRange uint32 `ch:"Vehicle_Powertrain_Range" json:"Vehicle_Powertrain_Range,omitempty"`
	// VehiclePowertrainTractionBatteryChargingChargeLimit Target charge limit (state of charge) for battery.
	VehiclePowertrainTractionBatteryChargingChargeLimit uint8 `ch:"Vehicle_Powertrain_TractionBattery_Charging_ChargeLimit" json:"Vehicle_Powertrain_TractionBattery_Charging_ChargeLimit,omitempty"`
	// VehiclePowertrainTractionBatteryChargingIsCharging True if charging is ongoing. Charging is considered to be ongoing if energy is flowing from charger to vehicle.
	VehiclePowertrainTractionBatteryChargingIsCharging bool `ch:"Vehicle_Powertrain_TractionBattery_Charging_IsCharging" json:"Vehicle_Powertrain_TractionBattery_Charging_IsCharging,omitempty"`
	// VehiclePowertrainTractionBatteryGrossCapacity Gross capacity of the battery.
	VehiclePowertrainTractionBatteryGrossCapacity uint16 `ch:"Vehicle_Powertrain_TractionBattery_GrossCapacity" json:"Vehicle_Powertrain_TractionBattery_GrossCapacity,omitempty"`
	// VehiclePowertrainTractionBatteryStateOfChargeCurrent Physical state of charge of the high voltage battery, relative to net capacity. This is not necessarily the state of charge being displayed to the customer.
	VehiclePowertrainTractionBatteryStateOfChargeCurrent float32 `ch:"Vehicle_Powertrain_TractionBattery_StateOfCharge_Current" json:"Vehicle_Powertrain_TractionBattery_StateOfCharge_Current,omitempty"`
	// VehiclePowertrainTransmissionTravelledDistance Odometer reading, total distance travelled during the lifetime of the transmission.
	VehiclePowertrainTransmissionTravelledDistance float32 `ch:"Vehicle_Powertrain_Transmission_TravelledDistance" json:"Vehicle_Powertrain_Transmission_TravelledDistance,omitempty"`
	// VehicleSpeed Vehicle speed.
	VehicleSpeed float32 `ch:"Vehicle_Speed" json:"Vehicle_Speed,omitempty"`
	// VehicleVehicleIdentificationBrand Vehicle brand or manufacturer.
	VehicleVehicleIdentificationBrand string `ch:"Vehicle_VehicleIdentification_Brand" json:"Vehicle_VehicleIdentification_Brand,omitempty"`
	// VehicleVehicleIdentificationModel Vehicle model.
	VehicleVehicleIdentificationModel string `ch:"Vehicle_VehicleIdentification_Model" json:"Vehicle_VehicleIdentification_Model,omitempty"`
	// VehicleVehicleIdentificationYear Model year of the vehicle.
	VehicleVehicleIdentificationYear uint16 `ch:"Vehicle_VehicleIdentification_Year" json:"Vehicle_VehicleIdentification_Year,omitempty"`
}

// InsertStatement for inserting a vehicle and all its fields into the table.

const InsertStatment = "INSERT INTO vss (Vehicle_Chassis_Axle_Row1_Wheel_Left_Tire_Pressure, Vehicle_Chassis_Axle_Row1_Wheel_Right_Tire_Pressure, Vehicle_Chassis_Axle_Row2_Wheel_Left_Tire_Pressure, Vehicle_Chassis_Axle_Row2_Wheel_Right_Tire_Pressure, Vehicle_CurrentLocation_Altitude, Vehicle_CurrentLocation_Latitude, Vehicle_CurrentLocation_Longitude, Vehicle_CurrentLocation_Timestamp, Vehicle_DIMO_DefinitionID, Vehicle_DIMO_Source, Vehicle_DIMO_Subject, Vehicle_DIMO_Timestamp, Vehicle_DIMO_Type, Vehicle_DIMO_VehicleID, Vehicle_Exterior_AirTemperature, Vehicle_LowVoltageBattery_CurrentVoltage, Vehicle_OBD_BarometricPressure, Vehicle_OBD_EngineLoad, Vehicle_OBD_IntakeTemp, Vehicle_OBD_RunTime, Vehicle_Powertrain_CombustionEngine_ECT, Vehicle_Powertrain_CombustionEngine_EngineOilLevel, Vehicle_Powertrain_CombustionEngine_Speed, Vehicle_Powertrain_CombustionEngine_TPS, Vehicle_Powertrain_FuelSystem_AbsoluteLevel, Vehicle_Powertrain_FuelSystem_SupportedFuelTypes, Vehicle_Powertrain_Range, Vehicle_Powertrain_TractionBattery_Charging_ChargeLimit, Vehicle_Powertrain_TractionBattery_Charging_IsCharging, Vehicle_Powertrain_TractionBattery_GrossCapacity, Vehicle_Powertrain_TractionBattery_StateOfCharge_Current, Vehicle_Powertrain_Transmission_TravelledDistance, Vehicle_Speed, Vehicle_VehicleIdentification_Brand, Vehicle_VehicleIdentification_Model, Vehicle_VehicleIdentification_Year)"
const InsertStatmentValues = InsertStatment + " Values"

type Appender interface {
	Append(vals ...any) error
}

func (v *Vehicle) AppendTo(appender Appender) error {
	return appender.Append(v.VehicleChassisAxleRow1WheelLeftTirePressure,
		v.VehicleChassisAxleRow1WheelRightTirePressure,
		v.VehicleChassisAxleRow2WheelLeftTirePressure,
		v.VehicleChassisAxleRow2WheelRightTirePressure,
		v.VehicleCurrentLocationAltitude,
		v.VehicleCurrentLocationLatitude,
		v.VehicleCurrentLocationLongitude,
		v.VehicleCurrentLocationTimestamp,
		v.VehicleDIMODefinitionID,
		v.VehicleDIMOSource,
		v.VehicleDIMOSubject,
		v.VehicleDIMOTimestamp,
		v.VehicleDIMOType,
		v.VehicleDIMOVehicleID,
		v.VehicleExteriorAirTemperature,
		v.VehicleLowVoltageBatteryCurrentVoltage,
		v.VehicleOBDBarometricPressure,
		v.VehicleOBDEngineLoad,
		v.VehicleOBDIntakeTemp,
		v.VehicleOBDRunTime,
		v.VehiclePowertrainCombustionEngineECT,
		v.VehiclePowertrainCombustionEngineEngineOilLevel,
		v.VehiclePowertrainCombustionEngineSpeed,
		v.VehiclePowertrainCombustionEngineTPS,
		v.VehiclePowertrainFuelSystemAbsoluteLevel,
		v.VehiclePowertrainFuelSystemSupportedFuelTypes,
		v.VehiclePowertrainRange,
		v.VehiclePowertrainTractionBatteryChargingChargeLimit,
		v.VehiclePowertrainTractionBatteryChargingIsCharging,
		v.VehiclePowertrainTractionBatteryGrossCapacity,
		v.VehiclePowertrainTractionBatteryStateOfChargeCurrent,
		v.VehiclePowertrainTransmissionTravelledDistance,
		v.VehicleSpeed,
		v.VehicleVehicleIdentificationBrand,
		v.VehicleVehicleIdentificationModel,
		v.VehicleVehicleIdentificationYear,
	)
}

func GetColumns(vehicles []*Vehicle) []proto.InputColumn {
	var (
		colVehicleChassisAxleRow1WheelLeftTirePressure          = new(proto.ColUInt16)
		colVehicleChassisAxleRow1WheelRightTirePressure         = new(proto.ColUInt16)
		colVehicleChassisAxleRow2WheelLeftTirePressure          = new(proto.ColUInt16)
		colVehicleChassisAxleRow2WheelRightTirePressure         = new(proto.ColUInt16)
		colVehicleCurrentLocationAltitude                       = new(proto.ColFloat64)
		colVehicleCurrentLocationLatitude                       = new(proto.ColFloat64)
		colVehicleCurrentLocationLongitude                      = new(proto.ColFloat64)
		colVehicleCurrentLocationTimestamp                      = new(proto.ColStr)
		colVehicleDIMODefinitionID                              = new(proto.ColStr)
		colVehicleDIMOSource                                    = new(proto.ColStr)
		colVehicleDIMOSubject                                   = new(proto.ColStr)
		colVehicleDIMOTimestamp                                 = new(proto.ColDateTime)
		colVehicleDIMOType                                      = new(proto.ColStr)
		colVehicleDIMOVehicleID                                 = new(proto.ColStr)
		colVehicleExteriorAirTemperature                        = new(proto.ColFloat32)
		colVehicleLowVoltageBatteryCurrentVoltage               = new(proto.ColFloat32)
		colVehicleOBDBarometricPressure                         = new(proto.ColFloat32)
		colVehicleOBDEngineLoad                                 = new(proto.ColFloat32)
		colVehicleOBDIntakeTemp                                 = new(proto.ColFloat32)
		colVehicleOBDRunTime                                    = new(proto.ColFloat32)
		colVehiclePowertrainCombustionEngineECT                 = new(proto.ColFloat32)
		colVehiclePowertrainCombustionEngineEngineOilLevel      = new(proto.ColStr)
		colVehiclePowertrainCombustionEngineSpeed               = new(proto.ColUInt16)
		colVehiclePowertrainCombustionEngineTPS                 = new(proto.ColUInt8)
		colVehiclePowertrainFuelSystemAbsoluteLevel             = new(proto.ColFloat32)
		colVehiclePowertrainFuelSystemSupportedFuelTypes        = new(proto.ColStr).Array()
		colVehiclePowertrainRange                               = new(proto.ColUInt32)
		colVehiclePowertrainTractionBatteryChargingChargeLimit  = new(proto.ColUInt8)
		colVehiclePowertrainTractionBatteryChargingIsCharging   = new(proto.ColBool)
		colVehiclePowertrainTractionBatteryGrossCapacity        = new(proto.ColUInt16)
		colVehiclePowertrainTractionBatteryStateOfChargeCurrent = new(proto.ColFloat32)
		colVehiclePowertrainTransmissionTravelledDistance       = new(proto.ColFloat32)
		colVehicleSpeed                                         = new(proto.ColFloat32)
		colVehicleVehicleIdentificationBrand                    = new(proto.ColStr)
		colVehicleVehicleIdentificationModel                    = new(proto.ColStr)
		colVehicleVehicleIdentificationYear                     = new(proto.ColUInt16)
	)

	for _, vehicle := range vehicles {
		colVehicleChassisAxleRow1WheelLeftTirePressure.Append(vehicle.VehicleChassisAxleRow1WheelLeftTirePressure)
		colVehicleChassisAxleRow1WheelRightTirePressure.Append(vehicle.VehicleChassisAxleRow1WheelRightTirePressure)
		colVehicleChassisAxleRow2WheelLeftTirePressure.Append(vehicle.VehicleChassisAxleRow2WheelLeftTirePressure)
		colVehicleChassisAxleRow2WheelRightTirePressure.Append(vehicle.VehicleChassisAxleRow2WheelRightTirePressure)
		colVehicleCurrentLocationAltitude.Append(vehicle.VehicleCurrentLocationAltitude)
		colVehicleCurrentLocationLatitude.Append(vehicle.VehicleCurrentLocationLatitude)
		colVehicleCurrentLocationLongitude.Append(vehicle.VehicleCurrentLocationLongitude)
		colVehicleCurrentLocationTimestamp.Append(vehicle.VehicleCurrentLocationTimestamp)
		colVehicleDIMODefinitionID.Append(vehicle.VehicleDIMODefinitionID)
		colVehicleDIMOSource.Append(vehicle.VehicleDIMOSource)
		colVehicleDIMOSubject.Append(vehicle.VehicleDIMOSubject)
		colVehicleDIMOTimestamp.Append(vehicle.VehicleDIMOTimestamp)
		colVehicleDIMOType.Append(vehicle.VehicleDIMOType)
		colVehicleDIMOVehicleID.Append(vehicle.VehicleDIMOVehicleID)
		colVehicleExteriorAirTemperature.Append(vehicle.VehicleExteriorAirTemperature)
		colVehicleLowVoltageBatteryCurrentVoltage.Append(vehicle.VehicleLowVoltageBatteryCurrentVoltage)
		colVehicleOBDBarometricPressure.Append(vehicle.VehicleOBDBarometricPressure)
		colVehicleOBDEngineLoad.Append(vehicle.VehicleOBDEngineLoad)
		colVehicleOBDIntakeTemp.Append(vehicle.VehicleOBDIntakeTemp)
		colVehicleOBDRunTime.Append(vehicle.VehicleOBDRunTime)
		colVehiclePowertrainCombustionEngineECT.Append(vehicle.VehiclePowertrainCombustionEngineECT)
		colVehiclePowertrainCombustionEngineEngineOilLevel.Append(vehicle.VehiclePowertrainCombustionEngineEngineOilLevel)
		colVehiclePowertrainCombustionEngineSpeed.Append(vehicle.VehiclePowertrainCombustionEngineSpeed)
		colVehiclePowertrainCombustionEngineTPS.Append(vehicle.VehiclePowertrainCombustionEngineTPS)
		colVehiclePowertrainFuelSystemAbsoluteLevel.Append(vehicle.VehiclePowertrainFuelSystemAbsoluteLevel)
		colVehiclePowertrainFuelSystemSupportedFuelTypes.Append(vehicle.VehiclePowertrainFuelSystemSupportedFuelTypes)
		colVehiclePowertrainRange.Append(vehicle.VehiclePowertrainRange)
		colVehiclePowertrainTractionBatteryChargingChargeLimit.Append(vehicle.VehiclePowertrainTractionBatteryChargingChargeLimit)
		colVehiclePowertrainTractionBatteryChargingIsCharging.Append(vehicle.VehiclePowertrainTractionBatteryChargingIsCharging)
		colVehiclePowertrainTractionBatteryGrossCapacity.Append(vehicle.VehiclePowertrainTractionBatteryGrossCapacity)
		colVehiclePowertrainTractionBatteryStateOfChargeCurrent.Append(vehicle.VehiclePowertrainTractionBatteryStateOfChargeCurrent)
		colVehiclePowertrainTransmissionTravelledDistance.Append(vehicle.VehiclePowertrainTransmissionTravelledDistance)
		colVehicleSpeed.Append(vehicle.VehicleSpeed)
		colVehicleVehicleIdentificationBrand.Append(vehicle.VehicleVehicleIdentificationBrand)
		colVehicleVehicleIdentificationModel.Append(vehicle.VehicleVehicleIdentificationModel)
		colVehicleVehicleIdentificationYear.Append(vehicle.VehicleVehicleIdentificationYear)
	}

	input := []proto.InputColumn{{Name: FieldVehicleChassisAxleRow1WheelLeftTirePressure, Data: colVehicleChassisAxleRow1WheelLeftTirePressure},
		{Name: FieldVehicleChassisAxleRow1WheelRightTirePressure, Data: colVehicleChassisAxleRow1WheelRightTirePressure},
		{Name: FieldVehicleChassisAxleRow2WheelLeftTirePressure, Data: colVehicleChassisAxleRow2WheelLeftTirePressure},
		{Name: FieldVehicleChassisAxleRow2WheelRightTirePressure, Data: colVehicleChassisAxleRow2WheelRightTirePressure},
		{Name: FieldVehicleCurrentLocationAltitude, Data: colVehicleCurrentLocationAltitude},
		{Name: FieldVehicleCurrentLocationLatitude, Data: colVehicleCurrentLocationLatitude},
		{Name: FieldVehicleCurrentLocationLongitude, Data: colVehicleCurrentLocationLongitude},
		{Name: FieldVehicleCurrentLocationTimestamp, Data: colVehicleCurrentLocationTimestamp},
		{Name: FieldVehicleDIMODefinitionID, Data: colVehicleDIMODefinitionID},
		{Name: FieldVehicleDIMOSource, Data: colVehicleDIMOSource},
		{Name: FieldVehicleDIMOSubject, Data: colVehicleDIMOSubject},
		{Name: FieldVehicleDIMOTimestamp, Data: colVehicleDIMOTimestamp},
		{Name: FieldVehicleDIMOType, Data: colVehicleDIMOType},
		{Name: FieldVehicleDIMOVehicleID, Data: colVehicleDIMOVehicleID},
		{Name: FieldVehicleExteriorAirTemperature, Data: colVehicleExteriorAirTemperature},
		{Name: FieldVehicleLowVoltageBatteryCurrentVoltage, Data: colVehicleLowVoltageBatteryCurrentVoltage},
		{Name: FieldVehicleOBDBarometricPressure, Data: colVehicleOBDBarometricPressure},
		{Name: FieldVehicleOBDEngineLoad, Data: colVehicleOBDEngineLoad},
		{Name: FieldVehicleOBDIntakeTemp, Data: colVehicleOBDIntakeTemp},
		{Name: FieldVehicleOBDRunTime, Data: colVehicleOBDRunTime},
		{Name: FieldVehiclePowertrainCombustionEngineECT, Data: colVehiclePowertrainCombustionEngineECT},
		{Name: FieldVehiclePowertrainCombustionEngineEngineOilLevel, Data: colVehiclePowertrainCombustionEngineEngineOilLevel},
		{Name: FieldVehiclePowertrainCombustionEngineSpeed, Data: colVehiclePowertrainCombustionEngineSpeed},
		{Name: FieldVehiclePowertrainCombustionEngineTPS, Data: colVehiclePowertrainCombustionEngineTPS},
		{Name: FieldVehiclePowertrainFuelSystemAbsoluteLevel, Data: colVehiclePowertrainFuelSystemAbsoluteLevel},
		{Name: FieldVehiclePowertrainFuelSystemSupportedFuelTypes, Data: colVehiclePowertrainFuelSystemSupportedFuelTypes},
		{Name: FieldVehiclePowertrainRange, Data: colVehiclePowertrainRange},
		{Name: FieldVehiclePowertrainTractionBatteryChargingChargeLimit, Data: colVehiclePowertrainTractionBatteryChargingChargeLimit},
		{Name: FieldVehiclePowertrainTractionBatteryChargingIsCharging, Data: colVehiclePowertrainTractionBatteryChargingIsCharging},
		{Name: FieldVehiclePowertrainTractionBatteryGrossCapacity, Data: colVehiclePowertrainTractionBatteryGrossCapacity},
		{Name: FieldVehiclePowertrainTractionBatteryStateOfChargeCurrent, Data: colVehiclePowertrainTractionBatteryStateOfChargeCurrent},
		{Name: FieldVehiclePowertrainTransmissionTravelledDistance, Data: colVehiclePowertrainTransmissionTravelledDistance},
		{Name: FieldVehicleSpeed, Data: colVehicleSpeed},
		{Name: FieldVehicleVehicleIdentificationBrand, Data: colVehicleVehicleIdentificationBrand},
		{Name: FieldVehicleVehicleIdentificationModel, Data: colVehicleVehicleIdentificationModel},
		{Name: FieldVehicleVehicleIdentificationYear, Data: colVehicleVehicleIdentificationYear},
	}

	return input
}
