package entities

import (
	"github.com/google/uuid"
	"time"
)

type CustomerOrder struct {
	OrderID    uuid.UUID
	TotalPrice float64
	Date       time.Time
}
