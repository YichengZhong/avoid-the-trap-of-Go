package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// decodeArithmeticRequestMath decode request params to struct
func decodeArithmeticRequestMath(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	requestType, ok := vars["type"]
	if !ok {
		return nil, ErrorBadRequestMath
	}

	pa, ok := vars["a"]
	if !ok {
		return nil, ErrorBadRequestMath
	}

	pb, ok := vars["b"]
	if !ok {
		return nil, ErrorBadRequestMath
	}

	a, _ := strconv.Atoi(pa)
	b, _ := strconv.Atoi(pb)

	return ArithmeticRequest{
		RequestType: requestType,
		A:           a,
		B:           b,
	}, nil
}

// encodeArithmeticResponseMath encode response to return
func encodeArithmeticResponseMath(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
