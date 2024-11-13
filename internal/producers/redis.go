package producers

import (
	"golang_template/internal/utils"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient interface {
	GetRedisClient() *redis.Client
	Close() error
}

type Redis struct {
	client *redis.Client
}

func NewRedis(config *utils.RedisConfig) RedisClient {
	addr := utils.GetDSNRedis(config)
	opt, err := redis.ParseURL(addr)
	if err != nil {
		log.Fatal(err)
	}

	opt.PoolSize = 10
	opt.MinIdleConns = 2
	opt.DialTimeout = 5 * time.Second
	opt.ReadTimeout = 3 * time.Second
	opt.WriteTimeout = 3 * time.Second
	opt.IdleTimeout = 5 * time.Minute

	client := redis.NewClient(opt)
	return &Redis{
		client: client,
	}
}

func (r *Redis) GetRedisClient() *redis.Client {
	return r.client
}

func (r *Redis) Close() error {
	return r.client.Close()
}
