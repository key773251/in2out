package converters

import (
	"encoding/json"
)

type JsonConverter struct {
}

// Converts passed in data to json data
func (jc JsonConverter) Convert(data any) (any, error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")

	return jsonData, err
}

func newJsonConverter() Converter {
	return &JsonConverter{}
}
