package app

import (
	example "golang_template/grpc/gen/example/proto"
	services "golang_template/internal/services/"

	"google.golang.org/grpc"
)

func (a *application) InitGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	exampleService := services.NewExampleService()
	example.RegisterExampleServer(grpcServer, exampleService)
	return grpcServer
}
