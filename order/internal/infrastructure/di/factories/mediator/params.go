package mediator

import (
	commandOrder "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/command"
	queryOrder "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/order/interfaces/query"
	commandProduct "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/product/interfaces/command"
	queryProduct "github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/product/interfaces/query"
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
}
