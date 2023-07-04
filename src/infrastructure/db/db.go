package db

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func BuildConnection(config config.DBConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.FullDNS()), &gorm.Config{})
	if err == nil {
		return db
	}
	panic(err.Error())
}
