package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
)

var SKAPING_MAP_URL = os.Getenv("SKAPING_MAP_URL")

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

func ReturnInternalServerError(err error) (*events.APIGatewayProxyResponse, error) {
	data := struct {
		Error string `json:"error"`
	}{err.Error()}
	bodyResponse := &SkapingDataBodyResponse{"Internal Server Error", &data}
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

	scraper, err := NewSkapingLocationScraperFromUrl(SKAPING_MAP_URL)
	if err != nil {
		return ReturnInternalServerError(err)
	}

	skapingLocations := scraper.ScrapLocations()

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
