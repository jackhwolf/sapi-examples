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
	Name   string `json:"name"`
	Breed  string `json:"breed"`
	Age    int    `json:"age"`
	Weight int    `json:"weight"`
}

// route we will use to post and get dogs
// GET will return the default dog
// POST will echo back the dog the user sends
func dogRoute(ctx context.Context, payload sapi.Payload) *sapi.HandlerReturn {
	if payload.HTTPMethod == http.MethodPost {
		dog := &Dog{}
		err := json.Unmarshal([]byte(payload.Body), &dog)
		if err != nil {
			return &sapi.HandlerReturn{&Dog{}, http.StatusInternalServerError, err}
		}
		return &sapi.HandlerReturn{&dog, http.StatusOK, nil}
	}
	return &sapi.HandlerReturn{&Dog{Name: "skippy", Breed: "doodle", Age: 14, Weight: 25}, http.StatusOK, nil}
}

func main() {
	rtr := sapi.NewRouter("/")

	rtr.AddRoute(dogRoute, "/dog", http.MethodGet, http.MethodPost)

	lambda.Start(rtr.HandleLambda)
}
