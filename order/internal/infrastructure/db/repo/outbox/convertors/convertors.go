package convertors

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/domain/common/events"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/db/models"
	"github.com/google/uuid"
	"reflect"
)

type PayloadEnhanced struct {
	payload       string
	reflect       reflect.Type
	eventUniqueID uuid.UUID
}
type EventToOutbox struct {
	payloads []PayloadEnhanced
}

func (e EventToOutbox) Create(events []events.Event) (EventToOutbox, error) {
	payloads := make([]PayloadEnhanced, len(events))
	for _, event := range events {
		binary, binaryErr := event.Bytes()
		if binaryErr != nil {
			return EventToOutbox{}, binaryErr
		}

		payloads = append(payloads, PayloadEnhanced{payload: string(binary), reflect: reflect.TypeOf(event), eventUniqueID: event.UniqueAggregateID()})
	}
	return EventToOutbox{payloads: payloads}, nil
}
func (e *EventToOutbox) Convert() []models.Outbox {
	var modelsOutbox []models.Outbox
	for _, payload := range e.payloads {
		if OrdersEventsHandler(&modelsOutbox, payload) {
			continue
		}
	}
	return modelsOutbox
}
