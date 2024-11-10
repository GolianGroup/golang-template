package controllers

import "github.com/gofiber/fiber/v2"

type IUserController interface {
	Login(ctx *fiber.Ctx) error
}

type UserController struct {
}

// inject user service to user controller

func NewUserController() IUserController {
	return &UserController{}
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	err := ctx.JSON("succeed")
	if err != nil {
		return err
	}
	return nil
	//make dto
}
