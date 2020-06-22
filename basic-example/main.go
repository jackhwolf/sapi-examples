package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jackhwolf/sapi"
)

func main() {
	// create new router with url prefix `/prefix`
	rtr := sapi.NewRouter("/prefix/")

	// assign sampleData function to GET /prefix/sample
	rtr.AddRoute(func(ctx context.Context, payload sapi.Payload) *sapi.HandlerReturn {
		sample := struct {
			Message string
			Time    int
		}{"Hello", 123}
		return &sapi.HandlerReturn{sample, 200, nil}
	}, "/sample", http.MethodGet)

	lambda.Start(rtr.HandleLambda)
}
