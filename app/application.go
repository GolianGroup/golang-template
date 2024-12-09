package app

import (
	"context"
	"golang_template/handler/routers"
	"golang_template/internal/config"
	"golang_template/internal/database/postgres"
	"golang_template/internal/logging"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
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
	)
	app.Run()
}
