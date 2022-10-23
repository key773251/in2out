package converters

type Converter interface {
	Convert(any) (any, error)
}
