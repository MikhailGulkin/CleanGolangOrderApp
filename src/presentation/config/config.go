package config

import (
	db "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/logger/config"
	broker "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/messageBroker/config"
	api "github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/config"
	cron "github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/cron/config"
	"go.uber.org/fx"
)

type Config struct {
	db.DBConfig                `toml:"db"`
	api.APIConfig              `toml:"api"`
	broker.MessageBrokerConfig `toml:"broker"`
	cron.CronConfig            `toml:"handlers"`
	config.LoggerConfig        `toml:"logging"`
}

var Module = fx.Provide(
	NewConfig,
	NewDBConfig,
	NewAPIConfig,
	NewBrokerConfig,
	NewCronConfig,
	NewLoggerConfig,
)
