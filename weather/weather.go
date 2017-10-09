package weather

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

type WeatherGroup struct {
	Descriptor    WeatherDescriptor
	Precipitation WeatherPrecipitation
	Obscuration   WeatherObscuration
	Other         WeatherOtherPhen
	Vecinity      bool
}
