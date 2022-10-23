package converters

type Converter interface {
	Convert(any) ([]byte, error)
}
