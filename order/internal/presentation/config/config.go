package config

import (
	cache "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/cache/config"
	db "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/config"
	logger "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/logger/config"
	broker "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/messageBroker/config"
	api "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/config"
	cron "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/cron/config"
)

type AppConfig struct {
	Mode string
}
type Config struct {
	AppConfig                  `toml:"app"`
	db.DBConfig                `toml:"db"`
	api.APIConfig              `toml:"api"`
	broker.MessageBrokerConfig `toml:"broker"`
	cron.CronConfig            `toml:"handlers"`
	logger.LoggerConfig        `toml:"logging"`
	cache.RedisConfig          `toml:"cache"`
}
