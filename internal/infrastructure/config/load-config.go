package config

import (
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

const DefaultConfigPath = "./config/dev.toml"

func LoadConfig(val interface{}, absolutePath string, relativePath string) {
	relativeEnv := getEnv("DEFAULT_CONFIG_PATH", "")
	if relativeEnv == "" {
		relativeEnv = DefaultConfigPath
	}
	if relativePath != "" && getEnv("DEFAULT_CONFIG_PATH", "") == "" {
		relativeEnv = relativePath
	}

	var pathConf string
	if absolutePath != "" {
		pathConf = filepath.Join(absolutePath, relativeEnv)
	} else {
		pathConf = relativeEnv
	}
	_, err := toml.DecodeFile(pathConf, val)
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
