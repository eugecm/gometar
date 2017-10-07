package time

import "time"

type Parser interface {
	Parse(string) (time.Time, error)
}
