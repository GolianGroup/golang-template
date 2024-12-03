package app

import (
	"golang_template/internal/grpc_services"

	"google.golang.org/grpc"
)

func (a *application) InitGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	userServiceServer := grpc_services.NewUserServiceServer(a.InitController().UserController())
	grpc_services.RegisterUserServiceServer(grpcServer, userServiceServer)
	return grpcServer
}
