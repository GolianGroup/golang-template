package app

import (
	"go.opentelemetry.io/otel"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func (a *application) InitTracerProvider() oteltrace.Tracer {
	return otel.Tracer(a.config.Tracer.ServiceName)
}
