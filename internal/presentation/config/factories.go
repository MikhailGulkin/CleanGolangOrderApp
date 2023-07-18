package config

import (
	cache "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/cache/config"
	load "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/config"
	db "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/config"
	logger "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/logger/config"
	broker "github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/messageBroker/config"
	api "github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/api/config"
	cron "github.com/MikhailGulkin/simpleGoOrderApp/internal/presentation/cron/config"
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
func NewLoggerConfig(config Config) logger.LoggerConfig {
	return config.LoggerConfig
}
func NewCacheConfig(config Config) cache.RedisConfig {
	return config.RedisConfig
}
