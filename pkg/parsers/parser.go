package parsers

type Parser interface {
	Parse(file string) error
}
