package query

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/common/interfaces/persistence/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/product/dto"
)

type GetAllProductsQuery struct {
	query.BaseListQueryParams
}

type GetAllProducts interface {
	Get(query GetAllProductsQuery) (dto.Products, error)
}