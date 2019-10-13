package parser

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/eugecm/gometar/metar"
	"github.com/eugecm/gometar/qnh"
	"github.com/eugecm/gometar/sky"
	"github.com/eugecm/gometar/temperature"
	"github.com/eugecm/gometar/visibility"
	"github.com/eugecm/gometar/visibility/runway"
	"github.com/eugecm/gometar/weather"
	"github.com/eugecm/gometar/wind"
)

func TestParser(t *testing.T) {
	now := time.Now()
	cases := []struct {
		input    string
		expected metar.Report
	}{
		{"BKPR 191800Z 00000KT CAVOK 13/06 Q1019 NOSIG", metar.Report{
			Station:  "BKPR",
			DateTime: time.Date(now.Year(), now.Month(), 19, 18, 0, 0, 0, time.UTC),
			Wind: wind.Group{
				Speed: wind.Speed{
					Speed: 0,
					Unit:  wind.SpeedUnitKnots,
				},
			},
			Cavok: true,
			Temperature: temperature.Group{
				Temperature: 13,
				DewPoint:    6,
			},
			Qnh: qnh.Group{
				Pressure: "1019",
				Unit:     qnh.PressureUnitHectoPascals,
			},
		}},
		{"CYVR 191813Z 15011KT 15SM -RA SCT012 BKN053 BKN100 BKN150 11/10 A2948 RMK SC3SC2AC1AC1 VIS NE-E 6 SLP984 DENSITY ALT 100FT", metar.Report{
			Station:  "CYVR",
			DateTime: time.Date(now.Year(), now.Month(), 19, 18, 13, 0, 0, time.UTC),
			Wind: wind.Group{
				Source: 150,
				Speed: wind.Speed{
					Speed: 11,
					Unit:  wind.SpeedUnitKnots,
				},
			},
			Visibility: visibility.Group{
				Distance: "15",
				Unit:     visibility.UnitStatuteMiles,
				Modifier: visibility.ModifierOrMore,
			},
			Weather: weather.Group{
				Intensity: weather.IntensityLight,
				Phenomena: []weather.Phenomenon{weather.PhenomenonRain},
			},
			Clouds: []sky.CloudInformation{
				{
					Height: "012",
					Amount: sky.CloudAmountScattered,
				},
				{
					Height: "053",
					Amount: sky.CloudAmountBroken,
				},
				{
					Height: "100",
					Amount: sky.CloudAmountBroken,
				},
				{
					Height: "150",
					Amount: sky.CloudAmountBroken,
				},
			},
			Temperature: temperature.Group{
				Temperature: 11,
				DewPoint:    10,
			},
			Qnh: qnh.Group{
				Pressure: "2948",
				Unit:     qnh.PressureUnitInchesOfMercury,
			},
		}},
		{"KHIB 191753Z AUTO 22007KT 190V250 10SM CLR 16/M02 A2996 RMK AO2 SLP151 T01611022 10161 20017 58005", metar.Report{
			Station:  "KHIB",
			DateTime: time.Date(now.Year(), now.Month(), 19, 17, 53, 0, 0, time.UTC),
			Auto:     true,
			Wind: wind.Group{
				Variable:     false,
				Source:       220,
				VarianceFrom: 190,
				VarianceTo:   250,
				Speed: wind.Speed{
					Speed: 7,
					Unit:  wind.SpeedUnitKnots,
				},
			},
			Visibility: visibility.Group{
				Distance: "10",
				Unit:     visibility.UnitStatuteMiles,
			},
			Clouds: []sky.CloudInformation{
				{
					Amount: sky.CloudAmountClear,
				},
			},
			Temperature: temperature.Group{
				Temperature: 16,
				DewPoint:    -2,
			},
			Qnh: qnh.Group{
				Pressure: "2996",
				Unit:     qnh.PressureUnitInchesOfMercury,
			},
		}},
		{"EGNM 191750Z 13011KT 1400 R14/P1500 RA BR SCT001 BKN002 13/13 Q0997", metar.Report{
			Station:  "EGNM",
			DateTime: time.Date(now.Year(), now.Month(), 19, 17, 50, 0, 0, time.UTC),
			Wind: wind.Group{
				Source: 130,
				Speed: wind.Speed{
					Speed: 11,
					Unit:  wind.SpeedUnitKnots,
				},
			},
			Visibility: visibility.Group{
				Distance: "1400",
				Unit:     visibility.UnitMeters,
			},
			RunwayVisualRange: runway.Group{
				Runway: "14",
				Visibility: visibility.Group{
					Distance: "1500",
					Unit:     visibility.UnitMeters,
					Modifier: visibility.ModifierOrMore,
				},
			},
			Weather: weather.Group{
				Phenomena: []weather.Phenomenon{
					weather.PhenomenonRain,
					weather.PhenomenonMist,
				},
			},
			Clouds: []sky.CloudInformation{
				{
					Height: "001",
					Amount: sky.CloudAmountScattered,
				},
				{
					Height: "002",
					Amount: sky.CloudAmountBroken,
				},
			},
			Temperature: temperature.Group{
				Temperature: 13,
				DewPoint:    13,
			},
			Qnh: qnh.Group{
				Pressure: "0997",
				Unit:     qnh.PressureUnitHectoPascals,
			},
		}},
		{"EHFD 191750Z AUTO 30016KT 9999 FEW026/// SCT033/// 06/M01 Q1001W///H///", metar.Report{
			Station:  "EHFD",
			DateTime: time.Date(now.Year(), now.Month(), 19, 17, 50, 0, 0, time.UTC),
			Auto: true,
			Wind: wind.Group{
				Source: 300,
				Speed: wind.Speed{
					Speed: 16,
					Unit:  wind.SpeedUnitKnots,
				},
			},
			Visibility: visibility.Group{
				Distance: "9999",
				Unit:     visibility.UnitMeters,
				Modifier: visibility.ModifierOrMore,
			},
			Clouds: []sky.CloudInformation{
				{
					Height: "026",
					Amount: sky.CloudAmountFew,
				},
				{
					Height: "033",
					Amount: sky.CloudAmountScattered,
				},
			},
			Temperature: temperature.Group{
				Temperature: 6,
				DewPoint:   -1,
			},
			Qnh: qnh.Group{
				Pressure: "1001",
				Unit:     qnh.PressureUnitHectoPascals,
			},
		}},
		{"EHJR 111825Z AUTO 29015KT 270V330 //// // ///////// 07/03 Q1022 RE// W07/H28", metar.Report{
			Station:  "EHJR",
			DateTime: time.Date(now.Year(), now.Month(), 11, 18, 25, 0, 0, time.UTC),
			Auto: true,
			Wind: wind.Group{
				Source: 290,
				VarianceFrom: 270,
				VarianceTo:   330,
				Speed: wind.Speed{
					Speed: 15,
					Unit:  wind.SpeedUnitKnots,
				},
			},
			Temperature: temperature.Group{
				Temperature: 7,
				DewPoint:    3,
			},
			Qnh: qnh.Group{
				Pressure: "1022",
				Unit:     qnh.PressureUnitHectoPascals,
			},
		}},
		{"EHKD 111925Z AUTO 30010KT 270V330 9999 NCD 06/03 Q1022 BLU 29015G25KT 9999 SCT025", metar.Report{
			Station:  "EHKD",
			DateTime: time.Date(now.Year(), now.Month(), 11, 19, 25, 0, 0, time.UTC),
			Auto: true,
			Wind: wind.Group{
				Source: 300,
				VarianceFrom: 270,
				VarianceTo:   330,
				Speed: wind.Speed{
					Speed: 10,
					Unit:  wind.SpeedUnitKnots,
				},
			},
			Visibility: visibility.Group{
				Distance: "9999",
				Unit:     visibility.UnitMeters,
				Modifier: visibility.ModifierOrMore,
			},
			Clouds: []sky.CloudInformation{
				{
					Amount: sky.CloudAmountNotDetected,
				},
			},
			Temperature: temperature.Group{
				Temperature: 6,
				DewPoint:    3,
			},
			Qnh: qnh.Group{
				Pressure: "1022",
				Unit:     qnh.PressureUnitHectoPascals,
			},
		}},
		{"EHFS 111925Z AUTO ///10KT 9999 NSC 08/03 Q//// W///H///", metar.Report{
			Station:  "EHFS",
			DateTime: time.Date(now.Year(), now.Month(), 11, 19, 25, 0, 0, time.UTC),
			Auto: true,
			Wind: wind.Group{
				Speed: wind.Speed{
					Speed: 10,
					Unit:  wind.SpeedUnitKnots,
				},
			},
			Visibility: visibility.Group{
				Distance: "9999",
				Unit:     visibility.UnitMeters,
				Modifier: visibility.ModifierOrMore,
			},
			Clouds: []sky.CloudInformation{
				{
					Amount: sky.CloudAmountNilSignificant,
				},
			},
			Temperature: temperature.Group{
				Temperature: 8,
				DewPoint:    3,
			},
		}},
	}

	p := New()

	for _, c := range cases {

		report, err := p.Parse(c.input)
		if err != nil {
			t.Errorf("could not parse %v: %v", c.input, err)
			t.FailNow()
		}

		testName := fmt.Sprintf("Station of %v is %v", c.input, c.expected.Station)
		t.Run(testName, func(t *testing.T) {
			if !reflect.DeepEqual(report.Station, c.expected.Station) {
				t.Errorf("expected %v but got %v", c.expected.Station, report.Station)
				t.FailNow()
			}
		})

		testName = fmt.Sprintf("Auto of %v is %v", c.input, c.expected.Auto)
		t.Run(testName, func(t *testing.T) {
			if !reflect.DeepEqual(report.Auto, c.expected.Auto) {
				t.Errorf("expected %v but got %v", c.expected.Auto, report.Auto)
				t.FailNow()
			}
		})

		testName = fmt.Sprintf("DateTime of %v is %v", c.input, c.expected.DateTime)
		t.Run(testName, func(t *testing.T) {
			if !reflect.DeepEqual(report.DateTime, c.expected.DateTime) {
				t.Errorf("expected %v but got %v", c.expected.DateTime, report.DateTime)
				t.FailNow()
			}
		})

		testName = fmt.Sprintf("Wind of %v is %v", c.input, c.expected.Wind)
		t.Run(testName, func(t *testing.T) {
			if !reflect.DeepEqual(report.Wind, c.expected.Wind) {
				t.Errorf("expected %v but got %v", c.expected.Wind, report.Wind)
				t.FailNow()
			}
		})

		testName = fmt.Sprintf("Cavok of %v is %v", c.input, c.expected.Cavok)
		t.Run(testName, func(t *testing.T) {
			if !reflect.DeepEqual(report.Cavok, c.expected.Cavok) {
				t.Errorf("expected %v but got %v", c.expected.Cavok, report.Cavok)
				t.FailNow()
			}
		})

		testName = fmt.Sprintf("Visibility of %v is %v", c.input, c.expected.Visibility)
		t.Run(testName, func(t *testing.T) {
			if !reflect.DeepEqual(report.Visibility, c.expected.Visibility) {
				t.Errorf("expected %v but got %v", c.expected.Visibility, report.Visibility)
				t.FailNow()
			}
		})

		testName = fmt.Sprintf("Weather of %v is %v", c.input, c.expected.Weather)
		t.Run(testName, func(t *testing.T) {
			if !reflect.DeepEqual(report.Weather, c.expected.Weather) {
				t.Errorf("expected %v but got %v", c.expected.Weather, report.Weather)
				t.FailNow()
			}
		})

		testName = fmt.Sprintf("RunwayVisualRange of %v is %v", c.input, c.expected.RunwayVisualRange)
		t.Run(testName, func(t *testing.T) {
			if !reflect.DeepEqual(report.RunwayVisualRange, c.expected.RunwayVisualRange) {
				t.Errorf("expected %v but got %v", c.expected.RunwayVisualRange, report.RunwayVisualRange)
				t.FailNow()
			}
		})

		testName = fmt.Sprintf("Clouds of %v is %v", c.input, c.expected.Clouds)
		t.Run(testName, func(t *testing.T) {
			if !reflect.DeepEqual(report.Clouds, c.expected.Clouds) {
				t.Errorf("expected %v but got %v", c.expected.Clouds, report.Clouds)
				t.FailNow()
			}
		})

		testName = fmt.Sprintf("Temperature of %v is %v", c.input, c.expected.Temperature)
		t.Run(testName, func(t *testing.T) {
			if !reflect.DeepEqual(report.Temperature, c.expected.Temperature) {
				t.Errorf("expected %v but got %v", c.expected.Temperature, report.Temperature)
				t.FailNow()
			}
		})

		testName = fmt.Sprintf("Qnh of %v is %v", c.input, c.expected.Qnh)
		t.Run(testName, func(t *testing.T) {
			if !reflect.DeepEqual(report.Qnh, c.expected.Qnh) {
				t.Errorf("expected %v but got %v", c.expected.Qnh, report.Qnh)
				t.FailNow()
			}
		})

	}
}
