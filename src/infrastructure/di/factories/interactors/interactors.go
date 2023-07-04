package interactors

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/command/product"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence"
	"go.uber.org/fx"
)

func NewCreateProductImpl(dao persistence.ProductDAO) command.CreateProduct {
	interactorProduct := product.CreateProductImpl{
		ProductDAO: dao,
	}
	return &interactorProduct
}

var Module = fx.Provide(
	NewCreateProductImpl,
)
