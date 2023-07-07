package dao

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/filters"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/value_object"
)

type ProductDAO interface {
	AcquireProductByID(productID value_object.ProductID) (product.Product, error)
	Create(product product.Product, tx interface{}) error
	UpdateProduct(product product.Product, tx interface{}) error
}
type ProductReader interface {
	GetAllProducts(filters filters.BaseFilters) ([]dto.Product, error)
	GetProductByName(name string) (dto.Product, error)
}
