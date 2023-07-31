package repo

import (
	order "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/order/aggregate"
	"github.com/google/uuid"
)

type OrderRepo interface {
	AcquireLastOrder() (order.Order, error)
	AddOrder(order order.Order, tx interface{}) error
	AcquiredOrder(uuid uuid.UUID) (order.Order, error)
}
