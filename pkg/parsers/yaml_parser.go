package parsers

import (
	"gopkg.in/yaml.v3"
)

type YamlParser struct {
}

func (yp YamlParser) Parse(fileContents []byte, data any) any {
	yaml.Unmarshal(fileContents, &data)

	return data
}

func newYamlParser() Parser {
	return &YamlParser{}
}
