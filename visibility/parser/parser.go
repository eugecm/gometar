package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/eugecm/gometar/visibility"
)

var groupRegexps = []string{
	`(?P<visibility>M?[0-9 /]{1,4})`,
	`(?P<unit>SM)?`,
}

type VisibilityParser struct {
	groupRegexp *regexp.Regexp
}

func New() visibility.Parser {
	groupRegexpString := strings.Join(groupRegexps, "")
	groupRegexp := regexp.MustCompile(groupRegexpString)

	return &VisibilityParser{groupRegexp: groupRegexp}
}

func (v *VisibilityParser) Parse(s string) (visibility.Group, error) {

	matches := v.groupRegexp.FindStringSubmatch(s)

	// get unit
	unit := visibility.UnitMeters
	if matches[2] == "SM" {
		unit = visibility.UnitStatuteMiles
	}

	// get distance
	distance := ""
	if matches[1] == "" {
		return visibility.Group{}, oops("could not determine distance")
	} else if strings.HasPrefix(matches[1], "M") {
		distance = matches[1][1:]
	} else {
		distance = matches[1]
	}

	modifier := visibility.VisibilityModifierExactly
	if matches[1] == "9999" || matches[1] == "15" {
		modifier = visibility.VisibilityModifierOrMore
	} else if strings.HasPrefix(matches[1], "M") {
		modifier = visibility.VisibilityModifierOrLess
	}

	return visibility.Group{Distance: distance, Unit: unit, Modifier: modifier}, nil
}

func oops(msg string) error {
	return fmt.Errorf("visibility parser: %v", msg)
}
