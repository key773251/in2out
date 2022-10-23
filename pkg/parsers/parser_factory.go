package parsers

import (
	"fmt"

	"golang.org/x/exp/maps"
)

// Returns map of currently supported file extensions for parsing
func getSupportedParsers() map[string]func() Parser {
	return map[string]func() Parser{
		"json": newJsonParser,
		"yaml": newYamlParser,
	}
}

func GetParser(fileType string) (Parser, error) {
	parserMap := getSupportedParsers()

	if _, ok := parserMap[fileType]; ok {
		// Get the parser using the function defined in parserMap
		parser := parserMap[fileType]()
		return parser, nil
	}

	return nil, fmt.Errorf("input file type '%s' not currently supported.\n Supported file types are: %s", fileType, maps.Keys(parserMap))
}
