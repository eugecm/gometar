package parser

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/eugecm/gometar/weather"
)

func TestWeatherParser(t *testing.T) {
	cases := []struct {
		input    string
		expected weather.Group
	}{
		{"-RADZ", weather.Group{
			Intensity: weather.IntensityLight,
			Phenomena: []weather.Phenomenon{
				weather.PhenomenonRain,
				weather.PhenomenonDrizzle,
			},
			Vicinity: false,
		}},
		{"-RA BR", weather.Group{
			Intensity: weather.IntensityLight,
			Phenomena: []weather.Phenomenon{
				weather.PhenomenonRain,
				weather.PhenomenonMist,
			},
			Vicinity: false,
		}},
		{"-RA", weather.Group{
			Intensity: weather.IntensityLight,
			Phenomena: []weather.Phenomenon{
				weather.PhenomenonRain,
			},
			Vicinity: false,
		}},
		{"+RA", weather.Group{
			Intensity: weather.IntensityHeavy,
			Phenomena: []weather.Phenomenon{
				weather.PhenomenonRain,
			},
			Vicinity: false,
		}},
		{"BR", weather.Group{
			Intensity: weather.IntensityModerate,
			Phenomena: []weather.Phenomenon{
				weather.PhenomenonMist,
			},
			Vicinity: false,
		}},
		{"SHRA", weather.Group{
			Intensity: weather.IntensityModerate,
			Phenomena: []weather.Phenomenon{
				weather.PhenomenonShowers,
				weather.PhenomenonRain,
			},
			Vicinity: false,
		}},
		{"DZRA", weather.Group{
			Intensity: weather.IntensityModerate,
			Phenomena: []weather.Phenomenon{
				weather.PhenomenonDrizzle,
				weather.PhenomenonRain,
			},
			Vicinity: false,
		}},
		{"VCFG", weather.Group{
			Intensity: weather.IntensityModerate,
			Phenomena: []weather.Phenomenon{
				weather.PhenomenonFog,
			},
			Vicinity: true,
		}},
		{"HZ", weather.Group{
			Intensity: weather.IntensityModerate,
			Phenomena: []weather.Phenomenon{
				weather.PhenomenonHaze,
			},
			Vicinity: false,
		}},
	}

	p := New()
	for _, c := range cases {
		group, err := p.Parse(c.input)

		if err != nil {
			t.Errorf("could not parse %v: %v", c.input, err)
			t.FailNow()
		}

		testName := fmt.Sprintf("intensity of %v is %v", c.input, c.expected.Intensity)
		t.Run(testName, func(t *testing.T) {
			if group.Intensity != c.expected.Intensity {
				t.Errorf("expected %v but got %v", c.expected.Intensity, group.Intensity)
				t.FailNow()
			}
		})

		testName = fmt.Sprintf("phenomena of %v is %v", c.input, c.expected.Phenomena)
		t.Run(testName, func(t *testing.T) {
			if !reflect.DeepEqual(group.Phenomena, c.expected.Phenomena) {
				t.Errorf("expected %v but got %v", c.expected.Phenomena, group.Phenomena)
				t.FailNow()
			}
		})

		testName = fmt.Sprintf("vicinity of %v is %v", c.input, c.expected.Vicinity)
		t.Run(testName, func(t *testing.T) {
			if group.Vicinity != c.expected.Vicinity {
				t.Errorf("expected %v but got %v", c.expected.Vicinity, group.Vicinity)
				t.FailNow()
			}
		})

	}
}
