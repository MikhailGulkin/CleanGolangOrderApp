package saga

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/consts"
	"github.com/google/uuid"
)

type Message struct {
	OrderID   uuid.UUID         `json:"orderID"`
	OrderType consts.SagaStatus `json:"orderType"`
}

type CreateOrder interface {
	CheckStatus(message Message)
}
