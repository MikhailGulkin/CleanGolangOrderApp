package models

type JSONB map[string]interface{}
type Outbox struct {
	Base
	Exchange  string
	Route     string
	Payload   JSONB `gorm:"type:jsonb"`
	Processed bool
}
