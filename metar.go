/*
Package metar includes abstractions to represent meteorological reports.
*/
package metar

import (
	"time"

	"github.com/eugecm/gometar/qnh"
	"github.com/eugecm/gometar/remarks"
	"github.com/eugecm/gometar/sky"
	"github.com/eugecm/gometar/temperature"
	"github.com/eugecm/gometar/visibility"
	"github.com/eugecm/gometar/visibility/runway"
	"github.com/eugecm/gometar/weather"
	"github.com/eugecm/gometar/wind"
)

// Report is a meteorological report at an airfield (METAR)
type Report struct {
	// Station is the ICAO location indicator that this report describes.
	Station string

	// DateTime represents the date and time (UTC) of this report.
	DateTime time.Time

	// Auto indicates if the report contains only automated observations.
	Auto bool

	// Wind describes wind conditions such as speed and direction.
	Wind wind.Group

	// Cavok indicates Cloud and Visbility OK. If set to true then Visibility
	// RunwayVisualRange, Weather and Cloud sections can be ignored.
	Cavok bool

	// Visibility describes the visibility in the sky at the airfield.
	Visibility visibility.Group

	// RunwayVisualRange gives information about the visibility on the runway
	RunwayVisualRange runway.Group

	// Weather describes the weather conditions at (or near) the airfield.
	Weather weather.Group

	// Clouds describes the clouds in the sky at different altitudes above
	// the airfield.
	Clouds []sky.CloudInformation

	// Temperature indicates temperature and dew point information.
	Temperature temperature.Group

	// Qnh represents the atmospheric pressure adjusted to sea level.
	Qnh qnh.Group

	// Remarks contains significant information not covered by the rest of the
	// METAR.
	Remarks string
}

type Parser interface {
	Parse(input string) (Report, error)
}
