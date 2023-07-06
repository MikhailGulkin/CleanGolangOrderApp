package dao

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/filters"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/product"
)

type ProductDAO interface {
	Create(product product.Product, tx interface{}) error
}
type ProductReader interface {
	GetAllProducts(filters filters.BaseFilters) ([]dto.Product, error)
	GetProductByName(name string) (dto.Product, error)
}
