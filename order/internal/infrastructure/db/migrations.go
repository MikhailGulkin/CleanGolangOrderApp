package db

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/models"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	_ = db.AutoMigrate(&models.Product{}, &models.Order{}, &models.Outbox{})
}
