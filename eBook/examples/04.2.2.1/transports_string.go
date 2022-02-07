package main

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

// MakeHttpHandlerMath make http handler use mux
func MakeHttpHandlerString(ctx context.Context, endpoint endpoint.Endpoint, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(kithttp.DefaultErrorEncoder),
	}

	r.Methods("POST").Path("/string/{type}/{a}/{b}").Handler(kithttp.NewServer(
		endpoint,
		decodeArithmeticRequestString,
		encodeArithmeticResponseString,
		options...,
	))

	return r
}

// decodeArithmeticRequestMath decode request params to struct
func decodeArithmeticRequestString(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	requestType, ok := vars["type"]
	if !ok {
		return nil, ErrorBadRequestString
	}

	pa, ok := vars["a"]
	if !ok {
		return nil, ErrorBadRequestString
	}

	pb, ok := vars["b"]
	if !ok {
		return nil, ErrorBadRequestString
	}

	return ArithStringRequest{
		RequestType: requestType,
		A:           pa,
		B:           pb,
	}, nil
}

// encodeArithmeticResponseMath encode response to return
func encodeArithmeticResponseString(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
