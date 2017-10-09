package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/eugecm/gometar/visibility"
)

var groupRegexps = []string{
	`M?(?P<visibility>[0-9 /]{1,4})`,
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
	}

	distance = matches[1]

	return visibility.Group{Distance: distance, Unit: unit}, nil
}

func oops(msg string) error {
	return fmt.Errorf("visibility parser: %v", msg)
}
