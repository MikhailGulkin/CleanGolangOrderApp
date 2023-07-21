package mediator

import (
	commandAddress "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/address/interfaces/command"
	commandOrder "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/command"
	queryOrder "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/order/interfaces/query"
	commandProduct "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/product/interfaces/command"
	queryProduct "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/product/interfaces/query"
	commandUser "github.com/MikhailGulkin/simpleGoOrderApp/internal/application/user/interfaces/command"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	queryProduct.GetAllProducts
	queryProduct.GetProductByName
	commandProduct.UpdateProductName
	commandProduct.CreateProduct
	commandOrder.CreateOrder
	commandOrder.DeleteOrder
	queryOrder.GetAllOrders
	queryOrder.GetAllOrdersByUserID
	queryOrder.GetOrderByID
	commandAddress.CreateAddress
	commandUser.CreateUser
}
