package parser

import (
	"testing"

	"github.com/eugecm/gometar/visibility"
)

func TestParserDistance(t *testing.T) {
	cases := []struct {
		input            string
		expectedDistance string
		expectedUnit     visibility.Unit
	}{
		{"10SM", "10", visibility.UnitStatuteMiles},
		{"M1/4SM", "1/4", visibility.UnitStatuteMiles},
		{"9999", "9999", visibility.UnitMeters},
		{"10SM", "10", visibility.UnitStatuteMiles},
	}

	p := New()
	for _, c := range cases {
		group, err := p.Parse(c.input)

		if err != nil {
			t.Error(err)
			t.Fail()
		}

		if group.Distance != c.expectedDistance {
			t.Errorf("expected distance to be %v, got %v instead", c.expectedDistance, group.Distance)
			t.Fail()
		}

		if group.Unit != c.expectedUnit {
			t.Errorf("expected unit to be %v, got %v instead", c.expectedUnit, group.Unit)
			t.Fail()
		}

	}
}

func TestParserModifier(t *testing.T) {
	cases := []struct {
		input            string
		expectedModifier visibility.Modifier
	}{
		{"10SM", visibility.ModifierExactly},
		{"M1/4SM", visibility.ModifierOrLess},
		{"9999", visibility.ModifierOrMore},
		{"15SM", visibility.ModifierOrMore},
	}

	p := New()
	for _, c := range cases {

		group, err := p.Parse(c.input)

		if err != nil {
			t.Error(err)
			t.Fail()
		}

		if group.Modifier != c.expectedModifier {
			t.Errorf("expected modifier for %v to be %v, got %v instead", c.input, c.expectedModifier, group.Modifier)
			t.Fail()
		}

	}

}
