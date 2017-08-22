/*
Package metar includes abstractions to help parse METAR strings.
 */
package metar

// Field is an abstraction that describes a meteorological feature of a METAR string
type Field interface {
	// Name is
	Name() string
	HumanString() string
	String() string
}

// FieldParser parses a METAR token string and returns a Field (if token is valid)
type FieldParser interface {
	GetField(field string) (Field, error)
}
