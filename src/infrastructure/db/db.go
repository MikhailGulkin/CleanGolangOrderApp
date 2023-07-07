package db

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func BuildConnection(config config.DBConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.FullDNS()), &gorm.Config{})
	if err == nil {
		if config.Migration {
			_ = db.AutoMigrate(&models.Product{}, &models.Address{})
		}

		return db
	}
	panic(err.Error())
}
