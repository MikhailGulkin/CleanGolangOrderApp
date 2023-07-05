package db

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Product{})
	if err != nil {
		return
	}
}

func BuildConnection(config config.DBConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.FullDNS()), &gorm.Config{})
	if err == nil {
		_ = db.AutoMigrate(&models.Product{})
		return db
	}
	panic(err.Error())
}
