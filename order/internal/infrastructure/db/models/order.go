package models

import (
	"github.com/google/uuid"
)

type Order struct {
	Base
	OrderStatus   string
	ClientID      uuid.UUID
	PaymentMethod string
	AddressID     uuid.UUID
	Closed        bool
	SerialNumber  int

	SagaStatus string `gorm:"default:Pending"`

	Products []Product `gorm:"many2many:order_products;"`
}
