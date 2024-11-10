package controllers

import "github.com/gofiber/fiber/v2"

type UserController interface {
	Login(ctx *fiber.Ctx) error
}

type userController struct {
}

// inject user service to user controller

func NewUserController() UserController {
	return &userController{}
}

func (c *userController) Login(ctx *fiber.Ctx) error {
	err := ctx.JSON("succeed")
	if err != nil {
		return err
	}
	return nil
	//make dto
}
