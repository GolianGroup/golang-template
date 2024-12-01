package controllers

import (
	"errors"
	err "golang_template/handler/errors"
	"golang_template/internal/services"
	"golang_template/internal/services/dto"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrInternal = &err.AppError{
		Message: services.ErrInternal.Message(),
		Err:     services.ErrInternal.Unwrap(),
		Code:    fiber.StatusInternalServerError,
	}

	ErrInvalidCredentials = &err.AppError{
		Message: services.ErrInvalidCredentials.Message(),
		Err:     services.ErrInvalidCredentials.Unwrap(),
		Code:    fiber.StatusBadRequest,
	}

	ErrUserNotFound = &err.AppError{
		Message: services.ErrUserNotFound.Message(),
		Err:     services.ErrUserNotFound.Unwrap(),
		Code:    fiber.StatusNotFound,
	}

	ErrParseRequest = &err.AppError{
		Message: "Parsing error occured",
		Err:     errors.New("parsing error"),
		Code:    fiber.StatusBadRequest,
	}
)

type UserController interface {
	Login(ctx *fiber.Ctx) error
}

type userController struct {
	service services.UserService
}

// inject user service to user controller

func NewUserController(service services.UserService) UserController {
	return &userController{service: service}
}

func (c *userController) Login(ctx *fiber.Ctx) error {
	userDto := dto.User{}
	err := ctx.BodyParser(&userDto)

	if err != nil {
		return ctx.Status(ErrParseRequest.Code).JSON(ErrParseRequest)
	}

	user, err := c.service.Login(ctx, userDto)

	if err == nil {
		return ctx.Status(fiber.StatusOK).JSON(user)
	}
	if errors.Is(err, services.ErrUserNotFound) {
		return ctx.Status(ErrUserNotFound.Code).JSON(ErrUserNotFound)
	}

	return ctx.Status(ErrInternal.Code).JSON(ErrInternal)
}
