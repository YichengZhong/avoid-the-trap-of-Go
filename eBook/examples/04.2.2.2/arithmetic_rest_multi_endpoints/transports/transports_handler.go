package transports

import (
	"GoURL/arithmetic_rest_multi_endpoints/endpoints"
	"context"
	"errors"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ErrorBadRequestMath   = errors.New("invalid request parameter")
	ErrorBadRequestString = errors.New("invalid request parameter")
)

// MakeHttpHandlerMath make http handler use mux
func MakeHttpHandler(ctx context.Context, endpoint endpoints.EndpointAll, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(kithttp.DefaultErrorEncoder),
	}

	r.Methods("POST").Path("/calculate/{type}/{a}/{b}").Handler(kithttp.NewServer(
		endpoint.MathEndpoint,
		decodeArithmeticRequestMath,
		encodeArithmeticResponseMath,
		options...,
	))

	r.Methods("POST").Path("/string/{type}/{a}/{b}").Handler(kithttp.NewServer(
		endpoint.StringEndpoint,
		decodeArithmeticRequestString,
		encodeArithmeticResponseString,
		options...,
	))

	return r
}
