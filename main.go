package main

import (
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"github.com/alvaroloes/sdkgen/gen"
	"github.com/alvaroloes/sdkgen/parser"
	"github.com/juju/errors"
)

func main() {
	log.SetLevel(log.DebugLevel)
	// This will be extracted from command line flags
	config := gen.Config{
		ApiName:       "Test",
		ApiPrefix:     "TT",
		ModelsRelPath: "Models",
		OutputDir:     "./",
	}

	specBytes, err := ioutil.ReadFile("./testFiles/api.spec")
	if err != nil {
		log.Fatal(errors.Annotate(err, "when reading API spec file"))
	}

	api, err := parser.NewApi(specBytes)
	if err != nil {
		log.Fatal(errors.ErrorStack(err))
	}

	gen, err := gen.New(gen.ObjC, api, config)
	if err != nil {
		log.Fatal(errors.ErrorStack(err))
	}

	if err := gen.Generate(); err != nil {
		log.Fatal(errors.ErrorStack(err))
	}

	//TODO:
	// - JSON arrays of arrays may not be properly handled

}
