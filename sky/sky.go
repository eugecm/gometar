package sky

type CloudAmount string

const (
	CloudAmountFew            CloudAmount = "FEW"
	CloudAmountScattered      CloudAmount = "SCT"
	CloudAmountBroken         CloudAmount = "BKN"
	CloudAmountOvercast       CloudAmount = "OVC"
	CloudAmountNilSignificant CloudAmount = "NSC"
	CloudAmountSkyClear       CloudAmount = "SKC"
	CloudAmountClear          CloudAmount = "CLR"
	CloudAmountCannotBeSeen   CloudAmount = "VV"
	CloudAmountNotDetected    CloudAmount = "NCD"
)

type CloudType string

const (
	CloudTypeNone            CloudType = ""
	CloudTypeToweringCumulus CloudType = "TCU"
	CloudTypeCumulonimbus    CloudType = "CB"
)

type CloudInformation struct {
	Height string
	Amount CloudAmount
	Type   CloudType
}

type Parser interface {
	Parse(input string) ([]CloudInformation, error)
}
