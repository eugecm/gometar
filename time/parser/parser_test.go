package parser

import (
	"testing"

	"time"
)

func TestParser(t *testing.T) {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	cases := []struct {
		input    string
		expected time.Time
	}{
		{"240830Z", time.Date(year, month, 24, 8, 30, 0, 0, time.UTC)},
		{"251900Z", time.Date(year, month, 25, 19, 0, 0, 0, time.UTC)},
		{"010856Z", time.Date(year, month, 1, 8, 56, 0, 0, time.UTC)},
		{"122355Z", time.Date(year, month, 12, 23, 55, 0, 0, time.UTC)},
	}

	parser := New()
	for _, c := range cases {
		_t, err := parser.Parse(c.input)
		if err != nil {
			t.Errorf("couldn't parse time %v: %v", c.input, err)
			t.Fail()
			continue
		}

		if _t != c.expected {
			t.Errorf("expected time to be %v. Got %v instead", c.expected, _t)
			t.Fail()
			continue
		}
	}
}
