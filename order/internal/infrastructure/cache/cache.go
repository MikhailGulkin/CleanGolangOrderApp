package cache

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/cache/config"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	*redis.Client
}

func NewClient(redisConf config.RedisConfig) Cache {
	return Cache{Client: redis.NewClient(
		&redis.Options{
			Addr:     redisConf.FullAddress(),
			Password: redisConf.Password,
			DB:       redisConf.DB,
		},
	)}
}
