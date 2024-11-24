package controllers

import (
	"golang_template/internal/logging"

	oteltrace "go.opentelemetry.io/otel/trace"
)

type Controllers interface {
	UserController() UserController
}

type controllers struct {
	userController UserController
}

func NewControllers(logger logging.Logger, tracer oteltrace.Tracer) Controllers {
	userController := NewUserController(logger, tracer)
	return &controllers{userController: userController}
}

func (c *controllers) UserController() UserController {
	return c.userController
}
