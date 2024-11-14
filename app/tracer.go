package app

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.4.0"
	"log"
)

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
