package db

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	_ = db.AutoMigrate(&models.Product{}, &models.User{}, &models.Address{}, &models.Order{})
}
