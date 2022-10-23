package parsers

import (
	"fmt"

	"golang.org/x/exp/maps"
)

// Returns map of currently supported file extensions for parsing
func getSupportedParsers() map[string]func() Parser {
	return map[string]func() Parser{
		"json":    newJsonParser,
		"jsonnet": newJsonnetParser,
		"yaml":    newYamlParser,
	}
}

func GetParser(fileType string, extVars map[string]string) (Parser, error) {
	parserMap := getSupportedParsers()

	if _, ok := parserMap[fileType]; ok {
		var parser Parser
		// Special case for jsonnet because of passed in extVars.
		// Kinda gross. Too lazy to find a better way
		if fileType == "jsonnet" && extVars != nil {
			parser = newJsonnetParserExt(extVars)

		} else {
			// Get the parser using the function defined in parserMap
			parser = parserMap[fileType]()
		}

		return parser, nil
	}

	return nil, fmt.Errorf("input file type '%s' not currently supported. Supported file types are: %s", fileType, maps.Keys(parserMap))
}
