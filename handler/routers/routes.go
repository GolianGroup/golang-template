package routers

import (
	"golang_template/handler/controllers"

	"github.com/gofiber/fiber/v2"
)

type Router interface {
	AddRoutes(router fiber.Router)
}

type router struct {
	userRouter UserRouter
}

func NewRouter(controllers controllers.Controllers) Router {
	userRouter := NewUserRouter(controllers.UserController())
	return &router{userRouter: userRouter}
}

func (r router) AddRoutes(router fiber.Router) {

	// router
	// init user router, etc ...
	// rate limiter
	// CORS
	r.userRouter.AddRoutes(router)

}
