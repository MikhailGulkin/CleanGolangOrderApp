package events

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/common/events"
	"github.com/google/uuid"
)

type OrderAddProduct struct {
	events.BaseEvent
	Product    uuid.UUID
	OrderPrice float64
}

func (OrderAddProduct) Create(product uuid.UUID, orderPrice float64) events.Event {
	return &OrderAddProduct{
		BaseEvent:  events.BaseEvent{}.Create(),
		Product:    product,
		OrderPrice: orderPrice,
	}
}
func (o *OrderAddProduct) Jsonify() ([]byte, error) {
	return events.Jsonify(o)
}
