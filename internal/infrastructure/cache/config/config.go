package config

import "fmt"

type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func (r *RedisConfig) FullAddress() string {
	return fmt.Sprintf("%s:%d", r.Host, r.Port)
}
