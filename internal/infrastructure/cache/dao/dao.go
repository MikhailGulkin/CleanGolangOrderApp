package dao

import "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/cache"

type BaseRedisDAO struct {
	cache.Cache
}
