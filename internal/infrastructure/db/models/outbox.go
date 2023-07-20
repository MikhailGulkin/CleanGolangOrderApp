package models

import "github.com/google/uuid"

type Outbox struct {
	Base
	Exchange    string
	Route       string
	Payload     string `gorm:"type:jsonb;"`
	AggregateID uuid.UUID
	EventStatus int `gorm:"default:1"`
}
type OutboxEventStatus uint

const (
	Undefined OutboxEventStatus = iota
	Awaiting                    = 1
	Processed                   = 2
	Sagas                       = 3
	Rejected                    = 4
)
