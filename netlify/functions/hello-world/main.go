package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("This message will show up in the CLI console.")

	type HelloWorldResponse struct {
		Message string `json:"message"`
	}
	helloWorldResponse := &HelloWorldResponse{
		Message: "Hello, world!",
	}
	jsonData, _ := json.Marshal(helloWorldResponse)

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            string(jsonData),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}
