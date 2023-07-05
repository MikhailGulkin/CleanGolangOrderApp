package providers

import (
	load "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/config"
	db "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/config"
)

func NewConfig() config.Config {
	var conf config.Config
	load.LoadConfig(&conf, "", "")
	return conf
}
func NewDBConfig(config config.Config) db.DBConfig {
	return config.DBConfig
}
func NewAPIConfig(config config.Config) config.APIConfig {
	return config.APIConfig
}
