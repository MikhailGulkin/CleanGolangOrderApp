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
