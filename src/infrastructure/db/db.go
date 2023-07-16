package db

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func BuildConnection(logger logger.Logger, config config.DBConfig) *gorm.DB {
	gormConfig := gorm.Config{
		Logger: logger.GetGormLogger(),
	}
	if !config.Logging {
		gormConfig.Logger = gormLogger.Default.LogMode(gormLogger.Silent)
	}
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
