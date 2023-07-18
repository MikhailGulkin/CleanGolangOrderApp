package models

type Product struct {
	Base
	Price        float64
	Name         string `gorm:"unique"`
	Discount     int32
	Quantity     int32
	Description  string
	Availability bool
}
