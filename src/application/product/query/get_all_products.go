package query

import (
	base "github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/dto"
	baseFilters "github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence/filters"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/persistence/filters"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/persistence/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/query"
)

type GetAllProductsImpl struct {
	DAO reader.ProductReader
	query.GetAllProducts
}

func (interactor *GetAllProductsImpl) Get(q query.GetAllProductsQuery) (dto.Products, error) {
	products, err := interactor.DAO.GetAllProducts(
		filters.GetAllProductsFilters{BaseFilters: baseFilters.BaseFilters{Limit: q.Limit, Offset: q.Offset, Order: q.Order}},
	)
	if err != nil {
		return dto.Products{}, err
	}
	return dto.Products{Products: products, BaseSequence: base.BaseSequence{Limit: q.Limit, Offset: q.Offset, Order: q.Order, Count: uint(len(products))}}, nil
}
