package config

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/infrastructure/db"
)

type App struct {
	Mode string
}
type Config struct {
	App         `toml:"app"`
	db.DBConfig `toml:"db"`
}
