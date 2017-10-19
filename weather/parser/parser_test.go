package parser

import (
	"fmt"
	"github.com/eugecm/gometar/weather"
	"testing"
)

func TestWeatherParser(t *testing.T) {
	cases := []struct {
		input    string
		expected weather.Group
	}{
		{"-RADZ", weather.Group{
			Intensity:     weather.IntensityLight,
			Descriptor:    weather.DescriptorNone,
			Precipitation: weather.PrecipitationRain,
			Obscuration:   weather.ObscurationNone,
			Other:         weather.OtherPhenNone,
			Vicinity:      false,
		}},
		{"-RA BR", weather.Group{}},
		{"-RA", weather.Group{}},
		{"+RA", weather.Group{}},
		{"BR", weather.Group{}},
		{"SHRA", weather.Group{}},
		{"DZRA", weather.Group{}},
		{"VCFG", weather.Group{}},
		{"HZ", weather.Group{}},
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

		testName := fmt.Sprintf("descritor of %v is %v", c.input, c.expected.Descriptor)
		t.Run(testName, func(t *testing.T) {
			if group.Descriptor != c.expected.Descriptor {
				t.Errorf("expected %v but got %v", c.expected.Descriptor, group.Descriptor)
				t.FailNow()
			}
		})

		testName := fmt.Sprintf("precipitation of %v is %v", c.input, c.expected.Precipitation)
		t.Run(testName, func(t *testing.T) {
			if group.Precipitation != c.expected.Precipitation {
				t.Errorf("expected %v but got %v", c.expected.Precipitation, group.Precipitation)
				t.FailNow()
			}
		})

		testName := fmt.Sprintf("obscuration of %v is %v", c.input, c.expected.Obscuration)
		t.Run(testName, func(t *testing.T) {
			if group.Obscuration != c.expected.Obscuration {
				t.Errorf("expected %v but got %v", c.expected.Obscuration, group.Obscuration)
				t.FailNow()
			}
		})

		testName := fmt.Sprintf("other weather of %v is %v", c.input, c.expected.Other)
		t.Run(testName, func(t *testing.T) {
			if group.Other != c.expected.Other {
				t.Errorf("expected %v but got %v", c.expected.Other, group.Other)
				t.FailNow()
			}
		})

		testName := fmt.Sprintf("vicinity of %v is %v", c.input, c.expected.Vicinity)
		t.Run(testName, func(t *testing.T) {
			if group.Vicinity != c.expected.Vicinity {
				t.Errorf("expected %v but got %v", c.expected.Vicinity, group.Vicinity)
				t.FailNow()
			}
		})

	}
}
