package config

import (
	cache "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/cache/config"
	load "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/config"
	db "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/config"
	logger "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/logger/config"
	broker "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/messageBroker/config"
	api "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/api/config"
	cron "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/cron/config"
)

func NewConfig() Config {
	var conf Config
	load.LoadConfig(&conf, "", "")
	return conf
}
func NewDBConfig(config Config) db.DBConfig {
	return config.DBConfig
}
func NewAppConfig(config Config) AppConfig {
	return config.AppConfig
}
func NewAPIConfig(config Config) api.APIConfig {
	config.APIConfig.Mode = config.AppConfig.Mode
	return config.APIConfig
}
func NewBrokerConfig(config Config) broker.MessageBrokerConfig {
	return config.MessageBrokerConfig
}
func NewCronConfig(config Config) cron.CronConfig {
	return config.CronConfig
}
func NewLoggerConfig(config Config) logger.LoggerConfig {
	config.LoggerConfig.Mode = config.AppConfig.Mode
	return config.LoggerConfig
}
func NewCacheConfig(config Config) cache.RedisConfig {
	return config.RedisConfig
}
