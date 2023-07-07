package db

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func BuildConnection(config config.DBConfig) *gorm.DB {
	gormConfig := gorm.Config{}
	if !config.Logging {
		gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	}
	db, err := gorm.Open(postgres.Open(config.FullDNS()), &gormConfig)
	sqlDB, errSQL := db.DB()
	if errSQL != nil {
		panic(errSQL)
	}
	sqlDB.SetMaxIdleConns(config.MaxIdleConnection)
	if err == nil {
		if config.Migration {
			_ = db.AutoMigrate(&models.Product{}, &models.Address{})
		}

		return db
	}
	panic(err.Error())
}
