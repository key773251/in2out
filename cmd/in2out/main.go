package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/key773251/in2out/pkg/converters"
	"github.com/key773251/in2out/pkg/parsers"
)

const version = "0.1.0"

func main() {
	inputFile := flag.String("i", "", "Input file path (Required)")
	outputFile := flag.String("o", "", "Output file path (Required)")
	flag.Parse()

	if *inputFile == "" || *outputFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	inputFileExt := strings.Split(*inputFile, ".")[1]
	outputFileExt := strings.Split(*outputFile, ".")[1]

	fmt.Println("in2out version:", version)

	fmt.Println("\nInput File:", *inputFile)
	fmt.Println("Output File:", *outputFile)

	// Try to read the file
	data := map[string]interface{}{}
	contents, err := ioutil.ReadFile(*inputFile)

	if err != nil {
		fmt.Println("\nFailed to", err)
		os.Exit(1)
	}

	// Get the parser from the factory and parse the file
	parser, err := parsers.GetParser(inputFileExt)
	var inputData any

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Printf("\nUsing %s file parser.", inputFileExt)
		inputData = parser.Parse(contents, &data)

		fmt.Println("\nInput Data:", inputData)
	}

	// Get the converter from the factory and convert the data
	converter, err := converters.GetConverter(outputFileExt)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Printf("\nUsing %s data converter.", outputFileExt)
		outputData, err := converter.Convert(inputData)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println("\nOutput Data:", string(outputData))
			ioutil.WriteFile(*outputFile, outputData, 0644)
		}
	}
}
