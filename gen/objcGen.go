package gen

import (
	"fmt"

	"path"

	"github.com/alvaroloes/sdkgen/parser"
)

type ObjCGen struct {
	tplDir string
}

func (gen *ObjCGen) setTemplateDir(tplDir string) {
	gen.tplDir = tplDir
}

func (gen *ObjCGen) Generate(api *parser.Api, config Config) error {
	modelTplDir := path.Join(gen.tplDir, modelTplPath)
	fmt.Println(modelTplDir)

	return nil
}
