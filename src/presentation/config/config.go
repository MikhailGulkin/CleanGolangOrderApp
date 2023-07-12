package config

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/config"
	c "github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/config"
	"go.uber.org/fx"
)

type Config struct {
	config.DBConfig `toml:"db"`
	c.APIConfig     `toml:"api"`
}

var Module = fx.Provide(
	NewConfig,
	NewDBConfig,
	NewAPIConfig,
)
