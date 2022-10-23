package parsers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Parser struct {
}

func (p Parser) Parse(file string) error {
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
