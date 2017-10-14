package parser

import (
	"fmt"
	"testing"

	"github.com/eugecm/gometar/visibility"
	"github.com/eugecm/gometar/visibility/runway"
)

func TestParser(t *testing.T) {

	cases := []struct {
		input    string
		expected runway.Group
	}{
		{"R36/4000FT", runway.Group{
			Runway: "36",
			Visibility: visibility.Group{
				Distance: "4000",
				Unit:     visibility.UnitFeet,
				Modifier: visibility.ModifierExactly,
			},
		}},
		{"R01L/0600V1000FT", runway.Group{
			Runway: "01L",
			Visibility: visibility.Group{
				Distance: "0600",
				Unit:     visibility.UnitFeet,
				Modifier: visibility.ModifierExactly,
			},
			IsVariable: true,
			Variable: visibility.Group{
				Distance: "1000",
				Unit:     visibility.UnitFeet,
				Modifier: visibility.ModifierExactly,
			},
		}},
		{"R27/P6000FT", runway.Group{
			Runway: "27",
			Visibility: visibility.Group{
				Distance: "6000",
				Unit:     visibility.UnitFeet,
				Modifier: visibility.ModifierOrMore,
			},
		}},
		{"R23/0200N", runway.Group{
			Runway: "23",
			Visibility: visibility.Group{
				Distance: "0200",
				Unit:     visibility.UnitMeters,
				Modifier: visibility.ModifierExactly,
			},
			Trend: visibility.TrendNil,
		}},
		{"R04/P1500U", runway.Group{
			Runway: "04",
			Visibility: visibility.Group{
				Distance: "1500",
				Unit:     visibility.UnitMeters,
				Modifier: visibility.ModifierOrMore,
			},
			Trend: visibility.TrendUp,
		}},
		{"R36/0300V750D", runway.Group{
			Runway: "36",
			Visibility: visibility.Group{
				Distance: "0300",
				Unit:     visibility.UnitMeters,
				Modifier: visibility.ModifierExactly,
			},
			IsVariable: true,
			Variable: visibility.Group{
				Distance: "750",
				Unit:     visibility.UnitMeters,
				Modifier: visibility.ModifierExactly,
			},
			Trend: visibility.TrendDown,
		}},
	}

	parser := New()
	for _, c := range cases {
		grp, err := parser.Parse(c.input)
		if err != nil {
			t.Errorf("could not parse RVR string %#v: %#v", c.input, err)
			t.Fail()
			continue
		}

		t.Run(fmt.Sprint("runway of ", c.input), func(t *testing.T) {
			if grp.Runway != c.expected.Runway {
				t.Errorf("is not '%#v' (got %#v)", c.expected.Runway, grp.Runway)
				t.Fail()
			}
		})

		t.Run(fmt.Sprint("visibility of ", c.input), func(t *testing.T) {
			if grp.Visibility != c.expected.Visibility {
				t.Errorf("is not %#v (got %#v)", c.expected.Visibility, grp.Visibility)
				t.Fail()
			}
		})

		t.Run(fmt.Sprint("variability of ", c.input), func(t *testing.T) {
			if grp.IsVariable != c.expected.IsVariable {
				t.Errorf("is not %v (got %v)", c.expected.IsVariable, grp.IsVariable)
				t.Fail()
			}
			if grp.Variable != c.expected.Variable {
				t.Errorf("is not %#v (got %#v)", c.expected.Variable, grp.Variable)
				t.Fail()
			}
		})

		t.Run(fmt.Sprint("trend of ", c.input), func(t *testing.T) {
			if grp.Trend != c.expected.Trend {
				t.Errorf("is not %#v (got %#v)", c.expected.Trend, grp.Trend)
				t.Fail()
			}
		})
	}
}
