package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/eugecm/gometar/sky"
)

func TestSkyParser(t *testing.T) {
	cases := []struct {
		input    string
		expected []sky.CloudInformation
	}{
		{"FEW012 BKN040 OVC140",
			[]sky.CloudInformation{
				sky.CloudInformation{
					Height: "012",
					Amount: sky.CloudAmountFew,
					Type:   sky.CloudTypeNone,
				},
				sky.CloudInformation{
					Height: "040",
					Amount: sky.CloudAmountBroken,
					Type:   sky.CloudTypeNone,
				},
				sky.CloudInformation{
					Height: "140",
					Amount: sky.CloudAmountOvercast,
					Type:   sky.CloudTypeNone,
				},
			},
		},
		{"SCT018TCU BKN200",
			[]sky.CloudInformation{
				sky.CloudInformation{
					Height: "018",
					Amount: sky.CloudAmountScattered,
					Type:   sky.CloudTypeToweringCumulus,
				},
				sky.CloudInformation{
					Height: "200",
					Amount: sky.CloudAmountBroken,
					Type:   sky.CloudTypeNone,
				},
			},
		},
		{"NSC",
			[]sky.CloudInformation{
				sky.CloudInformation{
					Amount: sky.CloudAmountNilSignificant,
				},
			},
		},
	}

	p := New()
	for _, c := range cases {
		clouds, err := p.Parse(c.input)
		if err != nil {
			t.Errorf("failed to parse '%v': %v", c.input, err)
			t.FailNow()
		}

		testName := fmt.Sprintf("number of clouds in %v: %v", c.input, len(c.expected))
		t.Run(testName, func(t *testing.T) {
			if len(clouds) != len(c.expected) {
				t.Errorf("amounts don't match: %#v", clouds)
				t.FailNow()
			}
		})

		cloudStrings := strings.Split(c.input, " ")
		for i := 0; i < len(c.expected); i++ {
			testName = fmt.Sprintf("height of %v is %v", cloudStrings[i], c.expected[i].Height)
			t.Run(testName, func(t *testing.T) {
				if clouds[i].Height != c.expected[i].Height {
					t.Errorf("height is %v", clouds[i].Height)
					t.FailNow()
				}
			})

			testName = fmt.Sprintf("amount of %v is %v", cloudStrings[i], c.expected[i].Amount)
			t.Run(testName, func(t *testing.T) {
				if clouds[i].Amount != c.expected[i].Amount {
					t.Errorf("amount is %v", clouds[i].Amount)
					t.FailNow()
				}
			})

			testName = fmt.Sprintf("type of %v is %v", cloudStrings[i], c.expected[i].Type)
			t.Run(testName, func(t *testing.T) {
				if clouds[i].Type != c.expected[i].Type {
					t.Errorf("type is %v", clouds[i].Type)
					t.FailNow()
				}
			})
		}
	}
}
