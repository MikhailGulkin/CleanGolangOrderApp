package user

import (
	"github.com/google/uuid"
	"time"
)

type UserOrder struct {
	OrderID    uuid.UUID
	TotalPrice float64
	Date       time.Time
}
