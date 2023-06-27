package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	dsn := "host=localhost user=postgres password=1234 dbname=OrderApp port=5431"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err == nil {
		err := db.AutoMigrate(&Product{})
		if err != nil {
			return
		}
	}

}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
