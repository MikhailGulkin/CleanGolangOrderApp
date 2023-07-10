package repo

import (
	domain "github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/order"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"
)

type OrderRepo interface {
	AcquireLastOrderByID(orderID vo.OrderID) (domain.Order, error)
	Add(order *domain.Order, tx interface{}) error
}
