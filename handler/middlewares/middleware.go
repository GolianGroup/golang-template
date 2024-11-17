package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/storage/redis/v3"
)

func LimiterConfig(max int, expiration int, storage *redis.Storage) limiter.Config {
	expirationDuration := time.Duration(expiration) * time.Second
	return limiter.Config{
		Max:        max,
		Expiration: expirationDuration,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests, please try again later.",
			})
		},
		Storage: storage,
	}
}
