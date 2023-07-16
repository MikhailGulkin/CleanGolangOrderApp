package dao

import (
	order "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/order/entities"
	"github.com/google/uuid"
)

type OrderDAO interface {
	GetProductsByIDs([]uuid.UUID) ([]order.OrderProduct, error)
	GetClientByID(uuid.UUID) (order.OrderClient, error)
	GetAddressByID(uuid.UUID) (order.OrderAddress, error)
}
