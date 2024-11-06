package app

import "github.com/gofiber/fiber/v3"

func (a *application) InitFramework() *fiber.App {
	return fiber.New()
}
