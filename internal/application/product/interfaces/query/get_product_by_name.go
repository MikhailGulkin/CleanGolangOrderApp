package query

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/internal/application/product/dto"
)

type GetProductByNameQuery struct {
	Name string
}
type GetProductByName interface {
	Get(query GetProductByNameQuery) (dto.Product, error)
}