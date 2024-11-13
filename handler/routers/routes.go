package routers

import (
	"github.com/gofiber/fiber/v2"
	"golang_template/handler/controllers"
)

type Router interface {
	AddRoutes(router fiber.Router)
}

type router struct {
	userRouter    UserRouter
	monitorRouter MonitorRouter
}

func NewRouter(controllers controllers.Controllers) Router {
	userRouter := NewUserRouter(controllers.GetUserController())
	monitorRouter := NewMonitorRouter()
	return &router{userRouter: userRouter,
		monitorRouter: monitorRouter,
	}
}

func (r router) AddRoutes(router fiber.Router) {

	// router
	// init user router, etc ...
	// rate limiter
	// CORS
	r.userRouter.AddRoutes(router)
	r.monitorRouter.AddRoutes(router)

}
