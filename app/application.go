package app

import (
	"context"
	"golang_template/handler/routers"
	"golang_template/internal/config"
	"golang_template/internal/database/postgres"
	"golang_template/internal/logging"
	"log"
	"net"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type Application interface {
	Setup()
}

type application struct {
	ctx    context.Context
	config *config.Config
}

func NewApplication(ctx context.Context, config *config.Config) Application {
	return &application{ctx: ctx, config: config}
}

// bootstrap

func (a *application) Setup() {
	app := fx.New(
		fx.Provide(
			a.InitRouter,
			a.InitFramework,
			a.InitController,
			a.InitServices,
			a.InitRepositories,
			a.InitRedis,
			a.InitDatabase,
			a.InitArangoDB,
			a.InitLogger,
			a.InitTracerProvider,
			a.InitGRPCServer,
		),
		fx.Invoke(func(lc fx.Lifecycle, db postgres.Database) {
			// Init Tracer
			shutdownTracer := a.InitTracer()
			lc.Append(fx.Hook{
				OnStop: func(ctx context.Context) error {
					if err := shutdownTracer(ctx); err != nil {
						log.Printf("Error shutting down tracer: %v", err) // this should change after logging branch get merged
					}
					log.Println(db.Close())
					return nil
				},
			})
		}),

		fx.Invoke(func(app *fiber.App, router routers.Router, logger logging.Logger) {
			logger.Info("Server Started")
			log.Fatal(app.Listen(a.config.Server.Host + ":" + a.config.Server.Port))
		}),

		fx.Invoke(func(lc fx.Lifecycle, grpcServer *grpc.Server) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					listener, err := net.Listen("tcp", a.config.GRPC.Host+":"+a.config.GRPC.Port)
					if err != nil {
						return err
					}
					go func() {
						if err := grpcServer.Serve(listener); err != nil {
							log.Fatalf("failed to serve gRPC: %v", err)
						}
					}()
					log.Println("gRPC server started on", a.config.GRPC.Host+":"+a.config.GRPC.Port)
					return nil
				},
				OnStop: func(ctx context.Context) error {
					grpcServer.Stop()
					log.Println("gRPC server stopped")
					return nil
				},
			})
		}),
	)
	app.Run()
}
