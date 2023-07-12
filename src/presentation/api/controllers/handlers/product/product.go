package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/query"
)

func NewProductHandler(
	createProduct command.CreateProduct,
	getAllProducts query.GetAllProducts,
	getProductByName query.GetProductByName,
	updateProductByName command.UpdateProductName,
) Handler {
	return Handler{
		createProduct:     createProduct,
		getAllProducts:    getAllProducts,
		getProductByName:  getProductByName,
		updateProductName: updateProductByName,
	}
}
