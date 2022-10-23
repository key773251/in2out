package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/key773251/in2out/pkg/converters"
	"github.com/key773251/in2out/pkg/parsers"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const version = "0.1.0"

func parseExtVars(extVarsList string) map[string]string {

	extVarsMap := make(map[string]string)

	keyValueStrings := strings.Split(extVarsList, ",")

	for _, element := range keyValueStrings {
		keyValuePairs := strings.Split(element, "=")
		key := keyValuePairs[0]
		value := keyValuePairs[1]

		extVarsMap[key] = value
	}

	return extVarsMap
}

func main() {
	fmt.Print("in2out version:", version, "\n")

	// Setup logging
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	inputFile := flag.String("i", "", "Input file path (Required)")
	outputFile := flag.String("o", "", "Output file path (Required)")
	extVarsArg := flag.String("e", "", "External Variables for jsonnet substitution")
	debugFlag := flag.Bool("d", false, "Enable debugging output")
	flag.Parse()

	if *debugFlag {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	var extVars map[string]string
	if *extVarsArg != "" {
		extVars = parseExtVars(*extVarsArg)
		log.Debug().Msgf("External Variables: %s", extVars)
	}

	if *inputFile == "" || *outputFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	inputFileExt := strings.Split(*inputFile, ".")[1]
	outputFileExt := strings.Split(*outputFile, ".")[1]

	log.Info().Msgf("Input file: %s", *inputFile)

	// Try to read the file
	data := map[string]interface{}{}
	contents, err := ioutil.ReadFile(*inputFile)

	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to open input file:")
	}

	// Get the parser from the factory and parse the file
	parser, err := parsers.GetParser(inputFileExt, extVars)
	var inputData any

	if err != nil {
		log.Fatal().Err(err).Msgf("Error finding parser:")
		os.Exit(1)
	} else {
		log.Debug().Msgf("Using %s file parser.", inputFileExt)
		inputData = parser.Parse(contents, &data)

		log.Trace().Msgf("Input Data: %s", inputData)
	}

	// Get the converter from the factory and convert the data
	converter, err := converters.GetConverter(outputFileExt)

	if err != nil {
		log.Fatal().Err(err).Msgf("Error finding data converter:")
	} else {
		log.Debug().Msgf("Using %s data converter.", outputFileExt)
		outputData, err := converter.Convert(inputData)

		if err != nil {
			log.Fatal().Err(err).Msg("Error with data conversion.")
		} else {
			log.Trace().Msgf("Output Data:\n", string(outputData))
			log.Info().Msgf("Writing output file to: %s", *outputFile)
			ioutil.WriteFile(*outputFile, outputData, 0644)
		}
	}
}
