package gen

import (
	"reflect"

	"net/url"

	"strings"

	"github.com/alvaroloes/sdkgen/parser"
	"github.com/jinzhu/inflection"
	"github.com/juju/errors"
)

type ResponseKind int

const (
	ObjectResponse ResponseKind = iota
	MapResponse
	ArrayResponse
	EmptyResponse
)

const (
	propertySpecSeparator = ":"
	attrSeparator         = ";"
	attrKeyValueSeparator = "="
)
const (
	attrKeyName = "name"
	attrKeyType = "type"
	attrKeyMap  = "map"
)

var crudNamePerMethod = map[parser.HTTPMethod]string{
	parser.GET:    "fetch",
	parser.POST:   "create",
	parser.PUT:    "update",
	parser.DELETE: "delete",
}

type modelInfo struct {
	Name                  string
	OriginalName          string
	Properties            map[string]property
	EndpointsInfo         []endpointInfo
	ModelDependencies     map[*modelInfo]struct{}
	EndpointsDependencies map[*modelInfo]struct{}

	// These are responsibility of language specific generators
	LangSpecificData map[string]interface{}
}

func (mi *modelInfo) DependsOnModel(modelName string) bool {
	for dep, _ := range mi.ModelDependencies {
		if dep.Name == modelName {
			return true
		}
	}
	return false
}

func newModelInfo(name string) *modelInfo {
	return &modelInfo{
		Name:                  name,
		OriginalName:          name,
		Properties:            make(map[string]property),
		ModelDependencies:     make(map[*modelInfo]struct{}),
		EndpointsDependencies: make(map[*modelInfo]struct{}),
	}
}

type property struct {
	Name      string
	NameLabel string
	Type      string
	TypeLabel string
	IsArray   bool
	IsMap     bool
}

func newProperty(propertySpec string, val interface{}) property {
	var p property
	attributes := newPropertyAttributes(propertySpec)

	p.Name = attributes.name
	if attributes.nameLabel != "" {
		p.NameLabel = attributes.nameLabel
	} else {
		p.NameLabel = p.Name
	}
	p.extractType(attributes, val)
	return p
}

func (p *property) extractType(attributes propertyAttributes, val interface{}) {
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
	if attributes.forcedType != "" {
		p.Type = attributes.forcedType
	}
	p.TypeLabel = p.Type
	p.IsMap = attributes.forceAsMap
}

type propertyAttributes struct {
	name       string
	nameLabel  string
	forcedType string
	forceAsMap bool
}

func newPropertyAttributes(propertySpec string) (res propertyAttributes) {
	nameAndAttributes := strings.Split(propertySpec, propertySpecSeparator)
	res.name = strings.TrimSpace(nameAndAttributes[0])
	if len(nameAndAttributes) < 2 {
		return
	}
	attributes := strings.Split(nameAndAttributes[1], attrSeparator)
	for _, attr := range attributes {
		keyVal := strings.Split(attr, attrKeyValueSeparator)
		val := ""
		if len(keyVal) > 1 {
			val = keyVal[1]
		}
		switch strings.TrimSpace(keyVal[0]) {
		case attrKeyName:
			res.nameLabel = strings.TrimSpace(val)
		case attrKeyType:
			res.forcedType = strings.TrimSpace(val)
		case attrKeyMap:
			res.forceAsMap = true
		}
	}
	return
}

type endpointInfo struct {
	ResourceModel  *modelInfo
	RequestModel   *modelInfo
	ResponseModel  *modelInfo
	Method         parser.HTTPMethod
	URLPath        string
	URLQueryParams url.Values
	SegmentParams  []string
	ResponseKind   ResponseKind
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
	return epi.ResponseKind == ArrayResponse
}

func (epi *endpointInfo) IsMapResponse() bool {
	return epi.ResponseKind == MapResponse
}

func (epi *endpointInfo) HasResponse() bool {
	return epi.ResponseKind != EmptyResponse
}
