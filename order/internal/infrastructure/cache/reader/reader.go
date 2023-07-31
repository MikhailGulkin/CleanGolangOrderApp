package reader

import "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/cache"

type BaseRedisReader struct {
	cache.Cache
}
