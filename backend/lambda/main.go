package main

import (
	"backend/api"
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(router)
}

var corsHeaders = map[string]string{
	"Access-Control-Allow-Origin":      "*",
	"Access-Control-Allow-Headers":     "*",
	"Access-Control-Allow-Methods":     "OPTIONS, POST, GET",
	"Access-Control-Allow-Credentials": "true",
}
var apiError, _ = json.Marshal(api.ApiErrorResponseContent{
	ErrorMessage: "Not Found",
})
var errorEvent = events.APIGatewayProxyResponse{
	Headers:    corsHeaders,
	StatusCode: http.StatusNotFound,
	Body:       string(apiError),
}

func router(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body []byte
	var err error
	switch req.RequestContext.OperationName {
	case "SayHello":
		body, err = json.Marshal(sayHello(req.QueryStringParameters["name"]))
	case "Info":
		body, err = json.Marshal(handleInfo())
	default:
		return errorEvent, nil
	}
	if err != nil {
		return errorEvent, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    corsHeaders,
		Body:       string(body),
	}, nil
}

func sayHello(input string) api.SayHelloResponseContent {
	return api.SayHelloResponseContent{Message: input}
}

func handleInfo() api.InfoResponseContent {
	return api.InfoResponseContent{
		Region:                 Ptr(os.Getenv("REGION")),
		UserPoolId:             Ptr(os.Getenv("USER_POOL_ID")),
		UserPoolWebClientId:    Ptr(os.Getenv("USER_POOL_WEB_CLIENT_ID")),
		AuthenticationFlowType: Ptr("USER_PASSWORD_AUTH"),
	}
}

func Ptr[T any](v T) *T {
	return &v
}
