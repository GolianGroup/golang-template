package routers

import (
	"golang_template/handler/controllers"

	"github.com/gofiber/fiber/v2"
)

type IUserRouter interface {
	AddRoutes(router fiber.Router)
}

type UserRouter struct {
	Controller controllers.IUserController
}

func NewUserRouter(userController controllers.IUserController) IUserRouter {
	return &UserRouter{Controller: userController}
}

func (r UserRouter) AddRoutes(router fiber.Router) {
	// init routes for user
	// has controller
	router.Get("/user", r.Controller.Login)
}
