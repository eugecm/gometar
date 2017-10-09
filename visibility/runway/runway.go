package runway

type RunwayVisualRangeGroup struct {
	Runway     string
	Unit       VisibilityUnit
	Visibility float32
	// Modifier is used for expressing bounds within a distance
	// measurement. Examples: 3000 meters OR MORE, 3/4 Statute miles OR
	// LESS etc.
	Modifier VisibilityModifier
}

type Parser interface {
	Parse(s string) (RunwayVisualRangeGroup, error)
}
