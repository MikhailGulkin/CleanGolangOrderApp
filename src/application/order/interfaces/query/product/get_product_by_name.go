package product

import "github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/dto"

type GetProductByNameQuery struct {
	Name string
}
type GetProductByName interface {
	Get(query GetProductByNameQuery) (dto.Product, error)
}
