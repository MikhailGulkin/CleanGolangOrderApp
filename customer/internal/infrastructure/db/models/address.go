package models

import (
	"fmt"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/models"
)

type Address struct {
	models.Base
	BuildingNumber int    `gorm:"index:UniqueAddress, unique"`
	StreetName     string `gorm:"index:UniqueAddress, unique"`
	City           string `gorm:"index:UniqueAddress, unique"`
	Country        string `gorm:"index:UniqueAddress, unique"`
	Orders         []models.Order
}

func (a *Address) GetFullAddress() string {
	return fmt.Sprintf("%s, %s, %s, %d", a.Country, a.City, a.StreetName, a.BuildingNumber)
}
