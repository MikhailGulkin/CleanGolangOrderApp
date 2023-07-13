package command

import (
	"github.com/google/uuid"
)

type CreateOrderCommand struct {
	OrderID         uuid.UUID   `json:"orderID" binding:"required"`
	PaymentMethod   string      `json:"paymentMethod" binding:"required"`
	ProductsIDs     []uuid.UUID `json:"productsIDs" binding:"required"`
	UserID          uuid.UUID   `json:"userID" binding:"required"`
	DeliveryAddress uuid.UUID   `json:"deliveryAddress" binding:"required"`
}

type CreateOrder interface {
	Create(command CreateOrderCommand) error
}
