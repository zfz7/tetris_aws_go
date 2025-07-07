package controllers

import (
	"backend/api"
	"backend/server/services"
)

type InfoController interface {
	Info() (*api.InfoResponseContent, error)
}

type infoController struct {
	infoService services.InfoService
}

func NewInfoController(infoService services.InfoService) *infoController {
	return &infoController{
		infoService: infoService,
	}
}

func (c *infoController) Info() (*api.InfoResponseContent, error) {
	return c.infoService.Info()
}
