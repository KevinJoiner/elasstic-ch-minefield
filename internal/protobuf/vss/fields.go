// Code generated from clickhouse vss table DO NOT EDIT.
package vss

import "github.com/ClickHouse/ch-go/proto"

const (
	// FieldADAS is the name for the ADAS column in the Clickhouse.
	FieldADAS string = "ADAS"
	// FieldADASABS is the name for the ADAS.ABS column in the Clickhouse.
	FieldADASABS string = "ADAS.ABS"
	// FieldADASABSIsEnabled is the name for the ADAS.ABS.IsEnabled column in the Clickhouse.
	FieldADASABSIsEnabled string = "ADAS.ABS.IsEnabled"
	// FieldADASABSIsEngaged is the name for the ADAS.ABS.IsEngaged column in the Clickhouse.
	FieldADASABSIsEngaged string = "ADAS.ABS.IsEngaged"
	// FieldADASABSIsError is the name for the ADAS.ABS.IsError column in the Clickhouse.
	FieldADASABSIsError string = "ADAS.ABS.IsError"
	// FieldADASActiveAutonomyLevel is the name for the ADAS.ActiveAutonomyLevel column in the Clickhouse.
	FieldADASActiveAutonomyLevel string = "ADAS.ActiveAutonomyLevel"
	// FieldADASCruiseControl is the name for the ADAS.CruiseControl column in the Clickhouse.
	FieldADASCruiseControl string = "ADAS.CruiseControl"
	// FieldADASCruiseControlIsActive is the name for the ADAS.CruiseControl.IsActive column in the Clickhouse.
	FieldADASCruiseControlIsActive string = "ADAS.CruiseControl.IsActive"
	// FieldADASCruiseControlIsEnabled is the name for the ADAS.CruiseControl.IsEnabled column in the Clickhouse.
	FieldADASCruiseControlIsEnabled string = "ADAS.CruiseControl.IsEnabled"
	// FieldADASCruiseControlIsError is the name for the ADAS.CruiseControl.IsError column in the Clickhouse.
	FieldADASCruiseControlIsError string = "ADAS.CruiseControl.IsError"
	// FieldADASCruiseControlSpeedSet is the name for the ADAS.CruiseControl.SpeedSet column in the Clickhouse.
	FieldADASCruiseControlSpeedSet string = "ADAS.CruiseControl.SpeedSet"
	// FieldADASDMS is the name for the ADAS.DMS column in the Clickhouse.
	FieldADASDMS string = "ADAS.DMS"
	// FieldADASDMSIsEnabled is the name for the ADAS.DMS.IsEnabled column in the Clickhouse.
	FieldADASDMSIsEnabled string = "ADAS.DMS.IsEnabled"
	// FieldADASDMSIsError is the name for the ADAS.DMS.IsError column in the Clickhouse.
	FieldADASDMSIsError string = "ADAS.DMS.IsError"
	// FieldADASDMSIsWarning is the name for the ADAS.DMS.IsWarning column in the Clickhouse.
	FieldADASDMSIsWarning string = "ADAS.DMS.IsWarning"
	// FieldADASEBA is the name for the ADAS.EBA column in the Clickhouse.
	FieldADASEBA string = "ADAS.EBA"
	// FieldADASEBAIsEnabled is the name for the ADAS.EBA.IsEnabled column in the Clickhouse.
	FieldADASEBAIsEnabled string = "ADAS.EBA.IsEnabled"
	// FieldADASEBAIsEngaged is the name for the ADAS.EBA.IsEngaged column in the Clickhouse.
	FieldADASEBAIsEngaged string = "ADAS.EBA.IsEngaged"
	// FieldADASEBAIsError is the name for the ADAS.EBA.IsError column in the Clickhouse.
	FieldADASEBAIsError string = "ADAS.EBA.IsError"
	// FieldADASEBD is the name for the ADAS.EBD column in the Clickhouse.
	FieldADASEBD string = "ADAS.EBD"
	// FieldADASEBDIsEnabled is the name for the ADAS.EBD.IsEnabled column in the Clickhouse.
	FieldADASEBDIsEnabled string = "ADAS.EBD.IsEnabled"
	// FieldADASEBDIsEngaged is the name for the ADAS.EBD.IsEngaged column in the Clickhouse.
	FieldADASEBDIsEngaged string = "ADAS.EBD.IsEngaged"
	// FieldADASEBDIsError is the name for the ADAS.EBD.IsError column in the Clickhouse.
	FieldADASEBDIsError string = "ADAS.EBD.IsError"
	// FieldADASESC is the name for the ADAS.ESC column in the Clickhouse.
	FieldADASESC string = "ADAS.ESC"
	// FieldADASESCIsEnabled is the name for the ADAS.ESC.IsEnabled column in the Clickhouse.
	FieldADASESCIsEnabled string = "ADAS.ESC.IsEnabled"
	// FieldADASESCIsEngaged is the name for the ADAS.ESC.IsEngaged column in the Clickhouse.
	FieldADASESCIsEngaged string = "ADAS.ESC.IsEngaged"
	// FieldADASESCIsError is the name for the ADAS.ESC.IsError column in the Clickhouse.
	FieldADASESCIsError string = "ADAS.ESC.IsError"
	// FieldADASESCIsStrongCrossWindDetected is the name for the ADAS.ESC.IsStrongCrossWindDetected column in the Clickhouse.
	FieldADASESCIsStrongCrossWindDetected string = "ADAS.ESC.IsStrongCrossWindDetected"
	// FieldADASESCRoadFriction is the name for the ADAS.ESC.RoadFriction column in the Clickhouse.
	FieldADASESCRoadFriction string = "ADAS.ESC.RoadFriction"
	// FieldADASESCRoadFrictionLowerBound is the name for the ADAS.ESC.RoadFriction.LowerBound column in the Clickhouse.
	FieldADASESCRoadFrictionLowerBound string = "ADAS.ESC.RoadFriction.LowerBound"
	// FieldADASESCRoadFrictionMostProbable is the name for the ADAS.ESC.RoadFriction.MostProbable column in the Clickhouse.
	FieldADASESCRoadFrictionMostProbable string = "ADAS.ESC.RoadFriction.MostProbable"
	// FieldADASESCRoadFrictionUpperBound is the name for the ADAS.ESC.RoadFriction.UpperBound column in the Clickhouse.
	FieldADASESCRoadFrictionUpperBound string = "ADAS.ESC.RoadFriction.UpperBound"
	// FieldADASIsAutoPowerOptimize is the name for the ADAS.IsAutoPowerOptimize column in the Clickhouse.
	FieldADASIsAutoPowerOptimize string = "ADAS.IsAutoPowerOptimize"
	// FieldADASLaneDepartureDetection is the name for the ADAS.LaneDepartureDetection column in the Clickhouse.
	FieldADASLaneDepartureDetection string = "ADAS.LaneDepartureDetection"
	// FieldADASLaneDepartureDetectionIsEnabled is the name for the ADAS.LaneDepartureDetection.IsEnabled column in the Clickhouse.
	FieldADASLaneDepartureDetectionIsEnabled string = "ADAS.LaneDepartureDetection.IsEnabled"
	// FieldADASLaneDepartureDetectionIsError is the name for the ADAS.LaneDepartureDetection.IsError column in the Clickhouse.
	FieldADASLaneDepartureDetectionIsError string = "ADAS.LaneDepartureDetection.IsError"
	// FieldADASLaneDepartureDetectionIsWarning is the name for the ADAS.LaneDepartureDetection.IsWarning column in the Clickhouse.
	FieldADASLaneDepartureDetectionIsWarning string = "ADAS.LaneDepartureDetection.IsWarning"
	// FieldADASObstacleDetection is the name for the ADAS.ObstacleDetection column in the Clickhouse.
	FieldADASObstacleDetection string = "ADAS.ObstacleDetection"
	// FieldADASObstacleDetectionIsEnabled is the name for the ADAS.ObstacleDetection.IsEnabled column in the Clickhouse.
	FieldADASObstacleDetectionIsEnabled string = "ADAS.ObstacleDetection.IsEnabled"
	// FieldADASObstacleDetectionIsError is the name for the ADAS.ObstacleDetection.IsError column in the Clickhouse.
	FieldADASObstacleDetectionIsError string = "ADAS.ObstacleDetection.IsError"
	// FieldADASObstacleDetectionIsWarning is the name for the ADAS.ObstacleDetection.IsWarning column in the Clickhouse.
	FieldADASObstacleDetectionIsWarning string = "ADAS.ObstacleDetection.IsWarning"
	// FieldADASPowerOptimizeLevel is the name for the ADAS.PowerOptimizeLevel column in the Clickhouse.
	FieldADASPowerOptimizeLevel string = "ADAS.PowerOptimizeLevel"
	// FieldADASSupportedAutonomyLevel is the name for the ADAS.SupportedAutonomyLevel column in the Clickhouse.
	FieldADASSupportedAutonomyLevel string = "ADAS.SupportedAutonomyLevel"
	// FieldADASTCS is the name for the ADAS.TCS column in the Clickhouse.
	FieldADASTCS string = "ADAS.TCS"
	// FieldADASTCSIsEnabled is the name for the ADAS.TCS.IsEnabled column in the Clickhouse.
	FieldADASTCSIsEnabled string = "ADAS.TCS.IsEnabled"
	// FieldADASTCSIsEngaged is the name for the ADAS.TCS.IsEngaged column in the Clickhouse.
	FieldADASTCSIsEngaged string = "ADAS.TCS.IsEngaged"
	// FieldADASTCSIsError is the name for the ADAS.TCS.IsError column in the Clickhouse.
	FieldADASTCSIsError string = "ADAS.TCS.IsError"
	// FieldAcceleration is the name for the Acceleration column in the Clickhouse.
	FieldAcceleration string = "Acceleration"
	// FieldAccelerationLateral is the name for the Acceleration.Lateral column in the Clickhouse.
	FieldAccelerationLateral string = "Acceleration.Lateral"
	// FieldAccelerationLongitudinal is the name for the Acceleration.Longitudinal column in the Clickhouse.
	FieldAccelerationLongitudinal string = "Acceleration.Longitudinal"
	// FieldAccelerationVertical is the name for the Acceleration.Vertical column in the Clickhouse.
	FieldAccelerationVertical string = "Acceleration.Vertical"
	// FieldAngularVelocity is the name for the AngularVelocity column in the Clickhouse.
	FieldAngularVelocity string = "AngularVelocity"
	// FieldAngularVelocityPitch is the name for the AngularVelocity.Pitch column in the Clickhouse.
	FieldAngularVelocityPitch string = "AngularVelocity.Pitch"
	// FieldAngularVelocityRoll is the name for the AngularVelocity.Roll column in the Clickhouse.
	FieldAngularVelocityRoll string = "AngularVelocity.Roll"
	// FieldAngularVelocityYaw is the name for the AngularVelocity.Yaw column in the Clickhouse.
	FieldAngularVelocityYaw string = "AngularVelocity.Yaw"
	// FieldAverageSpeed is the name for the AverageSpeed column in the Clickhouse.
	FieldAverageSpeed string = "AverageSpeed"
	// FieldBody is the name for the Body column in the Clickhouse.
	FieldBody string = "Body"
	// FieldBodyBodyType is the name for the Body.BodyType column in the Clickhouse.
	FieldBodyBodyType string = "Body.BodyType"
	// FieldBodyHood is the name for the Body.Hood column in the Clickhouse.
	FieldBodyHood string = "Body.Hood"
	// FieldBodyHoodIsOpen is the name for the Body.Hood.IsOpen column in the Clickhouse.
	FieldBodyHoodIsOpen string = "Body.Hood.IsOpen"
	// FieldBodyHoodPosition is the name for the Body.Hood.Position column in the Clickhouse.
	FieldBodyHoodPosition string = "Body.Hood.Position"
	// FieldBodyHoodSwitch is the name for the Body.Hood.Switch column in the Clickhouse.
	FieldBodyHoodSwitch string = "Body.Hood.Switch"
	// FieldBodyHornIsActive is the name for the Body.Horn_IsActive column in the Clickhouse.
	FieldBodyHornIsActive string = "Body.Horn_IsActive"
	// FieldBodyIsAutoPowerOptimize is the name for the Body.IsAutoPowerOptimize column in the Clickhouse.
	FieldBodyIsAutoPowerOptimize string = "Body.IsAutoPowerOptimize"
	// FieldBodyLights is the name for the Body.Lights column in the Clickhouse.
	FieldBodyLights string = "Body.Lights"
	// FieldBodyLightsBackup is the name for the Body.Lights.Backup column in the Clickhouse.
	FieldBodyLightsBackup string = "Body.Lights.Backup"
	// FieldBodyLightsBackupIsDefect is the name for the Body.Lights.Backup.IsDefect column in the Clickhouse.
	FieldBodyLightsBackupIsDefect string = "Body.Lights.Backup.IsDefect"
	// FieldBodyLightsBackupIsOn is the name for the Body.Lights.Backup.IsOn column in the Clickhouse.
	FieldBodyLightsBackupIsOn string = "Body.Lights.Backup.IsOn"
	// FieldBodyLightsBeam is the name for the Body.Lights.Beam column in the Clickhouse.
	FieldBodyLightsBeam string = "Body.Lights.Beam"
	// FieldBodyLightsBeamHigh is the name for the Body.Lights.Beam.High column in the Clickhouse.
	FieldBodyLightsBeamHigh string = "Body.Lights.Beam.High"
	// FieldBodyLightsBeamHighIsDefect is the name for the Body.Lights.Beam.High.IsDefect column in the Clickhouse.
	FieldBodyLightsBeamHighIsDefect string = "Body.Lights.Beam.High.IsDefect"
	// FieldBodyLightsBeamHighIsOn is the name for the Body.Lights.Beam.High.IsOn column in the Clickhouse.
	FieldBodyLightsBeamHighIsOn string = "Body.Lights.Beam.High.IsOn"
	// FieldBodyLightsBeamLow is the name for the Body.Lights.Beam.Low column in the Clickhouse.
	FieldBodyLightsBeamLow string = "Body.Lights.Beam.Low"
	// FieldBodyLightsBeamLowIsDefect is the name for the Body.Lights.Beam.Low.IsDefect column in the Clickhouse.
	FieldBodyLightsBeamLowIsDefect string = "Body.Lights.Beam.Low.IsDefect"
	// FieldBodyLightsBeamLowIsOn is the name for the Body.Lights.Beam.Low.IsOn column in the Clickhouse.
	FieldBodyLightsBeamLowIsOn string = "Body.Lights.Beam.Low.IsOn"
	// FieldBodyLightsBrake is the name for the Body.Lights.Brake column in the Clickhouse.
	FieldBodyLightsBrake string = "Body.Lights.Brake"
	// FieldBodyLightsBrakeIsActive is the name for the Body.Lights.Brake.IsActive column in the Clickhouse.
	FieldBodyLightsBrakeIsActive string = "Body.Lights.Brake.IsActive"
	// FieldBodyLightsBrakeIsDefect is the name for the Body.Lights.Brake.IsDefect column in the Clickhouse.
	FieldBodyLightsBrakeIsDefect string = "Body.Lights.Brake.IsDefect"
	// FieldBodyLightsDirectionIndicator is the name for the Body.Lights.DirectionIndicator column in the Clickhouse.
	FieldBodyLightsDirectionIndicator string = "Body.Lights.DirectionIndicator"
	// FieldBodyLightsDirectionIndicatorLeft is the name for the Body.Lights.DirectionIndicator.Left column in the Clickhouse.
	FieldBodyLightsDirectionIndicatorLeft string = "Body.Lights.DirectionIndicator.Left"
	// FieldBodyLightsDirectionIndicatorLeftIsDefect is the name for the Body.Lights.DirectionIndicator.Left.IsDefect column in the Clickhouse.
	FieldBodyLightsDirectionIndicatorLeftIsDefect string = "Body.Lights.DirectionIndicator.Left.IsDefect"
	// FieldBodyLightsDirectionIndicatorLeftIsSignaling is the name for the Body.Lights.DirectionIndicator.Left.IsSignaling column in the Clickhouse.
	FieldBodyLightsDirectionIndicatorLeftIsSignaling string = "Body.Lights.DirectionIndicator.Left.IsSignaling"
	// FieldBodyLightsDirectionIndicatorRight is the name for the Body.Lights.DirectionIndicator.Right column in the Clickhouse.
	FieldBodyLightsDirectionIndicatorRight string = "Body.Lights.DirectionIndicator.Right"
	// FieldBodyLightsDirectionIndicatorRightIsDefect is the name for the Body.Lights.DirectionIndicator.Right.IsDefect column in the Clickhouse.
	FieldBodyLightsDirectionIndicatorRightIsDefect string = "Body.Lights.DirectionIndicator.Right.IsDefect"
	// FieldBodyLightsDirectionIndicatorRightIsSignaling is the name for the Body.Lights.DirectionIndicator.Right.IsSignaling column in the Clickhouse.
	FieldBodyLightsDirectionIndicatorRightIsSignaling string = "Body.Lights.DirectionIndicator.Right.IsSignaling"
	// FieldBodyLightsFog is the name for the Body.Lights.Fog column in the Clickhouse.
	FieldBodyLightsFog string = "Body.Lights.Fog"
	// FieldBodyLightsFogFront is the name for the Body.Lights.Fog.Front column in the Clickhouse.
	FieldBodyLightsFogFront string = "Body.Lights.Fog.Front"
	// FieldBodyLightsFogFrontIsDefect is the name for the Body.Lights.Fog.Front.IsDefect column in the Clickhouse.
	FieldBodyLightsFogFrontIsDefect string = "Body.Lights.Fog.Front.IsDefect"
	// FieldBodyLightsFogFrontIsOn is the name for the Body.Lights.Fog.Front.IsOn column in the Clickhouse.
	FieldBodyLightsFogFrontIsOn string = "Body.Lights.Fog.Front.IsOn"
	// FieldBodyLightsFogRear is the name for the Body.Lights.Fog.Rear column in the Clickhouse.
	FieldBodyLightsFogRear string = "Body.Lights.Fog.Rear"
	// FieldBodyLightsFogRearIsDefect is the name for the Body.Lights.Fog.Rear.IsDefect column in the Clickhouse.
	FieldBodyLightsFogRearIsDefect string = "Body.Lights.Fog.Rear.IsDefect"
	// FieldBodyLightsFogRearIsOn is the name for the Body.Lights.Fog.Rear.IsOn column in the Clickhouse.
	FieldBodyLightsFogRearIsOn string = "Body.Lights.Fog.Rear.IsOn"
	// FieldBodyLightsHazard is the name for the Body.Lights.Hazard column in the Clickhouse.
	FieldBodyLightsHazard string = "Body.Lights.Hazard"
	// FieldBodyLightsHazardIsDefect is the name for the Body.Lights.Hazard.IsDefect column in the Clickhouse.
	FieldBodyLightsHazardIsDefect string = "Body.Lights.Hazard.IsDefect"
	// FieldBodyLightsHazardIsSignaling is the name for the Body.Lights.Hazard.IsSignaling column in the Clickhouse.
	FieldBodyLightsHazardIsSignaling string = "Body.Lights.Hazard.IsSignaling"
	// FieldBodyLightsIsHighBeamSwitchOn is the name for the Body.Lights.IsHighBeamSwitchOn column in the Clickhouse.
	FieldBodyLightsIsHighBeamSwitchOn string = "Body.Lights.IsHighBeamSwitchOn"
	// FieldBodyLightsLicensePlate is the name for the Body.Lights.LicensePlate column in the Clickhouse.
	FieldBodyLightsLicensePlate string = "Body.Lights.LicensePlate"
	// FieldBodyLightsLicensePlateIsDefect is the name for the Body.Lights.LicensePlate.IsDefect column in the Clickhouse.
	FieldBodyLightsLicensePlateIsDefect string = "Body.Lights.LicensePlate.IsDefect"
	// FieldBodyLightsLicensePlateIsOn is the name for the Body.Lights.LicensePlate.IsOn column in the Clickhouse.
	FieldBodyLightsLicensePlateIsOn string = "Body.Lights.LicensePlate.IsOn"
	// FieldBodyLightsLightSwitch is the name for the Body.Lights.LightSwitch column in the Clickhouse.
	FieldBodyLightsLightSwitch string = "Body.Lights.LightSwitch"
	// FieldBodyLightsParking is the name for the Body.Lights.Parking column in the Clickhouse.
	FieldBodyLightsParking string = "Body.Lights.Parking"
	// FieldBodyLightsParkingIsDefect is the name for the Body.Lights.Parking.IsDefect column in the Clickhouse.
	FieldBodyLightsParkingIsDefect string = "Body.Lights.Parking.IsDefect"
	// FieldBodyLightsParkingIsOn is the name for the Body.Lights.Parking.IsOn column in the Clickhouse.
	FieldBodyLightsParkingIsOn string = "Body.Lights.Parking.IsOn"
	// FieldBodyLightsRunning is the name for the Body.Lights.Running column in the Clickhouse.
	FieldBodyLightsRunning string = "Body.Lights.Running"
	// FieldBodyLightsRunningIsDefect is the name for the Body.Lights.Running.IsDefect column in the Clickhouse.
	FieldBodyLightsRunningIsDefect string = "Body.Lights.Running.IsDefect"
	// FieldBodyLightsRunningIsOn is the name for the Body.Lights.Running.IsOn column in the Clickhouse.
	FieldBodyLightsRunningIsOn string = "Body.Lights.Running.IsOn"
	// FieldBodyMirrors is the name for the Body.Mirrors column in the Clickhouse.
	FieldBodyMirrors string = "Body.Mirrors"
	// FieldBodyMirrorsDriverSide is the name for the Body.Mirrors.DriverSide column in the Clickhouse.
	FieldBodyMirrorsDriverSide string = "Body.Mirrors.DriverSide"
	// FieldBodyMirrorsDriverSideIsFolded is the name for the Body.Mirrors.DriverSide.IsFolded column in the Clickhouse.
	FieldBodyMirrorsDriverSideIsFolded string = "Body.Mirrors.DriverSide.IsFolded"
	// FieldBodyMirrorsDriverSideIsHeatingOn is the name for the Body.Mirrors.DriverSide.IsHeatingOn column in the Clickhouse.
	FieldBodyMirrorsDriverSideIsHeatingOn string = "Body.Mirrors.DriverSide.IsHeatingOn"
	// FieldBodyMirrorsDriverSideIsLocked is the name for the Body.Mirrors.DriverSide.IsLocked column in the Clickhouse.
	FieldBodyMirrorsDriverSideIsLocked string = "Body.Mirrors.DriverSide.IsLocked"
	// FieldBodyMirrorsDriverSidePan is the name for the Body.Mirrors.DriverSide.Pan column in the Clickhouse.
	FieldBodyMirrorsDriverSidePan string = "Body.Mirrors.DriverSide.Pan"
	// FieldBodyMirrorsDriverSideTilt is the name for the Body.Mirrors.DriverSide.Tilt column in the Clickhouse.
	FieldBodyMirrorsDriverSideTilt string = "Body.Mirrors.DriverSide.Tilt"
	// FieldBodyMirrorsPassengerSide is the name for the Body.Mirrors.PassengerSide column in the Clickhouse.
	FieldBodyMirrorsPassengerSide string = "Body.Mirrors.PassengerSide"
	// FieldBodyMirrorsPassengerSideIsFolded is the name for the Body.Mirrors.PassengerSide.IsFolded column in the Clickhouse.
	FieldBodyMirrorsPassengerSideIsFolded string = "Body.Mirrors.PassengerSide.IsFolded"
	// FieldBodyMirrorsPassengerSideIsHeatingOn is the name for the Body.Mirrors.PassengerSide.IsHeatingOn column in the Clickhouse.
	FieldBodyMirrorsPassengerSideIsHeatingOn string = "Body.Mirrors.PassengerSide.IsHeatingOn"
	// FieldBodyMirrorsPassengerSideIsLocked is the name for the Body.Mirrors.PassengerSide.IsLocked column in the Clickhouse.
	FieldBodyMirrorsPassengerSideIsLocked string = "Body.Mirrors.PassengerSide.IsLocked"
	// FieldBodyMirrorsPassengerSidePan is the name for the Body.Mirrors.PassengerSide.Pan column in the Clickhouse.
	FieldBodyMirrorsPassengerSidePan string = "Body.Mirrors.PassengerSide.Pan"
	// FieldBodyMirrorsPassengerSideTilt is the name for the Body.Mirrors.PassengerSide.Tilt column in the Clickhouse.
	FieldBodyMirrorsPassengerSideTilt string = "Body.Mirrors.PassengerSide.Tilt"
	// FieldBodyPowerOptimizeLevel is the name for the Body.PowerOptimizeLevel column in the Clickhouse.
	FieldBodyPowerOptimizeLevel string = "Body.PowerOptimizeLevel"
	// FieldBodyRaindetectionIntensity is the name for the Body.Raindetection_Intensity column in the Clickhouse.
	FieldBodyRaindetectionIntensity string = "Body.Raindetection_Intensity"
	// FieldBodyRearMainSpoilerPosition is the name for the Body.RearMainSpoilerPosition column in the Clickhouse.
	FieldBodyRearMainSpoilerPosition string = "Body.RearMainSpoilerPosition"
	// FieldBodyRefuelPosition is the name for the Body.RefuelPosition column in the Clickhouse.
	FieldBodyRefuelPosition string = "Body.RefuelPosition"
	// FieldBodyTrunk is the name for the Body.Trunk column in the Clickhouse.
	FieldBodyTrunk string = "Body.Trunk"
	// FieldBodyTrunkFront is the name for the Body.Trunk.Front column in the Clickhouse.
	FieldBodyTrunkFront string = "Body.Trunk.Front"
	// FieldBodyTrunkFrontIsLightOn is the name for the Body.Trunk.Front.IsLightOn column in the Clickhouse.
	FieldBodyTrunkFrontIsLightOn string = "Body.Trunk.Front.IsLightOn"
	// FieldBodyTrunkFrontIsLocked is the name for the Body.Trunk.Front.IsLocked column in the Clickhouse.
	FieldBodyTrunkFrontIsLocked string = "Body.Trunk.Front.IsLocked"
	// FieldBodyTrunkFrontIsOpen is the name for the Body.Trunk.Front.IsOpen column in the Clickhouse.
	FieldBodyTrunkFrontIsOpen string = "Body.Trunk.Front.IsOpen"
	// FieldBodyTrunkFrontPosition is the name for the Body.Trunk.Front.Position column in the Clickhouse.
	FieldBodyTrunkFrontPosition string = "Body.Trunk.Front.Position"
	// FieldBodyTrunkFrontSwitch is the name for the Body.Trunk.Front.Switch column in the Clickhouse.
	FieldBodyTrunkFrontSwitch string = "Body.Trunk.Front.Switch"
	// FieldBodyTrunkRear is the name for the Body.Trunk.Rear column in the Clickhouse.
	FieldBodyTrunkRear string = "Body.Trunk.Rear"
	// FieldBodyTrunkRearIsLightOn is the name for the Body.Trunk.Rear.IsLightOn column in the Clickhouse.
	FieldBodyTrunkRearIsLightOn string = "Body.Trunk.Rear.IsLightOn"
	// FieldBodyTrunkRearIsLocked is the name for the Body.Trunk.Rear.IsLocked column in the Clickhouse.
	FieldBodyTrunkRearIsLocked string = "Body.Trunk.Rear.IsLocked"
	// FieldBodyTrunkRearIsOpen is the name for the Body.Trunk.Rear.IsOpen column in the Clickhouse.
	FieldBodyTrunkRearIsOpen string = "Body.Trunk.Rear.IsOpen"
	// FieldBodyTrunkRearPosition is the name for the Body.Trunk.Rear.Position column in the Clickhouse.
	FieldBodyTrunkRearPosition string = "Body.Trunk.Rear.Position"
	// FieldBodyTrunkRearSwitch is the name for the Body.Trunk.Rear.Switch column in the Clickhouse.
	FieldBodyTrunkRearSwitch string = "Body.Trunk.Rear.Switch"
	// FieldBodyWindshield is the name for the Body.Windshield column in the Clickhouse.
	FieldBodyWindshield string = "Body.Windshield"
	// FieldBodyWindshieldFront is the name for the Body.Windshield.Front column in the Clickhouse.
	FieldBodyWindshieldFront string = "Body.Windshield.Front"
	// FieldBodyWindshieldFrontIsHeatingOn is the name for the Body.Windshield.Front.IsHeatingOn column in the Clickhouse.
	FieldBodyWindshieldFrontIsHeatingOn string = "Body.Windshield.Front.IsHeatingOn"
	// FieldBodyWindshieldFrontWasherFluid is the name for the Body.Windshield.Front.WasherFluid column in the Clickhouse.
	FieldBodyWindshieldFrontWasherFluid string = "Body.Windshield.Front.WasherFluid"
	// FieldBodyWindshieldFrontWasherFluidIsLevelLow is the name for the Body.Windshield.Front.WasherFluid.IsLevelLow column in the Clickhouse.
	FieldBodyWindshieldFrontWasherFluidIsLevelLow string = "Body.Windshield.Front.WasherFluid.IsLevelLow"
	// FieldBodyWindshieldFrontWasherFluidLevel is the name for the Body.Windshield.Front.WasherFluid.Level column in the Clickhouse.
	FieldBodyWindshieldFrontWasherFluidLevel string = "Body.Windshield.Front.WasherFluid.Level"
	// FieldBodyWindshieldFrontWiping is the name for the Body.Windshield.Front.Wiping column in the Clickhouse.
	FieldBodyWindshieldFrontWiping string = "Body.Windshield.Front.Wiping"
	// FieldBodyWindshieldFrontWipingIntensity is the name for the Body.Windshield.Front.Wiping.Intensity column in the Clickhouse.
	FieldBodyWindshieldFrontWipingIntensity string = "Body.Windshield.Front.Wiping.Intensity"
	// FieldBodyWindshieldFrontWipingIsWipersWorn is the name for the Body.Windshield.Front.Wiping.IsWipersWorn column in the Clickhouse.
	FieldBodyWindshieldFrontWipingIsWipersWorn string = "Body.Windshield.Front.Wiping.IsWipersWorn"
	// FieldBodyWindshieldFrontWipingMode is the name for the Body.Windshield.Front.Wiping.Mode column in the Clickhouse.
	FieldBodyWindshieldFrontWipingMode string = "Body.Windshield.Front.Wiping.Mode"
	// FieldBodyWindshieldFrontWipingSystem is the name for the Body.Windshield.Front.Wiping.System column in the Clickhouse.
	FieldBodyWindshieldFrontWipingSystem string = "Body.Windshield.Front.Wiping.System"
	// FieldBodyWindshieldFrontWipingSystemActualPosition is the name for the Body.Windshield.Front.Wiping.System.ActualPosition column in the Clickhouse.
	FieldBodyWindshieldFrontWipingSystemActualPosition string = "Body.Windshield.Front.Wiping.System.ActualPosition"
	// FieldBodyWindshieldFrontWipingSystemDriveCurrent is the name for the Body.Windshield.Front.Wiping.System.DriveCurrent column in the Clickhouse.
	FieldBodyWindshieldFrontWipingSystemDriveCurrent string = "Body.Windshield.Front.Wiping.System.DriveCurrent"
	// FieldBodyWindshieldFrontWipingSystemFrequency is the name for the Body.Windshield.Front.Wiping.System.Frequency column in the Clickhouse.
	FieldBodyWindshieldFrontWipingSystemFrequency string = "Body.Windshield.Front.Wiping.System.Frequency"
	// FieldBodyWindshieldFrontWipingSystemIsBlocked is the name for the Body.Windshield.Front.Wiping.System.IsBlocked column in the Clickhouse.
	FieldBodyWindshieldFrontWipingSystemIsBlocked string = "Body.Windshield.Front.Wiping.System.IsBlocked"
	// FieldBodyWindshieldFrontWipingSystemIsEndingWipeCycle is the name for the Body.Windshield.Front.Wiping.System.IsEndingWipeCycle column in the Clickhouse.
	FieldBodyWindshieldFrontWipingSystemIsEndingWipeCycle string = "Body.Windshield.Front.Wiping.System.IsEndingWipeCycle"
	// FieldBodyWindshieldFrontWipingSystemIsOverheated is the name for the Body.Windshield.Front.Wiping.System.IsOverheated column in the Clickhouse.
	FieldBodyWindshieldFrontWipingSystemIsOverheated string = "Body.Windshield.Front.Wiping.System.IsOverheated"
	// FieldBodyWindshieldFrontWipingSystemIsPositionReached is the name for the Body.Windshield.Front.Wiping.System.IsPositionReached column in the Clickhouse.
	FieldBodyWindshieldFrontWipingSystemIsPositionReached string = "Body.Windshield.Front.Wiping.System.IsPositionReached"
	// FieldBodyWindshieldFrontWipingSystemIsWiperError is the name for the Body.Windshield.Front.Wiping.System.IsWiperError column in the Clickhouse.
	FieldBodyWindshieldFrontWipingSystemIsWiperError string = "Body.Windshield.Front.Wiping.System.IsWiperError"
	// FieldBodyWindshieldFrontWipingSystemIsWiping is the name for the Body.Windshield.Front.Wiping.System.IsWiping column in the Clickhouse.
	FieldBodyWindshieldFrontWipingSystemIsWiping string = "Body.Windshield.Front.Wiping.System.IsWiping"
	// FieldBodyWindshieldFrontWipingSystemMode is the name for the Body.Windshield.Front.Wiping.System.Mode column in the Clickhouse.
	FieldBodyWindshieldFrontWipingSystemMode string = "Body.Windshield.Front.Wiping.System.Mode"
	// FieldBodyWindshieldFrontWipingSystemTargetPosition is the name for the Body.Windshield.Front.Wiping.System.TargetPosition column in the Clickhouse.
	FieldBodyWindshieldFrontWipingSystemTargetPosition string = "Body.Windshield.Front.Wiping.System.TargetPosition"
	// FieldBodyWindshieldFrontWipingWiperWear is the name for the Body.Windshield.Front.Wiping.WiperWear column in the Clickhouse.
	FieldBodyWindshieldFrontWipingWiperWear string = "Body.Windshield.Front.Wiping.WiperWear"
	// FieldBodyWindshieldRear is the name for the Body.Windshield.Rear column in the Clickhouse.
	FieldBodyWindshieldRear string = "Body.Windshield.Rear"
	// FieldBodyWindshieldRearIsHeatingOn is the name for the Body.Windshield.Rear.IsHeatingOn column in the Clickhouse.
	FieldBodyWindshieldRearIsHeatingOn string = "Body.Windshield.Rear.IsHeatingOn"
	// FieldBodyWindshieldRearWasherFluid is the name for the Body.Windshield.Rear.WasherFluid column in the Clickhouse.
	FieldBodyWindshieldRearWasherFluid string = "Body.Windshield.Rear.WasherFluid"
	// FieldBodyWindshieldRearWasherFluidIsLevelLow is the name for the Body.Windshield.Rear.WasherFluid.IsLevelLow column in the Clickhouse.
	FieldBodyWindshieldRearWasherFluidIsLevelLow string = "Body.Windshield.Rear.WasherFluid.IsLevelLow"
	// FieldBodyWindshieldRearWasherFluidLevel is the name for the Body.Windshield.Rear.WasherFluid.Level column in the Clickhouse.
	FieldBodyWindshieldRearWasherFluidLevel string = "Body.Windshield.Rear.WasherFluid.Level"
	// FieldBodyWindshieldRearWiping is the name for the Body.Windshield.Rear.Wiping column in the Clickhouse.
	FieldBodyWindshieldRearWiping string = "Body.Windshield.Rear.Wiping"
	// FieldBodyWindshieldRearWipingIntensity is the name for the Body.Windshield.Rear.Wiping.Intensity column in the Clickhouse.
	FieldBodyWindshieldRearWipingIntensity string = "Body.Windshield.Rear.Wiping.Intensity"
	// FieldBodyWindshieldRearWipingIsWipersWorn is the name for the Body.Windshield.Rear.Wiping.IsWipersWorn column in the Clickhouse.
	FieldBodyWindshieldRearWipingIsWipersWorn string = "Body.Windshield.Rear.Wiping.IsWipersWorn"
	// FieldBodyWindshieldRearWipingMode is the name for the Body.Windshield.Rear.Wiping.Mode column in the Clickhouse.
	FieldBodyWindshieldRearWipingMode string = "Body.Windshield.Rear.Wiping.Mode"
	// FieldBodyWindshieldRearWipingSystem is the name for the Body.Windshield.Rear.Wiping.System column in the Clickhouse.
	FieldBodyWindshieldRearWipingSystem string = "Body.Windshield.Rear.Wiping.System"
	// FieldBodyWindshieldRearWipingSystemActualPosition is the name for the Body.Windshield.Rear.Wiping.System.ActualPosition column in the Clickhouse.
	FieldBodyWindshieldRearWipingSystemActualPosition string = "Body.Windshield.Rear.Wiping.System.ActualPosition"
	// FieldBodyWindshieldRearWipingSystemDriveCurrent is the name for the Body.Windshield.Rear.Wiping.System.DriveCurrent column in the Clickhouse.
	FieldBodyWindshieldRearWipingSystemDriveCurrent string = "Body.Windshield.Rear.Wiping.System.DriveCurrent"
	// FieldBodyWindshieldRearWipingSystemFrequency is the name for the Body.Windshield.Rear.Wiping.System.Frequency column in the Clickhouse.
	FieldBodyWindshieldRearWipingSystemFrequency string = "Body.Windshield.Rear.Wiping.System.Frequency"
	// FieldBodyWindshieldRearWipingSystemIsBlocked is the name for the Body.Windshield.Rear.Wiping.System.IsBlocked column in the Clickhouse.
	FieldBodyWindshieldRearWipingSystemIsBlocked string = "Body.Windshield.Rear.Wiping.System.IsBlocked"
	// FieldBodyWindshieldRearWipingSystemIsEndingWipeCycle is the name for the Body.Windshield.Rear.Wiping.System.IsEndingWipeCycle column in the Clickhouse.
	FieldBodyWindshieldRearWipingSystemIsEndingWipeCycle string = "Body.Windshield.Rear.Wiping.System.IsEndingWipeCycle"
	// FieldBodyWindshieldRearWipingSystemIsOverheated is the name for the Body.Windshield.Rear.Wiping.System.IsOverheated column in the Clickhouse.
	FieldBodyWindshieldRearWipingSystemIsOverheated string = "Body.Windshield.Rear.Wiping.System.IsOverheated"
	// FieldBodyWindshieldRearWipingSystemIsPositionReached is the name for the Body.Windshield.Rear.Wiping.System.IsPositionReached column in the Clickhouse.
	FieldBodyWindshieldRearWipingSystemIsPositionReached string = "Body.Windshield.Rear.Wiping.System.IsPositionReached"
	// FieldBodyWindshieldRearWipingSystemIsWiperError is the name for the Body.Windshield.Rear.Wiping.System.IsWiperError column in the Clickhouse.
	FieldBodyWindshieldRearWipingSystemIsWiperError string = "Body.Windshield.Rear.Wiping.System.IsWiperError"
	// FieldBodyWindshieldRearWipingSystemIsWiping is the name for the Body.Windshield.Rear.Wiping.System.IsWiping column in the Clickhouse.
	FieldBodyWindshieldRearWipingSystemIsWiping string = "Body.Windshield.Rear.Wiping.System.IsWiping"
	// FieldBodyWindshieldRearWipingSystemMode is the name for the Body.Windshield.Rear.Wiping.System.Mode column in the Clickhouse.
	FieldBodyWindshieldRearWipingSystemMode string = "Body.Windshield.Rear.Wiping.System.Mode"
	// FieldBodyWindshieldRearWipingSystemTargetPosition is the name for the Body.Windshield.Rear.Wiping.System.TargetPosition column in the Clickhouse.
	FieldBodyWindshieldRearWipingSystemTargetPosition string = "Body.Windshield.Rear.Wiping.System.TargetPosition"
	// FieldBodyWindshieldRearWipingWiperWear is the name for the Body.Windshield.Rear.Wiping.WiperWear column in the Clickhouse.
	FieldBodyWindshieldRearWipingWiperWear string = "Body.Windshield.Rear.Wiping.WiperWear"
	// FieldCabin is the name for the Cabin column in the Clickhouse.
	FieldCabin string = "Cabin"
	// FieldCabinConvertibleStatus is the name for the Cabin.Convertible_Status column in the Clickhouse.
	FieldCabinConvertibleStatus string = "Cabin.Convertible_Status"
	// FieldCabinDoor is the name for the Cabin.Door column in the Clickhouse.
	FieldCabinDoor string = "Cabin.Door"
	// FieldCabinDoorCount is the name for the Cabin.DoorCount column in the Clickhouse.
	FieldCabinDoorCount string = "Cabin.DoorCount"
	// FieldCabinDoorRow1 is the name for the Cabin.Door.Row1 column in the Clickhouse.
	FieldCabinDoorRow1 string = "Cabin.Door.Row1"
	// FieldCabinDoorRow1DriverSide is the name for the Cabin.Door.Row1.DriverSide column in the Clickhouse.
	FieldCabinDoorRow1DriverSide string = "Cabin.Door.Row1.DriverSide"
	// FieldCabinDoorRow1DriverSideIsChildLockActive is the name for the Cabin.Door.Row1.DriverSide.IsChildLockActive column in the Clickhouse.
	FieldCabinDoorRow1DriverSideIsChildLockActive string = "Cabin.Door.Row1.DriverSide.IsChildLockActive"
	// FieldCabinDoorRow1DriverSideIsLocked is the name for the Cabin.Door.Row1.DriverSide.IsLocked column in the Clickhouse.
	FieldCabinDoorRow1DriverSideIsLocked string = "Cabin.Door.Row1.DriverSide.IsLocked"
	// FieldCabinDoorRow1DriverSideIsOpen is the name for the Cabin.Door.Row1.DriverSide.IsOpen column in the Clickhouse.
	FieldCabinDoorRow1DriverSideIsOpen string = "Cabin.Door.Row1.DriverSide.IsOpen"
	// FieldCabinDoorRow1DriverSidePosition is the name for the Cabin.Door.Row1.DriverSide.Position column in the Clickhouse.
	FieldCabinDoorRow1DriverSidePosition string = "Cabin.Door.Row1.DriverSide.Position"
	// FieldCabinDoorRow1DriverSideShade is the name for the Cabin.Door.Row1.DriverSide.Shade column in the Clickhouse.
	FieldCabinDoorRow1DriverSideShade string = "Cabin.Door.Row1.DriverSide.Shade"
	// FieldCabinDoorRow1DriverSideShadeIsOpen is the name for the Cabin.Door.Row1.DriverSide.Shade.IsOpen column in the Clickhouse.
	FieldCabinDoorRow1DriverSideShadeIsOpen string = "Cabin.Door.Row1.DriverSide.Shade.IsOpen"
	// FieldCabinDoorRow1DriverSideShadePosition is the name for the Cabin.Door.Row1.DriverSide.Shade.Position column in the Clickhouse.
	FieldCabinDoorRow1DriverSideShadePosition string = "Cabin.Door.Row1.DriverSide.Shade.Position"
	// FieldCabinDoorRow1DriverSideShadeSwitch is the name for the Cabin.Door.Row1.DriverSide.Shade.Switch column in the Clickhouse.
	FieldCabinDoorRow1DriverSideShadeSwitch string = "Cabin.Door.Row1.DriverSide.Shade.Switch"
	// FieldCabinDoorRow1DriverSideSwitch is the name for the Cabin.Door.Row1.DriverSide.Switch column in the Clickhouse.
	FieldCabinDoorRow1DriverSideSwitch string = "Cabin.Door.Row1.DriverSide.Switch"
	// FieldCabinDoorRow1DriverSideWindow is the name for the Cabin.Door.Row1.DriverSide.Window column in the Clickhouse.
	FieldCabinDoorRow1DriverSideWindow string = "Cabin.Door.Row1.DriverSide.Window"
	// FieldCabinDoorRow1DriverSideWindowIsOpen is the name for the Cabin.Door.Row1.DriverSide.Window.IsOpen column in the Clickhouse.
	FieldCabinDoorRow1DriverSideWindowIsOpen string = "Cabin.Door.Row1.DriverSide.Window.IsOpen"
	// FieldCabinDoorRow1DriverSideWindowPosition is the name for the Cabin.Door.Row1.DriverSide.Window.Position column in the Clickhouse.
	FieldCabinDoorRow1DriverSideWindowPosition string = "Cabin.Door.Row1.DriverSide.Window.Position"
	// FieldCabinDoorRow1DriverSideWindowSwitch is the name for the Cabin.Door.Row1.DriverSide.Window.Switch column in the Clickhouse.
	FieldCabinDoorRow1DriverSideWindowSwitch string = "Cabin.Door.Row1.DriverSide.Window.Switch"
	// FieldCabinDoorRow1PassengerSide is the name for the Cabin.Door.Row1.PassengerSide column in the Clickhouse.
	FieldCabinDoorRow1PassengerSide string = "Cabin.Door.Row1.PassengerSide"
	// FieldCabinDoorRow1PassengerSideIsChildLockActive is the name for the Cabin.Door.Row1.PassengerSide.IsChildLockActive column in the Clickhouse.
	FieldCabinDoorRow1PassengerSideIsChildLockActive string = "Cabin.Door.Row1.PassengerSide.IsChildLockActive"
	// FieldCabinDoorRow1PassengerSideIsLocked is the name for the Cabin.Door.Row1.PassengerSide.IsLocked column in the Clickhouse.
	FieldCabinDoorRow1PassengerSideIsLocked string = "Cabin.Door.Row1.PassengerSide.IsLocked"
	// FieldCabinDoorRow1PassengerSideIsOpen is the name for the Cabin.Door.Row1.PassengerSide.IsOpen column in the Clickhouse.
	FieldCabinDoorRow1PassengerSideIsOpen string = "Cabin.Door.Row1.PassengerSide.IsOpen"
	// FieldCabinDoorRow1PassengerSidePosition is the name for the Cabin.Door.Row1.PassengerSide.Position column in the Clickhouse.
	FieldCabinDoorRow1PassengerSidePosition string = "Cabin.Door.Row1.PassengerSide.Position"
	// FieldCabinDoorRow1PassengerSideShade is the name for the Cabin.Door.Row1.PassengerSide.Shade column in the Clickhouse.
	FieldCabinDoorRow1PassengerSideShade string = "Cabin.Door.Row1.PassengerSide.Shade"
	// FieldCabinDoorRow1PassengerSideShadeIsOpen is the name for the Cabin.Door.Row1.PassengerSide.Shade.IsOpen column in the Clickhouse.
	FieldCabinDoorRow1PassengerSideShadeIsOpen string = "Cabin.Door.Row1.PassengerSide.Shade.IsOpen"
	// FieldCabinDoorRow1PassengerSideShadePosition is the name for the Cabin.Door.Row1.PassengerSide.Shade.Position column in the Clickhouse.
	FieldCabinDoorRow1PassengerSideShadePosition string = "Cabin.Door.Row1.PassengerSide.Shade.Position"
	// FieldCabinDoorRow1PassengerSideShadeSwitch is the name for the Cabin.Door.Row1.PassengerSide.Shade.Switch column in the Clickhouse.
	FieldCabinDoorRow1PassengerSideShadeSwitch string = "Cabin.Door.Row1.PassengerSide.Shade.Switch"
	// FieldCabinDoorRow1PassengerSideSwitch is the name for the Cabin.Door.Row1.PassengerSide.Switch column in the Clickhouse.
	FieldCabinDoorRow1PassengerSideSwitch string = "Cabin.Door.Row1.PassengerSide.Switch"
	// FieldCabinDoorRow1PassengerSideWindow is the name for the Cabin.Door.Row1.PassengerSide.Window column in the Clickhouse.
	FieldCabinDoorRow1PassengerSideWindow string = "Cabin.Door.Row1.PassengerSide.Window"
	// FieldCabinDoorRow1PassengerSideWindowIsOpen is the name for the Cabin.Door.Row1.PassengerSide.Window.IsOpen column in the Clickhouse.
	FieldCabinDoorRow1PassengerSideWindowIsOpen string = "Cabin.Door.Row1.PassengerSide.Window.IsOpen"
	// FieldCabinDoorRow1PassengerSideWindowPosition is the name for the Cabin.Door.Row1.PassengerSide.Window.Position column in the Clickhouse.
	FieldCabinDoorRow1PassengerSideWindowPosition string = "Cabin.Door.Row1.PassengerSide.Window.Position"
	// FieldCabinDoorRow1PassengerSideWindowSwitch is the name for the Cabin.Door.Row1.PassengerSide.Window.Switch column in the Clickhouse.
	FieldCabinDoorRow1PassengerSideWindowSwitch string = "Cabin.Door.Row1.PassengerSide.Window.Switch"
	// FieldCabinDoorRow2 is the name for the Cabin.Door.Row2 column in the Clickhouse.
	FieldCabinDoorRow2 string = "Cabin.Door.Row2"
	// FieldCabinDoorRow2DriverSide is the name for the Cabin.Door.Row2.DriverSide column in the Clickhouse.
	FieldCabinDoorRow2DriverSide string = "Cabin.Door.Row2.DriverSide"
	// FieldCabinDoorRow2DriverSideIsChildLockActive is the name for the Cabin.Door.Row2.DriverSide.IsChildLockActive column in the Clickhouse.
	FieldCabinDoorRow2DriverSideIsChildLockActive string = "Cabin.Door.Row2.DriverSide.IsChildLockActive"
	// FieldCabinDoorRow2DriverSideIsLocked is the name for the Cabin.Door.Row2.DriverSide.IsLocked column in the Clickhouse.
	FieldCabinDoorRow2DriverSideIsLocked string = "Cabin.Door.Row2.DriverSide.IsLocked"
	// FieldCabinDoorRow2DriverSideIsOpen is the name for the Cabin.Door.Row2.DriverSide.IsOpen column in the Clickhouse.
	FieldCabinDoorRow2DriverSideIsOpen string = "Cabin.Door.Row2.DriverSide.IsOpen"
	// FieldCabinDoorRow2DriverSidePosition is the name for the Cabin.Door.Row2.DriverSide.Position column in the Clickhouse.
	FieldCabinDoorRow2DriverSidePosition string = "Cabin.Door.Row2.DriverSide.Position"
	// FieldCabinDoorRow2DriverSideShade is the name for the Cabin.Door.Row2.DriverSide.Shade column in the Clickhouse.
	FieldCabinDoorRow2DriverSideShade string = "Cabin.Door.Row2.DriverSide.Shade"
	// FieldCabinDoorRow2DriverSideShadeIsOpen is the name for the Cabin.Door.Row2.DriverSide.Shade.IsOpen column in the Clickhouse.
	FieldCabinDoorRow2DriverSideShadeIsOpen string = "Cabin.Door.Row2.DriverSide.Shade.IsOpen"
	// FieldCabinDoorRow2DriverSideShadePosition is the name for the Cabin.Door.Row2.DriverSide.Shade.Position column in the Clickhouse.
	FieldCabinDoorRow2DriverSideShadePosition string = "Cabin.Door.Row2.DriverSide.Shade.Position"
	// FieldCabinDoorRow2DriverSideShadeSwitch is the name for the Cabin.Door.Row2.DriverSide.Shade.Switch column in the Clickhouse.
	FieldCabinDoorRow2DriverSideShadeSwitch string = "Cabin.Door.Row2.DriverSide.Shade.Switch"
	// FieldCabinDoorRow2DriverSideSwitch is the name for the Cabin.Door.Row2.DriverSide.Switch column in the Clickhouse.
	FieldCabinDoorRow2DriverSideSwitch string = "Cabin.Door.Row2.DriverSide.Switch"
	// FieldCabinDoorRow2DriverSideWindow is the name for the Cabin.Door.Row2.DriverSide.Window column in the Clickhouse.
	FieldCabinDoorRow2DriverSideWindow string = "Cabin.Door.Row2.DriverSide.Window"
	// FieldCabinDoorRow2DriverSideWindowIsOpen is the name for the Cabin.Door.Row2.DriverSide.Window.IsOpen column in the Clickhouse.
	FieldCabinDoorRow2DriverSideWindowIsOpen string = "Cabin.Door.Row2.DriverSide.Window.IsOpen"
	// FieldCabinDoorRow2DriverSideWindowPosition is the name for the Cabin.Door.Row2.DriverSide.Window.Position column in the Clickhouse.
	FieldCabinDoorRow2DriverSideWindowPosition string = "Cabin.Door.Row2.DriverSide.Window.Position"
	// FieldCabinDoorRow2DriverSideWindowSwitch is the name for the Cabin.Door.Row2.DriverSide.Window.Switch column in the Clickhouse.
	FieldCabinDoorRow2DriverSideWindowSwitch string = "Cabin.Door.Row2.DriverSide.Window.Switch"
	// FieldCabinDoorRow2PassengerSide is the name for the Cabin.Door.Row2.PassengerSide column in the Clickhouse.
	FieldCabinDoorRow2PassengerSide string = "Cabin.Door.Row2.PassengerSide"
	// FieldCabinDoorRow2PassengerSideIsChildLockActive is the name for the Cabin.Door.Row2.PassengerSide.IsChildLockActive column in the Clickhouse.
	FieldCabinDoorRow2PassengerSideIsChildLockActive string = "Cabin.Door.Row2.PassengerSide.IsChildLockActive"
	// FieldCabinDoorRow2PassengerSideIsLocked is the name for the Cabin.Door.Row2.PassengerSide.IsLocked column in the Clickhouse.
	FieldCabinDoorRow2PassengerSideIsLocked string = "Cabin.Door.Row2.PassengerSide.IsLocked"
	// FieldCabinDoorRow2PassengerSideIsOpen is the name for the Cabin.Door.Row2.PassengerSide.IsOpen column in the Clickhouse.
	FieldCabinDoorRow2PassengerSideIsOpen string = "Cabin.Door.Row2.PassengerSide.IsOpen"
	// FieldCabinDoorRow2PassengerSidePosition is the name for the Cabin.Door.Row2.PassengerSide.Position column in the Clickhouse.
	FieldCabinDoorRow2PassengerSidePosition string = "Cabin.Door.Row2.PassengerSide.Position"
	// FieldCabinDoorRow2PassengerSideShade is the name for the Cabin.Door.Row2.PassengerSide.Shade column in the Clickhouse.
	FieldCabinDoorRow2PassengerSideShade string = "Cabin.Door.Row2.PassengerSide.Shade"
	// FieldCabinDoorRow2PassengerSideShadeIsOpen is the name for the Cabin.Door.Row2.PassengerSide.Shade.IsOpen column in the Clickhouse.
	FieldCabinDoorRow2PassengerSideShadeIsOpen string = "Cabin.Door.Row2.PassengerSide.Shade.IsOpen"
	// FieldCabinDoorRow2PassengerSideShadePosition is the name for the Cabin.Door.Row2.PassengerSide.Shade.Position column in the Clickhouse.
	FieldCabinDoorRow2PassengerSideShadePosition string = "Cabin.Door.Row2.PassengerSide.Shade.Position"
	// FieldCabinDoorRow2PassengerSideShadeSwitch is the name for the Cabin.Door.Row2.PassengerSide.Shade.Switch column in the Clickhouse.
	FieldCabinDoorRow2PassengerSideShadeSwitch string = "Cabin.Door.Row2.PassengerSide.Shade.Switch"
	// FieldCabinDoorRow2PassengerSideSwitch is the name for the Cabin.Door.Row2.PassengerSide.Switch column in the Clickhouse.
	FieldCabinDoorRow2PassengerSideSwitch string = "Cabin.Door.Row2.PassengerSide.Switch"
	// FieldCabinDoorRow2PassengerSideWindow is the name for the Cabin.Door.Row2.PassengerSide.Window column in the Clickhouse.
	FieldCabinDoorRow2PassengerSideWindow string = "Cabin.Door.Row2.PassengerSide.Window"
	// FieldCabinDoorRow2PassengerSideWindowIsOpen is the name for the Cabin.Door.Row2.PassengerSide.Window.IsOpen column in the Clickhouse.
	FieldCabinDoorRow2PassengerSideWindowIsOpen string = "Cabin.Door.Row2.PassengerSide.Window.IsOpen"
	// FieldCabinDoorRow2PassengerSideWindowPosition is the name for the Cabin.Door.Row2.PassengerSide.Window.Position column in the Clickhouse.
	FieldCabinDoorRow2PassengerSideWindowPosition string = "Cabin.Door.Row2.PassengerSide.Window.Position"
	// FieldCabinDoorRow2PassengerSideWindowSwitch is the name for the Cabin.Door.Row2.PassengerSide.Window.Switch column in the Clickhouse.
	FieldCabinDoorRow2PassengerSideWindowSwitch string = "Cabin.Door.Row2.PassengerSide.Window.Switch"
	// FieldCabinDriverPosition is the name for the Cabin.DriverPosition column in the Clickhouse.
	FieldCabinDriverPosition string = "Cabin.DriverPosition"
	// FieldCabinHVAC is the name for the Cabin.HVAC column in the Clickhouse.
	FieldCabinHVAC string = "Cabin.HVAC"
	// FieldCabinHVACAmbientAirTemperature is the name for the Cabin.HVAC.AmbientAirTemperature column in the Clickhouse.
	FieldCabinHVACAmbientAirTemperature string = "Cabin.HVAC.AmbientAirTemperature"
	// FieldCabinHVACIsAirConditioningActive is the name for the Cabin.HVAC.IsAirConditioningActive column in the Clickhouse.
	FieldCabinHVACIsAirConditioningActive string = "Cabin.HVAC.IsAirConditioningActive"
	// FieldCabinHVACIsAutoPowerOptimize is the name for the Cabin.HVAC.IsAutoPowerOptimize column in the Clickhouse.
	FieldCabinHVACIsAutoPowerOptimize string = "Cabin.HVAC.IsAutoPowerOptimize"
	// FieldCabinHVACIsFrontDefrosterActive is the name for the Cabin.HVAC.IsFrontDefrosterActive column in the Clickhouse.
	FieldCabinHVACIsFrontDefrosterActive string = "Cabin.HVAC.IsFrontDefrosterActive"
	// FieldCabinHVACIsRearDefrosterActive is the name for the Cabin.HVAC.IsRearDefrosterActive column in the Clickhouse.
	FieldCabinHVACIsRearDefrosterActive string = "Cabin.HVAC.IsRearDefrosterActive"
	// FieldCabinHVACIsRecirculationActive is the name for the Cabin.HVAC.IsRecirculationActive column in the Clickhouse.
	FieldCabinHVACIsRecirculationActive string = "Cabin.HVAC.IsRecirculationActive"
	// FieldCabinHVACPowerOptimizeLevel is the name for the Cabin.HVAC.PowerOptimizeLevel column in the Clickhouse.
	FieldCabinHVACPowerOptimizeLevel string = "Cabin.HVAC.PowerOptimizeLevel"
	// FieldCabinHVACStation is the name for the Cabin.HVAC.Station column in the Clickhouse.
	FieldCabinHVACStation string = "Cabin.HVAC.Station"
	// FieldCabinHVACStationRow1 is the name for the Cabin.HVAC.Station.Row1 column in the Clickhouse.
	FieldCabinHVACStationRow1 string = "Cabin.HVAC.Station.Row1"
	// FieldCabinHVACStationRow1Driver is the name for the Cabin.HVAC.Station.Row1.Driver column in the Clickhouse.
	FieldCabinHVACStationRow1Driver string = "Cabin.HVAC.Station.Row1.Driver"
	// FieldCabinHVACStationRow1DriverAirDistribution is the name for the Cabin.HVAC.Station.Row1.Driver.AirDistribution column in the Clickhouse.
	FieldCabinHVACStationRow1DriverAirDistribution string = "Cabin.HVAC.Station.Row1.Driver.AirDistribution"
	// FieldCabinHVACStationRow1DriverFanSpeed is the name for the Cabin.HVAC.Station.Row1.Driver.FanSpeed column in the Clickhouse.
	FieldCabinHVACStationRow1DriverFanSpeed string = "Cabin.HVAC.Station.Row1.Driver.FanSpeed"
	// FieldCabinHVACStationRow1DriverTemperature is the name for the Cabin.HVAC.Station.Row1.Driver.Temperature column in the Clickhouse.
	FieldCabinHVACStationRow1DriverTemperature string = "Cabin.HVAC.Station.Row1.Driver.Temperature"
	// FieldCabinHVACStationRow1Passenger is the name for the Cabin.HVAC.Station.Row1.Passenger column in the Clickhouse.
	FieldCabinHVACStationRow1Passenger string = "Cabin.HVAC.Station.Row1.Passenger"
	// FieldCabinHVACStationRow1PassengerAirDistribution is the name for the Cabin.HVAC.Station.Row1.Passenger.AirDistribution column in the Clickhouse.
	FieldCabinHVACStationRow1PassengerAirDistribution string = "Cabin.HVAC.Station.Row1.Passenger.AirDistribution"
	// FieldCabinHVACStationRow1PassengerFanSpeed is the name for the Cabin.HVAC.Station.Row1.Passenger.FanSpeed column in the Clickhouse.
	FieldCabinHVACStationRow1PassengerFanSpeed string = "Cabin.HVAC.Station.Row1.Passenger.FanSpeed"
	// FieldCabinHVACStationRow1PassengerTemperature is the name for the Cabin.HVAC.Station.Row1.Passenger.Temperature column in the Clickhouse.
	FieldCabinHVACStationRow1PassengerTemperature string = "Cabin.HVAC.Station.Row1.Passenger.Temperature"
	// FieldCabinHVACStationRow2 is the name for the Cabin.HVAC.Station.Row2 column in the Clickhouse.
	FieldCabinHVACStationRow2 string = "Cabin.HVAC.Station.Row2"
	// FieldCabinHVACStationRow2Driver is the name for the Cabin.HVAC.Station.Row2.Driver column in the Clickhouse.
	FieldCabinHVACStationRow2Driver string = "Cabin.HVAC.Station.Row2.Driver"
	// FieldCabinHVACStationRow2DriverAirDistribution is the name for the Cabin.HVAC.Station.Row2.Driver.AirDistribution column in the Clickhouse.
	FieldCabinHVACStationRow2DriverAirDistribution string = "Cabin.HVAC.Station.Row2.Driver.AirDistribution"
	// FieldCabinHVACStationRow2DriverFanSpeed is the name for the Cabin.HVAC.Station.Row2.Driver.FanSpeed column in the Clickhouse.
	FieldCabinHVACStationRow2DriverFanSpeed string = "Cabin.HVAC.Station.Row2.Driver.FanSpeed"
	// FieldCabinHVACStationRow2DriverTemperature is the name for the Cabin.HVAC.Station.Row2.Driver.Temperature column in the Clickhouse.
	FieldCabinHVACStationRow2DriverTemperature string = "Cabin.HVAC.Station.Row2.Driver.Temperature"
	// FieldCabinHVACStationRow2Passenger is the name for the Cabin.HVAC.Station.Row2.Passenger column in the Clickhouse.
	FieldCabinHVACStationRow2Passenger string = "Cabin.HVAC.Station.Row2.Passenger"
	// FieldCabinHVACStationRow2PassengerAirDistribution is the name for the Cabin.HVAC.Station.Row2.Passenger.AirDistribution column in the Clickhouse.
	FieldCabinHVACStationRow2PassengerAirDistribution string = "Cabin.HVAC.Station.Row2.Passenger.AirDistribution"
	// FieldCabinHVACStationRow2PassengerFanSpeed is the name for the Cabin.HVAC.Station.Row2.Passenger.FanSpeed column in the Clickhouse.
	FieldCabinHVACStationRow2PassengerFanSpeed string = "Cabin.HVAC.Station.Row2.Passenger.FanSpeed"
	// FieldCabinHVACStationRow2PassengerTemperature is the name for the Cabin.HVAC.Station.Row2.Passenger.Temperature column in the Clickhouse.
	FieldCabinHVACStationRow2PassengerTemperature string = "Cabin.HVAC.Station.Row2.Passenger.Temperature"
	// FieldCabinHVACStationRow3 is the name for the Cabin.HVAC.Station.Row3 column in the Clickhouse.
	FieldCabinHVACStationRow3 string = "Cabin.HVAC.Station.Row3"
	// FieldCabinHVACStationRow3Driver is the name for the Cabin.HVAC.Station.Row3.Driver column in the Clickhouse.
	FieldCabinHVACStationRow3Driver string = "Cabin.HVAC.Station.Row3.Driver"
	// FieldCabinHVACStationRow3DriverAirDistribution is the name for the Cabin.HVAC.Station.Row3.Driver.AirDistribution column in the Clickhouse.
	FieldCabinHVACStationRow3DriverAirDistribution string = "Cabin.HVAC.Station.Row3.Driver.AirDistribution"
	// FieldCabinHVACStationRow3DriverFanSpeed is the name for the Cabin.HVAC.Station.Row3.Driver.FanSpeed column in the Clickhouse.
	FieldCabinHVACStationRow3DriverFanSpeed string = "Cabin.HVAC.Station.Row3.Driver.FanSpeed"
	// FieldCabinHVACStationRow3DriverTemperature is the name for the Cabin.HVAC.Station.Row3.Driver.Temperature column in the Clickhouse.
	FieldCabinHVACStationRow3DriverTemperature string = "Cabin.HVAC.Station.Row3.Driver.Temperature"
	// FieldCabinHVACStationRow3Passenger is the name for the Cabin.HVAC.Station.Row3.Passenger column in the Clickhouse.
	FieldCabinHVACStationRow3Passenger string = "Cabin.HVAC.Station.Row3.Passenger"
	// FieldCabinHVACStationRow3PassengerAirDistribution is the name for the Cabin.HVAC.Station.Row3.Passenger.AirDistribution column in the Clickhouse.
	FieldCabinHVACStationRow3PassengerAirDistribution string = "Cabin.HVAC.Station.Row3.Passenger.AirDistribution"
	// FieldCabinHVACStationRow3PassengerFanSpeed is the name for the Cabin.HVAC.Station.Row3.Passenger.FanSpeed column in the Clickhouse.
	FieldCabinHVACStationRow3PassengerFanSpeed string = "Cabin.HVAC.Station.Row3.Passenger.FanSpeed"
	// FieldCabinHVACStationRow3PassengerTemperature is the name for the Cabin.HVAC.Station.Row3.Passenger.Temperature column in the Clickhouse.
	FieldCabinHVACStationRow3PassengerTemperature string = "Cabin.HVAC.Station.Row3.Passenger.Temperature"
	// FieldCabinHVACStationRow4 is the name for the Cabin.HVAC.Station.Row4 column in the Clickhouse.
	FieldCabinHVACStationRow4 string = "Cabin.HVAC.Station.Row4"
	// FieldCabinHVACStationRow4Driver is the name for the Cabin.HVAC.Station.Row4.Driver column in the Clickhouse.
	FieldCabinHVACStationRow4Driver string = "Cabin.HVAC.Station.Row4.Driver"
	// FieldCabinHVACStationRow4DriverAirDistribution is the name for the Cabin.HVAC.Station.Row4.Driver.AirDistribution column in the Clickhouse.
	FieldCabinHVACStationRow4DriverAirDistribution string = "Cabin.HVAC.Station.Row4.Driver.AirDistribution"
	// FieldCabinHVACStationRow4DriverFanSpeed is the name for the Cabin.HVAC.Station.Row4.Driver.FanSpeed column in the Clickhouse.
	FieldCabinHVACStationRow4DriverFanSpeed string = "Cabin.HVAC.Station.Row4.Driver.FanSpeed"
	// FieldCabinHVACStationRow4DriverTemperature is the name for the Cabin.HVAC.Station.Row4.Driver.Temperature column in the Clickhouse.
	FieldCabinHVACStationRow4DriverTemperature string = "Cabin.HVAC.Station.Row4.Driver.Temperature"
	// FieldCabinHVACStationRow4Passenger is the name for the Cabin.HVAC.Station.Row4.Passenger column in the Clickhouse.
	FieldCabinHVACStationRow4Passenger string = "Cabin.HVAC.Station.Row4.Passenger"
	// FieldCabinHVACStationRow4PassengerAirDistribution is the name for the Cabin.HVAC.Station.Row4.Passenger.AirDistribution column in the Clickhouse.
	FieldCabinHVACStationRow4PassengerAirDistribution string = "Cabin.HVAC.Station.Row4.Passenger.AirDistribution"
	// FieldCabinHVACStationRow4PassengerFanSpeed is the name for the Cabin.HVAC.Station.Row4.Passenger.FanSpeed column in the Clickhouse.
	FieldCabinHVACStationRow4PassengerFanSpeed string = "Cabin.HVAC.Station.Row4.Passenger.FanSpeed"
	// FieldCabinHVACStationRow4PassengerTemperature is the name for the Cabin.HVAC.Station.Row4.Passenger.Temperature column in the Clickhouse.
	FieldCabinHVACStationRow4PassengerTemperature string = "Cabin.HVAC.Station.Row4.Passenger.Temperature"
	// FieldCabinInfotainment is the name for the Cabin.Infotainment column in the Clickhouse.
	FieldCabinInfotainment string = "Cabin.Infotainment"
	// FieldCabinInfotainmentHMI is the name for the Cabin.Infotainment.HMI column in the Clickhouse.
	FieldCabinInfotainmentHMI string = "Cabin.Infotainment.HMI"
	// FieldCabinInfotainmentHMIBrightness is the name for the Cabin.Infotainment.HMI.Brightness column in the Clickhouse.
	FieldCabinInfotainmentHMIBrightness string = "Cabin.Infotainment.HMI.Brightness"
	// FieldCabinInfotainmentHMICurrentLanguage is the name for the Cabin.Infotainment.HMI.CurrentLanguage column in the Clickhouse.
	FieldCabinInfotainmentHMICurrentLanguage string = "Cabin.Infotainment.HMI.CurrentLanguage"
	// FieldCabinInfotainmentHMIDateFormat is the name for the Cabin.Infotainment.HMI.DateFormat column in the Clickhouse.
	FieldCabinInfotainmentHMIDateFormat string = "Cabin.Infotainment.HMI.DateFormat"
	// FieldCabinInfotainmentHMIDayNightMode is the name for the Cabin.Infotainment.HMI.DayNightMode column in the Clickhouse.
	FieldCabinInfotainmentHMIDayNightMode string = "Cabin.Infotainment.HMI.DayNightMode"
	// FieldCabinInfotainmentHMIDisplayOffDuration is the name for the Cabin.Infotainment.HMI.DisplayOffDuration column in the Clickhouse.
	FieldCabinInfotainmentHMIDisplayOffDuration string = "Cabin.Infotainment.HMI.DisplayOffDuration"
	// FieldCabinInfotainmentHMIDistanceUnit is the name for the Cabin.Infotainment.HMI.DistanceUnit column in the Clickhouse.
	FieldCabinInfotainmentHMIDistanceUnit string = "Cabin.Infotainment.HMI.DistanceUnit"
	// FieldCabinInfotainmentHMIEVEconomyUnits is the name for the Cabin.Infotainment.HMI.EVEconomyUnits column in the Clickhouse.
	FieldCabinInfotainmentHMIEVEconomyUnits string = "Cabin.Infotainment.HMI.EVEconomyUnits"
	// FieldCabinInfotainmentHMIEVEnergyUnits is the name for the Cabin.Infotainment.HMI.EVEnergyUnits column in the Clickhouse.
	FieldCabinInfotainmentHMIEVEnergyUnits string = "Cabin.Infotainment.HMI.EVEnergyUnits"
	// FieldCabinInfotainmentHMIFontSize is the name for the Cabin.Infotainment.HMI.FontSize column in the Clickhouse.
	FieldCabinInfotainmentHMIFontSize string = "Cabin.Infotainment.HMI.FontSize"
	// FieldCabinInfotainmentHMIFuelEconomyUnits is the name for the Cabin.Infotainment.HMI.FuelEconomyUnits column in the Clickhouse.
	FieldCabinInfotainmentHMIFuelEconomyUnits string = "Cabin.Infotainment.HMI.FuelEconomyUnits"
	// FieldCabinInfotainmentHMIFuelVolumeUnit is the name for the Cabin.Infotainment.HMI.FuelVolumeUnit column in the Clickhouse.
	FieldCabinInfotainmentHMIFuelVolumeUnit string = "Cabin.Infotainment.HMI.FuelVolumeUnit"
	// FieldCabinInfotainmentHMIIsScreenAlwaysOn is the name for the Cabin.Infotainment.HMI.IsScreenAlwaysOn column in the Clickhouse.
	FieldCabinInfotainmentHMIIsScreenAlwaysOn string = "Cabin.Infotainment.HMI.IsScreenAlwaysOn"
	// FieldCabinInfotainmentHMILastActionTime is the name for the Cabin.Infotainment.HMI.LastActionTime column in the Clickhouse.
	FieldCabinInfotainmentHMILastActionTime string = "Cabin.Infotainment.HMI.LastActionTime"
	// FieldCabinInfotainmentHMISpeedUnit is the name for the Cabin.Infotainment.HMI.SpeedUnit column in the Clickhouse.
	FieldCabinInfotainmentHMISpeedUnit string = "Cabin.Infotainment.HMI.SpeedUnit"
	// FieldCabinInfotainmentHMITemperatureUnit is the name for the Cabin.Infotainment.HMI.TemperatureUnit column in the Clickhouse.
	FieldCabinInfotainmentHMITemperatureUnit string = "Cabin.Infotainment.HMI.TemperatureUnit"
	// FieldCabinInfotainmentHMITimeFormat is the name for the Cabin.Infotainment.HMI.TimeFormat column in the Clickhouse.
	FieldCabinInfotainmentHMITimeFormat string = "Cabin.Infotainment.HMI.TimeFormat"
	// FieldCabinInfotainmentHMITirePressureUnit is the name for the Cabin.Infotainment.HMI.TirePressureUnit column in the Clickhouse.
	FieldCabinInfotainmentHMITirePressureUnit string = "Cabin.Infotainment.HMI.TirePressureUnit"
	// FieldCabinInfotainmentIsAutoPowerOptimize is the name for the Cabin.Infotainment.IsAutoPowerOptimize column in the Clickhouse.
	FieldCabinInfotainmentIsAutoPowerOptimize string = "Cabin.Infotainment.IsAutoPowerOptimize"
	// FieldCabinInfotainmentMedia is the name for the Cabin.Infotainment.Media column in the Clickhouse.
	FieldCabinInfotainmentMedia string = "Cabin.Infotainment.Media"
	// FieldCabinInfotainmentMediaAction is the name for the Cabin.Infotainment.Media.Action column in the Clickhouse.
	FieldCabinInfotainmentMediaAction string = "Cabin.Infotainment.Media.Action"
	// FieldCabinInfotainmentMediaDeclinedURI is the name for the Cabin.Infotainment.Media.DeclinedURI column in the Clickhouse.
	FieldCabinInfotainmentMediaDeclinedURI string = "Cabin.Infotainment.Media.DeclinedURI"
	// FieldCabinInfotainmentMediaPlayed is the name for the Cabin.Infotainment.Media.Played column in the Clickhouse.
	FieldCabinInfotainmentMediaPlayed string = "Cabin.Infotainment.Media.Played"
	// FieldCabinInfotainmentMediaPlayedAlbum is the name for the Cabin.Infotainment.Media.Played.Album column in the Clickhouse.
	FieldCabinInfotainmentMediaPlayedAlbum string = "Cabin.Infotainment.Media.Played.Album"
	// FieldCabinInfotainmentMediaPlayedArtist is the name for the Cabin.Infotainment.Media.Played.Artist column in the Clickhouse.
	FieldCabinInfotainmentMediaPlayedArtist string = "Cabin.Infotainment.Media.Played.Artist"
	// FieldCabinInfotainmentMediaPlayedPlaybackRate is the name for the Cabin.Infotainment.Media.Played.PlaybackRate column in the Clickhouse.
	FieldCabinInfotainmentMediaPlayedPlaybackRate string = "Cabin.Infotainment.Media.Played.PlaybackRate"
	// FieldCabinInfotainmentMediaPlayedSource is the name for the Cabin.Infotainment.Media.Played.Source column in the Clickhouse.
	FieldCabinInfotainmentMediaPlayedSource string = "Cabin.Infotainment.Media.Played.Source"
	// FieldCabinInfotainmentMediaPlayedTrack is the name for the Cabin.Infotainment.Media.Played.Track column in the Clickhouse.
	FieldCabinInfotainmentMediaPlayedTrack string = "Cabin.Infotainment.Media.Played.Track"
	// FieldCabinInfotainmentMediaPlayedURI is the name for the Cabin.Infotainment.Media.Played.URI column in the Clickhouse.
	FieldCabinInfotainmentMediaPlayedURI string = "Cabin.Infotainment.Media.Played.URI"
	// FieldCabinInfotainmentMediaSelectedURI is the name for the Cabin.Infotainment.Media.SelectedURI column in the Clickhouse.
	FieldCabinInfotainmentMediaSelectedURI string = "Cabin.Infotainment.Media.SelectedURI"
	// FieldCabinInfotainmentMediaVolume is the name for the Cabin.Infotainment.Media.Volume column in the Clickhouse.
	FieldCabinInfotainmentMediaVolume string = "Cabin.Infotainment.Media.Volume"
	// FieldCabinInfotainmentNavigation is the name for the Cabin.Infotainment.Navigation column in the Clickhouse.
	FieldCabinInfotainmentNavigation string = "Cabin.Infotainment.Navigation"
	// FieldCabinInfotainmentNavigationDestinationSet is the name for the Cabin.Infotainment.Navigation.DestinationSet column in the Clickhouse.
	FieldCabinInfotainmentNavigationDestinationSet string = "Cabin.Infotainment.Navigation.DestinationSet"
	// FieldCabinInfotainmentNavigationDestinationSetLatitude is the name for the Cabin.Infotainment.Navigation.DestinationSet.Latitude column in the Clickhouse.
	FieldCabinInfotainmentNavigationDestinationSetLatitude string = "Cabin.Infotainment.Navigation.DestinationSet.Latitude"
	// FieldCabinInfotainmentNavigationDestinationSetLongitude is the name for the Cabin.Infotainment.Navigation.DestinationSet.Longitude column in the Clickhouse.
	FieldCabinInfotainmentNavigationDestinationSetLongitude string = "Cabin.Infotainment.Navigation.DestinationSet.Longitude"
	// FieldCabinInfotainmentNavigationGuidanceVoice is the name for the Cabin.Infotainment.Navigation.GuidanceVoice column in the Clickhouse.
	FieldCabinInfotainmentNavigationGuidanceVoice string = "Cabin.Infotainment.Navigation.GuidanceVoice"
	// FieldCabinInfotainmentNavigationMapIsAutoScaleModeUsed is the name for the Cabin.Infotainment.Navigation.Map_IsAutoScaleModeUsed column in the Clickhouse.
	FieldCabinInfotainmentNavigationMapIsAutoScaleModeUsed string = "Cabin.Infotainment.Navigation.Map_IsAutoScaleModeUsed"
	// FieldCabinInfotainmentNavigationMute is the name for the Cabin.Infotainment.Navigation.Mute column in the Clickhouse.
	FieldCabinInfotainmentNavigationMute string = "Cabin.Infotainment.Navigation.Mute"
	// FieldCabinInfotainmentNavigationVolume is the name for the Cabin.Infotainment.Navigation.Volume column in the Clickhouse.
	FieldCabinInfotainmentNavigationVolume string = "Cabin.Infotainment.Navigation.Volume"
	// FieldCabinInfotainmentPowerOptimizeLevel is the name for the Cabin.Infotainment.PowerOptimizeLevel column in the Clickhouse.
	FieldCabinInfotainmentPowerOptimizeLevel string = "Cabin.Infotainment.PowerOptimizeLevel"
	// FieldCabinInfotainmentSmartphoneProjection is the name for the Cabin.Infotainment.SmartphoneProjection column in the Clickhouse.
	FieldCabinInfotainmentSmartphoneProjection string = "Cabin.Infotainment.SmartphoneProjection"
	// FieldCabinInfotainmentSmartphoneProjectionActive is the name for the Cabin.Infotainment.SmartphoneProjection.Active column in the Clickhouse.
	FieldCabinInfotainmentSmartphoneProjectionActive string = "Cabin.Infotainment.SmartphoneProjection.Active"
	// FieldCabinInfotainmentSmartphoneProjectionSource is the name for the Cabin.Infotainment.SmartphoneProjection.Source column in the Clickhouse.
	FieldCabinInfotainmentSmartphoneProjectionSource string = "Cabin.Infotainment.SmartphoneProjection.Source"
	// FieldCabinInfotainmentSmartphoneProjectionSupportedMode is the name for the Cabin.Infotainment.SmartphoneProjection.SupportedMode column in the Clickhouse.
	FieldCabinInfotainmentSmartphoneProjectionSupportedMode string = "Cabin.Infotainment.SmartphoneProjection.SupportedMode"
	// FieldCabinInfotainmentSmartphoneProjectionSupportedModesize0 is the name for the Cabin.Infotainment.SmartphoneProjection.SupportedMode.size0 column in the Clickhouse.
	FieldCabinInfotainmentSmartphoneProjectionSupportedModesize0 string = "Cabin.Infotainment.SmartphoneProjection.SupportedMode.size0"
	// FieldCabinIsAutoPowerOptimize is the name for the Cabin.IsAutoPowerOptimize column in the Clickhouse.
	FieldCabinIsAutoPowerOptimize string = "Cabin.IsAutoPowerOptimize"
	// FieldCabinIsWindowChildLockEngaged is the name for the Cabin.IsWindowChildLockEngaged column in the Clickhouse.
	FieldCabinIsWindowChildLockEngaged string = "Cabin.IsWindowChildLockEngaged"
	// FieldCabinLight is the name for the Cabin.Light column in the Clickhouse.
	FieldCabinLight string = "Cabin.Light"
	// FieldCabinLightAmbientLight is the name for the Cabin.Light.AmbientLight column in the Clickhouse.
	FieldCabinLightAmbientLight string = "Cabin.Light.AmbientLight"
	// FieldCabinLightAmbientLightRow1 is the name for the Cabin.Light.AmbientLight.Row1 column in the Clickhouse.
	FieldCabinLightAmbientLightRow1 string = "Cabin.Light.AmbientLight.Row1"
	// FieldCabinLightAmbientLightRow1DriverSide is the name for the Cabin.Light.AmbientLight.Row1.DriverSide column in the Clickhouse.
	FieldCabinLightAmbientLightRow1DriverSide string = "Cabin.Light.AmbientLight.Row1.DriverSide"
	// FieldCabinLightAmbientLightRow1DriverSideColor is the name for the Cabin.Light.AmbientLight.Row1.DriverSide.Color column in the Clickhouse.
	FieldCabinLightAmbientLightRow1DriverSideColor string = "Cabin.Light.AmbientLight.Row1.DriverSide.Color"
	// FieldCabinLightAmbientLightRow1DriverSideIntensity is the name for the Cabin.Light.AmbientLight.Row1.DriverSide.Intensity column in the Clickhouse.
	FieldCabinLightAmbientLightRow1DriverSideIntensity string = "Cabin.Light.AmbientLight.Row1.DriverSide.Intensity"
	// FieldCabinLightAmbientLightRow1DriverSideIsLightOn is the name for the Cabin.Light.AmbientLight.Row1.DriverSide.IsLightOn column in the Clickhouse.
	FieldCabinLightAmbientLightRow1DriverSideIsLightOn string = "Cabin.Light.AmbientLight.Row1.DriverSide.IsLightOn"
	// FieldCabinLightAmbientLightRow1PassengerSide is the name for the Cabin.Light.AmbientLight.Row1.PassengerSide column in the Clickhouse.
	FieldCabinLightAmbientLightRow1PassengerSide string = "Cabin.Light.AmbientLight.Row1.PassengerSide"
	// FieldCabinLightAmbientLightRow1PassengerSideColor is the name for the Cabin.Light.AmbientLight.Row1.PassengerSide.Color column in the Clickhouse.
	FieldCabinLightAmbientLightRow1PassengerSideColor string = "Cabin.Light.AmbientLight.Row1.PassengerSide.Color"
	// FieldCabinLightAmbientLightRow1PassengerSideIntensity is the name for the Cabin.Light.AmbientLight.Row1.PassengerSide.Intensity column in the Clickhouse.
	FieldCabinLightAmbientLightRow1PassengerSideIntensity string = "Cabin.Light.AmbientLight.Row1.PassengerSide.Intensity"
	// FieldCabinLightAmbientLightRow1PassengerSideIsLightOn is the name for the Cabin.Light.AmbientLight.Row1.PassengerSide.IsLightOn column in the Clickhouse.
	FieldCabinLightAmbientLightRow1PassengerSideIsLightOn string = "Cabin.Light.AmbientLight.Row1.PassengerSide.IsLightOn"
	// FieldCabinLightAmbientLightRow2 is the name for the Cabin.Light.AmbientLight.Row2 column in the Clickhouse.
	FieldCabinLightAmbientLightRow2 string = "Cabin.Light.AmbientLight.Row2"
	// FieldCabinLightAmbientLightRow2DriverSide is the name for the Cabin.Light.AmbientLight.Row2.DriverSide column in the Clickhouse.
	FieldCabinLightAmbientLightRow2DriverSide string = "Cabin.Light.AmbientLight.Row2.DriverSide"
	// FieldCabinLightAmbientLightRow2DriverSideColor is the name for the Cabin.Light.AmbientLight.Row2.DriverSide.Color column in the Clickhouse.
	FieldCabinLightAmbientLightRow2DriverSideColor string = "Cabin.Light.AmbientLight.Row2.DriverSide.Color"
	// FieldCabinLightAmbientLightRow2DriverSideIntensity is the name for the Cabin.Light.AmbientLight.Row2.DriverSide.Intensity column in the Clickhouse.
	FieldCabinLightAmbientLightRow2DriverSideIntensity string = "Cabin.Light.AmbientLight.Row2.DriverSide.Intensity"
	// FieldCabinLightAmbientLightRow2DriverSideIsLightOn is the name for the Cabin.Light.AmbientLight.Row2.DriverSide.IsLightOn column in the Clickhouse.
	FieldCabinLightAmbientLightRow2DriverSideIsLightOn string = "Cabin.Light.AmbientLight.Row2.DriverSide.IsLightOn"
	// FieldCabinLightAmbientLightRow2PassengerSide is the name for the Cabin.Light.AmbientLight.Row2.PassengerSide column in the Clickhouse.
	FieldCabinLightAmbientLightRow2PassengerSide string = "Cabin.Light.AmbientLight.Row2.PassengerSide"
	// FieldCabinLightAmbientLightRow2PassengerSideColor is the name for the Cabin.Light.AmbientLight.Row2.PassengerSide.Color column in the Clickhouse.
	FieldCabinLightAmbientLightRow2PassengerSideColor string = "Cabin.Light.AmbientLight.Row2.PassengerSide.Color"
	// FieldCabinLightAmbientLightRow2PassengerSideIntensity is the name for the Cabin.Light.AmbientLight.Row2.PassengerSide.Intensity column in the Clickhouse.
	FieldCabinLightAmbientLightRow2PassengerSideIntensity string = "Cabin.Light.AmbientLight.Row2.PassengerSide.Intensity"
	// FieldCabinLightAmbientLightRow2PassengerSideIsLightOn is the name for the Cabin.Light.AmbientLight.Row2.PassengerSide.IsLightOn column in the Clickhouse.
	FieldCabinLightAmbientLightRow2PassengerSideIsLightOn string = "Cabin.Light.AmbientLight.Row2.PassengerSide.IsLightOn"
	// FieldCabinLightInteractiveLightBar is the name for the Cabin.Light.InteractiveLightBar column in the Clickhouse.
	FieldCabinLightInteractiveLightBar string = "Cabin.Light.InteractiveLightBar"
	// FieldCabinLightInteractiveLightBarColor is the name for the Cabin.Light.InteractiveLightBar.Color column in the Clickhouse.
	FieldCabinLightInteractiveLightBarColor string = "Cabin.Light.InteractiveLightBar.Color"
	// FieldCabinLightInteractiveLightBarEffect is the name for the Cabin.Light.InteractiveLightBar.Effect column in the Clickhouse.
	FieldCabinLightInteractiveLightBarEffect string = "Cabin.Light.InteractiveLightBar.Effect"
	// FieldCabinLightInteractiveLightBarIntensity is the name for the Cabin.Light.InteractiveLightBar.Intensity column in the Clickhouse.
	FieldCabinLightInteractiveLightBarIntensity string = "Cabin.Light.InteractiveLightBar.Intensity"
	// FieldCabinLightInteractiveLightBarIsLightOn is the name for the Cabin.Light.InteractiveLightBar.IsLightOn column in the Clickhouse.
	FieldCabinLightInteractiveLightBarIsLightOn string = "Cabin.Light.InteractiveLightBar.IsLightOn"
	// FieldCabinLightIsDomeOn is the name for the Cabin.Light.IsDomeOn column in the Clickhouse.
	FieldCabinLightIsDomeOn string = "Cabin.Light.IsDomeOn"
	// FieldCabinLightIsGloveBoxOn is the name for the Cabin.Light.IsGloveBoxOn column in the Clickhouse.
	FieldCabinLightIsGloveBoxOn string = "Cabin.Light.IsGloveBoxOn"
	// FieldCabinLightPerceivedAmbientLight is the name for the Cabin.Light.PerceivedAmbientLight column in the Clickhouse.
	FieldCabinLightPerceivedAmbientLight string = "Cabin.Light.PerceivedAmbientLight"
	// FieldCabinLightSpotlight is the name for the Cabin.Light.Spotlight column in the Clickhouse.
	FieldCabinLightSpotlight string = "Cabin.Light.Spotlight"
	// FieldCabinLightSpotlightRow1 is the name for the Cabin.Light.Spotlight.Row1 column in the Clickhouse.
	FieldCabinLightSpotlightRow1 string = "Cabin.Light.Spotlight.Row1"
	// FieldCabinLightSpotlightRow1DriverSide is the name for the Cabin.Light.Spotlight.Row1.DriverSide column in the Clickhouse.
	FieldCabinLightSpotlightRow1DriverSide string = "Cabin.Light.Spotlight.Row1.DriverSide"
	// FieldCabinLightSpotlightRow1DriverSideColor is the name for the Cabin.Light.Spotlight.Row1.DriverSide.Color column in the Clickhouse.
	FieldCabinLightSpotlightRow1DriverSideColor string = "Cabin.Light.Spotlight.Row1.DriverSide.Color"
	// FieldCabinLightSpotlightRow1DriverSideIntensity is the name for the Cabin.Light.Spotlight.Row1.DriverSide.Intensity column in the Clickhouse.
	FieldCabinLightSpotlightRow1DriverSideIntensity string = "Cabin.Light.Spotlight.Row1.DriverSide.Intensity"
	// FieldCabinLightSpotlightRow1DriverSideIsLightOn is the name for the Cabin.Light.Spotlight.Row1.DriverSide.IsLightOn column in the Clickhouse.
	FieldCabinLightSpotlightRow1DriverSideIsLightOn string = "Cabin.Light.Spotlight.Row1.DriverSide.IsLightOn"
	// FieldCabinLightSpotlightRow1PassengerSide is the name for the Cabin.Light.Spotlight.Row1.PassengerSide column in the Clickhouse.
	FieldCabinLightSpotlightRow1PassengerSide string = "Cabin.Light.Spotlight.Row1.PassengerSide"
	// FieldCabinLightSpotlightRow1PassengerSideColor is the name for the Cabin.Light.Spotlight.Row1.PassengerSide.Color column in the Clickhouse.
	FieldCabinLightSpotlightRow1PassengerSideColor string = "Cabin.Light.Spotlight.Row1.PassengerSide.Color"
	// FieldCabinLightSpotlightRow1PassengerSideIntensity is the name for the Cabin.Light.Spotlight.Row1.PassengerSide.Intensity column in the Clickhouse.
	FieldCabinLightSpotlightRow1PassengerSideIntensity string = "Cabin.Light.Spotlight.Row1.PassengerSide.Intensity"
	// FieldCabinLightSpotlightRow1PassengerSideIsLightOn is the name for the Cabin.Light.Spotlight.Row1.PassengerSide.IsLightOn column in the Clickhouse.
	FieldCabinLightSpotlightRow1PassengerSideIsLightOn string = "Cabin.Light.Spotlight.Row1.PassengerSide.IsLightOn"
	// FieldCabinLightSpotlightRow2 is the name for the Cabin.Light.Spotlight.Row2 column in the Clickhouse.
	FieldCabinLightSpotlightRow2 string = "Cabin.Light.Spotlight.Row2"
	// FieldCabinLightSpotlightRow2DriverSide is the name for the Cabin.Light.Spotlight.Row2.DriverSide column in the Clickhouse.
	FieldCabinLightSpotlightRow2DriverSide string = "Cabin.Light.Spotlight.Row2.DriverSide"
	// FieldCabinLightSpotlightRow2DriverSideColor is the name for the Cabin.Light.Spotlight.Row2.DriverSide.Color column in the Clickhouse.
	FieldCabinLightSpotlightRow2DriverSideColor string = "Cabin.Light.Spotlight.Row2.DriverSide.Color"
	// FieldCabinLightSpotlightRow2DriverSideIntensity is the name for the Cabin.Light.Spotlight.Row2.DriverSide.Intensity column in the Clickhouse.
	FieldCabinLightSpotlightRow2DriverSideIntensity string = "Cabin.Light.Spotlight.Row2.DriverSide.Intensity"
	// FieldCabinLightSpotlightRow2DriverSideIsLightOn is the name for the Cabin.Light.Spotlight.Row2.DriverSide.IsLightOn column in the Clickhouse.
	FieldCabinLightSpotlightRow2DriverSideIsLightOn string = "Cabin.Light.Spotlight.Row2.DriverSide.IsLightOn"
	// FieldCabinLightSpotlightRow2PassengerSide is the name for the Cabin.Light.Spotlight.Row2.PassengerSide column in the Clickhouse.
	FieldCabinLightSpotlightRow2PassengerSide string = "Cabin.Light.Spotlight.Row2.PassengerSide"
	// FieldCabinLightSpotlightRow2PassengerSideColor is the name for the Cabin.Light.Spotlight.Row2.PassengerSide.Color column in the Clickhouse.
	FieldCabinLightSpotlightRow2PassengerSideColor string = "Cabin.Light.Spotlight.Row2.PassengerSide.Color"
	// FieldCabinLightSpotlightRow2PassengerSideIntensity is the name for the Cabin.Light.Spotlight.Row2.PassengerSide.Intensity column in the Clickhouse.
	FieldCabinLightSpotlightRow2PassengerSideIntensity string = "Cabin.Light.Spotlight.Row2.PassengerSide.Intensity"
	// FieldCabinLightSpotlightRow2PassengerSideIsLightOn is the name for the Cabin.Light.Spotlight.Row2.PassengerSide.IsLightOn column in the Clickhouse.
	FieldCabinLightSpotlightRow2PassengerSideIsLightOn string = "Cabin.Light.Spotlight.Row2.PassengerSide.IsLightOn"
	// FieldCabinLightSpotlightRow3 is the name for the Cabin.Light.Spotlight.Row3 column in the Clickhouse.
	FieldCabinLightSpotlightRow3 string = "Cabin.Light.Spotlight.Row3"
	// FieldCabinLightSpotlightRow3DriverSide is the name for the Cabin.Light.Spotlight.Row3.DriverSide column in the Clickhouse.
	FieldCabinLightSpotlightRow3DriverSide string = "Cabin.Light.Spotlight.Row3.DriverSide"
	// FieldCabinLightSpotlightRow3DriverSideColor is the name for the Cabin.Light.Spotlight.Row3.DriverSide.Color column in the Clickhouse.
	FieldCabinLightSpotlightRow3DriverSideColor string = "Cabin.Light.Spotlight.Row3.DriverSide.Color"
	// FieldCabinLightSpotlightRow3DriverSideIntensity is the name for the Cabin.Light.Spotlight.Row3.DriverSide.Intensity column in the Clickhouse.
	FieldCabinLightSpotlightRow3DriverSideIntensity string = "Cabin.Light.Spotlight.Row3.DriverSide.Intensity"
	// FieldCabinLightSpotlightRow3DriverSideIsLightOn is the name for the Cabin.Light.Spotlight.Row3.DriverSide.IsLightOn column in the Clickhouse.
	FieldCabinLightSpotlightRow3DriverSideIsLightOn string = "Cabin.Light.Spotlight.Row3.DriverSide.IsLightOn"
	// FieldCabinLightSpotlightRow3PassengerSide is the name for the Cabin.Light.Spotlight.Row3.PassengerSide column in the Clickhouse.
	FieldCabinLightSpotlightRow3PassengerSide string = "Cabin.Light.Spotlight.Row3.PassengerSide"
	// FieldCabinLightSpotlightRow3PassengerSideColor is the name for the Cabin.Light.Spotlight.Row3.PassengerSide.Color column in the Clickhouse.
	FieldCabinLightSpotlightRow3PassengerSideColor string = "Cabin.Light.Spotlight.Row3.PassengerSide.Color"
	// FieldCabinLightSpotlightRow3PassengerSideIntensity is the name for the Cabin.Light.Spotlight.Row3.PassengerSide.Intensity column in the Clickhouse.
	FieldCabinLightSpotlightRow3PassengerSideIntensity string = "Cabin.Light.Spotlight.Row3.PassengerSide.Intensity"
	// FieldCabinLightSpotlightRow3PassengerSideIsLightOn is the name for the Cabin.Light.Spotlight.Row3.PassengerSide.IsLightOn column in the Clickhouse.
	FieldCabinLightSpotlightRow3PassengerSideIsLightOn string = "Cabin.Light.Spotlight.Row3.PassengerSide.IsLightOn"
	// FieldCabinLightSpotlightRow4 is the name for the Cabin.Light.Spotlight.Row4 column in the Clickhouse.
	FieldCabinLightSpotlightRow4 string = "Cabin.Light.Spotlight.Row4"
	// FieldCabinLightSpotlightRow4DriverSide is the name for the Cabin.Light.Spotlight.Row4.DriverSide column in the Clickhouse.
	FieldCabinLightSpotlightRow4DriverSide string = "Cabin.Light.Spotlight.Row4.DriverSide"
	// FieldCabinLightSpotlightRow4DriverSideColor is the name for the Cabin.Light.Spotlight.Row4.DriverSide.Color column in the Clickhouse.
	FieldCabinLightSpotlightRow4DriverSideColor string = "Cabin.Light.Spotlight.Row4.DriverSide.Color"
	// FieldCabinLightSpotlightRow4DriverSideIntensity is the name for the Cabin.Light.Spotlight.Row4.DriverSide.Intensity column in the Clickhouse.
	FieldCabinLightSpotlightRow4DriverSideIntensity string = "Cabin.Light.Spotlight.Row4.DriverSide.Intensity"
	// FieldCabinLightSpotlightRow4DriverSideIsLightOn is the name for the Cabin.Light.Spotlight.Row4.DriverSide.IsLightOn column in the Clickhouse.
	FieldCabinLightSpotlightRow4DriverSideIsLightOn string = "Cabin.Light.Spotlight.Row4.DriverSide.IsLightOn"
	// FieldCabinLightSpotlightRow4PassengerSide is the name for the Cabin.Light.Spotlight.Row4.PassengerSide column in the Clickhouse.
	FieldCabinLightSpotlightRow4PassengerSide string = "Cabin.Light.Spotlight.Row4.PassengerSide"
	// FieldCabinLightSpotlightRow4PassengerSideColor is the name for the Cabin.Light.Spotlight.Row4.PassengerSide.Color column in the Clickhouse.
	FieldCabinLightSpotlightRow4PassengerSideColor string = "Cabin.Light.Spotlight.Row4.PassengerSide.Color"
	// FieldCabinLightSpotlightRow4PassengerSideIntensity is the name for the Cabin.Light.Spotlight.Row4.PassengerSide.Intensity column in the Clickhouse.
	FieldCabinLightSpotlightRow4PassengerSideIntensity string = "Cabin.Light.Spotlight.Row4.PassengerSide.Intensity"
	// FieldCabinLightSpotlightRow4PassengerSideIsLightOn is the name for the Cabin.Light.Spotlight.Row4.PassengerSide.IsLightOn column in the Clickhouse.
	FieldCabinLightSpotlightRow4PassengerSideIsLightOn string = "Cabin.Light.Spotlight.Row4.PassengerSide.IsLightOn"
	// FieldCabinPowerOptimizeLevel is the name for the Cabin.PowerOptimizeLevel column in the Clickhouse.
	FieldCabinPowerOptimizeLevel string = "Cabin.PowerOptimizeLevel"
	// FieldCabinRearShade is the name for the Cabin.RearShade column in the Clickhouse.
	FieldCabinRearShade string = "Cabin.RearShade"
	// FieldCabinRearShadeIsOpen is the name for the Cabin.RearShade.IsOpen column in the Clickhouse.
	FieldCabinRearShadeIsOpen string = "Cabin.RearShade.IsOpen"
	// FieldCabinRearShadePosition is the name for the Cabin.RearShade.Position column in the Clickhouse.
	FieldCabinRearShadePosition string = "Cabin.RearShade.Position"
	// FieldCabinRearShadeSwitch is the name for the Cabin.RearShade.Switch column in the Clickhouse.
	FieldCabinRearShadeSwitch string = "Cabin.RearShade.Switch"
	// FieldCabinRearviewMirrorDimmingLevel is the name for the Cabin.RearviewMirror_DimmingLevel column in the Clickhouse.
	FieldCabinRearviewMirrorDimmingLevel string = "Cabin.RearviewMirror_DimmingLevel"
	// FieldCabinSeat is the name for the Cabin.Seat column in the Clickhouse.
	FieldCabinSeat string = "Cabin.Seat"
	// FieldCabinSeatPosCount is the name for the Cabin.SeatPosCount column in the Clickhouse.
	FieldCabinSeatPosCount string = "Cabin.SeatPosCount"
	// FieldCabinSeatPosCountsize0 is the name for the Cabin.SeatPosCount.size0 column in the Clickhouse.
	FieldCabinSeatPosCountsize0 string = "Cabin.SeatPosCount.size0"
	// FieldCabinSeatRow1 is the name for the Cabin.Seat.Row1 column in the Clickhouse.
	FieldCabinSeatRow1 string = "Cabin.Seat.Row1"
	// FieldCabinSeatRow1DriverSide is the name for the Cabin.Seat.Row1.DriverSide column in the Clickhouse.
	FieldCabinSeatRow1DriverSide string = "Cabin.Seat.Row1.DriverSide"
	// FieldCabinSeatRow1DriverSideAirbagIsDeployed is the name for the Cabin.Seat.Row1.DriverSide.Airbag_IsDeployed column in the Clickhouse.
	FieldCabinSeatRow1DriverSideAirbagIsDeployed string = "Cabin.Seat.Row1.DriverSide.Airbag_IsDeployed"
	// FieldCabinSeatRow1DriverSideBackrest is the name for the Cabin.Seat.Row1.DriverSide.Backrest column in the Clickhouse.
	FieldCabinSeatRow1DriverSideBackrest string = "Cabin.Seat.Row1.DriverSide.Backrest"
	// FieldCabinSeatRow1DriverSideBackrestLumbar is the name for the Cabin.Seat.Row1.DriverSide.Backrest.Lumbar column in the Clickhouse.
	FieldCabinSeatRow1DriverSideBackrestLumbar string = "Cabin.Seat.Row1.DriverSide.Backrest.Lumbar"
	// FieldCabinSeatRow1DriverSideBackrestLumbarHeight is the name for the Cabin.Seat.Row1.DriverSide.Backrest.Lumbar.Height column in the Clickhouse.
	FieldCabinSeatRow1DriverSideBackrestLumbarHeight string = "Cabin.Seat.Row1.DriverSide.Backrest.Lumbar.Height"
	// FieldCabinSeatRow1DriverSideBackrestLumbarSupport is the name for the Cabin.Seat.Row1.DriverSide.Backrest.Lumbar.Support column in the Clickhouse.
	FieldCabinSeatRow1DriverSideBackrestLumbarSupport string = "Cabin.Seat.Row1.DriverSide.Backrest.Lumbar.Support"
	// FieldCabinSeatRow1DriverSideBackrestRecline is the name for the Cabin.Seat.Row1.DriverSide.Backrest.Recline column in the Clickhouse.
	FieldCabinSeatRow1DriverSideBackrestRecline string = "Cabin.Seat.Row1.DriverSide.Backrest.Recline"
	// FieldCabinSeatRow1DriverSideBackrestSideBolsterSupport is the name for the Cabin.Seat.Row1.DriverSide.Backrest.SideBolster_Support column in the Clickhouse.
	FieldCabinSeatRow1DriverSideBackrestSideBolsterSupport string = "Cabin.Seat.Row1.DriverSide.Backrest.SideBolster_Support"
	// FieldCabinSeatRow1DriverSideHeadrest is the name for the Cabin.Seat.Row1.DriverSide.Headrest column in the Clickhouse.
	FieldCabinSeatRow1DriverSideHeadrest string = "Cabin.Seat.Row1.DriverSide.Headrest"
	// FieldCabinSeatRow1DriverSideHeadrestAngle is the name for the Cabin.Seat.Row1.DriverSide.Headrest.Angle column in the Clickhouse.
	FieldCabinSeatRow1DriverSideHeadrestAngle string = "Cabin.Seat.Row1.DriverSide.Headrest.Angle"
	// FieldCabinSeatRow1DriverSideHeadrestHeight is the name for the Cabin.Seat.Row1.DriverSide.Headrest.Height column in the Clickhouse.
	FieldCabinSeatRow1DriverSideHeadrestHeight string = "Cabin.Seat.Row1.DriverSide.Headrest.Height"
	// FieldCabinSeatRow1DriverSideHeating is the name for the Cabin.Seat.Row1.DriverSide.Heating column in the Clickhouse.
	FieldCabinSeatRow1DriverSideHeating string = "Cabin.Seat.Row1.DriverSide.Heating"
	// FieldCabinSeatRow1DriverSideHeatingCooling is the name for the Cabin.Seat.Row1.DriverSide.HeatingCooling column in the Clickhouse.
	FieldCabinSeatRow1DriverSideHeatingCooling string = "Cabin.Seat.Row1.DriverSide.HeatingCooling"
	// FieldCabinSeatRow1DriverSideHeight is the name for the Cabin.Seat.Row1.DriverSide.Height column in the Clickhouse.
	FieldCabinSeatRow1DriverSideHeight string = "Cabin.Seat.Row1.DriverSide.Height"
	// FieldCabinSeatRow1DriverSideIsBelted is the name for the Cabin.Seat.Row1.DriverSide.IsBelted column in the Clickhouse.
	FieldCabinSeatRow1DriverSideIsBelted string = "Cabin.Seat.Row1.DriverSide.IsBelted"
	// FieldCabinSeatRow1DriverSideIsOccupied is the name for the Cabin.Seat.Row1.DriverSide.IsOccupied column in the Clickhouse.
	FieldCabinSeatRow1DriverSideIsOccupied string = "Cabin.Seat.Row1.DriverSide.IsOccupied"
	// FieldCabinSeatRow1DriverSideMassage is the name for the Cabin.Seat.Row1.DriverSide.Massage column in the Clickhouse.
	FieldCabinSeatRow1DriverSideMassage string = "Cabin.Seat.Row1.DriverSide.Massage"
	// FieldCabinSeatRow1DriverSideOccupantIdentifier is the name for the Cabin.Seat.Row1.DriverSide.Occupant_Identifier column in the Clickhouse.
	FieldCabinSeatRow1DriverSideOccupantIdentifier string = "Cabin.Seat.Row1.DriverSide.Occupant_Identifier"
	// FieldCabinSeatRow1DriverSideOccupantIdentifierIssuer is the name for the Cabin.Seat.Row1.DriverSide.Occupant_Identifier.Issuer column in the Clickhouse.
	FieldCabinSeatRow1DriverSideOccupantIdentifierIssuer string = "Cabin.Seat.Row1.DriverSide.Occupant_Identifier.Issuer"
	// FieldCabinSeatRow1DriverSideOccupantIdentifierSubject is the name for the Cabin.Seat.Row1.DriverSide.Occupant_Identifier.Subject column in the Clickhouse.
	FieldCabinSeatRow1DriverSideOccupantIdentifierSubject string = "Cabin.Seat.Row1.DriverSide.Occupant_Identifier.Subject"
	// FieldCabinSeatRow1DriverSidePosition is the name for the Cabin.Seat.Row1.DriverSide.Position column in the Clickhouse.
	FieldCabinSeatRow1DriverSidePosition string = "Cabin.Seat.Row1.DriverSide.Position"
	// FieldCabinSeatRow1DriverSideSeatBeltHeight is the name for the Cabin.Seat.Row1.DriverSide.SeatBeltHeight column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSeatBeltHeight string = "Cabin.Seat.Row1.DriverSide.SeatBeltHeight"
	// FieldCabinSeatRow1DriverSideSeatingLength is the name for the Cabin.Seat.Row1.DriverSide.Seating_Length column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSeatingLength string = "Cabin.Seat.Row1.DriverSide.Seating_Length"
	// FieldCabinSeatRow1DriverSideSwitch is the name for the Cabin.Seat.Row1.DriverSide.Switch column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitch string = "Cabin.Seat.Row1.DriverSide.Switch"
	// FieldCabinSeatRow1DriverSideSwitchBackrest is the name for the Cabin.Seat.Row1.DriverSide.Switch.Backrest column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchBackrest string = "Cabin.Seat.Row1.DriverSide.Switch.Backrest"
	// FieldCabinSeatRow1DriverSideSwitchBackrestIsReclineBackwardEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.Backrest.IsReclineBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchBackrestIsReclineBackwardEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.Backrest.IsReclineBackwardEngaged"
	// FieldCabinSeatRow1DriverSideSwitchBackrestIsReclineForwardEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.Backrest.IsReclineForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchBackrestIsReclineForwardEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.Backrest.IsReclineForwardEngaged"
	// FieldCabinSeatRow1DriverSideSwitchBackrestLumbar is the name for the Cabin.Seat.Row1.DriverSide.Switch.Backrest.Lumbar column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchBackrestLumbar string = "Cabin.Seat.Row1.DriverSide.Switch.Backrest.Lumbar"
	// FieldCabinSeatRow1DriverSideSwitchBackrestLumbarIsDownEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.Backrest.Lumbar.IsDownEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchBackrestLumbarIsDownEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.Backrest.Lumbar.IsDownEngaged"
	// FieldCabinSeatRow1DriverSideSwitchBackrestLumbarIsLessSupportEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.Backrest.Lumbar.IsLessSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchBackrestLumbarIsLessSupportEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.Backrest.Lumbar.IsLessSupportEngaged"
	// FieldCabinSeatRow1DriverSideSwitchBackrestLumbarIsMoreSupportEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.Backrest.Lumbar.IsMoreSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchBackrestLumbarIsMoreSupportEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.Backrest.Lumbar.IsMoreSupportEngaged"
	// FieldCabinSeatRow1DriverSideSwitchBackrestLumbarIsUpEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.Backrest.Lumbar.IsUpEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchBackrestLumbarIsUpEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.Backrest.Lumbar.IsUpEngaged"
	// FieldCabinSeatRow1DriverSideSwitchBackrestSideBolster is the name for the Cabin.Seat.Row1.DriverSide.Switch.Backrest.SideBolster column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchBackrestSideBolster string = "Cabin.Seat.Row1.DriverSide.Switch.Backrest.SideBolster"
	// FieldCabinSeatRow1DriverSideSwitchBackrestSideBolsterIsLessSupportEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.Backrest.SideBolster.IsLessSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchBackrestSideBolsterIsLessSupportEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.Backrest.SideBolster.IsLessSupportEngaged"
	// FieldCabinSeatRow1DriverSideSwitchBackrestSideBolsterIsMoreSupportEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.Backrest.SideBolster.IsMoreSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchBackrestSideBolsterIsMoreSupportEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.Backrest.SideBolster.IsMoreSupportEngaged"
	// FieldCabinSeatRow1DriverSideSwitchHeadrest is the name for the Cabin.Seat.Row1.DriverSide.Switch.Headrest column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchHeadrest string = "Cabin.Seat.Row1.DriverSide.Switch.Headrest"
	// FieldCabinSeatRow1DriverSideSwitchHeadrestIsBackwardEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.Headrest.IsBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchHeadrestIsBackwardEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.Headrest.IsBackwardEngaged"
	// FieldCabinSeatRow1DriverSideSwitchHeadrestIsDownEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.Headrest.IsDownEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchHeadrestIsDownEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.Headrest.IsDownEngaged"
	// FieldCabinSeatRow1DriverSideSwitchHeadrestIsForwardEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.Headrest.IsForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchHeadrestIsForwardEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.Headrest.IsForwardEngaged"
	// FieldCabinSeatRow1DriverSideSwitchHeadrestIsUpEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.Headrest.IsUpEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchHeadrestIsUpEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.Headrest.IsUpEngaged"
	// FieldCabinSeatRow1DriverSideSwitchIsBackwardEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.IsBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchIsBackwardEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.IsBackwardEngaged"
	// FieldCabinSeatRow1DriverSideSwitchIsCoolerEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.IsCoolerEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchIsCoolerEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.IsCoolerEngaged"
	// FieldCabinSeatRow1DriverSideSwitchIsDownEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.IsDownEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchIsDownEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.IsDownEngaged"
	// FieldCabinSeatRow1DriverSideSwitchIsForwardEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.IsForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchIsForwardEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.IsForwardEngaged"
	// FieldCabinSeatRow1DriverSideSwitchIsTiltBackwardEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.IsTiltBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchIsTiltBackwardEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.IsTiltBackwardEngaged"
	// FieldCabinSeatRow1DriverSideSwitchIsTiltForwardEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.IsTiltForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchIsTiltForwardEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.IsTiltForwardEngaged"
	// FieldCabinSeatRow1DriverSideSwitchIsUpEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.IsUpEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchIsUpEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.IsUpEngaged"
	// FieldCabinSeatRow1DriverSideSwitchIsWarmerEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.IsWarmerEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchIsWarmerEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.IsWarmerEngaged"
	// FieldCabinSeatRow1DriverSideSwitchMassage is the name for the Cabin.Seat.Row1.DriverSide.Switch.Massage column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchMassage string = "Cabin.Seat.Row1.DriverSide.Switch.Massage"
	// FieldCabinSeatRow1DriverSideSwitchMassageIsDecreaseEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.Massage.IsDecreaseEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchMassageIsDecreaseEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.Massage.IsDecreaseEngaged"
	// FieldCabinSeatRow1DriverSideSwitchMassageIsIncreaseEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.Massage.IsIncreaseEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchMassageIsIncreaseEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.Massage.IsIncreaseEngaged"
	// FieldCabinSeatRow1DriverSideSwitchSeating is the name for the Cabin.Seat.Row1.DriverSide.Switch.Seating column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchSeating string = "Cabin.Seat.Row1.DriverSide.Switch.Seating"
	// FieldCabinSeatRow1DriverSideSwitchSeatingIsBackwardEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.Seating.IsBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchSeatingIsBackwardEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.Seating.IsBackwardEngaged"
	// FieldCabinSeatRow1DriverSideSwitchSeatingIsForwardEngaged is the name for the Cabin.Seat.Row1.DriverSide.Switch.Seating.IsForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1DriverSideSwitchSeatingIsForwardEngaged string = "Cabin.Seat.Row1.DriverSide.Switch.Seating.IsForwardEngaged"
	// FieldCabinSeatRow1DriverSideTilt is the name for the Cabin.Seat.Row1.DriverSide.Tilt column in the Clickhouse.
	FieldCabinSeatRow1DriverSideTilt string = "Cabin.Seat.Row1.DriverSide.Tilt"
	// FieldCabinSeatRow1Middle is the name for the Cabin.Seat.Row1.Middle column in the Clickhouse.
	FieldCabinSeatRow1Middle string = "Cabin.Seat.Row1.Middle"
	// FieldCabinSeatRow1MiddleAirbagIsDeployed is the name for the Cabin.Seat.Row1.Middle.Airbag_IsDeployed column in the Clickhouse.
	FieldCabinSeatRow1MiddleAirbagIsDeployed string = "Cabin.Seat.Row1.Middle.Airbag_IsDeployed"
	// FieldCabinSeatRow1MiddleBackrest is the name for the Cabin.Seat.Row1.Middle.Backrest column in the Clickhouse.
	FieldCabinSeatRow1MiddleBackrest string = "Cabin.Seat.Row1.Middle.Backrest"
	// FieldCabinSeatRow1MiddleBackrestLumbar is the name for the Cabin.Seat.Row1.Middle.Backrest.Lumbar column in the Clickhouse.
	FieldCabinSeatRow1MiddleBackrestLumbar string = "Cabin.Seat.Row1.Middle.Backrest.Lumbar"
	// FieldCabinSeatRow1MiddleBackrestLumbarHeight is the name for the Cabin.Seat.Row1.Middle.Backrest.Lumbar.Height column in the Clickhouse.
	FieldCabinSeatRow1MiddleBackrestLumbarHeight string = "Cabin.Seat.Row1.Middle.Backrest.Lumbar.Height"
	// FieldCabinSeatRow1MiddleBackrestLumbarSupport is the name for the Cabin.Seat.Row1.Middle.Backrest.Lumbar.Support column in the Clickhouse.
	FieldCabinSeatRow1MiddleBackrestLumbarSupport string = "Cabin.Seat.Row1.Middle.Backrest.Lumbar.Support"
	// FieldCabinSeatRow1MiddleBackrestRecline is the name for the Cabin.Seat.Row1.Middle.Backrest.Recline column in the Clickhouse.
	FieldCabinSeatRow1MiddleBackrestRecline string = "Cabin.Seat.Row1.Middle.Backrest.Recline"
	// FieldCabinSeatRow1MiddleBackrestSideBolsterSupport is the name for the Cabin.Seat.Row1.Middle.Backrest.SideBolster_Support column in the Clickhouse.
	FieldCabinSeatRow1MiddleBackrestSideBolsterSupport string = "Cabin.Seat.Row1.Middle.Backrest.SideBolster_Support"
	// FieldCabinSeatRow1MiddleHeadrest is the name for the Cabin.Seat.Row1.Middle.Headrest column in the Clickhouse.
	FieldCabinSeatRow1MiddleHeadrest string = "Cabin.Seat.Row1.Middle.Headrest"
	// FieldCabinSeatRow1MiddleHeadrestAngle is the name for the Cabin.Seat.Row1.Middle.Headrest.Angle column in the Clickhouse.
	FieldCabinSeatRow1MiddleHeadrestAngle string = "Cabin.Seat.Row1.Middle.Headrest.Angle"
	// FieldCabinSeatRow1MiddleHeadrestHeight is the name for the Cabin.Seat.Row1.Middle.Headrest.Height column in the Clickhouse.
	FieldCabinSeatRow1MiddleHeadrestHeight string = "Cabin.Seat.Row1.Middle.Headrest.Height"
	// FieldCabinSeatRow1MiddleHeating is the name for the Cabin.Seat.Row1.Middle.Heating column in the Clickhouse.
	FieldCabinSeatRow1MiddleHeating string = "Cabin.Seat.Row1.Middle.Heating"
	// FieldCabinSeatRow1MiddleHeatingCooling is the name for the Cabin.Seat.Row1.Middle.HeatingCooling column in the Clickhouse.
	FieldCabinSeatRow1MiddleHeatingCooling string = "Cabin.Seat.Row1.Middle.HeatingCooling"
	// FieldCabinSeatRow1MiddleHeight is the name for the Cabin.Seat.Row1.Middle.Height column in the Clickhouse.
	FieldCabinSeatRow1MiddleHeight string = "Cabin.Seat.Row1.Middle.Height"
	// FieldCabinSeatRow1MiddleIsBelted is the name for the Cabin.Seat.Row1.Middle.IsBelted column in the Clickhouse.
	FieldCabinSeatRow1MiddleIsBelted string = "Cabin.Seat.Row1.Middle.IsBelted"
	// FieldCabinSeatRow1MiddleIsOccupied is the name for the Cabin.Seat.Row1.Middle.IsOccupied column in the Clickhouse.
	FieldCabinSeatRow1MiddleIsOccupied string = "Cabin.Seat.Row1.Middle.IsOccupied"
	// FieldCabinSeatRow1MiddleMassage is the name for the Cabin.Seat.Row1.Middle.Massage column in the Clickhouse.
	FieldCabinSeatRow1MiddleMassage string = "Cabin.Seat.Row1.Middle.Massage"
	// FieldCabinSeatRow1MiddleOccupantIdentifier is the name for the Cabin.Seat.Row1.Middle.Occupant_Identifier column in the Clickhouse.
	FieldCabinSeatRow1MiddleOccupantIdentifier string = "Cabin.Seat.Row1.Middle.Occupant_Identifier"
	// FieldCabinSeatRow1MiddleOccupantIdentifierIssuer is the name for the Cabin.Seat.Row1.Middle.Occupant_Identifier.Issuer column in the Clickhouse.
	FieldCabinSeatRow1MiddleOccupantIdentifierIssuer string = "Cabin.Seat.Row1.Middle.Occupant_Identifier.Issuer"
	// FieldCabinSeatRow1MiddleOccupantIdentifierSubject is the name for the Cabin.Seat.Row1.Middle.Occupant_Identifier.Subject column in the Clickhouse.
	FieldCabinSeatRow1MiddleOccupantIdentifierSubject string = "Cabin.Seat.Row1.Middle.Occupant_Identifier.Subject"
	// FieldCabinSeatRow1MiddlePosition is the name for the Cabin.Seat.Row1.Middle.Position column in the Clickhouse.
	FieldCabinSeatRow1MiddlePosition string = "Cabin.Seat.Row1.Middle.Position"
	// FieldCabinSeatRow1MiddleSeatBeltHeight is the name for the Cabin.Seat.Row1.Middle.SeatBeltHeight column in the Clickhouse.
	FieldCabinSeatRow1MiddleSeatBeltHeight string = "Cabin.Seat.Row1.Middle.SeatBeltHeight"
	// FieldCabinSeatRow1MiddleSeatingLength is the name for the Cabin.Seat.Row1.Middle.Seating_Length column in the Clickhouse.
	FieldCabinSeatRow1MiddleSeatingLength string = "Cabin.Seat.Row1.Middle.Seating_Length"
	// FieldCabinSeatRow1MiddleSwitch is the name for the Cabin.Seat.Row1.Middle.Switch column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitch string = "Cabin.Seat.Row1.Middle.Switch"
	// FieldCabinSeatRow1MiddleSwitchBackrest is the name for the Cabin.Seat.Row1.Middle.Switch.Backrest column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchBackrest string = "Cabin.Seat.Row1.Middle.Switch.Backrest"
	// FieldCabinSeatRow1MiddleSwitchBackrestIsReclineBackwardEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.Backrest.IsReclineBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchBackrestIsReclineBackwardEngaged string = "Cabin.Seat.Row1.Middle.Switch.Backrest.IsReclineBackwardEngaged"
	// FieldCabinSeatRow1MiddleSwitchBackrestIsReclineForwardEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.Backrest.IsReclineForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchBackrestIsReclineForwardEngaged string = "Cabin.Seat.Row1.Middle.Switch.Backrest.IsReclineForwardEngaged"
	// FieldCabinSeatRow1MiddleSwitchBackrestLumbar is the name for the Cabin.Seat.Row1.Middle.Switch.Backrest.Lumbar column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchBackrestLumbar string = "Cabin.Seat.Row1.Middle.Switch.Backrest.Lumbar"
	// FieldCabinSeatRow1MiddleSwitchBackrestLumbarIsDownEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.Backrest.Lumbar.IsDownEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchBackrestLumbarIsDownEngaged string = "Cabin.Seat.Row1.Middle.Switch.Backrest.Lumbar.IsDownEngaged"
	// FieldCabinSeatRow1MiddleSwitchBackrestLumbarIsLessSupportEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.Backrest.Lumbar.IsLessSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchBackrestLumbarIsLessSupportEngaged string = "Cabin.Seat.Row1.Middle.Switch.Backrest.Lumbar.IsLessSupportEngaged"
	// FieldCabinSeatRow1MiddleSwitchBackrestLumbarIsMoreSupportEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.Backrest.Lumbar.IsMoreSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchBackrestLumbarIsMoreSupportEngaged string = "Cabin.Seat.Row1.Middle.Switch.Backrest.Lumbar.IsMoreSupportEngaged"
	// FieldCabinSeatRow1MiddleSwitchBackrestLumbarIsUpEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.Backrest.Lumbar.IsUpEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchBackrestLumbarIsUpEngaged string = "Cabin.Seat.Row1.Middle.Switch.Backrest.Lumbar.IsUpEngaged"
	// FieldCabinSeatRow1MiddleSwitchBackrestSideBolster is the name for the Cabin.Seat.Row1.Middle.Switch.Backrest.SideBolster column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchBackrestSideBolster string = "Cabin.Seat.Row1.Middle.Switch.Backrest.SideBolster"
	// FieldCabinSeatRow1MiddleSwitchBackrestSideBolsterIsLessSupportEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.Backrest.SideBolster.IsLessSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchBackrestSideBolsterIsLessSupportEngaged string = "Cabin.Seat.Row1.Middle.Switch.Backrest.SideBolster.IsLessSupportEngaged"
	// FieldCabinSeatRow1MiddleSwitchBackrestSideBolsterIsMoreSupportEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.Backrest.SideBolster.IsMoreSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchBackrestSideBolsterIsMoreSupportEngaged string = "Cabin.Seat.Row1.Middle.Switch.Backrest.SideBolster.IsMoreSupportEngaged"
	// FieldCabinSeatRow1MiddleSwitchHeadrest is the name for the Cabin.Seat.Row1.Middle.Switch.Headrest column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchHeadrest string = "Cabin.Seat.Row1.Middle.Switch.Headrest"
	// FieldCabinSeatRow1MiddleSwitchHeadrestIsBackwardEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.Headrest.IsBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchHeadrestIsBackwardEngaged string = "Cabin.Seat.Row1.Middle.Switch.Headrest.IsBackwardEngaged"
	// FieldCabinSeatRow1MiddleSwitchHeadrestIsDownEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.Headrest.IsDownEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchHeadrestIsDownEngaged string = "Cabin.Seat.Row1.Middle.Switch.Headrest.IsDownEngaged"
	// FieldCabinSeatRow1MiddleSwitchHeadrestIsForwardEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.Headrest.IsForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchHeadrestIsForwardEngaged string = "Cabin.Seat.Row1.Middle.Switch.Headrest.IsForwardEngaged"
	// FieldCabinSeatRow1MiddleSwitchHeadrestIsUpEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.Headrest.IsUpEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchHeadrestIsUpEngaged string = "Cabin.Seat.Row1.Middle.Switch.Headrest.IsUpEngaged"
	// FieldCabinSeatRow1MiddleSwitchIsBackwardEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.IsBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchIsBackwardEngaged string = "Cabin.Seat.Row1.Middle.Switch.IsBackwardEngaged"
	// FieldCabinSeatRow1MiddleSwitchIsCoolerEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.IsCoolerEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchIsCoolerEngaged string = "Cabin.Seat.Row1.Middle.Switch.IsCoolerEngaged"
	// FieldCabinSeatRow1MiddleSwitchIsDownEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.IsDownEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchIsDownEngaged string = "Cabin.Seat.Row1.Middle.Switch.IsDownEngaged"
	// FieldCabinSeatRow1MiddleSwitchIsForwardEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.IsForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchIsForwardEngaged string = "Cabin.Seat.Row1.Middle.Switch.IsForwardEngaged"
	// FieldCabinSeatRow1MiddleSwitchIsTiltBackwardEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.IsTiltBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchIsTiltBackwardEngaged string = "Cabin.Seat.Row1.Middle.Switch.IsTiltBackwardEngaged"
	// FieldCabinSeatRow1MiddleSwitchIsTiltForwardEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.IsTiltForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchIsTiltForwardEngaged string = "Cabin.Seat.Row1.Middle.Switch.IsTiltForwardEngaged"
	// FieldCabinSeatRow1MiddleSwitchIsUpEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.IsUpEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchIsUpEngaged string = "Cabin.Seat.Row1.Middle.Switch.IsUpEngaged"
	// FieldCabinSeatRow1MiddleSwitchIsWarmerEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.IsWarmerEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchIsWarmerEngaged string = "Cabin.Seat.Row1.Middle.Switch.IsWarmerEngaged"
	// FieldCabinSeatRow1MiddleSwitchMassage is the name for the Cabin.Seat.Row1.Middle.Switch.Massage column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchMassage string = "Cabin.Seat.Row1.Middle.Switch.Massage"
	// FieldCabinSeatRow1MiddleSwitchMassageIsDecreaseEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.Massage.IsDecreaseEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchMassageIsDecreaseEngaged string = "Cabin.Seat.Row1.Middle.Switch.Massage.IsDecreaseEngaged"
	// FieldCabinSeatRow1MiddleSwitchMassageIsIncreaseEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.Massage.IsIncreaseEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchMassageIsIncreaseEngaged string = "Cabin.Seat.Row1.Middle.Switch.Massage.IsIncreaseEngaged"
	// FieldCabinSeatRow1MiddleSwitchSeating is the name for the Cabin.Seat.Row1.Middle.Switch.Seating column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchSeating string = "Cabin.Seat.Row1.Middle.Switch.Seating"
	// FieldCabinSeatRow1MiddleSwitchSeatingIsBackwardEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.Seating.IsBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchSeatingIsBackwardEngaged string = "Cabin.Seat.Row1.Middle.Switch.Seating.IsBackwardEngaged"
	// FieldCabinSeatRow1MiddleSwitchSeatingIsForwardEngaged is the name for the Cabin.Seat.Row1.Middle.Switch.Seating.IsForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1MiddleSwitchSeatingIsForwardEngaged string = "Cabin.Seat.Row1.Middle.Switch.Seating.IsForwardEngaged"
	// FieldCabinSeatRow1MiddleTilt is the name for the Cabin.Seat.Row1.Middle.Tilt column in the Clickhouse.
	FieldCabinSeatRow1MiddleTilt string = "Cabin.Seat.Row1.Middle.Tilt"
	// FieldCabinSeatRow1PassengerSide is the name for the Cabin.Seat.Row1.PassengerSide column in the Clickhouse.
	FieldCabinSeatRow1PassengerSide string = "Cabin.Seat.Row1.PassengerSide"
	// FieldCabinSeatRow1PassengerSideAirbagIsDeployed is the name for the Cabin.Seat.Row1.PassengerSide.Airbag_IsDeployed column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideAirbagIsDeployed string = "Cabin.Seat.Row1.PassengerSide.Airbag_IsDeployed"
	// FieldCabinSeatRow1PassengerSideBackrest is the name for the Cabin.Seat.Row1.PassengerSide.Backrest column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideBackrest string = "Cabin.Seat.Row1.PassengerSide.Backrest"
	// FieldCabinSeatRow1PassengerSideBackrestLumbar is the name for the Cabin.Seat.Row1.PassengerSide.Backrest.Lumbar column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideBackrestLumbar string = "Cabin.Seat.Row1.PassengerSide.Backrest.Lumbar"
	// FieldCabinSeatRow1PassengerSideBackrestLumbarHeight is the name for the Cabin.Seat.Row1.PassengerSide.Backrest.Lumbar.Height column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideBackrestLumbarHeight string = "Cabin.Seat.Row1.PassengerSide.Backrest.Lumbar.Height"
	// FieldCabinSeatRow1PassengerSideBackrestLumbarSupport is the name for the Cabin.Seat.Row1.PassengerSide.Backrest.Lumbar.Support column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideBackrestLumbarSupport string = "Cabin.Seat.Row1.PassengerSide.Backrest.Lumbar.Support"
	// FieldCabinSeatRow1PassengerSideBackrestRecline is the name for the Cabin.Seat.Row1.PassengerSide.Backrest.Recline column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideBackrestRecline string = "Cabin.Seat.Row1.PassengerSide.Backrest.Recline"
	// FieldCabinSeatRow1PassengerSideBackrestSideBolsterSupport is the name for the Cabin.Seat.Row1.PassengerSide.Backrest.SideBolster_Support column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideBackrestSideBolsterSupport string = "Cabin.Seat.Row1.PassengerSide.Backrest.SideBolster_Support"
	// FieldCabinSeatRow1PassengerSideHeadrest is the name for the Cabin.Seat.Row1.PassengerSide.Headrest column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideHeadrest string = "Cabin.Seat.Row1.PassengerSide.Headrest"
	// FieldCabinSeatRow1PassengerSideHeadrestAngle is the name for the Cabin.Seat.Row1.PassengerSide.Headrest.Angle column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideHeadrestAngle string = "Cabin.Seat.Row1.PassengerSide.Headrest.Angle"
	// FieldCabinSeatRow1PassengerSideHeadrestHeight is the name for the Cabin.Seat.Row1.PassengerSide.Headrest.Height column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideHeadrestHeight string = "Cabin.Seat.Row1.PassengerSide.Headrest.Height"
	// FieldCabinSeatRow1PassengerSideHeating is the name for the Cabin.Seat.Row1.PassengerSide.Heating column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideHeating string = "Cabin.Seat.Row1.PassengerSide.Heating"
	// FieldCabinSeatRow1PassengerSideHeatingCooling is the name for the Cabin.Seat.Row1.PassengerSide.HeatingCooling column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideHeatingCooling string = "Cabin.Seat.Row1.PassengerSide.HeatingCooling"
	// FieldCabinSeatRow1PassengerSideHeight is the name for the Cabin.Seat.Row1.PassengerSide.Height column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideHeight string = "Cabin.Seat.Row1.PassengerSide.Height"
	// FieldCabinSeatRow1PassengerSideIsBelted is the name for the Cabin.Seat.Row1.PassengerSide.IsBelted column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideIsBelted string = "Cabin.Seat.Row1.PassengerSide.IsBelted"
	// FieldCabinSeatRow1PassengerSideIsOccupied is the name for the Cabin.Seat.Row1.PassengerSide.IsOccupied column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideIsOccupied string = "Cabin.Seat.Row1.PassengerSide.IsOccupied"
	// FieldCabinSeatRow1PassengerSideMassage is the name for the Cabin.Seat.Row1.PassengerSide.Massage column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideMassage string = "Cabin.Seat.Row1.PassengerSide.Massage"
	// FieldCabinSeatRow1PassengerSideOccupantIdentifier is the name for the Cabin.Seat.Row1.PassengerSide.Occupant_Identifier column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideOccupantIdentifier string = "Cabin.Seat.Row1.PassengerSide.Occupant_Identifier"
	// FieldCabinSeatRow1PassengerSideOccupantIdentifierIssuer is the name for the Cabin.Seat.Row1.PassengerSide.Occupant_Identifier.Issuer column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideOccupantIdentifierIssuer string = "Cabin.Seat.Row1.PassengerSide.Occupant_Identifier.Issuer"
	// FieldCabinSeatRow1PassengerSideOccupantIdentifierSubject is the name for the Cabin.Seat.Row1.PassengerSide.Occupant_Identifier.Subject column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideOccupantIdentifierSubject string = "Cabin.Seat.Row1.PassengerSide.Occupant_Identifier.Subject"
	// FieldCabinSeatRow1PassengerSidePosition is the name for the Cabin.Seat.Row1.PassengerSide.Position column in the Clickhouse.
	FieldCabinSeatRow1PassengerSidePosition string = "Cabin.Seat.Row1.PassengerSide.Position"
	// FieldCabinSeatRow1PassengerSideSeatBeltHeight is the name for the Cabin.Seat.Row1.PassengerSide.SeatBeltHeight column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSeatBeltHeight string = "Cabin.Seat.Row1.PassengerSide.SeatBeltHeight"
	// FieldCabinSeatRow1PassengerSideSeatingLength is the name for the Cabin.Seat.Row1.PassengerSide.Seating_Length column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSeatingLength string = "Cabin.Seat.Row1.PassengerSide.Seating_Length"
	// FieldCabinSeatRow1PassengerSideSwitch is the name for the Cabin.Seat.Row1.PassengerSide.Switch column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitch string = "Cabin.Seat.Row1.PassengerSide.Switch"
	// FieldCabinSeatRow1PassengerSideSwitchBackrest is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Backrest column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchBackrest string = "Cabin.Seat.Row1.PassengerSide.Switch.Backrest"
	// FieldCabinSeatRow1PassengerSideSwitchBackrestIsReclineBackwardEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Backrest.IsReclineBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchBackrestIsReclineBackwardEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.Backrest.IsReclineBackwardEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchBackrestIsReclineForwardEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Backrest.IsReclineForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchBackrestIsReclineForwardEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.Backrest.IsReclineForwardEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchBackrestLumbar is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Backrest.Lumbar column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchBackrestLumbar string = "Cabin.Seat.Row1.PassengerSide.Switch.Backrest.Lumbar"
	// FieldCabinSeatRow1PassengerSideSwitchBackrestLumbarIsDownEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Backrest.Lumbar.IsDownEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchBackrestLumbarIsDownEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.Backrest.Lumbar.IsDownEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchBackrestLumbarIsLessSupportEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Backrest.Lumbar.IsLessSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchBackrestLumbarIsLessSupportEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.Backrest.Lumbar.IsLessSupportEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchBackrestLumbarIsMoreSupportEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Backrest.Lumbar.IsMoreSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchBackrestLumbarIsMoreSupportEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.Backrest.Lumbar.IsMoreSupportEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchBackrestLumbarIsUpEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Backrest.Lumbar.IsUpEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchBackrestLumbarIsUpEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.Backrest.Lumbar.IsUpEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchBackrestSideBolster is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Backrest.SideBolster column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchBackrestSideBolster string = "Cabin.Seat.Row1.PassengerSide.Switch.Backrest.SideBolster"
	// FieldCabinSeatRow1PassengerSideSwitchBackrestSideBolsterIsLessSupportEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Backrest.SideBolster.IsLessSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchBackrestSideBolsterIsLessSupportEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.Backrest.SideBolster.IsLessSupportEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchBackrestSideBolsterIsMoreSupportEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Backrest.SideBolster.IsMoreSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchBackrestSideBolsterIsMoreSupportEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.Backrest.SideBolster.IsMoreSupportEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchHeadrest is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Headrest column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchHeadrest string = "Cabin.Seat.Row1.PassengerSide.Switch.Headrest"
	// FieldCabinSeatRow1PassengerSideSwitchHeadrestIsBackwardEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Headrest.IsBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchHeadrestIsBackwardEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.Headrest.IsBackwardEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchHeadrestIsDownEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Headrest.IsDownEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchHeadrestIsDownEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.Headrest.IsDownEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchHeadrestIsForwardEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Headrest.IsForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchHeadrestIsForwardEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.Headrest.IsForwardEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchHeadrestIsUpEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Headrest.IsUpEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchHeadrestIsUpEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.Headrest.IsUpEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchIsBackwardEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.IsBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchIsBackwardEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.IsBackwardEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchIsCoolerEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.IsCoolerEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchIsCoolerEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.IsCoolerEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchIsDownEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.IsDownEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchIsDownEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.IsDownEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchIsForwardEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.IsForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchIsForwardEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.IsForwardEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchIsTiltBackwardEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.IsTiltBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchIsTiltBackwardEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.IsTiltBackwardEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchIsTiltForwardEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.IsTiltForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchIsTiltForwardEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.IsTiltForwardEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchIsUpEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.IsUpEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchIsUpEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.IsUpEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchIsWarmerEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.IsWarmerEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchIsWarmerEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.IsWarmerEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchMassage is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Massage column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchMassage string = "Cabin.Seat.Row1.PassengerSide.Switch.Massage"
	// FieldCabinSeatRow1PassengerSideSwitchMassageIsDecreaseEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Massage.IsDecreaseEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchMassageIsDecreaseEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.Massage.IsDecreaseEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchMassageIsIncreaseEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Massage.IsIncreaseEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchMassageIsIncreaseEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.Massage.IsIncreaseEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchSeating is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Seating column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchSeating string = "Cabin.Seat.Row1.PassengerSide.Switch.Seating"
	// FieldCabinSeatRow1PassengerSideSwitchSeatingIsBackwardEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Seating.IsBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchSeatingIsBackwardEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.Seating.IsBackwardEngaged"
	// FieldCabinSeatRow1PassengerSideSwitchSeatingIsForwardEngaged is the name for the Cabin.Seat.Row1.PassengerSide.Switch.Seating.IsForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideSwitchSeatingIsForwardEngaged string = "Cabin.Seat.Row1.PassengerSide.Switch.Seating.IsForwardEngaged"
	// FieldCabinSeatRow1PassengerSideTilt is the name for the Cabin.Seat.Row1.PassengerSide.Tilt column in the Clickhouse.
	FieldCabinSeatRow1PassengerSideTilt string = "Cabin.Seat.Row1.PassengerSide.Tilt"
	// FieldCabinSeatRow2 is the name for the Cabin.Seat.Row2 column in the Clickhouse.
	FieldCabinSeatRow2 string = "Cabin.Seat.Row2"
	// FieldCabinSeatRow2DriverSide is the name for the Cabin.Seat.Row2.DriverSide column in the Clickhouse.
	FieldCabinSeatRow2DriverSide string = "Cabin.Seat.Row2.DriverSide"
	// FieldCabinSeatRow2DriverSideAirbagIsDeployed is the name for the Cabin.Seat.Row2.DriverSide.Airbag_IsDeployed column in the Clickhouse.
	FieldCabinSeatRow2DriverSideAirbagIsDeployed string = "Cabin.Seat.Row2.DriverSide.Airbag_IsDeployed"
	// FieldCabinSeatRow2DriverSideBackrest is the name for the Cabin.Seat.Row2.DriverSide.Backrest column in the Clickhouse.
	FieldCabinSeatRow2DriverSideBackrest string = "Cabin.Seat.Row2.DriverSide.Backrest"
	// FieldCabinSeatRow2DriverSideBackrestLumbar is the name for the Cabin.Seat.Row2.DriverSide.Backrest.Lumbar column in the Clickhouse.
	FieldCabinSeatRow2DriverSideBackrestLumbar string = "Cabin.Seat.Row2.DriverSide.Backrest.Lumbar"
	// FieldCabinSeatRow2DriverSideBackrestLumbarHeight is the name for the Cabin.Seat.Row2.DriverSide.Backrest.Lumbar.Height column in the Clickhouse.
	FieldCabinSeatRow2DriverSideBackrestLumbarHeight string = "Cabin.Seat.Row2.DriverSide.Backrest.Lumbar.Height"
	// FieldCabinSeatRow2DriverSideBackrestLumbarSupport is the name for the Cabin.Seat.Row2.DriverSide.Backrest.Lumbar.Support column in the Clickhouse.
	FieldCabinSeatRow2DriverSideBackrestLumbarSupport string = "Cabin.Seat.Row2.DriverSide.Backrest.Lumbar.Support"
	// FieldCabinSeatRow2DriverSideBackrestRecline is the name for the Cabin.Seat.Row2.DriverSide.Backrest.Recline column in the Clickhouse.
	FieldCabinSeatRow2DriverSideBackrestRecline string = "Cabin.Seat.Row2.DriverSide.Backrest.Recline"
	// FieldCabinSeatRow2DriverSideBackrestSideBolsterSupport is the name for the Cabin.Seat.Row2.DriverSide.Backrest.SideBolster_Support column in the Clickhouse.
	FieldCabinSeatRow2DriverSideBackrestSideBolsterSupport string = "Cabin.Seat.Row2.DriverSide.Backrest.SideBolster_Support"
	// FieldCabinSeatRow2DriverSideHeadrest is the name for the Cabin.Seat.Row2.DriverSide.Headrest column in the Clickhouse.
	FieldCabinSeatRow2DriverSideHeadrest string = "Cabin.Seat.Row2.DriverSide.Headrest"
	// FieldCabinSeatRow2DriverSideHeadrestAngle is the name for the Cabin.Seat.Row2.DriverSide.Headrest.Angle column in the Clickhouse.
	FieldCabinSeatRow2DriverSideHeadrestAngle string = "Cabin.Seat.Row2.DriverSide.Headrest.Angle"
	// FieldCabinSeatRow2DriverSideHeadrestHeight is the name for the Cabin.Seat.Row2.DriverSide.Headrest.Height column in the Clickhouse.
	FieldCabinSeatRow2DriverSideHeadrestHeight string = "Cabin.Seat.Row2.DriverSide.Headrest.Height"
	// FieldCabinSeatRow2DriverSideHeating is the name for the Cabin.Seat.Row2.DriverSide.Heating column in the Clickhouse.
	FieldCabinSeatRow2DriverSideHeating string = "Cabin.Seat.Row2.DriverSide.Heating"
	// FieldCabinSeatRow2DriverSideHeatingCooling is the name for the Cabin.Seat.Row2.DriverSide.HeatingCooling column in the Clickhouse.
	FieldCabinSeatRow2DriverSideHeatingCooling string = "Cabin.Seat.Row2.DriverSide.HeatingCooling"
	// FieldCabinSeatRow2DriverSideHeight is the name for the Cabin.Seat.Row2.DriverSide.Height column in the Clickhouse.
	FieldCabinSeatRow2DriverSideHeight string = "Cabin.Seat.Row2.DriverSide.Height"
	// FieldCabinSeatRow2DriverSideIsBelted is the name for the Cabin.Seat.Row2.DriverSide.IsBelted column in the Clickhouse.
	FieldCabinSeatRow2DriverSideIsBelted string = "Cabin.Seat.Row2.DriverSide.IsBelted"
	// FieldCabinSeatRow2DriverSideIsOccupied is the name for the Cabin.Seat.Row2.DriverSide.IsOccupied column in the Clickhouse.
	FieldCabinSeatRow2DriverSideIsOccupied string = "Cabin.Seat.Row2.DriverSide.IsOccupied"
	// FieldCabinSeatRow2DriverSideMassage is the name for the Cabin.Seat.Row2.DriverSide.Massage column in the Clickhouse.
	FieldCabinSeatRow2DriverSideMassage string = "Cabin.Seat.Row2.DriverSide.Massage"
	// FieldCabinSeatRow2DriverSideOccupantIdentifier is the name for the Cabin.Seat.Row2.DriverSide.Occupant_Identifier column in the Clickhouse.
	FieldCabinSeatRow2DriverSideOccupantIdentifier string = "Cabin.Seat.Row2.DriverSide.Occupant_Identifier"
	// FieldCabinSeatRow2DriverSideOccupantIdentifierIssuer is the name for the Cabin.Seat.Row2.DriverSide.Occupant_Identifier.Issuer column in the Clickhouse.
	FieldCabinSeatRow2DriverSideOccupantIdentifierIssuer string = "Cabin.Seat.Row2.DriverSide.Occupant_Identifier.Issuer"
	// FieldCabinSeatRow2DriverSideOccupantIdentifierSubject is the name for the Cabin.Seat.Row2.DriverSide.Occupant_Identifier.Subject column in the Clickhouse.
	FieldCabinSeatRow2DriverSideOccupantIdentifierSubject string = "Cabin.Seat.Row2.DriverSide.Occupant_Identifier.Subject"
	// FieldCabinSeatRow2DriverSidePosition is the name for the Cabin.Seat.Row2.DriverSide.Position column in the Clickhouse.
	FieldCabinSeatRow2DriverSidePosition string = "Cabin.Seat.Row2.DriverSide.Position"
	// FieldCabinSeatRow2DriverSideSeatBeltHeight is the name for the Cabin.Seat.Row2.DriverSide.SeatBeltHeight column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSeatBeltHeight string = "Cabin.Seat.Row2.DriverSide.SeatBeltHeight"
	// FieldCabinSeatRow2DriverSideSeatingLength is the name for the Cabin.Seat.Row2.DriverSide.Seating_Length column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSeatingLength string = "Cabin.Seat.Row2.DriverSide.Seating_Length"
	// FieldCabinSeatRow2DriverSideSwitch is the name for the Cabin.Seat.Row2.DriverSide.Switch column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitch string = "Cabin.Seat.Row2.DriverSide.Switch"
	// FieldCabinSeatRow2DriverSideSwitchBackrest is the name for the Cabin.Seat.Row2.DriverSide.Switch.Backrest column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchBackrest string = "Cabin.Seat.Row2.DriverSide.Switch.Backrest"
	// FieldCabinSeatRow2DriverSideSwitchBackrestIsReclineBackwardEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.Backrest.IsReclineBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchBackrestIsReclineBackwardEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.Backrest.IsReclineBackwardEngaged"
	// FieldCabinSeatRow2DriverSideSwitchBackrestIsReclineForwardEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.Backrest.IsReclineForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchBackrestIsReclineForwardEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.Backrest.IsReclineForwardEngaged"
	// FieldCabinSeatRow2DriverSideSwitchBackrestLumbar is the name for the Cabin.Seat.Row2.DriverSide.Switch.Backrest.Lumbar column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchBackrestLumbar string = "Cabin.Seat.Row2.DriverSide.Switch.Backrest.Lumbar"
	// FieldCabinSeatRow2DriverSideSwitchBackrestLumbarIsDownEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.Backrest.Lumbar.IsDownEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchBackrestLumbarIsDownEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.Backrest.Lumbar.IsDownEngaged"
	// FieldCabinSeatRow2DriverSideSwitchBackrestLumbarIsLessSupportEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.Backrest.Lumbar.IsLessSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchBackrestLumbarIsLessSupportEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.Backrest.Lumbar.IsLessSupportEngaged"
	// FieldCabinSeatRow2DriverSideSwitchBackrestLumbarIsMoreSupportEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.Backrest.Lumbar.IsMoreSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchBackrestLumbarIsMoreSupportEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.Backrest.Lumbar.IsMoreSupportEngaged"
	// FieldCabinSeatRow2DriverSideSwitchBackrestLumbarIsUpEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.Backrest.Lumbar.IsUpEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchBackrestLumbarIsUpEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.Backrest.Lumbar.IsUpEngaged"
	// FieldCabinSeatRow2DriverSideSwitchBackrestSideBolster is the name for the Cabin.Seat.Row2.DriverSide.Switch.Backrest.SideBolster column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchBackrestSideBolster string = "Cabin.Seat.Row2.DriverSide.Switch.Backrest.SideBolster"
	// FieldCabinSeatRow2DriverSideSwitchBackrestSideBolsterIsLessSupportEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.Backrest.SideBolster.IsLessSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchBackrestSideBolsterIsLessSupportEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.Backrest.SideBolster.IsLessSupportEngaged"
	// FieldCabinSeatRow2DriverSideSwitchBackrestSideBolsterIsMoreSupportEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.Backrest.SideBolster.IsMoreSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchBackrestSideBolsterIsMoreSupportEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.Backrest.SideBolster.IsMoreSupportEngaged"
	// FieldCabinSeatRow2DriverSideSwitchHeadrest is the name for the Cabin.Seat.Row2.DriverSide.Switch.Headrest column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchHeadrest string = "Cabin.Seat.Row2.DriverSide.Switch.Headrest"
	// FieldCabinSeatRow2DriverSideSwitchHeadrestIsBackwardEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.Headrest.IsBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchHeadrestIsBackwardEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.Headrest.IsBackwardEngaged"
	// FieldCabinSeatRow2DriverSideSwitchHeadrestIsDownEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.Headrest.IsDownEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchHeadrestIsDownEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.Headrest.IsDownEngaged"
	// FieldCabinSeatRow2DriverSideSwitchHeadrestIsForwardEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.Headrest.IsForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchHeadrestIsForwardEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.Headrest.IsForwardEngaged"
	// FieldCabinSeatRow2DriverSideSwitchHeadrestIsUpEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.Headrest.IsUpEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchHeadrestIsUpEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.Headrest.IsUpEngaged"
	// FieldCabinSeatRow2DriverSideSwitchIsBackwardEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.IsBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchIsBackwardEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.IsBackwardEngaged"
	// FieldCabinSeatRow2DriverSideSwitchIsCoolerEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.IsCoolerEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchIsCoolerEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.IsCoolerEngaged"
	// FieldCabinSeatRow2DriverSideSwitchIsDownEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.IsDownEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchIsDownEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.IsDownEngaged"
	// FieldCabinSeatRow2DriverSideSwitchIsForwardEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.IsForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchIsForwardEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.IsForwardEngaged"
	// FieldCabinSeatRow2DriverSideSwitchIsTiltBackwardEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.IsTiltBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchIsTiltBackwardEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.IsTiltBackwardEngaged"
	// FieldCabinSeatRow2DriverSideSwitchIsTiltForwardEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.IsTiltForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchIsTiltForwardEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.IsTiltForwardEngaged"
	// FieldCabinSeatRow2DriverSideSwitchIsUpEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.IsUpEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchIsUpEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.IsUpEngaged"
	// FieldCabinSeatRow2DriverSideSwitchIsWarmerEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.IsWarmerEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchIsWarmerEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.IsWarmerEngaged"
	// FieldCabinSeatRow2DriverSideSwitchMassage is the name for the Cabin.Seat.Row2.DriverSide.Switch.Massage column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchMassage string = "Cabin.Seat.Row2.DriverSide.Switch.Massage"
	// FieldCabinSeatRow2DriverSideSwitchMassageIsDecreaseEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.Massage.IsDecreaseEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchMassageIsDecreaseEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.Massage.IsDecreaseEngaged"
	// FieldCabinSeatRow2DriverSideSwitchMassageIsIncreaseEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.Massage.IsIncreaseEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchMassageIsIncreaseEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.Massage.IsIncreaseEngaged"
	// FieldCabinSeatRow2DriverSideSwitchSeating is the name for the Cabin.Seat.Row2.DriverSide.Switch.Seating column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchSeating string = "Cabin.Seat.Row2.DriverSide.Switch.Seating"
	// FieldCabinSeatRow2DriverSideSwitchSeatingIsBackwardEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.Seating.IsBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchSeatingIsBackwardEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.Seating.IsBackwardEngaged"
	// FieldCabinSeatRow2DriverSideSwitchSeatingIsForwardEngaged is the name for the Cabin.Seat.Row2.DriverSide.Switch.Seating.IsForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2DriverSideSwitchSeatingIsForwardEngaged string = "Cabin.Seat.Row2.DriverSide.Switch.Seating.IsForwardEngaged"
	// FieldCabinSeatRow2DriverSideTilt is the name for the Cabin.Seat.Row2.DriverSide.Tilt column in the Clickhouse.
	FieldCabinSeatRow2DriverSideTilt string = "Cabin.Seat.Row2.DriverSide.Tilt"
	// FieldCabinSeatRow2Middle is the name for the Cabin.Seat.Row2.Middle column in the Clickhouse.
	FieldCabinSeatRow2Middle string = "Cabin.Seat.Row2.Middle"
	// FieldCabinSeatRow2MiddleAirbagIsDeployed is the name for the Cabin.Seat.Row2.Middle.Airbag_IsDeployed column in the Clickhouse.
	FieldCabinSeatRow2MiddleAirbagIsDeployed string = "Cabin.Seat.Row2.Middle.Airbag_IsDeployed"
	// FieldCabinSeatRow2MiddleBackrest is the name for the Cabin.Seat.Row2.Middle.Backrest column in the Clickhouse.
	FieldCabinSeatRow2MiddleBackrest string = "Cabin.Seat.Row2.Middle.Backrest"
	// FieldCabinSeatRow2MiddleBackrestLumbar is the name for the Cabin.Seat.Row2.Middle.Backrest.Lumbar column in the Clickhouse.
	FieldCabinSeatRow2MiddleBackrestLumbar string = "Cabin.Seat.Row2.Middle.Backrest.Lumbar"
	// FieldCabinSeatRow2MiddleBackrestLumbarHeight is the name for the Cabin.Seat.Row2.Middle.Backrest.Lumbar.Height column in the Clickhouse.
	FieldCabinSeatRow2MiddleBackrestLumbarHeight string = "Cabin.Seat.Row2.Middle.Backrest.Lumbar.Height"
	// FieldCabinSeatRow2MiddleBackrestLumbarSupport is the name for the Cabin.Seat.Row2.Middle.Backrest.Lumbar.Support column in the Clickhouse.
	FieldCabinSeatRow2MiddleBackrestLumbarSupport string = "Cabin.Seat.Row2.Middle.Backrest.Lumbar.Support"
	// FieldCabinSeatRow2MiddleBackrestRecline is the name for the Cabin.Seat.Row2.Middle.Backrest.Recline column in the Clickhouse.
	FieldCabinSeatRow2MiddleBackrestRecline string = "Cabin.Seat.Row2.Middle.Backrest.Recline"
	// FieldCabinSeatRow2MiddleBackrestSideBolsterSupport is the name for the Cabin.Seat.Row2.Middle.Backrest.SideBolster_Support column in the Clickhouse.
	FieldCabinSeatRow2MiddleBackrestSideBolsterSupport string = "Cabin.Seat.Row2.Middle.Backrest.SideBolster_Support"
	// FieldCabinSeatRow2MiddleHeadrest is the name for the Cabin.Seat.Row2.Middle.Headrest column in the Clickhouse.
	FieldCabinSeatRow2MiddleHeadrest string = "Cabin.Seat.Row2.Middle.Headrest"
	// FieldCabinSeatRow2MiddleHeadrestAngle is the name for the Cabin.Seat.Row2.Middle.Headrest.Angle column in the Clickhouse.
	FieldCabinSeatRow2MiddleHeadrestAngle string = "Cabin.Seat.Row2.Middle.Headrest.Angle"
	// FieldCabinSeatRow2MiddleHeadrestHeight is the name for the Cabin.Seat.Row2.Middle.Headrest.Height column in the Clickhouse.
	FieldCabinSeatRow2MiddleHeadrestHeight string = "Cabin.Seat.Row2.Middle.Headrest.Height"
	// FieldCabinSeatRow2MiddleHeating is the name for the Cabin.Seat.Row2.Middle.Heating column in the Clickhouse.
	FieldCabinSeatRow2MiddleHeating string = "Cabin.Seat.Row2.Middle.Heating"
	// FieldCabinSeatRow2MiddleHeatingCooling is the name for the Cabin.Seat.Row2.Middle.HeatingCooling column in the Clickhouse.
	FieldCabinSeatRow2MiddleHeatingCooling string = "Cabin.Seat.Row2.Middle.HeatingCooling"
	// FieldCabinSeatRow2MiddleHeight is the name for the Cabin.Seat.Row2.Middle.Height column in the Clickhouse.
	FieldCabinSeatRow2MiddleHeight string = "Cabin.Seat.Row2.Middle.Height"
	// FieldCabinSeatRow2MiddleIsBelted is the name for the Cabin.Seat.Row2.Middle.IsBelted column in the Clickhouse.
	FieldCabinSeatRow2MiddleIsBelted string = "Cabin.Seat.Row2.Middle.IsBelted"
	// FieldCabinSeatRow2MiddleIsOccupied is the name for the Cabin.Seat.Row2.Middle.IsOccupied column in the Clickhouse.
	FieldCabinSeatRow2MiddleIsOccupied string = "Cabin.Seat.Row2.Middle.IsOccupied"
	// FieldCabinSeatRow2MiddleMassage is the name for the Cabin.Seat.Row2.Middle.Massage column in the Clickhouse.
	FieldCabinSeatRow2MiddleMassage string = "Cabin.Seat.Row2.Middle.Massage"
	// FieldCabinSeatRow2MiddleOccupantIdentifier is the name for the Cabin.Seat.Row2.Middle.Occupant_Identifier column in the Clickhouse.
	FieldCabinSeatRow2MiddleOccupantIdentifier string = "Cabin.Seat.Row2.Middle.Occupant_Identifier"
	// FieldCabinSeatRow2MiddleOccupantIdentifierIssuer is the name for the Cabin.Seat.Row2.Middle.Occupant_Identifier.Issuer column in the Clickhouse.
	FieldCabinSeatRow2MiddleOccupantIdentifierIssuer string = "Cabin.Seat.Row2.Middle.Occupant_Identifier.Issuer"
	// FieldCabinSeatRow2MiddleOccupantIdentifierSubject is the name for the Cabin.Seat.Row2.Middle.Occupant_Identifier.Subject column in the Clickhouse.
	FieldCabinSeatRow2MiddleOccupantIdentifierSubject string = "Cabin.Seat.Row2.Middle.Occupant_Identifier.Subject"
	// FieldCabinSeatRow2MiddlePosition is the name for the Cabin.Seat.Row2.Middle.Position column in the Clickhouse.
	FieldCabinSeatRow2MiddlePosition string = "Cabin.Seat.Row2.Middle.Position"
	// FieldCabinSeatRow2MiddleSeatBeltHeight is the name for the Cabin.Seat.Row2.Middle.SeatBeltHeight column in the Clickhouse.
	FieldCabinSeatRow2MiddleSeatBeltHeight string = "Cabin.Seat.Row2.Middle.SeatBeltHeight"
	// FieldCabinSeatRow2MiddleSeatingLength is the name for the Cabin.Seat.Row2.Middle.Seating_Length column in the Clickhouse.
	FieldCabinSeatRow2MiddleSeatingLength string = "Cabin.Seat.Row2.Middle.Seating_Length"
	// FieldCabinSeatRow2MiddleSwitch is the name for the Cabin.Seat.Row2.Middle.Switch column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitch string = "Cabin.Seat.Row2.Middle.Switch"
	// FieldCabinSeatRow2MiddleSwitchBackrest is the name for the Cabin.Seat.Row2.Middle.Switch.Backrest column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchBackrest string = "Cabin.Seat.Row2.Middle.Switch.Backrest"
	// FieldCabinSeatRow2MiddleSwitchBackrestIsReclineBackwardEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.Backrest.IsReclineBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchBackrestIsReclineBackwardEngaged string = "Cabin.Seat.Row2.Middle.Switch.Backrest.IsReclineBackwardEngaged"
	// FieldCabinSeatRow2MiddleSwitchBackrestIsReclineForwardEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.Backrest.IsReclineForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchBackrestIsReclineForwardEngaged string = "Cabin.Seat.Row2.Middle.Switch.Backrest.IsReclineForwardEngaged"
	// FieldCabinSeatRow2MiddleSwitchBackrestLumbar is the name for the Cabin.Seat.Row2.Middle.Switch.Backrest.Lumbar column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchBackrestLumbar string = "Cabin.Seat.Row2.Middle.Switch.Backrest.Lumbar"
	// FieldCabinSeatRow2MiddleSwitchBackrestLumbarIsDownEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.Backrest.Lumbar.IsDownEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchBackrestLumbarIsDownEngaged string = "Cabin.Seat.Row2.Middle.Switch.Backrest.Lumbar.IsDownEngaged"
	// FieldCabinSeatRow2MiddleSwitchBackrestLumbarIsLessSupportEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.Backrest.Lumbar.IsLessSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchBackrestLumbarIsLessSupportEngaged string = "Cabin.Seat.Row2.Middle.Switch.Backrest.Lumbar.IsLessSupportEngaged"
	// FieldCabinSeatRow2MiddleSwitchBackrestLumbarIsMoreSupportEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.Backrest.Lumbar.IsMoreSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchBackrestLumbarIsMoreSupportEngaged string = "Cabin.Seat.Row2.Middle.Switch.Backrest.Lumbar.IsMoreSupportEngaged"
	// FieldCabinSeatRow2MiddleSwitchBackrestLumbarIsUpEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.Backrest.Lumbar.IsUpEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchBackrestLumbarIsUpEngaged string = "Cabin.Seat.Row2.Middle.Switch.Backrest.Lumbar.IsUpEngaged"
	// FieldCabinSeatRow2MiddleSwitchBackrestSideBolster is the name for the Cabin.Seat.Row2.Middle.Switch.Backrest.SideBolster column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchBackrestSideBolster string = "Cabin.Seat.Row2.Middle.Switch.Backrest.SideBolster"
	// FieldCabinSeatRow2MiddleSwitchBackrestSideBolsterIsLessSupportEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.Backrest.SideBolster.IsLessSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchBackrestSideBolsterIsLessSupportEngaged string = "Cabin.Seat.Row2.Middle.Switch.Backrest.SideBolster.IsLessSupportEngaged"
	// FieldCabinSeatRow2MiddleSwitchBackrestSideBolsterIsMoreSupportEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.Backrest.SideBolster.IsMoreSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchBackrestSideBolsterIsMoreSupportEngaged string = "Cabin.Seat.Row2.Middle.Switch.Backrest.SideBolster.IsMoreSupportEngaged"
	// FieldCabinSeatRow2MiddleSwitchHeadrest is the name for the Cabin.Seat.Row2.Middle.Switch.Headrest column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchHeadrest string = "Cabin.Seat.Row2.Middle.Switch.Headrest"
	// FieldCabinSeatRow2MiddleSwitchHeadrestIsBackwardEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.Headrest.IsBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchHeadrestIsBackwardEngaged string = "Cabin.Seat.Row2.Middle.Switch.Headrest.IsBackwardEngaged"
	// FieldCabinSeatRow2MiddleSwitchHeadrestIsDownEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.Headrest.IsDownEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchHeadrestIsDownEngaged string = "Cabin.Seat.Row2.Middle.Switch.Headrest.IsDownEngaged"
	// FieldCabinSeatRow2MiddleSwitchHeadrestIsForwardEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.Headrest.IsForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchHeadrestIsForwardEngaged string = "Cabin.Seat.Row2.Middle.Switch.Headrest.IsForwardEngaged"
	// FieldCabinSeatRow2MiddleSwitchHeadrestIsUpEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.Headrest.IsUpEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchHeadrestIsUpEngaged string = "Cabin.Seat.Row2.Middle.Switch.Headrest.IsUpEngaged"
	// FieldCabinSeatRow2MiddleSwitchIsBackwardEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.IsBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchIsBackwardEngaged string = "Cabin.Seat.Row2.Middle.Switch.IsBackwardEngaged"
	// FieldCabinSeatRow2MiddleSwitchIsCoolerEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.IsCoolerEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchIsCoolerEngaged string = "Cabin.Seat.Row2.Middle.Switch.IsCoolerEngaged"
	// FieldCabinSeatRow2MiddleSwitchIsDownEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.IsDownEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchIsDownEngaged string = "Cabin.Seat.Row2.Middle.Switch.IsDownEngaged"
	// FieldCabinSeatRow2MiddleSwitchIsForwardEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.IsForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchIsForwardEngaged string = "Cabin.Seat.Row2.Middle.Switch.IsForwardEngaged"
	// FieldCabinSeatRow2MiddleSwitchIsTiltBackwardEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.IsTiltBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchIsTiltBackwardEngaged string = "Cabin.Seat.Row2.Middle.Switch.IsTiltBackwardEngaged"
	// FieldCabinSeatRow2MiddleSwitchIsTiltForwardEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.IsTiltForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchIsTiltForwardEngaged string = "Cabin.Seat.Row2.Middle.Switch.IsTiltForwardEngaged"
	// FieldCabinSeatRow2MiddleSwitchIsUpEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.IsUpEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchIsUpEngaged string = "Cabin.Seat.Row2.Middle.Switch.IsUpEngaged"
	// FieldCabinSeatRow2MiddleSwitchIsWarmerEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.IsWarmerEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchIsWarmerEngaged string = "Cabin.Seat.Row2.Middle.Switch.IsWarmerEngaged"
	// FieldCabinSeatRow2MiddleSwitchMassage is the name for the Cabin.Seat.Row2.Middle.Switch.Massage column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchMassage string = "Cabin.Seat.Row2.Middle.Switch.Massage"
	// FieldCabinSeatRow2MiddleSwitchMassageIsDecreaseEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.Massage.IsDecreaseEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchMassageIsDecreaseEngaged string = "Cabin.Seat.Row2.Middle.Switch.Massage.IsDecreaseEngaged"
	// FieldCabinSeatRow2MiddleSwitchMassageIsIncreaseEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.Massage.IsIncreaseEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchMassageIsIncreaseEngaged string = "Cabin.Seat.Row2.Middle.Switch.Massage.IsIncreaseEngaged"
	// FieldCabinSeatRow2MiddleSwitchSeating is the name for the Cabin.Seat.Row2.Middle.Switch.Seating column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchSeating string = "Cabin.Seat.Row2.Middle.Switch.Seating"
	// FieldCabinSeatRow2MiddleSwitchSeatingIsBackwardEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.Seating.IsBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchSeatingIsBackwardEngaged string = "Cabin.Seat.Row2.Middle.Switch.Seating.IsBackwardEngaged"
	// FieldCabinSeatRow2MiddleSwitchSeatingIsForwardEngaged is the name for the Cabin.Seat.Row2.Middle.Switch.Seating.IsForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2MiddleSwitchSeatingIsForwardEngaged string = "Cabin.Seat.Row2.Middle.Switch.Seating.IsForwardEngaged"
	// FieldCabinSeatRow2MiddleTilt is the name for the Cabin.Seat.Row2.Middle.Tilt column in the Clickhouse.
	FieldCabinSeatRow2MiddleTilt string = "Cabin.Seat.Row2.Middle.Tilt"
	// FieldCabinSeatRow2PassengerSide is the name for the Cabin.Seat.Row2.PassengerSide column in the Clickhouse.
	FieldCabinSeatRow2PassengerSide string = "Cabin.Seat.Row2.PassengerSide"
	// FieldCabinSeatRow2PassengerSideAirbagIsDeployed is the name for the Cabin.Seat.Row2.PassengerSide.Airbag_IsDeployed column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideAirbagIsDeployed string = "Cabin.Seat.Row2.PassengerSide.Airbag_IsDeployed"
	// FieldCabinSeatRow2PassengerSideBackrest is the name for the Cabin.Seat.Row2.PassengerSide.Backrest column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideBackrest string = "Cabin.Seat.Row2.PassengerSide.Backrest"
	// FieldCabinSeatRow2PassengerSideBackrestLumbar is the name for the Cabin.Seat.Row2.PassengerSide.Backrest.Lumbar column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideBackrestLumbar string = "Cabin.Seat.Row2.PassengerSide.Backrest.Lumbar"
	// FieldCabinSeatRow2PassengerSideBackrestLumbarHeight is the name for the Cabin.Seat.Row2.PassengerSide.Backrest.Lumbar.Height column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideBackrestLumbarHeight string = "Cabin.Seat.Row2.PassengerSide.Backrest.Lumbar.Height"
	// FieldCabinSeatRow2PassengerSideBackrestLumbarSupport is the name for the Cabin.Seat.Row2.PassengerSide.Backrest.Lumbar.Support column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideBackrestLumbarSupport string = "Cabin.Seat.Row2.PassengerSide.Backrest.Lumbar.Support"
	// FieldCabinSeatRow2PassengerSideBackrestRecline is the name for the Cabin.Seat.Row2.PassengerSide.Backrest.Recline column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideBackrestRecline string = "Cabin.Seat.Row2.PassengerSide.Backrest.Recline"
	// FieldCabinSeatRow2PassengerSideBackrestSideBolsterSupport is the name for the Cabin.Seat.Row2.PassengerSide.Backrest.SideBolster_Support column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideBackrestSideBolsterSupport string = "Cabin.Seat.Row2.PassengerSide.Backrest.SideBolster_Support"
	// FieldCabinSeatRow2PassengerSideHeadrest is the name for the Cabin.Seat.Row2.PassengerSide.Headrest column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideHeadrest string = "Cabin.Seat.Row2.PassengerSide.Headrest"
	// FieldCabinSeatRow2PassengerSideHeadrestAngle is the name for the Cabin.Seat.Row2.PassengerSide.Headrest.Angle column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideHeadrestAngle string = "Cabin.Seat.Row2.PassengerSide.Headrest.Angle"
	// FieldCabinSeatRow2PassengerSideHeadrestHeight is the name for the Cabin.Seat.Row2.PassengerSide.Headrest.Height column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideHeadrestHeight string = "Cabin.Seat.Row2.PassengerSide.Headrest.Height"
	// FieldCabinSeatRow2PassengerSideHeating is the name for the Cabin.Seat.Row2.PassengerSide.Heating column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideHeating string = "Cabin.Seat.Row2.PassengerSide.Heating"
	// FieldCabinSeatRow2PassengerSideHeatingCooling is the name for the Cabin.Seat.Row2.PassengerSide.HeatingCooling column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideHeatingCooling string = "Cabin.Seat.Row2.PassengerSide.HeatingCooling"
	// FieldCabinSeatRow2PassengerSideHeight is the name for the Cabin.Seat.Row2.PassengerSide.Height column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideHeight string = "Cabin.Seat.Row2.PassengerSide.Height"
	// FieldCabinSeatRow2PassengerSideIsBelted is the name for the Cabin.Seat.Row2.PassengerSide.IsBelted column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideIsBelted string = "Cabin.Seat.Row2.PassengerSide.IsBelted"
	// FieldCabinSeatRow2PassengerSideIsOccupied is the name for the Cabin.Seat.Row2.PassengerSide.IsOccupied column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideIsOccupied string = "Cabin.Seat.Row2.PassengerSide.IsOccupied"
	// FieldCabinSeatRow2PassengerSideMassage is the name for the Cabin.Seat.Row2.PassengerSide.Massage column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideMassage string = "Cabin.Seat.Row2.PassengerSide.Massage"
	// FieldCabinSeatRow2PassengerSideOccupantIdentifier is the name for the Cabin.Seat.Row2.PassengerSide.Occupant_Identifier column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideOccupantIdentifier string = "Cabin.Seat.Row2.PassengerSide.Occupant_Identifier"
	// FieldCabinSeatRow2PassengerSideOccupantIdentifierIssuer is the name for the Cabin.Seat.Row2.PassengerSide.Occupant_Identifier.Issuer column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideOccupantIdentifierIssuer string = "Cabin.Seat.Row2.PassengerSide.Occupant_Identifier.Issuer"
	// FieldCabinSeatRow2PassengerSideOccupantIdentifierSubject is the name for the Cabin.Seat.Row2.PassengerSide.Occupant_Identifier.Subject column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideOccupantIdentifierSubject string = "Cabin.Seat.Row2.PassengerSide.Occupant_Identifier.Subject"
	// FieldCabinSeatRow2PassengerSidePosition is the name for the Cabin.Seat.Row2.PassengerSide.Position column in the Clickhouse.
	FieldCabinSeatRow2PassengerSidePosition string = "Cabin.Seat.Row2.PassengerSide.Position"
	// FieldCabinSeatRow2PassengerSideSeatBeltHeight is the name for the Cabin.Seat.Row2.PassengerSide.SeatBeltHeight column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSeatBeltHeight string = "Cabin.Seat.Row2.PassengerSide.SeatBeltHeight"
	// FieldCabinSeatRow2PassengerSideSeatingLength is the name for the Cabin.Seat.Row2.PassengerSide.Seating_Length column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSeatingLength string = "Cabin.Seat.Row2.PassengerSide.Seating_Length"
	// FieldCabinSeatRow2PassengerSideSwitch is the name for the Cabin.Seat.Row2.PassengerSide.Switch column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitch string = "Cabin.Seat.Row2.PassengerSide.Switch"
	// FieldCabinSeatRow2PassengerSideSwitchBackrest is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Backrest column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchBackrest string = "Cabin.Seat.Row2.PassengerSide.Switch.Backrest"
	// FieldCabinSeatRow2PassengerSideSwitchBackrestIsReclineBackwardEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Backrest.IsReclineBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchBackrestIsReclineBackwardEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.Backrest.IsReclineBackwardEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchBackrestIsReclineForwardEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Backrest.IsReclineForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchBackrestIsReclineForwardEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.Backrest.IsReclineForwardEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchBackrestLumbar is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Backrest.Lumbar column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchBackrestLumbar string = "Cabin.Seat.Row2.PassengerSide.Switch.Backrest.Lumbar"
	// FieldCabinSeatRow2PassengerSideSwitchBackrestLumbarIsDownEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Backrest.Lumbar.IsDownEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchBackrestLumbarIsDownEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.Backrest.Lumbar.IsDownEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchBackrestLumbarIsLessSupportEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Backrest.Lumbar.IsLessSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchBackrestLumbarIsLessSupportEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.Backrest.Lumbar.IsLessSupportEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchBackrestLumbarIsMoreSupportEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Backrest.Lumbar.IsMoreSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchBackrestLumbarIsMoreSupportEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.Backrest.Lumbar.IsMoreSupportEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchBackrestLumbarIsUpEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Backrest.Lumbar.IsUpEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchBackrestLumbarIsUpEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.Backrest.Lumbar.IsUpEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchBackrestSideBolster is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Backrest.SideBolster column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchBackrestSideBolster string = "Cabin.Seat.Row2.PassengerSide.Switch.Backrest.SideBolster"
	// FieldCabinSeatRow2PassengerSideSwitchBackrestSideBolsterIsLessSupportEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Backrest.SideBolster.IsLessSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchBackrestSideBolsterIsLessSupportEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.Backrest.SideBolster.IsLessSupportEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchBackrestSideBolsterIsMoreSupportEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Backrest.SideBolster.IsMoreSupportEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchBackrestSideBolsterIsMoreSupportEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.Backrest.SideBolster.IsMoreSupportEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchHeadrest is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Headrest column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchHeadrest string = "Cabin.Seat.Row2.PassengerSide.Switch.Headrest"
	// FieldCabinSeatRow2PassengerSideSwitchHeadrestIsBackwardEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Headrest.IsBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchHeadrestIsBackwardEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.Headrest.IsBackwardEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchHeadrestIsDownEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Headrest.IsDownEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchHeadrestIsDownEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.Headrest.IsDownEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchHeadrestIsForwardEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Headrest.IsForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchHeadrestIsForwardEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.Headrest.IsForwardEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchHeadrestIsUpEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Headrest.IsUpEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchHeadrestIsUpEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.Headrest.IsUpEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchIsBackwardEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.IsBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchIsBackwardEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.IsBackwardEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchIsCoolerEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.IsCoolerEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchIsCoolerEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.IsCoolerEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchIsDownEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.IsDownEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchIsDownEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.IsDownEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchIsForwardEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.IsForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchIsForwardEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.IsForwardEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchIsTiltBackwardEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.IsTiltBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchIsTiltBackwardEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.IsTiltBackwardEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchIsTiltForwardEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.IsTiltForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchIsTiltForwardEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.IsTiltForwardEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchIsUpEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.IsUpEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchIsUpEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.IsUpEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchIsWarmerEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.IsWarmerEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchIsWarmerEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.IsWarmerEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchMassage is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Massage column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchMassage string = "Cabin.Seat.Row2.PassengerSide.Switch.Massage"
	// FieldCabinSeatRow2PassengerSideSwitchMassageIsDecreaseEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Massage.IsDecreaseEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchMassageIsDecreaseEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.Massage.IsDecreaseEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchMassageIsIncreaseEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Massage.IsIncreaseEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchMassageIsIncreaseEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.Massage.IsIncreaseEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchSeating is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Seating column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchSeating string = "Cabin.Seat.Row2.PassengerSide.Switch.Seating"
	// FieldCabinSeatRow2PassengerSideSwitchSeatingIsBackwardEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Seating.IsBackwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchSeatingIsBackwardEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.Seating.IsBackwardEngaged"
	// FieldCabinSeatRow2PassengerSideSwitchSeatingIsForwardEngaged is the name for the Cabin.Seat.Row2.PassengerSide.Switch.Seating.IsForwardEngaged column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideSwitchSeatingIsForwardEngaged string = "Cabin.Seat.Row2.PassengerSide.Switch.Seating.IsForwardEngaged"
	// FieldCabinSeatRow2PassengerSideTilt is the name for the Cabin.Seat.Row2.PassengerSide.Tilt column in the Clickhouse.
	FieldCabinSeatRow2PassengerSideTilt string = "Cabin.Seat.Row2.PassengerSide.Tilt"
	// FieldCabinSeatRowCount is the name for the Cabin.SeatRowCount column in the Clickhouse.
	FieldCabinSeatRowCount string = "Cabin.SeatRowCount"
	// FieldCabinSunroof is the name for the Cabin.Sunroof column in the Clickhouse.
	FieldCabinSunroof string = "Cabin.Sunroof"
	// FieldCabinSunroofPosition is the name for the Cabin.Sunroof.Position column in the Clickhouse.
	FieldCabinSunroofPosition string = "Cabin.Sunroof.Position"
	// FieldCabinSunroofShade is the name for the Cabin.Sunroof.Shade column in the Clickhouse.
	FieldCabinSunroofShade string = "Cabin.Sunroof.Shade"
	// FieldCabinSunroofShadeIsOpen is the name for the Cabin.Sunroof.Shade.IsOpen column in the Clickhouse.
	FieldCabinSunroofShadeIsOpen string = "Cabin.Sunroof.Shade.IsOpen"
	// FieldCabinSunroofShadePosition is the name for the Cabin.Sunroof.Shade.Position column in the Clickhouse.
	FieldCabinSunroofShadePosition string = "Cabin.Sunroof.Shade.Position"
	// FieldCabinSunroofShadeSwitch is the name for the Cabin.Sunroof.Shade.Switch column in the Clickhouse.
	FieldCabinSunroofShadeSwitch string = "Cabin.Sunroof.Shade.Switch"
	// FieldCabinSunroofSwitch is the name for the Cabin.Sunroof.Switch column in the Clickhouse.
	FieldCabinSunroofSwitch string = "Cabin.Sunroof.Switch"
	// FieldCargoVolume is the name for the CargoVolume column in the Clickhouse.
	FieldCargoVolume string = "CargoVolume"
	// FieldChassis is the name for the Chassis column in the Clickhouse.
	FieldChassis string = "Chassis"
	// FieldChassisAcceleratorPedalPosition is the name for the Chassis.Accelerator_PedalPosition column in the Clickhouse.
	FieldChassisAcceleratorPedalPosition string = "Chassis.Accelerator_PedalPosition"
	// FieldChassisAxle is the name for the Chassis.Axle column in the Clickhouse.
	FieldChassisAxle string = "Chassis.Axle"
	// FieldChassisAxleCount is the name for the Chassis.AxleCount column in the Clickhouse.
	FieldChassisAxleCount string = "Chassis.AxleCount"
	// FieldChassisAxleRow1 is the name for the Chassis.Axle.Row1 column in the Clickhouse.
	FieldChassisAxleRow1 string = "Chassis.Axle.Row1"
	// FieldChassisAxleRow1AxleWidth is the name for the Chassis.Axle.Row1.AxleWidth column in the Clickhouse.
	FieldChassisAxleRow1AxleWidth string = "Chassis.Axle.Row1.AxleWidth"
	// FieldChassisAxleRow1SteeringAngle is the name for the Chassis.Axle.Row1.SteeringAngle column in the Clickhouse.
	FieldChassisAxleRow1SteeringAngle string = "Chassis.Axle.Row1.SteeringAngle"
	// FieldChassisAxleRow1TireAspectRatio is the name for the Chassis.Axle.Row1.TireAspectRatio column in the Clickhouse.
	FieldChassisAxleRow1TireAspectRatio string = "Chassis.Axle.Row1.TireAspectRatio"
	// FieldChassisAxleRow1TireDiameter is the name for the Chassis.Axle.Row1.TireDiameter column in the Clickhouse.
	FieldChassisAxleRow1TireDiameter string = "Chassis.Axle.Row1.TireDiameter"
	// FieldChassisAxleRow1TireWidth is the name for the Chassis.Axle.Row1.TireWidth column in the Clickhouse.
	FieldChassisAxleRow1TireWidth string = "Chassis.Axle.Row1.TireWidth"
	// FieldChassisAxleRow1TrackWidth is the name for the Chassis.Axle.Row1.TrackWidth column in the Clickhouse.
	FieldChassisAxleRow1TrackWidth string = "Chassis.Axle.Row1.TrackWidth"
	// FieldChassisAxleRow1TreadWidth is the name for the Chassis.Axle.Row1.TreadWidth column in the Clickhouse.
	FieldChassisAxleRow1TreadWidth string = "Chassis.Axle.Row1.TreadWidth"
	// FieldChassisAxleRow1Wheel is the name for the Chassis.Axle.Row1.Wheel column in the Clickhouse.
	FieldChassisAxleRow1Wheel string = "Chassis.Axle.Row1.Wheel"
	// FieldChassisAxleRow1WheelCount is the name for the Chassis.Axle.Row1.WheelCount column in the Clickhouse.
	FieldChassisAxleRow1WheelCount string = "Chassis.Axle.Row1.WheelCount"
	// FieldChassisAxleRow1WheelDiameter is the name for the Chassis.Axle.Row1.WheelDiameter column in the Clickhouse.
	FieldChassisAxleRow1WheelDiameter string = "Chassis.Axle.Row1.WheelDiameter"
	// FieldChassisAxleRow1WheelLeft is the name for the Chassis.Axle.Row1.Wheel.Left column in the Clickhouse.
	FieldChassisAxleRow1WheelLeft string = "Chassis.Axle.Row1.Wheel.Left"
	// FieldChassisAxleRow1WheelLeftBrake is the name for the Chassis.Axle.Row1.Wheel.Left.Brake column in the Clickhouse.
	FieldChassisAxleRow1WheelLeftBrake string = "Chassis.Axle.Row1.Wheel.Left.Brake"
	// FieldChassisAxleRow1WheelLeftBrakeFluidLevel is the name for the Chassis.Axle.Row1.Wheel.Left.Brake.FluidLevel column in the Clickhouse.
	FieldChassisAxleRow1WheelLeftBrakeFluidLevel string = "Chassis.Axle.Row1.Wheel.Left.Brake.FluidLevel"
	// FieldChassisAxleRow1WheelLeftBrakeIsBrakesWorn is the name for the Chassis.Axle.Row1.Wheel.Left.Brake.IsBrakesWorn column in the Clickhouse.
	FieldChassisAxleRow1WheelLeftBrakeIsBrakesWorn string = "Chassis.Axle.Row1.Wheel.Left.Brake.IsBrakesWorn"
	// FieldChassisAxleRow1WheelLeftBrakeIsFluidLevelLow is the name for the Chassis.Axle.Row1.Wheel.Left.Brake.IsFluidLevelLow column in the Clickhouse.
	FieldChassisAxleRow1WheelLeftBrakeIsFluidLevelLow string = "Chassis.Axle.Row1.Wheel.Left.Brake.IsFluidLevelLow"
	// FieldChassisAxleRow1WheelLeftBrakePadWear is the name for the Chassis.Axle.Row1.Wheel.Left.Brake.PadWear column in the Clickhouse.
	FieldChassisAxleRow1WheelLeftBrakePadWear string = "Chassis.Axle.Row1.Wheel.Left.Brake.PadWear"
	// FieldChassisAxleRow1WheelLeftSpeed is the name for the Chassis.Axle.Row1.Wheel.Left.Speed column in the Clickhouse.
	FieldChassisAxleRow1WheelLeftSpeed string = "Chassis.Axle.Row1.Wheel.Left.Speed"
	// FieldChassisAxleRow1WheelLeftTire is the name for the Chassis.Axle.Row1.Wheel.Left.Tire column in the Clickhouse.
	FieldChassisAxleRow1WheelLeftTire string = "Chassis.Axle.Row1.Wheel.Left.Tire"
	// FieldChassisAxleRow1WheelLeftTireIsPressureLow is the name for the Chassis.Axle.Row1.Wheel.Left.Tire.IsPressureLow column in the Clickhouse.
	FieldChassisAxleRow1WheelLeftTireIsPressureLow string = "Chassis.Axle.Row1.Wheel.Left.Tire.IsPressureLow"
	// FieldChassisAxleRow1WheelLeftTirePressure is the name for the Chassis.Axle.Row1.Wheel.Left.Tire.Pressure column in the Clickhouse.
	FieldChassisAxleRow1WheelLeftTirePressure string = "Chassis.Axle.Row1.Wheel.Left.Tire.Pressure"
	// FieldChassisAxleRow1WheelLeftTireTemperature is the name for the Chassis.Axle.Row1.Wheel.Left.Tire.Temperature column in the Clickhouse.
	FieldChassisAxleRow1WheelLeftTireTemperature string = "Chassis.Axle.Row1.Wheel.Left.Tire.Temperature"
	// FieldChassisAxleRow1WheelRight is the name for the Chassis.Axle.Row1.Wheel.Right column in the Clickhouse.
	FieldChassisAxleRow1WheelRight string = "Chassis.Axle.Row1.Wheel.Right"
	// FieldChassisAxleRow1WheelRightBrake is the name for the Chassis.Axle.Row1.Wheel.Right.Brake column in the Clickhouse.
	FieldChassisAxleRow1WheelRightBrake string = "Chassis.Axle.Row1.Wheel.Right.Brake"
	// FieldChassisAxleRow1WheelRightBrakeFluidLevel is the name for the Chassis.Axle.Row1.Wheel.Right.Brake.FluidLevel column in the Clickhouse.
	FieldChassisAxleRow1WheelRightBrakeFluidLevel string = "Chassis.Axle.Row1.Wheel.Right.Brake.FluidLevel"
	// FieldChassisAxleRow1WheelRightBrakeIsBrakesWorn is the name for the Chassis.Axle.Row1.Wheel.Right.Brake.IsBrakesWorn column in the Clickhouse.
	FieldChassisAxleRow1WheelRightBrakeIsBrakesWorn string = "Chassis.Axle.Row1.Wheel.Right.Brake.IsBrakesWorn"
	// FieldChassisAxleRow1WheelRightBrakeIsFluidLevelLow is the name for the Chassis.Axle.Row1.Wheel.Right.Brake.IsFluidLevelLow column in the Clickhouse.
	FieldChassisAxleRow1WheelRightBrakeIsFluidLevelLow string = "Chassis.Axle.Row1.Wheel.Right.Brake.IsFluidLevelLow"
	// FieldChassisAxleRow1WheelRightBrakePadWear is the name for the Chassis.Axle.Row1.Wheel.Right.Brake.PadWear column in the Clickhouse.
	FieldChassisAxleRow1WheelRightBrakePadWear string = "Chassis.Axle.Row1.Wheel.Right.Brake.PadWear"
	// FieldChassisAxleRow1WheelRightSpeed is the name for the Chassis.Axle.Row1.Wheel.Right.Speed column in the Clickhouse.
	FieldChassisAxleRow1WheelRightSpeed string = "Chassis.Axle.Row1.Wheel.Right.Speed"
	// FieldChassisAxleRow1WheelRightTire is the name for the Chassis.Axle.Row1.Wheel.Right.Tire column in the Clickhouse.
	FieldChassisAxleRow1WheelRightTire string = "Chassis.Axle.Row1.Wheel.Right.Tire"
	// FieldChassisAxleRow1WheelRightTireIsPressureLow is the name for the Chassis.Axle.Row1.Wheel.Right.Tire.IsPressureLow column in the Clickhouse.
	FieldChassisAxleRow1WheelRightTireIsPressureLow string = "Chassis.Axle.Row1.Wheel.Right.Tire.IsPressureLow"
	// FieldChassisAxleRow1WheelRightTirePressure is the name for the Chassis.Axle.Row1.Wheel.Right.Tire.Pressure column in the Clickhouse.
	FieldChassisAxleRow1WheelRightTirePressure string = "Chassis.Axle.Row1.Wheel.Right.Tire.Pressure"
	// FieldChassisAxleRow1WheelRightTireTemperature is the name for the Chassis.Axle.Row1.Wheel.Right.Tire.Temperature column in the Clickhouse.
	FieldChassisAxleRow1WheelRightTireTemperature string = "Chassis.Axle.Row1.Wheel.Right.Tire.Temperature"
	// FieldChassisAxleRow1WheelWidth is the name for the Chassis.Axle.Row1.WheelWidth column in the Clickhouse.
	FieldChassisAxleRow1WheelWidth string = "Chassis.Axle.Row1.WheelWidth"
	// FieldChassisAxleRow2 is the name for the Chassis.Axle.Row2 column in the Clickhouse.
	FieldChassisAxleRow2 string = "Chassis.Axle.Row2"
	// FieldChassisAxleRow2AxleWidth is the name for the Chassis.Axle.Row2.AxleWidth column in the Clickhouse.
	FieldChassisAxleRow2AxleWidth string = "Chassis.Axle.Row2.AxleWidth"
	// FieldChassisAxleRow2SteeringAngle is the name for the Chassis.Axle.Row2.SteeringAngle column in the Clickhouse.
	FieldChassisAxleRow2SteeringAngle string = "Chassis.Axle.Row2.SteeringAngle"
	// FieldChassisAxleRow2TireAspectRatio is the name for the Chassis.Axle.Row2.TireAspectRatio column in the Clickhouse.
	FieldChassisAxleRow2TireAspectRatio string = "Chassis.Axle.Row2.TireAspectRatio"
	// FieldChassisAxleRow2TireDiameter is the name for the Chassis.Axle.Row2.TireDiameter column in the Clickhouse.
	FieldChassisAxleRow2TireDiameter string = "Chassis.Axle.Row2.TireDiameter"
	// FieldChassisAxleRow2TireWidth is the name for the Chassis.Axle.Row2.TireWidth column in the Clickhouse.
	FieldChassisAxleRow2TireWidth string = "Chassis.Axle.Row2.TireWidth"
	// FieldChassisAxleRow2TrackWidth is the name for the Chassis.Axle.Row2.TrackWidth column in the Clickhouse.
	FieldChassisAxleRow2TrackWidth string = "Chassis.Axle.Row2.TrackWidth"
	// FieldChassisAxleRow2TreadWidth is the name for the Chassis.Axle.Row2.TreadWidth column in the Clickhouse.
	FieldChassisAxleRow2TreadWidth string = "Chassis.Axle.Row2.TreadWidth"
	// FieldChassisAxleRow2Wheel is the name for the Chassis.Axle.Row2.Wheel column in the Clickhouse.
	FieldChassisAxleRow2Wheel string = "Chassis.Axle.Row2.Wheel"
	// FieldChassisAxleRow2WheelCount is the name for the Chassis.Axle.Row2.WheelCount column in the Clickhouse.
	FieldChassisAxleRow2WheelCount string = "Chassis.Axle.Row2.WheelCount"
	// FieldChassisAxleRow2WheelDiameter is the name for the Chassis.Axle.Row2.WheelDiameter column in the Clickhouse.
	FieldChassisAxleRow2WheelDiameter string = "Chassis.Axle.Row2.WheelDiameter"
	// FieldChassisAxleRow2WheelLeft is the name for the Chassis.Axle.Row2.Wheel.Left column in the Clickhouse.
	FieldChassisAxleRow2WheelLeft string = "Chassis.Axle.Row2.Wheel.Left"
	// FieldChassisAxleRow2WheelLeftBrake is the name for the Chassis.Axle.Row2.Wheel.Left.Brake column in the Clickhouse.
	FieldChassisAxleRow2WheelLeftBrake string = "Chassis.Axle.Row2.Wheel.Left.Brake"
	// FieldChassisAxleRow2WheelLeftBrakeFluidLevel is the name for the Chassis.Axle.Row2.Wheel.Left.Brake.FluidLevel column in the Clickhouse.
	FieldChassisAxleRow2WheelLeftBrakeFluidLevel string = "Chassis.Axle.Row2.Wheel.Left.Brake.FluidLevel"
	// FieldChassisAxleRow2WheelLeftBrakeIsBrakesWorn is the name for the Chassis.Axle.Row2.Wheel.Left.Brake.IsBrakesWorn column in the Clickhouse.
	FieldChassisAxleRow2WheelLeftBrakeIsBrakesWorn string = "Chassis.Axle.Row2.Wheel.Left.Brake.IsBrakesWorn"
	// FieldChassisAxleRow2WheelLeftBrakeIsFluidLevelLow is the name for the Chassis.Axle.Row2.Wheel.Left.Brake.IsFluidLevelLow column in the Clickhouse.
	FieldChassisAxleRow2WheelLeftBrakeIsFluidLevelLow string = "Chassis.Axle.Row2.Wheel.Left.Brake.IsFluidLevelLow"
	// FieldChassisAxleRow2WheelLeftBrakePadWear is the name for the Chassis.Axle.Row2.Wheel.Left.Brake.PadWear column in the Clickhouse.
	FieldChassisAxleRow2WheelLeftBrakePadWear string = "Chassis.Axle.Row2.Wheel.Left.Brake.PadWear"
	// FieldChassisAxleRow2WheelLeftSpeed is the name for the Chassis.Axle.Row2.Wheel.Left.Speed column in the Clickhouse.
	FieldChassisAxleRow2WheelLeftSpeed string = "Chassis.Axle.Row2.Wheel.Left.Speed"
	// FieldChassisAxleRow2WheelLeftTire is the name for the Chassis.Axle.Row2.Wheel.Left.Tire column in the Clickhouse.
	FieldChassisAxleRow2WheelLeftTire string = "Chassis.Axle.Row2.Wheel.Left.Tire"
	// FieldChassisAxleRow2WheelLeftTireIsPressureLow is the name for the Chassis.Axle.Row2.Wheel.Left.Tire.IsPressureLow column in the Clickhouse.
	FieldChassisAxleRow2WheelLeftTireIsPressureLow string = "Chassis.Axle.Row2.Wheel.Left.Tire.IsPressureLow"
	// FieldChassisAxleRow2WheelLeftTirePressure is the name for the Chassis.Axle.Row2.Wheel.Left.Tire.Pressure column in the Clickhouse.
	FieldChassisAxleRow2WheelLeftTirePressure string = "Chassis.Axle.Row2.Wheel.Left.Tire.Pressure"
	// FieldChassisAxleRow2WheelLeftTireTemperature is the name for the Chassis.Axle.Row2.Wheel.Left.Tire.Temperature column in the Clickhouse.
	FieldChassisAxleRow2WheelLeftTireTemperature string = "Chassis.Axle.Row2.Wheel.Left.Tire.Temperature"
	// FieldChassisAxleRow2WheelRight is the name for the Chassis.Axle.Row2.Wheel.Right column in the Clickhouse.
	FieldChassisAxleRow2WheelRight string = "Chassis.Axle.Row2.Wheel.Right"
	// FieldChassisAxleRow2WheelRightBrake is the name for the Chassis.Axle.Row2.Wheel.Right.Brake column in the Clickhouse.
	FieldChassisAxleRow2WheelRightBrake string = "Chassis.Axle.Row2.Wheel.Right.Brake"
	// FieldChassisAxleRow2WheelRightBrakeFluidLevel is the name for the Chassis.Axle.Row2.Wheel.Right.Brake.FluidLevel column in the Clickhouse.
	FieldChassisAxleRow2WheelRightBrakeFluidLevel string = "Chassis.Axle.Row2.Wheel.Right.Brake.FluidLevel"
	// FieldChassisAxleRow2WheelRightBrakeIsBrakesWorn is the name for the Chassis.Axle.Row2.Wheel.Right.Brake.IsBrakesWorn column in the Clickhouse.
	FieldChassisAxleRow2WheelRightBrakeIsBrakesWorn string = "Chassis.Axle.Row2.Wheel.Right.Brake.IsBrakesWorn"
	// FieldChassisAxleRow2WheelRightBrakeIsFluidLevelLow is the name for the Chassis.Axle.Row2.Wheel.Right.Brake.IsFluidLevelLow column in the Clickhouse.
	FieldChassisAxleRow2WheelRightBrakeIsFluidLevelLow string = "Chassis.Axle.Row2.Wheel.Right.Brake.IsFluidLevelLow"
	// FieldChassisAxleRow2WheelRightBrakePadWear is the name for the Chassis.Axle.Row2.Wheel.Right.Brake.PadWear column in the Clickhouse.
	FieldChassisAxleRow2WheelRightBrakePadWear string = "Chassis.Axle.Row2.Wheel.Right.Brake.PadWear"
	// FieldChassisAxleRow2WheelRightSpeed is the name for the Chassis.Axle.Row2.Wheel.Right.Speed column in the Clickhouse.
	FieldChassisAxleRow2WheelRightSpeed string = "Chassis.Axle.Row2.Wheel.Right.Speed"
	// FieldChassisAxleRow2WheelRightTire is the name for the Chassis.Axle.Row2.Wheel.Right.Tire column in the Clickhouse.
	FieldChassisAxleRow2WheelRightTire string = "Chassis.Axle.Row2.Wheel.Right.Tire"
	// FieldChassisAxleRow2WheelRightTireIsPressureLow is the name for the Chassis.Axle.Row2.Wheel.Right.Tire.IsPressureLow column in the Clickhouse.
	FieldChassisAxleRow2WheelRightTireIsPressureLow string = "Chassis.Axle.Row2.Wheel.Right.Tire.IsPressureLow"
	// FieldChassisAxleRow2WheelRightTirePressure is the name for the Chassis.Axle.Row2.Wheel.Right.Tire.Pressure column in the Clickhouse.
	FieldChassisAxleRow2WheelRightTirePressure string = "Chassis.Axle.Row2.Wheel.Right.Tire.Pressure"
	// FieldChassisAxleRow2WheelRightTireTemperature is the name for the Chassis.Axle.Row2.Wheel.Right.Tire.Temperature column in the Clickhouse.
	FieldChassisAxleRow2WheelRightTireTemperature string = "Chassis.Axle.Row2.Wheel.Right.Tire.Temperature"
	// FieldChassisAxleRow2WheelWidth is the name for the Chassis.Axle.Row2.WheelWidth column in the Clickhouse.
	FieldChassisAxleRow2WheelWidth string = "Chassis.Axle.Row2.WheelWidth"
	// FieldChassisBrake is the name for the Chassis.Brake column in the Clickhouse.
	FieldChassisBrake string = "Chassis.Brake"
	// FieldChassisBrakeIsDriverEmergencyBrakingDetected is the name for the Chassis.Brake.IsDriverEmergencyBrakingDetected column in the Clickhouse.
	FieldChassisBrakeIsDriverEmergencyBrakingDetected string = "Chassis.Brake.IsDriverEmergencyBrakingDetected"
	// FieldChassisBrakePedalPosition is the name for the Chassis.Brake.PedalPosition column in the Clickhouse.
	FieldChassisBrakePedalPosition string = "Chassis.Brake.PedalPosition"
	// FieldChassisParkingBrake is the name for the Chassis.ParkingBrake column in the Clickhouse.
	FieldChassisParkingBrake string = "Chassis.ParkingBrake"
	// FieldChassisParkingBrakeIsAutoApplyEnabled is the name for the Chassis.ParkingBrake.IsAutoApplyEnabled column in the Clickhouse.
	FieldChassisParkingBrakeIsAutoApplyEnabled string = "Chassis.ParkingBrake.IsAutoApplyEnabled"
	// FieldChassisParkingBrakeIsEngaged is the name for the Chassis.ParkingBrake.IsEngaged column in the Clickhouse.
	FieldChassisParkingBrakeIsEngaged string = "Chassis.ParkingBrake.IsEngaged"
	// FieldChassisSteeringWheel is the name for the Chassis.SteeringWheel column in the Clickhouse.
	FieldChassisSteeringWheel string = "Chassis.SteeringWheel"
	// FieldChassisSteeringWheelAngle is the name for the Chassis.SteeringWheel.Angle column in the Clickhouse.
	FieldChassisSteeringWheelAngle string = "Chassis.SteeringWheel.Angle"
	// FieldChassisSteeringWheelExtension is the name for the Chassis.SteeringWheel.Extension column in the Clickhouse.
	FieldChassisSteeringWheelExtension string = "Chassis.SteeringWheel.Extension"
	// FieldChassisSteeringWheelHeatingCooling is the name for the Chassis.SteeringWheel.HeatingCooling column in the Clickhouse.
	FieldChassisSteeringWheelHeatingCooling string = "Chassis.SteeringWheel.HeatingCooling"
	// FieldChassisSteeringWheelTilt is the name for the Chassis.SteeringWheel.Tilt column in the Clickhouse.
	FieldChassisSteeringWheelTilt string = "Chassis.SteeringWheel.Tilt"
	// FieldChassisWheelbase is the name for the Chassis.Wheelbase column in the Clickhouse.
	FieldChassisWheelbase string = "Chassis.Wheelbase"
	// FieldConnectivityIsConnectivityAvailable is the name for the Connectivity_IsConnectivityAvailable column in the Clickhouse.
	FieldConnectivityIsConnectivityAvailable string = "Connectivity_IsConnectivityAvailable"
	// FieldCurbWeight is the name for the CurbWeight column in the Clickhouse.
	FieldCurbWeight string = "CurbWeight"
	// FieldCurrentLocation is the name for the CurrentLocation column in the Clickhouse.
	FieldCurrentLocation string = "CurrentLocation"
	// FieldCurrentLocationAltitude is the name for the CurrentLocation.Altitude column in the Clickhouse.
	FieldCurrentLocationAltitude string = "CurrentLocation.Altitude"
	// FieldCurrentLocationGNSSReceiver is the name for the CurrentLocation.GNSSReceiver column in the Clickhouse.
	FieldCurrentLocationGNSSReceiver string = "CurrentLocation.GNSSReceiver"
	// FieldCurrentLocationGNSSReceiverFixType is the name for the CurrentLocation.GNSSReceiver.FixType column in the Clickhouse.
	FieldCurrentLocationGNSSReceiverFixType string = "CurrentLocation.GNSSReceiver.FixType"
	// FieldCurrentLocationGNSSReceiverMountingPosition is the name for the CurrentLocation.GNSSReceiver.MountingPosition column in the Clickhouse.
	FieldCurrentLocationGNSSReceiverMountingPosition string = "CurrentLocation.GNSSReceiver.MountingPosition"
	// FieldCurrentLocationGNSSReceiverMountingPositionX is the name for the CurrentLocation.GNSSReceiver.MountingPosition.X column in the Clickhouse.
	FieldCurrentLocationGNSSReceiverMountingPositionX string = "CurrentLocation.GNSSReceiver.MountingPosition.X"
	// FieldCurrentLocationGNSSReceiverMountingPositionY is the name for the CurrentLocation.GNSSReceiver.MountingPosition.Y column in the Clickhouse.
	FieldCurrentLocationGNSSReceiverMountingPositionY string = "CurrentLocation.GNSSReceiver.MountingPosition.Y"
	// FieldCurrentLocationGNSSReceiverMountingPositionZ is the name for the CurrentLocation.GNSSReceiver.MountingPosition.Z column in the Clickhouse.
	FieldCurrentLocationGNSSReceiverMountingPositionZ string = "CurrentLocation.GNSSReceiver.MountingPosition.Z"
	// FieldCurrentLocationHeading is the name for the CurrentLocation.Heading column in the Clickhouse.
	FieldCurrentLocationHeading string = "CurrentLocation.Heading"
	// FieldCurrentLocationHorizontalAccuracy is the name for the CurrentLocation.HorizontalAccuracy column in the Clickhouse.
	FieldCurrentLocationHorizontalAccuracy string = "CurrentLocation.HorizontalAccuracy"
	// FieldCurrentLocationLatitude is the name for the CurrentLocation.Latitude column in the Clickhouse.
	FieldCurrentLocationLatitude string = "CurrentLocation.Latitude"
	// FieldCurrentLocationLongitude is the name for the CurrentLocation.Longitude column in the Clickhouse.
	FieldCurrentLocationLongitude string = "CurrentLocation.Longitude"
	// FieldCurrentLocationTimestamp is the name for the CurrentLocation.Timestamp column in the Clickhouse.
	FieldCurrentLocationTimestamp string = "CurrentLocation.Timestamp"
	// FieldCurrentLocationVerticalAccuracy is the name for the CurrentLocation.VerticalAccuracy column in the Clickhouse.
	FieldCurrentLocationVerticalAccuracy string = "CurrentLocation.VerticalAccuracy"
	// FieldCurrentOverallWeight is the name for the CurrentOverallWeight column in the Clickhouse.
	FieldCurrentOverallWeight string = "CurrentOverallWeight"
	// FieldDIMO is the name for the DIMO column in the Clickhouse.
	FieldDIMO string = "DIMO"
	// FieldDIMODefinitionID is the name for the DIMO.DefinitionID column in the Clickhouse.
	FieldDIMODefinitionID string = "DIMO.DefinitionID"
	// FieldDIMOSource is the name for the DIMO.Source column in the Clickhouse.
	FieldDIMOSource string = "DIMO.Source"
	// FieldDIMOType is the name for the DIMO.Type column in the Clickhouse.
	FieldDIMOType string = "DIMO.Type"
	// FieldDIMOVehicleID is the name for the DIMO.VehicleID column in the Clickhouse.
	FieldDIMOVehicleID string = "DIMO.VehicleID"
	// FieldDriver is the name for the Driver column in the Clickhouse.
	FieldDriver string = "Driver"
	// FieldDriverAttentiveProbability is the name for the Driver.AttentiveProbability column in the Clickhouse.
	FieldDriverAttentiveProbability string = "Driver.AttentiveProbability"
	// FieldDriverDistractionLevel is the name for the Driver.DistractionLevel column in the Clickhouse.
	FieldDriverDistractionLevel string = "Driver.DistractionLevel"
	// FieldDriverFatigueLevel is the name for the Driver.FatigueLevel column in the Clickhouse.
	FieldDriverFatigueLevel string = "Driver.FatigueLevel"
	// FieldDriverHeartRate is the name for the Driver.HeartRate column in the Clickhouse.
	FieldDriverHeartRate string = "Driver.HeartRate"
	// FieldDriverIdentifier is the name for the Driver.Identifier column in the Clickhouse.
	FieldDriverIdentifier string = "Driver.Identifier"
	// FieldDriverIdentifierIssuer is the name for the Driver.Identifier.Issuer column in the Clickhouse.
	FieldDriverIdentifierIssuer string = "Driver.Identifier.Issuer"
	// FieldDriverIdentifierSubject is the name for the Driver.Identifier.Subject column in the Clickhouse.
	FieldDriverIdentifierSubject string = "Driver.Identifier.Subject"
	// FieldDriverIsEyesOnRoad is the name for the Driver.IsEyesOnRoad column in the Clickhouse.
	FieldDriverIsEyesOnRoad string = "Driver.IsEyesOnRoad"
	// FieldDriverIsHandsOnWheel is the name for the Driver.IsHandsOnWheel column in the Clickhouse.
	FieldDriverIsHandsOnWheel string = "Driver.IsHandsOnWheel"
	// FieldEmissionsCO2 is the name for the EmissionsCO2 column in the Clickhouse.
	FieldEmissionsCO2 string = "EmissionsCO2"
	// FieldExterior is the name for the Exterior column in the Clickhouse.
	FieldExterior string = "Exterior"
	// FieldExteriorAirTemperature is the name for the Exterior.AirTemperature column in the Clickhouse.
	FieldExteriorAirTemperature string = "Exterior.AirTemperature"
	// FieldExteriorHumidity is the name for the Exterior.Humidity column in the Clickhouse.
	FieldExteriorHumidity string = "Exterior.Humidity"
	// FieldExteriorLightIntensity is the name for the Exterior.LightIntensity column in the Clickhouse.
	FieldExteriorLightIntensity string = "Exterior.LightIntensity"
	// FieldGrossWeight is the name for the GrossWeight column in the Clickhouse.
	FieldGrossWeight string = "GrossWeight"
	// FieldHeight is the name for the Height column in the Clickhouse.
	FieldHeight string = "Height"
	// FieldIsAutoPowerOptimize is the name for the IsAutoPowerOptimize column in the Clickhouse.
	FieldIsAutoPowerOptimize string = "IsAutoPowerOptimize"
	// FieldIsBrokenDown is the name for the IsBrokenDown column in the Clickhouse.
	FieldIsBrokenDown string = "IsBrokenDown"
	// FieldIsMoving is the name for the IsMoving column in the Clickhouse.
	FieldIsMoving string = "IsMoving"
	// FieldLength is the name for the Length column in the Clickhouse.
	FieldLength string = "Length"
	// FieldLowVoltageBattery is the name for the LowVoltageBattery column in the Clickhouse.
	FieldLowVoltageBattery string = "LowVoltageBattery"
	// FieldLowVoltageBatteryCurrentCurrent is the name for the LowVoltageBattery.CurrentCurrent column in the Clickhouse.
	FieldLowVoltageBatteryCurrentCurrent string = "LowVoltageBattery.CurrentCurrent"
	// FieldLowVoltageBatteryCurrentVoltage is the name for the LowVoltageBattery.CurrentVoltage column in the Clickhouse.
	FieldLowVoltageBatteryCurrentVoltage string = "LowVoltageBattery.CurrentVoltage"
	// FieldLowVoltageBatteryNominalCapacity is the name for the LowVoltageBattery.NominalCapacity column in the Clickhouse.
	FieldLowVoltageBatteryNominalCapacity string = "LowVoltageBattery.NominalCapacity"
	// FieldLowVoltageBatteryNominalVoltage is the name for the LowVoltageBattery.NominalVoltage column in the Clickhouse.
	FieldLowVoltageBatteryNominalVoltage string = "LowVoltageBattery.NominalVoltage"
	// FieldLowVoltageSystemState is the name for the LowVoltageSystemState column in the Clickhouse.
	FieldLowVoltageSystemState string = "LowVoltageSystemState"
	// FieldMaxTowBallWeight is the name for the MaxTowBallWeight column in the Clickhouse.
	FieldMaxTowBallWeight string = "MaxTowBallWeight"
	// FieldMaxTowWeight is the name for the MaxTowWeight column in the Clickhouse.
	FieldMaxTowWeight string = "MaxTowWeight"
	// FieldOBD is the name for the OBD column in the Clickhouse.
	FieldOBD string = "OBD"
	// FieldOBDAbsoluteLoad is the name for the OBD.AbsoluteLoad column in the Clickhouse.
	FieldOBDAbsoluteLoad string = "OBD.AbsoluteLoad"
	// FieldOBDAcceleratorPositionD is the name for the OBD.AcceleratorPositionD column in the Clickhouse.
	FieldOBDAcceleratorPositionD string = "OBD.AcceleratorPositionD"
	// FieldOBDAcceleratorPositionE is the name for the OBD.AcceleratorPositionE column in the Clickhouse.
	FieldOBDAcceleratorPositionE string = "OBD.AcceleratorPositionE"
	// FieldOBDAcceleratorPositionF is the name for the OBD.AcceleratorPositionF column in the Clickhouse.
	FieldOBDAcceleratorPositionF string = "OBD.AcceleratorPositionF"
	// FieldOBDAirStatus is the name for the OBD.AirStatus column in the Clickhouse.
	FieldOBDAirStatus string = "OBD.AirStatus"
	// FieldOBDAmbientAirTemperature is the name for the OBD.AmbientAirTemperature column in the Clickhouse.
	FieldOBDAmbientAirTemperature string = "OBD.AmbientAirTemperature"
	// FieldOBDBarometricPressure is the name for the OBD.BarometricPressure column in the Clickhouse.
	FieldOBDBarometricPressure string = "OBD.BarometricPressure"
	// FieldOBDCatalyst is the name for the OBD.Catalyst column in the Clickhouse.
	FieldOBDCatalyst string = "OBD.Catalyst"
	// FieldOBDCatalystBank1 is the name for the OBD.Catalyst.Bank1 column in the Clickhouse.
	FieldOBDCatalystBank1 string = "OBD.Catalyst.Bank1"
	// FieldOBDCatalystBank1Temperature1 is the name for the OBD.Catalyst.Bank1.Temperature1 column in the Clickhouse.
	FieldOBDCatalystBank1Temperature1 string = "OBD.Catalyst.Bank1.Temperature1"
	// FieldOBDCatalystBank1Temperature2 is the name for the OBD.Catalyst.Bank1.Temperature2 column in the Clickhouse.
	FieldOBDCatalystBank1Temperature2 string = "OBD.Catalyst.Bank1.Temperature2"
	// FieldOBDCatalystBank2 is the name for the OBD.Catalyst.Bank2 column in the Clickhouse.
	FieldOBDCatalystBank2 string = "OBD.Catalyst.Bank2"
	// FieldOBDCatalystBank2Temperature1 is the name for the OBD.Catalyst.Bank2.Temperature1 column in the Clickhouse.
	FieldOBDCatalystBank2Temperature1 string = "OBD.Catalyst.Bank2.Temperature1"
	// FieldOBDCatalystBank2Temperature2 is the name for the OBD.Catalyst.Bank2.Temperature2 column in the Clickhouse.
	FieldOBDCatalystBank2Temperature2 string = "OBD.Catalyst.Bank2.Temperature2"
	// FieldOBDCommandedEGR is the name for the OBD.CommandedEGR column in the Clickhouse.
	FieldOBDCommandedEGR string = "OBD.CommandedEGR"
	// FieldOBDCommandedEVAP is the name for the OBD.CommandedEVAP column in the Clickhouse.
	FieldOBDCommandedEVAP string = "OBD.CommandedEVAP"
	// FieldOBDCommandedEquivalenceRatio is the name for the OBD.CommandedEquivalenceRatio column in the Clickhouse.
	FieldOBDCommandedEquivalenceRatio string = "OBD.CommandedEquivalenceRatio"
	// FieldOBDControlModuleVoltage is the name for the OBD.ControlModuleVoltage column in the Clickhouse.
	FieldOBDControlModuleVoltage string = "OBD.ControlModuleVoltage"
	// FieldOBDCoolantTemperature is the name for the OBD.CoolantTemperature column in the Clickhouse.
	FieldOBDCoolantTemperature string = "OBD.CoolantTemperature"
	// FieldOBDDTCList is the name for the OBD.DTCList column in the Clickhouse.
	FieldOBDDTCList string = "OBD.DTCList"
	// FieldOBDDTCListsize0 is the name for the OBD.DTCList.size0 column in the Clickhouse.
	FieldOBDDTCListsize0 string = "OBD.DTCList.size0"
	// FieldOBDDistanceSinceDTCClear is the name for the OBD.DistanceSinceDTCClear column in the Clickhouse.
	FieldOBDDistanceSinceDTCClear string = "OBD.DistanceSinceDTCClear"
	// FieldOBDDistanceWithMIL is the name for the OBD.DistanceWithMIL column in the Clickhouse.
	FieldOBDDistanceWithMIL string = "OBD.DistanceWithMIL"
	// FieldOBDDriveCycleStatus is the name for the OBD.DriveCycleStatus column in the Clickhouse.
	FieldOBDDriveCycleStatus string = "OBD.DriveCycleStatus"
	// FieldOBDDriveCycleStatusDTCCount is the name for the OBD.DriveCycleStatus.DTCCount column in the Clickhouse.
	FieldOBDDriveCycleStatusDTCCount string = "OBD.DriveCycleStatus.DTCCount"
	// FieldOBDDriveCycleStatusIgnitionType is the name for the OBD.DriveCycleStatus.IgnitionType column in the Clickhouse.
	FieldOBDDriveCycleStatusIgnitionType string = "OBD.DriveCycleStatus.IgnitionType"
	// FieldOBDDriveCycleStatusIsMILOn is the name for the OBD.DriveCycleStatus.IsMILOn column in the Clickhouse.
	FieldOBDDriveCycleStatusIsMILOn string = "OBD.DriveCycleStatus.IsMILOn"
	// FieldOBDEGRError is the name for the OBD.EGRError column in the Clickhouse.
	FieldOBDEGRError string = "OBD.EGRError"
	// FieldOBDEVAPVaporPressure is the name for the OBD.EVAPVaporPressure column in the Clickhouse.
	FieldOBDEVAPVaporPressure string = "OBD.EVAPVaporPressure"
	// FieldOBDEVAPVaporPressureAbsolute is the name for the OBD.EVAPVaporPressureAbsolute column in the Clickhouse.
	FieldOBDEVAPVaporPressureAbsolute string = "OBD.EVAPVaporPressureAbsolute"
	// FieldOBDEVAPVaporPressureAlternate is the name for the OBD.EVAPVaporPressureAlternate column in the Clickhouse.
	FieldOBDEVAPVaporPressureAlternate string = "OBD.EVAPVaporPressureAlternate"
	// FieldOBDEngineLoad is the name for the OBD.EngineLoad column in the Clickhouse.
	FieldOBDEngineLoad string = "OBD.EngineLoad"
	// FieldOBDEngineSpeed is the name for the OBD.EngineSpeed column in the Clickhouse.
	FieldOBDEngineSpeed string = "OBD.EngineSpeed"
	// FieldOBDEthanolPercent is the name for the OBD.EthanolPercent column in the Clickhouse.
	FieldOBDEthanolPercent string = "OBD.EthanolPercent"
	// FieldOBDFreezeDTC is the name for the OBD.FreezeDTC column in the Clickhouse.
	FieldOBDFreezeDTC string = "OBD.FreezeDTC"
	// FieldOBDFuelInjectionTiming is the name for the OBD.FuelInjectionTiming column in the Clickhouse.
	FieldOBDFuelInjectionTiming string = "OBD.FuelInjectionTiming"
	// FieldOBDFuelLevel is the name for the OBD.FuelLevel column in the Clickhouse.
	FieldOBDFuelLevel string = "OBD.FuelLevel"
	// FieldOBDFuelPressure is the name for the OBD.FuelPressure column in the Clickhouse.
	FieldOBDFuelPressure string = "OBD.FuelPressure"
	// FieldOBDFuelRailPressureAbsolute is the name for the OBD.FuelRailPressureAbsolute column in the Clickhouse.
	FieldOBDFuelRailPressureAbsolute string = "OBD.FuelRailPressureAbsolute"
	// FieldOBDFuelRailPressureDirect is the name for the OBD.FuelRailPressureDirect column in the Clickhouse.
	FieldOBDFuelRailPressureDirect string = "OBD.FuelRailPressureDirect"
	// FieldOBDFuelRailPressureVac is the name for the OBD.FuelRailPressureVac column in the Clickhouse.
	FieldOBDFuelRailPressureVac string = "OBD.FuelRailPressureVac"
	// FieldOBDFuelRate is the name for the OBD.FuelRate column in the Clickhouse.
	FieldOBDFuelRate string = "OBD.FuelRate"
	// FieldOBDFuelStatus is the name for the OBD.FuelStatus column in the Clickhouse.
	FieldOBDFuelStatus string = "OBD.FuelStatus"
	// FieldOBDFuelType is the name for the OBD.FuelType column in the Clickhouse.
	FieldOBDFuelType string = "OBD.FuelType"
	// FieldOBDHybridBatteryRemaining is the name for the OBD.HybridBatteryRemaining column in the Clickhouse.
	FieldOBDHybridBatteryRemaining string = "OBD.HybridBatteryRemaining"
	// FieldOBDIntakeTemp is the name for the OBD.IntakeTemp column in the Clickhouse.
	FieldOBDIntakeTemp string = "OBD.IntakeTemp"
	// FieldOBDIsPTOActive is the name for the OBD.IsPTOActive column in the Clickhouse.
	FieldOBDIsPTOActive string = "OBD.IsPTOActive"
	// FieldOBDLongTermFuelTrim1 is the name for the OBD.LongTermFuelTrim1 column in the Clickhouse.
	FieldOBDLongTermFuelTrim1 string = "OBD.LongTermFuelTrim1"
	// FieldOBDLongTermFuelTrim2 is the name for the OBD.LongTermFuelTrim2 column in the Clickhouse.
	FieldOBDLongTermFuelTrim2 string = "OBD.LongTermFuelTrim2"
	// FieldOBDLongTermO2Trim1 is the name for the OBD.LongTermO2Trim1 column in the Clickhouse.
	FieldOBDLongTermO2Trim1 string = "OBD.LongTermO2Trim1"
	// FieldOBDLongTermO2Trim2 is the name for the OBD.LongTermO2Trim2 column in the Clickhouse.
	FieldOBDLongTermO2Trim2 string = "OBD.LongTermO2Trim2"
	// FieldOBDLongTermO2Trim3 is the name for the OBD.LongTermO2Trim3 column in the Clickhouse.
	FieldOBDLongTermO2Trim3 string = "OBD.LongTermO2Trim3"
	// FieldOBDLongTermO2Trim4 is the name for the OBD.LongTermO2Trim4 column in the Clickhouse.
	FieldOBDLongTermO2Trim4 string = "OBD.LongTermO2Trim4"
	// FieldOBDMAF is the name for the OBD.MAF column in the Clickhouse.
	FieldOBDMAF string = "OBD.MAF"
	// FieldOBDMAP is the name for the OBD.MAP column in the Clickhouse.
	FieldOBDMAP string = "OBD.MAP"
	// FieldOBDMaxMAF is the name for the OBD.MaxMAF column in the Clickhouse.
	FieldOBDMaxMAF string = "OBD.MaxMAF"
	// FieldOBDO2 is the name for the OBD.O2 column in the Clickhouse.
	FieldOBDO2 string = "OBD.O2"
	// FieldOBDO2Sensor1 is the name for the OBD.O2.Sensor1 column in the Clickhouse.
	FieldOBDO2Sensor1 string = "OBD.O2.Sensor1"
	// FieldOBDO2Sensor1ShortTermFuelTrim is the name for the OBD.O2.Sensor1.ShortTermFuelTrim column in the Clickhouse.
	FieldOBDO2Sensor1ShortTermFuelTrim string = "OBD.O2.Sensor1.ShortTermFuelTrim"
	// FieldOBDO2Sensor1Voltage is the name for the OBD.O2.Sensor1.Voltage column in the Clickhouse.
	FieldOBDO2Sensor1Voltage string = "OBD.O2.Sensor1.Voltage"
	// FieldOBDO2Sensor2 is the name for the OBD.O2.Sensor2 column in the Clickhouse.
	FieldOBDO2Sensor2 string = "OBD.O2.Sensor2"
	// FieldOBDO2Sensor2ShortTermFuelTrim is the name for the OBD.O2.Sensor2.ShortTermFuelTrim column in the Clickhouse.
	FieldOBDO2Sensor2ShortTermFuelTrim string = "OBD.O2.Sensor2.ShortTermFuelTrim"
	// FieldOBDO2Sensor2Voltage is the name for the OBD.O2.Sensor2.Voltage column in the Clickhouse.
	FieldOBDO2Sensor2Voltage string = "OBD.O2.Sensor2.Voltage"
	// FieldOBDO2Sensor3 is the name for the OBD.O2.Sensor3 column in the Clickhouse.
	FieldOBDO2Sensor3 string = "OBD.O2.Sensor3"
	// FieldOBDO2Sensor3ShortTermFuelTrim is the name for the OBD.O2.Sensor3.ShortTermFuelTrim column in the Clickhouse.
	FieldOBDO2Sensor3ShortTermFuelTrim string = "OBD.O2.Sensor3.ShortTermFuelTrim"
	// FieldOBDO2Sensor3Voltage is the name for the OBD.O2.Sensor3.Voltage column in the Clickhouse.
	FieldOBDO2Sensor3Voltage string = "OBD.O2.Sensor3.Voltage"
	// FieldOBDO2Sensor4 is the name for the OBD.O2.Sensor4 column in the Clickhouse.
	FieldOBDO2Sensor4 string = "OBD.O2.Sensor4"
	// FieldOBDO2Sensor4ShortTermFuelTrim is the name for the OBD.O2.Sensor4.ShortTermFuelTrim column in the Clickhouse.
	FieldOBDO2Sensor4ShortTermFuelTrim string = "OBD.O2.Sensor4.ShortTermFuelTrim"
	// FieldOBDO2Sensor4Voltage is the name for the OBD.O2.Sensor4.Voltage column in the Clickhouse.
	FieldOBDO2Sensor4Voltage string = "OBD.O2.Sensor4.Voltage"
	// FieldOBDO2Sensor5 is the name for the OBD.O2.Sensor5 column in the Clickhouse.
	FieldOBDO2Sensor5 string = "OBD.O2.Sensor5"
	// FieldOBDO2Sensor5ShortTermFuelTrim is the name for the OBD.O2.Sensor5.ShortTermFuelTrim column in the Clickhouse.
	FieldOBDO2Sensor5ShortTermFuelTrim string = "OBD.O2.Sensor5.ShortTermFuelTrim"
	// FieldOBDO2Sensor5Voltage is the name for the OBD.O2.Sensor5.Voltage column in the Clickhouse.
	FieldOBDO2Sensor5Voltage string = "OBD.O2.Sensor5.Voltage"
	// FieldOBDO2Sensor6 is the name for the OBD.O2.Sensor6 column in the Clickhouse.
	FieldOBDO2Sensor6 string = "OBD.O2.Sensor6"
	// FieldOBDO2Sensor6ShortTermFuelTrim is the name for the OBD.O2.Sensor6.ShortTermFuelTrim column in the Clickhouse.
	FieldOBDO2Sensor6ShortTermFuelTrim string = "OBD.O2.Sensor6.ShortTermFuelTrim"
	// FieldOBDO2Sensor6Voltage is the name for the OBD.O2.Sensor6.Voltage column in the Clickhouse.
	FieldOBDO2Sensor6Voltage string = "OBD.O2.Sensor6.Voltage"
	// FieldOBDO2Sensor7 is the name for the OBD.O2.Sensor7 column in the Clickhouse.
	FieldOBDO2Sensor7 string = "OBD.O2.Sensor7"
	// FieldOBDO2Sensor7ShortTermFuelTrim is the name for the OBD.O2.Sensor7.ShortTermFuelTrim column in the Clickhouse.
	FieldOBDO2Sensor7ShortTermFuelTrim string = "OBD.O2.Sensor7.ShortTermFuelTrim"
	// FieldOBDO2Sensor7Voltage is the name for the OBD.O2.Sensor7.Voltage column in the Clickhouse.
	FieldOBDO2Sensor7Voltage string = "OBD.O2.Sensor7.Voltage"
	// FieldOBDO2Sensor8 is the name for the OBD.O2.Sensor8 column in the Clickhouse.
	FieldOBDO2Sensor8 string = "OBD.O2.Sensor8"
	// FieldOBDO2Sensor8ShortTermFuelTrim is the name for the OBD.O2.Sensor8.ShortTermFuelTrim column in the Clickhouse.
	FieldOBDO2Sensor8ShortTermFuelTrim string = "OBD.O2.Sensor8.ShortTermFuelTrim"
	// FieldOBDO2Sensor8Voltage is the name for the OBD.O2.Sensor8.Voltage column in the Clickhouse.
	FieldOBDO2Sensor8Voltage string = "OBD.O2.Sensor8.Voltage"
	// FieldOBDO2WR is the name for the OBD.O2WR column in the Clickhouse.
	FieldOBDO2WR string = "OBD.O2WR"
	// FieldOBDO2WRSensor1 is the name for the OBD.O2WR.Sensor1 column in the Clickhouse.
	FieldOBDO2WRSensor1 string = "OBD.O2WR.Sensor1"
	// FieldOBDO2WRSensor1Current is the name for the OBD.O2WR.Sensor1.Current column in the Clickhouse.
	FieldOBDO2WRSensor1Current string = "OBD.O2WR.Sensor1.Current"
	// FieldOBDO2WRSensor1Lambda is the name for the OBD.O2WR.Sensor1.Lambda column in the Clickhouse.
	FieldOBDO2WRSensor1Lambda string = "OBD.O2WR.Sensor1.Lambda"
	// FieldOBDO2WRSensor1Voltage is the name for the OBD.O2WR.Sensor1.Voltage column in the Clickhouse.
	FieldOBDO2WRSensor1Voltage string = "OBD.O2WR.Sensor1.Voltage"
	// FieldOBDO2WRSensor2 is the name for the OBD.O2WR.Sensor2 column in the Clickhouse.
	FieldOBDO2WRSensor2 string = "OBD.O2WR.Sensor2"
	// FieldOBDO2WRSensor2Current is the name for the OBD.O2WR.Sensor2.Current column in the Clickhouse.
	FieldOBDO2WRSensor2Current string = "OBD.O2WR.Sensor2.Current"
	// FieldOBDO2WRSensor2Lambda is the name for the OBD.O2WR.Sensor2.Lambda column in the Clickhouse.
	FieldOBDO2WRSensor2Lambda string = "OBD.O2WR.Sensor2.Lambda"
	// FieldOBDO2WRSensor2Voltage is the name for the OBD.O2WR.Sensor2.Voltage column in the Clickhouse.
	FieldOBDO2WRSensor2Voltage string = "OBD.O2WR.Sensor2.Voltage"
	// FieldOBDO2WRSensor3 is the name for the OBD.O2WR.Sensor3 column in the Clickhouse.
	FieldOBDO2WRSensor3 string = "OBD.O2WR.Sensor3"
	// FieldOBDO2WRSensor3Current is the name for the OBD.O2WR.Sensor3.Current column in the Clickhouse.
	FieldOBDO2WRSensor3Current string = "OBD.O2WR.Sensor3.Current"
	// FieldOBDO2WRSensor3Lambda is the name for the OBD.O2WR.Sensor3.Lambda column in the Clickhouse.
	FieldOBDO2WRSensor3Lambda string = "OBD.O2WR.Sensor3.Lambda"
	// FieldOBDO2WRSensor3Voltage is the name for the OBD.O2WR.Sensor3.Voltage column in the Clickhouse.
	FieldOBDO2WRSensor3Voltage string = "OBD.O2WR.Sensor3.Voltage"
	// FieldOBDO2WRSensor4 is the name for the OBD.O2WR.Sensor4 column in the Clickhouse.
	FieldOBDO2WRSensor4 string = "OBD.O2WR.Sensor4"
	// FieldOBDO2WRSensor4Current is the name for the OBD.O2WR.Sensor4.Current column in the Clickhouse.
	FieldOBDO2WRSensor4Current string = "OBD.O2WR.Sensor4.Current"
	// FieldOBDO2WRSensor4Lambda is the name for the OBD.O2WR.Sensor4.Lambda column in the Clickhouse.
	FieldOBDO2WRSensor4Lambda string = "OBD.O2WR.Sensor4.Lambda"
	// FieldOBDO2WRSensor4Voltage is the name for the OBD.O2WR.Sensor4.Voltage column in the Clickhouse.
	FieldOBDO2WRSensor4Voltage string = "OBD.O2WR.Sensor4.Voltage"
	// FieldOBDO2WRSensor5 is the name for the OBD.O2WR.Sensor5 column in the Clickhouse.
	FieldOBDO2WRSensor5 string = "OBD.O2WR.Sensor5"
	// FieldOBDO2WRSensor5Current is the name for the OBD.O2WR.Sensor5.Current column in the Clickhouse.
	FieldOBDO2WRSensor5Current string = "OBD.O2WR.Sensor5.Current"
	// FieldOBDO2WRSensor5Lambda is the name for the OBD.O2WR.Sensor5.Lambda column in the Clickhouse.
	FieldOBDO2WRSensor5Lambda string = "OBD.O2WR.Sensor5.Lambda"
	// FieldOBDO2WRSensor5Voltage is the name for the OBD.O2WR.Sensor5.Voltage column in the Clickhouse.
	FieldOBDO2WRSensor5Voltage string = "OBD.O2WR.Sensor5.Voltage"
	// FieldOBDO2WRSensor6 is the name for the OBD.O2WR.Sensor6 column in the Clickhouse.
	FieldOBDO2WRSensor6 string = "OBD.O2WR.Sensor6"
	// FieldOBDO2WRSensor6Current is the name for the OBD.O2WR.Sensor6.Current column in the Clickhouse.
	FieldOBDO2WRSensor6Current string = "OBD.O2WR.Sensor6.Current"
	// FieldOBDO2WRSensor6Lambda is the name for the OBD.O2WR.Sensor6.Lambda column in the Clickhouse.
	FieldOBDO2WRSensor6Lambda string = "OBD.O2WR.Sensor6.Lambda"
	// FieldOBDO2WRSensor6Voltage is the name for the OBD.O2WR.Sensor6.Voltage column in the Clickhouse.
	FieldOBDO2WRSensor6Voltage string = "OBD.O2WR.Sensor6.Voltage"
	// FieldOBDO2WRSensor7 is the name for the OBD.O2WR.Sensor7 column in the Clickhouse.
	FieldOBDO2WRSensor7 string = "OBD.O2WR.Sensor7"
	// FieldOBDO2WRSensor7Current is the name for the OBD.O2WR.Sensor7.Current column in the Clickhouse.
	FieldOBDO2WRSensor7Current string = "OBD.O2WR.Sensor7.Current"
	// FieldOBDO2WRSensor7Lambda is the name for the OBD.O2WR.Sensor7.Lambda column in the Clickhouse.
	FieldOBDO2WRSensor7Lambda string = "OBD.O2WR.Sensor7.Lambda"
	// FieldOBDO2WRSensor7Voltage is the name for the OBD.O2WR.Sensor7.Voltage column in the Clickhouse.
	FieldOBDO2WRSensor7Voltage string = "OBD.O2WR.Sensor7.Voltage"
	// FieldOBDO2WRSensor8 is the name for the OBD.O2WR.Sensor8 column in the Clickhouse.
	FieldOBDO2WRSensor8 string = "OBD.O2WR.Sensor8"
	// FieldOBDO2WRSensor8Current is the name for the OBD.O2WR.Sensor8.Current column in the Clickhouse.
	FieldOBDO2WRSensor8Current string = "OBD.O2WR.Sensor8.Current"
	// FieldOBDO2WRSensor8Lambda is the name for the OBD.O2WR.Sensor8.Lambda column in the Clickhouse.
	FieldOBDO2WRSensor8Lambda string = "OBD.O2WR.Sensor8.Lambda"
	// FieldOBDO2WRSensor8Voltage is the name for the OBD.O2WR.Sensor8.Voltage column in the Clickhouse.
	FieldOBDO2WRSensor8Voltage string = "OBD.O2WR.Sensor8.Voltage"
	// FieldOBDOBDStandards is the name for the OBD.OBDStandards column in the Clickhouse.
	FieldOBDOBDStandards string = "OBD.OBDStandards"
	// FieldOBDOilTemperature is the name for the OBD.OilTemperature column in the Clickhouse.
	FieldOBDOilTemperature string = "OBD.OilTemperature"
	// FieldOBDOxygenSensorsIn2Banks is the name for the OBD.OxygenSensorsIn2Banks column in the Clickhouse.
	FieldOBDOxygenSensorsIn2Banks string = "OBD.OxygenSensorsIn2Banks"
	// FieldOBDOxygenSensorsIn4Banks is the name for the OBD.OxygenSensorsIn4Banks column in the Clickhouse.
	FieldOBDOxygenSensorsIn4Banks string = "OBD.OxygenSensorsIn4Banks"
	// FieldOBDPidsA is the name for the OBD.PidsA column in the Clickhouse.
	FieldOBDPidsA string = "OBD.PidsA"
	// FieldOBDPidsAsize0 is the name for the OBD.PidsA.size0 column in the Clickhouse.
	FieldOBDPidsAsize0 string = "OBD.PidsA.size0"
	// FieldOBDPidsB is the name for the OBD.PidsB column in the Clickhouse.
	FieldOBDPidsB string = "OBD.PidsB"
	// FieldOBDPidsBsize0 is the name for the OBD.PidsB.size0 column in the Clickhouse.
	FieldOBDPidsBsize0 string = "OBD.PidsB.size0"
	// FieldOBDPidsC is the name for the OBD.PidsC column in the Clickhouse.
	FieldOBDPidsC string = "OBD.PidsC"
	// FieldOBDPidsCsize0 is the name for the OBD.PidsC.size0 column in the Clickhouse.
	FieldOBDPidsCsize0 string = "OBD.PidsC.size0"
	// FieldOBDRelativeAcceleratorPosition is the name for the OBD.RelativeAcceleratorPosition column in the Clickhouse.
	FieldOBDRelativeAcceleratorPosition string = "OBD.RelativeAcceleratorPosition"
	// FieldOBDRelativeThrottlePosition is the name for the OBD.RelativeThrottlePosition column in the Clickhouse.
	FieldOBDRelativeThrottlePosition string = "OBD.RelativeThrottlePosition"
	// FieldOBDRunTime is the name for the OBD.RunTime column in the Clickhouse.
	FieldOBDRunTime string = "OBD.RunTime"
	// FieldOBDRunTimeMIL is the name for the OBD.RunTimeMIL column in the Clickhouse.
	FieldOBDRunTimeMIL string = "OBD.RunTimeMIL"
	// FieldOBDShortTermFuelTrim1 is the name for the OBD.ShortTermFuelTrim1 column in the Clickhouse.
	FieldOBDShortTermFuelTrim1 string = "OBD.ShortTermFuelTrim1"
	// FieldOBDShortTermFuelTrim2 is the name for the OBD.ShortTermFuelTrim2 column in the Clickhouse.
	FieldOBDShortTermFuelTrim2 string = "OBD.ShortTermFuelTrim2"
	// FieldOBDShortTermO2Trim1 is the name for the OBD.ShortTermO2Trim1 column in the Clickhouse.
	FieldOBDShortTermO2Trim1 string = "OBD.ShortTermO2Trim1"
	// FieldOBDShortTermO2Trim2 is the name for the OBD.ShortTermO2Trim2 column in the Clickhouse.
	FieldOBDShortTermO2Trim2 string = "OBD.ShortTermO2Trim2"
	// FieldOBDShortTermO2Trim3 is the name for the OBD.ShortTermO2Trim3 column in the Clickhouse.
	FieldOBDShortTermO2Trim3 string = "OBD.ShortTermO2Trim3"
	// FieldOBDShortTermO2Trim4 is the name for the OBD.ShortTermO2Trim4 column in the Clickhouse.
	FieldOBDShortTermO2Trim4 string = "OBD.ShortTermO2Trim4"
	// FieldOBDSpeed is the name for the OBD.Speed column in the Clickhouse.
	FieldOBDSpeed string = "OBD.Speed"
	// FieldOBDStatus is the name for the OBD.Status column in the Clickhouse.
	FieldOBDStatus string = "OBD.Status"
	// FieldOBDStatusDTCCount is the name for the OBD.Status.DTCCount column in the Clickhouse.
	FieldOBDStatusDTCCount string = "OBD.Status.DTCCount"
	// FieldOBDStatusIgnitionType is the name for the OBD.Status.IgnitionType column in the Clickhouse.
	FieldOBDStatusIgnitionType string = "OBD.Status.IgnitionType"
	// FieldOBDStatusIsMILOn is the name for the OBD.Status.IsMILOn column in the Clickhouse.
	FieldOBDStatusIsMILOn string = "OBD.Status.IsMILOn"
	// FieldOBDThrottleActuator is the name for the OBD.ThrottleActuator column in the Clickhouse.
	FieldOBDThrottleActuator string = "OBD.ThrottleActuator"
	// FieldOBDThrottlePosition is the name for the OBD.ThrottlePosition column in the Clickhouse.
	FieldOBDThrottlePosition string = "OBD.ThrottlePosition"
	// FieldOBDThrottlePositionB is the name for the OBD.ThrottlePositionB column in the Clickhouse.
	FieldOBDThrottlePositionB string = "OBD.ThrottlePositionB"
	// FieldOBDThrottlePositionC is the name for the OBD.ThrottlePositionC column in the Clickhouse.
	FieldOBDThrottlePositionC string = "OBD.ThrottlePositionC"
	// FieldOBDTimeSinceDTCCleared is the name for the OBD.TimeSinceDTCCleared column in the Clickhouse.
	FieldOBDTimeSinceDTCCleared string = "OBD.TimeSinceDTCCleared"
	// FieldOBDTimingAdvance is the name for the OBD.TimingAdvance column in the Clickhouse.
	FieldOBDTimingAdvance string = "OBD.TimingAdvance"
	// FieldOBDWarmupsSinceDTCClear is the name for the OBD.WarmupsSinceDTCClear column in the Clickhouse.
	FieldOBDWarmupsSinceDTCClear string = "OBD.WarmupsSinceDTCClear"
	// FieldPowerOptimizeLevel is the name for the PowerOptimizeLevel column in the Clickhouse.
	FieldPowerOptimizeLevel string = "PowerOptimizeLevel"
	// FieldPowertrain is the name for the Powertrain column in the Clickhouse.
	FieldPowertrain string = "Powertrain"
	// FieldPowertrainAccumulatedBrakingEnergy is the name for the Powertrain.AccumulatedBrakingEnergy column in the Clickhouse.
	FieldPowertrainAccumulatedBrakingEnergy string = "Powertrain.AccumulatedBrakingEnergy"
	// FieldPowertrainCombustionEngine is the name for the Powertrain.CombustionEngine column in the Clickhouse.
	FieldPowertrainCombustionEngine string = "Powertrain.CombustionEngine"
	// FieldPowertrainCombustionEngineAspirationType is the name for the Powertrain.CombustionEngine.AspirationType column in the Clickhouse.
	FieldPowertrainCombustionEngineAspirationType string = "Powertrain.CombustionEngine.AspirationType"
	// FieldPowertrainCombustionEngineBore is the name for the Powertrain.CombustionEngine.Bore column in the Clickhouse.
	FieldPowertrainCombustionEngineBore string = "Powertrain.CombustionEngine.Bore"
	// FieldPowertrainCombustionEngineCompressionRatio is the name for the Powertrain.CombustionEngine.CompressionRatio column in the Clickhouse.
	FieldPowertrainCombustionEngineCompressionRatio string = "Powertrain.CombustionEngine.CompressionRatio"
	// FieldPowertrainCombustionEngineConfiguration is the name for the Powertrain.CombustionEngine.Configuration column in the Clickhouse.
	FieldPowertrainCombustionEngineConfiguration string = "Powertrain.CombustionEngine.Configuration"
	// FieldPowertrainCombustionEngineDieselExhaustFluid is the name for the Powertrain.CombustionEngine.DieselExhaustFluid column in the Clickhouse.
	FieldPowertrainCombustionEngineDieselExhaustFluid string = "Powertrain.CombustionEngine.DieselExhaustFluid"
	// FieldPowertrainCombustionEngineDieselExhaustFluidCapacity is the name for the Powertrain.CombustionEngine.DieselExhaustFluid.Capacity column in the Clickhouse.
	FieldPowertrainCombustionEngineDieselExhaustFluidCapacity string = "Powertrain.CombustionEngine.DieselExhaustFluid.Capacity"
	// FieldPowertrainCombustionEngineDieselExhaustFluidIsLevelLow is the name for the Powertrain.CombustionEngine.DieselExhaustFluid.IsLevelLow column in the Clickhouse.
	FieldPowertrainCombustionEngineDieselExhaustFluidIsLevelLow string = "Powertrain.CombustionEngine.DieselExhaustFluid.IsLevelLow"
	// FieldPowertrainCombustionEngineDieselExhaustFluidLevel is the name for the Powertrain.CombustionEngine.DieselExhaustFluid.Level column in the Clickhouse.
	FieldPowertrainCombustionEngineDieselExhaustFluidLevel string = "Powertrain.CombustionEngine.DieselExhaustFluid.Level"
	// FieldPowertrainCombustionEngineDieselExhaustFluidRange is the name for the Powertrain.CombustionEngine.DieselExhaustFluid.Range column in the Clickhouse.
	FieldPowertrainCombustionEngineDieselExhaustFluidRange string = "Powertrain.CombustionEngine.DieselExhaustFluid.Range"
	// FieldPowertrainCombustionEngineDieselParticulateFilter is the name for the Powertrain.CombustionEngine.DieselParticulateFilter column in the Clickhouse.
	FieldPowertrainCombustionEngineDieselParticulateFilter string = "Powertrain.CombustionEngine.DieselParticulateFilter"
	// FieldPowertrainCombustionEngineDieselParticulateFilterDeltaPressure is the name for the Powertrain.CombustionEngine.DieselParticulateFilter.DeltaPressure column in the Clickhouse.
	FieldPowertrainCombustionEngineDieselParticulateFilterDeltaPressure string = "Powertrain.CombustionEngine.DieselParticulateFilter.DeltaPressure"
	// FieldPowertrainCombustionEngineDieselParticulateFilterInletTemperature is the name for the Powertrain.CombustionEngine.DieselParticulateFilter.InletTemperature column in the Clickhouse.
	FieldPowertrainCombustionEngineDieselParticulateFilterInletTemperature string = "Powertrain.CombustionEngine.DieselParticulateFilter.InletTemperature"
	// FieldPowertrainCombustionEngineDieselParticulateFilterOutletTemperature is the name for the Powertrain.CombustionEngine.DieselParticulateFilter.OutletTemperature column in the Clickhouse.
	FieldPowertrainCombustionEngineDieselParticulateFilterOutletTemperature string = "Powertrain.CombustionEngine.DieselParticulateFilter.OutletTemperature"
	// FieldPowertrainCombustionEngineDisplacement is the name for the Powertrain.CombustionEngine.Displacement column in the Clickhouse.
	FieldPowertrainCombustionEngineDisplacement string = "Powertrain.CombustionEngine.Displacement"
	// FieldPowertrainCombustionEngineECT is the name for the Powertrain.CombustionEngine.ECT column in the Clickhouse.
	FieldPowertrainCombustionEngineECT string = "Powertrain.CombustionEngine.ECT"
	// FieldPowertrainCombustionEngineEOP is the name for the Powertrain.CombustionEngine.EOP column in the Clickhouse.
	FieldPowertrainCombustionEngineEOP string = "Powertrain.CombustionEngine.EOP"
	// FieldPowertrainCombustionEngineEOT is the name for the Powertrain.CombustionEngine.EOT column in the Clickhouse.
	FieldPowertrainCombustionEngineEOT string = "Powertrain.CombustionEngine.EOT"
	// FieldPowertrainCombustionEngineEngineCode is the name for the Powertrain.CombustionEngine.EngineCode column in the Clickhouse.
	FieldPowertrainCombustionEngineEngineCode string = "Powertrain.CombustionEngine.EngineCode"
	// FieldPowertrainCombustionEngineEngineCoolantCapacity is the name for the Powertrain.CombustionEngine.EngineCoolantCapacity column in the Clickhouse.
	FieldPowertrainCombustionEngineEngineCoolantCapacity string = "Powertrain.CombustionEngine.EngineCoolantCapacity"
	// FieldPowertrainCombustionEngineEngineHours is the name for the Powertrain.CombustionEngine.EngineHours column in the Clickhouse.
	FieldPowertrainCombustionEngineEngineHours string = "Powertrain.CombustionEngine.EngineHours"
	// FieldPowertrainCombustionEngineEngineOilCapacity is the name for the Powertrain.CombustionEngine.EngineOilCapacity column in the Clickhouse.
	FieldPowertrainCombustionEngineEngineOilCapacity string = "Powertrain.CombustionEngine.EngineOilCapacity"
	// FieldPowertrainCombustionEngineEngineOilLevel is the name for the Powertrain.CombustionEngine.EngineOilLevel column in the Clickhouse.
	FieldPowertrainCombustionEngineEngineOilLevel string = "Powertrain.CombustionEngine.EngineOilLevel"
	// FieldPowertrainCombustionEngineIdleHours is the name for the Powertrain.CombustionEngine.IdleHours column in the Clickhouse.
	FieldPowertrainCombustionEngineIdleHours string = "Powertrain.CombustionEngine.IdleHours"
	// FieldPowertrainCombustionEngineIsRunning is the name for the Powertrain.CombustionEngine.IsRunning column in the Clickhouse.
	FieldPowertrainCombustionEngineIsRunning string = "Powertrain.CombustionEngine.IsRunning"
	// FieldPowertrainCombustionEngineMAF is the name for the Powertrain.CombustionEngine.MAF column in the Clickhouse.
	FieldPowertrainCombustionEngineMAF string = "Powertrain.CombustionEngine.MAF"
	// FieldPowertrainCombustionEngineMAP is the name for the Powertrain.CombustionEngine.MAP column in the Clickhouse.
	FieldPowertrainCombustionEngineMAP string = "Powertrain.CombustionEngine.MAP"
	// FieldPowertrainCombustionEngineMaxPower is the name for the Powertrain.CombustionEngine.MaxPower column in the Clickhouse.
	FieldPowertrainCombustionEngineMaxPower string = "Powertrain.CombustionEngine.MaxPower"
	// FieldPowertrainCombustionEngineMaxTorque is the name for the Powertrain.CombustionEngine.MaxTorque column in the Clickhouse.
	FieldPowertrainCombustionEngineMaxTorque string = "Powertrain.CombustionEngine.MaxTorque"
	// FieldPowertrainCombustionEngineNumberOfCylinders is the name for the Powertrain.CombustionEngine.NumberOfCylinders column in the Clickhouse.
	FieldPowertrainCombustionEngineNumberOfCylinders string = "Powertrain.CombustionEngine.NumberOfCylinders"
	// FieldPowertrainCombustionEngineNumberOfValvesPerCylinder is the name for the Powertrain.CombustionEngine.NumberOfValvesPerCylinder column in the Clickhouse.
	FieldPowertrainCombustionEngineNumberOfValvesPerCylinder string = "Powertrain.CombustionEngine.NumberOfValvesPerCylinder"
	// FieldPowertrainCombustionEngineOilLifeRemaining is the name for the Powertrain.CombustionEngine.OilLifeRemaining column in the Clickhouse.
	FieldPowertrainCombustionEngineOilLifeRemaining string = "Powertrain.CombustionEngine.OilLifeRemaining"
	// FieldPowertrainCombustionEnginePower is the name for the Powertrain.CombustionEngine.Power column in the Clickhouse.
	FieldPowertrainCombustionEnginePower string = "Powertrain.CombustionEngine.Power"
	// FieldPowertrainCombustionEngineSpeed is the name for the Powertrain.CombustionEngine.Speed column in the Clickhouse.
	FieldPowertrainCombustionEngineSpeed string = "Powertrain.CombustionEngine.Speed"
	// FieldPowertrainCombustionEngineStrokeLength is the name for the Powertrain.CombustionEngine.StrokeLength column in the Clickhouse.
	FieldPowertrainCombustionEngineStrokeLength string = "Powertrain.CombustionEngine.StrokeLength"
	// FieldPowertrainCombustionEngineTPS is the name for the Powertrain.CombustionEngine.TPS column in the Clickhouse.
	FieldPowertrainCombustionEngineTPS string = "Powertrain.CombustionEngine.TPS"
	// FieldPowertrainCombustionEngineTorque is the name for the Powertrain.CombustionEngine.Torque column in the Clickhouse.
	FieldPowertrainCombustionEngineTorque string = "Powertrain.CombustionEngine.Torque"
	// FieldPowertrainElectricMotor is the name for the Powertrain.ElectricMotor column in the Clickhouse.
	FieldPowertrainElectricMotor string = "Powertrain.ElectricMotor"
	// FieldPowertrainElectricMotorCoolantTemperature is the name for the Powertrain.ElectricMotor.CoolantTemperature column in the Clickhouse.
	FieldPowertrainElectricMotorCoolantTemperature string = "Powertrain.ElectricMotor.CoolantTemperature"
	// FieldPowertrainElectricMotorEngineCode is the name for the Powertrain.ElectricMotor.EngineCode column in the Clickhouse.
	FieldPowertrainElectricMotorEngineCode string = "Powertrain.ElectricMotor.EngineCode"
	// FieldPowertrainElectricMotorMaxPower is the name for the Powertrain.ElectricMotor.MaxPower column in the Clickhouse.
	FieldPowertrainElectricMotorMaxPower string = "Powertrain.ElectricMotor.MaxPower"
	// FieldPowertrainElectricMotorMaxRegenPower is the name for the Powertrain.ElectricMotor.MaxRegenPower column in the Clickhouse.
	FieldPowertrainElectricMotorMaxRegenPower string = "Powertrain.ElectricMotor.MaxRegenPower"
	// FieldPowertrainElectricMotorMaxRegenTorque is the name for the Powertrain.ElectricMotor.MaxRegenTorque column in the Clickhouse.
	FieldPowertrainElectricMotorMaxRegenTorque string = "Powertrain.ElectricMotor.MaxRegenTorque"
	// FieldPowertrainElectricMotorMaxTorque is the name for the Powertrain.ElectricMotor.MaxTorque column in the Clickhouse.
	FieldPowertrainElectricMotorMaxTorque string = "Powertrain.ElectricMotor.MaxTorque"
	// FieldPowertrainElectricMotorPower is the name for the Powertrain.ElectricMotor.Power column in the Clickhouse.
	FieldPowertrainElectricMotorPower string = "Powertrain.ElectricMotor.Power"
	// FieldPowertrainElectricMotorSpeed is the name for the Powertrain.ElectricMotor.Speed column in the Clickhouse.
	FieldPowertrainElectricMotorSpeed string = "Powertrain.ElectricMotor.Speed"
	// FieldPowertrainElectricMotorTemperature is the name for the Powertrain.ElectricMotor.Temperature column in the Clickhouse.
	FieldPowertrainElectricMotorTemperature string = "Powertrain.ElectricMotor.Temperature"
	// FieldPowertrainElectricMotorTorque is the name for the Powertrain.ElectricMotor.Torque column in the Clickhouse.
	FieldPowertrainElectricMotorTorque string = "Powertrain.ElectricMotor.Torque"
	// FieldPowertrainFuelSystem is the name for the Powertrain.FuelSystem column in the Clickhouse.
	FieldPowertrainFuelSystem string = "Powertrain.FuelSystem"
	// FieldPowertrainFuelSystemAbsoluteLevel is the name for the Powertrain.FuelSystem.AbsoluteLevel column in the Clickhouse.
	FieldPowertrainFuelSystemAbsoluteLevel string = "Powertrain.FuelSystem.AbsoluteLevel"
	// FieldPowertrainFuelSystemAverageConsumption is the name for the Powertrain.FuelSystem.AverageConsumption column in the Clickhouse.
	FieldPowertrainFuelSystemAverageConsumption string = "Powertrain.FuelSystem.AverageConsumption"
	// FieldPowertrainFuelSystemConsumptionSinceStart is the name for the Powertrain.FuelSystem.ConsumptionSinceStart column in the Clickhouse.
	FieldPowertrainFuelSystemConsumptionSinceStart string = "Powertrain.FuelSystem.ConsumptionSinceStart"
	// FieldPowertrainFuelSystemHybridType is the name for the Powertrain.FuelSystem.HybridType column in the Clickhouse.
	FieldPowertrainFuelSystemHybridType string = "Powertrain.FuelSystem.HybridType"
	// FieldPowertrainFuelSystemInstantConsumption is the name for the Powertrain.FuelSystem.InstantConsumption column in the Clickhouse.
	FieldPowertrainFuelSystemInstantConsumption string = "Powertrain.FuelSystem.InstantConsumption"
	// FieldPowertrainFuelSystemIsEngineStopStartEnabled is the name for the Powertrain.FuelSystem.IsEngineStopStartEnabled column in the Clickhouse.
	FieldPowertrainFuelSystemIsEngineStopStartEnabled string = "Powertrain.FuelSystem.IsEngineStopStartEnabled"
	// FieldPowertrainFuelSystemIsFuelLevelLow is the name for the Powertrain.FuelSystem.IsFuelLevelLow column in the Clickhouse.
	FieldPowertrainFuelSystemIsFuelLevelLow string = "Powertrain.FuelSystem.IsFuelLevelLow"
	// FieldPowertrainFuelSystemIsFuelPortFlapOpen is the name for the Powertrain.FuelSystem.IsFuelPortFlapOpen column in the Clickhouse.
	FieldPowertrainFuelSystemIsFuelPortFlapOpen string = "Powertrain.FuelSystem.IsFuelPortFlapOpen"
	// FieldPowertrainFuelSystemRange is the name for the Powertrain.FuelSystem.Range column in the Clickhouse.
	FieldPowertrainFuelSystemRange string = "Powertrain.FuelSystem.Range"
	// FieldPowertrainFuelSystemRefuelPortPosition is the name for the Powertrain.FuelSystem.RefuelPortPosition column in the Clickhouse.
	FieldPowertrainFuelSystemRefuelPortPosition string = "Powertrain.FuelSystem.RefuelPortPosition"
	// FieldPowertrainFuelSystemRefuelPortPositionsize0 is the name for the Powertrain.FuelSystem.RefuelPortPosition.size0 column in the Clickhouse.
	FieldPowertrainFuelSystemRefuelPortPositionsize0 string = "Powertrain.FuelSystem.RefuelPortPosition.size0"
	// FieldPowertrainFuelSystemRelativeLevel is the name for the Powertrain.FuelSystem.RelativeLevel column in the Clickhouse.
	FieldPowertrainFuelSystemRelativeLevel string = "Powertrain.FuelSystem.RelativeLevel"
	// FieldPowertrainFuelSystemSupportedFuel is the name for the Powertrain.FuelSystem.SupportedFuel column in the Clickhouse.
	FieldPowertrainFuelSystemSupportedFuel string = "Powertrain.FuelSystem.SupportedFuel"
	// FieldPowertrainFuelSystemSupportedFuelTypes is the name for the Powertrain.FuelSystem.SupportedFuelTypes column in the Clickhouse.
	FieldPowertrainFuelSystemSupportedFuelTypes string = "Powertrain.FuelSystem.SupportedFuelTypes"
	// FieldPowertrainFuelSystemSupportedFuelTypessize0 is the name for the Powertrain.FuelSystem.SupportedFuelTypes.size0 column in the Clickhouse.
	FieldPowertrainFuelSystemSupportedFuelTypessize0 string = "Powertrain.FuelSystem.SupportedFuelTypes.size0"
	// FieldPowertrainFuelSystemSupportedFuelsize0 is the name for the Powertrain.FuelSystem.SupportedFuel.size0 column in the Clickhouse.
	FieldPowertrainFuelSystemSupportedFuelsize0 string = "Powertrain.FuelSystem.SupportedFuel.size0"
	// FieldPowertrainFuelSystemTankCapacity is the name for the Powertrain.FuelSystem.TankCapacity column in the Clickhouse.
	FieldPowertrainFuelSystemTankCapacity string = "Powertrain.FuelSystem.TankCapacity"
	// FieldPowertrainFuelSystemTimeRemaining is the name for the Powertrain.FuelSystem.TimeRemaining column in the Clickhouse.
	FieldPowertrainFuelSystemTimeRemaining string = "Powertrain.FuelSystem.TimeRemaining"
	// FieldPowertrainIsAutoPowerOptimize is the name for the Powertrain.IsAutoPowerOptimize column in the Clickhouse.
	FieldPowertrainIsAutoPowerOptimize string = "Powertrain.IsAutoPowerOptimize"
	// FieldPowertrainPowerOptimizeLevel is the name for the Powertrain.PowerOptimizeLevel column in the Clickhouse.
	FieldPowertrainPowerOptimizeLevel string = "Powertrain.PowerOptimizeLevel"
	// FieldPowertrainRange is the name for the Powertrain.Range column in the Clickhouse.
	FieldPowertrainRange string = "Powertrain.Range"
	// FieldPowertrainTimeRemaining is the name for the Powertrain.TimeRemaining column in the Clickhouse.
	FieldPowertrainTimeRemaining string = "Powertrain.TimeRemaining"
	// FieldPowertrainTractionBattery is the name for the Powertrain.TractionBattery column in the Clickhouse.
	FieldPowertrainTractionBattery string = "Powertrain.TractionBattery"
	// FieldPowertrainTractionBatteryAccumulatedChargedEnergy is the name for the Powertrain.TractionBattery.AccumulatedChargedEnergy column in the Clickhouse.
	FieldPowertrainTractionBatteryAccumulatedChargedEnergy string = "Powertrain.TractionBattery.AccumulatedChargedEnergy"
	// FieldPowertrainTractionBatteryAccumulatedChargedThroughput is the name for the Powertrain.TractionBattery.AccumulatedChargedThroughput column in the Clickhouse.
	FieldPowertrainTractionBatteryAccumulatedChargedThroughput string = "Powertrain.TractionBattery.AccumulatedChargedThroughput"
	// FieldPowertrainTractionBatteryAccumulatedConsumedEnergy is the name for the Powertrain.TractionBattery.AccumulatedConsumedEnergy column in the Clickhouse.
	FieldPowertrainTractionBatteryAccumulatedConsumedEnergy string = "Powertrain.TractionBattery.AccumulatedConsumedEnergy"
	// FieldPowertrainTractionBatteryAccumulatedConsumedThroughput is the name for the Powertrain.TractionBattery.AccumulatedConsumedThroughput column in the Clickhouse.
	FieldPowertrainTractionBatteryAccumulatedConsumedThroughput string = "Powertrain.TractionBattery.AccumulatedConsumedThroughput"
	// FieldPowertrainTractionBatteryCellVoltage is the name for the Powertrain.TractionBattery.CellVoltage column in the Clickhouse.
	FieldPowertrainTractionBatteryCellVoltage string = "Powertrain.TractionBattery.CellVoltage"
	// FieldPowertrainTractionBatteryCellVoltageCellVoltages is the name for the Powertrain.TractionBattery.CellVoltage.CellVoltages column in the Clickhouse.
	FieldPowertrainTractionBatteryCellVoltageCellVoltages string = "Powertrain.TractionBattery.CellVoltage.CellVoltages"
	// FieldPowertrainTractionBatteryCellVoltageCellVoltagessize0 is the name for the Powertrain.TractionBattery.CellVoltage.CellVoltages.size0 column in the Clickhouse.
	FieldPowertrainTractionBatteryCellVoltageCellVoltagessize0 string = "Powertrain.TractionBattery.CellVoltage.CellVoltages.size0"
	// FieldPowertrainTractionBatteryCellVoltageIdMax is the name for the Powertrain.TractionBattery.CellVoltage.IdMax column in the Clickhouse.
	FieldPowertrainTractionBatteryCellVoltageIdMax string = "Powertrain.TractionBattery.CellVoltage.IdMax"
	// FieldPowertrainTractionBatteryCellVoltageIdMin is the name for the Powertrain.TractionBattery.CellVoltage.IdMin column in the Clickhouse.
	FieldPowertrainTractionBatteryCellVoltageIdMin string = "Powertrain.TractionBattery.CellVoltage.IdMin"
	// FieldPowertrainTractionBatteryCellVoltageMax is the name for the Powertrain.TractionBattery.CellVoltage.Max column in the Clickhouse.
	FieldPowertrainTractionBatteryCellVoltageMax string = "Powertrain.TractionBattery.CellVoltage.Max"
	// FieldPowertrainTractionBatteryCellVoltageMin is the name for the Powertrain.TractionBattery.CellVoltage.Min column in the Clickhouse.
	FieldPowertrainTractionBatteryCellVoltageMin string = "Powertrain.TractionBattery.CellVoltage.Min"
	// FieldPowertrainTractionBatteryCharging is the name for the Powertrain.TractionBattery.Charging column in the Clickhouse.
	FieldPowertrainTractionBatteryCharging string = "Powertrain.TractionBattery.Charging"
	// FieldPowertrainTractionBatteryChargingAveragePower is the name for the Powertrain.TractionBattery.Charging.AveragePower column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingAveragePower string = "Powertrain.TractionBattery.Charging.AveragePower"
	// FieldPowertrainTractionBatteryChargingChargeCurrent is the name for the Powertrain.TractionBattery.Charging.ChargeCurrent column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargeCurrent string = "Powertrain.TractionBattery.Charging.ChargeCurrent"
	// FieldPowertrainTractionBatteryChargingChargeCurrentDC is the name for the Powertrain.TractionBattery.Charging.ChargeCurrent.DC column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargeCurrentDC string = "Powertrain.TractionBattery.Charging.ChargeCurrent.DC"
	// FieldPowertrainTractionBatteryChargingChargeCurrentPhase1 is the name for the Powertrain.TractionBattery.Charging.ChargeCurrent.Phase1 column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargeCurrentPhase1 string = "Powertrain.TractionBattery.Charging.ChargeCurrent.Phase1"
	// FieldPowertrainTractionBatteryChargingChargeCurrentPhase2 is the name for the Powertrain.TractionBattery.Charging.ChargeCurrent.Phase2 column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargeCurrentPhase2 string = "Powertrain.TractionBattery.Charging.ChargeCurrent.Phase2"
	// FieldPowertrainTractionBatteryChargingChargeCurrentPhase3 is the name for the Powertrain.TractionBattery.Charging.ChargeCurrent.Phase3 column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargeCurrentPhase3 string = "Powertrain.TractionBattery.Charging.ChargeCurrent.Phase3"
	// FieldPowertrainTractionBatteryChargingChargeLimit is the name for the Powertrain.TractionBattery.Charging.ChargeLimit column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargeLimit string = "Powertrain.TractionBattery.Charging.ChargeLimit"
	// FieldPowertrainTractionBatteryChargingChargePlugType is the name for the Powertrain.TractionBattery.Charging.ChargePlugType column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargePlugType string = "Powertrain.TractionBattery.Charging.ChargePlugType"
	// FieldPowertrainTractionBatteryChargingChargePlugTypesize0 is the name for the Powertrain.TractionBattery.Charging.ChargePlugType.size0 column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargePlugTypesize0 string = "Powertrain.TractionBattery.Charging.ChargePlugType.size0"
	// FieldPowertrainTractionBatteryChargingChargePortFlap is the name for the Powertrain.TractionBattery.Charging.ChargePortFlap column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargePortFlap string = "Powertrain.TractionBattery.Charging.ChargePortFlap"
	// FieldPowertrainTractionBatteryChargingChargePortPosition is the name for the Powertrain.TractionBattery.Charging.ChargePortPosition column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargePortPosition string = "Powertrain.TractionBattery.Charging.ChargePortPosition"
	// FieldPowertrainTractionBatteryChargingChargePortPositionsize0 is the name for the Powertrain.TractionBattery.Charging.ChargePortPosition.size0 column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargePortPositionsize0 string = "Powertrain.TractionBattery.Charging.ChargePortPosition.size0"
	// FieldPowertrainTractionBatteryChargingChargePortType is the name for the Powertrain.TractionBattery.Charging.ChargePortType column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargePortType string = "Powertrain.TractionBattery.Charging.ChargePortType"
	// FieldPowertrainTractionBatteryChargingChargePortTypesize0 is the name for the Powertrain.TractionBattery.Charging.ChargePortType.size0 column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargePortTypesize0 string = "Powertrain.TractionBattery.Charging.ChargePortType.size0"
	// FieldPowertrainTractionBatteryChargingChargeRate is the name for the Powertrain.TractionBattery.Charging.ChargeRate column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargeRate string = "Powertrain.TractionBattery.Charging.ChargeRate"
	// FieldPowertrainTractionBatteryChargingChargeVoltage is the name for the Powertrain.TractionBattery.Charging.ChargeVoltage column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargeVoltage string = "Powertrain.TractionBattery.Charging.ChargeVoltage"
	// FieldPowertrainTractionBatteryChargingChargeVoltageDC is the name for the Powertrain.TractionBattery.Charging.ChargeVoltage.DC column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargeVoltageDC string = "Powertrain.TractionBattery.Charging.ChargeVoltage.DC"
	// FieldPowertrainTractionBatteryChargingChargeVoltagePhase1 is the name for the Powertrain.TractionBattery.Charging.ChargeVoltage.Phase1 column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargeVoltagePhase1 string = "Powertrain.TractionBattery.Charging.ChargeVoltage.Phase1"
	// FieldPowertrainTractionBatteryChargingChargeVoltagePhase2 is the name for the Powertrain.TractionBattery.Charging.ChargeVoltage.Phase2 column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargeVoltagePhase2 string = "Powertrain.TractionBattery.Charging.ChargeVoltage.Phase2"
	// FieldPowertrainTractionBatteryChargingChargeVoltagePhase3 is the name for the Powertrain.TractionBattery.Charging.ChargeVoltage.Phase3 column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingChargeVoltagePhase3 string = "Powertrain.TractionBattery.Charging.ChargeVoltage.Phase3"
	// FieldPowertrainTractionBatteryChargingEvseId is the name for the Powertrain.TractionBattery.Charging.EvseId column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingEvseId string = "Powertrain.TractionBattery.Charging.EvseId"
	// FieldPowertrainTractionBatteryChargingIsChargePortFlapOpen is the name for the Powertrain.TractionBattery.Charging.IsChargePortFlapOpen column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingIsChargePortFlapOpen string = "Powertrain.TractionBattery.Charging.IsChargePortFlapOpen"
	// FieldPowertrainTractionBatteryChargingIsCharging is the name for the Powertrain.TractionBattery.Charging.IsCharging column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingIsCharging string = "Powertrain.TractionBattery.Charging.IsCharging"
	// FieldPowertrainTractionBatteryChargingIsChargingCableConnected is the name for the Powertrain.TractionBattery.Charging.IsChargingCableConnected column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingIsChargingCableConnected string = "Powertrain.TractionBattery.Charging.IsChargingCableConnected"
	// FieldPowertrainTractionBatteryChargingIsChargingCableLocked is the name for the Powertrain.TractionBattery.Charging.IsChargingCableLocked column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingIsChargingCableLocked string = "Powertrain.TractionBattery.Charging.IsChargingCableLocked"
	// FieldPowertrainTractionBatteryChargingIsDischarging is the name for the Powertrain.TractionBattery.Charging.IsDischarging column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingIsDischarging string = "Powertrain.TractionBattery.Charging.IsDischarging"
	// FieldPowertrainTractionBatteryChargingLocation is the name for the Powertrain.TractionBattery.Charging.Location column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingLocation string = "Powertrain.TractionBattery.Charging.Location"
	// FieldPowertrainTractionBatteryChargingLocationAltitude is the name for the Powertrain.TractionBattery.Charging.Location.Altitude column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingLocationAltitude string = "Powertrain.TractionBattery.Charging.Location.Altitude"
	// FieldPowertrainTractionBatteryChargingLocationLatitude is the name for the Powertrain.TractionBattery.Charging.Location.Latitude column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingLocationLatitude string = "Powertrain.TractionBattery.Charging.Location.Latitude"
	// FieldPowertrainTractionBatteryChargingLocationLongitude is the name for the Powertrain.TractionBattery.Charging.Location.Longitude column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingLocationLongitude string = "Powertrain.TractionBattery.Charging.Location.Longitude"
	// FieldPowertrainTractionBatteryChargingMaxPower is the name for the Powertrain.TractionBattery.Charging.MaxPower column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingMaxPower string = "Powertrain.TractionBattery.Charging.MaxPower"
	// FieldPowertrainTractionBatteryChargingMaximumChargingCurrent is the name for the Powertrain.TractionBattery.Charging.MaximumChargingCurrent column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingMaximumChargingCurrent string = "Powertrain.TractionBattery.Charging.MaximumChargingCurrent"
	// FieldPowertrainTractionBatteryChargingMaximumChargingCurrentDC is the name for the Powertrain.TractionBattery.Charging.MaximumChargingCurrent.DC column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingMaximumChargingCurrentDC string = "Powertrain.TractionBattery.Charging.MaximumChargingCurrent.DC"
	// FieldPowertrainTractionBatteryChargingMaximumChargingCurrentPhase1 is the name for the Powertrain.TractionBattery.Charging.MaximumChargingCurrent.Phase1 column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingMaximumChargingCurrentPhase1 string = "Powertrain.TractionBattery.Charging.MaximumChargingCurrent.Phase1"
	// FieldPowertrainTractionBatteryChargingMaximumChargingCurrentPhase2 is the name for the Powertrain.TractionBattery.Charging.MaximumChargingCurrent.Phase2 column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingMaximumChargingCurrentPhase2 string = "Powertrain.TractionBattery.Charging.MaximumChargingCurrent.Phase2"
	// FieldPowertrainTractionBatteryChargingMaximumChargingCurrentPhase3 is the name for the Powertrain.TractionBattery.Charging.MaximumChargingCurrent.Phase3 column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingMaximumChargingCurrentPhase3 string = "Powertrain.TractionBattery.Charging.MaximumChargingCurrent.Phase3"
	// FieldPowertrainTractionBatteryChargingMode is the name for the Powertrain.TractionBattery.Charging.Mode column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingMode string = "Powertrain.TractionBattery.Charging.Mode"
	// FieldPowertrainTractionBatteryChargingPowerLoss is the name for the Powertrain.TractionBattery.Charging.PowerLoss column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingPowerLoss string = "Powertrain.TractionBattery.Charging.PowerLoss"
	// FieldPowertrainTractionBatteryChargingStartStopCharging is the name for the Powertrain.TractionBattery.Charging.StartStopCharging column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingStartStopCharging string = "Powertrain.TractionBattery.Charging.StartStopCharging"
	// FieldPowertrainTractionBatteryChargingTemperature is the name for the Powertrain.TractionBattery.Charging.Temperature column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingTemperature string = "Powertrain.TractionBattery.Charging.Temperature"
	// FieldPowertrainTractionBatteryChargingTimeToComplete is the name for the Powertrain.TractionBattery.Charging.TimeToComplete column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingTimeToComplete string = "Powertrain.TractionBattery.Charging.TimeToComplete"
	// FieldPowertrainTractionBatteryChargingTimer is the name for the Powertrain.TractionBattery.Charging.Timer column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingTimer string = "Powertrain.TractionBattery.Charging.Timer"
	// FieldPowertrainTractionBatteryChargingTimerMode is the name for the Powertrain.TractionBattery.Charging.Timer.Mode column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingTimerMode string = "Powertrain.TractionBattery.Charging.Timer.Mode"
	// FieldPowertrainTractionBatteryChargingTimerTime is the name for the Powertrain.TractionBattery.Charging.Timer.Time column in the Clickhouse.
	FieldPowertrainTractionBatteryChargingTimerTime string = "Powertrain.TractionBattery.Charging.Timer.Time"
	// FieldPowertrainTractionBatteryCurrentCurrent is the name for the Powertrain.TractionBattery.CurrentCurrent column in the Clickhouse.
	FieldPowertrainTractionBatteryCurrentCurrent string = "Powertrain.TractionBattery.CurrentCurrent"
	// FieldPowertrainTractionBatteryCurrentPower is the name for the Powertrain.TractionBattery.CurrentPower column in the Clickhouse.
	FieldPowertrainTractionBatteryCurrentPower string = "Powertrain.TractionBattery.CurrentPower"
	// FieldPowertrainTractionBatteryCurrentVoltage is the name for the Powertrain.TractionBattery.CurrentVoltage column in the Clickhouse.
	FieldPowertrainTractionBatteryCurrentVoltage string = "Powertrain.TractionBattery.CurrentVoltage"
	// FieldPowertrainTractionBatteryDCDC is the name for the Powertrain.TractionBattery.DCDC column in the Clickhouse.
	FieldPowertrainTractionBatteryDCDC string = "Powertrain.TractionBattery.DCDC"
	// FieldPowertrainTractionBatteryDCDCPowerLoss is the name for the Powertrain.TractionBattery.DCDC.PowerLoss column in the Clickhouse.
	FieldPowertrainTractionBatteryDCDCPowerLoss string = "Powertrain.TractionBattery.DCDC.PowerLoss"
	// FieldPowertrainTractionBatteryDCDCTemperature is the name for the Powertrain.TractionBattery.DCDC.Temperature column in the Clickhouse.
	FieldPowertrainTractionBatteryDCDCTemperature string = "Powertrain.TractionBattery.DCDC.Temperature"
	// FieldPowertrainTractionBatteryErrorCodes is the name for the Powertrain.TractionBattery.ErrorCodes column in the Clickhouse.
	FieldPowertrainTractionBatteryErrorCodes string = "Powertrain.TractionBattery.ErrorCodes"
	// FieldPowertrainTractionBatteryErrorCodessize0 is the name for the Powertrain.TractionBattery.ErrorCodes.size0 column in the Clickhouse.
	FieldPowertrainTractionBatteryErrorCodessize0 string = "Powertrain.TractionBattery.ErrorCodes.size0"
	// FieldPowertrainTractionBatteryGrossCapacity is the name for the Powertrain.TractionBattery.GrossCapacity column in the Clickhouse.
	FieldPowertrainTractionBatteryGrossCapacity string = "Powertrain.TractionBattery.GrossCapacity"
	// FieldPowertrainTractionBatteryId is the name for the Powertrain.TractionBattery.Id column in the Clickhouse.
	FieldPowertrainTractionBatteryId string = "Powertrain.TractionBattery.Id"
	// FieldPowertrainTractionBatteryIsGroundConnected is the name for the Powertrain.TractionBattery.IsGroundConnected column in the Clickhouse.
	FieldPowertrainTractionBatteryIsGroundConnected string = "Powertrain.TractionBattery.IsGroundConnected"
	// FieldPowertrainTractionBatteryIsPowerConnected is the name for the Powertrain.TractionBattery.IsPowerConnected column in the Clickhouse.
	FieldPowertrainTractionBatteryIsPowerConnected string = "Powertrain.TractionBattery.IsPowerConnected"
	// FieldPowertrainTractionBatteryMaxVoltage is the name for the Powertrain.TractionBattery.MaxVoltage column in the Clickhouse.
	FieldPowertrainTractionBatteryMaxVoltage string = "Powertrain.TractionBattery.MaxVoltage"
	// FieldPowertrainTractionBatteryNetCapacity is the name for the Powertrain.TractionBattery.NetCapacity column in the Clickhouse.
	FieldPowertrainTractionBatteryNetCapacity string = "Powertrain.TractionBattery.NetCapacity"
	// FieldPowertrainTractionBatteryNominalVoltage is the name for the Powertrain.TractionBattery.NominalVoltage column in the Clickhouse.
	FieldPowertrainTractionBatteryNominalVoltage string = "Powertrain.TractionBattery.NominalVoltage"
	// FieldPowertrainTractionBatteryPowerLoss is the name for the Powertrain.TractionBattery.PowerLoss column in the Clickhouse.
	FieldPowertrainTractionBatteryPowerLoss string = "Powertrain.TractionBattery.PowerLoss"
	// FieldPowertrainTractionBatteryProductionDate is the name for the Powertrain.TractionBattery.ProductionDate column in the Clickhouse.
	FieldPowertrainTractionBatteryProductionDate string = "Powertrain.TractionBattery.ProductionDate"
	// FieldPowertrainTractionBatteryRange is the name for the Powertrain.TractionBattery.Range column in the Clickhouse.
	FieldPowertrainTractionBatteryRange string = "Powertrain.TractionBattery.Range"
	// FieldPowertrainTractionBatteryStateOfCharge is the name for the Powertrain.TractionBattery.StateOfCharge column in the Clickhouse.
	FieldPowertrainTractionBatteryStateOfCharge string = "Powertrain.TractionBattery.StateOfCharge"
	// FieldPowertrainTractionBatteryStateOfChargeCurrent is the name for the Powertrain.TractionBattery.StateOfCharge.Current column in the Clickhouse.
	FieldPowertrainTractionBatteryStateOfChargeCurrent string = "Powertrain.TractionBattery.StateOfCharge.Current"
	// FieldPowertrainTractionBatteryStateOfChargeCurrentEnergy is the name for the Powertrain.TractionBattery.StateOfCharge.CurrentEnergy column in the Clickhouse.
	FieldPowertrainTractionBatteryStateOfChargeCurrentEnergy string = "Powertrain.TractionBattery.StateOfCharge.CurrentEnergy"
	// FieldPowertrainTractionBatteryStateOfChargeDisplayed is the name for the Powertrain.TractionBattery.StateOfCharge.Displayed column in the Clickhouse.
	FieldPowertrainTractionBatteryStateOfChargeDisplayed string = "Powertrain.TractionBattery.StateOfCharge.Displayed"
	// FieldPowertrainTractionBatteryStateOfHealth is the name for the Powertrain.TractionBattery.StateOfHealth column in the Clickhouse.
	FieldPowertrainTractionBatteryStateOfHealth string = "Powertrain.TractionBattery.StateOfHealth"
	// FieldPowertrainTractionBatteryTemperature is the name for the Powertrain.TractionBattery.Temperature column in the Clickhouse.
	FieldPowertrainTractionBatteryTemperature string = "Powertrain.TractionBattery.Temperature"
	// FieldPowertrainTractionBatteryTemperatureAverage is the name for the Powertrain.TractionBattery.Temperature.Average column in the Clickhouse.
	FieldPowertrainTractionBatteryTemperatureAverage string = "Powertrain.TractionBattery.Temperature.Average"
	// FieldPowertrainTractionBatteryTemperatureCellTemperature is the name for the Powertrain.TractionBattery.Temperature.CellTemperature column in the Clickhouse.
	FieldPowertrainTractionBatteryTemperatureCellTemperature string = "Powertrain.TractionBattery.Temperature.CellTemperature"
	// FieldPowertrainTractionBatteryTemperatureCellTemperaturesize0 is the name for the Powertrain.TractionBattery.Temperature.CellTemperature.size0 column in the Clickhouse.
	FieldPowertrainTractionBatteryTemperatureCellTemperaturesize0 string = "Powertrain.TractionBattery.Temperature.CellTemperature.size0"
	// FieldPowertrainTractionBatteryTemperatureMax is the name for the Powertrain.TractionBattery.Temperature.Max column in the Clickhouse.
	FieldPowertrainTractionBatteryTemperatureMax string = "Powertrain.TractionBattery.Temperature.Max"
	// FieldPowertrainTractionBatteryTemperatureMin is the name for the Powertrain.TractionBattery.Temperature.Min column in the Clickhouse.
	FieldPowertrainTractionBatteryTemperatureMin string = "Powertrain.TractionBattery.Temperature.Min"
	// FieldPowertrainTractionBatteryTimeRemaining is the name for the Powertrain.TractionBattery.TimeRemaining column in the Clickhouse.
	FieldPowertrainTractionBatteryTimeRemaining string = "Powertrain.TractionBattery.TimeRemaining"
	// FieldPowertrainTransmission is the name for the Powertrain.Transmission column in the Clickhouse.
	FieldPowertrainTransmission string = "Powertrain.Transmission"
	// FieldPowertrainTransmissionClutchEngagement is the name for the Powertrain.Transmission.ClutchEngagement column in the Clickhouse.
	FieldPowertrainTransmissionClutchEngagement string = "Powertrain.Transmission.ClutchEngagement"
	// FieldPowertrainTransmissionClutchWear is the name for the Powertrain.Transmission.ClutchWear column in the Clickhouse.
	FieldPowertrainTransmissionClutchWear string = "Powertrain.Transmission.ClutchWear"
	// FieldPowertrainTransmissionCurrentGear is the name for the Powertrain.Transmission.CurrentGear column in the Clickhouse.
	FieldPowertrainTransmissionCurrentGear string = "Powertrain.Transmission.CurrentGear"
	// FieldPowertrainTransmissionDiffLockFrontEngagement is the name for the Powertrain.Transmission.DiffLockFrontEngagement column in the Clickhouse.
	FieldPowertrainTransmissionDiffLockFrontEngagement string = "Powertrain.Transmission.DiffLockFrontEngagement"
	// FieldPowertrainTransmissionDiffLockRearEngagement is the name for the Powertrain.Transmission.DiffLockRearEngagement column in the Clickhouse.
	FieldPowertrainTransmissionDiffLockRearEngagement string = "Powertrain.Transmission.DiffLockRearEngagement"
	// FieldPowertrainTransmissionDriveType is the name for the Powertrain.Transmission.DriveType column in the Clickhouse.
	FieldPowertrainTransmissionDriveType string = "Powertrain.Transmission.DriveType"
	// FieldPowertrainTransmissionGearChangeMode is the name for the Powertrain.Transmission.GearChangeMode column in the Clickhouse.
	FieldPowertrainTransmissionGearChangeMode string = "Powertrain.Transmission.GearChangeMode"
	// FieldPowertrainTransmissionGearCount is the name for the Powertrain.Transmission.GearCount column in the Clickhouse.
	FieldPowertrainTransmissionGearCount string = "Powertrain.Transmission.GearCount"
	// FieldPowertrainTransmissionIsElectricalPowertrainEngaged is the name for the Powertrain.Transmission.IsElectricalPowertrainEngaged column in the Clickhouse.
	FieldPowertrainTransmissionIsElectricalPowertrainEngaged string = "Powertrain.Transmission.IsElectricalPowertrainEngaged"
	// FieldPowertrainTransmissionIsLowRangeEngaged is the name for the Powertrain.Transmission.IsLowRangeEngaged column in the Clickhouse.
	FieldPowertrainTransmissionIsLowRangeEngaged string = "Powertrain.Transmission.IsLowRangeEngaged"
	// FieldPowertrainTransmissionIsParkLockEngaged is the name for the Powertrain.Transmission.IsParkLockEngaged column in the Clickhouse.
	FieldPowertrainTransmissionIsParkLockEngaged string = "Powertrain.Transmission.IsParkLockEngaged"
	// FieldPowertrainTransmissionPerformanceMode is the name for the Powertrain.Transmission.PerformanceMode column in the Clickhouse.
	FieldPowertrainTransmissionPerformanceMode string = "Powertrain.Transmission.PerformanceMode"
	// FieldPowertrainTransmissionSelectedGear is the name for the Powertrain.Transmission.SelectedGear column in the Clickhouse.
	FieldPowertrainTransmissionSelectedGear string = "Powertrain.Transmission.SelectedGear"
	// FieldPowertrainTransmissionTemperature is the name for the Powertrain.Transmission.Temperature column in the Clickhouse.
	FieldPowertrainTransmissionTemperature string = "Powertrain.Transmission.Temperature"
	// FieldPowertrainTransmissionTorqueDistribution is the name for the Powertrain.Transmission.TorqueDistribution column in the Clickhouse.
	FieldPowertrainTransmissionTorqueDistribution string = "Powertrain.Transmission.TorqueDistribution"
	// FieldPowertrainTransmissionTravelledDistance is the name for the Powertrain.Transmission.TravelledDistance column in the Clickhouse.
	FieldPowertrainTransmissionTravelledDistance string = "Powertrain.Transmission.TravelledDistance"
	// FieldPowertrainTransmissionType is the name for the Powertrain.Transmission.Type column in the Clickhouse.
	FieldPowertrainTransmissionType string = "Powertrain.Transmission.Type"
	// FieldPowertrainType is the name for the Powertrain.Type column in the Clickhouse.
	FieldPowertrainType string = "Powertrain.Type"
	// FieldRoofLoad is the name for the RoofLoad column in the Clickhouse.
	FieldRoofLoad string = "RoofLoad"
	// FieldService is the name for the Service column in the Clickhouse.
	FieldService string = "Service"
	// FieldServiceDistanceToService is the name for the Service.DistanceToService column in the Clickhouse.
	FieldServiceDistanceToService string = "Service.DistanceToService"
	// FieldServiceIsServiceDue is the name for the Service.IsServiceDue column in the Clickhouse.
	FieldServiceIsServiceDue string = "Service.IsServiceDue"
	// FieldServiceTimeToService is the name for the Service.TimeToService column in the Clickhouse.
	FieldServiceTimeToService string = "Service.TimeToService"
	// FieldSpeed is the name for the Speed column in the Clickhouse.
	FieldSpeed string = "Speed"
	// FieldStartTime is the name for the StartTime column in the Clickhouse.
	FieldStartTime string = "StartTime"
	// FieldSubject is the name for the Subject column in the Clickhouse.
	FieldSubject string = "Subject"
	// FieldTimestamp is the name for the Timestamp column in the Clickhouse.
	FieldTimestamp string = "Timestamp"
	// FieldTrailerIsConnected is the name for the Trailer_IsConnected column in the Clickhouse.
	FieldTrailerIsConnected string = "Trailer_IsConnected"
	// FieldTraveledDistance is the name for the TraveledDistance column in the Clickhouse.
	FieldTraveledDistance string = "TraveledDistance"
	// FieldTraveledDistanceSinceStart is the name for the TraveledDistanceSinceStart column in the Clickhouse.
	FieldTraveledDistanceSinceStart string = "TraveledDistanceSinceStart"
	// FieldTripDuration is the name for the TripDuration column in the Clickhouse.
	FieldTripDuration string = "TripDuration"
	// FieldTripMeterReading is the name for the TripMeterReading column in the Clickhouse.
	FieldTripMeterReading string = "TripMeterReading"
	// FieldTurningDiameter is the name for the TurningDiameter column in the Clickhouse.
	FieldTurningDiameter string = "TurningDiameter"
	// FieldVehicleIdentification is the name for the VehicleIdentification column in the Clickhouse.
	FieldVehicleIdentification string = "VehicleIdentification"
	// FieldVehicleIdentificationAcrissCode is the name for the VehicleIdentification.AcrissCode column in the Clickhouse.
	FieldVehicleIdentificationAcrissCode string = "VehicleIdentification.AcrissCode"
	// FieldVehicleIdentificationBodyType is the name for the VehicleIdentification.BodyType column in the Clickhouse.
	FieldVehicleIdentificationBodyType string = "VehicleIdentification.BodyType"
	// FieldVehicleIdentificationBrand is the name for the VehicleIdentification.Brand column in the Clickhouse.
	FieldVehicleIdentificationBrand string = "VehicleIdentification.Brand"
	// FieldVehicleIdentificationDateVehicleFirstRegistered is the name for the VehicleIdentification.DateVehicleFirstRegistered column in the Clickhouse.
	FieldVehicleIdentificationDateVehicleFirstRegistered string = "VehicleIdentification.DateVehicleFirstRegistered"
	// FieldVehicleIdentificationKnownVehicleDamages is the name for the VehicleIdentification.KnownVehicleDamages column in the Clickhouse.
	FieldVehicleIdentificationKnownVehicleDamages string = "VehicleIdentification.KnownVehicleDamages"
	// FieldVehicleIdentificationLicensePlate is the name for the VehicleIdentification.LicensePlate column in the Clickhouse.
	FieldVehicleIdentificationLicensePlate string = "VehicleIdentification.LicensePlate"
	// FieldVehicleIdentificationMeetsEmissionStandard is the name for the VehicleIdentification.MeetsEmissionStandard column in the Clickhouse.
	FieldVehicleIdentificationMeetsEmissionStandard string = "VehicleIdentification.MeetsEmissionStandard"
	// FieldVehicleIdentificationModel is the name for the VehicleIdentification.Model column in the Clickhouse.
	FieldVehicleIdentificationModel string = "VehicleIdentification.Model"
	// FieldVehicleIdentificationOptionalExtras is the name for the VehicleIdentification.OptionalExtras column in the Clickhouse.
	FieldVehicleIdentificationOptionalExtras string = "VehicleIdentification.OptionalExtras"
	// FieldVehicleIdentificationOptionalExtrassize0 is the name for the VehicleIdentification.OptionalExtras.size0 column in the Clickhouse.
	FieldVehicleIdentificationOptionalExtrassize0 string = "VehicleIdentification.OptionalExtras.size0"
	// FieldVehicleIdentificationProductionDate is the name for the VehicleIdentification.ProductionDate column in the Clickhouse.
	FieldVehicleIdentificationProductionDate string = "VehicleIdentification.ProductionDate"
	// FieldVehicleIdentificationPurchaseDate is the name for the VehicleIdentification.PurchaseDate column in the Clickhouse.
	FieldVehicleIdentificationPurchaseDate string = "VehicleIdentification.PurchaseDate"
	// FieldVehicleIdentificationVIN is the name for the VehicleIdentification.VIN column in the Clickhouse.
	FieldVehicleIdentificationVIN string = "VehicleIdentification.VIN"
	// FieldVehicleIdentificationVehicleConfiguration is the name for the VehicleIdentification.VehicleConfiguration column in the Clickhouse.
	FieldVehicleIdentificationVehicleConfiguration string = "VehicleIdentification.VehicleConfiguration"
	// FieldVehicleIdentificationVehicleExteriorColor is the name for the VehicleIdentification.VehicleExteriorColor column in the Clickhouse.
	FieldVehicleIdentificationVehicleExteriorColor string = "VehicleIdentification.VehicleExteriorColor"
	// FieldVehicleIdentificationVehicleInteriorColor is the name for the VehicleIdentification.VehicleInteriorColor column in the Clickhouse.
	FieldVehicleIdentificationVehicleInteriorColor string = "VehicleIdentification.VehicleInteriorColor"
	// FieldVehicleIdentificationVehicleInteriorType is the name for the VehicleIdentification.VehicleInteriorType column in the Clickhouse.
	FieldVehicleIdentificationVehicleInteriorType string = "VehicleIdentification.VehicleInteriorType"
	// FieldVehicleIdentificationVehicleModelDate is the name for the VehicleIdentification.VehicleModelDate column in the Clickhouse.
	FieldVehicleIdentificationVehicleModelDate string = "VehicleIdentification.VehicleModelDate"
	// FieldVehicleIdentificationVehicleSeatingCapacity is the name for the VehicleIdentification.VehicleSeatingCapacity column in the Clickhouse.
	FieldVehicleIdentificationVehicleSeatingCapacity string = "VehicleIdentification.VehicleSeatingCapacity"
	// FieldVehicleIdentificationVehicleSpecialUsage is the name for the VehicleIdentification.VehicleSpecialUsage column in the Clickhouse.
	FieldVehicleIdentificationVehicleSpecialUsage string = "VehicleIdentification.VehicleSpecialUsage"
	// FieldVehicleIdentificationWMI is the name for the VehicleIdentification.WMI column in the Clickhouse.
	FieldVehicleIdentificationWMI string = "VehicleIdentification.WMI"
	// FieldVehicleIdentificationYear is the name for the VehicleIdentification.Year column in the Clickhouse.
	FieldVehicleIdentificationYear string = "VehicleIdentification.Year"
	// FieldVersionVSS is the name for the VersionVSS column in the Clickhouse.
	FieldVersionVSS string = "VersionVSS"
	// FieldVersionVSSLabel is the name for the VersionVSS.Label column in the Clickhouse.
	FieldVersionVSSLabel string = "VersionVSS.Label"
	// FieldVersionVSSMajor is the name for the VersionVSS.Major column in the Clickhouse.
	FieldVersionVSSMajor string = "VersionVSS.Major"
	// FieldVersionVSSMinor is the name for the VersionVSS.Minor column in the Clickhouse.
	FieldVersionVSSMinor string = "VersionVSS.Minor"
	// FieldVersionVSSPatch is the name for the VersionVSS.Patch column in the Clickhouse.
	FieldVersionVSSPatch string = "VersionVSS.Patch"
	// FieldWidth is the name for the Width column in the Clickhouse.
	FieldWidth string = "Width"
	// FieldWidthExcludingMirrors is the name for the WidthExcludingMirrors column in the Clickhouse.
	FieldWidthExcludingMirrors string = "WidthExcludingMirrors"
	// FieldWidthFoldedMirrors is the name for the WidthFoldedMirrors column in the Clickhouse.
	FieldWidthFoldedMirrors string = "WidthFoldedMirrors"
	// FieldWidthIncludingMirrors is the name for the WidthIncludingMirrors column in the Clickhouse.
	FieldWidthIncludingMirrors string = "WidthIncludingMirrors"
)
var (
	newColsFunc = map[string]func() proto.Column{
		FieldADASABSIsEnabled: MustGetColumn("UInt8"),
		FieldADASABSIsEngaged: MustGetColumn("UInt8"),
		FieldADASABSIsError: MustGetColumn("UInt8"),
		FieldADASActiveAutonomyLevel: MustGetColumn("String"),
		FieldADASCruiseControlIsActive: MustGetColumn("UInt8"),
		FieldADASCruiseControlIsEnabled: MustGetColumn("UInt8"),
		FieldADASCruiseControlIsError: MustGetColumn("UInt8"),
		FieldADASCruiseControlSpeedSet: MustGetColumn("Float32"),
		FieldADASDMSIsEnabled: MustGetColumn("UInt8"),
		FieldADASDMSIsError: MustGetColumn("UInt8"),
		FieldADASDMSIsWarning: MustGetColumn("UInt8"),
		FieldADASEBAIsEnabled: MustGetColumn("UInt8"),
		FieldADASEBAIsEngaged: MustGetColumn("UInt8"),
		FieldADASEBAIsError: MustGetColumn("UInt8"),
		FieldADASEBDIsEnabled: MustGetColumn("UInt8"),
		FieldADASEBDIsEngaged: MustGetColumn("UInt8"),
		FieldADASEBDIsError: MustGetColumn("UInt8"),
		FieldADASESCIsEnabled: MustGetColumn("UInt8"),
		FieldADASESCIsEngaged: MustGetColumn("UInt8"),
		FieldADASESCIsError: MustGetColumn("UInt8"),
		FieldADASESCIsStrongCrossWindDetected: MustGetColumn("UInt8"),
		FieldADASESCRoadFrictionLowerBound: MustGetColumn("Float32"),
		FieldADASESCRoadFrictionMostProbable: MustGetColumn("Float32"),
		FieldADASESCRoadFrictionUpperBound: MustGetColumn("Float32"),
		FieldADASIsAutoPowerOptimize: MustGetColumn("UInt8"),
		FieldADASLaneDepartureDetectionIsEnabled: MustGetColumn("UInt8"),
		FieldADASLaneDepartureDetectionIsError: MustGetColumn("UInt8"),
		FieldADASLaneDepartureDetectionIsWarning: MustGetColumn("UInt8"),
		FieldADASObstacleDetectionIsEnabled: MustGetColumn("UInt8"),
		FieldADASObstacleDetectionIsError: MustGetColumn("UInt8"),
		FieldADASObstacleDetectionIsWarning: MustGetColumn("UInt8"),
		FieldADASPowerOptimizeLevel: MustGetColumn("UInt32"),
		FieldADASSupportedAutonomyLevel: MustGetColumn("String"),
		FieldADASTCSIsEnabled: MustGetColumn("UInt8"),
		FieldADASTCSIsEngaged: MustGetColumn("UInt8"),
		FieldADASTCSIsError: MustGetColumn("UInt8"),
		FieldAccelerationLateral: MustGetColumn("Float32"),
		FieldAccelerationLongitudinal: MustGetColumn("Float32"),
		FieldAccelerationVertical: MustGetColumn("Float32"),
		FieldAngularVelocityPitch: MustGetColumn("Float32"),
		FieldAngularVelocityRoll: MustGetColumn("Float32"),
		FieldAngularVelocityYaw: MustGetColumn("Float32"),
		FieldAverageSpeed: MustGetColumn("Float32"),
		FieldBodyBodyType: MustGetColumn("String"),
		FieldBodyHoodIsOpen: MustGetColumn("UInt8"),
		FieldBodyHoodPosition: MustGetColumn("UInt32"),
		FieldBodyHoodSwitch: MustGetColumn("String"),
		FieldBodyHornIsActive: MustGetColumn("UInt8"),
		FieldBodyIsAutoPowerOptimize: MustGetColumn("UInt8"),
		FieldBodyLightsBackupIsDefect: MustGetColumn("UInt8"),
		FieldBodyLightsBackupIsOn: MustGetColumn("UInt8"),
		FieldBodyLightsBeamHighIsDefect: MustGetColumn("UInt8"),
		FieldBodyLightsBeamHighIsOn: MustGetColumn("UInt8"),
		FieldBodyLightsBeamLowIsDefect: MustGetColumn("UInt8"),
		FieldBodyLightsBeamLowIsOn: MustGetColumn("UInt8"),
		FieldBodyLightsBrakeIsActive: MustGetColumn("String"),
		FieldBodyLightsBrakeIsDefect: MustGetColumn("UInt8"),
		FieldBodyLightsDirectionIndicatorLeftIsDefect: MustGetColumn("UInt8"),
		FieldBodyLightsDirectionIndicatorLeftIsSignaling: MustGetColumn("UInt8"),
		FieldBodyLightsDirectionIndicatorRightIsDefect: MustGetColumn("UInt8"),
		FieldBodyLightsDirectionIndicatorRightIsSignaling: MustGetColumn("UInt8"),
		FieldBodyLightsFogFrontIsDefect: MustGetColumn("UInt8"),
		FieldBodyLightsFogFrontIsOn: MustGetColumn("UInt8"),
		FieldBodyLightsFogRearIsDefect: MustGetColumn("UInt8"),
		FieldBodyLightsFogRearIsOn: MustGetColumn("UInt8"),
		FieldBodyLightsHazardIsDefect: MustGetColumn("UInt8"),
		FieldBodyLightsHazardIsSignaling: MustGetColumn("UInt8"),
		FieldBodyLightsIsHighBeamSwitchOn: MustGetColumn("UInt8"),
		FieldBodyLightsLicensePlateIsDefect: MustGetColumn("UInt8"),
		FieldBodyLightsLicensePlateIsOn: MustGetColumn("UInt8"),
		FieldBodyLightsLightSwitch: MustGetColumn("String"),
		FieldBodyLightsParkingIsDefect: MustGetColumn("UInt8"),
		FieldBodyLightsParkingIsOn: MustGetColumn("UInt8"),
		FieldBodyLightsRunningIsDefect: MustGetColumn("UInt8"),
		FieldBodyLightsRunningIsOn: MustGetColumn("UInt8"),
		FieldBodyMirrorsDriverSideIsFolded: MustGetColumn("UInt8"),
		FieldBodyMirrorsDriverSideIsHeatingOn: MustGetColumn("UInt8"),
		FieldBodyMirrorsDriverSideIsLocked: MustGetColumn("UInt8"),
		FieldBodyMirrorsDriverSidePan: MustGetColumn("Int32"),
		FieldBodyMirrorsDriverSideTilt: MustGetColumn("Int32"),
		FieldBodyMirrorsPassengerSideIsFolded: MustGetColumn("UInt8"),
		FieldBodyMirrorsPassengerSideIsHeatingOn: MustGetColumn("UInt8"),
		FieldBodyMirrorsPassengerSideIsLocked: MustGetColumn("UInt8"),
		FieldBodyMirrorsPassengerSidePan: MustGetColumn("Int32"),
		FieldBodyMirrorsPassengerSideTilt: MustGetColumn("Int32"),
		FieldBodyPowerOptimizeLevel: MustGetColumn("UInt32"),
		FieldBodyRaindetectionIntensity: MustGetColumn("UInt32"),
		FieldBodyRearMainSpoilerPosition: MustGetColumn("Float32"),
		FieldBodyRefuelPosition: MustGetColumn("String"),
		FieldBodyTrunkFrontIsLightOn: MustGetColumn("UInt8"),
		FieldBodyTrunkFrontIsLocked: MustGetColumn("UInt8"),
		FieldBodyTrunkFrontIsOpen: MustGetColumn("UInt8"),
		FieldBodyTrunkFrontPosition: MustGetColumn("UInt32"),
		FieldBodyTrunkFrontSwitch: MustGetColumn("String"),
		FieldBodyTrunkRearIsLightOn: MustGetColumn("UInt8"),
		FieldBodyTrunkRearIsLocked: MustGetColumn("UInt8"),
		FieldBodyTrunkRearIsOpen: MustGetColumn("UInt8"),
		FieldBodyTrunkRearPosition: MustGetColumn("UInt32"),
		FieldBodyTrunkRearSwitch: MustGetColumn("String"),
		FieldBodyWindshieldFrontIsHeatingOn: MustGetColumn("UInt8"),
		FieldBodyWindshieldFrontWasherFluidIsLevelLow: MustGetColumn("UInt8"),
		FieldBodyWindshieldFrontWasherFluidLevel: MustGetColumn("UInt32"),
		FieldBodyWindshieldFrontWipingIntensity: MustGetColumn("UInt32"),
		FieldBodyWindshieldFrontWipingIsWipersWorn: MustGetColumn("UInt8"),
		FieldBodyWindshieldFrontWipingMode: MustGetColumn("String"),
		FieldBodyWindshieldFrontWipingSystemActualPosition: MustGetColumn("Float32"),
		FieldBodyWindshieldFrontWipingSystemDriveCurrent: MustGetColumn("Float32"),
		FieldBodyWindshieldFrontWipingSystemFrequency: MustGetColumn("UInt32"),
		FieldBodyWindshieldFrontWipingSystemIsBlocked: MustGetColumn("UInt8"),
		FieldBodyWindshieldFrontWipingSystemIsEndingWipeCycle: MustGetColumn("UInt8"),
		FieldBodyWindshieldFrontWipingSystemIsOverheated: MustGetColumn("UInt8"),
		FieldBodyWindshieldFrontWipingSystemIsPositionReached: MustGetColumn("UInt8"),
		FieldBodyWindshieldFrontWipingSystemIsWiperError: MustGetColumn("UInt8"),
		FieldBodyWindshieldFrontWipingSystemIsWiping: MustGetColumn("UInt8"),
		FieldBodyWindshieldFrontWipingSystemMode: MustGetColumn("String"),
		FieldBodyWindshieldFrontWipingSystemTargetPosition: MustGetColumn("Float32"),
		FieldBodyWindshieldFrontWipingWiperWear: MustGetColumn("UInt32"),
		FieldBodyWindshieldRearIsHeatingOn: MustGetColumn("UInt8"),
		FieldBodyWindshieldRearWasherFluidIsLevelLow: MustGetColumn("UInt8"),
		FieldBodyWindshieldRearWasherFluidLevel: MustGetColumn("UInt32"),
		FieldBodyWindshieldRearWipingIntensity: MustGetColumn("UInt32"),
		FieldBodyWindshieldRearWipingIsWipersWorn: MustGetColumn("UInt8"),
		FieldBodyWindshieldRearWipingMode: MustGetColumn("String"),
		FieldBodyWindshieldRearWipingSystemActualPosition: MustGetColumn("Float32"),
		FieldBodyWindshieldRearWipingSystemDriveCurrent: MustGetColumn("Float32"),
		FieldBodyWindshieldRearWipingSystemFrequency: MustGetColumn("UInt32"),
		FieldBodyWindshieldRearWipingSystemIsBlocked: MustGetColumn("UInt8"),
		FieldBodyWindshieldRearWipingSystemIsEndingWipeCycle: MustGetColumn("UInt8"),
		FieldBodyWindshieldRearWipingSystemIsOverheated: MustGetColumn("UInt8"),
		FieldBodyWindshieldRearWipingSystemIsPositionReached: MustGetColumn("UInt8"),
		FieldBodyWindshieldRearWipingSystemIsWiperError: MustGetColumn("UInt8"),
		FieldBodyWindshieldRearWipingSystemIsWiping: MustGetColumn("UInt8"),
		FieldBodyWindshieldRearWipingSystemMode: MustGetColumn("String"),
		FieldBodyWindshieldRearWipingSystemTargetPosition: MustGetColumn("Float32"),
		FieldBodyWindshieldRearWipingWiperWear: MustGetColumn("UInt32"),
		FieldCabinConvertibleStatus: MustGetColumn("String"),
		FieldCabinDoorCount: MustGetColumn("UInt32"),
		FieldCabinDoorRow1DriverSideIsChildLockActive: MustGetColumn("UInt8"),
		FieldCabinDoorRow1DriverSideIsLocked: MustGetColumn("UInt8"),
		FieldCabinDoorRow1DriverSideIsOpen: MustGetColumn("UInt8"),
		FieldCabinDoorRow1DriverSidePosition: MustGetColumn("UInt32"),
		FieldCabinDoorRow1DriverSideShadeIsOpen: MustGetColumn("UInt8"),
		FieldCabinDoorRow1DriverSideShadePosition: MustGetColumn("UInt32"),
		FieldCabinDoorRow1DriverSideShadeSwitch: MustGetColumn("String"),
		FieldCabinDoorRow1DriverSideSwitch: MustGetColumn("String"),
		FieldCabinDoorRow1DriverSideWindowIsOpen: MustGetColumn("UInt8"),
		FieldCabinDoorRow1DriverSideWindowPosition: MustGetColumn("UInt32"),
		FieldCabinDoorRow1DriverSideWindowSwitch: MustGetColumn("String"),
		FieldCabinDoorRow1PassengerSideIsChildLockActive: MustGetColumn("UInt8"),
		FieldCabinDoorRow1PassengerSideIsLocked: MustGetColumn("UInt8"),
		FieldCabinDoorRow1PassengerSideIsOpen: MustGetColumn("UInt8"),
		FieldCabinDoorRow1PassengerSidePosition: MustGetColumn("UInt32"),
		FieldCabinDoorRow1PassengerSideShadeIsOpen: MustGetColumn("UInt8"),
		FieldCabinDoorRow1PassengerSideShadePosition: MustGetColumn("UInt32"),
		FieldCabinDoorRow1PassengerSideShadeSwitch: MustGetColumn("String"),
		FieldCabinDoorRow1PassengerSideSwitch: MustGetColumn("String"),
		FieldCabinDoorRow1PassengerSideWindowIsOpen: MustGetColumn("UInt8"),
		FieldCabinDoorRow1PassengerSideWindowPosition: MustGetColumn("UInt32"),
		FieldCabinDoorRow1PassengerSideWindowSwitch: MustGetColumn("String"),
		FieldCabinDoorRow2DriverSideIsChildLockActive: MustGetColumn("UInt8"),
		FieldCabinDoorRow2DriverSideIsLocked: MustGetColumn("UInt8"),
		FieldCabinDoorRow2DriverSideIsOpen: MustGetColumn("UInt8"),
		FieldCabinDoorRow2DriverSidePosition: MustGetColumn("UInt32"),
		FieldCabinDoorRow2DriverSideShadeIsOpen: MustGetColumn("UInt8"),
		FieldCabinDoorRow2DriverSideShadePosition: MustGetColumn("UInt32"),
		FieldCabinDoorRow2DriverSideShadeSwitch: MustGetColumn("String"),
		FieldCabinDoorRow2DriverSideSwitch: MustGetColumn("String"),
		FieldCabinDoorRow2DriverSideWindowIsOpen: MustGetColumn("UInt8"),
		FieldCabinDoorRow2DriverSideWindowPosition: MustGetColumn("UInt32"),
		FieldCabinDoorRow2DriverSideWindowSwitch: MustGetColumn("String"),
		FieldCabinDoorRow2PassengerSideIsChildLockActive: MustGetColumn("UInt8"),
		FieldCabinDoorRow2PassengerSideIsLocked: MustGetColumn("UInt8"),
		FieldCabinDoorRow2PassengerSideIsOpen: MustGetColumn("UInt8"),
		FieldCabinDoorRow2PassengerSidePosition: MustGetColumn("UInt32"),
		FieldCabinDoorRow2PassengerSideShadeIsOpen: MustGetColumn("UInt8"),
		FieldCabinDoorRow2PassengerSideShadePosition: MustGetColumn("UInt32"),
		FieldCabinDoorRow2PassengerSideShadeSwitch: MustGetColumn("String"),
		FieldCabinDoorRow2PassengerSideSwitch: MustGetColumn("String"),
		FieldCabinDoorRow2PassengerSideWindowIsOpen: MustGetColumn("UInt8"),
		FieldCabinDoorRow2PassengerSideWindowPosition: MustGetColumn("UInt32"),
		FieldCabinDoorRow2PassengerSideWindowSwitch: MustGetColumn("String"),
		FieldCabinDriverPosition: MustGetColumn("String"),
		FieldCabinHVACAmbientAirTemperature: MustGetColumn("Float32"),
		FieldCabinHVACIsAirConditioningActive: MustGetColumn("UInt8"),
		FieldCabinHVACIsAutoPowerOptimize: MustGetColumn("UInt8"),
		FieldCabinHVACIsFrontDefrosterActive: MustGetColumn("UInt8"),
		FieldCabinHVACIsRearDefrosterActive: MustGetColumn("UInt8"),
		FieldCabinHVACIsRecirculationActive: MustGetColumn("UInt8"),
		FieldCabinHVACPowerOptimizeLevel: MustGetColumn("UInt32"),
		FieldCabinHVACStationRow1DriverAirDistribution: MustGetColumn("String"),
		FieldCabinHVACStationRow1DriverFanSpeed: MustGetColumn("UInt32"),
		FieldCabinHVACStationRow1DriverTemperature: MustGetColumn("Float32"),
		FieldCabinHVACStationRow1PassengerAirDistribution: MustGetColumn("String"),
		FieldCabinHVACStationRow1PassengerFanSpeed: MustGetColumn("UInt32"),
		FieldCabinHVACStationRow1PassengerTemperature: MustGetColumn("Float32"),
		FieldCabinHVACStationRow2DriverAirDistribution: MustGetColumn("String"),
		FieldCabinHVACStationRow2DriverFanSpeed: MustGetColumn("UInt32"),
		FieldCabinHVACStationRow2DriverTemperature: MustGetColumn("Float32"),
		FieldCabinHVACStationRow2PassengerAirDistribution: MustGetColumn("String"),
		FieldCabinHVACStationRow2PassengerFanSpeed: MustGetColumn("UInt32"),
		FieldCabinHVACStationRow2PassengerTemperature: MustGetColumn("Float32"),
		FieldCabinHVACStationRow3DriverAirDistribution: MustGetColumn("String"),
		FieldCabinHVACStationRow3DriverFanSpeed: MustGetColumn("UInt32"),
		FieldCabinHVACStationRow3DriverTemperature: MustGetColumn("Float32"),
		FieldCabinHVACStationRow3PassengerAirDistribution: MustGetColumn("String"),
		FieldCabinHVACStationRow3PassengerFanSpeed: MustGetColumn("UInt32"),
		FieldCabinHVACStationRow3PassengerTemperature: MustGetColumn("Float32"),
		FieldCabinHVACStationRow4DriverAirDistribution: MustGetColumn("String"),
		FieldCabinHVACStationRow4DriverFanSpeed: MustGetColumn("UInt32"),
		FieldCabinHVACStationRow4DriverTemperature: MustGetColumn("Float32"),
		FieldCabinHVACStationRow4PassengerAirDistribution: MustGetColumn("String"),
		FieldCabinHVACStationRow4PassengerFanSpeed: MustGetColumn("UInt32"),
		FieldCabinHVACStationRow4PassengerTemperature: MustGetColumn("Float32"),
		FieldCabinInfotainmentHMIBrightness: MustGetColumn("Float32"),
		FieldCabinInfotainmentHMICurrentLanguage: MustGetColumn("String"),
		FieldCabinInfotainmentHMIDateFormat: MustGetColumn("String"),
		FieldCabinInfotainmentHMIDayNightMode: MustGetColumn("String"),
		FieldCabinInfotainmentHMIDisplayOffDuration: MustGetColumn("UInt32"),
		FieldCabinInfotainmentHMIDistanceUnit: MustGetColumn("String"),
		FieldCabinInfotainmentHMIEVEconomyUnits: MustGetColumn("String"),
		FieldCabinInfotainmentHMIEVEnergyUnits: MustGetColumn("String"),
		FieldCabinInfotainmentHMIFontSize: MustGetColumn("String"),
		FieldCabinInfotainmentHMIFuelEconomyUnits: MustGetColumn("String"),
		FieldCabinInfotainmentHMIFuelVolumeUnit: MustGetColumn("String"),
		FieldCabinInfotainmentHMIIsScreenAlwaysOn: MustGetColumn("UInt8"),
		FieldCabinInfotainmentHMILastActionTime: MustGetColumn("String"),
		FieldCabinInfotainmentHMISpeedUnit: MustGetColumn("String"),
		FieldCabinInfotainmentHMITemperatureUnit: MustGetColumn("String"),
		FieldCabinInfotainmentHMITimeFormat: MustGetColumn("String"),
		FieldCabinInfotainmentHMITirePressureUnit: MustGetColumn("String"),
		FieldCabinInfotainmentIsAutoPowerOptimize: MustGetColumn("UInt8"),
		FieldCabinInfotainmentMediaAction: MustGetColumn("String"),
		FieldCabinInfotainmentMediaDeclinedURI: MustGetColumn("String"),
		FieldCabinInfotainmentMediaPlayedAlbum: MustGetColumn("String"),
		FieldCabinInfotainmentMediaPlayedArtist: MustGetColumn("String"),
		FieldCabinInfotainmentMediaPlayedPlaybackRate: MustGetColumn("Float32"),
		FieldCabinInfotainmentMediaPlayedSource: MustGetColumn("String"),
		FieldCabinInfotainmentMediaPlayedTrack: MustGetColumn("String"),
		FieldCabinInfotainmentMediaPlayedURI: MustGetColumn("String"),
		FieldCabinInfotainmentMediaSelectedURI: MustGetColumn("String"),
		FieldCabinInfotainmentMediaVolume: MustGetColumn("UInt32"),
		FieldCabinInfotainmentNavigationDestinationSetLatitude: MustGetColumn("Float64"),
		FieldCabinInfotainmentNavigationDestinationSetLongitude: MustGetColumn("Float64"),
		FieldCabinInfotainmentNavigationGuidanceVoice: MustGetColumn("String"),
		FieldCabinInfotainmentNavigationMapIsAutoScaleModeUsed: MustGetColumn("UInt8"),
		FieldCabinInfotainmentNavigationMute: MustGetColumn("String"),
		FieldCabinInfotainmentNavigationVolume: MustGetColumn("UInt32"),
		FieldCabinInfotainmentPowerOptimizeLevel: MustGetColumn("UInt32"),
		FieldCabinInfotainmentSmartphoneProjectionActive: MustGetColumn("String"),
		FieldCabinInfotainmentSmartphoneProjectionSource: MustGetColumn("String"),
		FieldCabinInfotainmentSmartphoneProjectionSupportedMode: MustGetColumn("Array(String)"),
		FieldCabinInfotainmentSmartphoneProjectionSupportedModesize0: MustGetColumn("UInt64"),
		FieldCabinIsAutoPowerOptimize: MustGetColumn("UInt8"),
		FieldCabinIsWindowChildLockEngaged: MustGetColumn("UInt8"),
		FieldCabinLightAmbientLightRow1DriverSideColor: MustGetColumn("String"),
		FieldCabinLightAmbientLightRow1DriverSideIntensity: MustGetColumn("UInt32"),
		FieldCabinLightAmbientLightRow1DriverSideIsLightOn: MustGetColumn("UInt8"),
		FieldCabinLightAmbientLightRow1PassengerSideColor: MustGetColumn("String"),
		FieldCabinLightAmbientLightRow1PassengerSideIntensity: MustGetColumn("UInt32"),
		FieldCabinLightAmbientLightRow1PassengerSideIsLightOn: MustGetColumn("UInt8"),
		FieldCabinLightAmbientLightRow2DriverSideColor: MustGetColumn("String"),
		FieldCabinLightAmbientLightRow2DriverSideIntensity: MustGetColumn("UInt32"),
		FieldCabinLightAmbientLightRow2DriverSideIsLightOn: MustGetColumn("UInt8"),
		FieldCabinLightAmbientLightRow2PassengerSideColor: MustGetColumn("String"),
		FieldCabinLightAmbientLightRow2PassengerSideIntensity: MustGetColumn("UInt32"),
		FieldCabinLightAmbientLightRow2PassengerSideIsLightOn: MustGetColumn("UInt8"),
		FieldCabinLightInteractiveLightBarColor: MustGetColumn("String"),
		FieldCabinLightInteractiveLightBarEffect: MustGetColumn("String"),
		FieldCabinLightInteractiveLightBarIntensity: MustGetColumn("UInt32"),
		FieldCabinLightInteractiveLightBarIsLightOn: MustGetColumn("UInt8"),
		FieldCabinLightIsDomeOn: MustGetColumn("UInt8"),
		FieldCabinLightIsGloveBoxOn: MustGetColumn("UInt8"),
		FieldCabinLightPerceivedAmbientLight: MustGetColumn("UInt32"),
		FieldCabinLightSpotlightRow1DriverSideColor: MustGetColumn("String"),
		FieldCabinLightSpotlightRow1DriverSideIntensity: MustGetColumn("UInt32"),
		FieldCabinLightSpotlightRow1DriverSideIsLightOn: MustGetColumn("UInt8"),
		FieldCabinLightSpotlightRow1PassengerSideColor: MustGetColumn("String"),
		FieldCabinLightSpotlightRow1PassengerSideIntensity: MustGetColumn("UInt32"),
		FieldCabinLightSpotlightRow1PassengerSideIsLightOn: MustGetColumn("UInt8"),
		FieldCabinLightSpotlightRow2DriverSideColor: MustGetColumn("String"),
		FieldCabinLightSpotlightRow2DriverSideIntensity: MustGetColumn("UInt32"),
		FieldCabinLightSpotlightRow2DriverSideIsLightOn: MustGetColumn("UInt8"),
		FieldCabinLightSpotlightRow2PassengerSideColor: MustGetColumn("String"),
		FieldCabinLightSpotlightRow2PassengerSideIntensity: MustGetColumn("UInt32"),
		FieldCabinLightSpotlightRow2PassengerSideIsLightOn: MustGetColumn("UInt8"),
		FieldCabinLightSpotlightRow3DriverSideColor: MustGetColumn("String"),
		FieldCabinLightSpotlightRow3DriverSideIntensity: MustGetColumn("UInt32"),
		FieldCabinLightSpotlightRow3DriverSideIsLightOn: MustGetColumn("UInt8"),
		FieldCabinLightSpotlightRow3PassengerSideColor: MustGetColumn("String"),
		FieldCabinLightSpotlightRow3PassengerSideIntensity: MustGetColumn("UInt32"),
		FieldCabinLightSpotlightRow3PassengerSideIsLightOn: MustGetColumn("UInt8"),
		FieldCabinLightSpotlightRow4DriverSideColor: MustGetColumn("String"),
		FieldCabinLightSpotlightRow4DriverSideIntensity: MustGetColumn("UInt32"),
		FieldCabinLightSpotlightRow4DriverSideIsLightOn: MustGetColumn("UInt8"),
		FieldCabinLightSpotlightRow4PassengerSideColor: MustGetColumn("String"),
		FieldCabinLightSpotlightRow4PassengerSideIntensity: MustGetColumn("UInt32"),
		FieldCabinLightSpotlightRow4PassengerSideIsLightOn: MustGetColumn("UInt8"),
		FieldCabinPowerOptimizeLevel: MustGetColumn("UInt32"),
		FieldCabinRearShadeIsOpen: MustGetColumn("UInt8"),
		FieldCabinRearShadePosition: MustGetColumn("UInt32"),
		FieldCabinRearShadeSwitch: MustGetColumn("String"),
		FieldCabinRearviewMirrorDimmingLevel: MustGetColumn("UInt32"),
		FieldCabinSeatPosCount: MustGetColumn("Array(UInt32)"),
		FieldCabinSeatPosCountsize0: MustGetColumn("UInt64"),
		FieldCabinSeatRow1DriverSideAirbagIsDeployed: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideBackrestLumbarHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow1DriverSideBackrestLumbarSupport: MustGetColumn("Float32"),
		FieldCabinSeatRow1DriverSideBackrestRecline: MustGetColumn("Float32"),
		FieldCabinSeatRow1DriverSideBackrestSideBolsterSupport: MustGetColumn("Float32"),
		FieldCabinSeatRow1DriverSideHeadrestAngle: MustGetColumn("Float32"),
		FieldCabinSeatRow1DriverSideHeadrestHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow1DriverSideHeating: MustGetColumn("Int32"),
		FieldCabinSeatRow1DriverSideHeatingCooling: MustGetColumn("Int32"),
		FieldCabinSeatRow1DriverSideHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow1DriverSideIsBelted: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideIsOccupied: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideMassage: MustGetColumn("UInt32"),
		FieldCabinSeatRow1DriverSideOccupantIdentifierIssuer: MustGetColumn("String"),
		FieldCabinSeatRow1DriverSideOccupantIdentifierSubject: MustGetColumn("String"),
		FieldCabinSeatRow1DriverSidePosition: MustGetColumn("UInt32"),
		FieldCabinSeatRow1DriverSideSeatBeltHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow1DriverSideSeatingLength: MustGetColumn("UInt32"),
		FieldCabinSeatRow1DriverSideSwitchBackrestIsReclineBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchBackrestIsReclineForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchBackrestLumbarIsDownEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchBackrestLumbarIsLessSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchBackrestLumbarIsMoreSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchBackrestLumbarIsUpEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchBackrestSideBolsterIsLessSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchBackrestSideBolsterIsMoreSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchHeadrestIsBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchHeadrestIsDownEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchHeadrestIsForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchHeadrestIsUpEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchIsBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchIsCoolerEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchIsDownEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchIsForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchIsTiltBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchIsTiltForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchIsUpEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchIsWarmerEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchMassageIsDecreaseEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchMassageIsIncreaseEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchSeatingIsBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideSwitchSeatingIsForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1DriverSideTilt: MustGetColumn("Float32"),
		FieldCabinSeatRow1MiddleAirbagIsDeployed: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleBackrestLumbarHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow1MiddleBackrestLumbarSupport: MustGetColumn("Float32"),
		FieldCabinSeatRow1MiddleBackrestRecline: MustGetColumn("Float32"),
		FieldCabinSeatRow1MiddleBackrestSideBolsterSupport: MustGetColumn("Float32"),
		FieldCabinSeatRow1MiddleHeadrestAngle: MustGetColumn("Float32"),
		FieldCabinSeatRow1MiddleHeadrestHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow1MiddleHeating: MustGetColumn("Int32"),
		FieldCabinSeatRow1MiddleHeatingCooling: MustGetColumn("Int32"),
		FieldCabinSeatRow1MiddleHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow1MiddleIsBelted: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleIsOccupied: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleMassage: MustGetColumn("UInt32"),
		FieldCabinSeatRow1MiddleOccupantIdentifierIssuer: MustGetColumn("String"),
		FieldCabinSeatRow1MiddleOccupantIdentifierSubject: MustGetColumn("String"),
		FieldCabinSeatRow1MiddlePosition: MustGetColumn("UInt32"),
		FieldCabinSeatRow1MiddleSeatBeltHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow1MiddleSeatingLength: MustGetColumn("UInt32"),
		FieldCabinSeatRow1MiddleSwitchBackrestIsReclineBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchBackrestIsReclineForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchBackrestLumbarIsDownEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchBackrestLumbarIsLessSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchBackrestLumbarIsMoreSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchBackrestLumbarIsUpEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchBackrestSideBolsterIsLessSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchBackrestSideBolsterIsMoreSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchHeadrestIsBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchHeadrestIsDownEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchHeadrestIsForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchHeadrestIsUpEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchIsBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchIsCoolerEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchIsDownEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchIsForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchIsTiltBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchIsTiltForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchIsUpEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchIsWarmerEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchMassageIsDecreaseEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchMassageIsIncreaseEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchSeatingIsBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleSwitchSeatingIsForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1MiddleTilt: MustGetColumn("Float32"),
		FieldCabinSeatRow1PassengerSideAirbagIsDeployed: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideBackrestLumbarHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow1PassengerSideBackrestLumbarSupport: MustGetColumn("Float32"),
		FieldCabinSeatRow1PassengerSideBackrestRecline: MustGetColumn("Float32"),
		FieldCabinSeatRow1PassengerSideBackrestSideBolsterSupport: MustGetColumn("Float32"),
		FieldCabinSeatRow1PassengerSideHeadrestAngle: MustGetColumn("Float32"),
		FieldCabinSeatRow1PassengerSideHeadrestHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow1PassengerSideHeating: MustGetColumn("Int32"),
		FieldCabinSeatRow1PassengerSideHeatingCooling: MustGetColumn("Int32"),
		FieldCabinSeatRow1PassengerSideHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow1PassengerSideIsBelted: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideIsOccupied: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideMassage: MustGetColumn("UInt32"),
		FieldCabinSeatRow1PassengerSideOccupantIdentifierIssuer: MustGetColumn("String"),
		FieldCabinSeatRow1PassengerSideOccupantIdentifierSubject: MustGetColumn("String"),
		FieldCabinSeatRow1PassengerSidePosition: MustGetColumn("UInt32"),
		FieldCabinSeatRow1PassengerSideSeatBeltHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow1PassengerSideSeatingLength: MustGetColumn("UInt32"),
		FieldCabinSeatRow1PassengerSideSwitchBackrestIsReclineBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchBackrestIsReclineForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchBackrestLumbarIsDownEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchBackrestLumbarIsLessSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchBackrestLumbarIsMoreSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchBackrestLumbarIsUpEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchBackrestSideBolsterIsLessSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchBackrestSideBolsterIsMoreSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchHeadrestIsBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchHeadrestIsDownEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchHeadrestIsForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchHeadrestIsUpEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchIsBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchIsCoolerEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchIsDownEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchIsForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchIsTiltBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchIsTiltForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchIsUpEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchIsWarmerEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchMassageIsDecreaseEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchMassageIsIncreaseEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchSeatingIsBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideSwitchSeatingIsForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow1PassengerSideTilt: MustGetColumn("Float32"),
		FieldCabinSeatRow2DriverSideAirbagIsDeployed: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideBackrestLumbarHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow2DriverSideBackrestLumbarSupport: MustGetColumn("Float32"),
		FieldCabinSeatRow2DriverSideBackrestRecline: MustGetColumn("Float32"),
		FieldCabinSeatRow2DriverSideBackrestSideBolsterSupport: MustGetColumn("Float32"),
		FieldCabinSeatRow2DriverSideHeadrestAngle: MustGetColumn("Float32"),
		FieldCabinSeatRow2DriverSideHeadrestHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow2DriverSideHeating: MustGetColumn("Int32"),
		FieldCabinSeatRow2DriverSideHeatingCooling: MustGetColumn("Int32"),
		FieldCabinSeatRow2DriverSideHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow2DriverSideIsBelted: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideIsOccupied: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideMassage: MustGetColumn("UInt32"),
		FieldCabinSeatRow2DriverSideOccupantIdentifierIssuer: MustGetColumn("String"),
		FieldCabinSeatRow2DriverSideOccupantIdentifierSubject: MustGetColumn("String"),
		FieldCabinSeatRow2DriverSidePosition: MustGetColumn("UInt32"),
		FieldCabinSeatRow2DriverSideSeatBeltHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow2DriverSideSeatingLength: MustGetColumn("UInt32"),
		FieldCabinSeatRow2DriverSideSwitchBackrestIsReclineBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchBackrestIsReclineForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchBackrestLumbarIsDownEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchBackrestLumbarIsLessSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchBackrestLumbarIsMoreSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchBackrestLumbarIsUpEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchBackrestSideBolsterIsLessSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchBackrestSideBolsterIsMoreSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchHeadrestIsBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchHeadrestIsDownEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchHeadrestIsForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchHeadrestIsUpEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchIsBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchIsCoolerEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchIsDownEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchIsForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchIsTiltBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchIsTiltForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchIsUpEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchIsWarmerEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchMassageIsDecreaseEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchMassageIsIncreaseEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchSeatingIsBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideSwitchSeatingIsForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2DriverSideTilt: MustGetColumn("Float32"),
		FieldCabinSeatRow2MiddleAirbagIsDeployed: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleBackrestLumbarHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow2MiddleBackrestLumbarSupport: MustGetColumn("Float32"),
		FieldCabinSeatRow2MiddleBackrestRecline: MustGetColumn("Float32"),
		FieldCabinSeatRow2MiddleBackrestSideBolsterSupport: MustGetColumn("Float32"),
		FieldCabinSeatRow2MiddleHeadrestAngle: MustGetColumn("Float32"),
		FieldCabinSeatRow2MiddleHeadrestHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow2MiddleHeating: MustGetColumn("Int32"),
		FieldCabinSeatRow2MiddleHeatingCooling: MustGetColumn("Int32"),
		FieldCabinSeatRow2MiddleHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow2MiddleIsBelted: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleIsOccupied: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleMassage: MustGetColumn("UInt32"),
		FieldCabinSeatRow2MiddleOccupantIdentifierIssuer: MustGetColumn("String"),
		FieldCabinSeatRow2MiddleOccupantIdentifierSubject: MustGetColumn("String"),
		FieldCabinSeatRow2MiddlePosition: MustGetColumn("UInt32"),
		FieldCabinSeatRow2MiddleSeatBeltHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow2MiddleSeatingLength: MustGetColumn("UInt32"),
		FieldCabinSeatRow2MiddleSwitchBackrestIsReclineBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchBackrestIsReclineForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchBackrestLumbarIsDownEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchBackrestLumbarIsLessSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchBackrestLumbarIsMoreSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchBackrestLumbarIsUpEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchBackrestSideBolsterIsLessSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchBackrestSideBolsterIsMoreSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchHeadrestIsBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchHeadrestIsDownEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchHeadrestIsForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchHeadrestIsUpEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchIsBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchIsCoolerEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchIsDownEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchIsForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchIsTiltBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchIsTiltForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchIsUpEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchIsWarmerEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchMassageIsDecreaseEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchMassageIsIncreaseEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchSeatingIsBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleSwitchSeatingIsForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2MiddleTilt: MustGetColumn("Float32"),
		FieldCabinSeatRow2PassengerSideAirbagIsDeployed: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideBackrestLumbarHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow2PassengerSideBackrestLumbarSupport: MustGetColumn("Float32"),
		FieldCabinSeatRow2PassengerSideBackrestRecline: MustGetColumn("Float32"),
		FieldCabinSeatRow2PassengerSideBackrestSideBolsterSupport: MustGetColumn("Float32"),
		FieldCabinSeatRow2PassengerSideHeadrestAngle: MustGetColumn("Float32"),
		FieldCabinSeatRow2PassengerSideHeadrestHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow2PassengerSideHeating: MustGetColumn("Int32"),
		FieldCabinSeatRow2PassengerSideHeatingCooling: MustGetColumn("Int32"),
		FieldCabinSeatRow2PassengerSideHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow2PassengerSideIsBelted: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideIsOccupied: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideMassage: MustGetColumn("UInt32"),
		FieldCabinSeatRow2PassengerSideOccupantIdentifierIssuer: MustGetColumn("String"),
		FieldCabinSeatRow2PassengerSideOccupantIdentifierSubject: MustGetColumn("String"),
		FieldCabinSeatRow2PassengerSidePosition: MustGetColumn("UInt32"),
		FieldCabinSeatRow2PassengerSideSeatBeltHeight: MustGetColumn("UInt32"),
		FieldCabinSeatRow2PassengerSideSeatingLength: MustGetColumn("UInt32"),
		FieldCabinSeatRow2PassengerSideSwitchBackrestIsReclineBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchBackrestIsReclineForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchBackrestLumbarIsDownEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchBackrestLumbarIsLessSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchBackrestLumbarIsMoreSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchBackrestLumbarIsUpEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchBackrestSideBolsterIsLessSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchBackrestSideBolsterIsMoreSupportEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchHeadrestIsBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchHeadrestIsDownEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchHeadrestIsForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchHeadrestIsUpEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchIsBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchIsCoolerEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchIsDownEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchIsForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchIsTiltBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchIsTiltForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchIsUpEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchIsWarmerEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchMassageIsDecreaseEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchMassageIsIncreaseEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchSeatingIsBackwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideSwitchSeatingIsForwardEngaged: MustGetColumn("UInt8"),
		FieldCabinSeatRow2PassengerSideTilt: MustGetColumn("Float32"),
		FieldCabinSeatRowCount: MustGetColumn("UInt32"),
		FieldCabinSunroofPosition: MustGetColumn("Int32"),
		FieldCabinSunroofShadeIsOpen: MustGetColumn("UInt8"),
		FieldCabinSunroofShadePosition: MustGetColumn("UInt32"),
		FieldCabinSunroofShadeSwitch: MustGetColumn("String"),
		FieldCabinSunroofSwitch: MustGetColumn("String"),
		FieldCargoVolume: MustGetColumn("Float32"),
		FieldChassisAcceleratorPedalPosition: MustGetColumn("UInt32"),
		FieldChassisAxleCount: MustGetColumn("UInt32"),
		FieldChassisAxleRow1AxleWidth: MustGetColumn("UInt32"),
		FieldChassisAxleRow1SteeringAngle: MustGetColumn("Float32"),
		FieldChassisAxleRow1TireAspectRatio: MustGetColumn("UInt32"),
		FieldChassisAxleRow1TireDiameter: MustGetColumn("Float32"),
		FieldChassisAxleRow1TireWidth: MustGetColumn("UInt32"),
		FieldChassisAxleRow1TrackWidth: MustGetColumn("UInt32"),
		FieldChassisAxleRow1TreadWidth: MustGetColumn("UInt32"),
		FieldChassisAxleRow1WheelCount: MustGetColumn("UInt32"),
		FieldChassisAxleRow1WheelDiameter: MustGetColumn("Float32"),
		FieldChassisAxleRow1WheelLeftBrakeFluidLevel: MustGetColumn("UInt32"),
		FieldChassisAxleRow1WheelLeftBrakeIsBrakesWorn: MustGetColumn("UInt8"),
		FieldChassisAxleRow1WheelLeftBrakeIsFluidLevelLow: MustGetColumn("UInt8"),
		FieldChassisAxleRow1WheelLeftBrakePadWear: MustGetColumn("UInt32"),
		FieldChassisAxleRow1WheelLeftSpeed: MustGetColumn("Float32"),
		FieldChassisAxleRow1WheelLeftTireIsPressureLow: MustGetColumn("UInt8"),
		FieldChassisAxleRow1WheelLeftTirePressure: MustGetColumn("UInt32"),
		FieldChassisAxleRow1WheelLeftTireTemperature: MustGetColumn("Float32"),
		FieldChassisAxleRow1WheelRightBrakeFluidLevel: MustGetColumn("UInt32"),
		FieldChassisAxleRow1WheelRightBrakeIsBrakesWorn: MustGetColumn("UInt8"),
		FieldChassisAxleRow1WheelRightBrakeIsFluidLevelLow: MustGetColumn("UInt8"),
		FieldChassisAxleRow1WheelRightBrakePadWear: MustGetColumn("UInt32"),
		FieldChassisAxleRow1WheelRightSpeed: MustGetColumn("Float32"),
		FieldChassisAxleRow1WheelRightTireIsPressureLow: MustGetColumn("UInt8"),
		FieldChassisAxleRow1WheelRightTirePressure: MustGetColumn("UInt32"),
		FieldChassisAxleRow1WheelRightTireTemperature: MustGetColumn("Float32"),
		FieldChassisAxleRow1WheelWidth: MustGetColumn("Float32"),
		FieldChassisAxleRow2AxleWidth: MustGetColumn("UInt32"),
		FieldChassisAxleRow2SteeringAngle: MustGetColumn("Float32"),
		FieldChassisAxleRow2TireAspectRatio: MustGetColumn("UInt32"),
		FieldChassisAxleRow2TireDiameter: MustGetColumn("Float32"),
		FieldChassisAxleRow2TireWidth: MustGetColumn("UInt32"),
		FieldChassisAxleRow2TrackWidth: MustGetColumn("UInt32"),
		FieldChassisAxleRow2TreadWidth: MustGetColumn("UInt32"),
		FieldChassisAxleRow2WheelCount: MustGetColumn("UInt32"),
		FieldChassisAxleRow2WheelDiameter: MustGetColumn("Float32"),
		FieldChassisAxleRow2WheelLeftBrakeFluidLevel: MustGetColumn("UInt32"),
		FieldChassisAxleRow2WheelLeftBrakeIsBrakesWorn: MustGetColumn("UInt8"),
		FieldChassisAxleRow2WheelLeftBrakeIsFluidLevelLow: MustGetColumn("UInt8"),
		FieldChassisAxleRow2WheelLeftBrakePadWear: MustGetColumn("UInt32"),
		FieldChassisAxleRow2WheelLeftSpeed: MustGetColumn("Float32"),
		FieldChassisAxleRow2WheelLeftTireIsPressureLow: MustGetColumn("UInt8"),
		FieldChassisAxleRow2WheelLeftTirePressure: MustGetColumn("UInt32"),
		FieldChassisAxleRow2WheelLeftTireTemperature: MustGetColumn("Float32"),
		FieldChassisAxleRow2WheelRightBrakeFluidLevel: MustGetColumn("UInt32"),
		FieldChassisAxleRow2WheelRightBrakeIsBrakesWorn: MustGetColumn("UInt8"),
		FieldChassisAxleRow2WheelRightBrakeIsFluidLevelLow: MustGetColumn("UInt8"),
		FieldChassisAxleRow2WheelRightBrakePadWear: MustGetColumn("UInt32"),
		FieldChassisAxleRow2WheelRightSpeed: MustGetColumn("Float32"),
		FieldChassisAxleRow2WheelRightTireIsPressureLow: MustGetColumn("UInt8"),
		FieldChassisAxleRow2WheelRightTirePressure: MustGetColumn("UInt32"),
		FieldChassisAxleRow2WheelRightTireTemperature: MustGetColumn("Float32"),
		FieldChassisAxleRow2WheelWidth: MustGetColumn("Float32"),
		FieldChassisBrakeIsDriverEmergencyBrakingDetected: MustGetColumn("UInt8"),
		FieldChassisBrakePedalPosition: MustGetColumn("UInt32"),
		FieldChassisParkingBrakeIsAutoApplyEnabled: MustGetColumn("UInt8"),
		FieldChassisParkingBrakeIsEngaged: MustGetColumn("UInt8"),
		FieldChassisSteeringWheelAngle: MustGetColumn("Int32"),
		FieldChassisSteeringWheelExtension: MustGetColumn("UInt32"),
		FieldChassisSteeringWheelHeatingCooling: MustGetColumn("Int32"),
		FieldChassisSteeringWheelTilt: MustGetColumn("UInt32"),
		FieldChassisWheelbase: MustGetColumn("UInt32"),
		FieldConnectivityIsConnectivityAvailable: MustGetColumn("UInt8"),
		FieldCurbWeight: MustGetColumn("UInt32"),
		FieldCurrentLocationAltitude: MustGetColumn("Float64"),
		FieldCurrentLocationGNSSReceiverFixType: MustGetColumn("String"),
		FieldCurrentLocationGNSSReceiverMountingPositionX: MustGetColumn("Int32"),
		FieldCurrentLocationGNSSReceiverMountingPositionY: MustGetColumn("Int32"),
		FieldCurrentLocationGNSSReceiverMountingPositionZ: MustGetColumn("Int32"),
		FieldCurrentLocationHeading: MustGetColumn("Float64"),
		FieldCurrentLocationHorizontalAccuracy: MustGetColumn("Float64"),
		FieldCurrentLocationLatitude: MustGetColumn("Float64"),
		FieldCurrentLocationLongitude: MustGetColumn("Float64"),
		FieldCurrentLocationTimestamp: MustGetColumn("String"),
		FieldCurrentLocationVerticalAccuracy: MustGetColumn("Float64"),
		FieldCurrentOverallWeight: MustGetColumn("UInt32"),
		FieldDIMODefinitionID: MustGetColumn("String"),
		FieldDIMOSource: MustGetColumn("String"),
		FieldDIMOType: MustGetColumn("String"),
		FieldDIMOVehicleID: MustGetColumn("String"),
		FieldDriverAttentiveProbability: MustGetColumn("Float32"),
		FieldDriverDistractionLevel: MustGetColumn("Float32"),
		FieldDriverFatigueLevel: MustGetColumn("Float32"),
		FieldDriverHeartRate: MustGetColumn("UInt32"),
		FieldDriverIdentifierIssuer: MustGetColumn("String"),
		FieldDriverIdentifierSubject: MustGetColumn("String"),
		FieldDriverIsEyesOnRoad: MustGetColumn("UInt8"),
		FieldDriverIsHandsOnWheel: MustGetColumn("UInt8"),
		FieldEmissionsCO2: MustGetColumn("Int32"),
		FieldExteriorAirTemperature: MustGetColumn("Float32"),
		FieldExteriorHumidity: MustGetColumn("Float32"),
		FieldExteriorLightIntensity: MustGetColumn("Float32"),
		FieldGrossWeight: MustGetColumn("UInt32"),
		FieldHeight: MustGetColumn("UInt32"),
		FieldIsAutoPowerOptimize: MustGetColumn("UInt8"),
		FieldIsBrokenDown: MustGetColumn("UInt8"),
		FieldIsMoving: MustGetColumn("UInt8"),
		FieldLength: MustGetColumn("UInt32"),
		FieldLowVoltageBatteryCurrentCurrent: MustGetColumn("Float32"),
		FieldLowVoltageBatteryCurrentVoltage: MustGetColumn("Float32"),
		FieldLowVoltageBatteryNominalCapacity: MustGetColumn("UInt32"),
		FieldLowVoltageBatteryNominalVoltage: MustGetColumn("UInt32"),
		FieldLowVoltageSystemState: MustGetColumn("String"),
		FieldMaxTowBallWeight: MustGetColumn("UInt32"),
		FieldMaxTowWeight: MustGetColumn("UInt32"),
		FieldOBDAbsoluteLoad: MustGetColumn("Float32"),
		FieldOBDAcceleratorPositionD: MustGetColumn("Float32"),
		FieldOBDAcceleratorPositionE: MustGetColumn("Float32"),
		FieldOBDAcceleratorPositionF: MustGetColumn("Float32"),
		FieldOBDAirStatus: MustGetColumn("String"),
		FieldOBDAmbientAirTemperature: MustGetColumn("Float32"),
		FieldOBDBarometricPressure: MustGetColumn("Float32"),
		FieldOBDCatalystBank1Temperature1: MustGetColumn("Float32"),
		FieldOBDCatalystBank1Temperature2: MustGetColumn("Float32"),
		FieldOBDCatalystBank2Temperature1: MustGetColumn("Float32"),
		FieldOBDCatalystBank2Temperature2: MustGetColumn("Float32"),
		FieldOBDCommandedEGR: MustGetColumn("Float32"),
		FieldOBDCommandedEVAP: MustGetColumn("Float32"),
		FieldOBDCommandedEquivalenceRatio: MustGetColumn("Float32"),
		FieldOBDControlModuleVoltage: MustGetColumn("Float32"),
		FieldOBDCoolantTemperature: MustGetColumn("Float32"),
		FieldOBDDTCList: MustGetColumn("Array(String)"),
		FieldOBDDTCListsize0: MustGetColumn("UInt64"),
		FieldOBDDistanceSinceDTCClear: MustGetColumn("Float32"),
		FieldOBDDistanceWithMIL: MustGetColumn("Float32"),
		FieldOBDDriveCycleStatusDTCCount: MustGetColumn("UInt32"),
		FieldOBDDriveCycleStatusIgnitionType: MustGetColumn("String"),
		FieldOBDDriveCycleStatusIsMILOn: MustGetColumn("UInt8"),
		FieldOBDEGRError: MustGetColumn("Float32"),
		FieldOBDEVAPVaporPressure: MustGetColumn("Float32"),
		FieldOBDEVAPVaporPressureAbsolute: MustGetColumn("Float32"),
		FieldOBDEVAPVaporPressureAlternate: MustGetColumn("Float32"),
		FieldOBDEngineLoad: MustGetColumn("Float32"),
		FieldOBDEngineSpeed: MustGetColumn("Float32"),
		FieldOBDEthanolPercent: MustGetColumn("Float32"),
		FieldOBDFreezeDTC: MustGetColumn("String"),
		FieldOBDFuelInjectionTiming: MustGetColumn("Float32"),
		FieldOBDFuelLevel: MustGetColumn("Float32"),
		FieldOBDFuelPressure: MustGetColumn("Float32"),
		FieldOBDFuelRailPressureAbsolute: MustGetColumn("Float32"),
		FieldOBDFuelRailPressureDirect: MustGetColumn("Float32"),
		FieldOBDFuelRailPressureVac: MustGetColumn("Float32"),
		FieldOBDFuelRate: MustGetColumn("Float32"),
		FieldOBDFuelStatus: MustGetColumn("String"),
		FieldOBDFuelType: MustGetColumn("UInt32"),
		FieldOBDHybridBatteryRemaining: MustGetColumn("Float32"),
		FieldOBDIntakeTemp: MustGetColumn("Float32"),
		FieldOBDIsPTOActive: MustGetColumn("UInt8"),
		FieldOBDLongTermFuelTrim1: MustGetColumn("Float32"),
		FieldOBDLongTermFuelTrim2: MustGetColumn("Float32"),
		FieldOBDLongTermO2Trim1: MustGetColumn("Float32"),
		FieldOBDLongTermO2Trim2: MustGetColumn("Float32"),
		FieldOBDLongTermO2Trim3: MustGetColumn("Float32"),
		FieldOBDLongTermO2Trim4: MustGetColumn("Float32"),
		FieldOBDMAF: MustGetColumn("Float32"),
		FieldOBDMAP: MustGetColumn("Float32"),
		FieldOBDMaxMAF: MustGetColumn("Float32"),
		FieldOBDO2Sensor1ShortTermFuelTrim: MustGetColumn("Float32"),
		FieldOBDO2Sensor1Voltage: MustGetColumn("Float32"),
		FieldOBDO2Sensor2ShortTermFuelTrim: MustGetColumn("Float32"),
		FieldOBDO2Sensor2Voltage: MustGetColumn("Float32"),
		FieldOBDO2Sensor3ShortTermFuelTrim: MustGetColumn("Float32"),
		FieldOBDO2Sensor3Voltage: MustGetColumn("Float32"),
		FieldOBDO2Sensor4ShortTermFuelTrim: MustGetColumn("Float32"),
		FieldOBDO2Sensor4Voltage: MustGetColumn("Float32"),
		FieldOBDO2Sensor5ShortTermFuelTrim: MustGetColumn("Float32"),
		FieldOBDO2Sensor5Voltage: MustGetColumn("Float32"),
		FieldOBDO2Sensor6ShortTermFuelTrim: MustGetColumn("Float32"),
		FieldOBDO2Sensor6Voltage: MustGetColumn("Float32"),
		FieldOBDO2Sensor7ShortTermFuelTrim: MustGetColumn("Float32"),
		FieldOBDO2Sensor7Voltage: MustGetColumn("Float32"),
		FieldOBDO2Sensor8ShortTermFuelTrim: MustGetColumn("Float32"),
		FieldOBDO2Sensor8Voltage: MustGetColumn("Float32"),
		FieldOBDO2WRSensor1Current: MustGetColumn("Float32"),
		FieldOBDO2WRSensor1Lambda: MustGetColumn("Float32"),
		FieldOBDO2WRSensor1Voltage: MustGetColumn("Float32"),
		FieldOBDO2WRSensor2Current: MustGetColumn("Float32"),
		FieldOBDO2WRSensor2Lambda: MustGetColumn("Float32"),
		FieldOBDO2WRSensor2Voltage: MustGetColumn("Float32"),
		FieldOBDO2WRSensor3Current: MustGetColumn("Float32"),
		FieldOBDO2WRSensor3Lambda: MustGetColumn("Float32"),
		FieldOBDO2WRSensor3Voltage: MustGetColumn("Float32"),
		FieldOBDO2WRSensor4Current: MustGetColumn("Float32"),
		FieldOBDO2WRSensor4Lambda: MustGetColumn("Float32"),
		FieldOBDO2WRSensor4Voltage: MustGetColumn("Float32"),
		FieldOBDO2WRSensor5Current: MustGetColumn("Float32"),
		FieldOBDO2WRSensor5Lambda: MustGetColumn("Float32"),
		FieldOBDO2WRSensor5Voltage: MustGetColumn("Float32"),
		FieldOBDO2WRSensor6Current: MustGetColumn("Float32"),
		FieldOBDO2WRSensor6Lambda: MustGetColumn("Float32"),
		FieldOBDO2WRSensor6Voltage: MustGetColumn("Float32"),
		FieldOBDO2WRSensor7Current: MustGetColumn("Float32"),
		FieldOBDO2WRSensor7Lambda: MustGetColumn("Float32"),
		FieldOBDO2WRSensor7Voltage: MustGetColumn("Float32"),
		FieldOBDO2WRSensor8Current: MustGetColumn("Float32"),
		FieldOBDO2WRSensor8Lambda: MustGetColumn("Float32"),
		FieldOBDO2WRSensor8Voltage: MustGetColumn("Float32"),
		FieldOBDOBDStandards: MustGetColumn("UInt32"),
		FieldOBDOilTemperature: MustGetColumn("Float32"),
		FieldOBDOxygenSensorsIn2Banks: MustGetColumn("UInt32"),
		FieldOBDOxygenSensorsIn4Banks: MustGetColumn("UInt32"),
		FieldOBDPidsA: MustGetColumn("Array(String)"),
		FieldOBDPidsAsize0: MustGetColumn("UInt64"),
		FieldOBDPidsB: MustGetColumn("Array(String)"),
		FieldOBDPidsBsize0: MustGetColumn("UInt64"),
		FieldOBDPidsC: MustGetColumn("Array(String)"),
		FieldOBDPidsCsize0: MustGetColumn("UInt64"),
		FieldOBDRelativeAcceleratorPosition: MustGetColumn("Float32"),
		FieldOBDRelativeThrottlePosition: MustGetColumn("Float32"),
		FieldOBDRunTime: MustGetColumn("Float32"),
		FieldOBDRunTimeMIL: MustGetColumn("Float32"),
		FieldOBDShortTermFuelTrim1: MustGetColumn("Float32"),
		FieldOBDShortTermFuelTrim2: MustGetColumn("Float32"),
		FieldOBDShortTermO2Trim1: MustGetColumn("Float32"),
		FieldOBDShortTermO2Trim2: MustGetColumn("Float32"),
		FieldOBDShortTermO2Trim3: MustGetColumn("Float32"),
		FieldOBDShortTermO2Trim4: MustGetColumn("Float32"),
		FieldOBDSpeed: MustGetColumn("Float32"),
		FieldOBDStatusDTCCount: MustGetColumn("UInt32"),
		FieldOBDStatusIgnitionType: MustGetColumn("String"),
		FieldOBDStatusIsMILOn: MustGetColumn("UInt8"),
		FieldOBDThrottleActuator: MustGetColumn("Float32"),
		FieldOBDThrottlePosition: MustGetColumn("Float32"),
		FieldOBDThrottlePositionB: MustGetColumn("Float32"),
		FieldOBDThrottlePositionC: MustGetColumn("Float32"),
		FieldOBDTimeSinceDTCCleared: MustGetColumn("Float32"),
		FieldOBDTimingAdvance: MustGetColumn("Float32"),
		FieldOBDWarmupsSinceDTCClear: MustGetColumn("UInt32"),
		FieldPowerOptimizeLevel: MustGetColumn("UInt32"),
		FieldPowertrainAccumulatedBrakingEnergy: MustGetColumn("Float32"),
		FieldPowertrainCombustionEngineAspirationType: MustGetColumn("String"),
		FieldPowertrainCombustionEngineBore: MustGetColumn("Float32"),
		FieldPowertrainCombustionEngineCompressionRatio: MustGetColumn("String"),
		FieldPowertrainCombustionEngineConfiguration: MustGetColumn("String"),
		FieldPowertrainCombustionEngineDieselExhaustFluidCapacity: MustGetColumn("Float32"),
		FieldPowertrainCombustionEngineDieselExhaustFluidIsLevelLow: MustGetColumn("UInt8"),
		FieldPowertrainCombustionEngineDieselExhaustFluidLevel: MustGetColumn("UInt32"),
		FieldPowertrainCombustionEngineDieselExhaustFluidRange: MustGetColumn("UInt32"),
		FieldPowertrainCombustionEngineDieselParticulateFilterDeltaPressure: MustGetColumn("Float32"),
		FieldPowertrainCombustionEngineDieselParticulateFilterInletTemperature: MustGetColumn("Float32"),
		FieldPowertrainCombustionEngineDieselParticulateFilterOutletTemperature: MustGetColumn("Float32"),
		FieldPowertrainCombustionEngineDisplacement: MustGetColumn("UInt32"),
		FieldPowertrainCombustionEngineECT: MustGetColumn("Float32"),
		FieldPowertrainCombustionEngineEOP: MustGetColumn("UInt32"),
		FieldPowertrainCombustionEngineEOT: MustGetColumn("Float32"),
		FieldPowertrainCombustionEngineEngineCode: MustGetColumn("String"),
		FieldPowertrainCombustionEngineEngineCoolantCapacity: MustGetColumn("Float32"),
		FieldPowertrainCombustionEngineEngineHours: MustGetColumn("Float32"),
		FieldPowertrainCombustionEngineEngineOilCapacity: MustGetColumn("Float32"),
		FieldPowertrainCombustionEngineEngineOilLevel: MustGetColumn("String"),
		FieldPowertrainCombustionEngineIdleHours: MustGetColumn("Float32"),
		FieldPowertrainCombustionEngineIsRunning: MustGetColumn("UInt8"),
		FieldPowertrainCombustionEngineMAF: MustGetColumn("UInt32"),
		FieldPowertrainCombustionEngineMAP: MustGetColumn("UInt32"),
		FieldPowertrainCombustionEngineMaxPower: MustGetColumn("UInt32"),
		FieldPowertrainCombustionEngineMaxTorque: MustGetColumn("UInt32"),
		FieldPowertrainCombustionEngineNumberOfCylinders: MustGetColumn("UInt32"),
		FieldPowertrainCombustionEngineNumberOfValvesPerCylinder: MustGetColumn("UInt32"),
		FieldPowertrainCombustionEngineOilLifeRemaining: MustGetColumn("Int32"),
		FieldPowertrainCombustionEnginePower: MustGetColumn("UInt32"),
		FieldPowertrainCombustionEngineSpeed: MustGetColumn("UInt32"),
		FieldPowertrainCombustionEngineStrokeLength: MustGetColumn("Float32"),
		FieldPowertrainCombustionEngineTPS: MustGetColumn("UInt32"),
		FieldPowertrainCombustionEngineTorque: MustGetColumn("Int32"),
		FieldPowertrainElectricMotorCoolantTemperature: MustGetColumn("Float32"),
		FieldPowertrainElectricMotorEngineCode: MustGetColumn("String"),
		FieldPowertrainElectricMotorMaxPower: MustGetColumn("UInt32"),
		FieldPowertrainElectricMotorMaxRegenPower: MustGetColumn("UInt32"),
		FieldPowertrainElectricMotorMaxRegenTorque: MustGetColumn("UInt32"),
		FieldPowertrainElectricMotorMaxTorque: MustGetColumn("UInt32"),
		FieldPowertrainElectricMotorPower: MustGetColumn("Int32"),
		FieldPowertrainElectricMotorSpeed: MustGetColumn("Int32"),
		FieldPowertrainElectricMotorTemperature: MustGetColumn("Float32"),
		FieldPowertrainElectricMotorTorque: MustGetColumn("Int32"),
		FieldPowertrainFuelSystemAbsoluteLevel: MustGetColumn("Float32"),
		FieldPowertrainFuelSystemAverageConsumption: MustGetColumn("Float32"),
		FieldPowertrainFuelSystemConsumptionSinceStart: MustGetColumn("Float32"),
		FieldPowertrainFuelSystemHybridType: MustGetColumn("String"),
		FieldPowertrainFuelSystemInstantConsumption: MustGetColumn("Float32"),
		FieldPowertrainFuelSystemIsEngineStopStartEnabled: MustGetColumn("UInt8"),
		FieldPowertrainFuelSystemIsFuelLevelLow: MustGetColumn("UInt8"),
		FieldPowertrainFuelSystemIsFuelPortFlapOpen: MustGetColumn("UInt8"),
		FieldPowertrainFuelSystemRange: MustGetColumn("UInt32"),
		FieldPowertrainFuelSystemRefuelPortPosition: MustGetColumn("Array(String)"),
		FieldPowertrainFuelSystemRefuelPortPositionsize0: MustGetColumn("UInt64"),
		FieldPowertrainFuelSystemRelativeLevel: MustGetColumn("UInt32"),
		FieldPowertrainFuelSystemSupportedFuel: MustGetColumn("Array(String)"),
		FieldPowertrainFuelSystemSupportedFuelTypes: MustGetColumn("Array(String)"),
		FieldPowertrainFuelSystemSupportedFuelTypessize0: MustGetColumn("UInt64"),
		FieldPowertrainFuelSystemSupportedFuelsize0: MustGetColumn("UInt64"),
		FieldPowertrainFuelSystemTankCapacity: MustGetColumn("Float32"),
		FieldPowertrainFuelSystemTimeRemaining: MustGetColumn("UInt32"),
		FieldPowertrainIsAutoPowerOptimize: MustGetColumn("UInt8"),
		FieldPowertrainPowerOptimizeLevel: MustGetColumn("UInt32"),
		FieldPowertrainRange: MustGetColumn("UInt32"),
		FieldPowertrainTimeRemaining: MustGetColumn("UInt32"),
		FieldPowertrainTractionBatteryAccumulatedChargedEnergy: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryAccumulatedChargedThroughput: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryAccumulatedConsumedEnergy: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryAccumulatedConsumedThroughput: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryCellVoltageCellVoltages: MustGetColumn("Array(Float32)"),
		FieldPowertrainTractionBatteryCellVoltageCellVoltagessize0: MustGetColumn("UInt64"),
		FieldPowertrainTractionBatteryCellVoltageIdMax: MustGetColumn("UInt32"),
		FieldPowertrainTractionBatteryCellVoltageIdMin: MustGetColumn("UInt32"),
		FieldPowertrainTractionBatteryCellVoltageMax: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryCellVoltageMin: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryChargingAveragePower: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryChargingChargeCurrentDC: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryChargingChargeCurrentPhase1: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryChargingChargeCurrentPhase2: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryChargingChargeCurrentPhase3: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryChargingChargeLimit: MustGetColumn("UInt32"),
		FieldPowertrainTractionBatteryChargingChargePlugType: MustGetColumn("Array(String)"),
		FieldPowertrainTractionBatteryChargingChargePlugTypesize0: MustGetColumn("UInt64"),
		FieldPowertrainTractionBatteryChargingChargePortFlap: MustGetColumn("String"),
		FieldPowertrainTractionBatteryChargingChargePortPosition: MustGetColumn("Array(String)"),
		FieldPowertrainTractionBatteryChargingChargePortPositionsize0: MustGetColumn("UInt64"),
		FieldPowertrainTractionBatteryChargingChargePortType: MustGetColumn("Array(String)"),
		FieldPowertrainTractionBatteryChargingChargePortTypesize0: MustGetColumn("UInt64"),
		FieldPowertrainTractionBatteryChargingChargeRate: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryChargingChargeVoltageDC: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryChargingChargeVoltagePhase1: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryChargingChargeVoltagePhase2: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryChargingChargeVoltagePhase3: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryChargingEvseId: MustGetColumn("String"),
		FieldPowertrainTractionBatteryChargingIsChargePortFlapOpen: MustGetColumn("UInt8"),
		FieldPowertrainTractionBatteryChargingIsCharging: MustGetColumn("UInt8"),
		FieldPowertrainTractionBatteryChargingIsChargingCableConnected: MustGetColumn("UInt8"),
		FieldPowertrainTractionBatteryChargingIsChargingCableLocked: MustGetColumn("UInt8"),
		FieldPowertrainTractionBatteryChargingIsDischarging: MustGetColumn("UInt8"),
		FieldPowertrainTractionBatteryChargingLocationAltitude: MustGetColumn("Float64"),
		FieldPowertrainTractionBatteryChargingLocationLatitude: MustGetColumn("Float64"),
		FieldPowertrainTractionBatteryChargingLocationLongitude: MustGetColumn("Float64"),
		FieldPowertrainTractionBatteryChargingMaxPower: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryChargingMaximumChargingCurrentDC: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryChargingMaximumChargingCurrentPhase1: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryChargingMaximumChargingCurrentPhase2: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryChargingMaximumChargingCurrentPhase3: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryChargingMode: MustGetColumn("String"),
		FieldPowertrainTractionBatteryChargingPowerLoss: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryChargingStartStopCharging: MustGetColumn("String"),
		FieldPowertrainTractionBatteryChargingTemperature: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryChargingTimeToComplete: MustGetColumn("UInt32"),
		FieldPowertrainTractionBatteryChargingTimerMode: MustGetColumn("String"),
		FieldPowertrainTractionBatteryChargingTimerTime: MustGetColumn("String"),
		FieldPowertrainTractionBatteryCurrentCurrent: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryCurrentPower: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryCurrentVoltage: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryDCDCPowerLoss: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryDCDCTemperature: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryErrorCodes: MustGetColumn("Array(String)"),
		FieldPowertrainTractionBatteryErrorCodessize0: MustGetColumn("UInt64"),
		FieldPowertrainTractionBatteryGrossCapacity: MustGetColumn("UInt32"),
		FieldPowertrainTractionBatteryId: MustGetColumn("String"),
		FieldPowertrainTractionBatteryIsGroundConnected: MustGetColumn("UInt8"),
		FieldPowertrainTractionBatteryIsPowerConnected: MustGetColumn("UInt8"),
		FieldPowertrainTractionBatteryMaxVoltage: MustGetColumn("UInt32"),
		FieldPowertrainTractionBatteryNetCapacity: MustGetColumn("UInt32"),
		FieldPowertrainTractionBatteryNominalVoltage: MustGetColumn("UInt32"),
		FieldPowertrainTractionBatteryPowerLoss: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryProductionDate: MustGetColumn("String"),
		FieldPowertrainTractionBatteryRange: MustGetColumn("UInt32"),
		FieldPowertrainTractionBatteryStateOfChargeCurrent: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryStateOfChargeCurrentEnergy: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryStateOfChargeDisplayed: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryStateOfHealth: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryTemperatureAverage: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryTemperatureCellTemperature: MustGetColumn("Array(Float32)"),
		FieldPowertrainTractionBatteryTemperatureCellTemperaturesize0: MustGetColumn("UInt64"),
		FieldPowertrainTractionBatteryTemperatureMax: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryTemperatureMin: MustGetColumn("Float32"),
		FieldPowertrainTractionBatteryTimeRemaining: MustGetColumn("UInt32"),
		FieldPowertrainTransmissionClutchEngagement: MustGetColumn("Float32"),
		FieldPowertrainTransmissionClutchWear: MustGetColumn("UInt32"),
		FieldPowertrainTransmissionCurrentGear: MustGetColumn("Int32"),
		FieldPowertrainTransmissionDiffLockFrontEngagement: MustGetColumn("Float32"),
		FieldPowertrainTransmissionDiffLockRearEngagement: MustGetColumn("Float32"),
		FieldPowertrainTransmissionDriveType: MustGetColumn("String"),
		FieldPowertrainTransmissionGearChangeMode: MustGetColumn("String"),
		FieldPowertrainTransmissionGearCount: MustGetColumn("Int32"),
		FieldPowertrainTransmissionIsElectricalPowertrainEngaged: MustGetColumn("UInt8"),
		FieldPowertrainTransmissionIsLowRangeEngaged: MustGetColumn("UInt8"),
		FieldPowertrainTransmissionIsParkLockEngaged: MustGetColumn("UInt8"),
		FieldPowertrainTransmissionPerformanceMode: MustGetColumn("String"),
		FieldPowertrainTransmissionSelectedGear: MustGetColumn("Int32"),
		FieldPowertrainTransmissionTemperature: MustGetColumn("Float32"),
		FieldPowertrainTransmissionTorqueDistribution: MustGetColumn("Float32"),
		FieldPowertrainTransmissionTravelledDistance: MustGetColumn("Float32"),
		FieldPowertrainTransmissionType: MustGetColumn("String"),
		FieldPowertrainType: MustGetColumn("String"),
		FieldRoofLoad: MustGetColumn("Int32"),
		FieldServiceDistanceToService: MustGetColumn("Float32"),
		FieldServiceIsServiceDue: MustGetColumn("UInt8"),
		FieldServiceTimeToService: MustGetColumn("Int32"),
		FieldSpeed: MustGetColumn("Float32"),
		FieldStartTime: MustGetColumn("String"),
		FieldSubject: MustGetColumn("String"),
		FieldTimestamp: MustGetColumn("String"),
		FieldTrailerIsConnected: MustGetColumn("UInt8"),
		FieldTraveledDistance: MustGetColumn("Float32"),
		FieldTraveledDistanceSinceStart: MustGetColumn("Float32"),
		FieldTripDuration: MustGetColumn("Float32"),
		FieldTripMeterReading: MustGetColumn("Float32"),
		FieldTurningDiameter: MustGetColumn("UInt32"),
		FieldVehicleIdentificationAcrissCode: MustGetColumn("String"),
		FieldVehicleIdentificationBodyType: MustGetColumn("String"),
		FieldVehicleIdentificationBrand: MustGetColumn("String"),
		FieldVehicleIdentificationDateVehicleFirstRegistered: MustGetColumn("String"),
		FieldVehicleIdentificationKnownVehicleDamages: MustGetColumn("String"),
		FieldVehicleIdentificationLicensePlate: MustGetColumn("String"),
		FieldVehicleIdentificationMeetsEmissionStandard: MustGetColumn("String"),
		FieldVehicleIdentificationModel: MustGetColumn("String"),
		FieldVehicleIdentificationOptionalExtras: MustGetColumn("Array(String)"),
		FieldVehicleIdentificationOptionalExtrassize0: MustGetColumn("UInt64"),
		FieldVehicleIdentificationProductionDate: MustGetColumn("String"),
		FieldVehicleIdentificationPurchaseDate: MustGetColumn("String"),
		FieldVehicleIdentificationVIN: MustGetColumn("String"),
		FieldVehicleIdentificationVehicleConfiguration: MustGetColumn("String"),
		FieldVehicleIdentificationVehicleExteriorColor: MustGetColumn("String"),
		FieldVehicleIdentificationVehicleInteriorColor: MustGetColumn("String"),
		FieldVehicleIdentificationVehicleInteriorType: MustGetColumn("String"),
		FieldVehicleIdentificationVehicleModelDate: MustGetColumn("String"),
		FieldVehicleIdentificationVehicleSeatingCapacity: MustGetColumn("UInt32"),
		FieldVehicleIdentificationVehicleSpecialUsage: MustGetColumn("String"),
		FieldVehicleIdentificationWMI: MustGetColumn("String"),
		FieldVehicleIdentificationYear: MustGetColumn("UInt32"),
		FieldVersionVSSLabel: MustGetColumn("String"),
		FieldVersionVSSMajor: MustGetColumn("UInt32"),
		FieldVersionVSSMinor: MustGetColumn("UInt32"),
		FieldVersionVSSPatch: MustGetColumn("UInt32"),
		FieldWidth: MustGetColumn("UInt32"),
		FieldWidthExcludingMirrors: MustGetColumn("UInt32"),
		FieldWidthFoldedMirrors: MustGetColumn("UInt32"),
		FieldWidthIncludingMirrors: MustGetColumn("UInt32"),
	}
)
