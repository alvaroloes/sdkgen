package gen

import (
	"strings"

	"github.com/alvaroloes/sdkgen/parser"
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
			prop.Type, propertyDependencies = objCType(prop, config)
			modelInfo.Properties[propSpec] = prop
			modelInfo.Dependencies = append(modelInfo.Dependencies, propertyDependencies...)
			// TODO: Property attributes
		}
	}
}

func objCType(prop property, config Config) (string, []string) {
	var res string
	var dependencies []string

	if prop.IsArray {
		res = "NSArray<"
	}

	objCType, typeFound := objCTypePerGoType[prop.Type]
	if typeFound {
		// In Objective C an array of booleans needs to be an array of NSNumbers
		if prop.IsArray && objCType.Name == typeBOOL {
			res += typeNSNumber
		} else {
			res += objCType.Name
		}
	} else {
		modelName := config.APIPrefix + strings.Title(prop.Type)
		dependencies = append(dependencies, modelName)
		res += modelName
	}

	if prop.IsArray {
		res += " *> *"
	} else if !typeFound || objCType.Pointer {
		// If type is not found, it means that the type is a class, so we need a pointer
		res += " *"
	} else {
		res += " "
	}

	return res, dependencies
}
