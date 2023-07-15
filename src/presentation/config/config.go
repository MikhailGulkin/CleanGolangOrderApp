package config

import (
	db "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/config"
	broker "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/message_broker/config"
	api "github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/config"
	"go.uber.org/fx"
)

type Config struct {
	db.DBConfig                `toml:"db"`
	api.APIConfig              `toml:"api"`
	broker.MessageBrokerConfig `toml:"broker"`
}

var Module = fx.Provide(
	NewConfig,
	NewDBConfig,
	NewAPIConfig,
	NewBrokerConfig,
)
