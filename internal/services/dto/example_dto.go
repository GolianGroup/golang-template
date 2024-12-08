package dto

import (
	example "golang_template/grpc/gen/example/proto"
)

// HelloRequestDTO represents the domain model for hello request
type HelloRequestDTO struct {
	Name string
}

// HelloReplyDTO represents the domain model for hello response
type HelloReplyDTO struct {
	Message string
}

// ToHelloRequestDTO converts gRPC HelloRequest to domain DTO
func ToHelloRequestDTO(req *example.HelloRequest) *HelloRequestDTO {
	return &HelloRequestDTO{
		Name: req.Name,
	}
}

// ToHelloReply converts domain DTO to gRPC HelloReply
func (dto *HelloReplyDTO) ToHelloReply() *example.HelloReply {
	return &example.HelloReply{
		Message: dto.Message,
	}
}
