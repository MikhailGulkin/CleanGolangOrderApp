package reader

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/interfaces/persistence/filters"
	"github.com/google/uuid"
)

type OrderReader interface {
	GetAllOrders(filters filters.GetAllOrdersFilters) ([]dto.Order, error)
}
type OrderCacheReader interface {
	GetAllOrdersByUserID(userID uuid.UUID, filters filters.GetAllOrdersByUserIDFilters) ([]dto.Order, error)
	GetOrderByID(uuid.UUID) (dto.Order, error)
}
