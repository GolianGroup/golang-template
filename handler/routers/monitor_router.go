package routers

import (
	"github.com/gofiber/fiber/v2"
	"golang_template/handler/middlewares"
)

// MonitorRouter defines the methods for setting up the monitor route
type MonitorRouter interface {
	AddRoutes(router fiber.Router)
}

// monitorRouter is the struct implementing MonitorRouter interface
type monitorRouter struct{}

// NewMonitorRouter creates a new instance of monitorRouter
func NewMonitorRouter() MonitorRouter {
	return &monitorRouter{}
}

// AddRoutes sets up the monitor route at `/metrics`
func (r monitorRouter) AddRoutes(router fiber.Router) {
	router.Get("/metrics", middlewares.MonitorMiddleware())
}
