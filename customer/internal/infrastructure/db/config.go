package db

import (
	"fmt"
	"github.com/MikhailGulkin/CleanGolangOrderApp/pkg/env"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

func NewConfig() Config {
	port, err := strconv.Atoi(env.GetEnv("POSTGRES_PORT", "5432"))
	if err != nil {
		panic(err)
	}
	return Config{
		Host:     env.GetEnv("POSTGRES_HOST", "localhost"),
		Port:     port,
		Database: env.GetEnv("POSTGRES_DB", "postgres"),
		User:     env.GetEnv("POSTGRES_USER", "postgres"),
		Password: env.GetEnv("POSTGRES_PASSWORD", "postgres"),
	}
}
func (conf *Config) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Database,
	)
}
