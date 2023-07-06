package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/query"
)

type GetAllProductsQuery struct {
	query.BaseListQueryParams
}

type GetAllProducts interface {
	Get(query GetAllProductsQuery) (dto.Products, error)
}
