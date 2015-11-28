package gen

import (
	"net/url"
	"path"
	"reflect"
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
	modelsMap := map[string]*modelInfo{}
	for _, endpoint := range g.api.Endpoints {
		//Extract the resource whose information is contained in this endpoint
		mainResource := endpoint.Resources[len(endpoint.Resources)-1]
		//Get or create the model info for this model name
		mInfo, modelExists := modelsMap[mainResource.Name]
		if !modelExists {
			mInfo = &modelInfo{
				Name: mainResource.Name,
			}
		}
		// Add this new endpoint to the model info with all the data needed
		mInfo.EndpointsInfo = append(mInfo.EndpointsInfo, endpointInfo{
			Method:        endpoint.Method,
			URLPath:       g.getURLPathForModels(endpoint.URL),
			SegmentParams: extractSegmentParamsRenamingDups(endpoint.Resources),
			ResponseType:  getResponseType(endpoint.ResponseBody),
		})
		// Merge the properties from request
		// TODO: Take into account that maybe I need to pass here the modelsMap to register nested models under the corresponding name
		mInfo.mergePropertiesFromBody(endpoint.RequestBody)
		mInfo.mergePropertiesFromBody(endpoint.ResponseBody)
	}
	return nil
}

func (g *Generator) getURLPathForModels(url *url.URL) string {
	//TODO: Strip version path when versioning is supported
	return url.Path
}

func getResponseType(body interface{}) ResponseType {
	if body == nil {
		return EmptyResponse
	}
	switch reflect.TypeOf(body).Kind() {
	case reflect.Map:
		return ObjectResponse
	case reflect.Array:
		return ArrayResponse
	default:
		return EmptyResponse
	}
}

func extractSegmentParamsRenamingDups(resources []parser.Resource) []string {
	segmentParams := []string{}
	for _, r := range resources {
		//TODO: use r.Name to avoid duplicates
		segmentParams = append(segmentParams, r.Parameters...)
	}
	return segmentParams
}

// New creates a new Generator for the API and configured for the language passed.
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
