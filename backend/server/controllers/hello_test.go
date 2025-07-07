package controllers

import (
	"backend/api"
	"context"
	"errors"
	"testing"
)

// Mock HelloService for testing
type mockHelloService struct {
	sayHelloFunc func(input string) (*api.SayHelloResponseContent, error)
}

func (m *mockHelloService) SayHello(ctx context.Context, input string) (*api.SayHelloResponseContent, error) {
	if m.sayHelloFunc != nil {
		return m.sayHelloFunc(input)
	}
	return &api.SayHelloResponseContent{Message: "Hello, World!"}, nil
}

func TestHelloController_SayHello(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		mockFunc    func(input string) (*api.SayHelloResponseContent, error)
		want        *api.SayHelloResponseContent
		wantErr     bool
		expectedErr error
	}{
		{
			name:  "successful hello",
			input: "World",
			mockFunc: func(input string) (*api.SayHelloResponseContent, error) {
				return &api.SayHelloResponseContent{Message: "World"}, nil
			},
			want:    &api.SayHelloResponseContent{Message: "World"},
			wantErr: false,
		},
		{
			name:  "successful hello with empty input",
			input: "",
			mockFunc: func(input string) (*api.SayHelloResponseContent, error) {
				return &api.SayHelloResponseContent{Message: ""}, nil
			},
			want:    &api.SayHelloResponseContent{Message: ""},
			wantErr: false,
		},
		{
			name:  "service returns invalid input error",
			input: "400",
			mockFunc: func(input string) (*api.SayHelloResponseContent, error) {
				return nil, api.InvalidInputErrorResponseContent{ErrorMessage: "Invalid input"}
			},
			want:        nil,
			wantErr:     true,
			expectedErr: api.InvalidInputErrorResponseContent{ErrorMessage: "Invalid input"},
		},
		{
			name:  "service returns internal server error",
			input: "500",
			mockFunc: func(input string) (*api.SayHelloResponseContent, error) {
				return nil, api.InternalServerErrorResponseContent{ErrorMessage: "Internal server error"}
			},
			want:        nil,
			wantErr:     true,
			expectedErr: api.InternalServerErrorResponseContent{ErrorMessage: "Internal server error"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &mockHelloService{
				sayHelloFunc: tt.mockFunc,
			}

			controller := NewHelloController(mockService)
			got, err := controller.SayHello(context.Background(), tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("HelloController.SayHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				if err == nil {
					t.Errorf("HelloController.SayHello() expected error but got nil")
					return
				}
				if err.Error() != tt.expectedErr.Error() {
					t.Errorf("HelloController.SayHello() error = %v, expected error %v", err.Error(), tt.expectedErr)
				}
				// Check that the error type matches the expected error type
				var expectedInvalidInputErr api.InvalidInputErrorResponseContent
				var expectedInternalServerErr api.InternalServerErrorResponseContent

				if errors.As(tt.expectedErr, &expectedInvalidInputErr) {
					var actualInvalidInputErr api.InvalidInputErrorResponseContent
					if !errors.As(err, &actualInvalidInputErr) {
						t.Errorf("HelloService.SayHello() error type mismatch: got %T, want %T", err, tt.expectedErr)
					}
				} else if errors.As(tt.expectedErr, &expectedInternalServerErr) {
					var actualInternalServerErr api.InternalServerErrorResponseContent
					if !errors.As(err, &actualInternalServerErr) {
						t.Errorf("HelloService.SayHello() error type mismatch: got %T, want %T", err, tt.expectedErr)
					}
				} else {
					t.Errorf("HelloService.SayHello() error type mismatch: got %T, want %T", err, tt.expectedErr)
				}
				return
			}

			if got == nil && tt.want != nil {
				t.Errorf("HelloController.SayHello() got nil, want %v", tt.want)
				return
			}

			if got != nil && tt.want == nil {
				t.Errorf("HelloController.SayHello() got %v, want nil", got)
				return
			}

			if got != nil && tt.want != nil && got.Message != tt.want.Message {
				t.Errorf("HelloController.SayHello() got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHelloController(t *testing.T) {
	mockService := &mockHelloService{}
	controller := NewHelloController(mockService)

	if controller == nil {
		t.Error("NewHelloController() returned nil")
		return
	}

	if controller.helloService != mockService {
		t.Error("NewHelloController() did not set service properly")
	}
}

func TestHelloController_Interface(t *testing.T) {
	var _ HelloController = (*helloController)(nil)
}

func TestHelloController_ServiceIntegration(t *testing.T) {
	// Test that controller correctly passes input to service
	var receivedInput string
	mockService := &mockHelloService{
		sayHelloFunc: func(input string) (*api.SayHelloResponseContent, error) {
			receivedInput = input
			return &api.SayHelloResponseContent{Message: input}, nil
		},
	}

	controller := NewHelloController(mockService)
	expectedInput := "test-input"

	_, err := controller.SayHello(context.Background(), expectedInput)
	if err != nil {
		t.Errorf("HelloController.SayHello() unexpected error = %v", err)
	}

	if receivedInput != expectedInput {
		t.Errorf("HelloController.SayHello() passed input %v to service, want %v", receivedInput, expectedInput)
	}
}

func TestHelloController_NilService(t *testing.T) {
	// This test ensures the controller handles nil service gracefully
	// In a real scenario, this might cause a panic, but it's good to document the behavior
	defer func() {
		if r := recover(); r == nil {
			t.Error("HelloController with nil service should panic")
		}
	}()

	controller := NewHelloController(nil)
	_, _ = controller.SayHello(context.Background(), "test")
}
