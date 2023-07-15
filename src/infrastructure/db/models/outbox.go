package models

type JSONB []interface{}

type Outbox struct {
	Base
	Exchange    string
	Route       string
	Payload     string `gorm:"type:jsonb;"`
	IsProcessed bool   `gorm:"default:false"`
}
