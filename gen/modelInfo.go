package gen

import "reflect"

type ResponseType int

const (
	ObjectResponse ResponseType = iota
	ArrayResponse
	EmptyResponse
)

type modelInfo struct {
	Name          string
	Properties    []property
	EndpointsInfo []endpointInfo
}

func (mi *modelInfo) getProperty(name string) (property, bool) {
	for _, prop := range mi.Properties {
		if prop.Name == name {
			return prop, true
		}
	}
	return property{}, false
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

func (p *property) extractType(propertySpec string, val interface{}) string {
	// TODO: Allow overriding the property type when nameSpec: "prop1: type=desiredType". This would have preference
	// TODO: Use an inflection to singularize the type (https://github.com/jinzhu/inflection)
	// TODO: Camelize the type? -> Better in specific generators as it depend on the language
	value := reflect.TypeOf(val)
	switch value.Kind() {
	case reflect.Map:
		// The value is an object, the type name is the property name
		return p.Name
	case reflect.Array, reflect.Slice:
		p.IsArray = true
		arrayVal := reflect.ValueOf(val)
		if arrayVal.Len() == 0 {
			return ""
		}
		return p.extractType(propertySpec, arrayVal.Index(0).Interface())
	default:
		return value.String()
	}
}

type endpointInfo struct {
	Method        string
	URLPath       string
	SegmentParams []string
	ResponseType  ResponseType
}
