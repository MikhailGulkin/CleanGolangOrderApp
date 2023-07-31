package mediator

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/command"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/query"
)

type CreateOrderCommandHandler struct {
	command.CreateOrder
}

func (c *CreateOrderCommandHandler) Handle(cmd interface{}) (interface{}, error) {
	return nil, c.Create(cmd.(command.CreateOrderCommand))
}

type DeleteOrderCommandHandler struct {
	command.DeleteOrder
}

func (c *DeleteOrderCommandHandler) Handle(cmd interface{}) (interface{}, error) {
	return nil, c.Delete(cmd.(command.DeleteOrderCommand))
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
