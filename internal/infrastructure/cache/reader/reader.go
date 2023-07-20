package reader

import "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/cache"

type BaseRedisReader struct {
	cache.Cache
}
