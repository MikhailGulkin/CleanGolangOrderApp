package cache

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/cache"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/di/factories/cache/dao"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/di/factories/cache/reader"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(cache.NewClient),
	dao.Module,
	reader.Module,
)
