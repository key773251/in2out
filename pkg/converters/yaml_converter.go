package converters

import "gopkg.in/yaml.v3"

type YamlConverter struct {
}

// Converts passed in data to json data
func (yc YamlConverter) Convert(data any) ([]byte, error) {
	jsonData, err := yaml.Marshal(data)

	return jsonData, err
}

func newYamlConverter() Converter {
	return &YamlConverter{}
}
