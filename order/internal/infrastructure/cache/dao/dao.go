package dao

import "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/cache"

type BaseRedisDAO struct {
	cache.Cache
}
