package common

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	EventID       string
	EventType     string
	Data          []byte
	Timestamp     time.Time
	AggregateType AggregateType
	AggregateID   string
	Version       int64
	Metadata      []byte
}

func NewBaseEvent(aggregate Aggregate, eventType string) Event {
	return Event{
		EventID:       uuid.New().String(),
		AggregateType: aggregate.GetType(),
		AggregateID:   aggregate.GetID(),
		Version:       aggregate.GetVersion(),
		EventType:     eventType,
		Timestamp:     time.Now().UTC(),
	}
}
