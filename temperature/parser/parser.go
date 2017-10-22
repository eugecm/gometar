package parser

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/eugecm/gometar/temperature"
)

var tempRegexp = `(?P<temperature>M?[0-9]{2})/(?P<dewpoint>M?[0-9]{2})`

type Parser struct {
	tempRegexp *regexp.Regexp
}

func New() temperature.Parser {
	return &Parser{tempRegexp: regexp.MustCompile(tempRegexp)}
}

func stringToInt(input string) (int, error) {

	if input[0] == 'M' {
		n, err := strconv.Atoi(input[1:len(input)])
		if err != nil {
			return 0, err
		}
		return -n, nil
	}

	n, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return n, nil

}

func (p *Parser) Parse(input string) (temperature.Group, error) {
	matches := p.tempRegexp.FindStringSubmatch(input)
	if len(matches) != 3 {
		return temperature.Group{}, fmt.Errorf("temp parser: could not extract temp/dewpoint from %v", input)
	}

	tempString, dewString := matches[1], matches[2]
	temp, err := stringToInt(tempString)
	if err != nil {
		return temperature.Group{}, fmt.Errorf("temp parser: could not extract temp: %v", err)
	}
	dewpoint, err := stringToInt(dewString)
	if err != nil {
		return temperature.Group{}, fmt.Errorf("temp parser: could not extract dewpoint: %v", err)
	}

	return temperature.Group{
		Temperature: temp,
		DewPoint:    dewpoint,
	}, nil
}
