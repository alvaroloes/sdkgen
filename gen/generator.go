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

type Generator struct {
	gen        specificGenerator
	api        *parser.Api
	modelsInfo map[string]*modelInfo
	config     Config
}

type specificGenerator interface {
	setTemplateDir(dir string)
	generate(config Config, api *parser.Api, modelsInfo map[string]*modelInfo) error
}

func (g *Generator) Generate() error {
	return g.gen.generate(g.config, g.api, g.modelsInfo)
}

func (g *Generator) extractModelsInfo() error {
	g.modelsInfo = map[string]*modelInfo{}
	for _, endpoint := range g.api.Endpoints {
		// Extract the resource whose information is contained in this endpoint
		mainResource := endpoint.Resources[len(endpoint.Resources)-1]
		modelName := mainResource.Name

		// Extract the endpoint info and set it to the corresponding model
		g.setEndpointInfo(modelName, endpoint)

		// Merge the properties form the request and response bodies into
		// the corresponding model
		g.mergeModelProperties(modelName, endpoint.RequestBody)
		g.mergeModelProperties(modelName, endpoint.ResponseBody)
	}
	return nil
}

func (g *Generator) getURLPathForModels(url *url.URL) string {
	//TODO: Strip version path when versioning is supported
	return url.Path
}

func (g *Generator) mergeModelProperties(modelName string, body interface{}) {
	if body == nil {
		return
	}

	mInfo := g.getModelOrCreate(modelName)

	switch reflect.TypeOf(body).Kind() {
	case reflect.Map:
		props := body.(map[string]interface{})
		for propSpec, val := range props {
			g.mergeModelProperty(mInfo, propSpec, val)
		}
	case reflect.Array:
		// Get the first object of the array and start again
		arrayVal := reflect.ValueOf(body)
		if arrayVal.Len() == 0 {
			return
		}
		g.mergeModelProperties(modelName, arrayVal.Index(0).Interface())
	default:
		// This means either an empty response or a non resource response. Ignore it
		return
	}
}

func (g *Generator) mergeModelProperty(mInfo *modelInfo, propSpec string, propVal interface{}) {
	prop := property{
		Name: g.getPropName(propSpec),
		Type: g.getPropType(propSpec, propVal),
	}

	_, found := mInfo.getProperty(prop.Name)
	if found {
		// TODO: What to do now?. Either the old or the new one must have preference
		// We could check if prop.Type's are equal. If not -> log a warning
		// Right now old one has preference

	} else {
		mInfo.Properties = append(mInfo.Properties, prop)
	}

	valKind := reflect.TypeOf(propVal).Kind()
	if valKind == reflect.Map || valKind == reflect.Array {
		g.mergeModelProperties(mInfo.Name, propVal)
	}
}

func (g *Generator) getPropName(nameSpec string) string {
	// TODO: Allow overriding the property name when nameSpec: "prop1: name=desiredName". This would have preference
	return nameSpec
}

func (g *Generator) getPropType(nameSpec string, propVal interface{}) string {
	// TODO: Allow overriding the property type when nameSpec: "prop1: type=desiredType". This would have preference
	value := reflect.TypeOf(propVal)
	switch value.Kind() {
	case reflect.Map:
		// The value is an object, the type name is the property name
		return nameSpec
	case reflect.Array:
		arrayVal := reflect.ValueOf(propVal)
		if arrayVal.Len() == 0 {
			return ""
		}
		return g.getPropType(nameSpec, arrayVal.Index(0).Interface())
	default:
		return value.String()
	}
}

func (g *Generator) setEndpointInfo(modelName string, endpoint parser.Endpoint) {
	mInfo := g.getModelOrCreate(modelName)
	mInfo.EndpointsInfo = append(mInfo.EndpointsInfo, endpointInfo{
		Method:        endpoint.Method,
		URLPath:       g.getURLPathForModels(endpoint.URL),
		SegmentParams: extractSegmentParamsRenamingDups(endpoint.Resources),
		ResponseType:  getResponseType(endpoint.ResponseBody),
	})
}

func (g *Generator) getModelOrCreate(modelName string) *modelInfo {
	mInfo, modelExists := g.modelsInfo[modelName]
	if !modelExists {
		mInfo = &modelInfo{
			Name: modelName,
		}
		g.modelsInfo[modelName] = mInfo
	}
	return mInfo
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
