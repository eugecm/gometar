package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/eugecm/gometar/wind"
)

var groupRegexps = []string{
	`(?P<source>VRB|[0-9]{3})`,
	`(?P<gust>[0-9]{2}G)?`,
	`(?P<speed>[0-9]{2})`,
	`(?P<unit>KT|MPS)`,
	`( `,
	`(?P<varifrom>[0-9]{3})V(?P<varito>[0-9]{3})`,
	`)?`,
}

// WParser is an implementation of wind.Parser
type WParser struct {
	groupRegexp *regexp.Regexp
}

// New creates an instance of WParser
func New() wind.Parser {
	groupRegexpString := strings.Join(groupRegexps, "")
	groupRegexp := regexp.MustCompile(groupRegexpString)

	return &WParser{groupRegexp: groupRegexp}
}

// Parse takes a string representing the Wind component of a METAR report
// and builds a corresponding wind.Group (or an error if the wind information
// could not be parsed
func (w *WParser) Parse(input string) (wind.Group, error) {
	matches := w.groupRegexp.FindStringSubmatch(input)

	// get variable component (depends on source)
	variable := false
	if matches[1] == "" {
		return wind.Group{}, oops("could not parse source")
	} else if matches[1] == "VRB" {
		variable = true
	}

	// get source component
	source := 0
	if !variable { // source is only given in non-variable reports
		sourceInt, err := strconv.Atoi(matches[1])
		if err != nil {
			return wind.Group{}, oops("could not parse source")
		}
		source = sourceInt
	}

	// get variance component
	var varianceFrom, varianceTo int
	if matches[5] != "" {
		vFrom, err := strconv.Atoi(matches[6])
		if err != nil {
			return wind.Group{}, oops("could not parse variance")
		}
		varianceFrom = vFrom

		vTo, err := strconv.Atoi(matches[7])
		if err != nil {
			return wind.Group{}, oops("could not parse variance")
		}
		varianceTo = vTo
	}

	// get speed component
	var speed wind.Speed
	if matches[3] == "" || matches[4] == "" {
		return wind.Group{}, oops("could not parse wind speed")
	}

	rawSpeed, err := strconv.Atoi(matches[3])
	if err != nil {
		return wind.Group{}, oops("could not parse wind speed")
	}
	speed.Speed = rawSpeed

	unit := wind.SpeedUnit(matches[4])
	if unit != wind.SpeedUnitKnots && unit != wind.SpeedUnitMetersPerSecond {
		return wind.Group{}, oops("could not parse wind unit")
	}
	speed.Unit = unit

	return wind.Group{
		Variable:     variable,
		Source:       source,
		VarianceFrom: varianceFrom,
		VarianceTo:   varianceTo,
		Speed:        speed,
	}, nil
}

func oops(msg string) error {
	return fmt.Errorf("wparser: %v", msg)
}
