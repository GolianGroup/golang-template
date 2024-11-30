package routers

import (
	"golang_template/handler/controllers"

	"golang_template/handler/middlewares"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type Router interface {
	AddRoutes(router fiber.Router)
}

type router struct {
	userRouter UserRouter
	tracer     trace.Tracer
}

func NewRouter(controllers controllers.Controllers, tracer oteltrace.Tracer) Router {
	userRouter := NewUserRouter(controllers.UserController())
	return &router{userRouter: userRouter,
		tracer: tracer}
}

func (r router) AddRoutes(router fiber.Router) {

	// router
	// init user router, etc ...
	// rate limiter
	// CORS
	router.Use(middlewares.TracingMiddleware(r.tracer))
	r.userRouter.AddRoutes(router)

}
