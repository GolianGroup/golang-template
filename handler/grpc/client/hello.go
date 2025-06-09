package grpc

import (
	"fmt"
	"golang_template/internal/config"
	hello_rpc_service "golang_template/proto"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type HelloClient interface {
	Connect() error
	Client() hello_rpc_service.RpcServiceClient
	Close() error
}
type helloClient struct {
	conn   *grpc.ClientConn
	conf   *config.Config
	logger *zap.Logger
}

// NewClient creates a new gRPC client with the given configuration
func NewHelloClient(conf *config.Config, logger *zap.Logger) HelloClient {
	return &helloClient{
		conf:   conf,
		logger: logger,
	}
}

// Close closes the gRPC connection
func (c *helloClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *helloClient) Connect() error {
	var retryOpts = []grpc_retry.CallOption{
		grpc_retry.WithMax(uint(c.conf.GRPCclient.MaxRetries)),                                                                        // Maximum of 5 retry attempts
		grpc_retry.WithPerRetryTimeout(time.Duration(c.conf.GRPCclient.PerRetryTimeout) * time.Second),                                // Timeout per retry attempt
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(time.Duration(c.conf.GRPCclient.BackoffExponential) * time.Millisecond)), // Exponential backoff
		// grpc_retry.BackoffLinear(),
		grpc_retry.WithCodes(codes.Unavailable, codes.DeadlineExceeded),
		grpc_retry.Disable(),
	}

	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOpts...)),
		grpc.WithDefaultServiceConfig(`{
			"methodConfig": [{
				"name": [{"service": ""}],
				"waitForReady" : true
			}]
		}`),
		grpc.WithDisableRetry(),
		grpc.WithDisableHealthCheck(),
		grpc.WithMaxCallAttempts(10),
		grpc.WithWriteBufferSize(1024 * 1024),
		grpc.WithReadBufferSize(1024 * 1024),
		grpc.WithConnectParams(grpc.ConnectParams{}),
	}

	url := fmt.Sprintf("%s:%s", c.conf.GRPCclient.Host, c.conf.GRPCclient.Port)
	conn, err := grpc.NewClient(url, options...)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.Unavailable:
				c.logger.Fatal("gRPC service is unavailable", zap.Error(err))
			case codes.DeadlineExceeded:
				c.logger.Error("gRPC request deadline exceeded", zap.Error(err))
			case codes.Internal:
				c.logger.Error("Internal server error during gRPC call", zap.Error(err))
			default:
				c.logger.Fatal("Unexpected gRPC error", zap.String("error_code", s.Code().String()), zap.Error(err))
			}
		} else {
			c.logger.Fatal("Failed to create gRPC client due to non-gRPC error", zap.Error(err))
		}

		return err
	}

	c.conn = conn

	return nil
}

func (c *helloClient) Client() hello_rpc_service.RpcServiceClient {
	return hello_rpc_service.NewRpcServiceClient(c.conn)
}
