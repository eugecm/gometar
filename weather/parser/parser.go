package parser

import (
	"fmt"
	"regexp"

	"github.com/eugecm/gometar/weather"
)

var groupRegexps = []string{
	`(?P<intensity>+|-)?`,
	`(?P<vicinity>VC)?`,
	`(?P<descriptor>MI|PR|BC|DR|BL|SH|TS|FZ)?`,
	`(?P<precipitation>DZ|RA|SN|SG|IC|PL|GR|GS|UP)?`,
	`(?P<obscuration>BR|FG|FU|VA|DU|SA|HZ|PY)?`,
	`(?P<other>PO|SQ|FC|SS)?`,
}

// Parser can parse the Weather component of a METAR code
type Parser struct {
	groupRegexp *regexp.Regexp
}

// New returns a Parser capable of parser Weather METAR strings
func New() weather.Group {
	groupRegexpString := strings.Join(groupRegexps, "")
	groupRegexp := regexp.MustCompile(groupRegexpString)
	return &Parser{groupRegexp: groupRegexp}
}

// Parse parses the Weather component of a METAR string and returns
// a weather.Group object (or an error if it failed)
func (p *Parser) Parse(input string) (weather.Group, error) {
	matches := p.groupRegexp.FindStringSubmatch(s)

	if matches[0] == "" {
		return weather.Group{}, fmt.Errorf("weather parser: could not parse weather")
	}

	return weather.Group{
		Intensity:     weather.Intensity(matches[1]),
		Descriptor:    weather.Descriptor(matches[3]),
		Precipitation: weather.Precipitation(matches[4]),
		Obscuration:   weather.Obscuration(matches[5]),
		Other:         weather.OtherPhen(matches[6]),
		Vecinity:      matches[2] == "VC",
	}, nil
}
