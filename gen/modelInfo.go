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

func (mi *modelInfo) mergePropertiesFromBody(body interface{}) {
	// TODO
}
