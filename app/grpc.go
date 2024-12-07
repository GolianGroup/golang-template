package app

import (
	example "golang_template/grpc/gen/example/proto"
	"golang_template/internal/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (a *application) InitGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	exampleService := services.NewExampleService()
	example.RegisterExampleServer(grpcServer, exampleService)

	reflection.Register(grpcServer)

	return grpcServer
}
