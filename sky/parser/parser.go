package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/eugecm/gometar/sky"
)

var cloudAmountRegexp = `(?P<amount>FEW|SCT|BKN|OVC|NSC|SKC|VV)(?P<height>[0-9]{3})(?P<cloud>TCU|CB)?`

// Parser is an implementation of sky.Parser
type Parser struct {
	cloudAmountRegexp *regexp.Regexp
}

// New creates an instance of WParser
func New() sky.Parser {
	return &Parser{
		cloudAmountRegexp: regexp.MustCompile(cloudAmountRegexp),
	}
}

func (p *Parser) getCloudInfo(input string) (sky.CloudInformation, error) {
	if input == "NSC" {
		return sky.CloudInformation{
			Amount: sky.CloudAmountNilSignificant,
		}, nil
	}

	if input == "CLR" {
		return sky.CloudInformation{
			Amount: sky.CloudAmountClear,
		}, nil
	}

	if !p.cloudAmountRegexp.MatchString(input) {
		return sky.CloudInformation{}, fmt.Errorf("invalid string")
	}

	matches := p.cloudAmountRegexp.FindStringSubmatch(input)
	return sky.CloudInformation{
		Height: matches[2],
		Amount: sky.CloudAmount(matches[1]),
		Type:   sky.CloudType(matches[3]),
	}, nil
}

func (p *Parser) Parse(input string) ([]sky.CloudInformation, error) {
	var clouds []sky.CloudInformation

	for _, cloudString := range strings.Split(input, " ") {
		cloud, err := p.getCloudInfo(cloudString)
		if err != nil {
			return nil, fmt.Errorf("cloud parser: could not parse cloud information: %v", err)
		}
		clouds = append(clouds, cloud)
	}

	return clouds, nil

}
