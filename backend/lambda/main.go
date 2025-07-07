package main

import (
	"backend/api"
	"backend/config"
	"backend/server/controllers"
	"backend/server/services"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var cfg *config.Config
var svc *services.Services
var ctrl *controllers.Controllers

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %s", err)
	}
	svc = services.InitServices(cfg)
	ctrl = controllers.InitControllers(svc, cfg)
	lambda.Start(router)
}

var corsHeaders = map[string]string{
	"Access-Control-Allow-Origin":      "*",
	"Access-Control-Allow-Headers":     "*",
	"Access-Control-Allow-Methods":     "OPTIONS, POST, GET",
	"Access-Control-Allow-Credentials": "true",
}

var apiError, _ = json.Marshal(api.InvalidInputErrorResponseContent{
	ErrorMessage: "Not Found",
})

var unknownMethod = events.APIGatewayProxyResponse{
	Headers:    corsHeaders,
	StatusCode: http.StatusMethodNotAllowed,
	Body:       string(apiError),
}

func router(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body []byte
	switch req.RequestContext.OperationName {
	case "SayHello":
		input := req.QueryStringParameters["name"]
		res, err := ctrl.HelloController.SayHello(input)
		if err != nil {
			return errorHandler(err), nil
		}
		body, err = json.Marshal(res)
	case "Info":
		res, err := ctrl.InfoController.Info()
		if err != nil {
			return errorHandler(err), nil
		}
		body, err = json.Marshal(res)
	default:
		return unknownMethod, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    corsHeaders,
		Body:       string(body),
	}, nil
}

func errorHandler(err error) events.APIGatewayProxyResponse {
	var statusCode int
	var body []byte

	if _, ok := err.(api.InvalidInputErrorResponseContent); ok {
		statusCode = http.StatusBadRequest
	} else if _, ok := err.(api.InternalServerErrorResponseContent); ok {
		statusCode = http.StatusInternalServerError
	} else {
		statusCode = http.StatusInternalServerError
	}

	body, marshalErr := json.Marshal(err)
	if marshalErr != nil {
		body = []byte(`{"errorMessage":"Internal Server Error"}`)
		statusCode = http.StatusInternalServerError
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers:    corsHeaders,
		Body:       string(body),
	}
}
