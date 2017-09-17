package qnh

type PressureUnit int

const (
	PressureUnitHectoPascals PressureUnit = iota
	PressureUnitInchesOfMercury
)

type QnhGroup struct {
	Pressure int
	Unit     PressureUnit
}
