package reader

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/filters"
)

type ProductReader interface {
	GetAllProducts(filters filters.BaseFilters) ([]dto.Product, error)
	GetProductByName(name string) (dto.Product, error)
}
