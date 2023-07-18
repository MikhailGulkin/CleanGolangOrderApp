package events

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/common/events"
	"github.com/google/uuid"
)

type OrderCreatedClient struct {
	ClientID uuid.UUID `json:"clientID"`
	Username string    `json:"username"`
}
type OrderCreated struct {
	events.BaseEvent
	OrderID         uuid.UUID          `json:"orderID"`
	Client          OrderCreatedClient `json:"client"`
	PaymentMethod   string             `json:"paymentMethod"`
	DeliveryAddress uuid.UUID          `json:"deliveryAddress"`
	SerialNumber    int                `json:"serialNumber"`
	TotalPrice      float64            `json:"totalPrice"`
}

func (OrderCreated) Create(
	orderID uuid.UUID,
	client OrderCreatedClient,
	paymentMethod string,
	address uuid.UUID,
	serialNumber int,
	totalPrice float64,
) events.Event {
	return &OrderCreated{
		OrderID:         orderID,
		BaseEvent:       events.BaseEvent{}.Create(),
		Client:          client,
		PaymentMethod:   paymentMethod,
		DeliveryAddress: address,
		SerialNumber:    serialNumber,
		TotalPrice:      totalPrice,
	}
}

func (o *OrderCreated) Bytes() ([]byte, error) {
	return events.Bytes(o)
}
