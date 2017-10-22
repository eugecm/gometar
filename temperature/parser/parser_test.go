package parser

import "fmt"
import "github.com/eugecm/gometar/temperature"
import "testing"

func TestTemperature(t *testing.T) {
	cases := []struct {
		input    string
		expected temperature.Group
	}{
		{"24/23", temperature.Group{
			Temperature: 24,
			DewPoint:    23,
		}},
		{"17/09", temperature.Group{
			Temperature: 17,
			DewPoint:    9,
		}},
		{"08/03", temperature.Group{
			Temperature: 8,
			DewPoint:    3,
		}},
		{"M00/M02", temperature.Group{
			Temperature: 0,
			DewPoint:    -2,
		}},
	}

	p := New()
	for _, c := range cases {
		group, err := p.Parse(c.input)
		if err != nil {
			t.Errorf("could not parse %v: %v", c.input, err)
			t.FailNow()
		}

		testName := fmt.Sprintf("temperature of %v is %v", c.input, c.expected.Temperature)
		t.Run(testName, func(t *testing.T) {
			if group.Temperature != c.expected.Temperature {
				t.Errorf("got %v", group.Temperature)
				t.FailNow()
			}
		})

		testName = fmt.Sprintf("dewpoint of %v is %v", c.input, c.expected.DewPoint)
		t.Run(testName, func(t *testing.T) {
			if group.DewPoint != c.expected.DewPoint {
				t.Errorf("got %v", group.DewPoint)
				t.FailNow()
			}
		})
	}
}
