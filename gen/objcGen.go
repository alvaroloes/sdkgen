package gen

import (
	"fmt"
	"html/template"
	"path"

	"strings"

	"github.com/alvaroloes/sdkgen/parser"
	"github.com/juju/errors"
)

type ObjCGen struct {
	tplDir string
}

func (gen *ObjCGen) setTemplateDir(tplDir string) {
	gen.tplDir = tplDir
}

func (gen *ObjCGen) generate(config Config, api *parser.Api, modelsInfo []modelInfo) error {
	modelTplDir := path.Join(gen.tplDir, modelTemplatePath)

	modelTpls, err := template.ParseGlob(path.Join(modelTplDir, "*"+templateExt))
	if err != nil {
		return errors.Annotate(err, "when reading model templates at "+modelTplDir)
	}
	otherTpls, err := template.ParseGlob(path.Join(gen.tplDir, "*"+templateExt))
	if err != nil {
		return errors.Annotate(err, "when reading templates at "+gen.tplDir)
	}

	apiDir := path.Join(config.OutputDir, config.ApiName)
	modelsDir := path.Join(apiDir, config.ModelsRelPath)

	for _, tpl := range modelTpls.Templates() {
		modelFileName := "Hola"
		tplFileName := tpl.Name()
		from := strings.Index(tplFileName, ".")
		to := strings.LastIndex(tplFileName, ".")
		if from > 0 && to > 0 {
			modelFileName += tplFileName[from:to]
		}
		modelFilePath := path.Join(modelsDir, modelFileName)
		fmt.Println(modelFilePath)
	}
	for _, tpl := range otherTpls.Templates() {
		fmt.Println(tpl.Name())
	}

	return nil
}
