package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/command"
	c "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/persistence/reader"
	q "github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/query"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/infrastructure/logger"
	"go.uber.org/fx"
)

func NewCreateProduct(dao dao.ProductDAO, uow persistence.UoW) c.CreateProduct {
	return &command.CreateProductImpl{
		ProductDAO: dao,
		UoW:        uow,
	}
}
func NewUpdateProductName(dao dao.ProductDAO, uow persistence.UoW) c.UpdateProductName {
	return &command.UpdateProductNameImpl{
		ProductDAO: dao,
		UoW:        uow,
	}
}

func NewGetALlProducts(dao reader.ProductReader, logger logger.Logger) q.GetAllProducts {
	return &query.GetAllProductsImpl{
		DAO:    dao,
		Logger: logger,
	}
}
func NewGetProductByName(dao reader.ProductReader) q.GetProductByName {
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
