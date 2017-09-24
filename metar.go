/*
Package metar includes abstractions to represent meteorological reports.
*/
package metar

import (
	"time"

	"github.com/eugecm/wind"
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
	Visibility VisibilityGroup

	// RunwayVisualRange gives information about the visibility on the runway
	RunwayVisualRange RunwayVisualRangeGroup

	// Weather describes the weather conditions at (or near) the airfield.
	Weather WeatherGroup

	// Clouds describes the clouds in the sky at different altitudes above
	// the airfield.
	Clouds []CloudInformation

	// Temperature indicates temperature and dew point information.
	Temperature TemperatureGroup

	// Qnh represents the atmospheric pressure adjusted to sea level.
	Qnh QnhGroup

	// Remarks contains significant information not covered by the rest of the
	// METAR.
	Remarks RemarksGroup
}
