package metar

import "bytes"

// Metar is the struct representation of a METAR code, broken down into Fields
type Metar struct {
	Type Field
	Station Field
	DateTime Field
	ReportModifier Field
	Wind Field
	Visibility Field
	RunwayVisualRange Field
	Weather Field
	Sky Field
	Temperature Field
	Altimeter Field
	Remark Field
}

// String returns the raw (coded) representation of a METAR string
func (m *Metar) String() string {
}
