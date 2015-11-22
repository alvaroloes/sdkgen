package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/alvaroloes/sdkgen/gen"
	"github.com/alvaroloes/sdkgen/parser"
	"github.com/juju/errors"
)

var logger = log.New(os.Stderr, "", 0)

func main() {
	// This will be extracted from command line flags
	config := gen.Config{
		ApiName:       "Test",
		ApiPrefix:     "TT",
		ModelsRelPath: "Models",
		OutputDir:     "./",
	}

	specBytes, err := ioutil.ReadFile("./testFiles/api.spec")
	if err != nil {
		logger.Fatal(errors.Annotate(err, "when reading API spec file"))
	}

	api, err := parser.NewApi(specBytes)
	if err != nil {
		logger.Fatal(errors.ErrorStack(err))
	}

	gen, err := gen.New(gen.ObjC, api, config)
	if err != nil {
		logger.Fatal(errors.ErrorStack(err))
	}

	if err := gen.Generate(); err != nil {
		logger.Fatal(errors.ErrorStack(err))
	}
}
