package gen

import (
	"strings"

	"github.com/alvaroloes/sdkgen/parser"
"text/template"
)

type objCTypeInfo struct {
	Name    string
	Pointer bool
}

const (
	typeBOOL     = "BOOL"
	typeNSNumber = "NSNumber"
	typeNSString = "NSString"
)

var objCTypePerGoType = map[string]objCTypeInfo{
	"bool":    {Name: typeBOOL, Pointer: false},
	"float64": {Name: typeNSNumber, Pointer: true},
	"string":  {Name: typeNSString, Pointer: true},
}

type ObjCGen struct {
}

func (gen *ObjCGen) adaptModelsInfo(modelsInfo map[string]*modelInfo, api *parser.API, config Config) {
	for _, modelInfo := range modelsInfo {
		modelInfo.Name = config.APIPrefix + strings.Title(modelInfo.Name)
		for propSpec, prop := range modelInfo.Properties {
			var propertyDependencies []string
			prop.Type, prop.TypeLabel, propertyDependencies = objCType(prop, config)
			modelInfo.Properties[propSpec] = prop
			modelInfo.ModelDependencies = append(modelInfo.ModelDependencies, propertyDependencies...)
			// TODO: Property attributes
		}
	}
}

func (gen *ObjCGen) funcMap() template.FuncMap {
	return objCFuncMap
}

func objCType(prop property, config Config) (string, string, []string) {
	var typeName, typeLabel string
	var dependencies []string

	if prop.IsArray {
		typeLabel = "NSArray<"
	}

	objCType, typeFound := objCTypePerGoType[prop.Type]
	if typeFound {
		typeName = objCType.Name
		// In Objective C an array of booleans needs to be an array of NSNumbers
		if prop.IsArray && typeName == typeBOOL {
			typeLabel += typeNSNumber
		} else {
			typeLabel += typeName
		}
	} else {
		typeName = config.APIPrefix + strings.Title(prop.Type)
		dependencies = append(dependencies, typeName)
		typeLabel += typeName
	}

	if prop.IsArray {
		typeLabel += " *> *"
	} else if !typeFound || objCType.Pointer {
		// If type is not found, it means that the type is a class, so we need a pointer
		typeLabel += " *"
	} else {
		typeLabel += " "
	}

	return typeName, typeLabel, dependencies
}
