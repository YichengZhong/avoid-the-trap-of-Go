package endpoints

import "github.com/go-kit/kit/endpoint"

type EndpointAll struct {
	MathEndpoint   endpoint.Endpoint
	StringEndpoint endpoint.Endpoint
}
