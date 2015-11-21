package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/alvaroloes/sdkgen/gen"
	"github.com/alvaroloes/sdkgen/parser"
	"github.com/juju/errors"
)

func main() {
	specBytes, err := ioutil.ReadFile("./testFiles/api.spec")
	if err != nil {
		log.Fatalln(err)
	}
	api, err := parser.NewApi(specBytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, errors.ErrorStack(err))
		return
	}

	fmt.Println(api)

	gen, err := gen.New(gen.ObjC)

	fmt.Println(gen, err)
}
