package services

import "backend/config"

type Services struct {
	InfoService  InfoService
	HelloService HelloService
}

func InitServices(cfg *config.Config) *Services {
	return &Services{
		InfoService:  NewInfoService(cfg),
		HelloService: NewHelloService(),
	}
}
