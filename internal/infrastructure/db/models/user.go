package models

import "github.com/google/uuid"

type User struct {
	Base
	Username  string
	Orders    []Order `gorm:"many2many:user_orders;"`
	AddressID uuid.UUID
	Address   Address
}
