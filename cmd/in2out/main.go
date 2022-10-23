package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

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

	fmt.Println("in2out version:", version)

	fmt.Println("\nInput File:", *inputFile)
	fmt.Println("Input File Extension:", inputFileExt)
	fmt.Println("Output File:", *outputFile)

	parser, err := parsers.GetParser(inputFileExt)

	if err != nil {
		fmt.Println(err)
	} else {
		parser.Parse(*inputFile)
	}
}
