package query

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/common/interfaces/persistence/query"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/dto"
)

type GetAllOrderQuery struct {
	query.BaseListQueryParams
}

type GetAllOrders interface {
	Get(query GetAllOrderQuery) (dto.Orders, error)
}
