package routers

import (
	"golang_template/handler/controllers"

	"github.com/gofiber/fiber/v2"
)

type UserRouter interface {
	AddRoutes(router fiber.Router)
}

type userRouter struct {
	Controller controllers.UserController
}

func NewUserRouter(userController controllers.UserController) UserRouter {
	return &userRouter{Controller: userController}
}

func (r userRouter) AddRoutes(router fiber.Router) {
	// init routes for user
	// has controller
	router.Get("/user", r.Controller.Login)
}
