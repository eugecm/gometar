package parser

import (
	"github.com/eugecm/gometar/metar"
	"github.com/eugecm/gometar/qnh"
	qnhparser "github.com/eugecm/gometar/qnh/parser"
	"github.com/eugecm/gometar/sky"
	skyparser "github.com/eugecm/gometar/sky/parser"
	"github.com/eugecm/gometar/temperature"
	temperatureparser "github.com/eugecm/gometar/temperature/parser"
	"github.com/eugecm/gometar/time"
	timeparser "github.com/eugecm/gometar/time/parser"
	"github.com/eugecm/gometar/visibility"
	visibilityparser "github.com/eugecm/gometar/visibility/parser"
	"github.com/eugecm/gometar/visibility/runway"
	runwayparser "github.com/eugecm/gometar/visibility/runway/parser"
	"github.com/eugecm/gometar/weather"
	weatherparser "github.com/eugecm/gometar/weather/parser"
	"github.com/eugecm/gometar/wind"
	windparser "github.com/eugecm/gometar/wind/parser"
	"strings"
)

type Parser struct {
	QnhParser         qnh.Parser
	SkyParser         sky.Parser
	TemperatureParser temperature.Parser
	TimeParser        time.Parser
	VisibilityParser  visibility.Parser
	RunwayParser      runway.Parser
	WeatherParser     weather.Parser
	WindParser        wind.Parser
}

func New() metar.Parser {
	return &Parser{
		QnhParser:         qnhparser.New(),
		SkyParser:         skyparser.New(),
		TemperatureParser: temperatureparser.New(),
		TimeParser:        timeparser.New(),
		VisibilityParser:  visibilityparser.New(),
		RunwayParser:      runwayparser.New(),
		WeatherParser:     weatherparser.New(),
		WindParser:        windparser.New(),
	}
}

func (p *Parser) Parse(input string) (metar.Report, error) {
	r := metar.Report{}
	tokens := strings.Split(input, " ")
	curToken := 0

	r.Station = tokens[curToken]
	curToken++

	t, err := p.TimeParser.Parse(tokens[curToken])
	if err != nil {
		return metar.Report{}, err
	}
	curToken++
	r.DateTime = t

	r.Auto = tokens[curToken] == "AUTO"
	if r.Auto {
		curToken++
	}

	windString := tokens[curToken]
	curToken++

	if len(tokens[curToken]) == 7 && tokens[curToken][3] == 'V' {
		windString = windString + " "
		windString = windString + tokens[curToken]
		curToken++
	}
	w, err := p.WindParser.Parse(windString)
	if err != nil {
		return metar.Report{}, err
	}
	r.Wind = w

	if tokens[curToken] == "CAVOK" {
		r.Cavok = true
		curToken++

		temp, err := p.TemperatureParser.Parse(tokens[curToken])
		if err != nil {
			return metar.Report{}, err
		}
		r.Temperature = temp
		curToken++

		pres, err := p.QnhParser.Parse(tokens[curToken])
		if err != nil {
			return metar.Report{}, err
		}
		r.Qnh = pres
		return r, nil
	}

	v, err := p.VisibilityParser.Parse(tokens[curToken])
	if err != nil {
		return metar.Report{}, err
	}
	r.Visibility = v
	curToken++

	if strings.Contains(tokens[curToken], "/") { // RVR
		rvr, err := p.RunwayParser.Parse(tokens[curToken])
		if err != nil {
			return metar.Report{}, err
		}
		r.RunwayVisualRange = rvr
		curToken++
	}

	var wxStrings []string
	for len(tokens[curToken]) != 6 && tokens[curToken] != "NCD" && tokens[curToken][0:2] != "VV" && tokens[curToken] != "SKC" && tokens[curToken] != "CLR" {
		wxStrings = append(wxStrings, tokens[curToken])
		curToken++
	}

	if len(wxStrings) > 0 {
		wx, err := p.WeatherParser.Parse(strings.Join(wxStrings, " "))
		if err != nil {
			return metar.Report{}, err
		}

		r.Weather = wx
	}

	var skyStrings []string
	for !strings.Contains(tokens[curToken], "/") {
		skyStrings = append(skyStrings, tokens[curToken])
		curToken++
	}

	if len(skyStrings) > 0 {
		clouds, err := p.SkyParser.Parse(strings.Join(skyStrings, " "))
		if err != nil {
			return metar.Report{}, err
		}
		r.Clouds = clouds
	}

	temp, err := p.TemperatureParser.Parse(tokens[curToken])
	if err != nil {
		return metar.Report{}, err
	}
	curToken++
	r.Temperature = temp

	pres, err := p.QnhParser.Parse(tokens[curToken])
	if err != nil {
		return metar.Report{}, err
	}
	r.Qnh = pres

	return r, nil
}
