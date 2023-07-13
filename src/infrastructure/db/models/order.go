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
	Address       Address
	AddressID     uuid.UUID
	Date          time.Time
	Closed        bool
	SerialNumber  int
}
