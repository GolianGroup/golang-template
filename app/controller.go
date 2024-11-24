package app

import (
	"golang_template/handler/controllers"
	"golang_template/internal/logging"

	oteltrace "go.opentelemetry.io/otel/trace"
)

func (a *application) InitController(logger logging.Logger, tracer oteltrace.Tracer) controllers.Controllers {
	return controllers.NewControllers(logger, tracer)
}
