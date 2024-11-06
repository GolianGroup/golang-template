package app

import (
	"github.com/gofiber/fiber/v3"
	"golang_template/handler/controllers"
	"golang_template/handler/routers"
)

func (a *application) InitRouter(app *fiber.App, controller controllers.UserController) routers.Router {
	return routers.NewRouter(app, controller)
}
