package order

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/query"
)

func NewOrderHandler(
	createOrder command.CreateOrder,
	getAllOrders query.GetAllOrders,
	getOrderByID query.GetOrderByID,
) Handler {
	return Handler{
		createOrder:  createOrder,
		getAllOrders: getAllOrders,
		getOrderByID: getOrderByID,
	}
}
