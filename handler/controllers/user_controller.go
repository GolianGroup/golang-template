package controllers

import (
	"golang_template/internal/logging"

	"github.com/gofiber/fiber/v2"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type UserController interface {
	Login(ctx *fiber.Ctx) error
}

type userController struct {
	logger logging.Logger
	tracer oteltrace.Tracer
}

// inject user service to user controller

func NewUserController(logger logging.Logger, tracer oteltrace.Tracer) UserController {
	return &userController{
		logger: logger,
		tracer: tracer,
	}
}

func (c *userController) Login(ctx *fiber.Ctx) error {
	c.logger.Info("Request recieved")

	_, span := c.tracer.Start(ctx.Context(), "Login")
	defer span.End()

	err := ctx.JSON("succeed")
	if err != nil {
		c.logger.Error("Request Failed")
		return err
	}
	return nil
	//make dto
}
