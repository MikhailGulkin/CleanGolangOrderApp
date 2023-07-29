package db

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/infrastructure/db/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func BuildConnection(config config.DBConfig) *gorm.DB {
	gormConfig := gorm.Config{}
	db, err := gorm.Open(postgres.Open(config.FullDNS()), &gormConfig)
	sqlDB, errSQL := db.DB()
	if errSQL != nil {
		panic(errSQL)
	}
	sqlDB.SetMaxIdleConns(config.MaxIdleConnection)
	if err == nil {
		if config.Migration {
			migrate(db)
		}

		return db
	}
	panic(err.Error())
}
