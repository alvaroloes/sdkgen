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
	ErrMultipleHosts  = errors.New("multiple hosts/scheme API is not supported")
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

const authToken = "AUTH"

var supportedMethods = GET.String() + "|" + POST.String() + "|" + PUT.String() + "|" + DELETE.String()

var endpointRegexp = regexp.MustCompile(`(?m)^\s*(` + authToken + `)?\s*?(` + supportedMethods + `)\s*(.*)$`)

const (
	endpointFullIndex = 2 * iota
	authIndex
	methodIndex
	urlIndex
)

var (
	requestBodyMarkRegexp  = regexp.MustCompile(`(?m)^\s*\-\>`)
	responseBodyMarkRegexp = regexp.MustCompile(`(?m)^\s*\<\-`)
)

const segmentParameterPrefix = ":"

type API struct {
	BaseURL   string
	Endpoints []Endpoint
}

func (api *API) extractBaseURL() error {
	var scheme, host string
	for _, ep := range api.Endpoints {
		if scheme == "" {
			scheme = ep.URL.Scheme
		} else if scheme != ep.URL.Scheme {
			return errors.Annotatef(ErrMultipleHosts, `found schemes "%s" and "%s"`, scheme, ep.URL.Scheme)
		}

		if host == "" {
			host = ep.URL.Host
		} else if host != ep.URL.Host {
			return errors.Annotatef(ErrMultipleHosts, `found hosts "%s" and "%s"`, host, ep.URL.Host)
		}
	}
	api.BaseURL = scheme + "://" + host
	return nil
}

type Endpoint struct {
	Authenticates bool
	Method        HTTPMethod
	URL           *url.URL
	Resources     []Resource
	RequestSpec   string
	RequestBody   interface{}
	ResponseSpec  string
	ResponseBody  interface{}
}

func (ep *Endpoint) extractResources() error {
	for _, segment := range strings.Split(ep.URL.Path, "/")[1:] {
		segment = strings.TrimSpace(segment)
		if segment == "" {
			continue
		}
		if strings.HasPrefix(segment, segmentParameterPrefix) {
			if len(ep.Resources) == 0 {
				return errors.Annotate(ErrNoRootResource, "in URL "+ep.URL.String())
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
		var requestBody []byte
		ep.RequestSpec, requestBody = findSpecAndJSONObject(endpointData[match[1]:])
		if err := json.Unmarshal(requestBody, &ep.RequestBody); err != nil {
			return errors.Annotate(err, "while parsing JSON request body of "+ep.URL.String())
		}
	}
	match = responseBodyMarkRegexp.FindIndex(endpointData)
	if match != nil {
		var responseBody []byte
		ep.ResponseSpec, responseBody = findSpecAndJSONObject(endpointData[match[1]:])
		if err := json.Unmarshal(responseBody, &ep.ResponseBody); err != nil {
			return errors.Annotate(err, "while parsing JSON response body of "+ep.URL.String())
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
		endpoint := Endpoint{}

		urlString := string(spec[match[urlIndex]:match[urlIndex+1]])
		parsedURL, err := url.Parse(urlString)
		if err != nil {
			return nil, errors.Annotate(err, "while parsing the URL "+urlString)
		}
		endpoint.URL = parsedURL

		httpMethod, err := HTTPMethodString(string(spec[match[methodIndex]:match[methodIndex+1]]))
		if err != nil {
			return nil, errors.Annotate(err, "while extracting the HTTP method of "+endpoint.URL.String())
		}
		endpoint.Method = httpMethod

		endpoint.Authenticates = match[authIndex] >= 0

		if err := endpoint.extractResources(); err != nil {
			return nil, errors.Annotate(err, "while extracting resources of "+endpoint.URL.String())
		}

		var endpointDataFinalIndex int
		if i < len(endpointMatches)-1 {
			endpointDataFinalIndex = endpointMatches[i+1][endpointFullIndex]
		} else {
			endpointDataFinalIndex = len(spec)
		}

		if err := endpoint.extractBodies(spec[match[endpointFullIndex+1]:endpointDataFinalIndex]); err != nil {
			return nil, errors.Annotate(err, "while extracting bodies of "+endpoint.URL.String())
		}

		// TODO: Check here if we understand the token response if authenticates
		api.Endpoints = append(api.Endpoints, endpoint)
	}

	if err := api.extractBaseURL(); err != nil {
		return nil, errors.Annotate(err, "while extracting the base URL")
	}

	return &api, nil
}

// findSpecAndJSONObject returns a string with the specification and
// a byte slice containing the first JSON object or array in the provided bytes
func findSpecAndJSONObject(bytes []byte) (string, []byte) {
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

	return strings.TrimSpace(string(bytes[:from])), bytes[from : to+1]
}
