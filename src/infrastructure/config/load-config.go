package config

import (
	"github.com/BurntSushi/toml"
)

const DefaultConfigPath = "./config/dev.toml"

func LoadConfig(val interface{}) {
	_, err := toml.DecodeFile(DefaultConfigPath, val)
	if err != nil {
		panic(err)
	}
}
