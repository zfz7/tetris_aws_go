package services

import (
	"backend/api"
)

type HelloService interface {
	SayHello(input string) (*api.SayHelloResponseContent, error)
}

type helloService struct {
}

func NewHelloService() *helloService {
	return &helloService{}
}

func (s *helloService) SayHello(input string) (*api.SayHelloResponseContent, error) {
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
