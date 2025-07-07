package main

import (
	"backend/api"
	"backend/config"
	"backend/server/controllers"
	"backend/server/services"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

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

func router(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.RequestContext.OperationName {
	case "SayHello":
		input := req.QueryStringParameters["name"]
		res, err := ctrl.HelloController.SayHello(ctx, input)
		return parseRes(res, err), nil
	case "Info":
		res, err := ctrl.InfoController.Info(ctx)
		return parseRes(res, err), nil
	default:
		return events.APIGatewayProxyResponse{
			Headers:    corsHeaders,
			StatusCode: http.StatusMethodNotAllowed,
			Body:       "{\"errorMessage\":\"Method not found.\"}",
		}, nil
	}
}

func parseRes[T any](res T, err error) events.APIGatewayProxyResponse {
	if err != nil {
		return errorHandler(err)
	}
	body, e := json.Marshal(res)

	if e != nil {
		return errorHandler(err) // Will return 500 on unknown errors
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    corsHeaders,
		Body:       string(body),
	}
}

func errorHandler(err error) events.APIGatewayProxyResponse {
	var statusCode int
	var body []byte

	var invalidInputError api.InvalidInputErrorResponseContent
	var internalServerError api.InternalServerErrorResponseContent
	if errors.As(err, &invalidInputError) {
		statusCode = http.StatusBadRequest
	} else if errors.As(err, &internalServerError) {
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
