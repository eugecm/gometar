package temperature

type Group struct {
	Temperature int
	DewPoint    int
}

type Parser interface {
	Parse(input string) (Group, error)
}
