package config

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/config"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"presentation.config",
	fx.Provide(
		config.NewConfig,
		config.NewDBConfig,
		config.NewAPIConfig,
		config.NewBrokerConfig,
		config.NewCronConfig,
		config.NewLoggerConfig,
		config.NewCacheConfig,
		config.NewAppConfig,
	),
)
