package repo

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/common/id"
)

type ProductRepo interface {
	AcquireProductsByIDs(productIDs []id.ID) ([]product.Product, error)
	AcquireProductByID(productID id.ID) (product.Product, error)
	AddProduct(product product.Product, tx interface{}) error
	UpdateProduct(product product.Product, tx interface{}) error
}
