package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("Get Skaping locations.")

	type SkapingLocationsResponse struct {
		Locations *SkapingLocations `json:"locations"`
	}

	skapingLocations := NewSkapingLocationScraperFromUrl("https://www.skaping.com/camera/map").ScrapLocations()

	skapingLocationsResponse := &SkapingLocationsResponse{&skapingLocations}

	skapingLocationsResponseJson, _ := json.Marshal(skapingLocationsResponse)

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            string(skapingLocationsResponseJson),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}
