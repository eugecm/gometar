package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/eugecm/gometar/visibility"
	"github.com/eugecm/gometar/visibility/runway"
)

var groupRegexps = []string{
	`R(?P<runway>[0-9]{2}[LCR]?)/`,
	`(?P<modifier>M|P)?`,
	`(?P<visibility>[0-9]{4})`,
	`(FT)?`,
	`(?P<variable>V[MP]?[0-9]+)?`,
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

func (r *RunwayParser) Parse(s string) (runway.Group, error) {
	matches := r.groupRegexp.FindStringSubmatch(s)

	// get runway
	rwy := matches[1]
	if rwy == "" {
		return runway.Group{}, oops("could not parse runway threshold")
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
		fmt.Printf("%#v\n", matches)
		return runway.Group{}, oops("could not parse distance")
	}

	// get unit
	unit := visibility.UnitMeters
	if matches[4] == "FT" || matches[7] == "FT" {
		unit = visibility.UnitFeet
	}

	// get variable
	variableDistance := ""
	variableModifier := visibility.ModifierExactly
	variableUnit := visibility.UnitUnknown
	if matches[5] != "" {
		if matches[5][1] == 'M' {
			variableModifier = visibility.ModifierOrLess
			variableDistance = matches[5][2:]
			variableUnit = unit
		} else if matches[5][1] == 'P' {
			variableModifier = visibility.ModifierOrMore
			variableDistance = matches[5][2:]
			variableUnit = unit
		} else {
			variableDistance = matches[5][1:]
			variableUnit = unit
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

	g := runway.Group{
		Runway: rwy,
		Visibility: visibility.Group{
			Distance: distance,
			Unit:     unit,
			Modifier: modifier,
		},
		IsVariable: variableDistance != "",
		Variable: visibility.Group{
			Distance: variableDistance,
			Unit:     variableUnit,
			Modifier: variableModifier,
		},
		Trend: trend,
	}

	return g, nil

}

func oops(msg string) error {
	return fmt.Errorf("runway parser: %v", msg)
}
