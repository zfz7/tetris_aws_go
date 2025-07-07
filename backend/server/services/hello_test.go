package services

import (
	"backend/api"
	"errors"
	"testing"
)

func TestHelloService_SayHello(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		want        *api.SayHelloResponseContent
		wantErr     bool
		expectedErr error
	}{
		{
			name:    "successful hello with normal input",
			input:   "World",
			want:    &api.SayHelloResponseContent{Message: "World"},
			wantErr: false,
		},
		{
			name:    "successful hello with empty input",
			input:   "",
			want:    &api.SayHelloResponseContent{Message: ""},
			wantErr: false,
		},
		{
			name:    "successful hello with special characters",
			input:   "Hello, 世界!",
			want:    &api.SayHelloResponseContent{Message: "Hello, 世界!"},
			wantErr: false,
		},
		{
			name:        "invalid input error - 400",
			input:       "400",
			want:        nil,
			wantErr:     true,
			expectedErr: api.InvalidInputErrorResponseContent{ErrorMessage: "Invalid Input."},
		},
		{
			name:        "internal server error - 500",
			input:       "500",
			want:        nil,
			wantErr:     true,
			expectedErr: api.InternalServerErrorResponseContent{ErrorMessage: "Internal Server Error."},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewHelloService()
			got, err := s.SayHello(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("HelloService.SayHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				if err == nil {
					t.Errorf("HelloService.SayHello() expected error but got nil")
					return
				}
				if err.Error() != tt.expectedErr.Error() {
					t.Errorf("HelloService.SayHello() error = %v, expected error %v", err.Error(), tt.expectedErr)
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
				t.Errorf("HelloService.SayHello() got nil, want %v", tt.want)
				return
			}

			if got != nil && tt.want == nil {
				t.Errorf("HelloService.SayHello() got %v, want nil", got)
				return
			}

			if got != nil && tt.want != nil && got.Message != tt.want.Message {
				t.Errorf("HelloService.SayHello() got %v, want %v", got, tt.want)
			}
		})
	}
}
