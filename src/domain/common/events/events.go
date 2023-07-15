package events

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type Event interface {
	Jsonify() ([]byte, error)
}
type BaseEvent struct {
	EventID        uuid.UUID `json:"eventID"`
	EventTimeStamp time.Time `json:"eventTimeStamp"`
}

func (BaseEvent) Create() BaseEvent {
	return BaseEvent{
		EventID:        uuid.New(),
		EventTimeStamp: time.Now(),
	}
}
func Jsonify(v any) ([]byte, error) {
	bin, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return bin, nil
}
func (o *BaseEvent) Jsonify() ([]byte, error) {
	return Jsonify(o)
}
