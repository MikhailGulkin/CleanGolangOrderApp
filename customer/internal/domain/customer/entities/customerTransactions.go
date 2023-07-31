package entities

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/user/consts"
	"github.com/google/uuid"
	"time"
)

type CustomerTransactions struct {
	TransactionId   uuid.UUID
	TransactionType consts.TransactionType
	TransactionDate time.Time
	TransactionSum  float64
	Comment         string
	OrderID         *uuid.UUID
}
