package parsers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type JsonParser struct {
}

func (p JsonParser) Parse(file string) error {
	data := map[string]interface{}{}
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Failed to ", err)
		return err
	}
	json.Unmarshal(contents, &data)

	fmt.Println(data)

	return nil
}

func newJsonParser() Parser {
	return &JsonParser{}
}
