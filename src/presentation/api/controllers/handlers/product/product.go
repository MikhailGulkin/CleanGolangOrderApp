package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/query"
	commandbus "github.com/MikhailGulkin/simpleGoOrderApp/src/infrastructure/commandBus"
)

func NewProductHandler(
	bus commandbus.CommandBus,
	getAllProducts query.GetAllProducts,
	getProductByName query.GetProductByName,
	updateProductByName command.UpdateProductName,
) Handler {
	return Handler{
		bus:               bus,
		getAllProducts:    getAllProducts,
		getProductByName:  getProductByName,
		updateProductName: updateProductByName,
	}
}
