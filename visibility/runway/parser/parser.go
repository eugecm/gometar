package parser

import (
	"regexp"

	"github.com/eugecm/gometar/visibility/runway"
)

var groupRegexps = []string{
	`R(?P<runway>[0-9]{2}(L|C|R)?)/`,
	`(?P<modifier>M|P)?`,
	`(?P<visibility>[0-9]{4})`,
	`(FT)?`,
	`(?P<variable>V(M|P)?[0-9]{4})?`,
	`(?P<tendency>U|D|N)?`,
	`(FT)?`,
}

type RunwayParser struct {
	groupRegexp *regexp.Regexp
}

func New() runway.Parser {
	groupRegexpString := strings.Join(groupRegexps, "")
	groupRegexp := regexp.MustCompile(groupRegexpString)

	return &RunwayParser{groupRegexp: groupRegexp}
}

func (r *RunwayParser) Parse(runway.Group, error) {
	matches := v.groupRegexp.FindStringSubmatch(s)

	// get runway
	runway := matches[1]
	if runway == "" {
		return Group{}, oops("could not parse runway threshold")
	}

	// get modifier
	modifier := visibility.ModifierExactly
	if matches[2] == "M" {
		modifier = visibility.ModifierOrLess
	} else if matches[2] == "P" {
		modifier = visibility.ModifierOrMore
	}

	// get distance
	distance := matches[3]
	if distance == "" {
		return Group{}, oops("could not parse distance")
	}

	// get unit
	unit := visibility.UnitMeters
	if matches[4] == "FT" || matches[7] == "FT" {
		unit = visibility.UnitFeet
	}

	// get variable
	variableDistance = ""
	variableModifier = visibility.ModifierExactly
	if matches[5] != "" {
		if matches[5][1] == "M" {
			variableModifier = visibility.ModifierOrLess
			variableDistance = matches[5][2:]
		} else if matches[5][1] == "P" {
			variableModifier = visibility.ModifierOrMore
			variableDistance = matches[5][2:]
		} else {
			variableDistance = matches[5][1:]
		}
	}

	// get trend
	var trend visibility.Trend
	switch matches[6] {
	case "U":
		trend = visibility.TrendUp
	case "D":
		trend = visibility.TrendDown
	case "N":
		trend = visibility.TrendNil
	default:
		trend = visibility.TrendNotProvided
	}

	g := Group{
		Runway: runway,
		Visibility: visibility.Group{
			Distance: distance,
			Unit:     unit,
			Modifier: modifier,
		},
		Variable: visibility.Group{
			Distance: variableDistance,
			Unit:     unit,
			Modifier: variableModifier,
		},
		Trend: trend,
	}

	return g, nil

}

func oops(msg string) error {
	return fmt.Errorf("runway parser: %v", msg)
}
