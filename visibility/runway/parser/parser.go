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

// RunwayParser parses RunwayVisualRange strings from METARs
type RunwayParser struct {
	groupRegexp *regexp.Regexp
}

// New returns a RunwayParser
func New() runway.Parser {
	groupRegexpString := strings.Join(groupRegexps, "")
	groupRegexp := regexp.MustCompile(groupRegexpString)

	return &RunwayParser{groupRegexp: groupRegexp}
}

func translateModifier(modifier string) visibility.Modifier {
	switch modifier {
	case "M":
		return visibility.ModifierOrLess
	case "P":
		return visibility.ModifierOrMore
	default:
		return visibility.ModifierExactly
	}
}

func translateTrend(trend string) visibility.Trend {
	switch trend {
	case "U":
		return visibility.TrendUp
	case "D":
		return visibility.TrendDown
	case "N":
		return visibility.TrendNil
	default:
		return visibility.TrendNotProvided
	}
}

// Parse builds a runway.Group from a Runway Visual Range METAR string
func (r *RunwayParser) Parse(s string) (runway.Group, error) {
	matches := r.groupRegexp.FindStringSubmatch(s)

	// get runway
	rwy := matches[1]
	if rwy == "" {
		return runway.Group{}, oops("could not parse runway threshold")
	}

	// get modifier
	modifier := translateModifier(matches[2])

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
	trend := translateTrend(matches[6])

	// build final runway object
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
