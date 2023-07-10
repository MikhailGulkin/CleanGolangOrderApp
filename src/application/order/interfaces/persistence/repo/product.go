package repo

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/value_object"
)

type ProductRepo interface {
	AcquireProductByID(productID value_object.ProductID) (product.Product, error)
	Create(product product.Product, tx interface{}) error
	UpdateProduct(product product.Product, tx interface{}) error
}
