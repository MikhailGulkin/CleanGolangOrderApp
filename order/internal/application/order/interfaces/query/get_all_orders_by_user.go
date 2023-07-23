package query

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/common/interfaces/persistence/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/dto"
	"github.com/google/uuid"
)

type GetAllOrderByUserIDQuery struct {
	UserID uuid.UUID `json:"userID"`
	query.BaseListQueryParams
}

type GetAllOrdersByUserID interface {
	Get(query GetAllOrderByUserIDQuery) (dto.Orders, error)
}
