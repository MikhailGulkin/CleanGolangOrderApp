package handlers

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/query"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewProductHandler),
)

func NewProductHandler(
	createProduct command.CreateProduct,
	getAllProducts query.GetAllProducts,
	getProductByName query.GetProductByName,
	updateProductByName command.UpdateProductName,
) ProductHandler {
	return ProductHandler{
		createProduct:     createProduct,
		getAllProducts:    getAllProducts,
		getProductByName:  getProductByName,
		updateProductName: updateProductByName,
	}
}
