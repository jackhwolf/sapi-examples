package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jackhwolf/sapi"
)

// Dog defines our data model
type Dog struct {
	Name, Breed string
	Age, Weight int
}

// route we will use to post and get dogs
// for both GET and POST, the user will get back the data they send
func dogRoute(ctx context.Context, payload sapi.Payload) *sapi.HandlerReturn {
	dog := &Dog{}
	err := json.Unmarshal([]byte(payload.Body), dog)
	if err != nil {
		return &sapi.HandlerReturn{dog, http.StatusInternalServerError, err}
	}
	return &sapi.HandlerReturn{dog, http.StatusOK, nil}
}

func main() {
	rtr := sapi.NewRouter("/")

	rtr.AddRoute(dogRoute, "/dog", http.MethodGet, http.MethodPost)

	lambda.Start(rtr.HandleLambda)
}
