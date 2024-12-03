package services

import (
	"context"
	"golang_template/grpc/gen/example"
)

// ExampleService implements the ExampleServer interface
type ExampleService struct {
	example.UnimplementedExampleServer
}

// NewExampleService creates a new instance of ExampleService
func NewExampleService() *ExampleService {
	return &ExampleService{}
}

// SayHello implements the SayHello method of the ExampleServer interface
func (s *ExampleService) SayHello(ctx context.Context, req *example.HelloRequest) (*example.HelloReply, error) {
	// Implement your logic here
	message := "Hello, " + req.GetName() // Assuming HelloRequest has a field 'Name'
	return &example.HelloReply{Message: message}, nil
}
