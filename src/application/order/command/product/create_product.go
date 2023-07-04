package product

import (
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/common/interfaces/persistence"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/command"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/application/order/interfaces/persistence/dao"
	"github.com/MikhailGulkin/simpleGoOrderApp/src/domain/entities/product"
)

type CreateProductImpl struct {
	dao.ProductDAO
	persistence.UoW
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
	interactor.StartTx()
	err = interactor.ProductDAO.Create(productEntity, interactor.GetTx())
	if err != nil {
		return err
	}
	interactor.Commit()

	return nil
}
