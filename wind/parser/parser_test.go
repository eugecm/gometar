package parser

import (
	"testing"
)

func TestWindSource(t *testing.T) {
	var cases = []struct {
		input    string
		expected int
	}{
		{"19020G26KT", 190},
		{"13009KT", 130},
		{"09004KT", 90},
		{"14002MPS 100V180", 140},
	}

	parser := New()

	for _, c := range cases {
		group, err := parser.Parse(c.input)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		if group.Source != c.expected {
			t.Errorf("expected source to be %v, got %v", c.expected, group.Source)
			t.Fail()
		}
	}
}
