package db

import (
	"time"
)

type Event struct {
	EventID    string
	EventType  string
	EventData  []byte
	EntityType string
	EntityID   string
	CreatedAt  time.Time
}
type Entity struct {
	EventType     string
	EntityID      string
	EntityVersion int64
	CreatedAt     time.Time
}
type OutboxMessage struct {
	Exchange    string
	Route       string
	Payload     []byte
	AggregateID string
	EventStatus int
}
