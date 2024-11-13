package middlewares

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type RateLimit interface {
	LimitMiddleware(c *fiber.Ctx) error
}

type RateLimiter struct {
	redisClient *redis.Client
	limit       int
	window      time.Duration
}

func NewRateLimiter(redisClient *redis.Client, limit int, window time.Duration) RateLimit {
	return &RateLimiter{
		redisClient: redisClient,
		limit:       limit,
		window:      window,
	}
}

func (r *RateLimiter) LimitMiddleware(c *fiber.Ctx) error {
	ip := c.IP()
	ctx := context.Background()

	allowRequests := true

	countStr, err := r.redisClient.Get(ctx, ip+":count").Result()
	if err == redis.Nil {
		countStr = "0"
	} else if err != nil {
		log.Printf("Error fetching count from Redis: %v", err)
		allowRequests = false
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		log.Printf("Error converting count to integer: %v", err)
		allowRequests = false
	}

	timestampStr, err := r.redisClient.Get(ctx, ip+":timestamp").Result()
	if err != nil && err != redis.Nil {
		log.Printf("Error fetching timestamp from Redis: %v", err)
		allowRequests = false
	}

	var timestamp time.Time
	if timestampStr == "" {
		timestamp = time.Now()
	} else {
		timestamp, err = time.Parse(time.RFC3339, timestampStr)
		if err != nil {
			log.Printf("Error parsing timestamp: %v", err)
			allowRequests = false
		}
	}

	now := time.Now()
	if now.Sub(timestamp) > r.window {
		count = 0
		timestamp = now
	}

	if allowRequests {
		count++

		if count > r.limit {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests, please try again later.",
			})
		}

		if err := r.redisClient.Set(ctx, ip+":count", count, r.window).Err(); err != nil {
			log.Printf("Error setting count in Redis: %v", err)
		}
		if err := r.redisClient.Set(ctx, ip+":timestamp", now.Format(time.RFC3339), r.window).Err(); err != nil {
			log.Printf("Error setting timestamp in Redis: %v", err)
		}
	}

	return c.Next()
}
