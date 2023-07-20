package events

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/common/events"
	"github.com/google/uuid"
)

type OrderCreateSaga struct {
	events.BaseEvent
	OrderID    uuid.UUID `json:"orderID"`
	TotalPrice float64   `json:"totalPrice"`
}

func (OrderCreateSaga) Create(orderID uuid.UUID, price float64) events.Event {
	return &OrderCreateSaga{
		BaseEvent:  events.BaseEvent{}.Create(),
		OrderID:    orderID,
		TotalPrice: price,
	}
}
func (o *OrderCreateSaga) Bytes() ([]byte, error) {
	return events.Bytes(o)
}
func (o *OrderCreateSaga) UniqueAggregateID() uuid.UUID {
	return o.OrderID
}
