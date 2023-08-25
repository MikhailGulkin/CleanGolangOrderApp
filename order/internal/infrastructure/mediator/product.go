package mediator

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/product/interfaces/command"
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/application/product/interfaces/query"
)

type CreateProductCommandHandler struct {
	command.CreateProduct
}

func (c *CreateProductCommandHandler) Handle(cmd interface{}) (interface{}, error) {
	return nil, c.Create(cmd.(command.CreateProductCommand))
}

type UpdateProductCommandHandler struct {
	command.UpdateProductName
}

func (c *UpdateProductCommandHandler) Handle(cmd interface{}) (interface{}, error) {
	return nil, c.Update(cmd.(command.UpdateProductNameCommand))
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
