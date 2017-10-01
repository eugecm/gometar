package parser

import (
	"testing"
)

func TestWindVariable(t *testing.T) {
	var cases = []struct {
		input    string
		expected bool
	}{
		{"00000KT", false},
		{"VRB01MPS", true},
		{"20005KT", false},
		{"05004KT", false},
	}

	parser := New()

	for _, c := range cases {
		group, err := parser.Parse(c.input)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		if group.Variable != c.expected {
			t.Errorf("expected variable to be %v, got %v", c.expected, group.Variable)
			t.Fail()
		}
	}
}

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

func TestWindVariance(t *testing.T) {
	var cases = []struct {
		input    string
		expected []int
	}{
		{"20005KT 130V260", []int{130, 260}},
		{"13016KT 100V160", []int{100, 160}},
		{"17008KT", []int{0, 0}},
		{"11005MPS 080V140", []int{80, 140}},
	}

	parser := New()

	for _, c := range cases {
		group, err := parser.Parse(c.input)
		expectedFrom := c.expected[0]
		expectedTo := c.expected[1]

		if err != nil {
			t.Error(err)
			t.Fail()
		}
		if group.VarianceFrom != expectedFrom {
			t.Errorf("expected VarianceFrom to be %v, got %v", group.VarianceFrom, expectedFrom)
			t.Fail()
		}
		if group.VarianceTo != expectedTo {
			t.Errorf("expected VarianceTo to be %v, got %v", group.VarianceTo, expectedTo)
			t.Fail()
		}
	}
}

//TODO: test for error handling

func BenchmarkWindParsing(b *testing.B) {
	cases := []string{
		"29008KT",
		"19011KT",
		"32007KT",
		"09007KT",
		"27006KT",
		"VRB04KT",
		"36001KT",
		"07004KT",
		"10003MPS",
		"35004KT",
		"28010KT",
		"00000KT",
	}
	casesN := len(cases)
	parser := New()

	for i := 0; i < b.N; i++ {
		parser.Parse(cases[i%casesN])
	}
}
