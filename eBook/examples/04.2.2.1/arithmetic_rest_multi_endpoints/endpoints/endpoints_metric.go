package endpoints

import (
	"GoURL/arithmetic_rest_multi_endpoints/service"
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"strings"
)

var (
	ErrMetricInvalidRequestType = errors.New("RequestType has only four type: Add,Subtract,Multiply,Divide")
)

// ArithmeticRequest define request struct
type ArithmeticRequest struct {
	RequestType string `json:"request_type"`
	A           int    `json:"a"`
	B           int    `json:"b"`
}

// ArithmeticResponse define response struct
type ArithmeticResponse struct {
	Result int   `json:"result"`
	Error  error `json:"error"`
}

// CalculateEndpoint define endpoint
type ArithmeticEndpoint endpoint.Endpoint

// MakeArithmeticEndpoint make endpoint
func MakeArithmeticEndpoint(svc service.ServiceMetricMath) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ArithmeticRequest)

		var (
			res, a, b int
			calError  error
		)

		a = req.A
		b = req.B

		if strings.EqualFold(req.RequestType, "Add") {
			res = svc.Add(a, b)
		} else if strings.EqualFold(req.RequestType, "Subtract") {
			res = svc.Subtract(a, b)
		} else if strings.EqualFold(req.RequestType, "Multiply") {
			res = svc.Multiply(a, b)
		} else if strings.EqualFold(req.RequestType, "Divide") {
			res, calError = svc.Divide(a, b)
		} else {
			return nil, ErrMetricInvalidRequestType
		}

		return ArithmeticResponse{Result: res, Error: calError}, nil
	}
}
