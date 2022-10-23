package parsers

import (
	"encoding/json"
)

type JsonParser struct {
}

func (jp JsonParser) Parse(fileContents []byte, data any) any {
	json.Unmarshal(fileContents, &data)

	return data
}

func newJsonParser() Parser {
	return &JsonParser{}
}
