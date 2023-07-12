package config

import (
	load "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/config"
	db "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/config"
	c "github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/config"
)

func NewConfig() Config {
	var conf Config
	load.LoadConfig(&conf, "", "")
	return conf
}
func NewDBConfig(config Config) db.DBConfig {
	return config.DBConfig
}
func NewAPIConfig(config Config) c.APIConfig {
	return config.APIConfig
}
