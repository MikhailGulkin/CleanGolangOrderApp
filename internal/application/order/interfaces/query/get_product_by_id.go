package query

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/dto"
	"github.com/google/uuid"
)

type GetOrderByIDQuery struct {
	ID uuid.UUID `json:"ID"`
}
type GetOrderByID interface {
	Get(query GetOrderByIDQuery) (dto.Order, error)
}
