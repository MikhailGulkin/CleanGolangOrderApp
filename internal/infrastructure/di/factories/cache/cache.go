package cache

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/cache"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	cache.NewClient,
)
