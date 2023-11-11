package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"os"
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
var apiError, _ = json.Marshal(ApiError{
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
		body, err = json.Marshal(handelInfo())
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
func sayHello(input string) SayHelloResponse {
	return SayHelloResponse{Message: input}
}

func handelInfo() InfoResponse {
	return InfoResponse{
		Region:                 os.Getenv("REGION"),
		UserPoolId:             os.Getenv("USER_POOL_ID"),
		UserPoolWebClientId:    os.Getenv("USER_POOL_WEB_CLIENT_ID"),
		AuthenticationFlowType: "USER_PASSWORD_AUTH",
	}
}

type SayHelloResponse struct {
	Message string `json:"message"`
}
type InfoResponse struct {
	AuthenticationFlowType string `json:"authenticationFlowType"`
	Region                 string `json:"region"`
	UserPoolId             string `json:"userPoolId"`
	UserPoolWebClientId    string `json:"userPoolWebClientId"`
}
type ApiError struct {
	ErrorMessage string `json:"errorMessage"`
}
