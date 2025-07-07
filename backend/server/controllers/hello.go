package controllers

import (
	"backend/api"
	"backend/server/services"
	"context"
)

type HelloController interface {
	SayHello(ctx context.Context, input string) (*api.SayHelloResponseContent, error)
}

type helloController struct {
	helloService services.HelloService
}

func NewHelloController(helloService services.HelloService) *helloController {
	return &helloController{
		helloService: helloService,
	}
}

func (c *helloController) SayHello(ctx context.Context, input string) (*api.SayHelloResponseContent, error) {
	return c.helloService.SayHello(ctx, input)
}
