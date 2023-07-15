package convertors

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/common/events"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/db/models"
	"reflect"
)

type PayloadEnhanced struct {
	payload string
	reflect reflect.Type
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

		payloads = append(payloads, PayloadEnhanced{payload: string(binary), reflect: reflect.TypeOf(event)})
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
