package qnh

type PressureUnit int

const (
	PressureUnitHectoPascals PressureUnit = iota
	PressureUnitInchesOfMercury
)

type Group struct {
	Pressure int
	Unit     PressureUnit
}
