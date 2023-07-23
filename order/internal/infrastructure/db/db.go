package db

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/db/config"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func BuildConnection(logger logger.Logger, config config.DBConfig) *gorm.DB {
	gormConfig := gorm.Config{}
	if !config.Logging {
		gormConfig.Logger = gormLogger.Default.LogMode(gormLogger.Silent)
	} else {
		gormConfig.Logger = logger.GetGormLogger()
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
