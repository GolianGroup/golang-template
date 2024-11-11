package api

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Router struct {
	fx.In
	Routers []IRouter `group:"routers"`
}
type IRouter interface {
	Handle(fiberRouter fiber.Router)
}
