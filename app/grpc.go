package app

import (
	example "golang_template/grpc/gen/example/proto"
	"golang_template/handler/controllers"
	"golang_template/internal/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (a *application) InitGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	// Initialize service and controller
	exampleService := services.NewExampleService()
	grpcController := controllers.NewGRPCController(exampleService)

	// Register server with controller
	example.RegisterExampleServer(grpcServer, grpcController)

	reflection.Register(grpcServer)

	return grpcServer
}
