package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/product"
)

type CreateProductImpl struct {
	persistence.ProductDAO
	command.CreateProduct
}

func (interactor *CreateProductImpl) Create(command command.CreateProductCommand) error {
	productEntity, err := product.Product{}.Create(
		command.Price,
		command.Discount,
		command.Description,
		command.Name,
	)
	if err != nil {
		return err
	}
	err = interactor.ProductDAO.Create(productEntity)
	if err != nil {
		return err
	}
	return nil
}
