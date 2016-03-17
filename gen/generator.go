package gen

import (
	"net/url"
	"os"
	"path"
	"reflect"
	"strings"
	"text/template"
	"time"

	"path/filepath"

	"fmt"

	"github.com/alvaroloes/sdkgen/parser"
	"github.com/jinzhu/inflection"
	"github.com/juju/errors"
)

//go:generate go-bindata -debug=$DEBUG -o=templates_bindata.go -pkg=$GOPACKAGE ../templates/...

var (
	ErrLangNotSupported = errors.New("language not supported")
)

type Language int

//go:generate stringer -type=Language

const (
	Android Language = iota
	ObjC
	Swift
)

const (
	templateDir                    = "./templates"
	commonTemplatesPath            = "common"
	modelTemplatePath              = "model"
	serviceTemplatePath            = "service"
	templateExt                    = ".tpl"
	fileNameModelNameInterpolation = "--ModelName--"
	fileNameAPINameInterpolation   = "--APIName--"
	fileNameAPIPrefixInterpolation = "--APIPrefix--"
	dirPermissions                 = 0777
)

// Config contains the needed configuration for the generator
type Config struct {
	OutputDir       string
	ModelsRelPath   string
	ServicesRelPath string
	APIName         string
	APIPrefix       string
}

type templateData struct {
	Config           Config
	API              *parser.API
	CurrentModelInfo *modelInfo
	AllModelsInfo    map[string]*modelInfo
	CurrentTime      time.Time
}

type languageSpecificGenerator interface {
	adaptModelsInfo(modelsInfo map[string]*modelInfo, api *parser.API, config Config)
	funcMap() template.FuncMap
}

// Generator contains all the information needed to generate the SDK in a specific language
type Generator struct {
	gen        languageSpecificGenerator
	api        *parser.API
	modelsInfo map[string]*modelInfo // Contains processed information to generate the models
	config     Config
	tplDir     string
}

func (g *Generator) Generate() error {
	// Extract the models info
	g.extractModelsInfo()
	// Adapt them to the specific language
	g.gen.adaptModelsInfo(g.modelsInfo, g.api, g.config)

	// Parse the base templates that contains common definitions
	baseTplsGlob := path.Join(g.tplDir, commonTemplatesPath, "*"+templateExt)
	baseTpls, err := template.New("base").Funcs(funcMap).Funcs(g.gen.funcMap()).ParseGlob(baseTplsGlob)
	if err != nil {
		return errors.Annotate(err, "when parsing common templates ("+baseTplsGlob+")")
	}

	// Read and parse the SDK general, model and service template files
	generalTplFileNames, generalTpls, err := g.parseTemplates(path.Join(g.tplDir, "*"+templateExt), baseTpls)
	if err != nil {
		return errors.Trace(err)
	}
	modelTplFileNames, modelTpls, err := g.parseTemplates(path.Join(g.tplDir, modelTemplatePath, "*"+templateExt), baseTpls)
	if err != nil {
		return errors.Trace(err)
	}
	serviceTplFileNames, serviceTpls, err := g.parseTemplates(path.Join(g.tplDir, serviceTemplatePath, "*"+templateExt), baseTpls)
	if err != nil {
		return errors.Trace(err)
	}

	// Create the needed directories
	apiDir := path.Join(g.config.OutputDir, g.config.APIName)
	modelsDir := path.Join(apiDir, g.config.ModelsRelPath)
	if err := os.MkdirAll(modelsDir, dirPermissions); err != nil {
		return errors.Annotatef(err, "when creating model directory")
	}
	servicesDir := path.Join(apiDir, g.config.ServicesRelPath)
	if err := os.MkdirAll(servicesDir, dirPermissions); err != nil {
		return errors.Annotatef(err, "when creating service directory")
	}

	// Generate the SDK files applying the templates
	err = g.generateGeneralFiles(generalTplFileNames, generalTpls, apiDir)
	if err != nil {
		return errors.Annotate(err, "when generating API files")
	}
	err = g.generatePerModelFiles(modelTplFileNames, modelTpls, modelsDir, func(modelInfo *modelInfo) bool {
		return len(modelInfo.Properties) == 0
	})
	if err != nil {
		return errors.Annotate(err, "when generating model files")
	}
	err = g.generatePerModelFiles(serviceTplFileNames, serviceTpls, servicesDir, func(modelInfo *modelInfo) bool {
		return len(modelInfo.EndpointsInfo) == 0
	})
	if err != nil {
		return errors.Annotate(err, "when generating service files")
	}

	return nil
}

func (g *Generator) parseTemplates(pathGlob string, baseTpl *template.Template) (fileNames []string, templates *template.Template, err error) {
	fileNames, err = filepath.Glob(pathGlob)
	if err != nil {
		return nil, nil, errors.Annotate(err, "when reading files in "+pathGlob)
	}
	templates, err = template.Must(baseTpl.Clone()).ParseFiles(fileNames...)
	if err != nil {
		return nil, nil, errors.Annotate(err, "when parsing service templates files in "+pathGlob)
	}
	return
}

func (g *Generator) generateGeneralFiles(templateFileNames []string, generalTpls *template.Template, apiDir string) error {
	for _, tplFileName := range templateFileNames {
		tplName := filepath.Base(tplFileName)
		// TODO: Do this concurrently
		// Get the name of the file, replacing some special strings in the template name
		repl := strings.NewReplacer(
			templateExt, "",
			fileNameAPINameInterpolation, g.config.APIName,
			fileNameAPIPrefixInterpolation, g.config.APIPrefix,
		)
		fileName := repl.Replace(tplName)
		err := generateFile(path.Join(apiDir, fileName), generalTpls.Lookup(tplName), templateData{
			Config:        g.config,
			API:           g.api,
			AllModelsInfo: g.modelsInfo,
			CurrentTime:   time.Now(),
		})
		if err != nil {
			return errors.Annotatef(err, "when generating API file %q", fileName)
		}
	}
	return nil
}

func (g *Generator) generatePerModelFiles(templateFileNames []string, modelTpls *template.Template, modelsDir string, filter func(modelInfo *modelInfo) bool) error {
	for _, tplFileName := range templateFileNames {
		tplName := filepath.Base(tplFileName)
		// Apply the templates to each model in the API
		for _, modelInfo := range g.modelsInfo {
			if filter(modelInfo) {
				continue
			}
			// TODO: Do this concurrently
			repl := strings.NewReplacer(
				templateExt, "",
				fileNameModelNameInterpolation, modelInfo.Name,
				fileNameAPINameInterpolation, g.config.APIName,
				fileNameAPIPrefixInterpolation, g.config.APIPrefix,
			)
			fileName := repl.Replace(tplName)
			err := generateFile(path.Join(modelsDir, fileName), modelTpls.Lookup(tplName), templateData{
				Config:           g.config,
				API:              g.api,
				CurrentModelInfo: modelInfo,
				AllModelsInfo:    g.modelsInfo,
				CurrentTime:      time.Now(),
			})
			if err != nil {
				return errors.Annotatef(err, "when generating model or service %q", modelInfo.Name)
			}
		}
	}
	return nil
}

func generateFile(filePath string, tpl *template.Template, data templateData) error {
	file, err := os.Create(filePath)
	if err != nil {
		return errors.Trace(err)
	}
	defer file.Close()
	return errors.Trace(tpl.Execute(file, data))
}

func (g *Generator) extractModelsInfo() {
	g.modelsInfo = map[string]*modelInfo{}
	for _, endpoint := range g.api.Endpoints {
		// Extract the resource whose information is contained in this endpoint
		mainResource := endpoint.Resources[len(endpoint.Resources)-1]
		resourceModelAttrs := modelAttributes{
			modelType:  mainResource.Name,
			forceAsMap: false,
		}
		requestModelAttrs := modelAttributesFromSpec(endpoint.RequestSpec)
		if requestModelAttrs.modelType == "" {
			requestModelAttrs.modelType = resourceModelAttrs.modelType
		}
		responseModelAttrs := modelAttributesFromSpec(endpoint.ResponseSpec)
		if responseModelAttrs.modelType == "" {
			responseModelAttrs.modelType = resourceModelAttrs.modelType
		}

		// TODO: forceASMap is not working until the template is updated
		// Extract the endpoint info and set it to the corresponding model
		g.setEndpointInfo(resourceModelAttrs, requestModelAttrs, responseModelAttrs, endpoint)

		// Merge the properties form the request and response bodies into
		// the corresponding model
		g.mergeModelProperties(requestModelAttrs.modelType, endpoint.RequestBody)
		g.mergeModelProperties(responseModelAttrs.modelType, endpoint.ResponseBody)
	}
	for name, mi := range g.modelsInfo {
		fmt.Printf("%s, %v\n", name, mi)
	}
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
	case reflect.Array, reflect.Slice:
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
	prop := newProperty(propSpec, propVal)

	_, found := mInfo.Properties[prop.Name]
	if found {
		// TODO: What to do now?. Either the old or the new one must have preference
		// We could check if prop.Type's are equal. If not -> log a warning
		// Right now old one has preference

	} else {
		mInfo.Properties[prop.Name] = prop
	}

	valKind := reflect.TypeOf(propVal).Kind()
	if valKind == reflect.Map || valKind == reflect.Array || valKind == reflect.Slice {
		//TODO: if !prop.IsRawMap {
		mInfo.ModelDependencies[g.getModelOrCreate(prop.Type)] = struct{}{}
		//TODO: }
		g.mergeModelProperties(prop.Type, propVal)
	}
}

func (g *Generator) setEndpointInfo(resourceModelAttrs, requestModelAttrs, responseModelAttrs modelAttributes, endpoint parser.Endpoint) {
	// Get/Create the needed models
	resourceModelInfo := g.getModelOrCreate(resourceModelAttrs.modelType)
	requestModelInfo := g.getModelOrCreate(requestModelAttrs.modelType)
	responseModelInfo := g.getModelOrCreate(responseModelAttrs.modelType)

	// Build the endpoint
	epi := endpointInfo{
		ResourceModel:  resourceModelInfo,
		RequestModel:   requestModelInfo,
		ResponseModel:  responseModelInfo,
		Method:         endpoint.Method,
		URLPath:        g.getURLPathForModels(endpoint.URL),
		URLQueryParams: endpoint.URL.Query(),
		SegmentParams:  extractSegmentParamsRenamingDups(endpoint.Resources),
		// TODO: Future: add RequestKind
		ResponseKind: getResponseKind(endpoint.ResponseBody, responseModelAttrs.forceAsMap), // TODO: response type can be "Map"
	}

	// Add the dependencies
	if epi.NeedsModelParam() {
		resourceModelInfo.EndpointsDependencies[requestModelInfo] = struct{}{}
	}
	if epi.HasResponse() /*TODO: && !epi.IsRawMapResponse()*/ {
		resourceModelInfo.EndpointsDependencies[responseModelInfo] = struct{}{}
	}

	resourceModelInfo.EndpointsInfo = append(resourceModelInfo.EndpointsInfo, epi)
}

func (g *Generator) getModelOrCreate(modelName string) *modelInfo {
	singularName := inflection.Singular(modelName)
	mInfo, modelExists := g.modelsInfo[singularName]
	if !modelExists {
		mInfo = newModelInfo(singularName)
		g.modelsInfo[singularName] = mInfo
	}
	return mInfo
}

func getResponseKind(body interface{}, forceAsMap bool) ResponseKind {
	if body == nil {
		return EmptyResponse
	}
	if forceAsMap {
		return MapResponse
	}

	switch reflect.TypeOf(body).Kind() {
	case reflect.Map:
		return ObjectResponse
	case reflect.Array, reflect.Slice:
		return ArrayResponse
	default:
		return EmptyResponse
	}
}

func extractSegmentParamsRenamingDups(resources []parser.Resource) []string {
	segmentParams := []string{}
	for _, r := range resources {
		//We assume that segment params have a unique name among the others in the same endpoint
		segmentParams = append(segmentParams, r.Parameters...)
	}
	return segmentParams
}

type modelAttributes struct {
	modelType  string
	forceAsMap bool
}

func modelAttributesFromSpec(modelSpec string) (res modelAttributes) {
	attributes := strings.Split(modelSpec, attrSeparator)
	for _, attr := range attributes {
		keyVal := strings.Split(attr, attrKeyValueSeparator)
		val := ""
		if len(keyVal) > 1 {
			val = keyVal[1]
		}
		switch strings.TrimSpace(keyVal[0]) {
		case attrKeyType:
			res.modelType = strings.TrimSpace(val)
		case attrKeyMap:
			res.forceAsMap = true
		}
	}
	return
}

// New creates a new Generator for the API and configured for the language passed.
func New(language Language, api *parser.API, config Config) (Generator, error) {
	var gen languageSpecificGenerator
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

	generator := Generator{
		gen:    gen,
		api:    api,
		config: config,
		tplDir: tplDir,
	}

	return generator, nil
}
