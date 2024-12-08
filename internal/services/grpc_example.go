package services

import (
	"context"
	"golang_template/internal/services/dto"
)

type ExampleService interface {
	SayHello(ctx context.Context, req *dto.HelloRequestDTO) (*dto.HelloReplyDTO, error)
}

type exampleService struct{}

func NewExampleService() ExampleService {
	return &exampleService{}
}

func (s *exampleService) SayHello(ctx context.Context, req *dto.HelloRequestDTO) (*dto.HelloReplyDTO, error) {
	message := "Hello, " + req.Name
	return &dto.HelloReplyDTO{
		Message: message,
	}, nil
}
