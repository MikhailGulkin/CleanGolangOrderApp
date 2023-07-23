package mediator

import (
	commandAddress "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/address/interfaces/command"
	commandOrder "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/interfaces/command"
	queryOrder "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/order/interfaces/query"
	commandProduct "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/command"
	queryProduct "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/query"
	commandUser "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/user/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/mediator"
	"go.uber.org/fx"
)

func NewMediator(p Params) mediator.Mediator {
	m := mediator.MediatorImpl{}.Create()
	m.RegisterCommandHandler(commandProduct.CreateProductCommand{}, &mediator.CreateProductCommandHandler{CreateProduct: p.CreateProduct})
	m.RegisterCommandHandler(commandProduct.UpdateProductNameCommand{}, &mediator.UpdateProductCommandHandler{UpdateProductName: p.UpdateProductName})
	m.RegisterQueryHandler(queryProduct.GetAllProductsQuery{}, &mediator.GetAllProductsQueryHandler{GetAllProducts: p.GetAllProducts})
	m.RegisterQueryHandler(queryProduct.GetProductByNameQuery{}, &mediator.GetProductByNameQueryHandler{GetProductByName: p.GetProductByName})

	m.RegisterCommandHandler(commandOrder.CreateOrderCommand{}, &mediator.CreateOrderCommandHandler{CreateOrder: p.CreateOrder})
	m.RegisterCommandHandler(commandOrder.DeleteOrderCommand{}, &mediator.DeleteOrderCommandHandler{DeleteOrder: p.DeleteOrder})
	m.RegisterQueryHandler(queryOrder.GetOrderByIDQuery{}, &mediator.GetOrdersByIDHandler{GetOrderByID: p.GetOrderByID})
	m.RegisterQueryHandler(queryOrder.GetAllOrderQuery{}, &mediator.GetAllOrdersQueryHandler{GetAllOrders: p.GetAllOrders})
	m.RegisterQueryHandler(queryOrder.GetAllOrderByUserIDQuery{}, &mediator.GetAllOrdersByUserIDQueryHandler{GetAllOrdersByUserID: p.GetAllOrdersByUserID})

	m.RegisterCommandHandler(commandUser.CreateUserCommand{}, &mediator.CreateUserCommandHandler{CreateUser: p.CreateUser})

	m.RegisterCommandHandler(commandAddress.CreateAddressCommand{}, &mediator.CreateAddressCommandHandler{CreateAddress: p.CreateAddress})

	return m
}

var Module = fx.Provide(
	NewMediator,
)
