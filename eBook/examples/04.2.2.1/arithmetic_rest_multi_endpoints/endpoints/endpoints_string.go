package endpoints

import (
	"GoURL/arithmetic_rest_multi_endpoints/service"
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"strings"
)

var (
	ErrStringInvalidRequestType = errors.New("RequestType has only four type: Add,Diff")
)

// ArithStringRequest define request struct
type ArithStringRequest struct {
	RequestType string `json:"request_type"`
	A           string `json:"a"`
	B           string `json:"b"`
}

// ArithStringResponse define response struct
type ArithStringResponse struct {
	Result string `json:"result"`
	Error  error  `json:"error"`
}

// CalculateEndpoint define endpoint
type ArithStringEndpoint endpoint.Endpoint

// MakeArithStringEndpoint make endpoint
func MakeArithStringEndpoint(svc service.ServiceMetricString) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ArithStringRequest)

		var (
			res, a, b string
			calError  error
		)

		a = req.A
		b = req.B

		if strings.EqualFold(req.RequestType, "Add") {
			res = svc.Add(a, b)
		} else {
			return nil, ErrStringInvalidRequestType
		}

		return ArithStringResponse{Result: res, Error: calError}, nil
	}
}
