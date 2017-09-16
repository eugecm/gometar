/*
Package metar includes abstractions to help parse METAR strings.
*/
package metar

import "time"

// Type of METAR report
type MetarType int8

const (
	// TypeMetar indicates that the report is of type METAR
	TypeMetar MetarType = iota
	// TypeSpeci specifies that the report is of type SPECI
	TypeSpeci
)

type VisibilityUnit int

const (
	// UnitStatuteMiles indicates distance is measured in SM (Only used in USA)
	UnitStatuteMiles VisibilityUnit = iota
	// UnitMeters indicates distance is measured in Meters
	UnitMeters
)

type VisibilityModifier int

const (
	VisibilityModifierOrLess VisibilityModifier = iota
	VisibilityModifierExactly
	VisibilityModifierOrMore
)

type VisibilityTrend int

const (
	VisibilityTrendUp VisibilityTrend = iota
	VisibilityTrendDown
	VisibilityTrendNil
	VisibilityTrendNotProvided
)

type WeatherIntensity int

const (
	WeatherIntensityLight WeatherIntensity = iota
	WeatherIntensityModerate
	WeatherIntensityHeavy
)

type WeatherDescriptor int

const (
	WeatherDescriptorShallow WeatherDescriptor = iota
	WeatherDescriptorPartial
	WeatherDescriptorPatches
	WeatherDescriptorLowDrifting
	WeatherDescriptorBlowing
	WeatherDescriptorShowers
	WeatherDescriptorThunderstorm
	WeatherDescriptorFreezing
)

type WeatherPrecipitation int

const (
	WeatherPrecipitationDrizzle WeatherPrecipitation = iota
	WeatherPrecipitationRain
	WeatherPrecipitationSnow
	WeatherPrecipitationSnowGrains
	WeatherPrecipitationIceCrystals
	WeatherPrecipitationIcePellets
	WeatherPrecipitationHail
	WeatherPrecipitationSmallHailandOrSnowPellets
	WeatherPrecipitationUnknownPrecipitation
)

type WeatherObscuration int

const (
	WeatherObscurationMist WeatherObscuration = iota
	WeatherObscurationFog
	WeatherObscurationSmoke
	WeatherObscurationVolcanicAsh
	WeatherObscurationWidespreadDust
	WeatherObscurationSand
	WeatherObscurationHaze
	WeatherObscurationSpray
)

type WeatherOtherPhen int

const (
	WeatherOtherPhenWellDevelopedDustSandWhirls WeatherOtherPhen = iota
	WeatherOtherPhenSqualls
	WeatherOtherPhenFunnelCloudTornadoWaterspout
	WeatherOtherPhenSandstorm
	WeatherOtherPhenDuststorm
)

type Report struct {
	// Station is the ICAO location indicator that this report describes.
	Station string

	// DateTime represents the date and time (UTC) of this report.
	DateTime time.Time

	// Auto indicates if the report contains only automated observations.
	Auto bool

	// Wind information for the report
	Wind struct {

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

		// Speed is the mean value for speed (in knots) observed in the sampling
		// period.
		Speed int

		// Gust is the maximum speed measured in the sampling period.
		Gust int
	}

	// Cavok indicates Cloud and Visbility OK. If set to true then Visibility
	// RunwayVisualRange, Weather and Cloud sections can be ignored.
	Cavok bool

	// Visibility describes the visibility conditions of the report.
	Visibility struct {
		Distance   int
		Unit       int
		Modifier   VisibilityModifier
		ToDistance int
		ToModifier VisibilityModifier
		Trend      VisibilityTrend
	}

	RunwayVisualRange struct {
		Runway     string
		Visibility VisibilityUnit
		Modifier
	}

	Weather struct {
		Descriptor    WeatherDescriptor
		Precipitation WeatherPrecipitation
		Obscuration   WeatherObscuration
		Other         WeatherOtherPhen
		Vecinity      bool
	}

	Temperature struct {
		Temperature int
		DewPoint    int
	}

	Pressure int

	Supplementary struct {
		RecentWeather string
		WindSheer     int
	}

	Remarks string
}
