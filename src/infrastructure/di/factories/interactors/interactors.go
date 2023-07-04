package interactors

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/command/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/dao"
	"go.uber.org/fx"
)

func NewCreateProductImpl(dao dao.ProductDAO, uow persistence.UoW) command.CreateProduct {
	interactorProduct := product.CreateProductImpl{
		ProductDAO: dao,
		UoW:        uow,
	}
	return &interactorProduct
}

var Module = fx.Provide(
	NewCreateProductImpl,
)
