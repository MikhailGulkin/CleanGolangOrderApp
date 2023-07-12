package models

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	Base
	Products      []Product `gorm:"many2many:order_products;"`
	OrderStatus   string
	ClientID      uuid.UUID
	Client        User
	PaymentMethod string
	AddressID     uuid.UUID
	Price         float64
	Date          time.Time
	SerialNumber  int
}
