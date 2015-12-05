package gen

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
	Name string
	Type string
}

type endpointInfo struct {
	Method        string
	URLPath       string
	SegmentParams []string
	ResponseType  ResponseType
}
