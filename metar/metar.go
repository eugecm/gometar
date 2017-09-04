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

// Type is either METAR os SPECI.
type Type int8

// Station is the ICAO location indicator that this report describes.
type Station string

// DateTime represents the date and time (UTC) of this report.
type DateTime time.Time

// Auto indicates if the report contains only automated observations.
type Auto bool

// Wind describes the wind conditions of the report.
type Wind interface {
	// Direction returns the direction of the wind in degrees from true north, or an error if
	// the direction is unknown.
	Direction() (int, error)
	// IsVariable determines whether the wind direction varies by more than 60 degrees.
	IsVariable() bool
	// Variance returns the starting and ending direction of the wind variance.
	Variance() (int, int)
	// Speed is the mean value of the speed during the sampling period (aprox 10 minutes).
	Speed() int
	// Gust returns the maximum wind speed measured during the sampling period. If the maximum
	// wind speed does not exceed the mean wind speed by more than 10 knots, error will be set
	// to true indicating that there is no gust.
	Gust() (int, error)
}

// DistanceUnit indicates the unit of measurement used to represent visibility distance
type DistanceUnit string

// Visibility describes the visibility conditions of the report.
type Visibility interface {
	// Cavok (Cloud And Visibility OK) indicates that the following conditions are observed
	// simultaneously:
	//   - Visibility of 10Km or more.
	//   - No clouds below 5000ft or highest minimum sector altitude. Whichever is greater.
	//   - No cumulonimbus or towering cumulus at any level.
	//   - No significant weather condition.
	Cavok() bool
	// Distance may be given in Statue
	Distance() int
	DistanceUnit()
	Runway() struct {
		Threshold int
	}
	Weather() string
	Cloud() string
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
