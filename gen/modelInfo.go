package gen

import (
	"reflect"

	"net/url"

	"strings"

	"github.com/alvaroloes/sdkgen/parser"
	"github.com/jinzhu/inflection"
	"github.com/juju/errors"
)

type ResponseType int

const (
	ObjectResponse ResponseType = iota
	ArrayResponse
	EmptyResponse
)

const (
	propertySpecSeparator         = ":"
	propertyAttrSeparator         = ";"
	propertyAttrKeyValueSeparator = "="
)
const (
	propertyAttrKeyName = "name"
	propertyAttrKeyType = "type"
	propertyAttrKeyMap = "map"
)

var crudNamePerMethod = map[parser.HTTPMethod]string{
	parser.GET:    "fetch",
	parser.POST:   "create",
	parser.PUT:    "update",
	parser.DELETE: "delete",
}

type modelInfo struct {
	Name          string
	OriginalName  string
	Properties    map[string]property
	EndpointsInfo []endpointInfo

	// These are responsibility of language specific generators
	ModelDependencies []string
	LangSpecificData  map[string]interface{}
}

func newModelInfo(name string) *modelInfo {
	singularName := inflection.Singular(name)
	return &modelInfo{
		Name:         singularName,
		OriginalName: singularName,
		Properties:   make(map[string]property),
	}
}

type property struct {
	Name      string
	Type      string
	TypeLabel string
	IsArray   bool
	IsMap   bool
}

func newProperty(propertySpec string, val interface{}) property {
	var p property
	attributes := NewPropertyAttributes(propertySpec)
	if attributes.forcedName != "" {
		p.Name = attributes.forcedName
	} else {
		p.Name = attributes.name
	}
	p.extractType(attributes, val)
	p.IsMap = attributes.isMap
	return p
}

func (p *property) extractType(attributes propertyAttributes, val interface{}) {
	if attributes.forcedType != "" {
		p.Type = attributes.forcedType
	} else {
		value := reflect.TypeOf(val)
		switch value.Kind() {
		case reflect.Map:
			// The value is an object, the type name is the property name
			p.Type = inflection.Singular(p.Name)
		case reflect.Array, reflect.Slice:
			p.IsArray = true
			arrayVal := reflect.ValueOf(val)
			if arrayVal.Len() > 0 {
				p.extractType(attributes, arrayVal.Index(0).Interface())
			}
		default:
			p.Type = value.String()
		}
	}
	p.TypeLabel = p.Type
}

type propertyAttributes struct {
	name       string
	forcedName string
	forcedType string
	isMap bool
}

func NewPropertyAttributes(propertySpec string) (res propertyAttributes) {
	nameAndAttributes := strings.Split(propertySpec, propertySpecSeparator)
	res.name = nameAndAttributes[0]
	if len(nameAndAttributes) < 2 {
		return
	}
	attributes := strings.Split(nameAndAttributes[1], propertyAttrSeparator)
	for _, attr := range attributes {
		keyVal := strings.Split(attr, propertyAttrKeyValueSeparator)
		val := ""
		if len(keyVal) > 1 {
			val = keyVal[1]
		}
		switch keyVal[0] {
		case propertyAttrKeyName:
			res.forcedName = val
		case propertyAttrKeyType:
			res.forcedType = val
		case propertyAttrKeyMap:
			res.isMap = true
		}
	}
	return
}

type endpointInfo struct {
	Model          *modelInfo
	Method         parser.HTTPMethod
	URLPath        string
	URLQueryParams url.Values
	SegmentParams  []string
	ResponseType   ResponseType
}

func (epi *endpointInfo) CRUDMethodName() (string, error) {
	if epi.Method == parser.UNKNOWN_HTTP_METHOD {
		return "", errors.Errorf("Unknown http method for endopint %s", epi.URLPath)
	}
	return crudNamePerMethod[epi.Method], nil
}

func (epi *endpointInfo) NeedsModelParam() bool {
	switch epi.Method {
	case parser.POST, parser.PUT:
		return true
	default:
		return false
	}
}

func (epi *endpointInfo) IsArrayResponse() bool {
	return epi.ResponseType == ArrayResponse
}

func (epi *endpointInfo) HasResponse() bool {
	return epi.ResponseType != EmptyResponse
}
