package wind

// SpeedUnit is the type of unit used to measure the speed in METAR
type SpeedUnit string

const (
	// SpeedUnitKnots = Knots (KN)
	SpeedUnitKnots SpeedUnit = "KN"
	// SpeedUnitMetersPerSecond = M/s
	SpeedUnitMetersPerSecond = "MPS"
)

// WindGroup gives wind information for a METAR
type WindGroup struct {

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
	Speed struct {
		Speed int
		Unit  SpeedUnit
	}

	// Gust is the maximum speed measured in the sampling period.
	Gust int
}

type WindParser interface {
	Parse(string) (WindGroup, error)
}
