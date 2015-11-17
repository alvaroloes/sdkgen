package gen

import (
	"github.com/alvaroloes/sdkgen/parser"
	"github.com/alvaroloes/sdkgen/gen/objc"
	"github.com/juju/errors"
)

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

type Generator interface {
	Generate(api *parser.Api) error
}

func New(language Language) (Generator, error) {
	var gen Generator
	switch (language) {
	case ObjC:
		gen = &objc.ObjCGen{}
//	case Android:
//	case Swift:
	default:
		return nil, errors.Annotate(ErrLangNotSupported, language.String())
	}
	return gen, nil
}