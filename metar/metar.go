/*
Package metar includes abstractions to help parse METAR strings.
 */
package metar

import "time"

const (
	TypeMetar = iota
	TypeSpeci
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
	// Direction (in degrees) where the wind is coming from.
	Direction() int
	// Mean speed. Given in knots.
	Speed() int
	Gust() int
	Variation() int
}
type Visibility interface {
	Cavok() bool
	Distance() int
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
