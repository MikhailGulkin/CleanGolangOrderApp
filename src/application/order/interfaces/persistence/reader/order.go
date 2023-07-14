package reader

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/filters"
	"github.com/google/uuid"
)

type OrderReader interface {
	GetAllOrders(filters filters.GetAllOrdersFilters) ([]dto.Order, error)
	GetOrderByID(id uuid.UUID) (dto.Order, error)
}
