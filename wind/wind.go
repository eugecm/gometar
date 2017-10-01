package wind

// SpeedUnit is the type of unit used to measure the speed in METAR
type SpeedUnit string

const (
	// SpeedUnitKnots = Knots (KT)
	SpeedUnitKnots SpeedUnit = "KT"
	// SpeedUnitMetersPerSecond = M/s
	SpeedUnitMetersPerSecond = "MPS"
)

// Speed represents wind speed value and unit
type Speed struct {
	// Speed is the velocity of the wind
	Speed int
	// Unit is the unit used for measuring the speed
	Unit SpeedUnit
}

// Group gives wind information for a METAR
type Group struct {

	// Variable indicates that the direction cannot be determined.
	Variable bool

	// Source of the wind in degrees from true north.
	Source int

	// VarianceFrom is the minimum observed wind direction represented
	// in degrees from true north. Only given if direction varies
	// substantially
	VarianceFrom int

	// VarianceFrom is the maximum observed wind direction represented
	// in degrees from true north. Only given if direction varies
	// substantially
	VarianceTo int

	// Speed is the mean value for speed observed in the sampling period.
	Speed Speed

	// Gust is the maximum speed measured in the sampling period.
	Gust int
}

// Parser is an interface for parsers that return wind.Group information
type Parser interface {
	Parse(string) (Group, error)
}
