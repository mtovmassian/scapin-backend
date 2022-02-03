package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
)

// const SKAPING_MAP_URL string = "http://example.com/"

const SKAPING_MAP_URL string = "https://www.skaping.com/camera/map"

type SkapingDataBodyResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	message := fmt.Sprintf("%s %s", request.HTTPMethod, request.Path)
	log.SetLevel(log.InfoLevel)
	log.WithFields(log.Fields{"params": request.QueryStringParameters}).Info(message)
	if val, ok := request.QueryStringParameters["data"]; ok && val == "locations" {
		return GetSkapingLocations()
	}
	return ReturnBadRequest()
}

func ReturnBadRequest() (*events.APIGatewayProxyResponse, error) {
	var data struct{}
	bodyResponse := &SkapingDataBodyResponse{"Bad Request", &data}
	bodyResponseJson, _ := json.Marshal(bodyResponse)
	return &events.APIGatewayProxyResponse{
		StatusCode:      400,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            string(bodyResponseJson),
		IsBase64Encoded: false,
	}, nil
}

func GetSkapingLocations() (*events.APIGatewayProxyResponse, error) {

	log.Info("Getting Skaping locations.")

	skapingLocations := NewSkapingLocationScraperFromUrl(SKAPING_MAP_URL).ScrapLocations()

	bodyResponse := &SkapingDataBodyResponse{"OK", &skapingLocations}

	bodyResponseJson, _ := json.Marshal(bodyResponse)

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            string(bodyResponseJson),
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}
