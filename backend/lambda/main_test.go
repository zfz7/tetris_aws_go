package main

import (
	"backend/api"
	"backend/config"
	"backend/server/controllers"
	"backend/server/services"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestRouter_SayHello(t *testing.T) {
	// Initialize controllers for testing
	cfg := &config.Config{
		Region:                 "us-west-2",
		UserPoolId:             "test-pool-id",
		UserPoolWebClientId:    "test-client-id",
		AuthenticationFlowType: "USER_PASSWORD_AUTH",
	}
	svc = services.InitServices(cfg)
	ctrl = controllers.InitControllers(svc, cfg)

	tests := []struct {
		name           string
		request        events.APIGatewayProxyRequest
		expectedStatus int
		expectedBody   string
		wantErr        bool
	}{
		{
			name: "successful SayHello request",
			request: events.APIGatewayProxyRequest{
				QueryStringParameters: map[string]string{"name": "World"},
				RequestContext: events.APIGatewayProxyRequestContext{
					OperationName: "SayHello",
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"World"}`,
			wantErr:        false,
		},
		{
			name: "SayHello with empty name",
			request: events.APIGatewayProxyRequest{
				QueryStringParameters: map[string]string{"name": ""},
				RequestContext: events.APIGatewayProxyRequestContext{
					OperationName: "SayHello",
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":""}`,
			wantErr:        false,
		},
		{
			name: "SayHello with no query parameters",
			request: events.APIGatewayProxyRequest{
				QueryStringParameters: nil,
				RequestContext: events.APIGatewayProxyRequestContext{
					OperationName: "SayHello",
				},
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":""}`,
			wantErr:        false,
		},
		{
			name: "SayHello with invalid input triggering 400 error",
			request: events.APIGatewayProxyRequest{
				QueryStringParameters: map[string]string{"name": "400"},
				RequestContext: events.APIGatewayProxyRequestContext{
					OperationName: "SayHello",
				},
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"errorMessage":"Invalid Input."}`,
			wantErr:        false,
		},
		{
			name: "SayHello with invalid input triggering 500 error",
			request: events.APIGatewayProxyRequest{
				QueryStringParameters: map[string]string{"name": "500"},
				RequestContext: events.APIGatewayProxyRequestContext{
					OperationName: "SayHello",
				},
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"errorMessage":"Internal Server Error."}`,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := router(context.Background(), tt.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("router() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.StatusCode != tt.expectedStatus {
				t.Errorf("router() StatusCode = %v, want %v", got.StatusCode, tt.expectedStatus)
			}

			if got.Body != tt.expectedBody {
				t.Errorf("router() Body = %v, want %v", got.Body, tt.expectedBody)
			}

			// Verify CORS headers are set
			if !reflect.DeepEqual(got.Headers, corsHeaders) {
				t.Errorf("router() Headers = %v, want %v", got.Headers, corsHeaders)
			}
		})
	}
}

func TestRouter_Info(t *testing.T) {
	// Initialize controllers for testing
	cfg := &config.Config{
		Region:                 "us-west-2",
		UserPoolId:             "test-pool-id",
		UserPoolWebClientId:    "test-client-id",
		AuthenticationFlowType: "USER_PASSWORD_AUTH",
	}
	svc = services.InitServices(cfg)
	ctrl = controllers.InitControllers(svc, cfg)

	tests := []struct {
		name           string
		request        events.APIGatewayProxyRequest
		expectedStatus int
		wantErr        bool
	}{
		{
			name: "successful Info request",
			request: events.APIGatewayProxyRequest{
				RequestContext: events.APIGatewayProxyRequestContext{
					OperationName: "Info",
				},
			},
			expectedStatus: http.StatusOK,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := router(context.Background(), tt.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("router() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.StatusCode != tt.expectedStatus {
				t.Errorf("router() StatusCode = %v, want %v", got.StatusCode, tt.expectedStatus)
			}

			// Verify the response body contains expected info structure
			var infoResponse api.InfoResponseContent
			if err := json.Unmarshal([]byte(got.Body), &infoResponse); err != nil {
				t.Errorf("router() Body is not valid JSON: %v", err)
			}

			// Verify CORS headers are set
			if !reflect.DeepEqual(got.Headers, corsHeaders) {
				t.Errorf("router() Headers = %v, want %v", got.Headers, corsHeaders)
			}
		})
	}
}

func TestErrorHandler(t *testing.T) {
	tests := []struct {
		name           string
		err            error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "InvalidInputErrorResponseContent",
			err:            api.InvalidInputErrorResponseContent{ErrorMessage: "Invalid input"},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"errorMessage":"Invalid input"}`,
		},
		{
			name:           "InternalServerErrorResponseContent",
			err:            api.InternalServerErrorResponseContent{ErrorMessage: "Internal server error"},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"errorMessage":"Internal server error"}`,
		},
		{
			name:           "generic error",
			err:            errors.New("generic error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := errorHandler(tt.err)

			if got.StatusCode != tt.expectedStatus {
				t.Errorf("errorHandler() StatusCode = %v, want %v", got.StatusCode, tt.expectedStatus)
			}

			if got.Body != tt.expectedBody {
				t.Errorf("errorHandler() Body = %v, want %v", got.Body, tt.expectedBody)
			}

			// Verify CORS headers are set
			if !reflect.DeepEqual(got.Headers, corsHeaders) {
				t.Errorf("errorHandler() Headers = %v, want %v", got.Headers, corsHeaders)
			}
		})
	}
}

func TestCorsHeaders(t *testing.T) {
	expected := map[string]string{
		"Access-Control-Allow-Origin":      "*",
		"Access-Control-Allow-Headers":     "*",
		"Access-Control-Allow-Methods":     "OPTIONS, POST, GET",
		"Access-Control-Allow-Credentials": "true",
	}

	if !reflect.DeepEqual(corsHeaders, expected) {
		t.Errorf("corsHeaders = %v, want %v", corsHeaders, expected)
	}
}
