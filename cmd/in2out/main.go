package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/key773251/in2out/pkg/parsers"
)

func main() {
	inputFile := flag.String("i", "", "Input file path (Required)")
	outputFile := flag.String("o", "", "Output file path (Required)")
	flag.Parse()

	if *inputFile == "" || *outputFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("Input File: ", *inputFile)
	fmt.Println("Output File: ", *outputFile)
	fmt.Println("-------------------------------")

	parser := parsers.Parser{}
	parser.Parse(*inputFile)
}
