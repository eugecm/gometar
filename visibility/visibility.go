package visibility

// VisibilityUnit is the measurement unit used for expressing visibility
// distance
type VisibilityUnit int

const (
	// UnitStatuteMiles indicates distance is measured in SM (Only used in USA)
	UnitStatuteMiles VisibilityUnit = iota
	// UnitMeters indicates distance is measured in Meters
	UnitMeters
)

// VisibilityModifier indicates the visibility range of the report
type VisibilityModifier int

const (
	VisibilityModifierOrLess VisibilityModifier = iota
	VisibilityModifierExactly
	VisibilityModifierOrMore
)

// VisibilityTrend indicates the trend of the visibility
type VisibilityTrend int

const (
	VisibilityTrendUp VisibilityTrend = iota
	VisibilityTrendDown
	VisibilityTrendNil
	VisibilityTrendNotProvided
)

// Visibility describes the visibility conditions of the report.
type Group struct {
	Distance   int
	Unit       VisibilityUnit
	Modifier   VisibilityModifier
	ToDistance int
	ToModifier VisibilityModifier
	Trend      VisibilityTrend
}

type Parser interface {
	Parse(s string) (Group, error)
}
