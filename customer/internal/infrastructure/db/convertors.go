package db

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/common"
)

func ConvertAggregateToEntity(customer common.Aggregate) Entity {
	return Entity{
		EntityID:      customer.GetID(),
		EventType:     string(customer.GetType()),
		EntityVersion: customer.GetVersion(),
	}
}

func ConvertDomainEventToEventModel(event common.Event) Event {
	return Event{
		EventID:    event.GetEventID(),
		EventType:  event.GetEventType(),
		EventData:  event.GetData(),
		EntityType: string(event.GetAggregateType()),
		EntityID:   event.GetAggregateID(),
		CreatedAt:  event.GetTimeStamp(),
	}
}
func ConvertDomainEventToOutboxMessage(event common.Event) OutboxMessage {
	return OutboxMessage{
		Exchange:    "Customer",
		Route:       "Customer",
		Payload:     event.GetData(),
		AggregateID: event.GetAggregateID(),
	}
}

func ConvertEventToDomainEvent(event Event) common.Event {
	return common.Event{
		EventID:       event.EventID,
		EventType:     event.EventType,
		Data:          event.EventData,
		AggregateType: common.AggregateType(event.EntityType),
		AggregateID:   event.EntityID,
		Timestamp:     event.CreatedAt,
	}
}
func ConvertEventsToDomainEvents(event []Event) []common.Event {
	events := make([]common.Event, 0, len(event))
	for _, e := range event {
		events = append(events, ConvertEventToDomainEvent(e))
	}
	return events
}
