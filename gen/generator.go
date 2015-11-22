package gen

import (
	"path"
	"strings"

	"github.com/alvaroloes/sdkgen/parser"
	"github.com/juju/errors"
)

//go:generate go-bindata -o templates_bindata.go -debug=$DEBUG -pkg $GOPACKAGE ../templates/...
//go:generate stringer -type=Language

var (
	ErrLangNotSupported = errors.New("language not supported")
)

type Language int

const (
	Android Language = iota
	ObjC
	Swift
)

const (
	templateDir       = "./templates"
	templateExt       = ".tpl"
	modelTemplatePath = "model"
)

type Config struct {
	OutputDir     string
	ModelsRelPath string
	ApiName       string
	ApiPrefix     string
}

type specificGenerator interface {
	setTemplateDir(dir string)
	generate(config Config, api *parser.Api, modelsInfo []modelInfo) error
}

type Generator struct {
	gen        specificGenerator
	api        *parser.Api
	modelsInfo []modelInfo
	config     Config
}

func (g *Generator) Generate() error {
	return g.gen.generate(g.config, g.api, g.modelsInfo)
}

func (g *Generator) extractModelsInfo() error {
	return nil
}

func New(language Language, api *parser.Api, config Config) (Generator, error) {
	var gen specificGenerator
	var tplDir string

	switch language {
	case ObjC:
		gen = &ObjCGen{}
		tplDir = path.Join(templateDir, strings.ToLower(language.String()))
		//	case Android:
		//	case Swift:
	default:
		return Generator{}, errors.Annotate(ErrLangNotSupported, language.String())
	}
	gen.setTemplateDir(tplDir)

	generator := Generator{
		gen:    gen,
		api:    api,
		config: config,
	}

	return generator, nil
}
