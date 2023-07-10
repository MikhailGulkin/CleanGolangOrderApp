package interactors

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/command"
	command2 "github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/persistence/reader"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/persistence/repo"
	query2 "github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/interfaces/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/product/query"
	"go.uber.org/fx"
)

func NewCreateProduct(dao repo.ProductRepo, uow persistence.UoW) command2.CreateProduct {
	return &command.CreateProductImpl{
		ProductRepo: dao,
		UoW:         uow,
	}
}
func NewUpdateProductName(dao repo.ProductRepo, uow persistence.UoW) command2.UpdateProductName {
	return &command.UpdateProductNameImpl{
		ProductRepo: dao,
		UoW:         uow,
	}
}

func NewGetALlProducts(dao reader.ProductReader) query2.GetAllProducts {
	return &query.GetAllProductsImpl{
		DAO: dao,
	}
}
func NewGetProductByName(dao reader.ProductReader) query2.GetProductByName {
	return &query.GetProductByNameImpl{
		DAO: dao,
	}
}

var Module = fx.Provide(
	NewCreateProduct,
	NewGetALlProducts,
	NewUpdateProductName,
	NewGetProductByName,
)
