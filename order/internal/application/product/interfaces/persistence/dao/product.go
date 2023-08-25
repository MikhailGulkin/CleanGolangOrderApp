package dao

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/common/id"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/domain/product/entities"
)

type ProductDAO interface {
	GetProductByID(productID id.ID) (entities.Product, error)
	CreateProduct(product entities.Product, tx interface{}) error
	UpdateProduct(product entities.Product, tx interface{}) error
}
