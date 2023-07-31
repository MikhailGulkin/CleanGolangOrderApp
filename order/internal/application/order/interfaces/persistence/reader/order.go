package reader

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/dto"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/persistence/filters"
	"github.com/google/uuid"
)

type OrderCacheReader interface {
	GetAllOrders(filters filters.GetAllOrdersFilters) ([]dto.Order, error)
	GetAllOrdersByUserID(userID uuid.UUID, filters filters.GetAllOrdersByUserIDFilters) ([]dto.Order, error)
	GetOrderByID(uuid.UUID) (dto.Order, error)
}
