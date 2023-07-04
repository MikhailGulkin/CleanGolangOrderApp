package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func BuildConnection(config ConfigDB) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.FullDNS()), &gorm.Config{})
	if err == nil {
		return db
	}
	panic(err.Error())
}
