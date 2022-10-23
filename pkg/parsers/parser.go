package parsers

type Parser interface {
	Parse([]byte, any) any
}
