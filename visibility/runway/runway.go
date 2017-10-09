package runway

import "github.com/eugecm/gometar/visibility"

// Group represents the Runway Visual Range information part of the report
type Group struct {
	Runway     string
	Visibility visibility.Group
	Variable   visibility.Group
	Trend      visibility.Trend
}

// Parser is an abstraction of entities that can parse Runway Visual Range
// strings from a METAR
type Parser interface {
	Parse(s string) (Group, error)
}
