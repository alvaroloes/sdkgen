package gen

import (
	"reflect"

	"github.com/alvaroloes/sdkgen/parser"
	"github.com/jinzhu/inflection"
)

type ResponseType int

const (
	ObjectResponse ResponseType = iota
	ArrayResponse
	EmptyResponse
)

type modelInfo struct {
	Name             string
	Properties       map[string]property
	EndpointsInfo    []endpointInfo
	LangSpecificData map[string]interface{}
}

func newModelInfo(name string) *modelInfo {
	return &modelInfo{
		Name:       inflection.Singular(name),
		Properties: make(map[string]property),
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
