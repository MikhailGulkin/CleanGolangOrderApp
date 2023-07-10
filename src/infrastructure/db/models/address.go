package models

type Address struct {
	Base
	BuildingNumber int    `gorm:"index:UniqueAddress, unique"`
	StreetName     string `gorm:"index:UniqueAddress, unique"`
	City           string `gorm:"index:UniqueAddress, unique"`
	Country        string `gorm:"index:UniqueAddress, unique"`
	Orders         []Order
}
