package events

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/common/events"
	"github.com/google/uuid"
)

type OrderDeleted struct {
	events.BaseEvent
	OrderID uuid.UUID `json:"orderID"`
}

func (OrderDeleted) Create(
	orderID uuid.UUID,
) events.Event {
	return &OrderDeleted{
		BaseEvent: events.BaseEvent{}.Create("OrderDeleted"),
		OrderID:   orderID,
	}
}
func (o *OrderDeleted) Bytes() ([]byte, error) {
	return events.Bytes(o)
}
func (o *OrderDeleted) UniqueAggregateID() uuid.UUID {
	return o.OrderID
}
