package parsers

import (
	"encoding/json"
	"log"

	"github.com/google/go-jsonnet"
)

type JsonnetParser struct {
	extVars map[string]string
}

func (jp JsonnetParser) Parse(fileContents []byte, data any) any {

	vm := jsonnet.MakeVM()

	// Bind jsonnet external variables
	for key, value := range jp.extVars {
		vm.ExtVar(key, value)
	}

	// Evaluate jsonnet and convert to json
	jsonStr, err := vm.EvaluateAnonymousSnippet("", string(fileContents))
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(jsonStr), &data)

	return data
}

func newJsonnetParser() Parser {
	return &JsonnetParser{}
}

// Extra function to create JsonnetParser and initialize extVars
func newJsonnetParserExt(extVars map[string]string) Parser {
	return &JsonnetParser{
		extVars: extVars,
	}
}
