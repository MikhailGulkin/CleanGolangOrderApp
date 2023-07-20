package cache

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/cache"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/di/factories/cache/dao"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(cache.NewClient),
	dao.Module,
)
