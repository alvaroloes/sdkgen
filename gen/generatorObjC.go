package gen

import "github.com/alvaroloes/sdkgen/parser"

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
		modelInfo.Name = config.APIPrefix + capitalize(modelInfo.Name)
		for propSpec, prop := range modelInfo.Properties {
			prop.Type = objCType(prop, config)
			modelInfo.Properties[propSpec] = prop
			// TODO: Property attributes
			// TODO: Insert the resource manager file as a template
			// TODO: Generate the model methods: Which name? param names?
		}
	}
}

func objCType(prop property, config Config) string {
	var res string

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
		res += config.APIPrefix + capitalize(prop.Type)
	}

	if prop.IsArray {
		res += " *> *"
	} else if !typeFound || objCType.Pointer {
		// If type is not found, it means that the type is a class, so we need a pointer
		res += " *"
	} else {
		res += " "
	}

	return res
}
