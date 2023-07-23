package dao

import "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/cache"

type BaseRedisDAO struct {
	cache.Cache
}
