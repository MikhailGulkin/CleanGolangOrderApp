package config

import db "github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/infrastructure/db/config"

type App struct {
	Mode string
}
type Config struct {
	App         `toml:"app"`
	db.DBConfig `toml:"db"`
}
