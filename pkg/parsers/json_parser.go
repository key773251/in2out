package parsers

import (
	"encoding/json"
)

type JsonParser struct {
}

func (p JsonParser) Parse(fileContents []byte, data any) error {
	json.Unmarshal(fileContents, &data)

	return nil
}

func newJsonParser() Parser {
	return &JsonParser{}
}
