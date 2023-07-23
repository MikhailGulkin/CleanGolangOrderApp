package cache

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/cache"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/di/factories/cache/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/di/factories/cache/reader"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(cache.NewClient),
	dao.Module,
	reader.Module,
)
