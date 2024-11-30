package app

import (
	"golang_template/handler/controllers"
	"golang_template/handler/routers"

	"github.com/gofiber/fiber/v2"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func (a *application) InitRouter(app *fiber.App, controller controllers.Controllers, tracer oteltrace.Tracer) routers.Router {
	router := routers.NewRouter(controller, tracer)
	router.AddRoutes(app.Group(""))
	return router
}
