package controllers

import (
	"context"
	example "golang_template/grpc/gen/example/proto"
	"golang_template/internal/services"
	"golang_template/internal/services/dto"
)

type GRPCController struct {
	example.UnimplementedExampleServer
	exampleService services.ExampleService
}

func NewGRPCController(service services.ExampleService) *GRPCController {
	return &GRPCController{
		exampleService: service,
	}
}

func (c *GRPCController) SayHello(ctx context.Context, req *example.HelloRequest) (*example.HelloReply, error) {
	// Convert gRPC request to DTO
	requestDTO := dto.ToHelloRequestDTO(req)

	// Call service with DTO
	responseDTO, err := c.exampleService.SayHello(ctx, requestDTO)
	if err != nil {
		return nil, err
	}

	// Convert DTO back to gRPC response
	return responseDTO.ToHelloReply(), nil
}
