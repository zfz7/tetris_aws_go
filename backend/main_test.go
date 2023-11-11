package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"reflect"
	"testing"
)

func Test_sayHello(t *testing.T) {
	got := sayHello("hi")
	want := SayHelloResponse{Message: "hi"}
	if got != want {
		t.Errorf("sayHello(hi) = %v; want %v", got, want)
	}
}

func Test_handelInfo(t *testing.T) {
	t.Setenv("REGION", "1")
	t.Setenv("USER_POOL_ID", "2")
	t.Setenv("USER_POOL_WEB_CLIENT_ID", "3")
	got := handelInfo()
	want := InfoResponse{
		Region:                 "1",
		UserPoolId:             "2",
		UserPoolWebClientId:    "3",
		AuthenticationFlowType: "USER_PASSWORD_AUTH",
	}
	if got != want {
		t.Errorf("handelInfo() = %v; want %v", got, want)
	}
}
func Test_routerForSayHello(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{"name": "hi"},
		RequestContext: events.APIGatewayProxyRequestContext{
			OperationName: "SayHello",
		},
	}
	responseBody, _ := json.Marshal(SayHelloResponse{
		Message: "hi",
	})
	want := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Headers":     "*",
			"Access-Control-Allow-Methods":     "OPTIONS, POST, GET",
			"Access-Control-Allow-Credentials": "true",
		},
		Body: string(responseBody),
	}
	got, _ := router(nil, request)
	if got.Body != want.Body {
		t.Errorf("router Body = %v; want %v", got.Body, want.Body)
	}
	if got.StatusCode != want.StatusCode {
		t.Errorf("router StatusCode = %v; want %v", got.StatusCode, want.StatusCode)
	}
	if !reflect.DeepEqual(got.Headers, want.Headers) {
		t.Errorf("router Headers = %v; want %v", got.Headers, want.Headers)
	}
}
func Test_routerForInfo(t *testing.T) {
	t.Setenv("REGION", "1")
	t.Setenv("USER_POOL_ID", "2")
	t.Setenv("USER_POOL_WEB_CLIENT_ID", "3")
	request := events.APIGatewayProxyRequest{
		RequestContext: events.APIGatewayProxyRequestContext{
			OperationName: "Info",
		},
	}
	responseBody, _ := json.Marshal(InfoResponse{
		Region:                 "1",
		UserPoolId:             "2",
		UserPoolWebClientId:    "3",
		AuthenticationFlowType: "USER_PASSWORD_AUTH",
	})
	want := events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Headers":     "*",
			"Access-Control-Allow-Methods":     "OPTIONS, POST, GET",
			"Access-Control-Allow-Credentials": "true",
		},
		Body: string(responseBody),
	}
	got, _ := router(nil, request)
	if got.Body != want.Body {
		t.Errorf("router Body = %v; want %v", got.Body, want.Body)
	}
	if got.StatusCode != want.StatusCode {
		t.Errorf("router StatusCode = %v; want %v", got.StatusCode, want.StatusCode)
	}
	if !reflect.DeepEqual(got.Headers, want.Headers) {
		t.Errorf("router Headers = %v; want %v", got.Headers, want.Headers)
	}
}

func Test_routerForError(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		RequestContext: events.APIGatewayProxyRequestContext{
			OperationName: "NotFound",
		},
	}
	responseBody, _ := json.Marshal(ApiError{
		ErrorMessage: "Not Found",
	})
	want := events.APIGatewayProxyResponse{
		StatusCode: http.StatusNotFound,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Headers":     "*",
			"Access-Control-Allow-Methods":     "OPTIONS, POST, GET",
			"Access-Control-Allow-Credentials": "true",
		},
		Body: string(responseBody),
	}
	got, _ := router(nil, request)
	if got.Body != want.Body {
		t.Errorf("router Body = %v; want %v", got.Body, want.Body)
	}
	if got.StatusCode != want.StatusCode {
		t.Errorf("router StatusCode = %v; want %v", got.StatusCode, want.StatusCode)
	}
	if !reflect.DeepEqual(got.Headers, want.Headers) {
		t.Errorf("router Headers = %v; want %v", got.Headers, want.Headers)
	}
}
