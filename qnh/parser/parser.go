package parser

import (
	"fmt"
	"regexp"

	"github.com/eugecm/gometar/qnh"
)

var qnhRegexp = `(?P<inhg>A[0-9]{4})?(?P<hpa>Q[0-9]{4})?`

type Parser struct {
	qnhRegexp *regexp.Regexp
}

func New() qnh.Parser {
	return &Parser{qnhRegexp: regexp.MustCompile(qnhRegexp)}
}

func (p *Parser) Parse(input string) (qnh.Group, error) {
	matches := p.qnhRegexp.FindStringSubmatch(input)
	if matches[1] == "" && matches[2] == "" {
		return qnh.Group{}, fmt.Errorf("qnh parser: could not extract qnh from %v", input)
	}

	var unit qnh.PressureUnit
	var pressure string
	if matches[1] != "" { // is inches of mercury
		unit = qnh.PressureUnitInchesOfMercury
		pressure = matches[1][1:len(matches[1])]
	} else { // is hectopascals
		unit = qnh.PressureUnitHectoPascals
		pressure = matches[2][1:len(matches[2])]
	}

	return qnh.Group{
		Pressure: pressure,
		Unit:     unit,
	}, nil
}
