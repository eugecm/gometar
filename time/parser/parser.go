package parser

import (
	"time"

	metarTime "github.com/eugecm/gometar/time"
)

// New returns a Parser capable of parsing METAR datetime groups
func New() metarTime.Parser {
	return &TParser{}
}

// TParser is an implementation of gometar.time.Parser
type TParser struct {
}

// Parse parses a METAR datetime string and returns a time.Time
func (t *TParser) Parse(s string) (time.Time, error) {
	return time.Parse("021504Z200601", s+time.Now().Format("200601"))
}
