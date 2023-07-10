package repo

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/aggregate/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/vo"
)

type ProductRepo interface {
	AcquireProductsByIDs(productIDs []vo.ProductID) ([]product.Product, error)
	AcquireProductByID(productID vo.ProductID) (product.Product, error)
	Create(product product.Product, tx interface{}) error
	UpdateProduct(product product.Product, tx interface{}) error
}
