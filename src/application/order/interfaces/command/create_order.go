package command

import (
	"github.com/google/uuid"
)

type CreateOrderCommand struct {
	OrderID         uuid.UUID
	PaymentMethod   string
	ProductsIDs     []uuid.UUID
	UserID          uuid.UUID
	DeliveryAddress uuid.UUID
}

type CreateOrder interface {
	Create(command CreateOrderCommand) error
}
