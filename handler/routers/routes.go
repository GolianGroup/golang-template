package routers

import (
	"golang_template/handler/controllers"

	"github.com/gofiber/fiber/v2"
)

type IRouter interface {
	AddRoutes(router fiber.Router)
}

type Router struct {
	userRouter IUserRouter
}

func NewRouter(controllers controllers.IControllers) IRouter {
	userRouter := NewUserRouter(controllers.GetUserController())
	return &Router{userRouter: userRouter}
}

func (r Router) AddRoutes(router fiber.Router) {

	// router
	// init user router, etc ...
	// rate limiter
	// CORS
	r.userRouter.AddRoutes(router)

}
