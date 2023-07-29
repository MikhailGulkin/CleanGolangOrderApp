package models

import "github.com/google/uuid"

type Customer struct {
	Base
	FirstName   string
	LastName    string
	PhoneNumber string
	Email       string
	Address     uuid.UUID `gorm:"type:uuid"`
}
