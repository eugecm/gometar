package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/eugecm/gometar/weather"
)

var validWxCodes = map[string]bool{
	"MI": true, "PR": true, "BC": true, "DR": true, "BL": true, "SH": true,
	"TS": true, "FZ": true, "DZ": true, "RA": true, "SN": true, "SG": true,
	"IC": true, "PL": true, "GR": true, "GS": true, "UP": true, "BR": true,
	"FG": true, "FU": true, "VA": true, "DU": true, "SA": true, "HZ": true,
	"PY": true, "PO": true, "SQ": true, "FC": true, "SS": true,
}

var groupRegexps = []string{
	`(?P<intensity>\+|-)?`,
	`(?P<vicinity>VC)?`,
	`(?P<phenomena>([A-Z]{2}\s?)+)`,
}

// Parser can parse the Weather component of a METAR code
type Parser struct {
	groupRegexp *regexp.Regexp
}

// New returns a Parser capable of parser Weather METAR strings
func New() weather.Parser {
	groupRegexpString := strings.Join(groupRegexps, "")
	groupRegexp := regexp.MustCompile(groupRegexpString)

	return &Parser{groupRegexp: groupRegexp}
}

// Parse parses the Weather component of a METAR string and returns
// a weather.Group object (or an error if it failed)
func (p *Parser) Parse(input string) (weather.Group, error) {
	matches := p.groupRegexp.FindStringSubmatch(input)

	if matches[0] == "" {
		return weather.Group{}, fmt.Errorf("weather parser: could not parse weather")
	}

	var phenomena []weather.Phenomenon
	wxString := strings.Replace(matches[3], " ", "", -1)
	for i := 2; i <= len(wxString); i = i + 2 {
		code := wxString[i-2 : i]
		_, ok := validWxCodes[code]
		if !ok {
			return weather.Group{}, fmt.Errorf("weather parser: invalid weather code %v", code)
		}

		phenomena = append(phenomena, weather.Phenomenon(code))
	}

	return weather.Group{
		Intensity: weather.Intensity(matches[1]),
		Phenomena: phenomena,
		Vicinity:  matches[2] == "VC",
	}, nil
}
