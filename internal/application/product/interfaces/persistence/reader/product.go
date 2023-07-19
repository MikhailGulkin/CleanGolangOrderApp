package reader

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/product/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/product/interfaces/persistence/filters"
)

type ProductReader interface {
	GetAllProducts(filters filters.GetAllProductsFilters) ([]dto.Product, error)
	GetProductByName(name string) (dto.Product, error)
}