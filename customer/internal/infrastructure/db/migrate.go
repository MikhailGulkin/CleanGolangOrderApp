package db

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/infrastructure/db/models"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	_ = db.AutoMigrate(&models.Customer{})
}
