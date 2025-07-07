package services

import (
	"backend/api"
	"context"
)

type HelloService interface {
	SayHello(ctx context.Context, input string) (*api.SayHelloResponseContent, error)
}

type helloService struct {
}

func NewHelloService() *helloService {
	return &helloService{}
}

func (s *helloService) SayHello(ctx context.Context, input string) (*api.SayHelloResponseContent, error) {
	if input == "400" {
		return nil, api.InvalidInputErrorResponseContent{ErrorMessage: "Invalid Input."}
	}
	if input == "500" {
		return nil, api.InternalServerErrorResponseContent{ErrorMessage: "Internal Server Error."}
	}
	return &api.SayHelloResponseContent{
		Message: input,
	}, nil
}
