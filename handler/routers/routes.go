package routers

import (
	"golang_template/handler/controllers"
	"golang_template/internal/producers"

	"github.com/gofiber/fiber/v2"
)

type Router interface {
	AddRoutes(router fiber.Router)
}

type router struct {
	redisClient producers.RedisClient

	userRouter UserRouter
}

func NewRouter(controllers controllers.Controllers, redisClient producers.RedisClient) Router {
	userRouter := NewUserRouter(controllers.GetUserController(), redisClient)
	return &router{userRouter: userRouter, redisClient: redisClient}
}

func (r router) AddRoutes(router fiber.Router) {

	// router
	// init user router, etc ...
	// rate limiter
	// CORS
	r.userRouter.AddRoutes(router)

}
