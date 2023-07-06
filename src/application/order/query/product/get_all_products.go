package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/dto"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/filters"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/query/product"
)

type GetAllProductsImpl struct {
	DAO dao.ProductReader
	product.GetAllProducts
}

func (interactor *GetAllProductsImpl) Get(query product.GetAllProductsQuery) (dto.Products, error) {
	products, err := interactor.DAO.GetAllProducts(
		filters.GetAllProductsFilters{}.Create(query.Limit, query.Offset, query.Order),
	)
	if err != nil {
		return dto.Products{}, err
	}
	return dto.Products{Products: products, Limit: query.Limit, Offset: query.Offset, Order: query.Order}, nil
}
