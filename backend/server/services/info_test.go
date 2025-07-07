package services

import (
	"backend/api"
	"backend/config"
	"backend/utils"
	"testing"
)

func TestInfoService_Info(t *testing.T) {
	tests := []struct {
		name   string
		config *config.Config
		want   *api.InfoResponseContent
	}{
		{
			name: "info with all fields populated",
			config: &config.Config{
				Region:                 "us-west-2",
				UserPoolId:             "us-west-2_123456789",
				UserPoolWebClientId:    "abcdef123456789",
				AuthenticationFlowType: "USER_PASSWORD_AUTH",
			},
			want: &api.InfoResponseContent{
				Region:                 utils.Ptr("us-west-2"),
				UserPoolId:             utils.Ptr("us-west-2_123456789"),
				UserPoolWebClientId:    utils.Ptr("abcdef123456789"),
				AuthenticationFlowType: utils.Ptr("USER_PASSWORD_AUTH"),
			},
		},
		{
			name: "info with empty fields",
			config: &config.Config{
				Region:                 "",
				UserPoolId:             "",
				UserPoolWebClientId:    "",
				AuthenticationFlowType: "",
			},
			want: &api.InfoResponseContent{
				Region:                 utils.Ptr(""),
				UserPoolId:             utils.Ptr(""),
				UserPoolWebClientId:    utils.Ptr(""),
				AuthenticationFlowType: utils.Ptr(""),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewInfoService(tt.config)
			got, err := s.Info()

			if err != nil {
				t.Errorf("InfoService.Info() error = %v, expected no error", err)
				return
			}

			if got == nil {
				t.Error("InfoService.Info() returned nil")
				return
			}

			// Check each field
			if !utils.StringEqual(got.Region, tt.want.Region) {
				t.Errorf("InfoService.Info() Region = %v, want %v", *got.Region, *tt.want.Region)
			}

			if !utils.StringEqual(got.UserPoolId, tt.want.UserPoolId) {
				t.Errorf("InfoService.Info() UserPoolId = %v, want %v", *got.UserPoolId, *tt.want.UserPoolId)
			}

			if !utils.StringEqual(got.UserPoolWebClientId, tt.want.UserPoolWebClientId) {
				t.Errorf("InfoService.Info() UserPoolWebClientId = %v, want %v", *got.UserPoolWebClientId, *tt.want.UserPoolWebClientId)
			}

			if !utils.StringEqual(got.AuthenticationFlowType, tt.want.AuthenticationFlowType) {
				t.Errorf("InfoService.Info() AuthenticationFlowType = %v, want %v", *got.AuthenticationFlowType, *tt.want.AuthenticationFlowType)
			}
		})
	}
}
