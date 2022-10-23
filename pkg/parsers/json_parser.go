package parsers

import (
	"encoding/json"
	"fmt"
)

type JsonParser struct {
}

func (jp JsonParser) Parse(fileContents []byte, data any) any {
	json.Unmarshal(fileContents, &data)

	jp.marshalJson(data)
	return data
}

func newJsonParser() Parser {
	return &JsonParser{}
}

//Joke code to try to marshal to json
func (jp JsonParser) marshalJson(data any) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("\nJSON Data:", jsonData)
	}
}
