package controllers

import (
	"golang_template/internal/logging"

	"math/rand"
	"time"

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

// func (c *userController) Login(ctx *fiber.Ctx) error {
// 	c.logger.Info("Request recieved")

//		err := ctx.JSON("succeed")
//		if err != nil {
//			c.logger.Error("Request Failed")
//			return err
//		}
//		return nil
//		//make dto
//	}
func (c *userController) Login(ctx *fiber.Ctx) error {
	c.logger.Info("Request received")

	// Parse phone number from request body
	var body struct {
		PhoneNumber string `json:"phone_number"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		c.logger.Error("Failed to parse request body")
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if body.PhoneNumber == "" {
		c.logger.Error("Phone number is missing in the request")
		return fiber.NewError(fiber.StatusBadRequest, "Phone number is required")
	}

	c.logger.Info("Phone number received: " + body.PhoneNumber)

	// Generate a random response status
	rand.NewSource(time.Now().UnixNano())
	statusOptions := []int{
		fiber.StatusTooManyRequests, // 429
		fiber.StatusBadRequest,      // 400
		fiber.StatusNotFound,        // 404
		fiber.StatusCreated,         // 201
	}
	randomStatus := statusOptions[rand.Intn(len(statusOptions))]

	switch randomStatus {
	case fiber.StatusTooManyRequests:
		c.logger.Error("Simulating rate limit")
		return fiber.NewError(randomStatus, "Too many requests, try again later")
	case fiber.StatusBadRequest:
		c.logger.Error("Simulating bad request")
		return fiber.NewError(randomStatus, "Bad request, check your data")
	case fiber.StatusNotFound:
		c.logger.Error("Simulating resource not found")
		return fiber.NewError(randomStatus, "Resource not found")
	case fiber.StatusCreated:
		c.logger.Info("Simulating successful request")
		return ctx.Status(randomStatus).JSON(fiber.Map{
			"message": "Login succeeded",
		})
	}

	// Fallback
	c.logger.Error("Unhandled status")
	return fiber.NewError(fiber.StatusInternalServerError, "Internal server error")
}
