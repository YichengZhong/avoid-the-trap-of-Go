package main

import (
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

	var svcmath ServiceMetricMath
	svcmath = ArithmeticServiceMath{}

	var svcstring ServiceMetricString
	svcstring = ArithmeticServiceString{}

	endpointall := EndpointAll{
		MathEndpoint:   MakeArithmeticEndpoint(svcmath),
		StringEndpoint: MakeArithStringEndpoint(svcstring),
	}

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	r := MakeHttpHandler(ctx, endpointall, logger)

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
