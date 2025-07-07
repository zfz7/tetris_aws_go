package services

import (
	"backend/api"
	"backend/config"
	"backend/utils"
)

type InfoService interface {
	Info() (*api.InfoResponseContent, error)
}

type infoService struct {
	cfg *config.Config
}

func NewInfoService(cfg *config.Config) *infoService {
	return &infoService{
		cfg: cfg,
	}
}

func (s *infoService) Info() (*api.InfoResponseContent, error) {
	if s.cfg == nil {
		return &api.InfoResponseContent{
			AuthenticationFlowType: utils.Ptr(""),
			Region:                 utils.Ptr(""),
			UserPoolId:             utils.Ptr(""),
			UserPoolWebClientId:    utils.Ptr(""),
		}, nil
	}
	return &api.InfoResponseContent{
		AuthenticationFlowType: utils.Ptr(s.cfg.AuthenticationFlowType),
		Region:                 utils.Ptr(s.cfg.Region),
		UserPoolId:             utils.Ptr(s.cfg.UserPoolId),
		UserPoolWebClientId:    utils.Ptr(s.cfg.UserPoolWebClientId),
	}, nil
}
