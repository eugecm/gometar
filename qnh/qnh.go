package qnh

type PressureUnit int

const (
	PressureUnitHectoPascals PressureUnit = iota
	PressureUnitInchesOfMercury
)

type Group struct {
	Pressure string
	Unit     PressureUnit
}

type Parser interface {
	Parse(input string) (Group, error)
}
