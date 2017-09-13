/*
Package metar includes abstractions to help parse METAR strings.
*/
package metar

import "time"

const (
	// TypeMetar indicates that the report is of type METAR
	TypeMetar = iota
	// TypeSpeci specifies that the report is of type SPECI
	TypeSpeci
)

const (
	// UnitStatuteMiles indicates distance is measured in SM (Only used in USA)
	UnitStatuteMiles = iota
	// UnitMeters indicates distance is measured in Meters
	UnitMeters
)

const (
	WeatherPatches = iota
	WeatherBlowing
	WeatherDrifting
	WeatherFreezing
	WeatherShallow
	WeatherPartial
	WeatherShowers
	WeatherThunderstorm
)

const (
	WeatherIntensityLight = iota
	WeatherIntensityModerate
	WeatherIntensityHeavy
)

// Type is either METAR os SPECI.
type Type int8

// Station is the ICAO location indicator that this report describes.
type Station string

// DateTime represents the date and time (UTC) of this report.
type DateTime time.Time

// Auto indicates if the report contains only automated observations.
type Auto bool

// Wind describes the wind conditions of the report.
type Wind struct {
	Direction struct {
		// Variable indicates that the direction cannot be determined.
		Variable bool
		// Source of the wind in degrees from true north.
		Source int
	}
	Variance struct {
		From int
		To   int
	}
	// Speed is the mean value for speed (in knots) observed in the sampling period.
	Speed int
	// Gust is the maximum speed measured in the sampling period.
	Gust int
}

// DistanceUnit indicates the unit of measurement used to represent visibility distance
type DistanceUnit string

// Visibility describes the visibility conditions of the report.
type Visibility struct {
	Cavok           bool
	Distance        int
	RunwayThreshold int
	Weather
	Cloud
}
type Temperature interface {
	Temperature() int
	DewPoint() int
}
type Pressure int
type Supplementary interface {
	RecentWeather() string
	WindSheer() struct {
		Runway int
	}
}
type Remarks string
