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

func NewCreateProduct(dao dao.ProductRepo, uow persistence.UoW) command.CreateProduct {
	return &product.CreateProductImpl{
		ProductRepo: dao,
		UoW:         uow,
	}
}
func NewUpdateProductName(dao dao.ProductRepo, uow persistence.UoW) command.UpdateProductName {
	return &product.UpdateProductNameImpl{
		ProductRepo: dao,
		UoW:         uow,
	}
}

func NewGetALlProducts(dao dao.ProductReader) query.GetAllProducts {
	return &p.GetAllProductsImpl{
		DAO: dao,
	}
}
func NewGetProductByName(dao dao.ProductReader) query.GetProductByName {
	return &p.GetProductByNameImpl{
		DAO: dao,
	}
}

var Module = fx.Provide(
	NewCreateProduct,
	NewGetALlProducts,
	NewUpdateProductName,
	NewGetProductByName,
)
