package models

import "fmt"

type Address struct {
	Base
	BuildingNumber int    `gorm:"index:UniqueAddress, unique"`
	StreetName     string `gorm:"index:UniqueAddress, unique"`
	City           string `gorm:"index:UniqueAddress, unique"`
	Country        string `gorm:"index:UniqueAddress, unique"`
	Orders         []Order
}

func (a *Address) GetFullAddress() string {
	return fmt.Sprintf("%s, %s, %s, %d", a.Country, a.City, a.StreetName, a.BuildingNumber)
}
