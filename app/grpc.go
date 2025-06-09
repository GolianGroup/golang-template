package app

import (
	"golang_template/handler/controllers"
	"golang_template/internal/logging"
	rpc_service "golang_template/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpc_client "golang_template/handler/grpc/client"
)

func (a *application) InitGRPCServer(controller controllers.Controllers, logger logging.Logger) *grpc.Server {
	grpcServer := grpc.NewServer()
	// Register server with controller
	rpc_service.RegisterRpcServiceServer(grpcServer, controller.RpcServiceController())

	if a.config.Environment == "development" {
		reflection.Register(grpcServer)
	}

	return grpcServer
}

func (a *application) InitGRPCClient(logger *zap.Logger) grpc_client.HelloClient {
	return grpc_client.NewHelloClient(a.config, logger)
}
