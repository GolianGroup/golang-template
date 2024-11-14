package app

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"golang_template/handler/routers"
	"golang_template/internal/config"
	"golang_template/internal/database"
	"log"
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
			a.InitDatabase,
			a.InitJaegerTracer,
		),
		fx.Invoke(func(lc fx.Lifecycle, db database.Database, cleanup func()) {
			lc.Append(fx.Hook{
				OnStop: func(ctx context.Context) error {
					log.Println("Shutting down tracing provider")
					cleanup() // Clean up Jaeger Tracer
					log.Println(db.Close())
					return nil
				},
			})
		}),
		fx.Invoke(func(app *fiber.App, router routers.Router) {
			log.Fatal(app.Listen(a.config.Server.Host + ":" + a.config.Server.Port))
		}),
	)
	app.Run()
}
