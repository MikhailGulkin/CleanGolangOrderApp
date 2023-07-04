package persistence

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/product"
)

type ProductDAO interface {
	Create(product.Product) error
}
