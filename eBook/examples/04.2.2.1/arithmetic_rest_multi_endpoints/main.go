package main

import (
	"GoURL/arithmetic_rest_multi_endpoints/endpoints"
	"GoURL/arithmetic_rest_multi_endpoints/service"
	"GoURL/arithmetic_rest_multi_endpoints/transports"
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()
	errChan := make(chan error)

	var svcmath service.ServiceMetricMath
	svcmath = service.ArithmeticServiceMath{}

	var svcstring service.ServiceMetricString
	svcstring = service.ArithmeticServiceString{}

	endpointall := endpoints.EndpointAll{
		MathEndpoint:   endpoints.MakeArithmeticEndpoint(svcmath),
		StringEndpoint: endpoints.MakeArithStringEndpoint(svcstring),
	}

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	r := transports.MakeHttpHandler(ctx, endpointall, logger)

	go func() {
		fmt.Println("Http Server start at port:9000")
		handler := r
		errChan <- http.ListenAndServe(":9000", handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)
}
