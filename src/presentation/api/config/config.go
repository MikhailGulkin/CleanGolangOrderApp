package config

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/config"
)

type APIConfig struct {
	Host string
	Port int
}
type Config struct {
	config.DBConfig `toml:"db"`
	APIConfig       `toml:"api"`
}
