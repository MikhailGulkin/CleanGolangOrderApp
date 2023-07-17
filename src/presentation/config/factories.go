package config

import (
	load "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/config"
	db "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/logger/config"
	broker "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/messageBroker/config"
	api "github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/api/config"
	cron "github.com/MikhailGulkin/simpleGoOrderApp/src/presentation/cron/config"
)

func NewConfig() Config {
	var conf Config
	load.LoadConfig(&conf, "", "")
	return conf
}
func NewDBConfig(config Config) db.DBConfig {
	return config.DBConfig
}
func NewAPIConfig(config Config) api.APIConfig {
	return config.APIConfig
}
func NewBrokerConfig(config Config) broker.MessageBrokerConfig {
	return config.MessageBrokerConfig
}
func NewCronConfig(config Config) cron.CronConfig {
	return config.CronConfig
}
func NewLoggerConfig(config Config) config.LoggerConfig {
	return config.LoggerConfig
}
