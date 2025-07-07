package controllers

import (
	"backend/config"
	"backend/server/services"
)

type Controllers struct {
	InfoController  InfoController
	HelloController HelloController
}

func InitControllers(services *services.Services, cfg *config.Config) *Controllers {
	return &Controllers{
		InfoController:  NewInfoController(services.InfoService),
		HelloController: NewHelloController(services.HelloService),
	}
}
