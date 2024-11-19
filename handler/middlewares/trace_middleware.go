package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

func TracingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tracer := otel.Tracer("myAppTracer")
		ctx, span := tracer.Start(c.Context(), c.Path())
		defer span.End()

		// Add attributes to the span
		span.SetAttributes(
			attribute.String("http.method", c.Method()),
			attribute.String("http.url", c.OriginalURL()),
			attribute.String("http.client_ip", c.IP()),
			attribute.String("http.user_agent", string(c.Request().Header.UserAgent())),
		)

		// Set the user context
		c.SetUserContext(ctx)

		// Proceed to the next handler and capture the status code
		err := c.Next() // Explicitly capture the error returned by c.Next()

		statusCode := c.Response().StatusCode()
		span.SetAttributes(attribute.Int("http.status_code", statusCode))
		// Record errors if any
		if err != nil {
			span.SetStatus(codes.Error, err.Error())
			span.RecordError(err)
		} else if statusCode >= 400 {
			span.SetStatus(codes.Error, "Client or Server Error")
		} else {
			span.SetStatus(codes.Ok, "Success")
		}

		return err // Return the error so it propagates correctly
	}
}
