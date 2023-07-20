package dao

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/dto"
	order "github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/order/entities"
	"github.com/google/uuid"
)

type OrderDAO interface {
	GetProductsByIDs([]uuid.UUID) ([]order.OrderProduct, error)
	GetClientByID(uuid.UUID) (order.OrderClient, error)
	GetAddressByID(uuid.UUID) (order.OrderAddress, error)
}
type OrderSagaDAO interface {
	CheckSagaCompletion(uuid.UUID) (bool, error)
	UpdateOrderSagaStatus(uuid.UUID, string, interface{}) error
	DeleteOrderCascade(uuid.UUID, interface{}) error
}
type OrderCacheDAO interface {
	GetOrder(uuid.UUID) dto.Order
	SaveOrder(dto.Order) error
}
