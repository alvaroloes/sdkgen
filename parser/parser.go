package parser

import (
	"encoding/json"
	"net/url"
	"regexp"
	"strings"

	"github.com/juju/errors"
)

var (
	ErrNoRootResource = errors.New("root REST resource not found")
)

//go:generate enumer -type=HTTPMethod

type HTTPMethod int

const (
	UNKNOWN_HTTP_METHOD HTTPMethod = iota
	GET
	POST
	PUT
	DELETE
)

var supportedMethods = GET.String() + "|" + POST.String() + "|" + PUT.String() + "|" + DELETE.String()

var endpointRegexp = regexp.MustCompile(`(?m)^\s*(` + supportedMethods + `)\s*(.*)$`)

const (
	endpointFullIndex = 2 * iota
	methodIndex
	urlIndex
)

var (
	//	queryMarkRegexp = regexp.MustCompile(`(?m)^\s*-?`) // TODO
	requestBodyMarkRegexp  = regexp.MustCompile(`(?m)^\s*\-\>`)
	responseBodyMarkRegexp = regexp.MustCompile(`(?m)^\s*\<\-`)
)

const segmentParameterPrefix = ":"

type API struct {
	//	Version   string //TODO
	Endpoints []Endpoint
}

type Endpoint struct {
	Method       HTTPMethod
	URLString    string //TODO: Remove this. It's redundant
	URL          *url.URL
	Resources    []Resource
	RequestBody  interface{}
	ResponseBody interface{}
}

func (ep *Endpoint) extractResources() error {
	var err error
	ep.URL, err = url.Parse(ep.URLString)
	if err != nil {
		return errors.Trace(err)
	}
	for _, segment := range strings.Split(ep.URL.Path, "/")[1:] {
		segment = strings.TrimSpace(segment)
		if segment == "" {
			continue
		}
		if strings.HasPrefix(segment, segmentParameterPrefix) {
			if len(ep.Resources) == 0 {
				return errors.Annotate(ErrNoRootResource, "in URL "+ep.URLString)
			}
			lastResource := &ep.Resources[len(ep.Resources)-1]
			parameterName := segment[len(segmentParameterPrefix):]
			lastResource.Parameters = append(lastResource.Parameters, parameterName)
		} else {
			ep.Resources = append(ep.Resources, Resource{
				Name: segment,
			})
		}
	}

	return nil
}

func (ep *Endpoint) extractBodies(endpointData []byte) error {
	match := requestBodyMarkRegexp.FindIndex(endpointData)
	if match != nil {
		requestBody := findJSONObject(endpointData[match[1]:])
		if err := json.Unmarshal(requestBody, &ep.RequestBody); err != nil {
			return errors.Annotate(err, "while parsing JSON request body of "+ep.URLString)
		}
	}
	match = responseBodyMarkRegexp.FindIndex(endpointData)
	if match != nil {
		responseBody := findJSONObject(endpointData[match[1]:])
		if err := json.Unmarshal(responseBody, &ep.ResponseBody); err != nil {
			return errors.Annotate(err, "while parsing JSON response body of "+ep.URLString)
		}
	}
	return nil
}

type Resource struct {
	Name       string
	Parameters []string
}

func NewAPI(spec []byte) (*API, error) {
	var api API
	endpointMatches := endpointRegexp.FindAllSubmatchIndex(spec, -1)
	for i, match := range endpointMatches {
		endpoint := Endpoint{
			URLString: string(spec[match[urlIndex]:match[urlIndex+1]]),
		}

		httpMethod, err := HTTPMethodString(string(spec[match[methodIndex]:match[methodIndex+1]]))
		if err != nil {
			return nil, errors.Annotate(err, "while extracting the HTTP method of "+endpoint.URLString)
		}
		endpoint.Method = httpMethod

		if err := endpoint.extractResources(); err != nil {
			return nil, errors.Annotate(err, "while extracting resources of "+endpoint.URLString)
		}

		var endpointDataFinalIndex int
		if i < len(endpointMatches)-1 {
			endpointDataFinalIndex = endpointMatches[i+1][endpointFullIndex]
		} else {
			endpointDataFinalIndex = len(spec)
		}

		if err := endpoint.extractBodies(spec[match[endpointFullIndex+1]:endpointDataFinalIndex]); err != nil {
			return nil, errors.Annotate(err, "while extracting bodies of "+endpoint.URLString)
		}
		api.Endpoints = append(api.Endpoints, endpoint)
	}
	return &api, nil
}

// findJSONObject returns a byte slice containing the first JSON object or array
// in the provided bytes
func findJSONObject(bytes []byte) []byte {
	var opening, closing byte
	var from, to int

	for i, b := range bytes {
		if b == '{' || b == '[' {
			from = i
			opening = b
			if b == '{' {
				closing = '}'
			} else {
				closing = ']'
			}
			break
		}
	}

	level := 1
	for i, b := range bytes[from+1:] {
		if b == opening {
			level++
		} else if b == closing {
			level--
		}
		if level <= 0 {
			to = from + 1 + i
			break
		}
	}
	return bytes[from : to+1]
}
