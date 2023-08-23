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
