package gen

import (
	"reflect"

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
	Dependencies     []string
	LangSpecificData map[string]interface{}
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
	Name    string
	Type    string
	IsArray bool
}

func newProperty(propertySpec string, val interface{}) property {
	var p property
	// TODO: Allow overriding the property name when nameSpec: "prop1: name=desiredName". This would have preference
	p.Name = propertySpec
	p.extractType(propertySpec, val)
	return p
}

func (p *property) extractType(propertySpec string, val interface{}) {
	// TODO: Allow overriding the property type when nameSpec: "prop1: type=desiredType". This would have preference
	value := reflect.TypeOf(val)
	switch value.Kind() {
	case reflect.Map:
		// The value is an object, the type name is the property name
		p.Type = inflection.Singular(p.Name)
	case reflect.Array, reflect.Slice:
		p.IsArray = true
		arrayVal := reflect.ValueOf(val)
		if arrayVal.Len() > 0 {
			p.extractType(propertySpec, arrayVal.Index(0).Interface())
		}
	default:
		p.Type = value.String()
	}
}

type endpointInfo struct {
	Method        parser.HTTPMethod
	URLPath       string
	SegmentParams []string
	ResponseType  ResponseType
}

func (epi *endpointInfo) CRUDMethodName() (string, error) {
	if epi.Method == parser.UNKNOWN_HTTP_METHOD {
		return "", errors.Errorf("Unknown http method for endopint %s", epi.URLPath)
	}
	return crudNamePerMethod[epi.Method], nil
}

func (epi *endpointInfo) IsArrayResponse() bool {
	return epi.ResponseType == ArrayResponse
}
