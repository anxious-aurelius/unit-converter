package main

import "errors"

type Unit int

const (
	millimeter Unit = iota
	centimeter
	meter
	kilometer
	inch
	foot
	yard
	mile
)

var unitToMeters = map[Unit]float64{
	millimeter: 0.001,
	centimeter: 0.01,
	meter:      1.0,
	kilometer:  1000.0,
	inch:       0.0254,
	foot:       0.3048,
	yard:       0.9144,
	mile:       1609.344,
}

var unitTypeConversion = map[string]Unit{
	"millimeter": millimeter,
	"centimeter": centimeter,
	"meter":      meter,
	"kilometer":  kilometer,
	"inch":       inch,
	"foot":       foot,
	"yard":       yard,
	"mile":       mile,
}

func convertLengthUnit(value float64, inputUnit Unit, targetUnit Unit) (float64, error) {
	//converting to meters(base unit)
	unit, ok := unitToMeters[inputUnit]
	if !ok {
		return 0, errors.New("invalid source unit provided")
	}
	tempMeter := value * unit
	unit, ok = unitToMeters[targetUnit]
	if !ok {
		return 0, errors.New("invalid target unit provided")
	}
	targetValue := tempMeter / unit
	return targetValue, nil
}

func stringToUnit(unit string) (Unit, error) {
	value, ok := unitTypeConversion[unit]
	if !ok {
		return 0, errors.New("the unit provided is not a valid unit")
	}
	return value, nil
}
