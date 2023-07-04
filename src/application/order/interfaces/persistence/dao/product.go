package dao

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/product"
)

type ProductDAO interface {
	Create(product product.Product, tx interface{}) error
}
