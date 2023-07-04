package config

import (
	"github.com/BurntSushi/toml"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db"
)

const DefaultConfigPath = "./config/dev.toml"

func LoadConfig(val interface{}, scope string) {
	var mapConf map[string]map[string]interface{}

	_, err := toml.DecodeFile(DefaultConfigPath, &mapConf)
	if err != nil {
		panic(err)
	}
	if dbConf, ok := val.(*db.ConfigDB); ok && scope == "db" {
		dbConf.Password = mapConf["db"]["password"].(string)
		dbConf.Host = mapConf["db"]["host"].(string)
		dbConf.Database = mapConf["db"]["database"].(string)
		dbConf.Port = int(mapConf["db"]["port"].(int64))
		dbConf.User = mapConf["db"]["user"].(string)
	}

	_, err = toml.DecodeFile(DefaultConfigPath, val)
	if err != nil {
		panic(err)
	}
}
