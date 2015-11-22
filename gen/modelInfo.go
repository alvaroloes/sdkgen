package gen

import "net/url"

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
	URL           url.URL
	SegmentParams []string
	EmptyResponse bool
}
