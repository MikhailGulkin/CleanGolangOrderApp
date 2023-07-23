package dao

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/cache"
	base "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/cache/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/cache/dao/order"
	"go.uber.org/fx"
)

func NewOrderCacheDAO(cache cache.Cache) dao.OrderCacheDAO {
	return &order.CacheDAOImpl{
		BaseRedisDAO: base.BaseRedisDAO{Cache: cache},
	}
}

var Module = fx.Provide(
	NewOrderCacheDAO,
)
