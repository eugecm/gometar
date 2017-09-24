package parser

import "github.com/eugecm/gometar/wind"

type WParser struct {
}

func New() wind.WindParser {
	return &WParser{}
}

func (w *WParser) Parse(input string) (wind.Group, error) {
	return wind.Group{}, nil
}
