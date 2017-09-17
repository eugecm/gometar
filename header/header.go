package header

// Type of METAR report
type MetarType int8

const (
	// TypeMetar indicates that the report is of type METAR
	TypeMetar MetarType = iota
	// TypeSpeci specifies that the report is of type SPECI
	TypeSpeci
)
