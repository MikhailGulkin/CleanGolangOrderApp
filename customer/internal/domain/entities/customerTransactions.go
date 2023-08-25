package entities

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/consts"
	"github.com/MikhailGulkin/simpleGoOrderApp/customer/internal/domain/vo"
	"github.com/google/uuid"
	"time"
)

type CustomerTransactions struct {
	TransactionId   uuid.UUID
	TransactionType consts.TransactionType
	TransactionDate time.Time
	TransactionSum  vo.Money
	Comment         string
	OrderID         *uuid.UUID
}
