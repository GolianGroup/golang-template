package routers

import (
	"github.com/gofiber/fiber/v3"
	"golang_template/handler/controllers"
)

type Router interface {
	Handle()
}

type router struct {
	app        *fiber.App
	controller controllers.UserController
}

func NewRouter(app *fiber.App, controller controllers.UserController) Router {
	return &router{app: app, controller: controller}
}

func (r router) Handle() {
	// router
	// init user router, etc ...
	// rate limiter
	// CORS
	//
	uRouter := NewUserRouter(r.app, r.controller)
	uRouter.Handle()

}
