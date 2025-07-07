package controllers

import (
	"backend/api"
	"backend/server/services"
)

type HelloController interface {
	SayHello(input string) (*api.SayHelloResponseContent, error)
}

type helloController struct {
	helloService services.HelloService
}

func NewHelloController(helloService services.HelloService) *helloController {
	return &helloController{
		helloService: helloService,
	}
}

func (c *helloController) SayHello(input string) (*api.SayHelloResponseContent, error) {
	return c.helloService.SayHello(input)
}
