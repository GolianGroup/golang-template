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
	userRouter  UserRouter
	videoRouter VideoRouter
}


func NewRouter(controllers controllers.Controllers, redisClient producers.RedisClient) Router {
	userRouter := NewUserRouter(controllers.UserController(), redisClient)
	videoRouter := NewVideoRouter(controllers.VideoController())
	return &router{
		userRouter:  userRouter,
		videoRouter: videoRouter,
	}
	return &router{userRouter: userRouter, redisClient: redisClient}
}

func (r router) AddRoutes(router fiber.Router) {

	// router
	// init user router, etc ...
	// rate limiter
	// CORS
	r.userRouter.AddRoutes(router)
	r.videoRouter.AddRoutes(router)

}
