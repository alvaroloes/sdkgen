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

type Generator interface {
	setTemplateDir(dir string)
	Generate(api *parser.Api, config Config) error
}

func New(language Language) (Generator, error) {
	var gen Generator
	var tplDir string

	switch language {
	case ObjC:
		gen = &ObjCGen{}
		tplDir = path.Join(templateDir, strings.ToLower(language.String()))
		//	case Android:
		//	case Swift:
	default:
		return nil, errors.Annotate(ErrLangNotSupported, language.String())
	}
	gen.setTemplateDir(tplDir)
	return gen, nil
}
