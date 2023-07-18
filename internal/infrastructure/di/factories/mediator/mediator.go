package mediator

import (
	commandAddress "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/address/interfaces/command"
	commandOrder "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/command"
	queryOrder "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/query"
	commandProduct "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/product/interfaces/command"
	queryProduct "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/product/interfaces/query"
	commandUser "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/user/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/infrastructure/mediator"
	"go.uber.org/fx"
)

func NewMediator(
	getAllProducts queryProduct.GetAllProducts,
	getProductByName queryProduct.GetProductByName,
	updateProductName commandProduct.UpdateProductName,
	createProduct commandProduct.CreateProduct,
	createOrder commandOrder.CreateOrder,
	getAllOrders queryOrder.GetAllOrders,
	getOrderByID queryOrder.GetOrderByID,
	createAddress commandAddress.CreateAddress,
	createUser commandUser.CreateUser,
) mediator.Mediator {
	m := mediator.MediatorImpl{}.Create()
	m.RegisterCommandHandler(commandProduct.CreateProductCommand{}, &mediator.CreateProductCommandHandler{CreateProduct: createProduct})
	m.RegisterCommandHandler(commandProduct.UpdateProductNameCommand{}, &mediator.UpdateProductCommandHandler{UpdateProductName: updateProductName})
	m.RegisterQueryHandler(queryProduct.GetAllProductsQuery{}, &mediator.GetAllProductsQueryHandler{GetAllProducts: getAllProducts})
	m.RegisterQueryHandler(queryProduct.GetProductByNameQuery{}, &mediator.GetProductByNameQueryHandler{GetProductByName: getProductByName})

	m.RegisterCommandHandler(commandOrder.CreateOrderCommand{}, &mediator.CreateOrderCommandHandler{CreateOrder: createOrder})
	m.RegisterQueryHandler(queryOrder.GetOrderByIDQuery{}, &mediator.GetOrdersByIDHandler{GetOrderByID: getOrderByID})
	m.RegisterQueryHandler(queryOrder.GetAllOrderQuery{}, &mediator.GetAllOrdersQueryHandler{GetAllOrders: getAllOrders})

	m.RegisterCommandHandler(commandUser.CreateUserCommand{}, &mediator.CreateUserCommandHandler{CreateUser: createUser})

	m.RegisterCommandHandler(commandAddress.CreateAddressCommand{}, &mediator.CreateAddressCommandHandler{CreateAddress: createAddress})

	return m
}

var Module = fx.Provide(
	NewMediator,
)
