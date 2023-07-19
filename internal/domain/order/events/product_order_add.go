package events

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/common/events"
	"github.com/google/uuid"
)

type OrderAddProduct struct {
	events.BaseEvent
	OrderID      uuid.UUID `json:"orderID"`
	ClientID     uuid.UUID `json:"clientID"`
	ProductID    uuid.UUID `json:"productID"`
	ProductName  string    `json:"productName"`
	ProductPrice float64   `json:"productPrice"`
	OrderPrice   float64   `json:"orderPrice"`
}

func (OrderAddProduct) Create(
	orderID uuid.UUID,
	product uuid.UUID,
	productPrice float64,
	orderPrice float64,
	productName string,
	clientID uuid.UUID,
) events.Event {
	return &OrderAddProduct{
		BaseEvent:    events.BaseEvent{}.Create(),
		OrderID:      orderID,
		ProductID:    product,
		ProductName:  productName,
		ProductPrice: productPrice,
		OrderPrice:   orderPrice,
		ClientID:     clientID,
	}
}
func (o *OrderAddProduct) Bytes() ([]byte, error) {
	return events.Bytes(o)
}