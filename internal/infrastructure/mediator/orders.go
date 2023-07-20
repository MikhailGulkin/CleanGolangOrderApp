package mediator

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/query"
)

type CreateOrderCommandHandler struct {
	command.CreateOrder
}

func (c *CreateOrderCommandHandler) Handle(cmd interface{}) error {
	return c.Create(cmd.(command.CreateOrderCommand))
}

type GetAllOrdersQueryHandler struct {
	query.GetAllOrders
}

func (c *GetAllOrdersQueryHandler) Handle(q interface{}) (interface{}, error) {
	return c.Get(q.(query.GetAllOrderQuery))
}

type GetAllOrdersByUserIDQueryHandler struct {
	query.GetAllOrdersByUserID
}

func (c *GetAllOrdersByUserIDQueryHandler) Handle(q interface{}) (interface{}, error) {
	return c.Get(q.(query.GetAllOrderByUserIDQuery))
}

type GetOrdersByIDHandler struct {
	query.GetOrderByID
}

func (c *GetOrdersByIDHandler) Handle(q interface{}) (interface{}, error) {
	return c.Get(q.(query.GetOrderByIDQuery))
}
