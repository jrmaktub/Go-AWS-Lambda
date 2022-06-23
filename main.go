package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json: "What isyour name?"`
	Age  int    `json: "How old are you"`
}

type MyResponse struct {
	Message string `json: "Answer:`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {

	return MyResponse{Message: fmt.Sprintf("%s is %dyears old!", event.Name, event.Age)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
