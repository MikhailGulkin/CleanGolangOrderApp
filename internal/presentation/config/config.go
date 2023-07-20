package config

import (
	cache "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/cache/config"
	db "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/config"
	logger "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/logger/config"
	broker "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/messageBroker/config"
	api "github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/config"
	cron "github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/cron/config"
	"go.uber.org/fx"
)

type App struct {
	Mode string
}
type Config struct {
	App                        `toml:"app"`
	db.DBConfig                `toml:"db"`
	api.APIConfig              `toml:"api"`
	broker.MessageBrokerConfig `toml:"broker"`
	cron.CronConfig            `toml:"handlers"`
	logger.LoggerConfig        `toml:"logging"`
	cache.RedisConfig          `toml:"cache"`
}

var Module = fx.Provide(
	NewConfig,
	NewDBConfig,
	NewAPIConfig,
	NewBrokerConfig,
	NewCronConfig,
	NewLoggerConfig,
	NewCacheConfig,
)
