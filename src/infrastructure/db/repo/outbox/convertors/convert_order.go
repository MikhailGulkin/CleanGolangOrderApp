package convertors

import (
	o "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/events"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	"reflect"
)

func OrdersEventsHandler(outboxes *[]models.Outbox, payload PayloadEnhanced) bool {
	if payload.reflect == reflect.TypeOf(new(o.OrderCreated)) {
		*outboxes = append(*outboxes, models.Outbox{
			Exchange: "Order",
			Route:    "order_create",
			Payload:  payload.payload,
		})
		return true
	}
	if payload.reflect == reflect.TypeOf(new(o.OrderAddProduct)) {
		*outboxes = append(*outboxes, models.Outbox{
			Exchange: "Order",
			Route:    "order_add_product",
			Payload:  payload.payload,
		})
		return true
	}
	return false
}
