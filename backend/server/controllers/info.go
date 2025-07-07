package controllers

import (
	"backend/api"
	"backend/server/services"
	"context"
)

type InfoController interface {
	Info(ctx context.Context) (*api.InfoResponseContent, error)
}

type infoController struct {
	infoService services.InfoService
}

func NewInfoController(infoService services.InfoService) *infoController {
	return &infoController{
		infoService: infoService,
	}
}

func (c *infoController) Info(ctx context.Context) (*api.InfoResponseContent, error) {
	return c.infoService.Info(ctx)
}
