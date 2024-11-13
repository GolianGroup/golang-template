package app

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.4.0"
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

func (a *application) InitJaegerTracer() func() {
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(a.config.Jaeger.Endpoint)))
	if err != nil {
		log.Fatalf("failed to create Jaeger exporter: %v", err)
	}

	// Set up the tracer provider
	tracerProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(a.config.Jaeger.Service),
		)),
	)
	otel.SetTracerProvider(tracerProvider)

	log.Println("Jaeger tracer initialized")

	// Return a cleanup function to shut down the tracer provider
	return func() {
		if err := tracerProvider.Shutdown(a.ctx); err != nil {
			log.Printf("failed to shut down tracer provider: %v", err)
		}
	}
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
