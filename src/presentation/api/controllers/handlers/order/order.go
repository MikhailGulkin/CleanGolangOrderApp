package order

import c "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"

func NewOrderHandler(
	createOrder c.CreateOrder,
) Handler {
	return Handler{
		createOrder: createOrder,
	}
}
