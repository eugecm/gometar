package visibility

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
	WeatherDescriptorNone WeatherDescriptor = iota
	WeatherDescriptorShallow
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
	WeatherPrecipitationNone WeatherPrecipitation = iota
	WeatherPrecipitationDrizzle
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
	WeatherObscurationNone WeatherObscuration = iota
	WeatherObscurationMist
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
	WeatherOtherPhenNone WeatherOtherPhen = iota
	WeatherOtherPhenWellDevelopedDustSandWhirls
	WeatherOtherPhenSqualls
	WeatherOtherPhenFunnelCloudTornadoWaterspout
	WeatherOtherPhenSandstorm
	WeatherOtherPhenDuststorm
)

type CloudAmount int

const (
	CloudAmountFew CloudAmount = iota
	CloudAmountScattered
	CloudAmountBroken
	CloudAmountOvercast
	CloudAmountNilSignificant
	CloudAmountNilDetected
)

type CloudType int

const (
	CloudTypeNone CloudType = iota
	CloudTypeToweringCumulus
	CloudTypeCumulonimbusOrashowerThunderstorm
	CloudTypeAltocumulusCastellanus
)

// Visibility describes the visibility conditions of the report.
type VisibilityGroup struct {
	Distance   int
	Unit       VisibilityUnit
	Modifier   VisibilityModifier
	ToDistance int
	ToModifier VisibilityModifier
	Trend      VisibilityTrend
}

type RunwayVisualRangeGroup struct {
	Runway     string
	Unit       VisibilityUnit
	Visibility float32
	// Modifier is used for expressing bounds within a distance
	// measurement. Examples: 3000 meters OR MORE, 3/4 Statute miles OR
	// LESS etc.
	Modifier VisibilityModifier
}

type WeatherGroup struct {
	Descriptor    WeatherDescriptor
	Precipitation WeatherPrecipitation
	Obscuration   WeatherObscuration
	Other         WeatherOtherPhen
	Vecinity      bool
}

type CloudInformation struct {
	Amount CloudAmount
	// Height is how high the cloud is from the airfield (in feet)
	Height int
	Type   CloudType
}
