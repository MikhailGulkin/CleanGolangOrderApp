package repo

import (
	order "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/order"
)

type OrderRepo interface {
	AcquireLastOrder() (order.Order, error)
	AddOrder(order order.Order, tx interface{}) error
}
