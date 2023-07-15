package events

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/common/events"
	"github.com/google/uuid"
)

type OrderCreated struct {
	events.BaseEvent
	Client          uuid.UUID
	PaymentMethod   string
	DeliveryAddress uuid.UUID
	SerialNumber    int
}

func (OrderCreated) Create(client uuid.UUID, paymentMethod string, address uuid.UUID, serialNumber int) events.Event {
	return &OrderCreated{
		BaseEvent:       events.BaseEvent{}.Create(),
		Client:          client,
		PaymentMethod:   paymentMethod,
		DeliveryAddress: address,
		SerialNumber:    serialNumber,
	}
}

func (o *OrderCreated) Bytes() ([]byte, error) {
	return events.Bytes(o)
}
