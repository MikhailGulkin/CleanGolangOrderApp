package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

const DefaultConfigPath = "./config/dev.toml"

func LoadConfig(val interface{}) {
	_, err := toml.DecodeFile(getEnv("DEFAULT_CONFIG_PATH", DefaultConfigPath), val)
	if err != nil {
		panic(err)
	}
}
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
