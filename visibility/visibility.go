package visibility

// Unit is the measurement unit used for expressing visibility
// distance
type Unit int

const (
	// UnitUnknown indicates that the unit is unknown or not supported
	UnitUnknown Unit = iota
	// UnitStatuteMiles indicates distance is measured in SM (Only used in USA)
	UnitStatuteMiles
	// UnitMeters indicates distance is measured in Meters
	UnitMeters
	// UnitFeet indicates distance is measued in Feet
	UnitFeet
)

// Modifier indicates the visibility range of the report
type Modifier int

const (
	// ModifierExactly indicates measured distance is accurate
	ModifierExactly Modifier = iota
	// ModifierOrLess indicates measured distance could be lower
	ModifierOrLess
	// ModifierOrMore indicates measured distance could be higher
	ModifierOrMore
)

// Trend indicates the trend of the visibility (up, down, nil)
type Trend int

const (
	// TrendNotProvided indicates trend was not specified
	TrendNotProvided = iota
	// TrendUp indicates distance is increasing
	TrendUp
	// TrendDown indicates distance is decreasing
	TrendDown
	// TrendNil indicates distance is not changing
	TrendNil
)

// Group describes the visibility conditions of the report.
type Group struct {
	Distance string
	Unit     Unit
	Modifier Modifier
}

// Parser is an abstraction of an object that can parse Visibility strings
type Parser interface {
	Parse(s string) (Group, error)
}
