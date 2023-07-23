package mediator

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/order/internal/application/product/interfaces/query"
)

type CreateProductCommandHandler struct {
	command.CreateProduct
}

func (c *CreateProductCommandHandler) Handle(cmd interface{}) error {
	return c.Create(cmd.(command.CreateProductCommand))
}

type UpdateProductCommandHandler struct {
	command.UpdateProductName
}

func (c *UpdateProductCommandHandler) Handle(cmd interface{}) error {
	return c.Update(cmd.(command.UpdateProductNameCommand))
}

type GetAllProductsQueryHandler struct {
	query.GetAllProducts
}

func (c *GetAllProductsQueryHandler) Handle(q interface{}) (interface{}, error) {
	return c.GetAllProducts.Get(q.(query.GetAllProductsQuery))
}

type GetProductByNameQueryHandler struct {
	query.GetProductByName
}

func (c *GetProductByNameQueryHandler) Handle(q interface{}) (interface{}, error) {
	return c.GetProductByName.Get(q.(query.GetProductByNameQuery))
}
