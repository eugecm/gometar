/*
Package metar includes abstractions to represent meteorological reports.
*/
package metar

import "time"

type Report struct {
	// Station is the ICAO location indicator that this report describes.
	Station string

	// DateTime represents the date and time (UTC) of this report.
	DateTime time.Time

	// Auto indicates if the report contains only automated observations.
	Auto bool

	Wind WindGroup

	// Cavok indicates Cloud and Visbility OK. If set to true then Visibility
	// RunwayVisualRange, Weather and Cloud sections can be ignored.
	Cavok bool

	// Visibility describes the visibility conditions of the report.
	Visibility VisibilityGroup

	// RunwayVisualRange gives information about the visibility on the runway
	RunwayVisualRange RunwayVisualRangeGroup

	Weather WeatherGroup

	Clouds []CloudInformation

	Temperature TemperatureGroup
	Qnh         QnhGroup
	Remarks     RemarksGroup
}
