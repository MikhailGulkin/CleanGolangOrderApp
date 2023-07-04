package dto

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID           uuid.UUID
	Name         string
	Price        float64
	Discount     int32
	Quantity     int32
	Description  string
	Availability bool
	CreatedAt    time.Time
}
