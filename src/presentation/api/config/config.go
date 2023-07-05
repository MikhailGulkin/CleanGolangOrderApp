package config

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/config"
	"go.uber.org/fx"
)

type APIConfig struct {
	Host          string
	Port          int
	BaseURLPrefix string `toml:"base_url_prefix"`
}
type Config struct {
	config.DBConfig `toml:"db"`
	APIConfig       `toml:"api"`
}

var Module = fx.Provide(
	NewConfig,
	NewDBConfig,
	NewAPIConfig,
)
