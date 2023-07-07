package config

import (
	"fmt"
)

type DBConfig struct {
	Host              string `toml:"host"`
	Port              int    `toml:"port"`
	Database          string `toml:"database"`
	User              string `toml:"user"`
	Password          string `toml:"password"`
	Migration         bool   `toml:"migration"`
	MaxIdleConnection int    `toml:"max_idle_connection"`
	Logging           bool   `toml:"logging"`
}

func (conf *DBConfig) FullDNS() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		conf.Host, conf.User, conf.Password, conf.Database, conf.Port,
	)
}
