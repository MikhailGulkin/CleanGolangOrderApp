package config

import (
	load "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/config"
	db "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/config"
)

func NewConfig() Config {
	var conf Config
	load.LoadConfig(&conf, "", "")
	return conf
}
func NewDBConfig(config Config) db.DBConfig {
	return config.DBConfig
}
func NewAPIConfig(config Config) APIConfig {
	return config.APIConfig
}
