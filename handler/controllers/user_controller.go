package controllers

import (
	"golang_template/internal/logging"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Login(ctx *fiber.Ctx) error
}

type userController struct {
	logger logging.Logger
}

// inject user service to user controller

func NewUserController(logger logging.Logger) UserController {
	return &userController{
		logger: logger,
	}
}

func (c *userController) Login(ctx *fiber.Ctx) error {
	c.logger.Info("Request recieved")
	err := ctx.JSON("succeed")
	if err != nil {
		c.logger.Error("Request Failed")
		return err
	}
	return nil
	//make dto
}
