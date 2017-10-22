package parser

import (
	"fmt"
	"testing"

	"github.com/eugecm/gometar/qnh"
)

func TestQnh(t *testing.T) {
	cases := []struct {
		input    string
		expected qnh.Group
	}{
		{"Q1008", qnh.Group{
			Pressure: "1008",
			Unit:     qnh.PressureUnitHectoPascals,
		}},
		{"Q0998", qnh.Group{
			Pressure: "0998",
			Unit:     qnh.PressureUnitHectoPascals,
		}},
		{"A3018", qnh.Group{
			Pressure: "3018",
			Unit:     qnh.PressureUnitInchesOfMercury,
		}},
		{"A2943", qnh.Group{
			Pressure: "2943",
			Unit:     qnh.PressureUnitInchesOfMercury,
		}},
	}

	p := New()
	for _, c := range cases {
		group, err := p.Parse(c.input)
		if err != nil {
			t.Errorf("could not parse %v: %v", c.input, err)
			t.FailNow()
		}

		testName := fmt.Sprintf("pressure of %v is %v", c.input, c.expected.Pressure)
		t.Run(testName, func(t *testing.T) {
			if group.Pressure != c.expected.Pressure {
				t.Errorf("got %v", group.Pressure)
				t.FailNow()
			}
		})

		testName = fmt.Sprintf("unit of %v is %v", c.input, c.expected.Unit)
		t.Run(testName, func(t *testing.T) {
			if group.Unit != c.expected.Unit {
				t.Errorf("got %v", group.Unit)
				t.FailNow()
			}
		})
	}
}
