package db

import "github.com/MikhailGulkin/CleanGolangOrderApp/pkg/env"

type Config struct {
	DSN string
}

func NewConfig() Config {
	dns := "postgres://postgres:postgres@localhost:5431/postgres"
	return Config{
		DSN: env.GetEnv("DSN", dns),
	}
}
func (conf *Config) GetDSN() string {
	return conf.DSN
}
