package main
import (
	"io/ioutil"
	"log"
	"github.com/alvaroloes/sdkgen/parser"
)

func main() {
	specBytes, err := ioutil.ReadFile("./api.spec")
	if err != nil {
		log.Fatalln(err)
	}
	parser.NewApi(specBytes)
}
