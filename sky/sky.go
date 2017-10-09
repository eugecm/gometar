package sky

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

type CloudInformation struct {
	Amount CloudAmount
	// Height is how high the cloud is from the airfield (in feet)
	Height int
	Type   CloudType
}
