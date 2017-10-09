package visibility

// Unit is the measurement unit used for expressing visibility
// distance
type Unit int

const (
	// UnitStatuteMiles indicates distance is measured in SM (Only used in USA)
	UnitStatuteMiles Unit = iota
	// UnitMeters indicates distance is measured in Meters
	UnitMeters
)

// Modifier indicates the visibility range of the report
type Modifier int

const (
	// ModifierOrLess indicates measured distance could be lower
	ModifierOrLess Modifier = iota
	// ModifierExactly indicates measured distance is accurate
	ModifierExactly
	// ModifierOrMore indicates measured distance could be higher
	ModifierOrMore
)

// Trend indicates the trend of the visibility (up, down, nil)
type Trend int

const (
	// TrendUp indicates distance is increasing
	TrendUp Trend = iota
	// TrendDown indicates distance is decreasing
	TrendDown
	// TrendNil indicates distance is not changing
	TrendNil
	// TrendNotProvided indicates trend was not specified
	TrendNotProvided
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
