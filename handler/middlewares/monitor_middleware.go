package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func MonitorMiddleware() fiber.Handler {
	return monitor.New(monitor.Config{Title: "MyService Metrics"})
}

//TODO: Get title from env.
