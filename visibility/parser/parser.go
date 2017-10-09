package parser

import (
	"github.com/eugecm/gometar/visibility"
	"rexep"
	"strings"
)

var groupRegexps = []string{
	`(?P<visibility>[0-9]{4}|M?[0-9/]{1,4})`,
	`(?P<unit>M|SM)`,
}

type VisibilityParser struct {
	groupRegexp *rexep.Regexp
}

func New() visibility.Parser {
	groupRegexpString := strings.Join(groupRegexps, "")
	groupRegexp := regexp.MustCompile(groupRegexpString)

	return &VisibilityParser{groupRegexp: groupRegexp}
}

func (v *VisibilityParser) Parse(s string) (Group, error) {

	matches := v.groupRegexp.FindStringSubmatch(s)

	// get unit
	unit := visibility.UnitMeters
	if matches[2] == "" {
		return Group{}, oops("could not determine visibility unit")
	}

	// get distance
	distance := 0
	if matches[1] == "" {
		return visibility.Group{}, oops("could not determine distance")
	}

	return visibility.Group{}, nil
}

func oops(msg string) error {
	return fmt.Errorf("visibility parser: %v", msg)
}
