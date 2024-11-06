package routers

import (
	"github.com/gofiber/fiber/v3"
	"golang_template/handler/controllers"
)

type UserRouter interface {
	Handle()
}

type userRouter struct {
	app        *fiber.App
	controller controllers.UserController
}

func NewUserRouter(app *fiber.App, userController controllers.UserController) UserRouter {
	return &userRouter{app: app, controller: userController}
}

func (r userRouter) Handle() {
	//init routes for user
	// has controller
	r.app.Get("/user", r.controller.Login)
}
