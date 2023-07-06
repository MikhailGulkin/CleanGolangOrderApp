package interactors

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/command/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/dao"
	query "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/query/product"
	p "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/query/product"
	"go.uber.org/fx"
)

func NewCreateProductImpl(dao dao.ProductDAO, uow persistence.UoW) command.CreateProduct {
	return &product.CreateProductImpl{
		ProductDAO: dao,
		UoW:        uow,
	}
}
func NewGetALlProductsImp(dao dao.ProductReader) query.GetAllProducts {
	return &p.GetAllProductsImpl{
		DAO: dao,
	}
}

var Module = fx.Provide(
	NewCreateProductImpl,
	NewGetALlProductsImp,
)
