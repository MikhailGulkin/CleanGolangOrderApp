package repo

import (
	order "github.com/MikhailGulkin/simpleGoOrderApp/internal/domain/order/aggregate"
)

type OrderRepo interface {
	AcquireLastOrder() (order.Order, error)
	AddOrder(order order.Order, tx interface{}) error
}
