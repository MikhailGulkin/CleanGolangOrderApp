package events

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Event interface {
	Bytes() ([]byte, error)
	UniqueAggregateID() uuid.UUID
}
type BaseEvent struct {
	EventID        uuid.UUID `json:"eventID,omitempty"`
	EventType      string    `json:"eventType"`
	EventTimeStamp time.Time `json:"eventTimeStamp,omitempty"`
}

func (BaseEvent) Create(eventType string) BaseEvent {
	return BaseEvent{
		EventType:      eventType,
		EventID:        uuid.New(),
		EventTimeStamp: time.Now(),
	}
}
func Bytes(v any) ([]byte, error) {
	bin, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bin, nil
}
func (o *BaseEvent) Bytes() ([]byte, error) {
	return Bytes(o)
}
func (o *BaseEvent) UniqueAggregateID() uuid.UUID {
	return o.EventID
}
