package app

import (
	"github.com/gofiber/fiber/v2"
	"golang_template/handler/controllers"
	"golang_template/handler/middlewares"
	"golang_template/handler/routers"
)

func (a *application) InitRouter(app *fiber.App, controller controllers.Controllers) routers.Router {
	// Use the tracing middleware globally
	app.Use(middlewares.TracingMiddleware())

	router := routers.NewRouter(controller)
	router.AddRoutes(app.Group(""))
	return router
}
