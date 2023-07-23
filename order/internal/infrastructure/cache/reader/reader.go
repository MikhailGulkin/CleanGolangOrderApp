package reader

import "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/cache"

type BaseRedisReader struct {
	cache.Cache
}
