package events

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/common/events"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/consts"
	"github.com/google/uuid"
)

type OrderCreateProductEvent struct {
	ProductID  uuid.UUID `json:"productID"`
	Name       string    `json:"name"`
	TotalPrice float64   `json:"totalPrice"`
}

type OrderCreated struct {
	events.BaseEvent
	OrderID           uuid.UUID                 `json:"orderID"`
	ClientID          uuid.UUID                 `json:"clientID"`
	OrderStatus       string                    `json:"orderStatus"`
	PaymentMethod     string                    `json:"paymentMethod"`
	Products          []OrderCreateProductEvent `json:"products"`
	DeliveryAddressID uuid.UUID                 `json:"deliveryAddressID"`
	SerialNumber      int                       `json:"serialNumber"`
	TotalPrice        float64                   `json:"totalPrice"`
}

func (OrderCreated) Create(
	orderID uuid.UUID,
	client uuid.UUID,
	paymentMethod string,
	serialNumber int,
	products []OrderCreateProductEvent,
	totalPrice float64,
	address uuid.UUID,
) events.Event {
	return &OrderCreated{
		OrderID:           orderID,
		BaseEvent:         events.BaseEvent{}.Create("OrderCreated"),
		ClientID:          client,
		PaymentMethod:     paymentMethod,
		Products:          products,
		DeliveryAddressID: address,
		OrderStatus:       string(consts.New),
		SerialNumber:      serialNumber,
		TotalPrice:        totalPrice,
	}
}

func (o *OrderCreated) Bytes() ([]byte, error) {
	return events.Bytes(o)
}
func (o *OrderCreated) UniqueAggregateID() uuid.UUID {
	return o.OrderID
}
