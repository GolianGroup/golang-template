package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type SystemController interface {
	HealthCheck() fiber.Handler
	ReadyCheck() fiber.Handler
}

type systemController struct {
}

func NewSystemController() SystemController {
	return &systemController{}
}

func (controller *systemController) HealthCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "UP",
			"time":   time.Now(),
		})
	}
}

func (controller *systemController) ReadyCheck() fiber.Handler {

	return func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "READY",
		})
	}
}
