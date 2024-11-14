package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
)

func TracingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tracer := otel.Tracer("myAppTracer")
		ctx, span := tracer.Start(c.Context(), c.Path())
		defer span.End()

		c.SetUserContext(ctx)
		return c.Next()
	}
}
