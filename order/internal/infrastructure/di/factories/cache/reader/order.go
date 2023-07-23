package reader

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/interfaces/persistence/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/cache"
	r "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/cache/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/cache/reader/order"
	"go.uber.org/fx"
)

func NewOrderCacheReader(cache cache.Cache) reader.OrderCacheReader {
	return &order.CacheReaderImpl{
		BaseRedisReader: r.BaseRedisReader{Cache: cache},
	}
}

var Module = fx.Provide(
	NewOrderCacheReader,
)
