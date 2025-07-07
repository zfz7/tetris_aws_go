package controllers

import (
	"backend/api"
	"backend/utils"
	"testing"
)

// Mock InfoService for testing
type mockInfoService struct {
	infoFunc func() (*api.InfoResponseContent, error)
}

func (m *mockInfoService) Info() (*api.InfoResponseContent, error) {
	if m.infoFunc != nil {
		return m.infoFunc()
	}
	return &api.InfoResponseContent{
		Region:                 utils.Ptr("us-west-2"),
		UserPoolId:             utils.Ptr("test-pool-id"),
		UserPoolWebClientId:    utils.Ptr("test-client-id"),
		AuthenticationFlowType: utils.Ptr("USER_PASSWORD_AUTH"),
	}, nil
}

func TestInfoController_Info(t *testing.T) {
	tests := []struct {
		name        string
		mockFunc    func() (*api.InfoResponseContent, error)
		want        *api.InfoResponseContent
		wantErr     bool
		expectedErr string
	}{
		{
			name: "successful info retrieval",
			mockFunc: func() (*api.InfoResponseContent, error) {
				return &api.InfoResponseContent{
					Region:                 utils.Ptr("us-west-2"),
					UserPoolId:             utils.Ptr("us-west-2_123456789"),
					UserPoolWebClientId:    utils.Ptr("abcdef123456789"),
					AuthenticationFlowType: utils.Ptr("USER_PASSWORD_AUTH"),
				}, nil
			},
			want: &api.InfoResponseContent{
				Region:                 utils.Ptr("us-west-2"),
				UserPoolId:             utils.Ptr("us-west-2_123456789"),
				UserPoolWebClientId:    utils.Ptr("abcdef123456789"),
				AuthenticationFlowType: utils.Ptr("USER_PASSWORD_AUTH"),
			},
			wantErr: false,
		},
		{
			name: "successful info retrieval with empty values",
			mockFunc: func() (*api.InfoResponseContent, error) {
				return &api.InfoResponseContent{
					Region:                 utils.Ptr(""),
					UserPoolId:             utils.Ptr(""),
					UserPoolWebClientId:    utils.Ptr(""),
					AuthenticationFlowType: utils.Ptr(""),
				}, nil
			},
			want: &api.InfoResponseContent{
				Region:                 utils.Ptr(""),
				UserPoolId:             utils.Ptr(""),
				UserPoolWebClientId:    utils.Ptr(""),
				AuthenticationFlowType: utils.Ptr(""),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := &mockInfoService{
				infoFunc: tt.mockFunc,
			}

			controller := NewInfoController(mockService)
			got, err := controller.Info()

			if (err != nil) != tt.wantErr {
				t.Errorf("InfoController.Info() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				if err == nil {
					t.Errorf("InfoController.Info() expected error but got nil")
					return
				}
				if err.Error() != tt.expectedErr {
					t.Errorf("InfoController.Info() error = %v, expected error %v", err.Error(), tt.expectedErr)
				}
				return
			}

			if got != nil && tt.want != nil {
				if !utils.StringEqual(got.Region, tt.want.Region) {
					t.Errorf("InfoController.Info() Region = %v, want %v", utils.Ptr(got.Region), utils.Ptr(tt.want.Region))
				}
				if !utils.StringEqual(got.UserPoolId, tt.want.UserPoolId) {
					t.Errorf("InfoController.Info() UserPoolId = %v, want %v", utils.Ptr(got.UserPoolId), utils.Ptr(tt.want.UserPoolId))
				}
				if !utils.StringEqual(got.UserPoolWebClientId, tt.want.UserPoolWebClientId) {
					t.Errorf("InfoController.Info() UserPoolWebClientId = %v, want %v", utils.Ptr(got.UserPoolWebClientId), utils.Ptr(tt.want.UserPoolWebClientId))
				}
				if !utils.StringEqual(got.AuthenticationFlowType, tt.want.AuthenticationFlowType) {
					t.Errorf("InfoController.Info() AuthenticationFlowType = %v, want %v", utils.Ptr(got.AuthenticationFlowType), utils.Ptr(tt.want.AuthenticationFlowType))
				}
			}
		})
	}
}
