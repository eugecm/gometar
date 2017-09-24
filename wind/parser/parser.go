package parser

import "github.com/eugecm/gometar/wind"

// WParser is an implementation of WindParser
type WParser struct {
}

// New creates an instance of WParser
func New() wind.WindParser {
	return &WParser{}
}

// Parse takes a string representing the Wind component of a METAR report
// and builds a corresponding wind.Group (or an error if the wind information
// could not be parsed
func (w *WParser) Parse(input string) (wind.Group, error) {
	return wind.Group{}, nil
}
