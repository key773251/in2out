package converters

import (
	"fmt"

	"golang.org/x/exp/maps"
)

// Returns map of currently supported file extensions for converting
func getSupportedConverters() map[string]func() Converter {
	return map[string]func() Converter{
		"json": newJsonConverter,
	}
}

func GetParser(fileType string) (Converter, error) {
	converterMap := getSupportedConverters()

	if _, ok := converterMap[fileType]; ok {
		// Get the converter using the function defined in converterMap
		parser := converterMap[fileType]()
		return parser, nil
	}

	return nil, fmt.Errorf("output file type '%s' not currently supported.\n Supported file types are: %s", fileType, maps.Keys(converterMap))
}
