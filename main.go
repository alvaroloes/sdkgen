package main
import (
	"io/ioutil"
	"log"
	"github.com/alvaroloes/sdkgen/parser"
	"fmt"
	"github.com/juju/errors"
	"os"
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
}
