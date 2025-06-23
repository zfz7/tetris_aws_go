package main

import (
	"backend/api"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func Test_sayHello(t *testing.T) {
	got := sayHello("hi")
	want := api.SayHelloResponseContent{Message: "hi"}
	if got != want {
		t.Errorf("sayHello(hi) = %v; want %v", got, want)
	}
}

func Test_handleInfo(t *testing.T) {
	t.Setenv("REGION", "1")
	t.Setenv("USER_POOL_ID", "2")
	t.Setenv("USER_POOL_WEB_CLIENT_ID", "3")
	got := handleInfo()
	want := api.InfoResponseContent{
		Region:                 Ptr("1"),
		UserPoolId:             Ptr("2"),
		UserPoolWebClientId:    Ptr("3"),
		AuthenticationFlowType: Ptr("USER_PASSWORD_AUTH"),
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("handleInfo() = %#v; want %#v", got, want)
	}
}
func Test_routerForSayHello(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		QueryStringParameters: map[string]string{"name": "hi"},
		RequestContext: events.APIGatewayProxyRequestContext{
			OperationName: "SayHello",
		},
	}
	responseBody, _ := json.Marshal(api.SayHelloResponseContent{
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
	responseBody, _ := json.Marshal(api.InfoResponseContent{
		Region:                 Ptr("1"),
		UserPoolId:             Ptr("2"),
		UserPoolWebClientId:    Ptr("3"),
		AuthenticationFlowType: Ptr("USER_PASSWORD_AUTH"),
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
	responseBody, _ := json.Marshal(api.ApiErrorResponseContent{
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
