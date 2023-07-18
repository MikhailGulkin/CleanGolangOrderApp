package dao

import (
	order "github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/order/entities"
	"github.com/google/uuid"
)

type OrderDAO interface {
	GetProductsByIDs([]uuid.UUID) ([]order.OrderProduct, error)
	GetClientByID(uuid.UUID) (order.OrderClient, error)
	GetAddressByID(uuid.UUID) (order.OrderAddress, error)
}
type OrderSagaDAO interface {
	UpdateOrderSagaStatus(uuid.UUID, string, interface{}) error
	DeleteOrderCascade(uuid.UUID, interface{}) error
}
