package parser

const (
	MethodGET = "GET"
	MethodPOST = "POST"
	MethodPUT = "PUT"
	MethodDELETE = "DELETE"
)

const (
	endpointNoBody = iota
	endpointRequestBody
	endpointResponseBody
)