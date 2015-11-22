package gen

import (
	"fmt"

	"path"

	"html/template"

	"os"

	"github.com/alvaroloes/sdkgen/parser"
)

type ObjCGen struct {
	tplDir string
}

func (gen *ObjCGen) setTemplateDir(tplDir string) {
	gen.tplDir = tplDir
}

func (gen *ObjCGen) Generate(api *parser.Api, config Config) error {
	modelTplDir := path.Join(gen.tplDir, modelTemplatePath)
	fmt.Println(modelTplDir)

	modelTpls, err := template.ParseGlob(path.Join(modelTplDir, "/*"+templateExt))

	fmt.Println(modelTpls, err)

	modelTpls.ExecuteTemplate(os.Stdout, "model.h.tpl", config)
	modelTpls.ExecuteTemplate(os.Stdout, "model.m.tpl", config)

	return nil
}
