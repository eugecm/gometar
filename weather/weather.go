package weather

type Intensity string

const (
	IntensityModerate Intensity = ""
	IntensityLight              = "-"
	IntensityHeavy              = "+"
)

type Phenomenon string

const (
	PhenomenonShallow                      = "MI"
	PhenomenonPartial                      = "PR"
	PhenomenonPatches                      = "BC"
	PhenomenonLowDrifting                  = "DR"
	PhenomenonBlowing                      = "BL"
	PhenomenonShowers                      = "SH"
	PhenomenonThunderstorm                 = "TS"
	PhenomenonFreezing                     = "FZ"
	PhenomenonDrizzle                      = "DZ"
	PhenomenonRain                         = "RA"
	PhenomenonSnow                         = "SN"
	PhenomenonSnowGrains                   = "SG"
	PhenomenonIceCrystals                  = "IC"
	PhenomenonIcePellets                   = "PL"
	PhenomenonHail                         = "GR"
	PhenomenonSmallHailandOrSnowPellets    = "GS"
	PhenomenonUnknown                      = "UP"
	PhenomenonMist                         = "BR"
	PhenomenonFog                          = "FG"
	PhenomenonSmoke                        = "FU"
	PhenomenonVolcanicAsh                  = "VA"
	PhenomenonWidespreadDust               = "DU"
	PhenomenonSand                         = "SA"
	PhenomenonHaze                         = "HZ"
	PhenomenonSpray                        = "PY"
	PhenomenonWellDevelopedDustSandWhirls  = "PO"
	PhenomenonSqualls                      = "SQ"
	PhenomenonFunnelCloudTornadoWaterspout = "FC"
	PhenomenonSandstormOrDuststorm         = "SS"
)

type Group struct {
	Intensity Intensity
	Phenomena []Phenomenon
	Vicinity  bool
}

type Parser interface {
	Parse(input string) (Group, error)
}

// BUG(eugeniocanom@gmail.com): Weather groups should not be represented as a
// list of Phenomena as the codes can form groups depending on the meaning of
// each code.
