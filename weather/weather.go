package weather

type Intensity string

const (
	IntensityModerate Intensity = ""
	IntensityLight              = "-"
	IntensityHeavy              = "+"
)

type Descriptor string

const (
	DescriptorNone         Descriptor = ""
	DescriptorShallow                 = "MI"
	DescriptorPartial                 = "PR"
	DescriptorPatches                 = "BC"
	DescriptorLowDrifting             = "DR"
	DescriptorBlowing                 = "BL"
	DescriptorShowers                 = "SH"
	DescriptorThunderstorm            = "TS"
	DescriptorFreezing                = "FZ"
)

type Precipitation string

const (
	PrecipitationNone                      Precipitation = ""
	PrecipitationDrizzle                                 = "DZ"
	PrecipitationRain                                    = "RA"
	PrecipitationSnow                                    = "SN"
	PrecipitationSnowGrains                              = "SG"
	PrecipitationIceCrystals                             = "IC"
	PrecipitationIcePellets                              = "PL"
	PrecipitationHail                                    = "GR"
	PrecipitationSmallHailandOrSnowPellets               = "GS"
	PrecipitationUnknownPrecipitation                    = "UP"
)

type Obscuration string

const (
	ObscurationNone           Obscuration = ""
	ObscurationMist                       = "BR"
	ObscurationFog                        = "FG"
	ObscurationSmoke                      = "FU"
	ObscurationVolcanicAsh                = "VA"
	ObscurationWidespreadDust             = "DU"
	ObscurationSand                       = "SA"
	ObscurationHaze                       = "HZ"
	ObscurationSpray                      = "PY"
)

type OtherPhen string

const (
	OtherPhenNone                         OtherPhen = ""
	OtherPhenWellDevelopedDustSandWhirls            = "PO"
	OtherPhenSqualls                                = "SQ"
	OtherPhenFunnelCloudTornadoWaterspout           = "FC"
	OtherPhenSandstormOrDuststorm                   = "SS"
)

type Group struct {
	Descriptor    Descriptor
	Precipitation Precipitation
	Obscuration   Obscuration
	Other         OtherPhen
	Vecinity      bool
}

type Parser interface {
	Parse(input string) (Group, error)
}
