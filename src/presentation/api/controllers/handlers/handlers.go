package handlers

import (
	interfaces "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/query/product"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewProductHandler),
)

func NewProductHandler(
	createProduct interfaces.CreateProduct,
	getAllProducts product.GetAllProducts,
	getProductByName product.GetProductByName,
	updateProductByName interfaces.UpdateProductName,
) ProductHandler {
	return ProductHandler{
		createProduct:     createProduct,
		getAllProducts:    getAllProducts,
		getProductByName:  getProductByName,
		updateProductName: updateProductByName,
	}
}
