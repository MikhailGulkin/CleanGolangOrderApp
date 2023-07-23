package query

import (
	"fmt"
	base "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/common/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/common/interfaces/logger"
	baseFilters "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/common/interfaces/persistence/filters"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/persistence/filters"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/persistence/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/query"
)

type GetAllProductsImpl struct {
	DAO reader.ProductReader
	query.GetAllProducts
	logger.Logger
}

func (interactor *GetAllProductsImpl) Get(q query.GetAllProductsQuery) (dto.Products, error) {
	products, err := interactor.DAO.GetAllProducts(
		filters.GetAllProductsFilters{BaseFilters: baseFilters.BaseFilters{Limit: q.Limit, Offset: q.Offset, Order: q.Order}},
	)
	if err != nil {
		return dto.Products{}, err
	}
	length := len(products)
	interactor.Logger.Info(
		fmt.Sprintf("Get all products, count: %d, order: %s, limit: %d, offset: %d",
			length, q.Order, q.Limit, q.Offset,
		))
	return dto.Products{
		Products:     products,
		BaseSequence: base.BaseSequence{Limit: q.Limit, Offset: q.Offset, Order: q.Order, Count: uint(length)},
	}, nil
}
