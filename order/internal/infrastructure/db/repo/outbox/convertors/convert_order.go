package convertors

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/consts/outbox"
	o "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/events"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/infrastructure/db/models"
	"reflect"
)

func OrdersEventsHandler(outboxes *[]models.Outbox, payload PayloadEnhanced) bool {
	if payload.reflect == reflect.TypeOf(new(o.OrderCreated)) {
		*outboxes = append(*outboxes, models.Outbox{
			Exchange:    "Orders",
			Route:       "Order.Create",
			Payload:     payload.payload,
			EventStatus: outbox.Sagas,
			AggregateID: payload.eventUniqueID,
		})
		return true
	}
	if payload.reflect == reflect.TypeOf(new(o.OrderAddProduct)) {
		*outboxes = append(*outboxes, models.Outbox{
			Exchange:    "Orders",
			Route:       "Order.AddProduct",
			Payload:     payload.payload,
			EventStatus: outbox.Sagas,
			AggregateID: payload.eventUniqueID,
		})
		return true
	}
	if payload.reflect == reflect.TypeOf(new(o.OrderDeleted)) {
		*outboxes = append(*outboxes, models.Outbox{
			Exchange:    "Orders",
			Route:       "Order.Delete",
			Payload:     payload.payload,
			AggregateID: payload.eventUniqueID,
		})
		return true
	}
	if payload.reflect == reflect.TypeOf(new(o.OrderCreateSaga)) {
		*outboxes = append(*outboxes, models.Outbox{
			Exchange:    "Orders",
			Route:       "Order.Saga.Create",
			Payload:     payload.payload,
			AggregateID: payload.eventUniqueID,
		})
		return true
	}

	return false
}
